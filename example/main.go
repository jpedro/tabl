package main

import (
	"github.com/jpedro/tables"
)

func main() {
	table := tables.New()

	table.Add("KEY", "DESCRIPTION", "NUMBER_NUMBER_NUMBER", "ALMOST A NUMBER")
	table.Add("Some metric", "a", "1", 1)
	table.Add("This is a looooong value", "Some text", "1", 1.23)
	table.Add("key", "And now for something completely different", "\033[31;1m333\033[0m", "3a")

	table.Print()
}
