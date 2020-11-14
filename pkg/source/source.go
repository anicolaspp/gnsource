package source

import "github.com/anicolaspp/gnsource/pkg/stage"

// From creates a new ComposableStage from some slice
func From(data []interface{}) stage.Composable {
	s := stage.NewComposable(data)

	return s
}

// FromIterator creates a Composable using an Iterator
func FromIterator(it *stage.Iterator) stage.Composable {
	return stage.NewComposableFromIterator(it)
}

// FromGenerator creates a Composable using an initial value seed, and a function succ that generates
// the next value.
// Notice that the values are generated lazily.
func FromGenerator(seed interface{}, succ stage.Succ) stage.Composable {
	return stage.NewComposableFromIterator(stage.NewGenerator(seed, succ))
}
