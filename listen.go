package service

import (
  zmq "github.com/pebbe/zmq4"
  "log"
)

// Listen will listen to all publication for registration from registry
func Listen(endpoint string, service Service) {
  subscriber, errSock := zmq.NewSocket(zmq.SUB)
  defer subscriber.Close()
  if errSock != nil {
    log.Fatal(errSock)
  }
  log.Println("subscriber created")

  errConn := subscriber.Connect(endpoint)
  if errConn != nil {
    log.Fatal(errConn)
  }
  log.Println("subscriber connected to", endpoint)

  errSub := subscriber.SetSubscribe("")
  if errSub != nil {
    log.Fatal(errSub)
  }
  log.Println("subscriber subscribing to all messages")

  for {
    log.Println("waiting")
    msg, errRecv := subscriber.Recv(0)
    log.Println("receive publication, registration will be done")
    log.Println(msg)
    if errRecv != nil {
      log.Fatal(errRecv)
    }
    Register(endpoint, service)
  }
}
