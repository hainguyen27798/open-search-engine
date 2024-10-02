package response

type ServerCode struct {
	code int
}

func ReturnCode(code int) *ServerCode {
	return &ServerCode{code}
}

func (sc *ServerCode) InValid() bool {
	return sc.Code() >= ErrBadRequest
}

func (sc *ServerCode) Code() int {
	return sc.code
}

// 200xx: success
// 400xx: bad request err

const (
	CodeSuccess         = 20000 // Success
	UpdatedSuccess      = 20002 // Updated Success
	DeletedSuccess      = 20001 // Deleted Success
	CreatedSuccess      = 20101 // Created success
	ErrBadRequest       = 40000 // Bad Request
	ErrCodeParamInvalid = 40001 // Param invalid
	ErrCreateFailed     = 40003 // Create failed
	ErrNotFound         = 40400 // Resource not found
	ErrInternalError    = 50000 // Internal error
	ErrJWTInternalError = 50001 // JWT internal error
)

var CodeMsg = map[int]string{
	CodeSuccess:         "Success",
	UpdatedSuccess:      "Updated Success",
	DeletedSuccess:      "Deleted Success",
	CreatedSuccess:      "Created Success",
	ErrBadRequest:       "Bad Request",
	ErrCodeParamInvalid: "Param invalid",
	ErrCreateFailed:     "Create failed",
	ErrNotFound:         "Resource not found",
	ErrInternalError:    "Internal error",
	ErrJWTInternalError: "JWT internal error",
}
