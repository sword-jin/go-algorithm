package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/sync/semaphore"
)

var empty *semaphore.Weighted
var filled *semaphore.Weighted

func printLeft() {
	for {
		empty.Acquire(context.TODO(), 1)
		fmt.Printf("(")
		filled.Release(1)
	}
}

func printRight() {
	for {
		filled.Acquire(context.TODO(), 1)
		fmt.Printf(")")
		empty.Release(1)
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("give limit")
	}
	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("limit integer")
	}
	empty = semaphore.NewWeighted(int64(limit))
	filled = semaphore.NewWeighted(int64(limit))
	filled.Acquire(context.TODO(), int64(limit))

	done := make(chan bool, 1)
	for i := 0; i < 8; i++ {
		go printLeft()
		go printRight()
	}
	<-done
}
