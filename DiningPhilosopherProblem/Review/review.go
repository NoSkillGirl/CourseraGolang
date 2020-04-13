package main

import (
	"fmt"
	"sync"
)

// Chopstick would be later put into a pool
type Chopstick struct {
}

// Philosopher is labelled with num
type Philosopher struct {
	num int
}

// chopstick pool
var pool = sync.Pool{
	New: func() interface{} {
		return new(Chopstick)
	},
}

var w sync.WaitGroup

// asking permission to the host
func askPermissionToHost(host chan int) {
	<-host
}

// just releasing the host
func sayThanksToHost(host chan int) {
	host <- 1
}

// eat!
func (p Philosopher) eat(host chan int) {
	defer w.Done()
	defer sayThanksToHost(host)

	askPermissionToHost(host)

	// pick up chopsticks, any order
	left := pool.Get()
	right := pool.Get()

	fmt.Printf("starting to eat %d\n", p.num)
	fmt.Printf("finishing eating %d\n", p.num)
	// time.Sleep(2 * time.Second)

	// then return the chopsticks
	pool.Put(right)
	pool.Put(left)
}

// initialize chopsticks pool
func initChopsticks(amount int) {
	for i := 0; i < 5; i++ {
		pool.Put(new(Chopstick))
	}
}

// creates and inits the philosophers with their numbers
func initPhilosophers(amount int) []*Philosopher {
	philosophers := make([]*Philosopher, amount)
	for i := 0; i < amount; i++ {
		// each philosopher is numbered
		philosophers[i] = &Philosopher{i + 1}
	}
	return philosophers
}

// Dinner time
func haveDinner(philosophers []*Philosopher, host chan int) {
	for _, philosopher := range philosophers {
		// lets eat 3 times
		for i := 0; i < 3; i++ {
			w.Add(1)
			go philosopher.eat(host)
		}
	}
}

func main() {
	// host allows no more than 2 philosophers to eat concurrently
	host := make(chan int, 2)

	// number of philosophers and chopsticks
	const amount = 5

	// creates the chopsticks
	initChopsticks(amount)
	// creates the philosophers
	philosophers := initPhilosophers(amount)

	// dinning!
	haveDinner(philosophers, host)

	// init the host with 2 possible slots to eat
	host <- 1
	host <- 1

	// just waiting everyone eat 3 times
	w.Wait()
}
