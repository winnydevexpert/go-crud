package domain

type CreateCustomerRequest struct {
	FirstName string `json:"firstName" validate:"required" binding:"required"`
	Body      string `json:"body"`
}
