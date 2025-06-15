package main

import (
	"github.com/jpedro/tabl"
)

func main() {
	table := tabl.New()

	table.Add("KEY", "DESCRIPTION", "COLORED_NUMBERS", "ALMOST_A_NUMBER")
	table.Add("Runes gone wrong", "", '1', 1) // "49" might be the wrong value
	table.Add("Array", [2]int{0, 1}, 'a', 1)
	table.Add("Slices", []int{0, 1}, 'A', 1)
	table.Add("Maps", map[string]int{"a": 1}, 1.23e2, 1)
	table.Add("Struct", struct{a int}{a: 123}, 1.23e-2, 1)
	table.Add("Nobody expects the Spanish Inquisition", "Ah!", string('1'), 1.23)
	table.Add("Uh", "Integer with color codes", "\033[32;1m333\033[0m", "3a")

	table.Print()
}
