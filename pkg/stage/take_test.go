package stage

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTake(t *testing.T) {

	s := NewComposable([]interface{}{1,2,3,4,5})

	result := s.Take(2).ToList()

	if size := len(result); size != 2  {
		t.Errorf("Take did not take 2!!!")
	}

	result = NewComposable([]interface{}{1,2,3,4,5}).Take(100).ToList()

	if size := len(result); size != 5  {
		t.Errorf("Take did not take 5!!!")
	}
}

func TestTake_Random(t *testing.T) {
	rand.Seed(int64(time.Now().Nanosecond()))

	it := NewRandomIterator()

	size := rand.Int() % 100

	t.Log(fmt.Sprintf("Taking %v random numbers", size))

	composable := NewComposableFromIterator(it)

	result := composable.Take(size).ToList()

	if len(result) != size {
		t.Errorf("Take did not take %v!!!", size)
	}
}
