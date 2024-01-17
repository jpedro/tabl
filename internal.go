package tables

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	alignLeft = iota
	alignRight
)

var (
	ansiCodeRegex = regexp.MustCompile(`\033\[(.*?)m`)
)

func countColumns(data [][]any) int {
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
	columns := countColumns(data)
	widths := make([]int, columns)
	aligns := make([]int, columns)

	// Assume they all are numeric, because you want to stop as soon
	// as you find the first non-numeric value and align it as string
	for i := range aligns {
		aligns[i] = alignRight
	}

	for i, row := range data {
		for j, val := range row {
			dirty, _ := getTextAlign(val)
			// fmt.Printf("  dirty (%T %d): %v\n", dirty, len(dirty), dirty)
			clean := cleanText(dirty)
			// fmt.Printf("  clean (%T %d): %v\n", clean, len(clean), clean)
			// Why twice? We want the real alignment of the clean value
			_, align := getTextAlign(clean)
			// fmt.Printf("  value (%T %d): %v\n", value, len(value), value)

			length := len(clean)
			if length > widths[j] {
				widths[j] = length
			}

			if i == 0 {
				// fmt.Printf("Skip the header for the aligns: %v\n", i)
				continue
			}

			if aligns[j] == alignLeft {
				// fmt.Printf("Left align means a we already marked it as such: %v\n", j)
				continue
			}

			aligns[j] = align
			// fmt.Printf("Doing align [%d]: %d\n", align, j)
		}
	}

	separator := "   "
	format := ""
	align := ""
	for i := range widths {
		align = ""
		if aligns[i] == alignLeft {
			align = "-"
		}
		format = fmt.Sprintf("%s%%%s%dv%s", format, align, widths[i], separator)
	}

	// Remove the trailing separator
	format = format[0:len(format)-len(separator)] + "\n"

	// fmt.Printf("==> Columns: %d\n", columns)
	// fmt.Printf("==> Widths:  %v\n", widths)
	// fmt.Printf("==> Aligns:  %v\n", aligns)
	// fmt.Printf("==> Format:  %s\n", format)

	return format, columns, widths
}

func padValue(text string, number int) string {
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

func getTextAlign(val any) (string, int) {
	switch v := val.(type) {
	case int, int8, int16, int64,
		uint, uint8, uint16, uint32, uint64:
		// `byte` is a type alias for `uint8`
		// `rune` is a type alias for `int32`
		// Excluded `uintptr`
		// https://github.com/golang/go/blob/master/src/builtin/builtin.go#L90-L94
		return fmt.Sprintf("%d", v), alignRight

	case float32, float64:
		return fmt.Sprintf("%v", v), alignRight

	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return fmt.Sprintf("%v", i), alignRight
		}

		f, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return fmt.Sprintf("%v", f), alignRight
		}
	}

	return fmt.Sprintf("%s", val), alignLeft
}
