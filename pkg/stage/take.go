package stage

type Take struct {
	prevStage Composable

	n int
}

func (t *Take) Map(fn func(interface{}) interface{}) Composable {
	return NewMapped(t, fn)
}

func (t *Take) ForEach(f func(interface{})) Runnable {
	return NewForEach(t, f)
}

func (t *Take) Filter(f func(interface{}) bool) Composable {
	return NewFiltered(t, f)
}

func (t *Take) Take(n int) Composable {
	return NewTake(t, n)
}

func (t *Take) TakeWhile(f func(interface{}) bool) Composable {
	return NewTakeWhile(t, f)
}

func (t *Take) First() Runnable {
	panic("implement me")
}

func (t *Take) ToList() []interface{} {
	return ToList(t)
}

func (t *Take) Fold(zero interface{}, fn func(interface{}, interface{}) interface{}) Runnable {
	panic("implement me")
}

func (t *Take) MoveNext() bool {
	if t.n > 0 && t.prevStage.MoveNext() {
		t.n--

		return true
	}

	return false
}

func (t *Take) Current() interface{} {
	return t.prevStage.Current()
}

func NewTake(prevStage Composable, n int) *Take {
	return &Take{prevStage: prevStage, n: n}
}

