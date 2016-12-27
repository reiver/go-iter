package iter

import (
	"reflect"
)

func (receiver *UTF8) Type() reflect.Type {
	return reflect.TypeOf((*rune)(nil)).Elem()
}
