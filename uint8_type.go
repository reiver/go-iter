package iter

import (
	"reflect"
)

func (receiver *Uint8) Type() reflect.Type {
	return reflect.TypeOf((*uint8)(nil)).Elem()
}
