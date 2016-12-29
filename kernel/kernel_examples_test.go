package iterkernel_test


import (
	"errors"
	"fmt"

	"github.com/reiver/go-iter/kernel"
)

var (
	errInternalError = errors.New("Internal Error")
	errNilReceiver   = errors.New("Nil Receiver")
)

type MyStruct struct {
	Apple  string
	Banana int
	Cherry float64
}

type MyStructIterator struct {
	kernel iterkernel.Kernel
	Slice []MyStruct
}

func (receiver *MyStructIterator) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.kernel.KernelClose()
}

func (receiver *MyStructIterator) Decode(x interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.kernel.KernelDecode(receiver.decode, x)
}

func (receiver *MyStructIterator) decode(x interface{}) (bool, error) {
	if nil == receiver {
		return false, errNilReceiver
	}

	if nil == x {
		return false, nil
	}

	datum, err := receiver.kernel.KernelDatum()
	if nil != err {
		return false, err
	}

	strt, ok := datum.(MyStruct)
	if !ok {
		return false, errInternalError
	}

	switch p := x.(type) {
	case *MyStruct:
		if nil == p {
			return true, nil
		}

		*p = strt

		return true, nil
	default:
		return false, nil
	}
}

func (receiver *MyStructIterator) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.kernel.KernelErr()
}

func (receiver *MyStructIterator) Next() bool {
	if nil == receiver {
		return false
	}

	return receiver.kernel.KernelNext(receiver.next)
}

func (receiver *MyStructIterator) next(index int, v interface{}) (bool, error) {
	if nil == receiver {
		return false, errNilReceiver
	}

	slice := receiver.Slice
	if nil == slice  {
		return false, nil
	}

	if len(slice) <= index {
		return false, nil
	}

	datum := slice[index]

	switch t := v.(type) {
	case *interface{}:
                *t = datum
	default:
		return false, fmt.Errorf("Bad Type: %T", t)
	}

	return true, nil
}

func ExampleSlice() {

	slice := []MyStruct{
		MyStruct{
			Apple: "ONE",
			Banana: 1,
			Cherry: 1.1,
		},
		MyStruct{
			Apple: "TWO",
			Banana: 2,
			Cherry: 2.2,
		},
		MyStruct{
			Apple: "THREE",
			Banana: 3,
			Cherry: 3.3,
		},
	}

	iterator := MyStructIterator{Slice:slice}

	for iterator.Next() {
		var datum MyStruct

		if err := iterator.Decode(&datum); nil != err {
			fmt.Printf("ERROR: (%T) %v \n", err, err)
			return
		}

		fmt.Printf("%v\n", datum)
	}
	if err := iterator.Err(); nil != err {
		fmt.Printf("ERROR: (%T) %v \n", err, err)
		return
	}

	// Output:
	//  {ONE 1 1.1}
	// {TWO 2 2.2}
	// {THREE 3 3.3}
}


