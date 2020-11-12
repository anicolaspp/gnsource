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
	it.started = true
	
	if it.idx < len(it.data)-1 {
		it.idx++

		return true
	}

	return false
}

func (it Iterator) Current() interface{} {
	if !it.started {
		panic("Current cannot be used before MoveNext()")
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

type Succ func(interface{}) interface{}

type Generator struct {
	seed interface{}

	next Succ

	current interface{}
}

func NewGenerator(seed interface{}, next Succ) *Generator {
	return &Generator{seed: seed, next: next, current: nil}
}

func (g *Generator) MoveNext() bool {
	return true
}

func (g *Generator) Current() interface{} {
	if g.current == nil {
		g.current = g.seed
	} else {
		g.current = g.next(g.current)
	}

	return g.current
}



