package main

import "fmt"

const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexStarving1
	mutexStarving2
	mutexStarving3
	mutexStarving4
	mutexStarving5
	mutexWaiterShift = iota
)

func main() {
	// println(mutexLocked)
	// println(mutexWoken)
	// println(mutexStarving)
	// println(mutexStarving1)
	// println(mutexStarving2)
	// println(mutexStarving3)
	// println(mutexStarving4)
	// println(mutexStarving5)
	// println(mutexWaiterShift)

	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"
}
