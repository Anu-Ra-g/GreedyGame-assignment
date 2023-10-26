package main

import (
	"sync"
	"time"
)

type Command struct {
	Command string `json:"command"`
}

type ListModel struct {
	Value []string
}

type Keystore struct {
	keys map[string]KeyModel
	mu   sync.Mutex
}

type KeyModel struct {
	Value      string
	ExTime     string
	InsertTime time.Time
}
