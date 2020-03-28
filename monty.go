package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// Set up the experiment
	dontchange := 0 // never change your choice win count
	change := 0     // always change your choice win count
	size := 3       // number of doors
	runs := 1000000 // number of experiments
	var result []bool
	for i := 0; i < runs; i++ {
		result = run(size) // run the experiment
		if result[0] {
			dontchange++ // How often my origonal selection would result in a win
		} else {
			change++ // How often changing my selection would result in a win
		}
	}

	fmt.Println("Don't change wins: ", float64(dontchange)/float64(runs)*100)
	fmt.Println("Always change wins: ", float64(change)/float64(runs)*100)

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
	doors := make([]bool, size)    // values default to false
	doors[selectDoor(size)] = true // randomly set the door that has the prize

	selection := selectDoor(size) // Contestant selects a door

	//Simulate the final options provided to the Contestant
	final := make([]bool, 2)
	// index 0 is the users origonal choice
	final[0] = doors[selection]
	// index 1 is the alternate choice presented which can only be
	// false if the users selection is true or true if the users selection is false
	// which is simulated by flipping the boolean value of users selection
	final[1] = !doors[selection]

	return final
}
