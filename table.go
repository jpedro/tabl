package table

import (
	"fmt"
	"regexp"
	"strconv"
	// "unicode"
)

const (
	ALIGN_STRING int = iota
	ALIGN_NUMBER
)


// Deprecated: Use `tablelize.Print` instead.
func Rows(data [][]any) {
	Print(data)
}

// Prints the table.
func Print(data [][]any) {
	fmt.Println(Render(data))
}

// Returns the rendered table.
func Render(data [][]any) string {
	var widths []int
	var aligns []int

	fields := len(data[0])
	widths = make([]int, fields)
	aligns = make([]int, fields)

	// Assume they all are numeric, because you want to stop as soon
	// as you find the first non-numeric value and align it as string
	for i := range aligns {
		aligns[i] = ALIGN_NUMBER
	}

	dirty := ""
	clean := ""
	for i, row := range data {
		for j, val := range row {
			dirty = fmt.Sprintf("%s", val)
			fmt.Printf("dirty (%v): %v\n", len(dirty), dirty)
			clean = cleanText(dirty)
			fmt.Printf("clean (%d): %v\n", len(clean), clean)
			len := len(clean)
			if len > widths[j] {
				widths[j] = len
			}

			// Skip the header for the aligns
			if i == 0 {
				continue
			}

			// String align means a value already failed to be numeric
			if aligns[j] == ALIGN_STRING {
				continue
			}

			switch val.(type) {
			case int, uint, int8, uint8, int16, uint16, int32, uint64, uintptr:
			case float32, float64:
			case string:
				if !isNumeric(clean) {
					fmt.Printf("clean is NOT NUM: %v\n", clean)
					aligns[j] = ALIGN_STRING
					break
				}
				fmt.Printf("clean is NUM: %v\n", clean)
			default:
				aligns[j] = ALIGN_STRING
			}
		}
	}

	format := ""
	align := ""
	for i := range widths {
		align = ""
		if aligns[i] == ALIGN_STRING {
			align = "-"
		}
		format = fmt.Sprintf("%s%%%s%dv   ", format, align, widths[i])
	}

	format = format[0:len(format)-3] + "\n"
	fmt.Printf("format: %v\n", format)
	fmt.Printf("widths: %v\n", widths)
	text := ""
	var row []any
	var val any
	for i := range data {
		row = data[i]
		args := make([]any, len(row))
		for i := range row {
			val = row[i]
			// dirty = fmt.Sprintf("%v", val)
			// fmt.Printf("dirty: %v\n", dirty)
			// clean = cleanText(dirty)
			// fmt.Printf("clean: %v\n", clean)
			args[i] = val

		}
		fmt.Printf("args: %#v\n", args)
		text = text + fmt.Sprintf(format, args...)
	}

	return text
}

func cleanText(text string) string {
    m := regexp.MustCompile(`\033\[(.*?)m`)
    res := m.ReplaceAllString(text, "")
	return res
	// // fmt.Printf("text: %s\n", text)
	// clean := []rune{}
	// for _, r := range text {
	// 	if !unicode.IsPrint(r) {
	// 		continue
	// 	}
	// 	clean = append(clean, r)
	// }

	// fmt.Printf("clean: (%d) %s\n", len(clean), string(clean))
	// return string(clean)
}

// Checks is a value can be cast into an integer or a float.
func isNumeric(val any) bool {
	switch v := val.(type) {
	case int, int8, int16, int64,
		uint, uint8, uint16, uint32, uint64:
		// https://github.com/golang/go/blob/master/src/builtin/builtin.go#L90-L94
		// `byte` is a type alias for `uint8`
		// https://github.com/golang/go/blob/master/src/builtin/builtin.go#L94
		// `rune` is a type alias for `int32`
		// Excluded `uintptr`
		return true

	case float32, float64:
		return true

	case string:
		_, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return true
		}

		_, err = strconv.ParseFloat(v, 64)
		if err == nil {
			return true
		}
	}

	return false
}
