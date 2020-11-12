package main

import (
	"fmt"
	"github.com/anicolaspp/gnsource/pkg/source"
	"math"
)

func main() {
	fmt.Println("Hello from gnsource module")

	stream := source.From([]interface{}{0, 1, 2, 3, 4, 5})

	runnable := stream.Map(func(value interface{}) interface{} {
		return math.Pow10(value.(int))
	}).ForEach(func(value interface{}) {
		fmt.Println(value)
	})

	fmt.Println(runnable.Run())
}
