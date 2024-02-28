package models

type TaxSettings struct {
	TaxRate    float64 `gorm:"column:tax_rate; type:decimal(20,6)"`
	CountryID  uint64  `gorm:"column:country_id; type:bigint(20) unsigned"`
	CustomerID uint64  `gorm:"column:customer_id; type:bigint(20) unsigned"`
}
