package main

import (
	"container/ring"
	"fmt"
)

func main() {

	nums := []int{4, 8, 15, 16, 23, 42}

	r := ring.New(len(nums))

	// Initialize the ring with values
	for i := 0; i < r.Len(); i++ {
		r.Value = nums[i]
		r = r.Next()
	}

	// Print the ring elements 5 times
	for i := 0; i < len(nums); i++ {
		r.Do(func(p interface{}) {
			fmt.Printf("%02d ", p)
		})

		// Move the ring to the next element for offset
		r = r.Next()

		// Print a new line after each iteration
		fmt.Println()
	}

}
