package utils

import "log"

func ErrorEndPoint(endPoint string, err error) {
	log.Fatal(endPoint + ": " + err.Error())
}
