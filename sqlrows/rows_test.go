package itersqlrows

import (
	"github.com/reiver/go-iter"

	"database/sql"

	"testing"
)

func TestRowsSqlRows(t *testing.T) {

	var rows *sql.Rows

	var iterator iter.Iterator = &Rows{
		Rows: rows,                 // THIS IS THE LINE THAT ACTUALLY MATTERS.
	}

	if nil == iterator {
		t.Errorf("This should never happen.")
	}
}
