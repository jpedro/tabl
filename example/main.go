package main

import (
	"github.com/jpedro/tablelize"
)

func main() {
	var data [][]any

	data = append(data, []any{"KEY", "DESCRIPTION", "NUMBER", "ALMOST A NUMBER"})
	data = append(data, []any{"Some metric", "a", "1", 1})
	data = append(data, []any{"This is a looooong value", "Some text", "-1.23", 1.23})
	data = append(data, []any{"key", "And now for something completely different", "3", "3a"})

	tablelize.Print(data)
}
