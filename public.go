package tabl

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
