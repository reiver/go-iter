package iteruint16

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*uint16)(nil)).Elem()
}
