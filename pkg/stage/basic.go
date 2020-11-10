package stage

type BasicComposable struct {
	it *Iterator
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
	mapped := &Mapped{prevStage: stage, mapping: fn}

	return mapped
}
