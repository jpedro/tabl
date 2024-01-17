package tables

import (
	"fmt"
)

type Table struct {
	rows [][]any
}

func New() *Table {
	return &Table{}
}

func (me *Table) Add(values ...any) {
	me.rows = append(me.rows, values)
}

func (me *Table) Print() {
	Print(me.rows)
}

// Prints the table.
func Print(data [][]any) {
	fmt.Println(Render(data))
}

// Returns the rendered table.
func Render(data [][]any) string {
	format, columns, widths := calcFormat(data)
	text := ""
	var row []any
	var val any
	for i := range data {
		row = data[i]
		len := len(row)
		args := make([]any, columns)
		for i := range row {
			val = row[i]
			raw, _ := getValueAlign(val)
			args[i] = pad(raw, widths[i])
		}

		if len < columns {
			for i := len; i < columns; i++ {
				args[i] = ""
			}
		}
		// fmt.Printf("args: %#v\n", args)
		// colored := format
		// if i%2 == 1 {
		// 	colored = "\033[48;5;237m" + format + "\033[0m"
		// }
		// colored = colored + "\n"
		text = text + fmt.Sprintf(format, args...)
	}

	return text
}
