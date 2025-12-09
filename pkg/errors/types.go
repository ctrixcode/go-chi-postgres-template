package errors

type ErrorType struct {
	Code    string
	Message string
}

var (
	ErrBadRequest          = ErrorType{Code: "BAD_REQUEST", Message: "Bad request"}
	ErrValidationFailed    = ErrorType{Code: "VALIDATION_FAILED", Message: "Validation failed"}
	ErrUnauthorized        = ErrorType{Code: "UNAUTHORIZED", Message: "Unauthorized: User not authenticated."}
	ErrNotFound            = ErrorType{Code: "NOT_FOUND", Message: "Resource not found."}
	ErrInternalServerError = ErrorType{Code: "INTERNAL_SERVER_ERROR", Message: "Internal server error"}
	ErrSomethingWentWrong  = ErrorType{Code: "SOMETHING_WENT_WRONG", Message: "Something went wrong"}
)
