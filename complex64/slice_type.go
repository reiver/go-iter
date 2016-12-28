package itercomplex64

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*complex64)(nil)).Elem()
}
