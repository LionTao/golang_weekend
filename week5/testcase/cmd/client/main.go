package main

import (
	"golang_weekend/week5/testcase/client"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			client.Client()
			wg.Done()
		}()
	}
	wg.Wait()
}
