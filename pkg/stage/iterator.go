package stage

import (
	"math/rand"
	"time"
)

type Lazy interface {
	MoveNext() bool
	Current() interface{}
}

type Iterator struct {
	data []interface{}

	idx int

	started bool
}

func NewIteratorFromData(data []interface{}) *Iterator {
	return &Iterator{data: data, idx: -1, started: false}
}

func (it *Iterator) MoveNext() bool {
	if it.idx < len(it.data)-1 {
		it.idx++

		return true
	}

	return false
}

func (it Iterator) Current() interface{} {
	if !it.started {
		return panic
	}
	
	return it.data[it.idx]
}

type RandomIterator struct {
	rnd *rand.Rand
}

func NewRandomIterator() *RandomIterator {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return &RandomIterator{r}
}

func (r RandomIterator) MoveNext() bool {
	return true
}

func (r RandomIterator) Current() interface{} {
	return r.rnd.Uint64()
}
