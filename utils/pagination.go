package utils

import "math"

type PaginationObject struct {
	Total        int64  `json:"total"`        // จำนวนรายการทั้งหมด
	PreviousPage *int64 `json:"previousPage"` // หน้าก่อนหน้า
	CurrentPage  int64  `json:"currentPage"`  // หน้าปัจจุบัน
	NextPage     *int64 `json:"nextPage"`     // หน้าถัดไป
	LastPage     *int64 `json:"lastPage"`     // หน้าสุดท้าย
	PerPage      int64  `json:"perPage"`      // จำนวนต่อหน้า
}

func CalculatorPreviousPage(page int64, total int64) *int64 {
	if page == 1 || total == 0 {
		return nil
	} else {
		previousPage := page - 1
		return &previousPage
	}
}

func CalculatorNextPage(page int64, totalPage int64) *int64 {
	if page == totalPage {
		return nil
	} else {
		nextPage := page + 1
		return &nextPage
	}
}

func CalculatorOffset(page int64, perPage int64) int64 {
	return (page - 1) * perPage
}

func Pagination(page int64, perPage int64, total int64) PaginationObject {
	if total == 0 {
		return PaginationObject{
			Total:        total,
			PreviousPage: nil,
			CurrentPage:  page,
			NextPage:     nil,
			LastPage:     nil,
			PerPage:      perPage,
		}
	} else {
		lastPage := int64(math.Ceil(float64(total) / float64(perPage)))
		return PaginationObject{
			Total:        total,
			PreviousPage: CalculatorPreviousPage(page, total),
			CurrentPage:  page,
			NextPage:     CalculatorNextPage(page, lastPage),
			LastPage:     &lastPage,
			PerPage:      perPage,
		}
	}
}
