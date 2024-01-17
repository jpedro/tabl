package main

import (
	"github.com/jpedro/tables"
)

func main() {
	table := tables.New()

	table.Add("KEY", "DESCRIPTION", "COLORED_NUMBERS", "ALMOST_A_NUMBER")
	table.Add("Some metric", "", '1', 1) // "49" might be the wrong value
	table.Add("Nobody expects the Spanish Inquisition", "Ah!", string('1'), 1.23)
	table.Add("Uh", "Integer with color codes", "\033[32;1m333\033[0m", "3a")

	table.Print()
}
