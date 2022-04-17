package main

import (
	"golang_weekend/week5/testcase/server"
	"time"
)

func main() {
	server.NewUpStreamServer(
		10,
		50,
		0.8,
		time.Second*5,
	).Run(":9000")
}
