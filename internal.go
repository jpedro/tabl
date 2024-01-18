package tables

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	alignLeft = iota
	alignRight
)

var (
	ansiCodeRegex = regexp.MustCompile(`\033\[(.*?)m`)
)

// We allow rows have diff number of columns
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

func log(message string, args ...any) {
	if !ShouldLog {
		return
	}
	fmt.Fprintf(os.Stderr, message, args...)
}

func calcFormat(data [][]any) (string, int, []int) {
	columns := countColumns(data)
	// Preallocate is always faster
	widths := make([]int, columns)
	aligns := make([]int, columns)
	log("\n")
	log("Columns: %v\n", columns)
	log("Widths:  %v\n", widths)
	log("Aligns:  %v\n", aligns)

	// Assume they all are numeric, because you want to stop as soon
	// as you find the first non-numeric value and align it as string
	for i := range aligns {
		aligns[i] = alignRight
	}

	for i, row := range data {
		for j, val := range row {
			log("\n")
			log("- VAL in row %d, col %d = '%v'\n", i, j, val)
			dirty, old := getTextAlign(val)
			log("  dirty ([%d]%T): '%v'\n", len(dirty), dirty, dirty)
			log("    align %v\n", old)
			clean := cleanText(dirty)
			log("  clean ([%d]%T): '%v'\n", len(clean), clean, clean)
			// Why twice? We want the real alignment of the clean value
			value, align := getTextAlign(clean)
			log("  value ([%d]%T): '%v'\n", len(value), value, value)
			log("    align %v\n", align)

			// length := len(clean)
			length := utf8.RuneCountInString(clean)
			if length > widths[j] {
				log("  extending width col=%d from %d to %d\n", j, widths[j], length)
				widths[j] = length
			}

			if i == 0 {
				log("  skip the header for col=%d\n", i)
				continue
			}

			if aligns[j] == alignLeft {
				log("  left align means already set for col=%d\n", j)
				continue
			}

			log("  BEFORE REALIGN col=%d, cur=%d, new=%d\n", j, aligns[j], align)
			aligns[j] = align
			log("  AFTER REALIGN cur=%d\n", aligns[j])
		}
	}

	// separator := " | "
	format := ""
	align := ""
	for i := range widths {
		align = ""
		if aligns[i] == alignLeft {
			align = "-"
		}
		format = fmt.Sprintf("%s%%%s%ds%s", format, align, widths[i], CellSeparator)
	}

	log("\n")
	log("Format: %s\n", format)
	log("Widths:  %v\n", widths)
	log("Aligns:  %v\n", aligns)

	// Remove the trailing separator
	// We use the raw bytes instead of the UTF-8 length
	// format = format[0:len(format)-utf8.RuneCountInString(RowSeparator)] + "\n"
	format = RowPadding + RowStarting + format[0:len(format)-len(CellSeparator)] + RowFinish+ "\n"

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

// Checks is a value can be cast into an integer or a float.
//lint:ignore U1000 Shhhh
func isNumeric(val any) bool {
	switch v := val.(type) {
	case int, int8, int16, int64,
		uint, uint8, uint16, uint32, uint64:
		// `byte` is a type alias for `uint8`
		// `rune` is a type alias for `int32`
		// Excluded `uintptr`
		// https://github.com/golang/go/blob/master/src/builtin/builtin.go#L90-L94
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

	return fmt.Sprintf("%v", val), alignLeft
}
