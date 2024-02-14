package utils

const (
	// AuthErrorMessage ...
	AuthErrorMessage = ""
	// AuthRequiredError is a message sending when there is no Authorization field in the header
	AuthRequiredError = "authorization required"
	// TokenError ...
	TokenError = "invalid token"
	// PermissionError ...
	PermissionError = "permission denied"
)

// ErrorMessagePrototype - a prototype for error message
type ErrorMessagePrototype struct {
	APIVersion string        `json:"apiVersion"`
	Error      errorResponse `json:"error"`
}

// SuccessMessagePrototype -- a prototype for success message
type SuccessMessagePrototype struct {
	APIVersion string     `json:"apiVersion"`
	Data       DataObject `json:"data"`
}

type errorResponse struct {
	Code    string  `json:"code"`
	Message *string `json:"message,omitempty"`
}

// DataObject - response any API
type DataObject struct {
	Kind        *string           `json:"kind,omitempty"`
	Title       *string           `json:"title,omitempty"`
	Description *string           `json:"description,omitempty"`
	Item        interface{}       `json:"item,omitempty"`
	Items       interface{}       `json:"items,omitempty"`
	Pagination  *PaginationObject `json:"pagination,omitempty"`
}

// ErrorMessage - return an error message
func ErrorMessage(code string, message *string) ErrorMessagePrototype {
	err := errorResponse{
		Code:    code,
		Message: message,
	}
	return ErrorMessagePrototype{APIVersion: "1.0.0", Error: err}
}

// SuccessMessage - return an success message
func SuccessMessage(data DataObject) SuccessMessagePrototype {
	return SuccessMessagePrototype{APIVersion: "1.0.0", Data: data}
}
