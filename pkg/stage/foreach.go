package stage

type ForEach struct {
	prev Composable

	fn func(interface{})
}

func (foreach ForEach) Run() interface{} {
	for foreach.prev.MoveNext() {
		foreach.fn(foreach.prev.Current())
	}

	return Done()
}

func NewForEach(prev Composable, fn func(interface{})) Runnable {
	return ForEach{prev: prev, fn: fn}
}

