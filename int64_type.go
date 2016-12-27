package iter

import (
	"reflect"
)

func (receiver *Int64) Type() reflect.Type {
	return reflect.TypeOf((*int64)(nil)).Elem()
}
