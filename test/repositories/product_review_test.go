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
			Status:     0,
		},
		{
			ProductID:  1,
			CustomerID: 1,
			Comment:    "comment1",
			Rate:       0,
			Status:     0,
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
		assert.Equal(t, 0, len(modelPublishedReviews))
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
		if assert.Equal(t, len(readPublishedReviews), len(modelReviews)) {
			readPublishedReviews[0].Model.ID = modelReviews[0].Model.ID
			readPublishedReviews[0].CreatedAt = modelReviews[0].CreatedAt
			readPublishedReviews[0].UpdatedAt = modelReviews[0].UpdatedAt

			assert.Equal(t, readPublishedReviews[0], modelReviews[0])
		}
	}
}
