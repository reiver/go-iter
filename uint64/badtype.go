package iteruint64

import (
	"fmt"
)

type internalBadTypeComplainer struct{
	actualType string
}

func (receiver internalBadTypeComplainer) Error() string {
	return fmt.Sprintf("Bad Type: %q", receiver.actualType)
}

func (internalBadTypeComplainer) BadTypeComplainer() {
	// Nothing here.
}
