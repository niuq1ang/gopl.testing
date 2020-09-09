package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"time"
)

const defaultWatchPath = "conf/"

var flag = false

func main() {
	watch()
	fmt.Println("process will go on")
	time.Sleep(time.Second * 20)
}

func watch() {
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Printf("watch file err: %v", err)
	}
	defer watch.Close()
	fileName := "conf/ones_license.crt"
	fmt.Println(fileName)
	err = watch.Add(defaultWatchPath)
	if err != nil {
		log.Printf("watch file err: %v", err)
	}
	go func() {
		count := 0
		for {
			count++
			fmt.Printf("preceed times=%d \n", count)
			select {
			case ev := <-watch.Events:
				{
					// if fileName == ev.Name && ((ev.Op&fsnotify.Create == fsnotify.Create) || (ev.Op&fsnotify.Write == fsnotify.Write) || (ev.Op == fsnotify.Remove) ||
					// 	(ev.Op == fsnotify.Rename) || (ev.Op == fsnotify.Chmod)) {
					// 	log.Printf(" operate file_name=%s \n", ev.Name)
					// }
					if ev.Op == fsnotify.Create && ev.Name == "conf/ones_license.crt" {
						log.Printf("create file file_name=%s \n", ev.Name)
						flag = true
						// continue
					}
					if ev.Op == fsnotify.Write && ev.Name == "conf/ones_license.crt" {
						flag = true
						log.Printf("write file file_name=%s \n", ev.Name)
						// return
					}
					if ev.Op == fsnotify.Remove && ev.Name == "conf/ones_license.crt" {
						log.Printf("delete file file_name=%s \n", ev.Name)
					}
					if ev.Op == fsnotify.Remove && ev.Name == "conf/ones_license.crt" {
						log.Println("删除文件 : ", ev.Name)
					}
					if ev.Op == fsnotify.Rename && ev.Name == "conf/ones_license.crt" {
						log.Println("重命名文件 : ", ev.Name)
					}
					if ev.Op == fsnotify.Chmod && ev.Name == "conf/ones_license.crt" {
						log.Println("修改权限 : ", ev.Name)
					}
				}
			case err := <-watch.Errors:
				{
					log.Printf("watch file err: %v \n", err)
					return
				}
			}
		}
	}()

	select {}
}
