package entity

type PaymentMethod struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	IsActive bool
}

func (PaymentMethod) TableName() string {
	return "payment_method"
}