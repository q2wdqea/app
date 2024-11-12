package ecode

const (
	ParamError     = 400
	RecordNotFound = 10010
)

var Errors = map[int]string{
	ParamError:     "param error",
	RecordNotFound: "record not found",
}
