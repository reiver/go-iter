package itersqlrows

import (
	"github.com/reiver/go-cast"

	"database/sql"
	"errors"
	"fmt"
	"time"
)

type mockColumnScanner struct {
	ColumnNames []string
	ColumnsErr    error
	ScanValues []interface{}
}

func (receiver *mockColumnScanner) Columns() ([]string, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	return receiver.ColumnNames, receiver.ColumnsErr
}

func (receiver *mockColumnScanner) Scan(dest ...interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}


	columnNames := receiver.ColumnNames
	scanValues  := receiver.ScanValues

	if len(columnNames) != len(scanValues) {
		return errors.New("The length of the scan values slice is not the same as the length of the column names. This is a problem with how the mock was set up.")
	}


	if nil == dest {
		return errors.New("The destination slice to Scan is nil.")
	}


	if len(columnNames) != len(dest) {
		return errors.New("The length of thje destination slice is not the same as the length of the column names.")
	}


	for i, columnName := range columnNames {

		target := dest[i]
		value  := scanValues[i]

		switch t := target.(type) {
		case sql.Scanner:
			if err := t.Scan(value); nil != err {
				return fmt.Errorf("Problem scanning for column #%d (column name %q) into sql.Scanner: (%T) %v", i, columnName, err, err)
			}
		case *bool:
			x, err := cast.Bool(value)
			if nil != err {
				return fmt.Errorf("Problem scanning for column #%d (column name %q) into string from type %T: (%T) %v", i, columnName, value, err, err)
			}

			*t = x
		case *[]byte:
			x, err := cast.String(value) //@TODO: Should be a cast.Bytes()
			if nil != err {
				return fmt.Errorf("Problem scanning for column #%d (column name %q) into string from type %T: (%T) %v", i, columnName, value, err, err)
			}

			*t = []byte(x) //@TODO: Should be a cast.Bytes()
		case *float64:
			x, err := cast.Float64(value)
			if nil != err {
				return fmt.Errorf("Problem scanning for column #%d (column name %q) into string from type %T: (%T) %v", i, columnName, value, err, err)
			}

			*t = x
		case *int64:
			x, err := cast.Int64(value)
			if nil != err {
				return fmt.Errorf("Problem scanning for column #%d (column name %q) into string from type %T: (%T) %v", i, columnName, value, err, err)
			}

			*t = x
		case *string:
			x, err := cast.String(value)
			if nil != err {
				return fmt.Errorf("Problem scanning for column #%d (column name %q) into string from type %T: (%T) %v", i, columnName, value, err, err)
			}

			*t = x
		case *time.Time:
			x, err := cast.Time(value)
			if nil != err {
				return fmt.Errorf("Problem scanning for column #%d (column name %q) into string from type %T: (%T) %v", i, columnName, value, err, err)
			}

			*t = x
		default:
			return fmt.Errorf("Problem with value to scan in #%d (column name %q), unscannable type: %T. This is a problem with how the mock was set up.", i, columnName, value)
		}
	}

	return nil
}
