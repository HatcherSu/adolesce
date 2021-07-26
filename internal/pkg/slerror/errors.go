package slerror

type errCode struct {
	err   error
	code  int
	cause error
}
