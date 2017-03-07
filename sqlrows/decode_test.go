package itersqlrows

import (
	"time"

	"testing"
)

func TestDecode(t *testing.T) {

	now  := time.Now()
	then := time.Now().Add( -5 * time.Hour )

	colScanner := &mockColumnScanner{
		ColumnNames: []string{
			"fruit_name",
			"name", "age",
			"height",
			"is_female",
			"when_created",
			"file_data",
			"greeting",
			"when_cleared",
			"when_reset",
		},
		ColumnsErr:    nil,
		ScanValues:  []interface{}{
			"cherry",
			"Wanda Doe",
			int64(23),
			float64(5.2),
			true,
			now,
			[]byte("apple banana cherry"),
			"Hello world!",
			nil,
			&then,
		},
	}

	type MyStruct struct {
		FruitName    string    `iter:"fruit_name"`
		PersonName   string    `iter.name:"name"`
		Age          int64     `iter.target.name:"age"`
		Height       float64   `iter:"height"`
		IsFemale     bool      `iter.name:"is_female"`
		WhenCreated  time.Time `iter.target.name:"when_created"`
		File       []byte      `iter:"file_data"`
		Greeting     target    `iter.name:"greeting"`
		WhenCleared *time.Time `iter:"when_cleared"`
		WhenReset   *time.Time `iter:"when_reset"`
	}

	var myStruct MyStruct

	if err := decode(colScanner, &myStruct); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	if expected, actual := "cherry", myStruct.FruitName; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "Wanda Doe", myStruct.PersonName; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := int64(23), myStruct.Age; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := float64(5.2), myStruct.Height; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := true, myStruct.IsFemale; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	{
		i, err := myStruct.Greeting.Interface()
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		s, ok := i.(string)
		if !ok {
			t.Errorf("Expected string, but actually got %t", i)
			return
		}

		if expected, actual := "Hello world!", s; expected != actual {
			t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
			return
		}
	}
}
