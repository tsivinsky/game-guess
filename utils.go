package main

import (
	"math/rand"
	"sync"
)

func GenerateRandomNumbers(count int, max int) []int {
	var nums []int
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)

		go func() {
			num := rand.Intn(max)
			nums = append(nums, num)

			wg.Done()
		}()
	}

	wg.Wait()

	return nums
}
