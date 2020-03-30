package main

import (
	"fmt"
	"log"
	"time"
)

func showPanics(number int) {
	defer recoverFromPanics()
	if number > 5 {
		panic(fmt.Sprintf("%d > 5", number))
	}
	time.Sleep(1 * time.Second)
	log.Printf("number: %d\n", number)
}

func main() {
	defer recoverFromPanics()
	for i := 0; i < 10; i++ {
		go showPanics(i)
	}
	time.Sleep(2 * time.Second)
	log.Println("Main thread done")
}

func recoverFromPanics() {
	if recover() != nil {
		log.Println("Recovered")
	}
}