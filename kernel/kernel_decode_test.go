package iterkernel

import (
	"math"

	"testing"
)

func TestKernelDecodeWithNilFuncIntoFloat32(t *testing.T) {

	tests := []struct {
		Src interface{}
		Expected float32
	}{
		{
			Src:      float32(-math.MaxFloat32),
			Expected:         -math.MaxFloat32,
		},
		{
			Src:      float32(-1.0),
			Expected:         -1.0,
		},
		{
			Src:      float32(-math.SmallestNonzeroFloat32),
			Expected:         -math.SmallestNonzeroFloat32,
		},
		{
			Src:      float32(0.0),
			Expected:         0.0,
		},
		{
			Src:      float32(math.SmallestNonzeroFloat32),
			Expected:         math.SmallestNonzeroFloat32,
		},
		{
			Src:      float32(1.0),
			Expected:         1.0,
		},
		{
			Src:      float32(math.MaxFloat32),
			Expected:         math.MaxFloat32,
		},



		{
			Src: int8(math.MinInt8),
			Expected: math.MinInt8,
		},
		{
			Src: int8(-123),
			Expected: -123,
		},
		{
			Src: int8(-1),
			Expected: -1,
		},
		{
			Src: int8(0),
			Expected: 0,
		},
		{
			Src: int8(1),
			Expected: 1,
		},
		{
			Src: int8(123),
			Expected: 123,
		},
		{
			Src: int8(math.MaxInt8),
			Expected: math.MaxInt8,
		},



		{
			Src: int16(math.MinInt16),
			Expected:  math.MinInt16,
		},
		{
			Src: int16(-123),
			Expected:  -123,
		},
		{
			Src: int16(-1),
			Expected:  -1,
		},
		{
			Src: int16(0),
			Expected:  0,
		},
		{
			Src: int16(1),
			Expected:  1,
		},
		{
			Src: int16(123),
			Expected:  123,
		},
		{
			Src: int16(math.MaxInt16),
			Expected:  math.MaxInt16,
		},



		{
			Src: uint8(0),
			Expected:  0,
		},
		{
			Src: uint8(1),
			Expected:  1,
		},
		{
			Src: uint8(123),
			Expected:  123,
		},
		{
			Src: uint8(math.MaxUint8),
			Expected:  math.MaxUint8,
		},



		{
			Src: uint16(0),
			Expected:   0,
		},
		{
			Src: uint16(1),
			Expected:   1,
		},
		{
			Src: uint16(123),
			Expected:   123,
		},
		{
			Src: uint16(math.MaxUint16),
			Expected:   math.MaxUint16,
		},



		{
			Src: uint32(0),
			Expected:   0,
		},
		{
			Src: uint32(1),
			Expected:   1,
		},
		{
			Src: uint32(123),
			Expected:   123,
		},
		{
			Src: uint32(math.MaxUint32),
			Expected:   math.MaxUint32,
		},
	}


	for testNumber, test := range tests {

		var kernel Kernel

		kernel.datum = test.Src

		var target float32

		if err := kernel.KernelDecode(nil, &target); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, target; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
	}
}

func TestKernelDecodeWithNilFuncIntoFloat64(t *testing.T) {

	tests := []struct {
		Src interface{}
		Expected float64
	}{
		{
			Src:      float32(-math.MaxFloat32),
			Expected:         -math.MaxFloat32,
		},
		{
			Src:      float32(-1.0),
			Expected:         -1.0,
		},
		{
			Src:      float32(-math.SmallestNonzeroFloat32),
			Expected:         -math.SmallestNonzeroFloat32,
		},
		{
			Src:      float32(0.0),
			Expected:         0.0,
		},
		{
			Src:      float32(math.SmallestNonzeroFloat32),
			Expected:         math.SmallestNonzeroFloat32,
		},
		{
			Src:      float32(1.0),
			Expected:         1.0,
		},
		{
			Src:      float32(math.MaxFloat32),
			Expected:         math.MaxFloat32,
		},



		{
			Src:      float64(-math.MaxFloat64),
			Expected:         -math.MaxFloat64,
		},
		{
			Src:      float64(-1.0),
			Expected:         -1.0,
		},
		{
			Src:      float64(-math.SmallestNonzeroFloat64),
			Expected:         -math.SmallestNonzeroFloat64,
		},
		{
			Src:      float64(0.0),
			Expected:         0.0,
		},
		{
			Src:      float64(math.SmallestNonzeroFloat64),
			Expected:         math.SmallestNonzeroFloat64,
		},
		{
			Src:      float64(1.0),
			Expected:         1.0,
		},
		{
			Src:      float64(math.MaxFloat64),
			Expected:         math.MaxFloat64,
		},



		{
			Src: int8(math.MinInt8),
			Expected: math.MinInt8,
		},
		{
			Src: int8(-123),
			Expected: -123,
		},
		{
			Src: int8(-1),
			Expected: -1,
		},
		{
			Src: int8(0),
			Expected: 0,
		},
		{
			Src: int8(1),
			Expected: 1,
		},
		{
			Src: int8(123),
			Expected: 123,
		},
		{
			Src: int8(math.MaxInt8),
			Expected: math.MaxInt8,
		},



		{
			Src: int16(math.MinInt16),
			Expected:  math.MinInt16,
		},
		{
			Src: int16(-123),
			Expected:  -123,
		},
		{
			Src: int16(-1),
			Expected:  -1,
		},
		{
			Src: int16(0),
			Expected:  0,
		},
		{
			Src: int16(1),
			Expected:  1,
		},
		{
			Src: int16(123),
			Expected:  123,
		},
		{
			Src: int16(math.MaxInt16),
			Expected:  math.MaxInt16,
		},



		{
			Src: uint8(0),
			Expected:  0,
		},
		{
			Src: uint8(1),
			Expected:  1,
		},
		{
			Src: uint8(123),
			Expected:  123,
		},
		{
			Src: uint8(math.MaxUint8),
			Expected:  math.MaxUint8,
		},



		{
			Src: uint16(0),
			Expected:   0,
		},
		{
			Src: uint16(1),
			Expected:   1,
		},
		{
			Src: uint16(123),
			Expected:   123,
		},
		{
			Src: uint16(math.MaxUint16),
			Expected:   math.MaxUint16,
		},



		{
			Src: uint32(0),
			Expected:   0,
		},
		{
			Src: uint32(1),
			Expected:   1,
		},
		{
			Src: uint32(123),
			Expected:   123,
		},
		{
			Src: uint32(math.MaxUint32),
			Expected:   math.MaxUint32,
		},
	}


	for testNumber, test := range tests {

		var kernel Kernel

		kernel.datum = test.Src

		var target float64

		if err := kernel.KernelDecode(nil, &target); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, target; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
	}
}

func TestKernelDecodeWithNilFuncIntoInt8(t *testing.T) {

	tests := []struct {
		Src interface{}
		Expected int8
	}{
		{
			Src: int8(math.MinInt8),
			Expected: math.MinInt8,
		},
		{
			Src: int8(-123),
			Expected: -123,
		},
		{
			Src: int8(-1),
			Expected: -1,
		},
		{
			Src: int8(0),
			Expected: 0,
		},
		{
			Src: int8(1),
			Expected: 1,
		},
		{
			Src: int8(123),
			Expected: 123,
		},
		{
			Src: int8(math.MaxInt8),
			Expected: math.MaxInt8,
		},
	}


	for testNumber, test := range tests {

		var kernel Kernel

		kernel.datum = test.Src

		var target int8

		if err := kernel.KernelDecode(nil, &target); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, target; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
	}
}

func TestKernelDecodeWithNilFuncIntoInt16(t *testing.T) {

	tests := []struct {
		Src interface{}
		Expected int16
	}{
		{
			Src: int8(math.MinInt8),
			Expected: math.MinInt8,
		},
		{
			Src: int8(-123),
			Expected: -123,
		},
		{
			Src: int8(-1),
			Expected: -1,
		},
		{
			Src: int8(0),
			Expected: 0,
		},
		{
			Src: int8(1),
			Expected: 1,
		},
		{
			Src: int8(123),
			Expected: 123,
		},
		{
			Src: int8(math.MaxInt8),
			Expected: math.MaxInt8,
		},



		{
			Src: int16(math.MinInt16),
			Expected:  math.MinInt16,
		},
		{
			Src: int16(-123),
			Expected:  -123,
		},
		{
			Src: int16(-1),
			Expected:  -1,
		},
		{
			Src: int16(0),
			Expected:  0,
		},
		{
			Src: int16(1),
			Expected:  1,
		},
		{
			Src: int16(123),
			Expected:  123,
		},
		{
			Src: int16(math.MaxInt16),
			Expected:  math.MaxInt16,
		},



		{
			Src: uint8(0),
			Expected:  0,
		},
		{
			Src: uint8(1),
			Expected:  1,
		},
		{
			Src: uint8(123),
			Expected:  123,
		},
		{
			Src: uint8(math.MaxUint8),
			Expected:  math.MaxUint8,
		},
	}


	for testNumber, test := range tests {

		var kernel Kernel

		kernel.datum = test.Src

		var target int16

		if err := kernel.KernelDecode(nil, &target); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, target; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
	}
}

func TestKernelDecodeWithNilFuncIntoInt32(t *testing.T) {

	tests := []struct {
		Src interface{}
		Expected int32
	}{
		{
			Src: int8(math.MinInt8),
			Expected: math.MinInt8,
		},
		{
			Src: int8(-123),
			Expected: -123,
		},
		{
			Src: int8(-1),
			Expected: -1,
		},
		{
			Src: int8(0),
			Expected: 0,
		},
		{
			Src: int8(1),
			Expected: 1,
		},
		{
			Src: int8(123),
			Expected: 123,
		},
		{
			Src: int8(math.MaxInt8),
			Expected: math.MaxInt8,
		},



		{
			Src: int16(math.MinInt16),
			Expected:  math.MinInt16,
		},
		{
			Src: int16(-123),
			Expected:  -123,
		},
		{
			Src: int16(-1),
			Expected:  -1,
		},
		{
			Src: int16(0),
			Expected:  0,
		},
		{
			Src: int16(1),
			Expected:  1,
		},
		{
			Src: int16(123),
			Expected:  123,
		},
		{
			Src: int16(math.MaxInt16),
			Expected:  math.MaxInt16,
		},



		{
			Src: int32(math.MinInt32),
			Expected:  math.MinInt32,
		},
		{
			Src: int32(-123),
			Expected:  -123,
		},
		{
			Src: int32(-1),
			Expected:  -1,
		},
		{
			Src: int32(0),
			Expected:  0,
		},
		{
			Src: int32(1),
			Expected:  1,
		},
		{
			Src: int32(123),
			Expected:  123,
		},
		{
			Src: int32(math.MaxInt32),
			Expected:  math.MaxInt32,
		},



		{
			Src: uint8(0),
			Expected:  0,
		},
		{
			Src: uint8(1),
			Expected:  1,
		},
		{
			Src: uint8(123),
			Expected:  123,
		},
		{
			Src: uint8(math.MaxUint8),
			Expected:  math.MaxUint8,
		},



		{
			Src: uint16(0),
			Expected:   0,
		},
		{
			Src: uint16(1),
			Expected:   1,
		},
		{
			Src: uint16(123),
			Expected:   123,
		},
		{
			Src: uint16(math.MaxUint16),
			Expected:   math.MaxUint16,
		},
	}


	for testNumber, test := range tests {

		var kernel Kernel

		kernel.datum = test.Src

		var target int32

		if err := kernel.KernelDecode(nil, &target); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, target; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
	}
}

func TestKernelDecodeWithNilFuncIntoInt64(t *testing.T) {

	tests := []struct {
		Src interface{}
		Expected int64
	}{
		{
			Src: int8(math.MinInt8),
			Expected: math.MinInt8,
		},
		{
			Src: int8(-123),
			Expected: -123,
		},
		{
			Src: int8(-1),
			Expected: -1,
		},
		{
			Src: int8(0),
			Expected: 0,
		},
		{
			Src: int8(1),
			Expected: 1,
		},
		{
			Src: int8(123),
			Expected: 123,
		},
		{
			Src: int8(math.MaxInt8),
			Expected: math.MaxInt8,
		},



		{
			Src: int16(math.MinInt16),
			Expected:  math.MinInt16,
		},
		{
			Src: int16(-123),
			Expected:  -123,
		},
		{
			Src: int16(-1),
			Expected:  -1,
		},
		{
			Src: int16(0),
			Expected:  0,
		},
		{
			Src: int16(1),
			Expected:  1,
		},
		{
			Src: int16(123),
			Expected:  123,
		},
		{
			Src: int16(math.MaxInt16),
			Expected:  math.MaxInt16,
		},



		{
			Src: int32(math.MinInt32),
			Expected:  math.MinInt32,
		},
		{
			Src: int32(-123),
			Expected:  -123,
		},
		{
			Src: int32(-1),
			Expected:  -1,
		},
		{
			Src: int32(0),
			Expected:  0,
		},
		{
			Src: int32(1),
			Expected:  1,
		},
		{
			Src: int32(123),
			Expected:  123,
		},
		{
			Src: int32(math.MaxInt32),
			Expected:  math.MaxInt32,
		},



		{
			Src: int64(math.MinInt64),
			Expected:  math.MinInt64,
		},
		{
			Src: int64(-123),
			Expected:  -123,
		},
		{
			Src: int64(-1),
			Expected:  -1,
		},
		{
			Src: int64(0),
			Expected:  0,
		},
		{
			Src: int64(1),
			Expected:  1,
		},
		{
			Src: int64(123),
			Expected:  123,
		},
		{
			Src: int64(math.MaxInt64),
			Expected:  math.MaxInt64,
		},



		{
			Src: uint8(0),
			Expected:  0,
		},
		{
			Src: uint8(1),
			Expected:  1,
		},
		{
			Src: uint8(123),
			Expected:  123,
		},
		{
			Src: uint8(math.MaxUint8),
			Expected:  math.MaxUint8,
		},



		{
			Src: uint16(0),
			Expected:   0,
		},
		{
			Src: uint16(1),
			Expected:   1,
		},
		{
			Src: uint16(123),
			Expected:   123,
		},
		{
			Src: uint16(math.MaxUint16),
			Expected:   math.MaxUint16,
		},



		{
			Src: uint32(0),
			Expected:   0,
		},
		{
			Src: uint32(1),
			Expected:   1,
		},
		{
			Src: uint32(123),
			Expected:   123,
		},
		{
			Src: uint32(math.MaxUint32),
			Expected:   math.MaxUint32,
		},
	}


	for testNumber, test := range tests {

		var kernel Kernel

		kernel.datum = test.Src

		var target int64

		if err := kernel.KernelDecode(nil, &target); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, target; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
	}
}

func TestKernelDecodeWithNilFuncIntoUint8(t *testing.T) {

	tests := []struct {
		Src interface{}
		Expected uint8
	}{
		{
			Src: uint8(0),
			Expected:  0,
		},
		{
			Src: uint8(1),
			Expected:  1,
		},
		{
			Src: uint8(123),
			Expected:  123,
		},
		{
			Src: uint8(math.MaxUint8),
			Expected:  math.MaxUint8,
		},
	}


	for testNumber, test := range tests {

		var kernel Kernel

		kernel.datum = test.Src

		var target uint8

		if err := kernel.KernelDecode(nil, &target); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, target; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
	}
}

func TestKernelDecodeWithNilFuncIntoUint16(t *testing.T) {

	tests := []struct {
		Src interface{}
		Expected uint16
	}{
		{
			Src: uint8(0),
			Expected:  0,
		},
		{
			Src: uint8(1),
			Expected:  1,
		},
		{
			Src: uint8(123),
			Expected:  123,
		},
		{
			Src: uint8(math.MaxUint8),
			Expected:  math.MaxUint8,
		},



		{
			Src: uint16(0),
			Expected:   0,
		},
		{
			Src: uint16(1),
			Expected:   1,
		},
		{
			Src: uint16(123),
			Expected:   123,
		},
		{
			Src: uint16(math.MaxUint16),
			Expected:   math.MaxUint16,
		},
	}


	for testNumber, test := range tests {

		var kernel Kernel

		kernel.datum = test.Src

		var target uint16

		if err := kernel.KernelDecode(nil, &target); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, target; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
	}
}

func TestKernelDecodeWithNilFuncIntoUint32(t *testing.T) {

	tests := []struct {
		Src interface{}
		Expected uint32
	}{
		{
			Src: uint8(0),
			Expected:  0,
		},
		{
			Src: uint8(1),
			Expected:  1,
		},
		{
			Src: uint8(123),
			Expected:  123,
		},
		{
			Src: uint8(math.MaxUint8),
			Expected:  math.MaxUint8,
		},



		{
			Src: uint16(0),
			Expected:   0,
		},
		{
			Src: uint16(1),
			Expected:   1,
		},
		{
			Src: uint16(123),
			Expected:   123,
		},
		{
			Src: uint16(math.MaxUint16),
			Expected:   math.MaxUint16,
		},



		{
			Src: uint32(0),
			Expected:   0,
		},
		{
			Src: uint32(1),
			Expected:   1,
		},
		{
			Src: uint32(123),
			Expected:   123,
		},
		{
			Src: uint32(math.MaxUint32),
			Expected:   math.MaxUint32,
		},
	}


	for testNumber, test := range tests {

		var kernel Kernel

		kernel.datum = test.Src

		var target uint32

		if err := kernel.KernelDecode(nil, &target); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, target; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
	}
}

func TestKernelDecodeWithNilFuncIntoUint64(t *testing.T) {

	tests := []struct {
		Src interface{}
		Expected uint64
	}{
		{
			Src: uint8(0),
			Expected:  0,
		},
		{
			Src: uint8(1),
			Expected:  1,
		},
		{
			Src: uint8(123),
			Expected:  123,
		},
		{
			Src: uint8(math.MaxUint8),
			Expected:  math.MaxUint8,
		},



		{
			Src: uint16(0),
			Expected:   0,
		},
		{
			Src: uint16(1),
			Expected:   1,
		},
		{
			Src: uint16(123),
			Expected:   123,
		},
		{
			Src: uint16(math.MaxUint16),
			Expected:   math.MaxUint16,
		},



		{
			Src: uint32(0),
			Expected:   0,
		},
		{
			Src: uint32(1),
			Expected:   1,
		},
		{
			Src: uint32(123),
			Expected:   123,
		},
		{
			Src: uint32(math.MaxUint32),
			Expected:   math.MaxUint32,
		},



		{
			Src: uint64(0),
			Expected:   0,
		},
		{
			Src: uint64(1),
			Expected:   1,
		},
		{
			Src: uint64(123),
			Expected:   123,
		},
		{
			Src: uint64(math.MaxUint64),
			Expected:   math.MaxUint64,
		},
	}


	for testNumber, test := range tests {

		var kernel Kernel

		kernel.datum = test.Src

		var target uint64

		if err := kernel.KernelDecode(nil, &target); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, target; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
	}
}
