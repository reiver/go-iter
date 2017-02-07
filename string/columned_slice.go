package iterstring

import (
	"reflect"
)

type ColumnedSlice struct {
	Slice
	ColumnName string
}

func (receiver *ColumnedSlice) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.Close()
}

func (receiver *ColumnedSlice) Columns() ([]string, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	return append([]string(nil), receiver.ColumnName), nil
}

func (receiver *ColumnedSlice) Decode(x interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.Decode(x)
}

func (receiver *ColumnedSlice) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.Err()
}

func (receiver *ColumnedSlice) Next() bool {
	if nil == receiver {
		return false
	}

	return receiver.Next()
}

func (receiver *ColumnedSlice) Scan(dest ...interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.Scan(dest...)
}

func (receiver *ColumnedSlice) Type() reflect.Type {
	if nil == receiver {
		return nil
	}

	return receiver.Type()
}

