package stage

type TakeWhile struct {
	prevStage Composable

	predicate func(interface{}) bool
}

func (t *TakeWhile) Map(fn func(interface{}) interface{}) Composable {
	return NewMapped(t, fn)
}

func (t *TakeWhile) ForEach(f func(interface{})) Runnable {
	return NewForEach(t, f)
}

func (t *TakeWhile) Filter(f func(interface{}) bool) Composable {
	return NewFiltered(t, f)
}

func (t *TakeWhile) Take(n int) Composable {
	return NewTake(t, n)
}

func (t *TakeWhile) TakeWhile(f func(interface{}) bool) Composable {
	return NewTakeWhile(t, f)
}

func (t *TakeWhile) First() Runnable {
	panic("implement me")
}

func (t *TakeWhile) ToList() []interface{} {
	return ToList(t)
}

func (t *TakeWhile) Fold(zero interface{}, fn func(interface{}, interface{}) interface{}) Runnable {
	panic("implement me")
}

func (t *TakeWhile) MoveNext() bool {
	if t.prevStage.MoveNext() && t.predicate(t.prevStage.Current()) {
		return true
	}

	return false
}

func (t *TakeWhile) Current() interface{} {
	return t.prevStage.Current()
}

func NewTakeWhile(prevStage Composable, predicate func(interface{}) bool) *TakeWhile {
	return &TakeWhile{prevStage: prevStage, predicate: predicate}
}
