package repositories

import (
	"OnlineStoreBackend/models"
	"sort"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type RepositoryVariation struct {
	DB *gorm.DB
}

func NewRepositoryVariation(db *gorm.DB) *RepositoryVariation {
	return &RepositoryVariation{DB: db}
}

func (repository *RepositoryVariation) ReadByID(modelVar *models.Variations, variationID uint64) error {
	return repository.DB.First(modelVar, variationID).Error
}

func (repository *RepositoryVariation) ReadBySku(modelVar *models.Variations, sku string) error {
	return repository.DB.Where("sku = ?", sku).First(modelVar).Error
}

func (repository *RepositoryVariation) ReadByAttributeValueIDs(modelVar *models.Variations, valueIDs []uint64, productID uint64) error {
	ids := make([]string, 0)
	sort.Slice(valueIDs, func(i, j int) bool { return valueIDs[i] > valueIDs[j] })
	for _, valueID := range valueIDs {
		ids = append(ids, strconv.FormatUint(uint64(valueID), 10))
	}
	temp := strings.Join(ids, ",")

	return repository.DB.
		Table("store_product_variations As vars").
		Select("vars.*, Group_Concat(dets.attribute_value_id) As ids").
		Joins("Left Join store_product_variation_details As dets ON dets.variation_id = vars.id").
		Group("vars.id").
		Where("vars.product_id = ? And vars.deleted_at Is Null And dets.deleted_at Is Null", productID).
		Having("(ids Is Null And ? = '') Or ids = ?", temp, temp).
		Limit(1).
		Scan(modelVar).
		Error
}

func (repository *RepositoryVariation) ReadByProduct(modelVars *[]models.VariationsWithAttributeValue, productID uint64) error {
	return repository.DB.
		Table("store_product_variations As vars").
		Select(`
			vars.*,
			vals.id As attribute_value_id,
			vals.attribute_value,
			attrs.attribute_name,
		`).
		Joins("Left Join store_product_variation_details As dets On dets.variation_id = vars.id").
		Where("vars.product_id = ?", productID).
		Where("vars.deleted_at Is Null And dets.deleted_at Is Null And vals.deleted_at Is Null And attrs.deleted_at Is Null").
		Scan(&modelVars).
		Error
}
