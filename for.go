package iter

import (
	"fmt"
	"reflect"
)

// For is useful for iterating through an iterator in a terse way.
//
// Example Usage
//
//	iterator := ...
//	
//	err := iter.For{iterator}.Each(func(datum string){
//		//@TODO
//	})
//
// Or:
//
//	type MyStruct {
//		GiveName   string `iter:"given_name"`
//		FamilyName string `iter:"family_name"`
//		Age        int64  `iter:"age"`
//	}
//	
//	iterator := ...
//	
//	err := iter.For{iterator}.Each(func(datum MyStruct){
//		//@TODO
//	})
//
// Etc.
type For struct {
	Iterator Iterator
}

func (receiver For) Each(fn interface{}) error {
	return receiver.each(fn)
}

func (receiver For) each(fn interface{}) (err error) {

	iterator := receiver.Iterator
	if nil == iterator {
		return errNilIterator
	}


	reflectedValue := reflect.ValueOf(fn)
	reflectedType  := reflectedValue.Type()
	if nil == reflectedType {
		return errInternalError
	}
	if reflect.Func != reflectedType.Kind() {
		return errNotFunc
	}
	if expected, actual := 1, reflectedType.NumIn(); expected != actual {
		return fmt.Errorf("Expected func passed as parameter to Each to have %d inputs, but actually has %d.", expected, actual)
	}
	if expected, actual := 0, reflectedType.NumOut(); expected != actual {
		return fmt.Errorf("Expected func passed as parameter to Each to have %d outputs, but actually has %d.", expected, actual)
	}

	paramType := reflectedType.In(0)
	if nil == paramType {
		return errInternalError
	}


	defer func(){
		if e := iterator.Close(); nil != e {
			if nil == err {
				err = e
			}
		}
	}()

	for iterator.Next() {

		paramValue := reflect.New(paramType)

		{
			var in [1]reflect.Value
			in[0] = paramValue

			out := reflect.ValueOf(iterator.Decode).Call(in[:])
			if expected, actual := 1, len(out); expected != actual {
				return fmt.Errorf("Expected call to iterator's Decode method to return %d outputs, but actually returned %d.", expected, actual)
			}
			out0 := out[0]
			i := out0.Interface()
			if nil != i {
				err, ok := i.(error)
				if !ok {
					return fmt.Errorf("Expected call to iterator's Decode method to return output of type error, but actually returned %T.", i)
				}
				if nil != err {
					return err
				}
			}
		}

		{
			var in [1]reflect.Value
			in[0] = paramValue.Elem()

			reflectedValue.Call(in[:])
		}
	}
	if err := iterator.Err(); nil != err {
		return err
	}

	return nil
}
