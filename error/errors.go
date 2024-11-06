package apierror

type ApiError int

// 0 --> 999 | GENERAL ERRORS
const (
	UnexpectedError ApiError = iota
	AccessDenied
	NotImplemented
	InvalidRequest
	BadRequest
)
