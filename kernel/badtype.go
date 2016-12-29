package iterkernel

import (
	"fmt"
)

type internalBadTypeComplainer struct{
	sourceType string
	actualType string
}

func (receiver internalBadTypeComplainer) Error() string {
	return fmt.Sprintf("Bad Type: %q; for source type %q", receiver.actualType, receiver.sourceType)
}

func (internalBadTypeComplainer) BadTypeComplainer() {
	// Nothing here.
}
