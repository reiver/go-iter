/*
Package iter (including its sub-packages) provides tools for creating iterators.

These iterators are intentionally made to resemble *sql.Rows from the "database/sql" package.
Including having the same Close, Err, Next, and Scan methods.

(Note that to turn something into an actual *sql.Rows from the "database/sql" package, instead of just resembling it,
use https://github.com/reiver/go-shunt instead.)

For example, we can turn a slice into an iterator, with code like the following:

	var slice []string = []string {
		"apple",
		"banana",
		"cherry",
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
	case float64:
		//@TODO
	case [2]float64:
		//@TODO
	case string:
		//@TODO
	default:
		//@TODO
	}
*/
package iter
