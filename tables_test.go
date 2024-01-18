package tables

import (
	"testing"
)

func TestRenderNumber(t *testing.T) {
	var data [][]any

	data = append(data, []any{"KEY", "VAL"})
	data = append(data, []any{"One", "1"})

	returned := Render(data)
	expected := "KEY   VAL\nOne     1\n"

	if returned != expected {
		t.Errorf("Failed to render properly")
	}
}

func TestRenderMixed(t *testing.T) {
	var data [][]any

	data = append(data, []any{"KEY", "VAL"})
	data = append(data, []any{"One", "1a"})

	returned := Render(data)
	expected := "KEY   VAL\nOne   1a \n"

	if returned != expected {
		t.Errorf("Failed to render properly")
	}
}

func TestPrint(t *testing.T) {
	var data [][]any

	data = append(data, []any{"KEY", "VAL"})
	data = append(data, []any{"One", "\033[1m1\033[0m]"})

	Print(data)
}
