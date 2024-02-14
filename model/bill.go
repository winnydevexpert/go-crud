package model

type Bill struct {
	ID         uint     `gorm:"primaryKey; NOT NULL;" json:"id"`
	CustomerId uint     `json:"CustomerId"`
	Customer   Customer `gorm:"foreignKey:CustomerId; NOT NULL;" json:"customer"`
	OrderId    uint     `json:"CrderId"`
	Order      Order    `gorm:"foreignKey:OrderId; NOT NULL;" json:"order"`
	Amount     int      `json:"amount"`
}
