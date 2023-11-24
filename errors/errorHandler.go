package errorHandler

import (
	"fmt"
	"log"
)

func CheckError(err error, action string) bool {
	if err != nil {
		log.Printf("%s: %s\n", action, err.Error())
		fmt.Printf("%s: %s\n", action, err.Error())
		return true
	}
	return false
}

func CheckFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
