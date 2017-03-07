package iter

import (
	"github.com/reiver/go-iter/string"

	"testing"
)

func TestForSingle(t *testing.T) {

	const expectedValue = "Hello world!"
	slice := []string{
		expectedValue,
	}

	iterator := &iterstring.Slice{
		Slice: slice,
	}

	var actualValue string
	numIteration := 0
	err := For{iterator}.Each(func(datum string){
		numIteration++

		actualValue = datum
	})
	if nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}
	if expected, actual := len(slice), numIteration; expected != actual {
		t.Errorf("Expected %d, but actually got %d.", expected, actual)
		return
	}
	if expected, actual := expectedValue, actualValue; expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}
}

func TestForMany(t *testing.T) {

	slice := []string{
		"Apple",
		"Banana",
		"Cherry",
		"Date",
	}

	iterator := &iterstring.Slice{
		Slice: slice,
	}

	var actualValues []string
	err := For{iterator}.Each(func(datum string){
		actualValues = append(actualValues, datum)
	})
	if nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}
	if expected, actual := len(slice), len(actualValues); expected != actual {
		t.Errorf("Expected %d, but actually got %d.", expected, actual)
		return
	}
	for valueNumber, expectedValue := range slice {
		actualValue := actualValues[valueNumber]

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("For value #%d, expected %q, but actually got %q.", valueNumber, expected, actual)
			continue
		}
	}
}

func TestForFail(t *testing.T) {

	const expectedValue = "Hello world!"
	slice := []string{
		expectedValue,
	}

	iterator := &iterstring.Slice{
		Slice: slice,
	}

	var actualValue int64
	numIteration := 0
	err := For{iterator}.Each(func(datum int64){
		numIteration++

		actualValue = datum
	})
	if nil == err {
		t.Errorf("Expect an error, but did not actually got one: %v", err)
		return
	}
	if expected, actual := 0, numIteration; expected != actual {
		t.Errorf("Expected %d, but actually got %d.", expected, actual)
		return
	}
}
