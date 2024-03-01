package main

import "beneficiary-validate/controller"

func main() {
	server := controller.NewServer()
	server.Start(":8080")
}
