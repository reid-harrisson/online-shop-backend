package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	revsvc "OnlineStoreBackend/services/reviews"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	updateReview = models.Reviews{
		ProductID:  1,
		CustomerID: 1,
		Comment:    "comment1",
		Rate:       1,
		Status:     1,
	}
)

func TestUpdateStatus(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetStoresDB(db)
	test_utils.ResetProductsDB(db)

	// Setup
	serviceReview := revsvc.NewServiceReview(db)

	// Assertions
	assert.NoError(t, serviceReview.UpdateStatus(1, &updateReview, "utils.Pending"))
}
