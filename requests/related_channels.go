package requests

type RequestProductChannel struct {
	ChannelIDs []uint64 `json:"channels" example:"173,174,175"`
}
