package service

import (
  "log"
)

func Whereis(endpoint, name string) Service {
  log.Println("Where is", name,"?")
  return Service{
    "youpi",
    "192.168.0.104",
    "REP",
    5000,
  }
}
