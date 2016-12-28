package iterfloat64x2

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*[2]float64)(nil)).Elem()
}
