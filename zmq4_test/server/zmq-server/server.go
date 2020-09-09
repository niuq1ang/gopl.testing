package zmq

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

const (
	REQUEST_TIMEOUT = 2500 * time.Millisecond //  msecs, (> 1000!)
	REQUEST_RETRIES = 8                       //  Before we abandon
	SERVER_ENDPOINT = "tcp://127.0.0.1:5555"
)

func Run() {
	fmt.Println("starting server...")
	srv, err := zmq.NewSocket(zmq.ROUTER)
	if err != nil {
		panic(err)
	}
	srv.Bind(SERVER_ENDPOINT)

	poller := zmq.NewPoller()
	poller.Add(srv, zmq.POLLIN)

	for {
		sockets, err := poller.Poll(REQUEST_TIMEOUT)
		if err != nil {
			fmt.Printf("Interrupted: %v\n", err)
			break //  Interrupted
		}

		if len(sockets) > 0 {
			reply, err := srv.RecvMessage(0)
			if err != nil {
				fmt.Printf("reply error: %+v\n", reply)
				break
			}
			fmt.Printf("accept: %+v\n", reply[1])
			response := []string{reply[0], "ok"}
			fmt.Println(srv.SendMessage(response))
		} else {
			fmt.Printf("nothing recived \n")
		}
		time.Sleep(1e9)
	}
}
