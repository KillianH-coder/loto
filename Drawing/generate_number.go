package Drawing // directory name

import (
	"math/rand/v2"
)

func GenerateNumbers() [6]int {
	numbers := [6]int{}
	for index := 0; index < len(numbers); {
		nb := rand.IntN(48) + 1
		if Verify(numbers, nb) {
			numbers[index] = nb
			index++
		}
	}
	numbers[5] = rand.IntN(9) + 1
	return numbers
}

func Verify(numbers [6]int, nb int) bool {
	for _, value := range numbers {
		if nb == value {
			return false
		}
	}
	return true
}
