package stage

type BasicComposable struct {
	it Lazy
}

func (stage *BasicComposable) Take(n int) Composable {
	return NewTake(stage, n)
}

func (stage *BasicComposable) TakeWhile(f func(interface{}) bool) Composable {
	panic("implement me")
}

func (stage *BasicComposable) First() Runnable {
	panic("implement me")
}

func (stage *BasicComposable) ToList() []interface{} {
	return ToList(stage)
}

func (stage *BasicComposable) Fold(zero interface{}, fn func(interface{}, interface{}) interface{}) Runnable {
	panic("implement me")
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

