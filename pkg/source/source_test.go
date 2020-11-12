package source

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFromGenerator(t *testing.T) {
	rand.Seed(int64(time.Now().Nanosecond()))

	succ := func(value interface{}) interface{} {
		return value.(int) + 1
	}

	seed := rand.Int() % 100
	toTake := rand.Int() % 1000

	result := FromGenerator(seed, succ).Take(toTake).ToList()

	t.Log(fmt.Sprintf("Starting at %v and taking %v", seed, toTake))

	if len(result) != toTake {
		t.Error("Wrong generator")
	}
}
