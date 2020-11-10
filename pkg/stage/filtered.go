package stage

type Filtered struct {
	prev Composable

	predicate func(interface{}) bool
}

func (f Filtered) Filter(fn func(interface{}) bool) Composable {
	return NewFiltered(f, fn)
}

func NewFiltered(prev Composable, predicate func(interface{}) bool) *Filtered {
	return &Filtered{prev: prev, predicate: predicate}
}

func (f Filtered) Map(fn func(interface{}) interface{}) Composable {
	return &Mapped{prevStage: f, mapping: fn}
}

func (f Filtered) ForEach(fn func(interface{})) Runnable {
	return NewForEach(f, fn)
}

func (f Filtered) MoveNext() bool {
	if !f.prev.MoveNext() {
		return false
	}

	if !f.predicate(f.prev.Current()) {
		return f.MoveNext()
	}

	return true
}

func (f Filtered) Current() interface{} {
	return f.prev.Current()
}
