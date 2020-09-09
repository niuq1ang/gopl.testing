package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/db"
	noticeModel "github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/notice"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/record"
	webhookModel "github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/webhook"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils"
	"github.com/cenkalti/backoff"
	"github.com/ngaut/log"
)

const (
	bufferSize = 1000
	// 约定在完全没有消息发送后的300s后发送心跳消息。若连续30分钟没有任何应答，则认为tunnel已经失效
	heartBeatSpan = 300
	maxWaitTimes  = 6 //  30min / 300s = 6
)

var backOff = &backoff.ExponentialBackOff{
	InitialInterval:     200 * time.Millisecond,
	RandomizationFactor: backoff.DefaultRandomizationFactor,
	Multiplier:          backoff.DefaultMultiplier,
	MaxInterval:         3 * time.Second,
	MaxElapsedTime:      9 * time.Second,
	Clock:               backoff.SystemClock,
}

type PayLoad struct {
	SequenceNo string                        `json:"sequence_no"`
	Message    []*noticeModel.MessagePayLoad `json:"messages"`
}

type Client struct {
	*webhookModel.WebHook
	signal              bool
	lock                sync.Mutex
	ConsecutiveFailures int64
	Notice              chan *noticeModel.Notice
	BackOff             backoff.BackOff
	StartWaitTime       time.Time
}

func NewClient(hook *webhookModel.WebHook) *Client {
	c := &Client{
		WebHook: hook,
		Notice:  make(chan *noticeModel.Notice, bufferSize),
		BackOff: &backoff.ExponentialBackOff{
			InitialInterval:     200 * time.Millisecond,
			RandomizationFactor: backoff.DefaultRandomizationFactor,
			Multiplier:          backoff.DefaultMultiplier,
			MaxInterval:         3 * time.Second,
			MaxElapsedTime:      9 * time.Second,
			Clock:               backoff.SystemClock,
		},
	}

	go c.Sender()
	return c
}

func (c *Client) AddNotice(n *noticeModel.Notice) {
	size := len(c.Notice)
	if size >= bufferSize {
		log.Warn("tunnel chan full")
		return
	}
	select {
	case c.Notice <- n: // SUCCESS
	default:
		log.Warn("no notice send")
	}
}

func (c *Client) Sender() {
	for {
		if c.Status == webhookModel.WebHookIsBatchUnable {
			break
		}
		var msgs []*noticeModel.MessagePayLoad
		var seqNos []int64

		timeNow := time.Now().Unix()
		for {
			length := len(c.Notice)
			if length == 0 {
				timeSpan := time.Now().Unix() - timeNow
				if timeSpan > 300 {
					// send heartbeat
					if err := PostHeartbeat(c.URL); err != nil {
						c.ConsecutiveFailures++
					} else {
						c.ConsecutiveFailures = 0
					}
					if c.ConsecutiveFailures > maxWaitTimes {
						c.Status = webhookModel.WebhookStatusUnavilable
						if err := UpdateCacheByWebhook(c.WebHook); err != nil {
							log.Error(errors.Trace(err))
						}
						break
					}
					timeNow = time.Now().Unix()
				}
				time.Sleep(time.Second)
				continue
			}
			n := <-c.Notice
			msgs = append(msgs, n.MessagePayLoad)
			seqNos = append(seqNos, n.SeqNo)
			if c.IsBatch != webhookModel.WebHookIsBatchEnable {
				break
			} else if len(msgs) == int(c.BatchCount) {
				break
			} else {
				continue
			}
		}

		sequenceNo := utils.Random62String(16)
		payload := PayLoad{
			Message:    msgs,
			SequenceNo: sequenceNo,
		}
		data, _ := json.Marshal(payload)

		if Post(c.BackOff, c.URL, sequenceNo, data) {
			if _, err := record.UpdateRecordsBySelNo(db.NoticeDBMap, record.RecordStatusSuccess, seqNos); err != nil {
				log.Error(err)
				continue
			}
			c.ConsecutiveFailures = 0
		} else {
			if _, err := record.UpdateRecordsBySelNo(db.NoticeDBMap, record.RecordStatusFailure, seqNos); err != nil {
				log.Error(err)
				continue
			}
			for {
				log.Warn("send failed", payload.Message)
				// send heartbeat
				time.Sleep(heartBeatSpan * time.Second)
				if err := PostHeartbeat(c.URL); err != nil {
					c.ConsecutiveFailures++
					if c.ConsecutiveFailures > maxWaitTimes {
						c.Status = webhookModel.WebhookStatusUnavilable
						UpdateCacheByWebhook(c.WebHook)
						break
					}
					continue
				} else {
					c.ConsecutiveFailures = 0
					break
				}
			}
		}
		timeNow = time.Now().Unix()
	}
	fmt.Println("sender heartbeat end", c.WebHook)
}

func Post(backOff backoff.BackOff, url, sequenceNo string, contents []byte) bool {
	client := new(http.Client)
	var resp *http.Response
	var err error
	op := func() error {
		var req *http.Request
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(contents))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err = client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		if resp.StatusCode == http.StatusOK {
			if string(body) == sequenceNo {
				return nil
			}
		}
		return errors.New(fmt.Sprintf("invalid return msg [%s]", string(body)))
	}

	err = backoff.Retry(op, backOff)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func PostHeartbeat(url string) error {
	client := new(http.Client)
	var resp *http.Response
	var err error
	sequenceNo := utils.Random62String(16)
	op := func() error {
		var req *http.Request
		requestBody := map[string]string{
			"sequence_no": sequenceNo,
		}
		byt, _ := json.Marshal(requestBody)
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(byt))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err = client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		if resp.StatusCode == http.StatusOK {
			if string(body) == sequenceNo {
				return nil
			}
		}
		return errors.New(fmt.Sprintf("invalid return msg [%s]", string(body)))
	}

	err = backoff.Retry(op, backOff)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}
