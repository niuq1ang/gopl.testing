package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/cenkalti/backoff"
)

type Client struct {
	BaseURL   string
	UserUUID  string
	UserToken string
	BackOff   backoff.BackOff
}

func NewClient(baseURL string) *Client {
	c := new(Client)
	c.UserUUID = userUUID
	c.UserToken = token
	c.BaseURL = baseURL
	c.BackOff = &backoff.ExponentialBackOff{
		InitialInterval:     200 * time.Millisecond,
		RandomizationFactor: backoff.DefaultRandomizationFactor,
		Multiplier:          backoff.DefaultMultiplier,
		MaxInterval:         5 * time.Second,
		MaxElapsedTime:      20 * time.Second,
		Clock:               backoff.SystemClock,
	}
	return c
}

func (c *Client) Post(api string, reqObj interface{}, respObj interface{}) (status int, err error) {
	return c.Request(http.MethodPost, api, reqObj, respObj)
}

func (c *Client) Request(method string, api string, reqObj interface{}, respObj interface{}) (status int, err error) {
	var requestBody []byte
	if reqObj != nil {
		requestBody, err = json.Marshal(reqObj)
		if err != nil {
			return
		}
		// fmt.Printf("req:\n%s\n", string(requestBody))
	}

	// u, err := url.Parse(c.BaseURL)
	// if err != nil {
	// 	return
	// }
	// u.Path = path.Join(u.Path, api)

	client := new(http.Client)
	var resp *http.Response
	op := func() error {
		var req *http.Request
		// req, err = http.NewRequest(method, u.String(), bytes.NewBuffer(requestBody))
		req, err = http.NewRequest(method, c.BaseURL+api, bytes.NewBuffer(requestBody))
		if err != nil {
			return err
		}
		if strings.Contains(api, "/team") {
			req.Header.Set("Ones-User-Id", c.UserUUID)
			req.Header.Set("Ones-Auth-Token", c.UserToken)
		}
		if reqObj != nil {
			req.Header.Set("Content-Type", "application/json")
		}
		resp, err = client.Do(req)
		if err != nil {
			return err
		}
		status = resp.StatusCode
		if resp.StatusCode != http.StatusOK {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			return fmt.Errorf("Call API failed: url=\"%s\", status=%d, body=\"%s\"", c.BaseURL+api, resp.StatusCode, body)
		}
		return nil
	}

	if c.BackOff == nil {
		err = op()
	} else {
		err = backoff.Retry(op, c.BackOff)
	}
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if respObj != nil && resp.Body != nil {
		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		// fmt.Printf("res:\n%s\n", string(body))
		err = json.Unmarshal(body, respObj)
		return
	}

	return
}
