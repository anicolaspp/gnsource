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
	}).Filter(func(value interface{}) bool {
		return value == 3 || value == 5
	})

	ret := mapped.ForEach(func(value interface{}) {
		fmt.Println(value)
	}).Run()

	fmt.Println(ret)
}
