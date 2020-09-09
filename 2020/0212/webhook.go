package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	// fileName := todayFilename()
	// 打开以当前日期为文件名的文件（不存在则创建文件，存在则追加内容）
	f := newLogFile()
	defer f.Close()

	log.SetOutput(f)
	Run("127.0.0.1:8888")
}

func Run(addr string) {
	app := iris.Default()
	app.Use(recover.New())

	f := newLogFile()
	defer f.Close()
	app.Logger().SetOutput(f)

	w := app.Party("/webhook")
	w.Post("/accept", Accept)
	app.Run(iris.Addr(addr))
}

type Base struct {
	SequenceNo string `json:"id"`
}

func Accept(c iris.Context) {
	s, _ := ioutil.ReadAll(c.Request().Body)
	log.Printf("RequestBody: %v\n", string(s))
	result := new(Base)
	_ = json.Unmarshal([]byte(s), result)
	log.Println(result.SequenceNo)
	c.WriteString(result.SequenceNo)
}

func newLogFile() *os.File {
	// 打开以当前日期为文件名的文件（不存在则创建文件，存在则追加内容）
	f, err := os.OpenFile("webhook.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
