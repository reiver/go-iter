package itercomplex128

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*complex128)(nil)).Elem()
}
