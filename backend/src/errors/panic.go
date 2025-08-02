package errors

import "fmt"

func RecoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered from panic:", r)
	}
}