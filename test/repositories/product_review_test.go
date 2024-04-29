package test

import (
	"OnlineStoreBackend/models"
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/repositories"
	"testing"

	// nolint
	"github.com/stretchr/testify/assert"
)

var (
	readPublishedReviews = []models.Reviews{
		{
			ProductID:  1,
			CustomerID: 1,
			Comment:    "comment1",
			Rate:       0,
			Status:     1,
		},
	}
)

func TestReadPublishReviews(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetProductReviewDB(db)

	// Setup
	modelPublishedReviews := []models.Reviews{}
	repositoryReview := repositories.NewRepositoryReview(db)

	// Assertions
	if assert.NoError(t, repositoryReview.ReadPublishReviews(&modelPublishedReviews, 1)) {
		readPublishedReviews[0].Model.ID = modelPublishedReviews[0].Model.ID
		readPublishedReviews[0].CreatedAt = modelPublishedReviews[0].CreatedAt
		readPublishedReviews[0].UpdatedAt = modelPublishedReviews[0].UpdatedAt

		assert.Equal(t, readPublishedReviews[0], modelPublishedReviews[0])
	}
}

func TestReadAllReviews(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetProductReviewDB(db)

	// Setup
	modelReviews := []models.Reviews{}
	repositoryReview := repositories.NewRepositoryReview(db)

	// Assertions
	if assert.NoError(t, repositoryReview.ReadReviews(&modelReviews, 1)) {
		readPublishedReviews[0].Model.ID = modelReviews[0].Model.ID
		readPublishedReviews[0].CreatedAt = modelReviews[0].CreatedAt
		readPublishedReviews[0].UpdatedAt = modelReviews[0].UpdatedAt

		assert.Equal(t, readPublishedReviews[0], modelReviews[0])
	}
}
