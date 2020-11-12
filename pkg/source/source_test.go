package source

import (
	"fmt"
	"testing"
)

func TestFromGenerator(t *testing.T) {
	succ := func(value interface{}) interface{} {
		return value.(int) + 1
	}

	result := FromGenerator(0, succ).Take(10).ToList()

	fmt.Println(result)

	if len(result) != 10  {
		t.Error("Wrong generator")
	}
}
