package main

func main() {
	var err error = App().Listen(":3000")
	if err != nil {
		panic(err.Error())
	}
}
