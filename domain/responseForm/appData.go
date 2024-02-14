package domain

import "go-crud-learn/utils"

type AppDataListResponse struct {
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Items       interface{}            `json:"items"`
	Pagination  utils.PaginationObject `json:"pagination"`
}
