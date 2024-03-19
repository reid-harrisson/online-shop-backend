package repositories

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"

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

func (repository *RepositoryCombo) ReadApproved(modelCombos *[]models.Combos, modelItems *[]models.ComboItems, storeID uint64) {
	repository.DB.Where("store_id = ? And status = ?", storeID, utils.Approved).Find(modelCombos)
	comboIDs := []uint64{}
	for _, modelCombo := range *modelCombos {
		comboIDs = append(comboIDs, uint64(modelCombo.ID))
	}
	repository.DB.Where("combo_id In (?)", comboIDs).Find(modelItems)
}

func (repository *RepositoryCombo) ReadPublished(modelCombos *[]models.Combos, modelItems *[]models.ComboItems, storeID uint64) {
	repository.DB.Where("store_id = ? And status = ?", storeID, utils.Pending).Find(modelCombos)
	comboIDs := []uint64{}
	for _, modelCombo := range *modelCombos {
		comboIDs = append(comboIDs, uint64(modelCombo.ID))
	}
	repository.DB.Where("combo_id In (?)", comboIDs).Find(modelItems)
}

func (repository *RepositoryCombo) ReadByID(modelCombo *models.Combos, comboID uint64) error {
	return repository.DB.First(modelCombo, comboID).Error
}

func (repository *RepositoryCombo) ReadDetail(modelItems *[]models.CartItemsWithDetail, comboID uint64) error {
	return repository.DB.Table("store_combo_items As items").
		Select(`items.id,
			items.variation_id,
			items.quantity,
			prods.store_id,
			vars.price,
			vars.discount_amount,
			vars.discount_type,
			vars.image_urls,
			vars.stock_level,
			vars.title As variation_name,
			ships.weight,
			ships.width,
			ships.length,
			ships.height,
			Group_Concat(Concat('"', cates.name,'"') Separator ', ') As categories`).
		Joins("Left Join store_product_variations As vars On vars.id = items.variation_id").
		Joins("Left Join store_products As prods On prods.id = vars.product_id").
		Joins("Left Join store_product_categories As prodcates On prodcates.product_id = prods.id").
		Joins("Left Join store_categories As cates On cates.id = prodcates.category_id").
		Joins("Left Join store_shipping_data As ships On ships.variation_id = vars.id").
		Group("items.id").
		Where(`items.combo_id = ?
			And items.deleted_at Is Null
			And vars.deleted_at Is Null
			And prods.deleted_at Is Null
			And prodcates.deleted_at Is Null
			And cates.deleted_at Is Null`, comboID).
		Scan(modelItems).
		Error
}

func (repository *RepositoryCombo) ReadStatus(status *utils.ProductStatus, comboID uint64) error {
	return repository.DB.Model(&models.Combos{}).Select("status").Where("id = ?", comboID).Scan(status).Error
}
