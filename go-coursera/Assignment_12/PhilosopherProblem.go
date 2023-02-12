package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
Explanation ->
	This idea is about using a buffered channel (with size 2) to act as 'host', which
means following the requirement of 'no more than 2 philosophers eat concurrently'.
	By using buffered channel, when a philo-goroutine try to eat, it first tries to put
data into the channel. It will be blocked if channel had already got 2 data (reach the
capacity), then after one philo-goroutine finished eating (1 time), it will consume data
from the channel and this would let other goroutine can put more data into it.
	In addition, when the philo-goroutine performs the "eating" operation, a 1s sleep ops
is added, for better observation on checking this channel-solution works.
	A simple random function by using rand.Float64() > 0.5 is for picking up chopsticks
'randomly'.

*/

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
}

// eat would try requiring both leftCs & rightCs randomly 3-times
// num -> number of the philosopher
// cn -> for getting the permission from 'host'
// wg -> for main goroutine not exit before all other go routines finished
func (p Philo) eat(num int, cn chan int, wg *sync.WaitGroup) {
	// eat 3-times
	for i := 0; i < 3; i++ {
		// try input into channel, as request permission from host
		cn <- 1
		// in any order
		if randomPicking() {
			p.leftCS.Lock()
			p.rightCS.Lock()
		} else {
			p.rightCS.Lock()
			p.leftCS.Lock()
		}
		fmt.Printf("starting to eat<%v>\n", num)
		time.Sleep(1 * time.Second)
		fmt.Printf("finishing eating<%v>\n", num)
		p.leftCS.Unlock()
		p.rightCS.Unlock()
		// consume as 'giving-back' permission to host
		<-cn
	}
	wg.Done()
}

// randomPicking is for return ture/false randomly
func randomPicking() bool {
	return rand.Float64() > 0.5
}

func main() {

	// 1. initialization
	chopSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{
			chopSticks[i],
			chopSticks[(i+1)%5]}
	}
	host := make(chan int, 2)
	wg := sync.WaitGroup{}

	// 2. run
	for i := 0; i < len(philos); i++ {
		wg.Add(1)
		go philos[i].eat(i+1, host, &wg)
	}
	wg.Wait()

}
