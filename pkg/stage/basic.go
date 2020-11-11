package stage

type BasicComposable struct {
	it *Iterator
}

func (stage *BasicComposable) Filter(f func(interface{}) bool) Composable {
	return NewFiltered(stage, f)
}

func (stage *BasicComposable) ForEach(f func(interface{})) Runnable {
	return NewForEach(stage, f)
}

func (stage *BasicComposable) MoveNext() bool {
	return stage.it.MoveNext()
}

func (stage *BasicComposable) Current() interface{} {
	return stage.it.Current()
}

func (stage *BasicComposable) Map(fn func(interface{}) interface{}) Composable {
	mapped := NewMapped(stage, fn)

	return mapped
}
