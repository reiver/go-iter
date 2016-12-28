# go-iter

Package **iter** provides tools for creating iterators, for the Go programming language.

These iterators are intentionally made to resemble `*sql.Rows` from the `"database/sql"` package.
Including having the same `Close`, `Err`, `Next`, and `Scan` methods.

(Note that to turn something into an actual `*sql.Rows` from the `"database/sql"` package,
instead of _just_ resembling it, use https://github.com/reiver/go-shunt instead.)


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-iter

[![GoDoc](https://godoc.org/github.com/reiver/go-iter?status.svg)](https://godoc.org/github.com/reiver/go-iter)


## Example

```go
var slice []string = []string {
	"apple",
	"banana",
	"cherry",
}

iterator := iterstring.Slice{
	Slice: slice,
}

defer iterator.Close()

for iterator.Next() {

	var datum string // Could have also used: var datum interface{}

	if err := iterator.Decode(&datum); nil != err {
		return err
	}

	fmt.Printf("Next datum: %v \n", datum)
}
if err := iterator.Err(); nil != err {
	return err
}
```
