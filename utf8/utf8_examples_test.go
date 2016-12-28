package iterutf8_test


import (
	"fmt"

	"github.com/reiver/go-iter/utf8"
)


func ExampleSlice() {

	var p []byte = []byte("Hello world! 🙂")

	iterator := iterutf8.Slice{Slice:p}

	for iterator.Next() {
		var r rune

		if err := iterator.Decode(&r); nil != err {
			fmt.Printf("ERROR: (%T) %v \n", err, err)
			return
		}

		fmt.Printf("%q (%d) \n", r, r)
	}
	if err := iterator.Err(); nil != err {
		fmt.Printf("ERROR: (%T) %v \n", err, err)
		return
	}

	// Output:
	// 'H' (72)
	// 'e' (101)
	// 'l' (108)
	// 'l' (108)
	// 'o' (111)
	// ' ' (32)
	// 'w' (119)
	// 'o' (111)
	// 'r' (114)
	// 'l' (108)
	// 'd' (100)
	// '!' (33)
	// ' ' (32)
	// '🙂' (128578)
}


