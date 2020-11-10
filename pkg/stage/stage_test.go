package stage

import (
	"fmt"
	"testing"
)

func TestBasicComposable_Map(t *testing.T) {
	data := make([]interface{}, 0)

	for i := 0; i < 5; i++ {
		data = append(data, i)
	}

	basic := NewComposable(data)

	mapped := basic.Map(func(value interface{}) interface{} {
		return value
	}).Map(func(value interface{}) interface{} {
		return value.(int) + 1
	})

	n := 0

	ret := mapped.ForEach(func(value interface{}) {
		n += 1
	}).Run()

	fmt.Println(ret)
	fmt.Println(n)

	fmt.Println(ToList(mapped))
}
