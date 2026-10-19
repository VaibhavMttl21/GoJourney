package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("race Condition")
	wg := &sync.WaitGroup{}
	// mut := &sync.RWMutex{}

	// var score = []int{0}

	wg.Add(2)

	// go func(wg *sync.WaitGroup, m *sync.RWMutex) {
	// 	defer wg.Done()
	// 	fmt.Println("One R")
	// 	mut.Lock()
	// 	score = append(score, 1)
	// 	mut.Unlock()
	// }(wg, mut)

	// go func(wg *sync.WaitGroup, m *sync.RWMutex) {
	// 	defer wg.Done()
	// 	fmt.Println("Two R")
	// 	mut.Lock()
	// 	score = append(score, 2)
	// 	mut.Unlock()
	// }(wg, mut)

	// go func(wg *sync.WaitGroup, m *sync.RWMutex) {
	// 	defer wg.Done()
	// 	fmt.Println("Three R")
	// 	mut.Lock()
	// 	score = append(score, 3)
	// 	mut.Unlock()
	// }(wg, mut)

	// go func(wg *sync.WaitGroup, m *sync.RWMutex) {
	// 	defer wg.Done()
	// 	fmt.Println("Four R")
	// 	mut.RLock()
	// 	fmt.Println(score)
	// 	mut.RUnlock()
	// }(wg, mut)

	var counter = 0

	go func(wg *sync.WaitGroup,) {
		defer wg.Done()
		counter++
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		counter++
	}(wg)

	wg.Wait()
	// fmt.Println(score)
	fmt.Println(counter)
}
