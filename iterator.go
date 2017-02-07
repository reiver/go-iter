package iter

// Iterator represents a "normal" iterator.
//
// Many of the sub-packages that are part of this export types that fit
// this interface.
//
// Some of those exported type have more methods than just these.
//
// But these can be relied on to exist.
//
// Example Usage
//
//	iterator := //...
//	
//	defer iterator.Close()
//	
//	for iterator.Next() {
//	
//		datum := struct{
//			GivenName  string `iter:"given_name"`
//			FamilyName string `iter:"family_name"`
//			Age        int64  `iter:"age"`
//		}{}
//	
//		if err := iterator.Decode(&datum); nil != err {
//			return err
//		}
//
//		//@TODO: do something with `datum`.
//	}
//	if err := iterator.Err(); nil != err {
//		return err
//	}
type Iterator interface {
	Close() error
	Decode(interface{}) error
	Err() error
	Next() bool
}
