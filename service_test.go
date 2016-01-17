package service

import (
  "testing"
  "fmt"
  zmq "github.com/pebbe/zmq4"
  "encoding/json"
)

func TestRegisterAllInterfaces(t *testing.T) {
    // Socket to receive message
    receiver, _ := zmq.NewSocket(zmq.PULL)
    defer receiver.Close()
    receiver.Bind("tcp://127.0.0.1:3001")

    endpoint := "tcp://127.0.0.1:3001"

    cases := []struct {
      nameIn string
      protocolIn string
      portIn int
      want string
    }{
      { "youpi", "HTTP", 80, `{"action":"register","service":{"name":"youpi","hostname":"localhost","protocol":"HTTP","port":80}}`},
    }

    for _, c := range cases {
      registeredIps := RegisterAllInterfaces(endpoint, c.nameIn, c.protocolIn, c.portIn)
      var registrations []string
      for _ = range registeredIps {
        got,_ := receiver.Recv(0)
        fmt.Println("Received message", got)
        registrations = append(registrations, got)
      }
      if len(registeredIps) != len(registrations) {
        t.Errorf("Didn't receive expected number of registrations. Expected: %d, received: %d", len(registeredIps), len(registrations))
      }
    }
}

func TestRegister(t *testing.T) {
    // Socket to receive message
    receiver, _ := zmq.NewSocket(zmq.PULL)
    defer receiver.Close()
    receiver.Bind("tcp://127.0.0.1:3001")

    endpoint := "tcp://127.0.0.1:3001"

    service := Service{
      Name: "youpi",
      Hostname: "localhost",
      Protocol: "HTTP",
      Port: 80,
    }
    cases := []struct {
      in Service
      want string
    }{
      { service, `{"action":"register","service":{"name":"youpi","hostname":"localhost","protocol":"HTTP","port":80}}`},
    }

    for _, c := range cases {
      Register(endpoint, c.in)

      got,_ := receiver.Recv(0)
      fmt.Println("Received message", got)
      if got != c.want {
        inString,_ := json.Marshal(c.in)
        t.Errorf("Register(%q, %s) sent %s, expected %s", endpoint, inString, got, c.want)
      }
    }
}
