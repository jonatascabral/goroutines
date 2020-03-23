package main

import (
	"log"
)

type Model struct {
	ID int
	Name string
}

func setId(m *Model, number int, result chan bool) {
	m.ID = number
	log.Println(m)
	result <- true
}

func main() {
	results := make([]chan bool, 0)
	model := Model{}
	for i := 0; i < 5; i++ {
		results = append(results, make(chan bool))
		go setId(&model, i, results[i])
	}
	log.Println("Main thread done")
	log.Println("Waiting channels")
	log.Println(<-results[len(results)-1])
	log.Println(model)
}
