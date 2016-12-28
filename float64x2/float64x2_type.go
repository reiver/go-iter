package iterfloat64x2

import (
	"reflect"
)

func (receiver *Float64x2) Type() reflect.Type {
	return reflect.TypeOf((*[2]float64)(nil)).Elem()
}
