package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {

	zctx, err := zmq.NewContext()

	if err != nil {
		panic(err)
	}

	defer zctx.Term()

	//  Socket to talk to clients

	s, err := zctx.NewSocket(zmq.REP)

	if err != nil {
		panic(err)
	}
	defer s.Close()

	s.Bind("shmem://test")

	fmt.Println("Server started")
	fmt.Println("Waiting for messages...")

	for {
		//  Wait for next request from client
		msg, err := s.Recv(0)

		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		println("Received ", msg)

		//  Do some 'work'
		//time.Sleep(time.Second)

		//  Send reply back to client
		reply := "Yeah, I got your message"
		s.Send(reply, 0)
		println("Sent ", reply)
	}

}
