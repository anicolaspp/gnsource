package stage

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTakeWhile_ToList(t *testing.T) {
	s := NewComposable([]interface{}{1, 2, 3, 4, 5})

	result := s.TakeWhile(func(value interface{}) bool {
		n := value.(int)

		return n >= 1 && n < 5
	}).ToList()

	fmt.Println(result)

	if len(result) != 4 {
		t.Error("Takewhile did not take correct values")
	}
}

func TestTakeWhile_RandomTake(t *testing.T) {
	rand.Seed(int64(time.Now().Nanosecond()))

	it := NewRandomIterator()

	size := rand.Int() % 100

	taken := 0

	predicate := func(value interface{}) bool {
		taken++

		return taken <= size/2
	}

	result := NewComposableFromIterator(it).TakeWhile(predicate).ToList()

	if n := len(result); n != size/2 {
		t.Error(fmt.Sprintf("Takewhile should have taken %v values but took %v", size/2, n))
	}
}
