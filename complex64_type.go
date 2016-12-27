package iter

import (
	"reflect"
)

func (receiver *Complex64) Type() reflect.Type {
	return reflect.TypeOf((*complex64)(nil)).Elem()
}
