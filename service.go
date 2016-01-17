package service

import (
  "log"
  "encoding/json"
  zmq "github.com/pebbe/zmq4"
  "github.com/LeReunionais/ip"
)

type message struct {
  Action string `json:"action"`
  Service Service`json:"service"`
}

type Service struct {
  Name string `json:"name"`
  Hostname string `json:"hostname"`
  Protocol string `json:"protocol"`
  Port int `json:"port"`
}

func RegisterAllInterfaces (endpoint, name, protocol string, port int) []string {
  ips := ip.Ips()
  for _, ip := range ips {
    service := Service{
      Name: name,
      Hostname: ip,
      Protocol: protocol,
      Port: port,
    }
    Register(endpoint, service)
  }
  return ips
}

func Register (endpoint string, service Service) {
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
    Action: "register",
    Service: service,
  }
  registerMsgByt,errMar := json.Marshal(registerMsg)
  if errMar != nil {
    log.Fatal(registerMsg)
    log.Fatal(errMar)
  }

  pusher.Send(string(registerMsgByt), 0)
  log.Println("pusher sent some message ", string(registerMsgByt))
}
