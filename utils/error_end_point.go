package utils

import "fmt"

func ErrorEndPoint(endPoint string, err error) {
	fmt.Println(endPoint + ": " + err.Error())
}
