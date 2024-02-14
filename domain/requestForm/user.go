package domain

type UserRegisterRequest struct {
	FirstName string  `json:"firstName" validate:"required" binding:"required"` // ชื่อ
	MidName   *string `json:"midName"`                                          // ชื่อกลาง
	LastName  string  `json:"lastName" validate:"required" binding:"required"`  // นามสกุล
}
