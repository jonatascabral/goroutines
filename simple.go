package main

import (
	"fmt"
	"log"
	"time"
)

func show(number int) {
	time.Sleep(5 * time.Second)
	fmt.Printf("number: %d\n", number)
}

func main() {
	for i := 0; i < 10; i++ {
		go show(i)
	}
	time.Sleep(2 * time.Second)
	log.Println("Main thread done")
}
