package main

import (
	"log"
)

var model *Model

type Model struct {
	ID int
	Name string
}

func setId(number int, result chan bool) {
	model.ID = number
	log.Println(model)
	result <- true
}

func main() {
	results := make([]chan bool, 0)
	for i := 0; i < 5; i++ {
		results = append(results, make(chan bool))
		go setId(i, results[i])
	}
	model = &Model{}
	log.Println("Main thread done")
	log.Println("Waiting channels")
	log.Println(<-results[len(results)-1])
	log.Println(model)
}
