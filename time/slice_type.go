package itertime

import (
	"reflect"
	"time"
)

func (receiver *Slice) Type() reflect.Type {
	return reflect.TypeOf((*time.Time)(nil)).Elem()
}
