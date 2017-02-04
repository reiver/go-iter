package iter

type Iterator interface {
	Close() error
	Decode(interface{}) error
	Err() error
	Next() bool
}
