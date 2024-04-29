package test

import (
	test_utils "OnlineStoreBackend/pkgs/test"
	"OnlineStoreBackend/requests"
	chansvc "OnlineStoreBackend/services/related_channels"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	channelRequest = requests.RequestProductChannel{
		ChannelIDs: []uint64{2},
	}
)

func TestUpdateRelatedChannels(t *testing.T) {
	cfg := test_utils.PrepareAllConfiguration("./../../../config.test.yaml")

	// DB Connection
	db := test_utils.InitTestDB(cfg)
	test_utils.ResetProductsDB(db)
	test_utils.ResetRelatedChannelsDB(db)

	// Setup
	var chanService = chansvc.NewServiceProductChannel(db)

	// Assertions
	assert.NoError(t, chanService.Update(1, &channelRequest))
}
