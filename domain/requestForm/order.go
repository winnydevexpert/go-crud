package domain

type CreateOrderRequest struct {
	ID          uint   `json:"id"`
	ProductName string `json:"productName"`
}
