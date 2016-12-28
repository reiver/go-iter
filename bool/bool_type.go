package iterbool

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*bool)(nil)).Elem()
}
