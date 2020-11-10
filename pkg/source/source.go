package source

import "github.com/anicolaspp/gnsource/pkg/stage"

// From creates a new ComposableStage from some slice
func From(data []interface{}) stage.Composable {
	s := stage.NewComposable(data)

	return s
}

func FromIterator(it *stage.Iterator) stage.Composable  {
	return stage.NewComposableFromIterator(it)
}
