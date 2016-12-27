package iter

import (
	"reflect"
)

func (receiver *String) Type() reflect.Type {
	return reflect.TypeOf((*string)(nil)).Elem()
}
