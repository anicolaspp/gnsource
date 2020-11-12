package stage

type Composable interface {
	Map(fn func(interface{}) interface{}) Composable

	ForEach(func(interface{})) Runnable

	Filter(func(interface{}) bool) Composable

	Take(n int) Composable

	TakeWhile(func(interface{}) bool) Composable

	First() Runnable

	ToList() []interface{}

	Fold(zero interface{}, fn func(interface{}, interface{}) interface{}) Runnable

	MoveNext() bool

	Current() interface{}
}

//Runnable represents a runnable operation on a Composable. Composable is a lazy structure and to get the final result,
//one should invoke a runnable operation.
type Runnable interface {
	Run() interface{}
}

//Done represent correct termination of Runnable execution
func Done() int {
	return 0
}

//NewComposable creates a Composable from slice
func NewComposable(data []interface{}) Composable {
	return NewComposableFromIterator(NewIteratorFromData(data))
}

//NewComposableFromIterator creates a Composable from an Iterator
func NewComposableFromIterator(iterator Lazy) Composable {
	return &BasicComposable{iterator}
}

//ToList converts the Composable Stage into a slice
func ToList(stage Composable) []interface{} {
	result := make([]interface{}, 0)

	for stage.MoveNext() {
		result = append(result, stage.Current())
	}

	return result
}
