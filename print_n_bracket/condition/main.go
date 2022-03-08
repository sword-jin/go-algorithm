package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var n, limit int
var cond = sync.NewCond(&sync.Mutex{})

func printLeft() {
	for {
		cond.L.Lock()
		for !(n < limit) {
			cond.Wait()
		}
		fmt.Printf("(")
		n++
		cond.Broadcast()
		cond.L.Unlock()
	}
}

func printRight() {
	for {
		cond.L.Lock()
		for !(n > 0) {
			cond.Wait()
		}
		fmt.Printf(")")
		n--
		cond.Broadcast()
		cond.L.Unlock()
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
