package main

import (
	"fmt"
	"log"
	"time"
)

func showChan(number int, result chan string) {
	time.Sleep(5 * time.Second)
	fmt.Printf("number: %d\n", number)
	result <- fmt.Sprintf("Done %d", number)
}

func main() {
	results := make([]chan string, 0)
	for i := 0; i < 10; i++ {
		results = append(results, make(chan string))
		go showChan(i, results[i])
	}
	time.Sleep(2 * time.Second)
	log.Println("Main thread done")
	log.Println("Waiting channels")
	log.Println(<-results[len(results)-1])
}
