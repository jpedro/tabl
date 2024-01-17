package main

import (
	"github.com/jpedro/table"
)

func main() {
	var data [][]any

	data = append(data, []any{"KEY", "DESCRIPTION", "NUMBER_NUMBER_NUMBER", "ALMOST A NUMBER"})
	data = append(data, []any{"Some metric", "a", "1", 1})
	data = append(data, []any{"This is a looooong value", "Some text", "1", 1.23})
	data = append(data, []any{"key", "And now for something completely different", "\033[31;1m333\033[0m", "3a"})

	table.Print(data)
}
