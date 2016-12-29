package iterfloat64

// Decode stores the next datum in the data represented by the empty interface value `x`.
// If `x` is nil, the value will be discarded.
// Otherwise, the value underlying `x` must be a pointer to the correct type for the next datum.
func (receiver *Slice) Decode(x interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.kernel.KernelDecode(nil, x)
}
