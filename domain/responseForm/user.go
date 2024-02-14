package domain

type UserRegisterResponse struct {
	ApiVersion string `json:"apiVersion"`
	Data       struct {
		Title       string       `json:"title"`
		Description string       `json:"description"`
		Item        UserRegister `json:"item"`
	} `json:"data"`
}

type UserRegister struct {
	FirstName string  `json:"firstName"` // ชื่อ
	MidName   *string `json:"midName"`   // ชื่อกลาง
	LastName  string  `json:"lastName"`  // นามสกุล
}
