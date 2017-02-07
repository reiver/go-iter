package itersqlrows

import (
	"fmt"
)

func errCouldNotSet(v interface{}, name string, err error) error {
	msg := fmt.Sprintf("Could not set value ([%T] %v) for struct field named %q because: (%T) %v", v, v, name, err, err)

	complainer := internalCouldNotSetComplainer{
		msg: msg,
		err: err,
	}

	return &complainer
}

type CouldNotSetComplainer interface {
	error
	CouldNotSetComplainer()
}

type internalCouldNotSetComplainer struct{
	msg string
	err error
}

func (receiver internalCouldNotSetComplainer) Error() string {
	return receiver.msg
}

func (internalCouldNotSetComplainer) CouldNotSetComplainer() {
	// Nothing here.
}

func (receiver internalCouldNotSetComplainer) Err() error {
	return receiver.err
}
