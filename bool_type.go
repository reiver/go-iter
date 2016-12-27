package iter

import (
	"reflect"
)

func (receiver *Bool) Type() reflect.Type {
	return reflect.TypeOf((*bool)(nil)).Elem()
}
