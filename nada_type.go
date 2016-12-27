package iter

import (
	"reflect"
)

func (receiver *Nada) Type() reflect.Type {
	return reflect.TypeOf((*struct{})(nil)).Elem()
}
