package repositories

import (
	"OnlineStoreBackend/models"
	"strconv"

	"gorm.io/gorm"
)

type RepositoryTax struct {
	DB *gorm.DB
}

func NewRepositoryTax(db *gorm.DB) *RepositoryTax {
	return &RepositoryTax{DB: db}
}

func (repository *RepositoryTax) ReadByCustomerID(modelTax *models.Taxes, customerID uint64) error {
	return repository.DB.Table("users").
		Select(`users.country_id As country_id,
			countries.tax_rate As tax_rate,
			? As customer_id`, customerID).
		Joins("Join countries On countries.id = users.country_id").
		Where("users.id = ?", customerID).
		Where("countries.deleted_at Is Null And users.deleted_at Is Null").
		Limit(1).
		Scan(modelTax).
		Error
}

func (repository *RepositoryTax) ReadByCountryID(modelTax *models.Taxes, countryID uint64) error {
	return repository.DB.Table("countries").
		Select("tax_rate, id As country_id").
		Where("id = ? And deleted_at Is Null", countryID).
		Limit(1).
		Scan(modelTax).
		Error
}

func (repository *RepositoryTax) ReadCurrency(currencySymbol *string, exchangeRate *float64, customerID uint64) error {
	currencyCode, temp := "", map[string]interface{}{}
	err := repository.DB.Table("users").
		Select("curs.code As code, curs.symbol As symbol").
		Joins("Join countries As couns On couns.id = users.country_id").
		Joins("Join currencies As curs On curs.code = couns.currency_code").
		Where("users.id = ?", customerID).
		Scan(&temp).
		Error
	if err != nil {
		return err
	}

	if temp["code"] != nil {
		currencyCode = temp["code"].(string)
	}
	if temp["symbol"] != nil {
		*currencySymbol = temp["symbol"].(string)
	}

	err = repository.DB.Table("exchange_rates").
		Order("id Desc").
		Limit(1).
		Scan(&temp).
		Error
	if err != nil {
		return err
	}

	if temp[currencyCode] != nil {
		*exchangeRate, _ = strconv.ParseFloat(temp[currencyCode].(string), 64)
	}
	if *exchangeRate == 0 {
		*exchangeRate = 1
	}
	if *currencySymbol == "" {
		*currencySymbol = "$"
	}

	return nil
}
