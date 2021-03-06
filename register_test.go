package service

import (
	"encoding/json"
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"regexp"
	"testing"
)

func TestRegister(t *testing.T) {
	// Socket to receive message
	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	error := receiver.Bind("tcp://127.0.0.1:6002")
	if error != nil {
		t.Fatal(error)
	}

	endpoint := "tcp://127.0.0.1:6002"

	service := Service{
		Name:     "youpi",
		Hostname: "localhost",
		Protocol: "HTTP",
		Port:     80,
	}
	cases := []struct {
		in   Service
		want string
	}{
		{service, `{"jsonrpc":"2.0","id":".*","method":"register","params":{"service":{"name":"youpi","hostname":"localhost","protocol":"HTTP","port":80}}}`},
	}

	for _, c := range cases {
		go Register(endpoint, c.in)

		fmt.Println("Waiting for something to be received")
		got, _ := receiver.Recv(0)
		fmt.Println("Received message", got)
		matched, _ := regexp.MatchString(c.want, got)
		if !matched {
			inString, _ := json.Marshal(c.in)
			t.Errorf("Register(%q, %s) sent %s, expected %s", endpoint, inString, got, c.want)
		}
	}
}
