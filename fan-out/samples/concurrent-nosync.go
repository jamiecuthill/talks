package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(2)
}

type T struct {
	ID int
}

// START OMIT
func main() {
	queue := make(chan T)

	n := 5
	for i := n; i > 0; i-- {
		go worker(queue)
	}

	fmt.Println("start")
	for _, d := range seed(10) {
		queue <- d
	}
	close(queue)
	fmt.Println("done")
}

func emit(d T) {
	fmt.Println("emit", d.ID)
	time.Sleep(1 * time.Second)
}

// END OMIT

func worker(queue chan T) {
	for d := range queue {
		emit(d)
	}
}

func seed(n int) []T {
	var data []T
	for i := 1; i <= n; i++ {
		data = append(data, T{i})
	}
	return data
}
