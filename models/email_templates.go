package models

import "gorm.io/gorm"

type EmailTemplate struct {
	gorm.Model

	StoreID                    uint64 `gorm:"column:store_id; type:bigint(20) unsigned"`
	OrderStatus                uint64 `gorm:"column:order_status; type:tinyint(4)"`
	CompanyName                string `gorm:"column:company_name; type:varchar(100)"`
	CompanyLink                string `gorm:"column:company_link; type:varchar(100)"`
	CompanyLogoUrl             string `gorm:"column:company_logo_url; type:varchar(100)"`
	CompanyPrimaryColor        string `gorm:"column:company_primary_color; type:varchar(100)"`
	EmailPretext               string `gorm:"column:email_pretext; type:varchar(100)"`
	HeaderPosterSloganTitle    string `gorm:"column:header_poster_slogan_title; type:varchar(100)"`
	HeaderPosterSloganSubtitle string `gorm:"column:header_poster_slogan_subtitle; type:varchar(100)"`
	BodyGreeting               string `gorm:"column:body_greeting; type:varchar(100)"`
	FirstName                  string `gorm:"column:first_name; type:varchar(100)"`
	BodyMessage                string `gorm:"column:body_message; type:varchar(100)"`
	BodyCtaBtnLink             string `gorm:"column:body_cta_btn_link; type:varchar(100)"`
	BodyCtaBtnLabel            string `gorm:"column:body_cta_btn_label; type:varchar(100)"`
	BodySecondaryMessage       string `gorm:"column:body_secondary_message; type:varchar(100)"`
	UnsubscribeLink            string `gorm:"column:unsubscribe_link; type:varchar(100)"`
	UnsubscribeSafeLink        string `gorm:"column:unsubscribe_safe_link; type:varchar(100)"`
}

func (EmailTemplate) TableName() string {
	return "store_email_templates"
}
