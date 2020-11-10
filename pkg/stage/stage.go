package stage

type Composable interface {
	Map(fn func(interface{}) interface{}) Composable

	ForEach(func(interface{})) Runnable

	Filter(func(interface{}) bool) Composable

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
	return NewComposableFromIterator(NewIteratorFromData(data))
}

func NewComposableFromIterator(iterator *Iterator) Composable {
	return &BasicComposable{iterator}
}

func ToList(stage Composable) []interface{} {
	result := make([]interface{}, 0)

	for stage.MoveNext() {
		result = append(result, stage.Current())
	}

	return result
}
