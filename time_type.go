package iter

import (
	"reflect"
	"time"
)

func (receiver *Time) Type() reflect.Type {
	return reflect.TypeOf((*time.Time)(nil)).Elem()
}
