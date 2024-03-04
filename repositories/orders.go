package repositories

import (
	"OnlineStoreBackend/models"

	"gorm.io/gorm"
)

type RepositoryOrder struct {
	DB *gorm.DB
}

func NewRepositoryOrder(db *gorm.DB) *RepositoryOrder {
	return &RepositoryOrder{DB: db}
}

func (repository *RepositoryOrder) ReadByStoreAndOrderID(modelOrder *models.StoreOrders, orderID uint64, storeID uint64) {
	repository.DB.Table("store_orders As ords").
		Select(`
			oitms.order_id,
			ords.customer_id,
			oitms.variation_id,
			oitms.price,
			oitms.quantity,
			oitms.sub_total_price,
			ords.billing_address,
			ords.shipping_address,
			oitms.tax_rate,
			oitms.tax_amount,
			oitms.shipping_method_id,
			oitms.shipping_price,
			oitms.total_price,
			oitms.status As product_status
		`).
		Joins(`Right Join store_order_items As oitms On oitms.order_id = ords.id`).
		Where("oitms.store_id = ? And oitms.order_id = ?", storeID, orderID).
		Where("ords.deleted_at Is Null And oitms.deleted_at Is Null").
		Scan(modelOrder)
}

func (repository *RepositoryOrder) ReadByStoreID(modelOrders *[]models.StoreOrders, storeID uint64) {
	repository.DB.Table("store_orders As ords").
		Select(`
			oitms.order_id,
			ords.customer_id,
			oitms.variation_id,
			oitms.price,
			oitms.quantity,
			oitms.sub_total_price,
			ords.billing_address_id,
			ords.shipping_address_id,
			oitms.tax_rate,
			oitms.tax_amount,
			oitms.shipping_method_id,
			oitms.shipping_price,
			oitms.total_price,
			oitms.status As product_status
		`).
		Joins(`Right Join store_order_items As oitms On oitms.order_id = ords.id`).
		Where("oitms.store_id = ?", storeID).
		Where("ords.deleted_at Is Null And oitms.deleted_at Is Null").
		Scan(modelOrders)
}
func (repository *RepositoryOrder) ReadByCustomerID(modelOrders *[]models.CustomerOrders, customerID uint64) {
	repository.DB.Table("store_orders As ords").
		Select(`
			ords.id AS order_id,
			ords.status AS order_status,
			Sum( oitms.total_price ) As total_price,
			billing_address_id,
			shipping_address_id
		`).
		Joins(`Right Join store_order_items As oitms On oitms.order_id = ords.id`).
		Where("ords.customer_id = ?", customerID).
		Where("ords.deleted_at Is Null And oitms.deleted_at Is Null").
		Group("oitms.order_id").
		Scan(modelOrders)
}

func (repository *RepositoryOrder) ReadByOrderID(modelOrder *models.CustomerOrdersWithAddress, orderID uint64) {
	modelOrder.Items = make([]models.CustomerOrderItems, 0)
	repository.DB.Table("store_orders As ords").
		Select(`
			ords.status As order_status,
			oitms.store_id,
			oitms.status As product_status,
			oitms.variation_id,
			oitms.price,
			oitms.quantity,
			oitms.sub_total_price,
			oitms.tax_rate,
			oitms.tax_amount,
			oitms.shipping_method_id,
			oitms.shipping_price,
			oitms.total_price,
			ords.billing_address_id,
			ords.shipping_address_id
		`).
		Joins(`Right Join store_order_items As oitms On oitms.order_id = ords.id`).
		Where("ords.id = ?", orderID).
		Where("ords.deleted_at Is Null And oitms.deleted_at Is Null").
		Scan(&modelOrder.Items)
	if len(modelOrder.Items) > 0 {
		billingAdddressID := modelOrder.Items[0].BillingAddressID
		shippingAdddressID := modelOrder.Items[0].ShippingAddressID
		addrRepo := NewRepositoryCustomer(repository.DB)
		addrRepo.ReadAddressByID(&modelOrder.BillingAddress, billingAdddressID)
		addrRepo.ReadAddressByID(&modelOrder.ShippingAddress, shippingAdddressID)
	}
}
