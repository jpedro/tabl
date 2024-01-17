package tables

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	// Should this be aligned to the right?
	ALIGN_LEFT = 0
	ALIGN_RIGHT = 1
)
var (
	ansiCodeRegex = regexp.MustCompile(`\033\[(.*?)m`)
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

// Deprecated: Use `tablelize.Print` instead.
func Rows(data [][]any) {
	Print(data)
}

// Prints the table.
func Print(data [][]any) {
	fmt.Println(Render(data))
}

func calcColumns(data [][]any) int {
	count := 0

	for i := range data {
		length := len(data[i])
		if length > count {
			count = length
		}
	}

	return count
}

func calcFormat(data [][]any) (string, int, []int) {
	columns := calcColumns(data)
	widths := make([]int, columns)
	aligns := make([]int, columns)

	// Assume they all are numeric, because you want to stop as soon
	// as you find the first non-numeric value and align it as string
	for i := range aligns {
		aligns[i] = ALIGN_RIGHT
	}

	// align := ALIGN_RIGHT
	// dirty := ""
	// clean := ""
	for i, row := range data {
		for j, val := range row {
			dirty, _ := getValueAlign(val)
			// fmt.Printf("  dirty (%T %d): %v\n", dirty, len(dirty), dirty)
			clean := cleanText(dirty)
			// fmt.Printf("  clean (%T %d): %v\n", clean, len(clean), clean)
			_, align := getValueAlign(clean)
			// fmt.Printf("  value (%T %d): %v\n", value, len(value), value)

			length := len(clean)
			if length > widths[j] {
				widths[j] = length
			}

			// Skip the header for the aligns
			if i == 0 {
				continue
			}

			// Left align means a we already marked it as such
			if aligns[j] == ALIGN_LEFT {
				// fmt.Printf("  Skipping another ALIGN_LEFT: %v\n", j)
				continue
			}

			aligns[j] = align
			// fmt.Printf("  Doing align [%d]: %d\n", align, j)
		}
	}

	format := ""
	align := ""
	for i := range widths {
		align = ""
		if aligns[i] == ALIGN_LEFT {
			align = "-"
		}
		format = fmt.Sprintf("%s%%%s%dv   ", format, align, widths[i])
	}
	// Remove the trailing separator `   `
	format = format[0:len(format)-3] + "\n"

	// fmt.Printf("==> Columns: %d\n", columns)
	// fmt.Printf("==> Widths:  %v\n", widths)
	// fmt.Printf("==> Aligns:  %v\n", aligns)
	// fmt.Printf("==> Format:  %s\n", format)

	return format, columns, widths
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

func pad(text string, number int) string {
	clean := cleanText(text)
	if len(clean) == len(text) {
		return text
	}
	result := text
	// fmt.Printf("  - Pad: %s vs %s (%d vs %d) with %d\n", text, clean, number, len(clean), number - len(clean))
	for i := 0 ; i < (number - len(clean)); i++ {
		result = " " + result
	}
	return result
}

func cleanText(text string) string {
	if !strings.Contains(text, "\033") {
		return text
	}

	res := ansiCodeRegex.ReplaceAllString(text, "")
	return res
}

// // Checks is a value can be cast into an integer or a float.
// func isNumeric(val any) bool {
// 	switch v := val.(type) {
// 	case int, int8, int16, int64,
// 		uint, uint8, uint16, uint32, uint64:
// 		// `byte` is a type alias for `uint8`
// 		// `rune` is a type alias for `int32`
// 		// Excluded `uintptr`
// 		// https://github.com/golang/go/blob/master/src/builtin/builtin.go#L90-L94
// 		return true

// 	case float32, float64:
// 		return true

// 	case string:
// 		_, err := strconv.ParseInt(v, 10, 64)
// 		if err == nil {
// 			return true
// 		}

// 		_, err = strconv.ParseFloat(v, 64)
// 		if err == nil {
// 			return true
// 		}
// 	}

// 	return false
// }

func getValueAlign(val any) (string, int) {
	switch v := val.(type) {
	case int, int8, int16, int64,
		uint, uint8, uint16, uint32, uint64:
		// `byte` is a type alias for `uint8`
		// `rune` is a type alias for `int32`
		// Excluded `uintptr`
		// https://github.com/golang/go/blob/master/src/builtin/builtin.go#L90-L94
		return fmt.Sprintf("%d", v), ALIGN_RIGHT

	case float32, float64:
		return fmt.Sprintf("%v", v), ALIGN_RIGHT

	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return fmt.Sprintf("%v", i), ALIGN_RIGHT
		}

		f, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return fmt.Sprintf("%v", f), ALIGN_RIGHT
		}
	}

	return fmt.Sprintf("%s", val), ALIGN_LEFT
}
