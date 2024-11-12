package ecode

const (
	ParamError           = 400
	WalletRecordNotFound = 10010
)

var Errors = map[int]string{
	ParamError:           "param error",
	WalletRecordNotFound: "wallet record not found",
}
