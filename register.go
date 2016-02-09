package service

import (
	"encoding/json"
	zmq "github.com/pebbe/zmq4"
	"github.com/satori/go.uuid"
	"log"
)

// Register will register one service
func Register(endpoint string, service Service) {
	pusher, errSock := zmq.NewSocket(zmq.PUSH)
	defer pusher.Close()
	if errSock != nil {
		log.Fatal(errSock)
	}
	log.Println("pusher created")
	errCon := pusher.Connect(endpoint)
	if errCon != nil {
		log.Fatal(errCon)
	}
	log.Println("pusher connected to", endpoint)

	registerMsg := message{
		Jsonrpc: "2.0",
		ID:      uuid.NewV4().String(),
		Method:  "register",
		Params:  registerparams{service},
	}
	registerMsgByt, errMar := json.Marshal(registerMsg)
	if errMar != nil {
		log.Fatal(registerMsg)
		log.Fatal(errMar)
	}

	pusher.Send(string(registerMsgByt), 0)
	log.Println("pusher sent some message ", string(registerMsgByt))
}
