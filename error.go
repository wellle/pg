package pg

import (
	"errors"
	"fmt"
)

var (
	ErrSSLNotSupported = errors.New("pg: SSL is not enabled on the server")

	ErrNoRows    = errors.New("pg: no rows in result set")
	ErrMultiRows = errors.New("pg: multiple rows in result set")

	errClosed         = errors.New("pg: database is closed")
	errTxDone         = errors.New("pg: transaction has already been committed or rolled back")
	errStmtClosed     = errors.New("pg: statement is closed")
	errListenerClosed = errors.New("pg: listener is closed")
)

var (
	_ Error = &pgError{}
	_ Error = &pgError{}
)

type Error interface {
	Field(byte) string
}

type pgError struct {
	c map[byte]string
}

func (err *pgError) Field(k byte) string {
	return err.c[k]
}

func (err *pgError) Error() string {
	return fmt.Sprintf(
		"%s #%s %s: %s",
		err.Field('S'), err.Field('C'), err.Field('M'), err.Field('D'),
	)
}

type IntegrityError struct {
	*pgError
}
