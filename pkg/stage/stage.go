package stage

type Composable interface {
	Map(fn func(interface{}) interface{}) Composable

	ForEach(func(interface{})) Runnable

	MoveNext() bool

	Current() interface{}
}

type Runnable interface {
	Run() interface{}
}

func Done() int {
	return 0
}

func NewComposable(data []interface{}) Composable {
	return &BasicComposable{NewIteratorFromData(data)}
}

func ToList(stage Composable) []interface{} {
	result := make([]interface{}, 0)

	for stage.MoveNext() {
		result = append(result, stage.Current())
	}

	return result
}

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
