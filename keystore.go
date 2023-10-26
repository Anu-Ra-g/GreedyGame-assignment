package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func init() {
	InitializeKeystore()
}

func get(cmd []string) (string, error) {

	keystore.mu.Lock()
	defer keystore.mu.Unlock()

	val, exists := keystore.keys[cmd[1]]

	if exists {
		return val.Value, nil
	}

	return "", errors.New("something went wrong")
}

func set(cmd []string) {

	keystore.mu.Lock()
	defer keystore.mu.Unlock()

	_, exists := keystore.keys[cmd[1]]

	match1 := containsStringElement(cmd, "NX")
	match2 := containsStringElement(cmd, "XX")

	match3 := containsStringElement(cmd, "EX")

	var dummy KeyModel

	if len(cmd) <= 6 {

		if !exists && match1 {
			dummy.Value = cmd[2]
		} else if exists && match2 {
			dummy.Value = cmd[2]
			dummy.InsertTime = time.Now()
		} else if match3 {
			dummy.Value = cmd[2]
			dummy.ExTime = cmd[4]
			dummy.InsertTime = time.Now()
		} else {
			dummy.Value = cmd[2]
		}

	}

	keystore.keys[cmd[1]] = dummy

	fmt.Println(keystore.keys)

}

func deleteExpiredKeys() {
	for {
		time.Sleep(10 * time.Second)

		currentTime := time.Now()

		keystore.mu.Lock()

		for key, item := range keystore.keys {
			if item.ExTime != "" {
				exTimeValue, err := strconv.Atoi(item.ExTime)
				if err != nil {
					continue
				}

				timeDifference := currentTime.Sub(item.InsertTime)
				if int(timeDifference.Seconds()) >= exTimeValue {
					delete(keystore.keys, key)
				}
			}
		}

		keystore.mu.Unlock()

	}
}

func containsStringElement(arr []string, target string) bool {
	for _, element := range arr {
		if element == target {
			return true
		}
	}
	return false
}
