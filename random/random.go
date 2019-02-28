package random

import (
	"math/rand"
	"time"
)

func Shuffle(target []int) []int {
	size := len(target)
	shuffled := append([]int{}, target...)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		randomIndex := rand.Int() % size
		value := shuffled[i]
		shuffled[i] = shuffled[randomIndex]
		shuffled[randomIndex] = value
	}

	return shuffled
}
