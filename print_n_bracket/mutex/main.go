package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var lock = sync.Mutex{}
var count int
var limit int

func printLeft() {
	for {
	retry:
		lock.Lock()
		if count == limit {
			lock.Unlock()
			goto retry
		}
		count++
		fmt.Printf("(")
		lock.Unlock()
	}
}

func printRight() {
	for {
	retry:
		lock.Lock()
		if count == 0 {
			lock.Unlock()
			goto retry
		}
		count--
		fmt.Printf(")")
		lock.Unlock()
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("give limit")
	}
	var err error
	limit, err = strconv.Atoi(os.Args[1])
	if err != nil {
		panic("limit integer")
	}

	done := make(chan bool, 1)
	for i := 0; i < 8; i++ {
		go printLeft()
		go printRight()
	}
	<-done
}
