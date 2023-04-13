package entity

type Order struct {
	ID                uint `gorm:"primaryKey"`
	CustomerID        uint
	CustomerAddressID uint
	OrderItems        []OrderItem
	Total             float64
}

type Items struct {
	ProductID       uint
	PaymentMethodID uint
}

type OrderRequest struct {
	CustomerID        uint
	CustomerAddressID uint
	OrderItems        []*Items
}