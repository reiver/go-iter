package iternada

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*struct{})(nil)).Elem()
}
