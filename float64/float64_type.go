package iterfloat64

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*float64)(nil)).Elem()
}
