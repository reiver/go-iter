package iter

import (
	"reflect"
)

func (receiver *Slice[T]) Type() reflect.Type {
	return reflect.TypeOf((*string)(nil)).Elem()
}
