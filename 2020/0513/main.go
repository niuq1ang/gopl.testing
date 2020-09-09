package main

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"github.com/ngaut/log"
)

func main() {
	err := InitLDAPAutoSync(true)
	if err != nil {
		panic(err)
	}

	go func() {
		for i := 0; i <= 10; i++ {
			time.Sleep(time.Second * 1)
			SyncLooper.ReLoad("12345678", i%2 == 0)
		}
	}()

	// go func() {
	// 	for i := 0; i <= 10; i++ {
	// 		time.Sleep(time.Second * 1)
	// 		SyncLooper.ReLoad("12345678", false)
	// 	}
	// }()

	time.Sleep(time.Minute * 10)
}

type ConfigStatus struct {
	OrgUUID string `db:"org_uuid"` // 团队UUID
	IsOpen  bool   `db:"is_open"`  // 是否开启
}

var SyncLooper *Looper

func InitLDAPAutoSync(isOpen bool) error {
	orgUUID := "12345678"
	SyncLooper = newSyncLooper()
	err := SyncLooper.ReLoad(orgUUID, isOpen)
	if err != nil {
		return err
	}
	return nil
}

type Looper struct {
	store map[string]struct{}
	lock  *sync.Mutex
}

func newSyncLooper() *Looper {
	return &Looper{lock: &sync.Mutex{}, store: map[string]struct{}{}}
}

func (looper *Looper) ReLoad(orgUUID string, isOpen bool) error {
	status := ConfigStatus{
		IsOpen:  isOpen,
		OrgUUID: orgUUID,
	}
	// fmt.Println("1", isOpen)
	if status.IsOpen {
		looper.lock.Lock()
		looper.store[orgUUID] = struct{}{}
		looper.lock.Unlock()

		looper.run(orgUUID)
	} else {
		looper.lock.Lock()
		delete(looper.store, orgUUID)
		looper.lock.Unlock()
	}
	// fmt.Println("66666")
	return nil
}

func (looper *Looper) run(orgUUID string) {
	// fmt.Printf("%+v\n", looper)
	// runFlag := orgUUID + "run"
	// if _, ok := looper.store[runFlag]; ok {
	// 	log.Info("LDAP sync is already exists ...")
	// 	return
	// }
	// fmt.Println("222")
	// sleepDuration := time.Second * 1

	// go func() {
	// 	defer func() {
	// 		delete(looper.store, runFlag)
	// 		fmt.Println("aaaaaaaaaaaaaaaaaaaaa")
	// 		if p := recover(); p != nil {
	// 			// delete(looper.store, runFlag)
	// 			fmt.Printf("[Unexpected LDAP Sync Error] %s: %s\n", p, debug.Stack())
	// 			switch p := p.(type) {
	// 			case error:
	// 				fmt.Println(p)
	// 			default:
	// 				fmt.Printf("%s\n", p)
	// 			}
	// 		}
	// 	}()
	// 	for {
	// 		fmt.Println("333")
	// 		time.Sleep(sleepDuration)
	// 		if _, ok := looper.store[orgUUID]; !ok {
	// 			looper.lock.Lock()
	// 			delete(looper.store, runFlag)
	// 			looper.lock.Unlock()
	// 			fmt.Println("Exit ...")
	// 			return
	// 		}
	// 		err := runLDAPSync(orgUUID)

	// 		// syncType := 6
	// 		// if err != nil {
	// 		// 	SetSyncStatus(orgUUID, syncType, true)
	// 		// } else {
	// 		// 	SetSyncStatus(orgUUID, syncType, false)
	// 		// }
	// 		fmt.Printf("[Looper.run] error: %v\n", err)
	// 	}
	// 	fmt.Println("over")
	// }()
	// looper.lock.Lock()
	// looper.store[runFlag] = struct{}{}
	// looper.lock.Unlock()

	runFlag := orgUUID + "run"
	if _, ok := looper.store[runFlag]; ok {
		log.Info("Looper.run] sync is already exist")
		return
	}
	log.Info("Looper.run] looper ldap start")

	sleepDuration := time.Second * 1
	go func() {
		defer func() {
			fmt.Println("[Looper.run] goroutine exit, clean up env start")
			looper.lock.Lock()
			delete(looper.store, runFlag)
			looper.lock.Unlock()
			if p := recover(); p != nil {
				log.Warnf("[Looper.run] error, %s: %s", p, debug.Stack())
				switch p := p.(type) {
				case error:
					log.Error(p)
				default:
					log.Errorf("%s", p)
				}
				err := looper.ReLoad(orgUUID, true)
				if err != nil {
					log.Error(err)
				}
			} else {
				log.Infof("[Looper.run] goroutine exit")
			}
		}()
		i := 1
		for {
			i++
			time.Sleep(sleepDuration)
			if _, ok := looper.store[orgUUID]; !ok {
				return
			}
			syncType := 6
			log.Infof("[Looper.run] ldap sync start, sync_type=%d", syncType)
			err := runLDAPSync(orgUUID)
			if err != nil {
				SetSyncStatus(orgUUID, syncType, true)
				log.Error(err)
			} else {
				SetSyncStatus(orgUUID, syncType, false)
			}
			log.Infof("[Looper.run] ldap sync end, sync_type=%d", syncType)
			if i == 5 {
				panic("AAA")
			}
		}

	}()
	looper.lock.Lock()
	looper.store[runFlag] = struct{}{}
	looper.lock.Unlock()
}

func runLDAPSync(orgUUID string) error {
	fmt.Println("run ldap sync")
	return nil
}

var lock = sync.Mutex{}
var mapOrgSyncStatus = map[string]bool{}

func SetSyncStatus(orgUUID string, syncType int, isFailed bool) {
	key := fmt.Sprintf("%s-%d", orgUUID, syncType)
	lock.Lock()
	mapOrgSyncStatus[key] = isFailed
	lock.Unlock()
}

func SyncFailed(orgUUID string, syncType int) bool {
	key := fmt.Sprintf("%s-%d", orgUUID, syncType)
	isFailed := false
	lock.Lock()
	isFailed = mapOrgSyncStatus[key]
	lock.Unlock()
	return isFailed
}
