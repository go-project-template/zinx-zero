package main

import (
	"fmt"
	"time"
)

// test for docker-modd
func main() {
	for {
		time.Sleep(time.Second)
		fmt.Println("test for docker-modd:", time.Now().String())
	}
}
