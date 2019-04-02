package random

import (
	"math/rand"
	"time"
)

// Shuffle is shuffle slice element
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

// RandomValues is create slice whitch have random values.
func RandomValues(maxValue int, length int) []int {
	rand.Seed(time.Now().UnixNano())

	randomValues := make([]int, length)
	for i := 0; i < length; i++ {
		randomValues[i] = rand.Intn(maxValue)
	}
	return randomValues
}
