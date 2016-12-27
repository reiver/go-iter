package iter

import (
	"reflect"
)

func (receiver *Byte) Type() reflect.Type {
	return reflect.TypeOf((*byte)(nil)).Elem()
}
