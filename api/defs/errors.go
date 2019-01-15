package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{400, Err{"Request body is not correct", "001"}}
	ErrorNotAuthUser = ErrorResponse{401, Err{"User authentication failed", "002"}}
)