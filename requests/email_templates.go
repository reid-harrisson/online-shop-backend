package requests

import validation "github.com/go-ozzo/ozzo-validation"

type RequestEmailTemplate struct {
	StoreID                    uint64 `json:"store_id" example:"1"`
	OrderStatus                uint64 `json:"order_status" example:"3"`
	CompanyName                string `json:"company_name" example:"Company Name"`
	CompanyLink                string `json:"company_link" example:"Company Link"`
	CompanyLogoUrl             string `json:"company_logo_url" example:"Company Logo Url"`
	CompanyPrimaryColor        string `json:"company_primary_color" example:"Company Primary Color"`
	EmailPretext               string `json:"email_pretext" example:"Email Pretext"`
	HeaderPosterSloganTitle    string `json:"header_poster_slogan_title" example:"Header Poster Slogan Title"`
	HeaderPosterSloganSubtitle string `json:"header_poster_slogan_subtitle" example:"Header Poster Slogan Subtitle"`
	BodyGreeting               string `json:"body_greeting" example:"Body Greeting"`
	FirstName                  string `json:"first_name" example:"First Name"`
	BodyMessage                string `json:"body_message" example:"Body Message"`
	BodyCtaBtnLink             string `json:"body_cta_btn_link" example:"Body Cta Btn Link"`
	BodyCtaBtnLabel            string `json:"body_cta_btn_label" example:"Body Cta Btn Label"`
	BodySecondaryMessage       string `json:"body_secondary_message" example:"Body Secondary Message"`
	UnsubscribeLink            string `json:"unsubscribe_link" example:"Unsubscribe Link"`
	UnsubscribeSafeLink        string `json:"unsubscribe_safe_link" example:"Unsubscribe Safe Link"`
}

func (request RequestEmailTemplate) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.StoreID, validation.Required),
		validation.Field(&request.OrderStatus, validation.Required),
	)
}
