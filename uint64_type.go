package iter

import (
	"reflect"
)

func (receiver *Uint64) Type() reflect.Type {
	return reflect.TypeOf((*uint64)(nil)).Elem()
}
