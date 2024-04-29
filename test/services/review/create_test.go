package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/requests"
	revsvc "OnlineStoreBackend/services/reviews"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	requestReview = requests.RequestReview{
		Rate:    1,
		Comment: "1",
	}
)

func TestCreateReview(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetProductReviewDB(db)

	// Setup
	serviceReview := revsvc.NewServiceReview(db)
	modelReview := models.Reviews{}

	// Assertions
	assert.NoError(t, serviceReview.Create(&modelReview, &requestReview, 1, 1))
}
