package itersqlrows

import (
	"database/sql"
	"sync"
)

type Rows struct {
	Rows *sql.Rows
	mutex sync.RWMutex
}

func (receiver *Rows) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	rows := receiver.Rows
	if nil == rows {
		return errNilSqlRows
	}

	return rows.Close()
}

func (receiver *Rows) Decode(v interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	rows := receiver.Rows
	if nil == rows {
		return errNilSqlRows
	}

	if err := decode(rows, v); nil != err {
		return err
	}

	return nil
}

func (receiver *Rows) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	rows := receiver.Rows
	if nil == rows {
		return errNilSqlRows
	}

	return rows.Err()
}

func (receiver *Rows) Next() bool {
	if nil == receiver {
		return false
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	rows := receiver.Rows
	if nil == rows {
		return false
	}

	return rows.Next()
}
