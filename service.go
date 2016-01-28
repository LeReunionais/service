package service

import (
	"encoding/json"
	"github.com/LeReunionais/ip"
	zmq "github.com/pebbe/zmq4"
	"github.com/satori/go.uuid"
	"log"
)

// RegisterAllInterfaces will automatically register host interfaces
func RegisterAllInterfaces(endpoint, name, protocol string, port int) []string {
	ips := ip.Ips()
	for _, ip := range ips {
		service := Service{
			Name:     name,
			Hostname: ip,
			Protocol: protocol,
			Port:     port,
		}
		Register(endpoint, service)
	}
	return ips
}

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
