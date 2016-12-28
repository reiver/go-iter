package iterstring

import (
	"bytes"

	"testing"
)

func TestSliceWriteTo(t *testing.T) {

	tests := []struct{
		Slice []string
	}{
		{
			Slice: []string(nil),
		},
		{
			Slice: []string{},
		},



		{
			Slice: []string{"apple"},
		},
		{
			Slice: []string{"banana"},
		},
		{
			Slice: []string{"cherry"},
		},



		{
			Slice: []string{"apple", "banana", "cherry"},
		},



		{
			Slice: []string{"apple banana cherry"},
		},



		{
			Slice: []string{"ONE"},
		},
		{
			Slice: []string{"TWO"},
		},
		{
			Slice: []string{"THREE"},
		},



		{
			Slice: []string{"ONE", "TWO", "THREE"},
		},



		{
			Slice: []string{"ONE TWO THREE"},
		},



		{
			Slice: []string{"😐"},
		},



		{
			Slice: []string{"💀", "👽"},
		},



		{
			Slice: []string{"💀👽"},
		},

	}


	for testNumber, test := range tests {

		slice := append([]string(nil), test.Slice...)

		iterator := Slice{
			Slice: slice,
		}

		var buffer bytes.Buffer

		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++

			iterator.WriteTo(&buffer)

			var buffer2 bytes.Buffer
			iterator.WriteTo(&buffer2)
			if expected, actual := len(test.Slice[iterationNumber]), buffer2.Len(); expected != actual {
				t.Errorf("For test #%d and iteration number #%d, expected %d, but actually got %d.", testNumber, iterationNumber, expected, actual)
				continue
			}
			if expected, actual := test.Slice[iterationNumber], buffer2.String(); expected != actual {
				t.Errorf("For test #%d and iteration number #%d, expected %d (%q), but actually got %d (%q).", testNumber, iterationNumber, expected, expected, actual, actual)
				continue
			}
		}
		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v" , testNumber, err, err)
			continue
		}

		{
			expected := 0
			for _, s := range test.Slice {
				expected += len(s)
			}

			if actual := len(buffer.String()); expected != actual {
				t.Errorf("For test #%d, expected %d, but actually got %d." , testNumber, expected, actual)
				continue
			}
		}
		{
			var expectedBuffer bytes.Buffer
			for _, s := range test.Slice {
				expectedBuffer.WriteString(s)
			}

			if expected, actual := expectedBuffer.String(), buffer.String(); expected != actual {
				t.Errorf("For test #%d, expected %q, but actually got %q." , testNumber, expected, actual)
				continue
			}
		}
	}
}
