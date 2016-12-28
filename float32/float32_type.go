package iterfloat32

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*float32)(nil)).Elem()
}
