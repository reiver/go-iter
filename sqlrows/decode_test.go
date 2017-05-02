package itersqlrows

import (
	"math/big"
	"time"

	"testing"
)

func TestDecode(t *testing.T) {

	now  := time.Now()
	then := time.Now().Add( -5 * time.Hour )

	later := time.Now().Add( -10 * time.Hour )
	evenLater := time.Now().Add( -20 * time.Hour )

	colScanner := &mockColumnScanner{
		ColumnNames: []string{
			"fruit_name",
			"name",
			"age",
			"height",
			"is_female",
			"when_created",
			"file_data",
			"greeting",
			"when_cleared",
			"when_reset",

			"my_bool",
			"my_bool_ptr",
			"my_bool_ptr_nil",

			"my_float32",
			"my_float32_ptr",
			"my_float32_ptr_nil",

			"my_float64",
			"my_float64_ptr",
			"my_float64_ptr_nil",

			"my_int",
			"my_int_ptr",
			"my_int_ptr_nil",

			"my_int8",
			"my_int8_ptr",
			"my_int8_ptr_nil",

			"my_int16",
			"my_int16_ptr",
			"my_int16_ptr_nil",

			"my_int32",
			"my_int32_ptr",
			"my_int32_ptr_nil",

			"my_int64",
			"my_int64_ptr",
			"my_int64_ptr_nil",

			"my_string",
			"my_string_ptr",
			"my_string_ptr_nil",

			"my_time",
			"my_time_ptr",
			"my_time_ptr_nil",

			"my_uint",
			"my_uint_ptr",
			"my_uint_ptr_nil",

			"my_uint8",
			"my_uint8_ptr",
			"my_uint8_ptr_nil",

			"my_uint16",
			"my_uint16_ptr",
			"my_uint16_ptr_nil",

			"my_uint32",
			"my_uint32_ptr",
			"my_uint32_ptr_nil",

			"my_uint64",
			"my_uint64_ptr",
			"my_uint64_ptr_nil",

			"my_string_from_byte_slice",
			"my_string_ptr_from_byte_slice",
			"my_string_ptr_from_byte_slice_nil",

			"my_math_big_rat_ptr",
			"my_math_big_float_ptr",
			"my_math_big_rat_ptr_from_byte_slice",
			"my_math_big_float_ptr_from_byte_slice",
			"my_other_math_big_float_ptr",
		},
		ColumnsErr:    nil,
		ScanValues:  []interface{}{
			"cherry",                           // fruit_name
			"Wanda Doe",                        // name
			int64(23),                          // age
			float64(5.2),                       // height
			true,                               // is_female
			now,                                // when_created
			[]byte("apple banana cherry"),      // file_data
			"Hello world!",                     // greeting
			nil,                                // when_cleared
			then,                               // when_reset

			true,                               // my_bool
			true,                               // my_bool_ptr
			nil,                                // my_bool_ptr_nil

			float32(-2.0),                      // my_float32
			float32(-3.0),                      // my_float32_ptr
			nil,                                // my_float32_ptr_nil

			float64(-4.0),                      // my_float64
			float64(-5.0),                      // my_float64_ptr
			nil,                                // my_float64_ptr_nil

			int(-1),                            // my_int
			int(-2),                            // my_int_ptr
			nil,                                // my_int_ptr_nil

			int8(-3),                           // my_int8
			int8(-4),                           // my_int8_ptr
			nil,                                // my_int8_ptr_nil

			int16(-5),                          // my_int16
			int16(-6),                          // my_int16_ptr
			nil,                                // my_int16_ptr_nil

			int32(-7),                          // my_int32
			int32(-8),                          // my_int32_ptr
			nil,                                // my_int32_ptr_nil

			int64(-9),                          // my_int64
			int64(-10),                         // my_int64_ptr
			nil,                                // my_int64_ptr_nil

			"apple",                            // my_string
			"banana",                           // my_string_ptr
			nil,                                // my_string_ptr_nil

			later,                              // my_time
			evenLater,                          // my_time_ptr
			nil,                                // my_time_ptr_nil

			uint(1),                            // my_uint
			uint(2),                            // my_uint_ptr
			nil,                                // my_uint_ptr_nil

			uint8(3),                           // my_uint8
			uint8(4),                           // my_uint8_ptr
			nil,                                // my_uint8_ptr_nil

			uint16(5),                          // my_uint16
			uint16(6),                          // my_uint16_ptr
			nil,                                // my_uint16_ptr_nil

			uint32(7),                          // my_uint32
			uint32(8),                          // my_uint32_ptr
			nil,                                // my_uint32_ptr_nil

			uint64(9),                          // my_uint64
			uint64(10),                         // my_uint64_ptr
			nil,                                // my_uint64_ptr_nil

			[]byte("ONE"),                      // my_string_from_byte_slice
			[]byte("two THREE"),                // my_string_ptr_from_byte_slice
			nil,                                // my_string_ptr_from_byte_slice_nil

			"3.14159265358979323846264338327950",         // "my_math_big_rat_ptr"
			"3.14159265358979323846264338327950",         // "my_math_big_float_ptr"
			[]byte("2.71828182845904523536028747135266"), // "my_math_big_rat_ptr_from_byte_slice"
			[]byte("2.71828182845904523536028747135266"), // "my_math_big_float_ptr_from_byte_slice"
			"0.1", // "my_other_math_big_float_ptr"
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
		MyBool           bool       `iter:"my_bool"`
		MyBoolPtr       *bool       `iter:"my_bool_ptr"`
		MyBoolPtrNil    *bool       `iter:"my_bool_ptr_nil"`

		MyFloat32        float32    `iter:"my_float32"`
		MyFloat32Ptr    *float32    `iter:"my_float32_ptr"`
		MyFloat32PtrNil *float32    `iter:"my_float32_ptr_nil"`

		MyFloat64        float64    `iter:"my_float64"`
		MyFloat64Ptr    *float64    `iter:"my_float64_ptr"`
		MyFloat64PtrNil *float64    `iter:"my_float64_ptr_nil"`

		MyInt            int        `iter:"my_int"`
		MyIntPtr        *int        `iter:"my_int_ptr"`
		MyIntPtrNil     *int        `iter:"my_int_ptr_nil"`

		MyInt8           int8       `iter:"my_int8"`
		MyInt8Ptr       *int8       `iter:"my_int8_ptr"`
		MyInt8PtrNil    *int8       `iter:"my_int8_ptr_nil"`

		MyInt16          int16      `iter:"my_int16"`
		MyInt16Ptr      *int16      `iter:"my_int16_ptr"`
		MyInt16PtrNil   *int16      `iter:"my_int16_ptr_nil"`

		MyInt32          int32      `iter:"my_int32"`
		MyInt32Ptr      *int32      `iter:"my_int32_ptr"`
		MyInt32PtrNil   *int32      `iter:"my_int32_ptr_nil"`

		MyInt64          int64      `iter:"my_int64"`
		MyInt64Ptr      *int64      `iter:"my_int64_ptr"`
		MyInt64PtrNil   *int64      `iter:"my_int64_ptr_nil"`

		MyString         string     `iter:"my_string"`
		MyStringPtr     *string     `iter:"my_string_ptr"`
		MyStringPtrNil  *string     `iter:"my_string_ptr_nil"`

		MyTime           time.Time  `iter:"my_time"`
		MyTimePtr       *time.Time  `iter:"my_time_ptr"`
		MyTimePtrNil    *time.Time  `iter:"my_time_ptr_nil"`

		MyUint            uint      `iter:"my_uint"`
		MyUintPtr        *uint      `iter:"my_uint_ptr"`
		MyUintPtrNil     *uint      `iter:"my_uint_ptr_nil"`

		MyUint8           uint8     `iter:"my_uint8"`
		MyUint8Ptr       *uint8     `iter:"my_uint8_ptr"`
		MyUint8PtrNil    *uint8     `iter:"my_uint8_ptr_nil"`

		MyUint16          uint16    `iter:"my_uint16"`
		MyUint16Ptr      *uint16    `iter:"my_uint16_ptr"`
		MyUint16PtrNil   *uint16    `iter:"my_uint16_ptr_nil"`

		MyUint32          uint32    `iter:"my_uint32"`
		MyUint32Ptr      *uint32    `iter:"my_uint32_ptr"`
		MyUint32PtrNil   *uint32    `iter:"my_uint32_ptr_nil"`

		MyUint64          uint64    `iter:"my_uint64"`
		MyUint64Ptr      *uint64    `iter:"my_uint64_ptr"`
		MyUint64PtrNil   *uint64    `iter:"my_uint64_ptr_nil"`

		MyStringFromByteSlice        string `iter:"my_string_from_byte_slice"`
		MyStringPtrFromByteSlice    *string `iter:"my_string_ptr_from_byte_slice"`
		MyStringPtrFromByteSliceNil *string `iter:"my_string_ptr_from_byte_slice_nil"`

		MyMathBigRatPtr   *big.Rat   `iter:"my_math_big_rat_ptr"`
		MyMathBigFloatPtr *big.Float `iter:"my_math_big_float_ptr"`
		MyMathBigRatPtrFromByteSlice   *big.Rat   `iter:"my_math_big_rat_ptr_from_byte_slice"`
		MyMathBigFloatPtrFromByteSlice *big.Float `iter:"my_math_big_float_ptr_from_byte_slice"`
		MyOtherMathBigFloatPtr *big.Float `iter:"my_other_math_big_float_ptr"`
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
	if expected, actual := now, myStruct.WhenCreated; !expected.Equal(actual) {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "apple banana cherry", string(myStruct.File); expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "Hello world!", myStruct.Greeting.String(); expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*time.Time)(nil), myStruct.WhenCleared; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := then, *myStruct.WhenReset; !expected.Equal(actual) {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := true, myStruct.MyBool; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := true, *myStruct.MyBoolPtr; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*bool)(nil), myStruct.MyBoolPtrNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := float32(-2.0), myStruct.MyFloat32; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := float32(-3.0), *myStruct.MyFloat32Ptr; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*float32)(nil), myStruct.MyFloat32PtrNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := float64(-4.0), myStruct.MyFloat64; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := float64(-5.0), *myStruct.MyFloat64Ptr; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*float64)(nil), myStruct.MyFloat64PtrNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := int(-1), myStruct.MyInt; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := int(-2), *myStruct.MyIntPtr; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*int)(nil), myStruct.MyIntPtrNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := int8(-3), myStruct.MyInt8; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := int8(-4), *myStruct.MyInt8Ptr; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*int8)(nil), myStruct.MyInt8PtrNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := int16(-5), myStruct.MyInt16; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := int16(-6), *myStruct.MyInt16Ptr; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*int16)(nil), myStruct.MyInt16PtrNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := int32(-7), myStruct.MyInt32; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := int32(-8), *myStruct.MyInt32Ptr; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*int32)(nil), myStruct.MyInt32PtrNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := int64(-9), myStruct.MyInt64; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := int64(-10), *myStruct.MyInt64Ptr; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*int64)(nil), myStruct.MyInt64PtrNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := "apple", myStruct.MyString; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "banana", *myStruct.MyStringPtr; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*string)(nil), myStruct.MyStringPtrNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := later, myStruct.MyTime; !expected.Equal(actual) {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := evenLater, *myStruct.MyTimePtr; !expected.Equal(actual) {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*time.Time)(nil), myStruct.MyTimePtrNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := uint(1), myStruct.MyUint; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := uint(2), *myStruct.MyUintPtr; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*uint)(nil), myStruct.MyUintPtrNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := uint8(3), myStruct.MyUint8; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := uint8(4), *myStruct.MyUint8Ptr; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*uint8)(nil), myStruct.MyUint8PtrNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := uint16(5), myStruct.MyUint16; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := uint16(6), *myStruct.MyUint16Ptr; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*uint16)(nil), myStruct.MyUint16PtrNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := uint32(7), myStruct.MyUint32; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := uint32(8), *myStruct.MyUint32Ptr; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*uint32)(nil), myStruct.MyUint32PtrNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := uint64(9), myStruct.MyUint64; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := uint64(10), *myStruct.MyUint64Ptr; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*uint64)(nil), myStruct.MyUint64PtrNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := "ONE", myStruct.MyStringFromByteSlice; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "two THREE", *myStruct.MyStringPtrFromByteSlice; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := (*string)(nil), myStruct.MyStringPtrFromByteSliceNil; expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}



	if expected, actual := "6283185307179586476925286766559/2000000000000000000000000000000", myStruct.MyMathBigRatPtr.String(); expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "3.141592654", myStruct.MyMathBigFloatPtr.Text('g', 10); expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "135914091422952261768014373567633/50000000000000000000000000000000", myStruct.MyMathBigRatPtrFromByteSlice.String(); expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "2.71828182845904523536028747135266", myStruct.MyMathBigFloatPtrFromByteSlice.Text('g', 33); expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "0.1", myStruct.MyOtherMathBigFloatPtr.Text('g', 10); expected != actual {
		t.Errorf("Expected (%T) ⟨%v⟩, but actually got (%T) ⟨%v⟩.", expected, expected, actual, actual)
		return
	}
}
