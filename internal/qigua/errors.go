package qigua

import "errors"

var (
	ErrInvalidCoinValues  = errors.New("coin values must be 6 integers each in range 6-9")
	ErrInvalidDayanValues = errors.New("dayan values must be 6 integers each in range 6-9")
)
