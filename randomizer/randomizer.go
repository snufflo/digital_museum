package randomizer

import (
	"math/rand/v2"
)

func RandNum(max int) int {
	return rand.IntN(max)
}
