package ricknmorty

import "errors"

var (
	ErrUnknown       = errors.New("unknown error on ricknmorty client")
	ErrNotFound      = errors.New("not found result to this query")
	ErrResponse      = errors.New("api returned error")
	ErrInvalidParams = errors.New("invalid params in this query")
)
