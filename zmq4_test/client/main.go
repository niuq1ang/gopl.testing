package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

// 定义一个插件的运行时环境
// 实现了简单的RPC协议

const (
	REQUEST_TIMEOUT = 2500 * time.Millisecond //  msecs, (> 1000!)
	REQUEST_RETRIES = 8                       //  Before we abandon
	SERVER_ENDPOINT = "tcp://127.0.0.1:5555"
)

func main() {
	fmt.Println("connecting to server...")
	client, err := zmq.NewSocket(zmq.DEALER)
	if err != nil {
		panic(err)
	}
	client.Connect(SERVER_ENDPOINT)

	poller := zmq.NewPoller()
	poller.Add(client, zmq.POLLIN)

	for {
		client.SendMessage("plugin02")
		//  Poll socket for a reply, with timeout
		sockets, err := poller.Poll(REQUEST_TIMEOUT)
		if err != nil {
			break //  Interrupted
		}

		if len(sockets) > 0 {
			//  We got a reply from the server, must match sequence
			reply, err := client.RecvMessage(0)
			if err != nil {
				fmt.Printf("reply error: %v\n", reply)
				break //  Interrupted
			}
			fmt.Printf("accept: %v\n", reply)

		} else {
			fmt.Println("W: no response from server, retrying...")
		}
		time.Sleep(1e9)
	}
	// defer client.Close()

	// fmt.Println("E: runtime env shutdowning...")
}
