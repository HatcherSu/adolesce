package slerror

// return successful code
const (
	SuccessCode int = 0
)

// return the failure code
const (
	InvalidParamErrCode int = iota + 1000
	InnerServerErrCode
)
