package main

import (
	"errors"
	"fmt"
)

func init() {
	InitializeListStore()
}

func qpush(cmd []string) {
	key := cmd[1]

	val, exists := liststore[key]

	if !exists {
		liststore[key] = ListModel{
			Value: cmd[2:],
		}
	} else {
		updatedArr := append(val.Value, cmd[2:]...)
		val := ListModel{
			Value: updatedArr,
		}
		liststore[key] = val
	}

	fmt.Println(liststore)
}

func qpop(cmd []string) (string, error) {

	key := cmd[1]

	val, exists := liststore[key]

	popped_value := ""

	if exists && len(val.Value) > 0 {
		index := len(val.Value) - 1
		popped_value = val.Value[index]
		newArr := val.Value[:index]

		liststore[key] = ListModel{
			Value: newArr,
		}

		return popped_value, nil
	}

	if exists && len(val.Value) == 0 {
		delete(liststore, key)
	}

	fmt.Println(liststore)

	return "", errors.New("something went wrong")

}
