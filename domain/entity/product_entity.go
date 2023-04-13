package entity

type Product struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `gorm:"not null"`
	Price float64 `gorm:"not null"`
}

func (Product) TableName() string {
	return "product"
}