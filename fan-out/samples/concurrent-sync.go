package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(2)
}

type T struct {
	ID int
}

func main() {
	queue := make(chan T)

	// START OMIT
	var wg sync.WaitGroup // HL

	n := 5
	for i := n; i > 0; i-- {
		wg.Add(1) // HL
		go worker(queue, &wg)
	}

	fmt.Println("start")
	for _, d := range seed(10) {
		queue <- d
	}
	close(queue)
	wg.Wait() // HL
	fmt.Println("done")
	// END OMIT
}

// STARTW OMIT
func worker(queue chan T, wg *sync.WaitGroup) {
	defer wg.Done() // HL
	for d := range queue {
		emit(d)
	}
}

// ENDW OMIT

func emit(d T) {
	fmt.Println("emit", d.ID)
	time.Sleep(1 * time.Second)
}

func seed(n int) []T {
	var data []T
	for i := 1; i <= n; i++ {
		data = append(data, T{i})
	}
	return data
}
