package iterfloat64

import (
	"reflect"
)

func (receiver *common) _type() reflect.Type {
	return reflect.TypeOf((*float64)(nil)).Elem()
}
