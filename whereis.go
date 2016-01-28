package service

import (
	"encoding/json"
	zmq "github.com/pebbe/zmq4"
	"github.com/satori/go.uuid"
	"log"
)

func Whereis(endpoint, name string) Service {
	log.Println("Where is", name, "?")
	requester, errSock := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	if errSock != nil {
		log.Fatal(errSock)
	}
	log.Println("requester created")

	errCon := requester.Connect(endpoint)
	if errCon != nil {
		log.Fatal(errCon)
	}
	log.Println("requester connected")

	whereIsMsg := message{
		Jsonrpc: "2.0",
		ID:      uuid.NewV4().String(),
		Method:  "find",
		Params:  findparams{"article"},
	}
	whereIsMsgByt, errMar := json.Marshal(whereIsMsg)
	if errMar != nil {
		log.Fatal(whereIsMsgByt)
	}
	requester.Send(string(whereIsMsgByt), 0)
	log.Println("requester asked for", string(whereIsMsgByt))

	reply, errRec := requester.Recv(0)
	log.Println("received", reply)
	if errRec != nil {
		log.Fatal(errRec)
	}

	replyMsg := findreply{}
	errJSON := json.Unmarshal([]byte(reply), &replyMsg)
	if errJSON != nil {
		log.Fatal(errJSON)
	}

	return replyMsg.Result
}
