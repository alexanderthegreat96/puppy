package puppy

import (
	"github.com/alexanderthegreat96/dog",
	"math/rand"
)

func Bark() string {
	return "Random Woofs"
}

func Barks() string {
	return "Woofs! Woofs! Woofs!"
}

// imported function WhenGrownUp
func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

// imported function WhenGrownUp
func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

// imported BarkAt
func BarkAt(s string) string {
	return dog.BarkAt(s)
}

func TagMe() string {
	return "Tagged You!"
}

func RandomNumber() int {
	return rand.Intn(100)
}