package iteruint32

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*uint32)(nil)).Elem()
}
