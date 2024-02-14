package domain

type CreateOrderResponse struct {
	ID          uint   `json:"id"`
	ProductName string `json:"productName"`
}
