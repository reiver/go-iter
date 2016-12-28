package iteruint8

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*uint8)(nil)).Elem()
}
