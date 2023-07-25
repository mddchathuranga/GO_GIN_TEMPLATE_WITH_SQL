package exutilities

// ErrorResponse is used to document error responses. not working for swagger annotations getting error
type ErrorResponse struct {
	Message string `json:"message"`
}
type ValidationErrorResponse struct {
	Message string `json:"message"`
}
