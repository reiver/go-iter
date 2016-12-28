package iterfloat64x2

import (
	"testing"
)

var (
	float64x2TestSlices = [][][2]float64{
		[][2]float64{},

		[][2]float64{ [2]float64{0.0,0.0} },
		[][2]float64{ [2]float64{1.0,0.1} },
		[][2]float64{ [2]float64{2.0,0.2} },
		[][2]float64{ [2]float64{3.0,0.3} },
		[][2]float64{ [2]float64{4.0,0.4} },
		[][2]float64{ [2]float64{5.0,0.5} },
		[][2]float64{ [2]float64{6.0,0.6} },
		[][2]float64{ [2]float64{7.0,0.7} },
		[][2]float64{ [2]float64{8.0,0.8} },
		[][2]float64{ [2]float64{9.0,0.9} },
		[][2]float64{ [2]float64{10.0,0.01} },

		[][2]float64{
			[2]float64{0.0,0.0},
			[2]float64{1.0,0.1},
			[2]float64{2.0,0.2},
			[2]float64{3.0,0.3},
			[2]float64{4.0,0.4},
			[2]float64{5.0,0.5},
			[2]float64{6.0,0.6},
			[2]float64{7.0,0.7},
			[2]float64{8.0,0.8},
			[2]float64{9.0,0.9},
			[2]float64{10.0,0.01},
		},

		[][2]float64{
			[2]float64{213.202,-12.345},
			[2]float64{18.179,-0.00001},
			[2]float64{4.000002,42.0},
		},
	}
)

func TestSlice(t *testing.T) {

	tests := []struct{
		Slice [][2]float64
	}{}

	for _, slice := range float64x2TestSlices {
		sliceCopy := append([][2]float64(nil), slice...)

		test := struct{
			Slice [][2]float64
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		slice := append([][2]float64(nil), test.Slice...)

		iterator := Slice{
			Slice: slice,
		}

		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		var actualData [][2]float64
		iterationNumber := -1
		for iterator.Next() {
			iterationNumber++


			var datum [2]float64

			if err := iterator.Decode(&datum); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}

			actualData = append(actualData, datum)


			if err := iterator.Decode((*[2]float64)(nil)); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}


			{
				var x interface{}

				if err := iterator.Decode(&x); nil != err {
					t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
					continue
				}

				datum2, ok := x.([2]float64)
				if !ok {
					t.Errorf("For test #%d and iteration #%d, expected to be able to cast, but actually could not. (%T)", testNumber, iterationNumber, x)
					continue
				}

				if expected, actual := datum, datum2; expected != actual {
					t.Errorf("For test #%d and iteration #%d, expected %v, but actually got %v.", testNumber, iterationNumber, expected, actual)
					continue
				}
			}


			if err := iterator.Decode((*interface{})(nil)); nil != err {
				t.Errorf("For test #%d and iteration #%d, did not expect an error, but actually got one: (%T) %v", testNumber, iterationNumber, err, err)
				continue
			}
		}
		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := len(test.Slice), len(actualData); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}
		for datumNumber, expected := range test.Slice {
			actual := actualData[datumNumber]

			if expected != actual {
				t.Errorf("For test #%d and datum #%d, expected %v, but actually got %v.", testNumber, datumNumber, expected, actual)
				continue
			}
		}

		if err := iterator.Close(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
	}
}

func TestSliceClose(t *testing.T) {

	tests := []struct{
		Slice [][2]float64
	}{}

	for _, slice := range float64x2TestSlices {
		sliceCopy := append([][2]float64(nil), slice...)

		test := struct{
			Slice [][2]float64
		}{
			Slice: sliceCopy,
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		for closeTestNumber:=0; closeTestNumber<len(test.Slice); closeTestNumber++ {
			slice := append([][2]float64(nil), test.Slice...)

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

func TestSliceErrNilReceiver(t *testing.T) {

	iterator := (*Slice)(nil)

	{
		err := iterator.Close()
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
		if expected, actual := errNilReceiver, err; expected != actual {
			t.Errorf("Expected an (%T) %q, but actually got (%T) %q.", expected, expected, actual, actual)
			return
		}
	}

	{
		var datum interface{}

		err := iterator.Decode(&datum)
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
		if expected, actual := errNilReceiver, err; expected != actual {
			t.Errorf("Expected an (%T) %q, but actually got (%T) %q.", expected, expected, actual, actual)
			return
		}
	}

	{
		err := iterator.Err()
		if nil == err {
			t.Errorf("Expected an error, but did not actually get one: %v", err)
			return
		}
		if expected, actual := errNilReceiver, err; expected != actual {
			t.Errorf("Expected an (%T) %q, but actually got (%T) %q.", expected, expected, actual, actual)
			return
		}
	}

	{
		actual := iterator.Next()
		if expected := false; expected != actual {
			t.Errorf("Expected an %t, but actually got %t.", expected, actual)
			return
		}
	}
}
