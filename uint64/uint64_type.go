package iteruint64

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*uint64)(nil)).Elem()
}
