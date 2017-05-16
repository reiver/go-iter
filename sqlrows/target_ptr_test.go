package itersqlrows

import (
	"math"
	"math/big"
	"time"

	"testing"
)

func TestTargetBoolPtr(t *testing.T) {

	boolFalse := false
	boolTrue  := true

	tests := []struct{
		Value *bool
	}{
		{
			Value: nil,
		},
		{
			Value: &boolFalse,
		},
		{
			Value: &boolTrue,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *bool:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := *test.Value, *casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetFloat32Ptr(t *testing.T) {

	negativeMax             := float32(-math.MaxFloat32)
	negativeOne             := float32(-1.0)
	negativeSmallestNonzero := float32(-math.SmallestNonzeroFloat32)
	zero                    := float32(0.0)
	smallestNonzero         := float32(math.SmallestNonzeroFloat32)
	one                     := float32(1.0)
	max                     := float32(math.MaxFloat32)

	tests := []struct{
		Value *float32
	}{
		{
			Value: nil,
		},
		{
			Value: &negativeMax,
		},
		{
			Value: &negativeOne,
		},
		{
			Value: &negativeSmallestNonzero,
		},
		{
			Value: &zero,
		},
		{
			Value: &smallestNonzero,
		},
		{
			Value: &one,
		},
		{
			Value: &max,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *float32:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := *test.Value, *casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetFloat64Ptr(t *testing.T) {

	negativeMax             := -math.MaxFloat64
	negativeOne             := -1.0
	negativeSmallestNonzero := -math.SmallestNonzeroFloat64
	zero                    := 0.0
	smallestNonzero         := math.SmallestNonzeroFloat64
	one                     := 1.0
	max                     := math.MaxFloat64

	tests := []struct{
		Value *float64
	}{
		{
			Value: nil,
		},
		{
			Value: &negativeMax,
		},
		{
			Value: &negativeOne,
		},
		{
			Value: &negativeSmallestNonzero,
		},
		{
			Value: &zero,
		},
		{
			Value: &smallestNonzero,
		},
		{
			Value: &one,
		},
		{
			Value: &max,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *float64:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := *test.Value, *casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetIntPtr(t *testing.T) {

	min         := int(math.MinInt32)
	negativeOne := int(-1)
	zero        := int(0)
	one         := int(1)
	max         := int(math.MaxInt32)

	tests := []struct{
		Value *int
	}{
		{
			Value: nil,
		},
		{
			Value: &min,
		},
		{
			Value: &negativeOne,
		},
		{
			Value: &zero,
		},
		{
			Value: &one,
		},
		{
			Value: &max,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *int:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := *test.Value, *casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetInt8Ptr(t *testing.T) {

	min         := int8(math.MinInt8)
	negativeOne := int8(-1)
	zero        := int8(0)
	one         := int8(1)
	max         := int8(math.MaxInt8)

	tests := []struct{
		Value *int8
	}{
		{
			Value: nil,
		},
		{
			Value: &min,
		},
		{
			Value: &negativeOne,
		},
		{
			Value: &zero,
		},
		{
			Value: &one,
		},
		{
			Value: &max,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *int8:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := *test.Value, *casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetInt16Ptr(t *testing.T) {

	min         := int16(math.MinInt16)
	negativeOne := int16(-1)
	zero        := int16(0)
	one         := int16(1)
	max         := int16(math.MaxInt16)

	tests := []struct{
		Value *int16
	}{
		{
			Value: nil,
		},
		{
			Value: &min,
		},
		{
			Value: &negativeOne,
		},
		{
			Value: &zero,
		},
		{
			Value: &one,
		},
		{
			Value: &max,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *int16:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := *test.Value, *casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetInt32Ptr(t *testing.T) {

	min         := int32(math.MinInt32)
	negativeOne := int32(-1)
	zero        := int32(0)
	one         := int32(1)
	max         := int32(math.MaxInt32)

	tests := []struct{
		Value *int32
	}{
		{
			Value: nil,
		},
		{
			Value: &min,
		},
		{
			Value: &negativeOne,
		},
		{
			Value: &zero,
		},
		{
			Value: &one,
		},
		{
			Value: &max,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *int32:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := *test.Value, *casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetInt64Ptr(t *testing.T) {

	min         := int64(math.MinInt64)
	negativeOne := int64(-1)
	zero        := int64(0)
	one         := int64(1)
	max         := int64(math.MaxInt64)

	tests := []struct{
		Value *int64
	}{
		{
			Value: nil,
		},
		{
			Value: &min,
		},
		{
			Value: &negativeOne,
		},
		{
			Value: &zero,
		},
		{
			Value: &one,
		},
		{
			Value: &max,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *int64:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := *test.Value, *casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetStringPtr(t *testing.T) {

	apple      := "apple"
	banana     := "banana"
	cherry     := "cherry"
	date       := "date"
	helloWorld := "Hello world"

	tests := []struct{
		Value *string
	}{
		{
			Value: nil,
		},
		{
			Value: &apple,
		},
		{
			Value: &banana,
		},
		{
			Value: &cherry,
		},
		{
			Value: &date,
		},
		{
			Value: &helloWorld,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *string:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := *test.Value, *casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetTimePtr(t *testing.T) {

	t0 := time.Now()
	t1 := time.Now().Add(1 * time.Millisecond)
	t2 := time.Now().Add(2 * time.Second)
	t3 := time.Now().Add(3 * time.Minute)
	t4 := time.Now().Add(4 * time.Hour)

	tests := []struct{
		Value *time.Time
	}{
		{
			Value: nil,
		},
		{
			&t0,
		},
		{
			Value: &t1,
		},
		{
			Value: &t2,
		},
		{
			Value: &t3,
		},
		{
			Value: &t4,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *time.Time:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := *test.Value, *casted; !expected.Equal(actual) {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetMathBigFloatPtr(t *testing.T) {

	negativeMax             := big.NewFloat( float64(-math.MaxFloat64) )
	negativeOne             := big.NewFloat( float64(-1.0) )
	negativeSmallestNonzero := big.NewFloat( float64(-math.SmallestNonzeroFloat64) )
	zero                    := big.NewFloat( float64(0.0) )
	smallestNonzero         := big.NewFloat( float64(math.SmallestNonzeroFloat64) )
	one                     := big.NewFloat( float64(1.0) )
	max                     := big.NewFloat( float64(math.MaxFloat64) )

	tests := []struct{
		Value *big.Float
	}{
		{
			Value: nil,
		},
		{
			Value: negativeMax,
		},
		{
			Value: negativeOne,
		},
		{
			Value: negativeSmallestNonzero,
		},
		{
			Value: zero,
		},
		{
			Value: smallestNonzero,
		},
		{
			Value: one,
		},
		{
			Value: max,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *big.Float:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := test.Value, casted; 0 != expected.Cmp(actual) {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetMathBigRatPtr(t *testing.T) {

	negativeTwo22           := big.NewRat(-2, 22)
	negativeOne             := big.NewRat(-1, 1)
	negativeOneHalf         := big.NewRat(1, 2)
	zero                    := big.NewRat(0, 1)
	oneHalf                 := big.NewRat(1, 2)
	one                     := big.NewRat(1, 1)
	two22                   := big.NewRat(2, 22)

	tests := []struct{
		Value *big.Rat
	}{
		{
			Value: nil,
		},
		{
			Value: negativeTwo22,
		},
		{
			Value: negativeOne,
		},
		{
			Value: negativeOneHalf,
		},
		{
			Value: zero,
		},
		{
			Value: oneHalf,
		},
		{
			Value: one,
		},
		{
			Value: two22,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *big.Rat:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := test.Value, casted; 0 != expected.Cmp(actual) {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetUintPtr(t *testing.T) {

	zero := uint(0)
	one  := uint(1)
	max  := uint(math.MaxUint32)

	tests := []struct{
		Value *uint
	}{
		{
			Value: nil,
		},
		{
			Value: &zero,
		},
		{
			Value: &one,
		},
		{
			Value: &max,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *uint:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := *test.Value, *casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetUint8Ptr(t *testing.T) {

	zero := uint8(0)
	one  := uint8(1)
	max  := uint8(math.MaxUint8)

	tests := []struct{
		Value *uint8
	}{
		{
			Value: nil,
		},
		{
			Value: &zero,
		},
		{
			Value: &one,
		},
		{
			Value: &max,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *uint8:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := *test.Value, *casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetUint16Ptr(t *testing.T) {

	zero := uint16(0)
	one  := uint16(1)
	max  := uint16(math.MaxUint16)

	tests := []struct{
		Value *uint16
	}{
		{
			Value: nil,
		},
		{
			Value: &zero,
		},
		{
			Value: &one,
		},
		{
			Value: &max,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *uint16:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := *test.Value, *casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetUint32Ptr(t *testing.T) {

	zero := uint32(0)
	one  := uint32(1)
	max  := uint32(math.MaxUint32)

	tests := []struct{
		Value *uint32
	}{
		{
			Value: nil,
		},
		{
			Value: &zero,
		},
		{
			Value: &one,
		},
		{
			Value: &max,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *uint32:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := *test.Value, *casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}

func TestTargetUint64Ptr(t *testing.T) {

	zero := uint64(0)
	one  := uint64(1)
	max  := uint64(math.MaxUint64)

	tests := []struct{
		Value *uint64
	}{
		{
			Value: nil,
		},
		{
			Value: &zero,
		},
		{
			Value: &one,
		},
		{
			Value: &max,
		},
	}


	for testNumber, test := range tests {

		var x target

		if err := x.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		i, err := x.Interface()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		switch casted := i.(type) {
		default:
			t.Errorf("For test #%d, unexpected type: %T", testNumber, i)
			continue
		case *uint64:
			if nil != casted {
				if expected, actual := test.Value, casted; expected == actual {
					t.Errorf("For test #%d, did not expect %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
				if expected, actual := *test.Value, *casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			} else {
				if expected, actual := test.Value, casted; expected != actual {
					t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
					continue
				}
			}
		}
	}
}
