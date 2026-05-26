package core

import "errors"

var (
	ErrInvalidGuaIndex = errors.New("invalid gua index")
	ErrInvalidYaoIndex = errors.New("invalid yao index")
	ErrInvalidDayanIdx = errors.New("invalid dayan index, must be positive")
	ErrGuaNotFound     = errors.New("gua not found")
)

var (
	ErrInvalidDayanNumber = errors.New("dayan number must be between 1 and 81")
)

var (
	ErrInvalidCoinValues  = newError("coin values must be 6 integers each in range 6-9")
	ErrInvalidDayanValues = newError("dayan values must be 6 integers each in range 6-9")
)

var (
	ErrInvalidJiaZiIndex = errors.New("jiazi index must be between 1 and 60")
	ErrInvalidTianGan    = errors.New("invalid tiangan index")
	ErrInvalidDiZhi      = errors.New("invalid dizhi index")
	ErrInvalidJiaZiCombo = errors.New("invalid jiazi combination: gan and zhi must have same parity")
)

func newError(msg string) error {
	return &divinationError{msg: msg}
}

type divinationError struct{ msg string }

func (e *divinationError) Error() string { return e.msg }
