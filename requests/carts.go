package requests

type RequestCart struct {
	AttributeValueIDs []uint64 `json:"attribute_value_ids" example:"1,2,3"`
}
