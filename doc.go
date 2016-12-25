/*
Package iter provides tools for creating iterators.

For example, we can turn a slice into an iterator, with code like the following:

	var slice []string = []string {
		"apple",
		"banana",
		"cherry",
	}
	
	iterator := iter.Strings{
		Slice: slice,
	}
	
	for iterator.Next() {

		var datum string

		if err := iterator.Decode(&datum); nil != err {
			return err
		}

		fmt.Printf("Next datum: %v \n", datum)
	}
	if err := iterator.Err(); nil != err {
		return err
	}

This can help to enable us to write more (run-time) generic code.
*/
package iter
