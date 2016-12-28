package iterstring

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*string)(nil)).Elem()
}
