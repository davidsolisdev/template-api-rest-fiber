package main

import "github.com/davidsolisdev/template-api-rest-fiber/internal/server"

func main() {
	var err error = server.App().Listen(":3005")
	if err != nil {
		panic(err.Error())
	}
}
