package iterfloat64

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return receiver.nucleus._type()
}
