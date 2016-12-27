package iter

import (
	"reflect"
)

func (receiver *Float32) Type() reflect.Type {
	return reflect.TypeOf((*float32)(nil)).Elem()
}
