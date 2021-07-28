package main

import (
	"fmt"
	"sync"
)

type Philosopher struct {
	ID      int
	cond    *sync.Cond
	forks   []int
	eatChan chan int
}

func (p Philosopher) Start() {
	leftForkIndex := (p.ID - 1) % len(p.forks)
	if leftForkIndex == -1 {
		leftForkIndex = len(p.forks) - 1
	}
	rightForkIndex := (p.ID + 1) % len(p.forks)
	go func() {
		p.cond.L.Lock()
		for p.forks[leftForkIndex] == 1 || p.forks[rightForkIndex] == 1 {
			p.cond.Wait()
		}
		p.forks[leftForkIndex] = 1
		p.forks[rightForkIndex] = 1
		p.cond.L.Unlock()

		p.Eat()

		p.cond.L.Lock()
		p.forks[leftForkIndex] = 0
		p.forks[rightForkIndex] = 0
		p.cond.L.Unlock()
		p.cond.Broadcast()
	}()
}

func (p Philosopher) Eat() {
	fmt.Printf("Philosopher#%d eating\n", p.ID)
	p.eatChan <- 1
}

func main() {
	var eatChan = make(chan int, 5)
	var cond = sync.NewCond(&sync.Mutex{})
	var forks = make([]int, 1)

	var phils []Philosopher
	for i := 0; i < 5; i++ {
		phils = append(phils, Philosopher{
			ID:      i,
			cond:    cond,
			forks:   forks,
			eatChan: eatChan,
		})
	}
	for _, p := range phils {
		p.Start()
	}

	i := 0
	for i < 5 {
		select {
		case <-eatChan:
			i++
		}
	}
}
