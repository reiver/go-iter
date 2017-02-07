package itersqlrows

import (
	"bytes"
	"math"
	"time"

	"testing"
)

func TestTargetBool(t *testing.T) {

	tests := []struct{
		Value bool
	}{
		{
			Value: false,
		},
		{
			Value: true,
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
		case bool:
			if expected, actual := test.Value, casted; expected != actual {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestTargetBytes(t *testing.T) {

	tests := []struct{
		Value []byte
	}{
		{
			Value: []byte("apple"),
		},
		{
			Value: []byte("banana"),
		},
		{
			Value: []byte("cherry"),
		},
		{
			Value: []byte("date"),
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
		case []byte:
			if expected, actual := test.Value, casted; !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, expected %v (%q), but actually got %v (%q).", testNumber, expected, string(expected), actual, string(actual))
				continue
			}
		}
	}
}

func TestTargetFloat64(t *testing.T) {

	tests := []struct{
		Value float64
	}{
		{
			Value: -math.MaxFloat64,
		},
		{
			Value: -1.0,
		},
		{
			Value: -math.SmallestNonzeroFloat64,
		},
		{
			Value: 0.0,
		},
		{
			Value: math.SmallestNonzeroFloat64,
		},
		{
			Value: 1.0,
		},
		{
			Value: math.MaxFloat64,
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
		case float64:
			if expected, actual := test.Value, casted; expected != actual {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestTargetInt64(t *testing.T) {

	tests := []struct{
		Value int64
	}{
		{
			Value: math.MinInt64,
		},
		{
			Value: -1,
		},
		{
			Value: 0,
		},
		{
			Value: 1,
		},
		{
			Value: math.MaxInt64,
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
		case int64:
			if expected, actual := test.Value, casted; expected != actual {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestTargetString(t *testing.T) {

	tests := []struct{
		Value string
	}{
		{
			Value: "apple",
		},
		{
			Value: "banana",
		},
		{
			Value: "cherry",
		},
		{
			Value: "date",
		},
		{
			Value: "Hello world!",
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
		case string:
			if expected, actual := test.Value, casted; expected != actual {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}

func TestTargetTime(t *testing.T) {

	tests := []struct{
		Value time.Time
	}{
		{
			Value: time.Now(),
		},
		{
			Value: time.Now().Add(1 * time.Millisecond),
		},
		{
			Value: time.Now().Add(2 * time.Second),
		},
		{
			Value: time.Now().Add(3 * time.Minute),
		},
		{
			Value: time.Now().Add(4 * time.Hour),
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
		case time.Time:
			if expected, actual := test.Value, casted; !expected.Equal(actual) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}
