package tabl

import (
	"testing"
)

func TestNewRender(t *testing.T) {
	table := New()

	table.Add("KEY", "VAL")
	table.Add("One", "1")

	returned := table.Render()
	expected := "KEY   VAL\nOne     1\n"

	if returned != expected {
		t.Errorf("Failed to render properly")
	}
}

func TestNewPrint(t *testing.T) {
	table := New()

	table.Add("KEY", "VAL")
	table.Add("One", nil)

	text := table.Render()
	t.Logf("Rendered: %s\n", text)
	table.Print()
}
