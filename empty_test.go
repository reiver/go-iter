package iter

import (
	"testing"
)

func TestEmpty(t *testing.T) {

	if nil == Empty {
		t.Errorf("Did not expect iter.Empty to be nil, but actually is: (%T) %v", Empty, Empty)
		return
	}

	for testNumber := 0; testNumber < 10; testNumber++ {

		func(){
			var iterator Iterator = Empty

			defer func(testNumber int) {
				if err := iterator.Close(); nil != err {
					t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
				}
			}(testNumber)
			for iterator.Next() {
				t.Errorf("For test #%d, did not expect an iteration, but actually got one.", testNumber)
				break
			}
			if err := iterator.Err(); nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			}
		}()

		{
			var iterator Iterator = Empty

			err := For{iterator}.Each(func(datum string){
				t.Errorf("For test #%d, did not expect an iteration, but actually got one.", testNumber)
			})
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			}

		}

		{
			var iterator Iterator = Empty

			err := For{iterator}.Each(func(datum int64){
				t.Errorf("For test #%d, did not expect an iteration, but actually got one.", testNumber)
			})
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			}

		}

		{
			var iterator Iterator = Empty

			type MyType struct {
				ID         int64  `iter:"id"`
				GivenName  string `iter:"given_name"`
				FamilyName string `iter:"family_name"`
				Locality   string `iter:"locality"`
			}

			err := For{iterator}.Each(func(datum MyType){
				t.Errorf("For test #%d, did not expect an iteration, but actually got one.", testNumber)
			})
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			}

		}
	}
}
