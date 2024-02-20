package models

type Users struct {
	ContactEmail string `gorm:"column:contact_email"`
	ContactPhone string `gorm:"column:contact_phone"`
}
