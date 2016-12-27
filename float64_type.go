package iter

import (
	"reflect"
)

func (receiver *Float64) Type() reflect.Type {
	return reflect.TypeOf((*float64)(nil)).Elem()
}
