package stage

type Mapped struct {
	prevStage Composable

	mapping func(interface{}) interface{}
}

func NewMapped(prevStage Composable, mapping func(interface{}) interface{}) *Mapped {
	return &Mapped{prevStage: prevStage, mapping: mapping}
}

func (mapped *Mapped) Filter(fn func(interface{}) bool) Composable {
	return NewFiltered(mapped, fn)
}

func (mapped *Mapped) ForEach(f func(interface{})) Runnable {
	return NewForEach(mapped, f)
}

func (mapped *Mapped) Map(fn func(interface{}) interface{}) Composable {
	result := NewMapped(mapped, fn)

	return result
}

func (mapped Mapped) MoveNext() bool {
	return mapped.prevStage.MoveNext()
}

func (mapped *Mapped) Current() interface{} {

	value := mapped.prevStage.Current()

	return mapped.mapping(value)
}
