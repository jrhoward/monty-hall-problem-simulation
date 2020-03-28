package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dontchange := 0.0 // never change your choice win count
	change := 0.0     // always change your choice win count
	size := 3         // number of doors
	runs := 1000000   // number of experiments
	var result []bool
	for i := 0; i < runs; i++ {
		result = run(size) // run the experiment
		if result[0] {
			dontchange++ // How often my origonal selection would result in a win
		} else {
			change++ // How often changing my selection would result in a win
		}
	}
	fmt.Println("Don't change wins: ", dontchange/float64(runs)*100)
	fmt.Println("Always change wins: ", change/float64(runs)*100)
}

// random index to simulate a door selection
func selectDoor(size int) int {
	min := 0
	max := size - 1
	index := rand.Intn(max-min+1) + min
	return index
}

// run the experiment
func run(size int) []bool {
	prize := selectDoor(size)     // randomly set the door that has the prize
	selection := selectDoor(size) // Contestant selects a door
	//The final two options provided to the Contestant
	final := make([]bool, 2)
	// index 0 is the users origonal choice which is a win if it is the prize
	final[0] = prize == selection
	// index 1 is the alternate choice presented
	final[1] = !final[0]
	return final
}
