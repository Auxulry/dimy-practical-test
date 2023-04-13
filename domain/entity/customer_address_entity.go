package entity

type CustomerAddress struct {
	ID         uint `gorm:"primaryKey"`
	CustomerID uint
	Address    string
}

func (CustomerAddress) TableName() string {
	return "customer_address"
}