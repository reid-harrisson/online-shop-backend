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

func (repository *RepositoryCombo) ReadByID(modelCombo *models.Combos, comboID uint64) error {
	return repository.DB.First(modelCombo, comboID).Error
}

func (repository *RepositoryCombo) ReadDetail(modelItems *[]models.CartItemsWithDetail, comboID uint64) error {
	return repository.DB.Table("store_combo_items As items").
		Select(`
			items.id,
			items.variation_id,
			items.quantity,
			prods.store_id,
			vars.price,
			vars.discount_amount,
			vars.discount_type,
			vars.image_urls,
			vars.stock_level,
			vars.title As variation_name,
			Group_Concat(Concat('"', cates.name,'"') Separator ', ') As categories
		`).
		Joins("Left Join store_product_variations As vars On vars.id = items.variation_id").
		Joins("Left Join store_products As prods On prods.id = vars.product_id").
		Joins("Left Join store_product_categories As prodcates On prodcates.product_id = prods.id").
		Joins("Left Join store_categories As cates On cates.id = prodcates.category_id").
		Group("items.variation_id").
		Where("items.combo_id = ?", comboID).
		Scan(modelItems).
		Error
}
