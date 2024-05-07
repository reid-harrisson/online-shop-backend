package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model

	// credentials
	Email       string `gorm:"column:email; type:varchar(100); default:null"`
	MobileNo    string `gorm:"column:mobile_no; type:varchar(45); default:null"`
	WebauthnID  string `gorm:"column:webauthn_id; type:varchar(100); default:null"`
	WebauthnKey string `gorm:"column:webauthn_key; type:text; default:null"`
	Password    string `gorm:"column:password; type:varchar(100); default:null"`
	AuthID      string `gorm:"column:auth_id; type:varchar(100); default:null"`

	// activity
	IPCreation     string    `gorm:"column:ip_creation; type:varchar(20); default:null"`
	IPLastModified string    `gorm:"column:ip_last_modified; type:varchar(20); default:null"`
	IPLastLogin    string    `gorm:"column:ip_last_login; type:varchar(20); default:null"`
	LastLoginDate  time.Time `gorm:"column:last_login_date; type:datetime; default:null"`

	// kyc info
	FirstName   string    `gorm:"column:first_name; type:varchar(45); default:null"`
	LastName    string    `gorm:"column:last_name; type:varchar(45); default:null"`
	Gender      string    `gorm:"column:gender; type:varchar(10); default:Female"`
	Age         int8      `gorm:"column:age; type:smallint; default:null"`
	DateOfBirth time.Time `gorm:"column:dob; type:date; default:null"`
	NationalID  string    `gorm:"column:national_id; type:varchar(25); default:null"`

	// location
	AddressLine1 string `gorm:"column:address_line1; type:varchar(100); default:null"`
	AddressLine2 string `gorm:"column:address_line2; type:varchar(100); default:null"`
	CityID       uint64 `gorm:"column:city_id; type:bigint(20); default:null"`
	RegionID     uint64 `gorm:"column:region_id; type:bigint(20); default:null"`
	CountryID    uint64 `gorm:"column:country_id; type:bigint(20); default:null"`
	CompanyID    uint64 `gorm:"column:company_id; type:bigint(20); default:null"`
	PostalCode   string `gorm:"column:postal_code; type:varchar(15); default:null"`
	// status
	Active int8 `gorm:"column:active; type:tinyint; default:1"`
}

func (Users) TableName() string {
	return "users"
}
