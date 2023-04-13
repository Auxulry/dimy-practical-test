package entity

type Customer struct {
	ID                uint `gorm:"primaryKey"`
	CustomerName      string
	CustomerAddresses []CustomerAddress
}

func (Customer) TableName() string {
	return "customer"
}