package iter

import (
	"fmt"
)

type internalBadTypeComplainer struct{
	actualType string
}

func (receiver internalBadTypeComplainer) Error() string {
	return fmt.Sprintf("iter: bad type: %q", receiver.actualType)
}

func (internalBadTypeComplainer) BadTypeComplainer() {
	// Nothing here.
}
