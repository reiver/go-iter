package iterfloat64


func (receiver *Slice) Scan(dest ...interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.kernel.KernelScan(nil, dest...)
}
