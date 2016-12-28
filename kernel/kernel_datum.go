package iterkernel

func (receiver *Kernel) KernelDatum() (interface{}, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	if err := receiver.err; nil != err {
		return nil, err
	}

	return receiver.datum, nil
}
