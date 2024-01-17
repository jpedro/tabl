package main

import (
	"github.com/jpedro/tables"
)

func main() {
	table := tables.New()

	table.Add("KEY", "DESCRIPTION", "COLORED_NUMBERS", "ALMOST_A_NUMBER")
	table.Add("Some metric", "", "1", 1)
	table.Add("This is a looooong key", "Some text", "1", 1.23)
	table.Add("Uh", "And now for something completely different", "\033[32;1m333\033[0m", "3")

	table.Print()
}
