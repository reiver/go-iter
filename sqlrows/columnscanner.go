package itersqlrows

type columnScanner interface {
	Columns() ([]string, error)
	Scan(...interface{}) error
}
