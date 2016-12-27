package iter

import (
	"reflect"
)

func (receiver *Complex128) Type() reflect.Type {
	return reflect.TypeOf((*complex128)(nil)).Elem()
}
