package service

import (
	zmq "github.com/pebbe/zmq4"
	"log"
	"testing"
)

func r(replier *zmq.Socket) {
	for {
		got, _ := replier.Recv(0)
		log.Println("Received", got)
		replier.Send(`{"jsonrpc":"2.0","id":"youpi","result":{"name":"test","hostname":"localtest","protocol":"test","port":0}}`, 0)
	}
}
func TestWhereis(t *testing.T) {
	replier, errSock := zmq.NewSocket(zmq.REP)
	defer replier.Close()
	if errSock != nil {
		t.Fatal(errSock)
	}

	endpoint := "tcp://127.0.0.1:6000"
	errBin := replier.Bind(endpoint)
	if errBin != nil {
		t.Fatal(errBin)
	}

	cases := []struct {
		nameIn string
	}{
		{"article"},
	}

	go r(replier)
	for _, c := range cases {
		service := Whereis(endpoint, c.nameIn)
		log.Println(service)
	}
}
