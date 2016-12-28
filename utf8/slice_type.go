package iterutf8

import (
	"reflect"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*rune)(nil)).Elem()
}
