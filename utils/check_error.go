package utils

import "log"

func CheckIfError(err error, msg string)  {
	if err != nil {
		log.Fatal(msg)
		return
	}
}