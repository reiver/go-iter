package iter

import (
	"reflect"
)

func (receiver *Uint32) Type() reflect.Type {
	return reflect.TypeOf((*uint32)(nil)).Elem()
}
