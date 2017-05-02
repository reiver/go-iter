package itersqlrows

import (
	"github.com/reiver/go-cast"

	"database/sql"
	"fmt"
	"math/big"
	"reflect"
	"time"
)

const (
	targetNameKeyLong   = "iter.target.name"
	targetNameKeyNormal = "iter.name"
	targetNameKeyShort  = "iter"
)

func decode(colScanner columnScanner, v interface{}) error {

	if nil == colScanner {
		return errNilColumnScanner
	}
	if nil == v {
		return errNilDestination
	}


	data := map[string][]interface{}{}
	{
		columns, err := colScanner.Columns()
		if nil != err {
			return err
		}

		var a []interface{}
		for _, name := range columns {
			var x interface{} = new(target)

			if _, ok := data[name]; !ok {
				data[name] = []interface{}{}
			}
			data[name] = append(data[name], x)

			a = append(a, x)
		}

		if err := colScanner.Scan(a...); nil != err {
			return err
		}
	}


	var reflectedStructValue reflect.Value
	var reflectedStructType  reflect.Type
	{
		// This needs to get at the struct.
		reflectedValue := reflect.ValueOf(v)
		reflectedType  := reflect.TypeOf(v)
		if nil == reflectedType {
			return errNilReflectedType
		}
		for reflect.Ptr == reflectedValue.Kind() {
			reflectedValue = reflectedValue.Elem()
			reflectedType = reflectedType.Elem()
			if nil == reflectedType {
				return errNilReflectedType
			}
		}
		if reflect.Struct != reflectedType.Kind() {
			return fmt.Errorf("Unsupported Type: %T", v)
		}

		reflectedStructValue = reflectedValue
		reflectedStructType  = reflectedType
	}

	numFields := reflectedStructType.NumField()
	Loop: for fieldNumber:=0; fieldNumber<numFields; fieldNumber++ {

		reflectedFieldValue := reflectedStructValue.Field(fieldNumber)
		if !reflectedFieldValue.CanSet() {
			continue Loop
		}

		// Figure out the "name" the user wants to use.
		var name string
		{
			reflectedStructField := reflectedStructType.Field(fieldNumber)

			var ok bool

			name, ok = reflectedStructField.Tag.Lookup(targetNameKeyLong)
			if !ok {
				name, ok = reflectedStructField.Tag.Lookup(targetNameKeyNormal)
				if !ok {
					name, ok = reflectedStructField.Tag.Lookup(targetNameKeyShort)
					if !ok {
						name = reflectedStructField.Name
					}
				}
			}
		}

		// Get the value.
		var value interface{}
		{
			values, ok := data[name]
			if !ok {
				continue Loop
			}
			if 1 > len(values) {
				continue Loop
			}
			valueInterface := values[0]
			valueTarget, ok := valueInterface.(*target)
			if !ok {
				return fmt.Errorf("Unexpected Type: %T", valueTarget)
			}

			var err error
			value, err = valueTarget.Interface()
			if nil != err {
				return err
			}
		}


		{
			pointerToReflectedFieldValue := reflectedFieldValue.Addr()
			switch x := pointerToReflectedFieldValue.Interface().(type) {
			case sql.Scanner:
				if err := x.Scan(value); nil != err {
					return err
				}
				continue Loop
			}
		}


		var castedValue interface{}
		{
			switch reflectedFieldValue.Kind() {
			case                reflect.Bool:
				casted, err := cast.Bool(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Complex64:
				casted, err := cast.Complex64(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Complex128:
				casted, err := cast.Complex128(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Float32:
				casted, err := cast.Float32(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Float64:
				casted, err := cast.Float64(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Int:
				casted, err := cast.Int(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Int8:
				casted, err := cast.Int8(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Int16:
				casted, err := cast.Int16(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Int32:
				casted, err := cast.Int32(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Int64:
				casted, err := cast.Int64(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.String:
				casted, err := cast.String(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Uint:
				casted, err := cast.Uint(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Uint8:
				casted, err := cast.Uint8(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Uint16:
				casted, err := cast.Uint16(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Uint32:
				casted, err := cast.Uint32(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Uint64:
				casted, err := cast.Uint64(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted

			case reflect.Slice:
				if nil == value {
					castedValue = nil
				} else {
					switch casted := value.(type) {
					case []byte:
						castedValue = casted
					default:
						err := fmt.Errorf("Cannot cast into something of type %T, for struct field target named %q (%T).", value, name, reflectedFieldValue.Interface())
						return err
					}
				}
			case reflect.Struct:
				switch casted := value.(type) {
				case time.Time:
					castedValue = casted
				default:
					err := fmt.Errorf("Cannot cast into something of type %T, for struct field target named %q (%T).", value, name, reflectedFieldValue.Interface())
					return err
				}
			case reflect.Ptr:
				if nil == value {
					castedValue = nil
				} else {
					switch casted := value.(type) {
					case bool:
						castedValue = &casted
					case float32:
						castedValue = &casted
					case float64:
						castedValue = &casted
					case int:
						castedValue = &casted
					case int8:
						castedValue = &casted
					case int16:
						castedValue = &casted
					case int32:
						castedValue = &casted
					case int64:
						castedValue = &casted
					case string:
						switch reflectedFieldValue.Interface().(type) {
						case *big.Rat:
							r := new(big.Rat)
							_, err := fmt.Sscan(casted, r)
							if nil != err {
								return fmt.Errorf("Problem parsing string into *math/big.Rat, because: (%T) %v", err, err)
							}

							castedValue = r
						case *big.Float:
							const prec = 400

							f, _, err := big.ParseFloat(casted, 10, prec, big.ToNearestEven)
							if nil != err {
								return err
							}

							castedValue = f
						default:
							castedValue = &casted
						}
					case time.Time:
						castedValue = &casted
					case uint:
						castedValue = &casted
					case uint8:
						castedValue = &casted
					case uint16:
						castedValue = &casted
					case uint32:
						castedValue = &casted
					case uint64:
						castedValue = &casted

					case []byte:
						switch reflectedFieldValue.Interface().(type) {
						case *big.Rat:
							r := new(big.Rat)
							_, err := fmt.Sscan(string(casted), r)
							if nil != err {
								return fmt.Errorf("Problem parsing string into *math/big.Rat, because: (%T) %v", err, err)
							}

							castedValue = r
						case *big.Float:
							const prec = 400

							f, _, err := big.ParseFloat(string(casted), 10, prec, big.ToNearestEven)
							if nil != err {
								return err
							}

							castedValue = f
						case []byte:
							castedValue = casted
						case string:
							castedValue = string(casted)
						case *string:
							x := string(casted)
							castedValue = &x
						default:
							err := fmt.Errorf("Cannot cast into something of type %T, for struct field target named %q (%T).", value, name, reflectedFieldValue.Interface())
							return err
						}

					default:
						err := fmt.Errorf("Cannot cast into something of type %T, for struct field target named %q (%T).", value, name, reflectedFieldValue.Interface())
						return err
					}
				}
			default:
				err := fmt.Errorf("Cannot cast into something of type %T, for struct field target named %q (%T).", value, name, reflectedFieldValue.Interface())
				return err
			}
		}

		if err := func() (err error) {
			defer func() {
				if r := recover(); nil != r {
					err = fmt.Errorf("Could not set value ([%T] %v) for struct field named %q (%T) because: (%T) %v", castedValue, castedValue, name, reflectedFieldValue.Interface(), r, r)
				}
			}()
			if nil == castedValue {
				reflectedFieldValue.Set( reflect.Zero(reflectedFieldValue.Type()) )
			} else {
				reflectedFieldValue.Set( reflect.ValueOf(castedValue) )
			}
			return nil
		}(); nil != err {
			return err
		}
	}


	return nil
}
