package itersqlrows

import (
	"errors"
)

var (
	errNilColumnScanner = errors.New("Nil Column Scanner")
	errNilDestination   = errors.New("Nil Destination")
	errNilReceiver      = errors.New("Nil Receiver")
	errNilReflectedType = errors.New("Nil Reflected Type")
	errNilSqlRows       = errors.New("Nil SQL Rows")
)
