/*
Package iterstring provides tools for creating iterators for the string type.

These iterators are intentionally made to resemble *sql.Rows from the "database/sql" package.
Including having the same Close, Err, Next, and Scan methods.

(Note that to turn something into an actual *sql.Rows from the "database/sql" package, instead of just resembling it,
use https://github.com/reiver/go-shunt instead.)

For example, we can turn a slice into an iterator, with code like the following:

	var slice []string = []string {
		false,
		true,
		true,
		false,
		true,
	}
	
	iterator := iterstring.Slice{
		Slice: slice,
	}
	
	defer iterator.Close()
	
	for iterator.Next() {
	
		var datum string // Could have also used: var datum interface{}
	
		if err := iterator.Decode(&datum); nil != err {
			return err
		}
	
		fmt.Printf("Next datum: %v \n", datum)
	}
	if err := iterator.Err(); nil != err {
		return err
	}

This can help to enable us to write more (run-time oriented) generic code.

To be able to distinguish one iterator type from another, in a (run-time oriented)
generic way, that is not so specific to this package, we can use the Type method.

For example:

	switch reflect.Zero( iterator.Type() ).Interface().(type) {
	case bool:
		//@TODO
	case byte:
		//@TODO
	case complex64:
		//@TODO
	case complex128:
		//@TODO
	case float32:
		//@TODO
	case float64:
		//@TODO
	case [2]float64:
		//@TODO
	case [3]float64:
		//@TODO
	case [4]float64:
		//@TODO
	case int8:
		//@TODO
	case int16:
		//@TODO
	case int32:
		//@TODO
	case int64:
		//@TODO
	case rune:
		//@TODO
	case string:
		//@TODO
	case struct{}:
		//@TODO
	case time.Time:
		//@TODO
	case uint8:
		//@TODO
	case uint16:
		//@TODO
	case uint32:
		//@TODO
	case uint64:
		//@TODO
	default:
		//@TODO
	}
*/
package iterstring
