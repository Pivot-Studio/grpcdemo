package errwraper

type ErrCode int

type ApiError struct {
	Msg  string
	Code ErrCode
}

func (ae ApiError) Error() string {
	return ae.Msg
}

func Simple(msg string, code ErrCode) error {
	return ApiError{Msg: msg, Code: code}
}
