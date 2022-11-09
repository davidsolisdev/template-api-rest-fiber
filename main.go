package main

func main() {
	var err error = App().Listen(":3005")
	if err != nil {
		panic(err.Error())
	}
}
