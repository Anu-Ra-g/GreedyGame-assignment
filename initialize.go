package main

import (
	"sync"
	"time"
)

var keystore Keystore
var liststore map[string]ListModel

func InitializeKeystore() {

	keystore = Keystore{
		keys: map[string]KeyModel{
			"key1": {
				Value:      "This is working",
				ExTime:     "20",
				InsertTime: time.Now(),
			},
			"key2": {
				Value:      "value2",
				ExTime:     "78",
				InsertTime: time.Now(),
			},
		},
		mu: sync.Mutex{},
	}
}

func InitializeListStore() {
	liststore = map[string]ListModel{
		"demo_list": {
			Value: []string{"1", "2", "3"},
		},
	}
}
