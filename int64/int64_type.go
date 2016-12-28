package iterint64

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*int64)(nil)).Elem()
}
