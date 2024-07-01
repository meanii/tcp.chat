package cmd

import "sync"

type Room struct {
	Id      string
	Name    string
	Members []user
	Public  bool
}

type Rooms struct {
	Mutext sync.RWMutex
	Rooms  map[string]Room
}
