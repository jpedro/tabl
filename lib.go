package tabl

import (
	"fmt"
)

var (
	ShowDebug     = false
	RowPadding    = ""
	RowStarting   = ""
	RowFinish     = ""
	CellSeparator = "   "
)

func New() *Table {
	return &Table{}
}

func Print(data [][]any) {
	fmt.Println(Render(data))
}

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
			raw, _ := alignText(val)
			args[i] = padValue(raw, widths[i])
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
		// text = text + RowPadding + RowStarting + fmt.Sprintf(format, args...) + RowFinish
		text = text + fmt.Sprintf(format, args...)
	}

	return text
}
