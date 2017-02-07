package itersqlrows

import (
	"fmt"
	"time"
)

const (
	targetTypeBool    = "bool"
	targetTypeBytes   = "[]byte"
	targetTypeFloat64 = "float64"
	targetTypeInt64   = "int64"
	targetTypeString  = "string"
	targetTypeTime    = "time.Time"
)

// target represents something that can be either a bool, []byte, float64, int64, string, or time.Time.
//
// Its usefulness is that it can be used as a sql.Scanner (from the "database/sql" package).
//
// As thus can be a target to receive data.
type target struct {
	typ string

	boolValue    bool
	bytesValue []byte
	float64Value float64
	int64Value   int64
	stringValue  string
	timeValue    time.Time
}


func (receiver target) Interface() (interface{}, error) {

	t := receiver.typ
	switch t {
	case targetTypeBool:
		return receiver.boolValue, nil
	case targetTypeBytes:
		return receiver.bytesValue, nil
	case targetTypeFloat64:
		return receiver.float64Value, nil
	case targetTypeInt64:
		return receiver.int64Value, nil
	case targetTypeString:
		return receiver.stringValue, nil
	case targetTypeTime:
		return receiver.timeValue, nil
	default:
		return nil, fmt.Errorf("Unsupported Type: %q", t)
	}
}

func (receiver *target) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch t := src.(type) {
	case bool:
		receiver.typ = targetTypeBool
		receiver.boolValue = t
		return nil
	case []byte:
		receiver.typ = targetTypeBytes
		receiver.bytesValue = t
		return nil
	case float64:
		receiver.typ = targetTypeFloat64
		receiver.float64Value = t
		return nil
	case int64:
		receiver.typ = targetTypeInt64
		receiver.int64Value = t
		return nil
	case string:
		receiver.typ = targetTypeString
		receiver.stringValue = t
		return nil
	case time.Time:
		receiver.typ = targetTypeTime
		receiver.timeValue = t
		return nil
	default:
		return fmt.Errorf("Unsupported Type: %T", t)
	}
}

func (receiver *target) String() string {
	if nil == receiver {
		panic(errNilReceiver)
	}

	switch receiver.typ {
	case targetTypeBool:
		return fmt.Sprintf("%v", receiver.boolValue)
	case targetTypeBytes:
		return fmt.Sprintf("%v", receiver.bytesValue)
	case targetTypeFloat64:
		return fmt.Sprintf("%v", receiver.float64Value)
	case targetTypeInt64:
		return fmt.Sprintf("%v", receiver.int64Value)
	case targetTypeString:
		return fmt.Sprintf("%v", receiver.stringValue)
	case targetTypeTime:
		return fmt.Sprintf("%v", receiver.timeValue)
	default:
		panic(fmt.Errorf("Internal Error: Unsupported Type: %T", receiver.typ))
	}
}
