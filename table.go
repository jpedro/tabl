package tabl

type Table struct {
	rows [][]any
	Log bool
}

func (me *Table) Add(values ...any) {
	me.rows = append(me.rows, values)
}

func (me *Table) Render() string {
	return render(me.rows)
}
