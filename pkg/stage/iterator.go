package stage

type Iterator struct {
	data []interface{}

	idx int
}

func NewIteratorFromData(data []interface{}) *Iterator {
	return &Iterator{data: data, idx: -1}
}

func (it *Iterator) MoveNext() bool {
	if it.idx < len(it.data)-1 {
		it.idx++

		return true
	}

	return false
}

func (it Iterator) Current() interface{} {
	return it.data[it.idx]
}
