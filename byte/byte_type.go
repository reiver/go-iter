package iterbyte

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*byte)(nil)).Elem()
}
