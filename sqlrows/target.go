package itersqlrows

import (
	"fmt"
	"time"
)

const (
	targetTypeBool       =  "bool"
	targetTypeBoolPtr    = "*bool"
	targetTypeBytes      = "[]byte"
	targetTypeFloat32    =  "float32"
	targetTypeFloat32Ptr = "*float32"
	targetTypeFloat64    =  "float64"
	targetTypeFloat64Ptr = "*float64"
	targetTypeInt        =  "int"
	targetTypeIntPtr     = "*int"
	targetTypeInt8       =  "int8"
	targetTypeInt8Ptr    = "*int8"
	targetTypeInt16      =  "int16"
	targetTypeInt16Ptr   = "*int16"
	targetTypeInt32      =  "int32"
	targetTypeInt32Ptr   = "*int32"
	targetTypeInt64      =  "int64"
	targetTypeInt64Ptr   = "*int64"
	targetTypeNil        = "nil"
	targetTypeString     =  "string"
	targetTypeStringPtr  = "*string"
	targetTypeTime       =  "time.Time"
	targetTypeTimePtr    = "*time.Time"
	targetTypeUint       =  "uint"
	targetTypeUintPtr    = "*uint"
	targetTypeUint8      =  "uint8"
	targetTypeUint8Ptr   = "*uint8"
	targetTypeUint16     =  "uint16"
	targetTypeUint16Ptr  = "*uint16"
	targetTypeUint32     =  "uint32"
	targetTypeUint32Ptr  = "*uint32"
	targetTypeUint64     =  "uint64"
	targetTypeUint64Ptr  = "*uint64"
)

// target represents something that can be either a bool, []byte, float64, int64, string, or time.Time.
//
// Its usefulness is that it can be used as a sql.Scanner (from the "database/sql" package).
//
// As thus can be a target to receive data.
type target struct {
	typ string

	boolValue        bool
	boolPtrValue    *bool
	bytesValue     []byte
	float32Value     float32
	float32PtrValue *float32
	float64Value     float64
	float64PtrValue *float64
	intValue         int
	intPtrValue     *int
	int8Value        int8
	int8PtrValue    *int8
	int16Value       int16
	int16PtrValue   *int16
	int32Value       int32
	int32PtrValue   *int32
	int64Value       int64
	int64PtrValue   *int64
	nilValue         interface{}
	stringValue      string
	stringPtrValue  *string
	timeValue        time.Time
	timePtrValue    *time.Time
	uintValue        uint
	uintPtrValue    *uint
	uint8Value       uint8
	uint8PtrValue   *uint8
	uint16Value      uint16
	uint16PtrValue  *uint16
	uint32Value      uint32
	uint32PtrValue  *uint32
	uint64Value      uint64
	uint64PtrValue  *uint64
}


func (receiver target) Interface() (interface{}, error) {

	t := receiver.typ
	switch t {
	case targetTypeBool:
		return receiver.boolValue, nil
	case targetTypeBoolPtr:
		return receiver.boolPtrValue, nil
	case targetTypeBytes:
		return receiver.bytesValue, nil
	case targetTypeFloat32:
		return receiver.float32Value, nil
	case targetTypeFloat32Ptr:
		return receiver.float32PtrValue, nil
	case targetTypeFloat64:
		return receiver.float64Value, nil
	case targetTypeFloat64Ptr:
		return receiver.float64PtrValue, nil
	case targetTypeInt:
		return receiver.intValue, nil
	case targetTypeIntPtr:
		return receiver.intPtrValue, nil
	case targetTypeInt8:
		return receiver.int8Value, nil
	case targetTypeInt8Ptr:
		return receiver.int8PtrValue, nil
	case targetTypeInt16:
		return receiver.int16Value, nil
	case targetTypeInt16Ptr:
		return receiver.int16PtrValue, nil
	case targetTypeInt32:
		return receiver.int32Value, nil
	case targetTypeInt32Ptr:
		return receiver.int32PtrValue, nil
	case targetTypeInt64:
		return receiver.int64Value, nil
	case targetTypeInt64Ptr:
		return receiver.int64PtrValue, nil
	case targetTypeNil:
		return receiver.nilValue, nil
	case targetTypeString:
		return receiver.stringValue, nil
	case targetTypeStringPtr:
		return receiver.stringPtrValue, nil
	case targetTypeTime:
		return receiver.timeValue, nil
	case targetTypeTimePtr:
		return receiver.timePtrValue, nil
	case targetTypeUint:
		return receiver.uintValue, nil
	case targetTypeUintPtr:
		return receiver.uintPtrValue, nil
	case targetTypeUint8:
		return receiver.uint8Value, nil
	case targetTypeUint8Ptr:
		return receiver.uint8PtrValue, nil
	case targetTypeUint16:
		return receiver.uint16Value, nil
	case targetTypeUint16Ptr:
		return receiver.uint16PtrValue, nil
	case targetTypeUint32:
		return receiver.uint32Value, nil
	case targetTypeUint32Ptr:
		return receiver.uint32PtrValue, nil
	case targetTypeUint64:
		return receiver.uint64Value, nil
	case targetTypeUint64Ptr:
		return receiver.uint64PtrValue, nil
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
	case *bool:
		receiver.typ = targetTypeBoolPtr
		receiver.boolPtrValue = t
		return nil
	case []byte:
		receiver.typ = targetTypeBytes
		receiver.bytesValue = t
		return nil
	case float32:
		receiver.typ = targetTypeFloat32
		receiver.float32Value = t
		return nil
	case *float32:
		receiver.typ = targetTypeFloat32Ptr
		receiver.float32PtrValue = t
		return nil
	case float64:
		receiver.typ = targetTypeFloat64
		receiver.float64Value = t
		return nil
	case *float64:
		receiver.typ = targetTypeFloat64Ptr
		receiver.float64PtrValue = t
		return nil
	case int:
		receiver.typ = targetTypeInt
		receiver.intValue = t
		return nil
	case *int:
		receiver.typ = targetTypeIntPtr
		receiver.intPtrValue = t
		return nil
	case int8:
		receiver.typ = targetTypeInt8
		receiver.int8Value = t
		return nil
	case *int8:
		receiver.typ = targetTypeInt8Ptr
		receiver.int8PtrValue = t
		return nil
	case int16:
		receiver.typ = targetTypeInt16
		receiver.int16Value = t
		return nil
	case *int16:
		receiver.typ = targetTypeInt16Ptr
		receiver.int16PtrValue = t
		return nil
	case int32:
		receiver.typ = targetTypeInt32
		receiver.int32Value = t
		return nil
	case *int32:
		receiver.typ = targetTypeInt32Ptr
		receiver.int32PtrValue = t
		return nil
	case int64:
		receiver.typ = targetTypeInt64
		receiver.int64Value = t
		return nil
	case *int64:
		receiver.typ = targetTypeInt64Ptr
		receiver.int64PtrValue = t
		return nil
	case string:
		receiver.typ = targetTypeString
		receiver.stringValue = t
		return nil
	case *string:
		receiver.typ = targetTypeStringPtr
		receiver.stringPtrValue = t
		return nil
	case time.Time:
		receiver.typ = targetTypeTime
		receiver.timeValue = t
		return nil
	case *time.Time:
		receiver.typ = targetTypeTimePtr
		receiver.timePtrValue = t
		return nil
	case uint:
		receiver.typ = targetTypeUint
		receiver.uintValue = t
		return nil
	case *uint:
		receiver.typ = targetTypeUintPtr
		receiver.uintPtrValue = t
		return nil
	case uint8:
		receiver.typ = targetTypeUint8
		receiver.uint8Value = t
		return nil
	case *uint8:
		receiver.typ = targetTypeUint8Ptr
		receiver.uint8PtrValue = t
		return nil
	case uint16:
		receiver.typ = targetTypeUint16
		receiver.uint16Value = t
		return nil
	case *uint16:
		receiver.typ = targetTypeUint16Ptr
		receiver.uint16PtrValue = t
		return nil
	case uint32:
		receiver.typ = targetTypeUint32
		receiver.uint32Value = t
		return nil
	case *uint32:
		receiver.typ = targetTypeUint32Ptr
		receiver.uint32PtrValue = t
		return nil
	case uint64:
		receiver.typ = targetTypeUint64
		receiver.uint64Value = t
		return nil
	case *uint64:
		receiver.typ = targetTypeUint64Ptr
		receiver.uint64PtrValue = t
		return nil
	default:
		if nil == t {
			receiver.typ = targetTypeNil
			receiver.nilValue = nil
			return nil
		}
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
	case targetTypeBoolPtr:
		return fmt.Sprintf("%v", receiver.boolPtrValue)
	case targetTypeBytes:
		return fmt.Sprintf("%v", receiver.bytesValue)
	case targetTypeFloat32:
		return fmt.Sprintf("%v", receiver.float32Value)
	case targetTypeFloat32Ptr:
		return fmt.Sprintf("%v", receiver.float32PtrValue)
	case targetTypeFloat64:
		return fmt.Sprintf("%v", receiver.float64Value)
	case targetTypeFloat64Ptr:
		return fmt.Sprintf("%v", receiver.float64PtrValue)
	case targetTypeInt:
		return fmt.Sprintf("%v", receiver.intValue)
	case targetTypeIntPtr:
		return fmt.Sprintf("%v", receiver.intPtrValue)
	case targetTypeInt8:
		return fmt.Sprintf("%v", receiver.int8Value)
	case targetTypeInt8Ptr:
		return fmt.Sprintf("%v", receiver.int8PtrValue)
	case targetTypeInt16:
		return fmt.Sprintf("%v", receiver.int16Value)
	case targetTypeInt16Ptr:
		return fmt.Sprintf("%v", receiver.int16PtrValue)
	case targetTypeInt32:
		return fmt.Sprintf("%v", receiver.int32Value)
	case targetTypeInt32Ptr:
		return fmt.Sprintf("%v", receiver.int32PtrValue)
	case targetTypeInt64:
		return fmt.Sprintf("%v", receiver.int64Value)
	case targetTypeInt64Ptr:
		return fmt.Sprintf("%v", receiver.int64PtrValue)
	case targetTypeNil:
		return fmt.Sprintf("%v", receiver.nilValue)
	case targetTypeString:
		return fmt.Sprintf("%v", receiver.stringValue)
	case targetTypeStringPtr:
		return fmt.Sprintf("%v", receiver.stringPtrValue)
	case targetTypeTime:
		return fmt.Sprintf("%v", receiver.timeValue)
	case targetTypeTimePtr:
		return fmt.Sprintf("%v", receiver.timePtrValue)
	case targetTypeUint:
		return fmt.Sprintf("%v", receiver.uintValue)
	case targetTypeUintPtr:
		return fmt.Sprintf("%v", receiver.uintPtrValue)
	case targetTypeUint8:
		return fmt.Sprintf("%v", receiver.uint8Value)
	case targetTypeUint8Ptr:
		return fmt.Sprintf("%v", receiver.uint8PtrValue)
	case targetTypeUint16:
		return fmt.Sprintf("%v", receiver.uint16Value)
	case targetTypeUint16Ptr:
		return fmt.Sprintf("%v", receiver.uint16PtrValue)
	case targetTypeUint32:
		return fmt.Sprintf("%v", receiver.uint32Value)
	case targetTypeUint32Ptr:
		return fmt.Sprintf("%v", receiver.uint32PtrValue)
	case targetTypeUint64:
		return fmt.Sprintf("%v", receiver.uint64Value)
	case targetTypeUint64Ptr:
		return fmt.Sprintf("%v", receiver.uint64PtrValue)
	default:
		panic(fmt.Errorf("Internal Error: Unsupported Type: %T", receiver.typ))
	}
}
