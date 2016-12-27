package iter

import (
	"reflect"
)

func (receiver *Uint16) Type() reflect.Type {
	return reflect.TypeOf((*uint16)(nil)).Elem()
}
