package iter

type Iterator interface {
	Decode(interface{}) error
	Err() error
	Next() bool
}
