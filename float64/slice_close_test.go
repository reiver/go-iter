package iterfloat64

import (
	"testing"
)

func TestSliceClose(t *testing.T) {

	tests := []struct{
		Slice []float64
	}{}

	for _, slice := range float64TestSlices {
		sliceCopy := append([]float64(nil), slice...)

		test := struct{
			Slice []float64
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		for closeTestNumber:=0; closeTestNumber<len(test.Slice); closeTestNumber++ {
			slice := append([]float64(nil), test.Slice...)

			iterator := Slice{
				Slice: slice,
			}

			for i:=0; i<closeTestNumber; i++ {
				if expected, actual := true, iterator.Next(); expected != actual {
					t.Errorf("For test #%d and close test #%d, expected %t, but actually got %t.", testNumber, closeTestNumber, expected, actual)
					continue
				}
			}

			if err := iterator.Close(); nil != err {
				t.Errorf("For test #%d, and close test #%d did not expect an error, but actually got one: (%T) %v", testNumber, closeTestNumber, err, err)
				continue
			}

			if expected, actual := false, iterator.Next(); expected != actual {
				t.Errorf("For test #%d and close test #%d, expected %t, but actually got %t.", testNumber, closeTestNumber, expected, actual)
				continue
			}

		}
	}
}
