package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryCombo struct {
	DB *gorm.DB
}

func NewRepositoryCombo(db *gorm.DB) *RepositoryCombo {
	return &RepositoryCombo{DB: db}
}

func (repository *RepositoryCombo) ReadByStoreID(modelCombos *[]models.Combos, modelItems *[]models.ComboItems, storeID uint64) {
	repository.DB.Where("store_id = ?", storeID).Find(modelCombos)
	comboIDs := []uint64{}
	for _, modelCombo := range *modelCombos {
		comboIDs = append(comboIDs, uint64(modelCombo.ID))
	}
	repository.DB.Where("combo_id In (?)", comboIDs).Find(modelItems)
}

func (repository *RepositoryCombo) ReadByID(modelCoupon *models.Coupons, couponID uint64) error {
	return repository.DB.First(modelCoupon, couponID).Error
}

func (repository *RepositoryCombo) ReadByCode(modelCoupon *models.Coupons, code string) error {
	return repository.DB.Where("coupon_code = ?", code).First(modelCoupon).Error
}

func (repository *RepositoryCombo) ReadByIDs(modelCoupon *[]models.Coupons, ids []uint64) error {
	return repository.DB.Where("id In (?)", ids).Find(modelCoupon).Error
}
