package model

type Order struct {
	ID          uint   `json:"id"`
	ProductName string `json:"productName"`
}

// Id			uint		`gorm:"primary_key; AUTO_INCREMENT" json:"id"`
// UserId		int 		`gorm:"NOT NULL; index" json:"user_id"`
// FinishedAt 	time.Time 	`gorm:"DEFAULT:now()" json:"finished_at"`
