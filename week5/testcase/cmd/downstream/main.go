package main

import "golang_weekend/week5/testcase/server"

func main() {
	server.NewDownStreamServer(0.2).Run(":8000")
}
