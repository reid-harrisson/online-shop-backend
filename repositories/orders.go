package repositories

import (
	"OnlineStoreBackend/models"

	"github.com/jinzhu/gorm"
)

type RepositoryOrder struct {
	DB *gorm.DB
}

func NewRepositoryOrder(db *gorm.DB) *RepositoryOrder {
	return &RepositoryOrder{DB: db}
}

func (repository *RepositoryOrder) ReadByStoreAndOrderID(modelOrder *models.StoreOrders, orderID uint64, storeID uint64) {
	repository.DB.Table("store_orders As so").
		Select(`
			soi.order_id,
			so.customer_id,
			soi.product_id,
			soi.price,
			soi.quantity,
			soi.sub_total_price,
			so.billing_address,
			so.shipping_address,
			soi.tax_rate,
			soi.tax_amount,
			soi.shipping_method,
			soi.shipping_price,
			soi.total_price,
			soi.status As product_status
		`).
		Joins(`Right Join store_order_items As soi On soi.order_id = so.id`).
		Where("soi.store_id = ? And soi.order_id = ?", storeID, orderID).
		Where("so.deleted_at Is Null And soi.deleted_at Is Null").
		Scan(modelOrder)
}

func (repository *RepositoryOrder) ReadByStoreID(modelOrders *[]models.StoreOrders, storeID uint64) {
	repository.DB.Table("store_orders As so").
		Select(`
			soi.order_id,
			so.customer_id,
			soi.product_id,
			soi.price,
			soi.quantity,
			soi.sub_total_price,
			so.billing_address_id,
			so.shipping_address_id,
			soi.tax_rate,
			soi.tax_amount,
			soi.shipping_method,
			soi.shipping_price,
			soi.total_price,
			soi.status As product_status
		`).
		Joins(`Right Join store_order_items As soi On soi.order_id = so.id`).
		Where("soi.store_id = ?", storeID).
		Where("so.deleted_at Is Null And soi.deleted_at Is Null").
		Scan(modelOrders)
}
func (repository *RepositoryOrder) ReadByCustomerID(modelOrders *[]models.CustomerOrders, customerID uint64) {
	repository.DB.Table("store_orders As so").
		Select(`
			so.id AS order_id,
			so.status AS order_status,
			Sum( soi.total_price ) As total_price,
			billing_address_id,
			shipping_address_id
		`).
		Joins(`Right Join store_order_items As soi On soi.order_id = so.id`).
		Where("so.customer_id = ?", customerID).
		Where("so.deleted_at Is Null And soi.deleted_at Is Null").
		Group("soi.order_id").
		Scan(modelOrders)
}

func (repository *RepositoryOrder) ReadByOrderID(modelOrder *models.CustomerOrdersWithAddress, orderID uint64) {
	modelOrder.Items = make([]models.CustomerOrderItems, 0)
	repository.DB.Table("store_orders As so").
		Select(`
			so.status As order_status,
			soi.store_id,
			soi.status As product_status,
			soi.variation_id,
			soi.price,
			soi.quantity,
			soi.sub_total_price,
			soi.tax_rate,
			soi.tax_amount,
			soi.shipping_method_id,
			soi.shipping_price,
			soi.total_price,
			so.billing_address_id,
			so.shipping_address_id
		`).
		Joins(`Right Join store_order_items As soi On soi.order_id = so.id`).
		Where("so.id = ?", orderID).
		Where("so.deleted_at Is Null And soi.deleted_at Is Null").
		Scan(&modelOrder.Items)
	if len(modelOrder.Items) > 0 {
		billingAdddressID := modelOrder.Items[0].BillingAddressID
		shippingAdddressID := modelOrder.Items[0].ShippingAddressID
		addrRepo := NewRepositoryCustomer(repository.DB)
		addrRepo.ReadAddressByID(&modelOrder.BillingAddress, billingAdddressID)
		addrRepo.ReadAddressByID(&modelOrder.ShippingAddress, shippingAdddressID)
	}
}
