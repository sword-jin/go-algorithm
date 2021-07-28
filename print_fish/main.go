package main

import (
	"fmt"
	"sync"
)

// <><_
// ><>_

type state uint8

const (
	Pending state = iota
	A             //     B   >   C
	B             //   <            <
	C             // A                D
	D             //   >            >
	E             //     E   <   F
	F             // D _ A
)

type rule struct {
	from  state
	print string
	to    state
}

var rules = []rule{
	{from: A, print: "<", to: B},
	{from: B, print: ">", to: C},
	{from: C, print: "<", to: D},
	{from: D, print: "_", to: A},
	{from: A, print: ">", to: E},
	{from: E, print: "<", to: F},
	{from: F, print: ">", to: D},
}

func getNextState(current state, print string) state {
	for _, r := range rules {
		if r.from == current && r.print == print {
			return r.to
		}
	}
	return Pending
}

func main() {
	var current = A
	var cond = sync.NewCond(new(sync.Mutex))

	var printCh = func(ch string) {
		for {
			cond.L.Lock()
			var next state
			for {
				next = getNextState(current, ch)
				if next == Pending {
					cond.Wait()
				} else {
					break
				}
			}

			fmt.Print(ch)

			current = next
			cond.L.Unlock()
			cond.Broadcast()
		}
	}

	for i := 0; i < 100; i++ {
		go printCh("_")
		go printCh("<")
		go printCh(">")
	}

	ch := make(chan int, 1)
	<-ch
}
