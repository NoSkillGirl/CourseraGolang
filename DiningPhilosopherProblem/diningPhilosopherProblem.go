package main

import (
	"fmt"
	"sync"
)

type chops struct {
	sync.Mutex
}

type philo struct {
	leftCS, rightCS *chops
}

func (p philo) eat(i int, wg *sync.WaitGroup) {

	k := 0
	for k < 3 {
		// time.Sleep(1 * time.Second)
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("starting to eat ", i)
		fmt.Println("finishing eating", i)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		k++
	}

	wg.Done()
}

func main() {
	CSticks := make([]*chops, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(chops)
	}

	philos := make([]*philo, 5)

	for i := 0; i < 5; i++ {
		philos[i] = &philo{
			CSticks[i],
			CSticks[(i+1)%5],
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			wg.Wait()
		}
		wg.Add(1)
		go philos[i].eat(i, &wg)
	}
	wg.Wait()
}
