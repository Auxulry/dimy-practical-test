package entity

type OrderItem struct {
	ID              uint `gorm:"primaryKey"`
	ProductID       uint
	PaymentMethodID uint
	OrderID         uint
	Total           float64
}