package service

import (
  zmq "github.com/pebbe/zmq4"
  "log"
  "testing"
  "time"
)

func TestListen(t *testing.T) {
  log.Println("Starting test on LISTEN")
  publisher, errSock := zmq.NewSocket(zmq.PUB)
  defer publisher.Close()
  if errSock != nil {
    t.Fatal(errSock)
  }
  log.Println("publisher created")

  endpoint := "tcp://127.0.0.1:7000"
  errBin := publisher.Bind(endpoint)
  if errBin != nil {
    t.Fatal(errBin)
  }
  log.Println("publisher bound to", endpoint)

  service := Service{ "name", "localhost", "REP", 7000 }
  go Listen("tcp://127.0.0.1:7000", service)
  publisher.Send("youpi", 0)
  log.Println("send")
}
