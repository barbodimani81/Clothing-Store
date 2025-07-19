package models

type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `json:"user_id"`
	Items     []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	Total     float64     `json:"total"`
	CreatedAt int64       `json:"created_at"`
}
