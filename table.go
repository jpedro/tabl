package tables

type Table struct {
	rows [][]any
}

func (me *Table) Add(values ...any) {
	me.rows = append(me.rows, values)
}

func (me *Table) Print() {
	Print(me.rows)
}
