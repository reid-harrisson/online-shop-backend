package test_utils

import (
	"OnlineStoreBackend/db"
	"OnlineStoreBackend/pkgs/config"

	"gorm.io/gorm"
)

func InitTestDB(cfg *config.Config) *gorm.DB {
	db := db.Init(cfg)

	return db
}

func PrepareAllConfiguration(path string) *config.Config {
	cfg, err := config.Load([]string{path}, true, nil)
	if err != nil {
		panic(err)
	}

	return cfg
}

func ResetUsersDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table users")

	db.Exec("INSERT INTO `users` (`id`, `deleted_at`, `created_at`, `updated_at`, `first_name`, `last_name`, `gender`, `age`, `dob`, `national_id`, `email`, `mobile_no`, `auth_id`, `webauthn_id`, `webauthn_key`, `password`, `website`, `opt_email`, `opt_sms`, `opt_news`, `opt_marketing`, `gdpr_consent`, `ip_creation`, `ip_last_login`, `ip_last_modified`, `last_login_date`, `profile_image`, `profile_level_id`, `profile_points`, `profile_bio`, `profile_status_id`, `address_line1`, `address_line2`, `suburb`, `city_id`, `region_id`, `country_id`, `postal_code`, `skills`, `active`, `company_id`, `activation_code`, `position`, `skype`, `tours_enabled`, `tours_first_play_status`, `force_page`, `trial_sku`, `ui_sounds`, `crm_contact_id`, `is_journo`, `save_billing_pref`, `wallet_balance`, `token_balance`, `platform_upload_token_youtube`, `platform_upload_refresh_token_youtube`, `platform_upload_default_channel_youtube`, `ad_license_status`, `license_start_date`, `license_end_date`, `reward_count`) VALUES (1, NULL, '2023-11-14 16:54:21', '2023-11-14 16:54:21', 'Steve', 'Etberg', 'Male', 39, '1981-03-26', '5501015009082', 'steve@pockittv.com', '+27.79.326.7663', 'steve@pockittv.com', NULL, NULL, '$2y$10$f9A42Mq5WEkVv50Ns05gde8ETUZsGJ6AyZkaCVvQQ8db5cGltXOj2', 'www.pockittv.com', 1, 1, 1, 1, 1, '160.153.128.39', '197.234.243.67', '197.234.243.67', '2023-09-20 10:56:19', '47.png', 2, 5, 'This is my bio...', 1, 'Mozart', 'Arizona', 'Honeydew', 1, 2, 1, '2000', 'These are my skills', 1, 1, '5e7e910e63432', 'CTO', 'steveetberg', 0, 'accedit|0,contentview|0,contentsyndicationview|0,contentviewlive|0,adupload|0,adedit|0,ad|0,admanage|0,contentupload|0,contentedit|0,content|0,contentmanage|0,livebroadcasts|0,livebroadcastedit|0,createliveevent|0,contentsyndication|0,contentsyndicationmanage|0,setgenres|0,setplevels|0,setpstatuses|0,setpbadges|0,setroles|0,setgroups|0,setlocations|0,setusers|0,setcompanies|0,setcompanysociallinks|0,setfooterlinkcategories|0,setcompanyfooterlinks|0,sysaudit|0,setsystemsocialplatforms|0,setexchangerates|0,setcompanytheme|0,setnotifications|0,accnotify|1,setaffordabilityindex|1,setcountries|1,setbillingcycles|1,setfeaturecategories|1,setfeatures|1,setlicenseunittypes|1,setlicensetypes|1,setproducts|1,setproductbundles|1,trial-setup|1,adminhome|1', NULL, NULL, 0, NULL, 0, 0, 0.0000, 0, 'ya29.a0AfB_byCb5Nr259xuEMZPIj32VGHEQwDbYVwYfAsq2pkeEryIx0OBKkJ2xoAnqcvwXTMTcFm_WNjvWJ4o-X1DN5HAghT-m6eekYZvAGci64T9FDLtVgd4JqlfhukngjyB-F6Vk39lnJL-H_YA1ZiiAGQ6ebC6GopXyrQl0gaCgYKAQUSARASFQGOcNnC9zp1dSMGjF6abakh1equgA0173', '1//03G-1mSobEWWkCgYIARAAGAMSNwF-L9IrkmJRtY8HqZlFxTvFvvdGXQPYu_OF_8xZBsjdeX3maXD_ijb-GAMzS30F55Ajj6iPBLE', NULL, NULL, NULL, NULL, NULL);")
	db.Exec("INSERT INTO `users` (`id`, `deleted_at`, `created_at`, `updated_at`, `first_name`, `last_name`, `gender`, `age`, `dob`, `national_id`, `email`, `mobile_no`, `auth_id`, `webauthn_id`, `webauthn_key`, `password`, `website`, `opt_email`, `opt_sms`, `opt_news`, `opt_marketing`, `gdpr_consent`, `ip_creation`, `ip_last_login`, `ip_last_modified`, `last_login_date`, `profile_image`, `profile_level_id`, `profile_points`, `profile_bio`, `profile_status_id`, `address_line1`, `address_line2`, `suburb`, `city_id`, `region_id`, `country_id`, `postal_code`, `skills`, `active`, `company_id`, `activation_code`, `position`, `skype`, `tours_enabled`, `tours_first_play_status`, `force_page`, `trial_sku`, `ui_sounds`, `crm_contact_id`, `is_journo`, `save_billing_pref`, `wallet_balance`, `token_balance`, `platform_upload_token_youtube`, `platform_upload_refresh_token_youtube`, `platform_upload_default_channel_youtube`, `ad_license_status`, `license_start_date`, `license_end_date`, `reward_count`) VALUES (2, NULL, '2023-11-14 16:54:21', '2023-11-14 16:54:21', 'Tade', 'Upton', 'Female', 34, NULL, NULL, 'jade@pockittv.com', '0824721073', 'jade@pockittv.com', NULL, NULL, '$2y$10$yMiOqjE3W5mmFFsewTPgjeyDMdHwDTJz9BD3Jk.vIbqjUhuU1EBiG', NULL, 0, 0, 0, 0, 0, '172.68.42.7', '197.234.242.42', '197.234.242.42', '2023-08-08 06:53:28', NULL, 1, 0, NULL, 1, NULL, NULL, NULL, 2, 1, 2, NULL, NULL, 1, 2, '5eea1b7033524', 'CCO', 'jade.upton18', 1, 'accnotify|0,accedit|0,contentview|0,contentsyndicationview|0,contentviewlive|0,adupload|0,adedit|0,ad|0,admanage|0,contentupload|0,contentedit|0,content|0,contentmanage|0,livebroadcasts|0,livebroadcastedit|0,createliveevent|0,contentsyndication|0,contentsyndicationmanage|0,setgenres|0,setplevels|0,setpstatuses|0,setpbadges|0,setroles|0,setgroups|0,setlocations|0,setusers|0,setcompanies|0,setcompanysociallinks|0,setfooterlinkcategories|0,setcompanyfooterlinks|0,sysaudit|0,setsystemsocialplatforms|0,setexchangerates|0,setcompanytheme|0,setnotifications|0,setaffordabilityindex|0,setcountries|0,setbillingcycles|0,setfeaturecategories|0,setfeatures|0,setlicenseunittypes|0,setlicensetypes|0,setproducts|0,setproductbundles|0,trial-setup|0,adminhome|0', NULL, NULL, 1, NULL, 0, 0, 0.0000, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetCompaniesDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table companies")

	db.Exec("INSERT INTO `companies` (`id`, `deleted_at`, `created_at`, `updated_at`, `name`, `city_id`, `region_id`, `country_id`, `active`, `company_type_id`, `theme_layout`, `theme_color_1`, `theme_color_2`, `theme_color_3`, `theme_color_4`, `theme_color_5`, `theme_color_6`, `theme_color_7`, `theme_color_8`, `theme_color_9`, `theme_color_10`, `theme_color_mode`, `logo_1`, `logo_2`, `logo_3`, `logo_4`, `logo_5`, `logo_6`, `logo_7`, `logo_8`, `logo_9`, `logo_10`, `system_mode`, `subdomain`, `domain`, `brand_logo_text`, `workflow_status`, `ad_rotation_setting`, `api_key`, `api_secret`, `api_domain`, `default_currency`, `brand_slogan`, `domain_status`, `trial_start_date`, `trial_duration_days`, `theme_background`, `theme_bg_color_1`, `theme_bg_color_2`, `use_bg_image`, `crm_company_id`, `video_bg_1`, `video_bg_2`, `video_bg_3`, `video_bg_4`, `video_bg_5`, `video_bg_6`, `video_bg_7`, `video_bg_8`, `video_bg_9`, `video_bg_10`, `play_pause_ads`, `ad_budget_amount`, `ad_budget_used`, `ad_budget_available`, `ad_price_per_view`, `ad_price_per_clickthrough`, `ad_budget_currency_code`, `rev_share_ad_server`, `rev_share_ad_server_min`, `rev_share_tx_server`, `rev_share_tx_server_min`, `force_single_currency`) VALUES (1, NULL, '2023-11-13 22:53:18', '2023-11-13 22:53:18', 'POCKIT TV Global', 1, 2, 1, 1, 2, 0, '#e4fe20', '#d8ff2c', '#ff1493', '#cb11e4', '#2eb9ff', '#fbc80e', '#7e2ae5', '#da2bc2', '#2eb9ff', 'dark', 'https://app.pockittv.com/images/companies/2/h', 'https://app.pockittv.com/images/companies/2/footer/logo140x34.png', 'https://app.pockittv.com/images/companies/2/favicon/Pockit Icon.png', 'https://www.pockittv.com/images/mail/companies/2/pockittv_logo_chrome_glow_260x64.png', 'https://app.pockittv.com/images/mail/companies/2/pockittv_logo_chrome_glow_260x64.png', NULL, NULL, NULL, NULL, NULL, 'live', 'app', 'pockittv.com', '', 'Complete', 'Randomize', 'e61d32dc2a872e8c8f302d2ce7fab6f43d18b16e61614', '', 'www.pockittv.com,biz.pockittv.com,app.pockittv.com,southafrica.pockittv.com,app.pockittv.dev', 'USD', 'You', '1', 0, NULL, NULL, '#333333', '#000000', '0', NULL, 0, '', '', '', '', '', '', '', '', '', '0', 0, 0.00, 0.00, 0.00, 0.00000, 0.00000, '0.0000', 0.0000, 0.0000, 0.0000, NULL, NULL);")
	db.Exec("INSERT INTO `companies` (`id`, `deleted_at`, `created_at`, `updated_at`, `name`, `city_id`, `region_id`, `country_id`, `active`, `company_type_id`, `theme_layout`, `theme_color_1`, `theme_color_2`, `theme_color_3`, `theme_color_4`, `theme_color_5`, `theme_color_6`, `theme_color_7`, `theme_color_8`, `theme_color_9`, `theme_color_10`, `theme_color_mode`, `logo_1`, `logo_2`, `logo_3`, `logo_4`, `logo_5`, `logo_6`, `logo_7`, `logo_8`, `logo_9`, `logo_10`, `system_mode`, `subdomain`, `domain`, `brand_logo_text`, `workflow_status`, `ad_rotation_setting`, `api_key`, `api_secret`, `api_domain`, `default_currency`, `brand_slogan`, `domain_status`, `trial_start_date`, `trial_duration_days`, `theme_background`, `theme_bg_color_1`, `theme_bg_color_2`, `use_bg_image`, `crm_company_id`, `video_bg_1`, `video_bg_2`, `video_bg_3`, `video_bg_4`, `video_bg_5`, `video_bg_6`, `video_bg_7`, `video_bg_8`, `video_bg_9`, `video_bg_10`, `play_pause_ads`, `ad_budget_amount`, `ad_budget_used`, `ad_budget_available`, `ad_price_per_view`, `ad_price_per_clickthrough`, `ad_budget_currency_code`, `rev_share_ad_server`, `rev_share_ad_server_min`, `rev_share_tx_server`, `rev_share_tx_server_min`, `force_single_currency`) VALUES (2, NULL, '2023-11-13 22:53:18', '2023-11-13 22:53:18', 'AFDA', 2, 1, 2, 5, 1, 0, '#efb538', '', '', '#c61d23', '#efb538', '#f5d184', '#efb538', '#c61d23', '#610005', 'dark', 'images/companies/60/header/download.png', 'images/companies/60/footer/download.png', 'images/companies/60/favicon/cropped-favicon-1-192x192.png', '', '', '', '', '', '', '', 'live', 'afda', 'pockittv.com', '', 'Complete', 'Randomize', 'e61d32dc2a872e8c8f302d2ce7fab6f43d18b16e61614', '', 'afda.pockittv.com', 'USD', 'The', '1', 0, NULL, 0, '#5c5c5c', '#262626', '1', 127, 0, '', '', '', '', '', '', '', '', '', '0', 0, 0.00, 0.00, 0.00, 0.00000, 0.00000, '0.0000', 0.0000, 0.0000, 0.0000, NULL, NULL);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetCitiesDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table cities")

	db.Exec("INSERT INTO `pockittv_test_db`.`cities` (`id`, `deleted_at`, `created_at`, `updated_at`, `region_id`, `country_id`, `latitude`, `longitude`, `name`, `active`) VALUES (1, NULL, '2023-11-13 20:14:26', '2023-11-13 20:14:26', 1, 1, 42.50000000, 1.51666670, 'Andorra', 1);")
	db.Exec("INSERT INTO `pockittv_test_db`.`cities` (`id`, `deleted_at`, `created_at`, `updated_at`, `region_id`, `country_id`, `latitude`, `longitude`, `name`, `active`) VALUES (2, NULL, '2023-11-13 20:14:26', '2023-11-13 20:14:26', 2, 2, 42.56666670, 1.51666670, 'Ansalonga', 1);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetStoresDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table stores")

	db.Exec("INSERT INTO `stores` (`id`, `deleted_at`, `created_at`, `updated_at`, `company_id`, `owner_id`, `contact_phone`, `contact_email`, `show_stock_level_status`, `show_out_of_stock_status`, `delivery_policy`, `returns_policy`, `terms`, `name`) VALUES (1, NULL, '2024-04-08 09:57:59', '2024-04-08 09:58:51', 1, 1, '+27793267663', 'steve@pockittv.com', 0, 0, '', '', '', 'Steve');")
	db.Exec("INSERT INTO `stores` (`id`, `deleted_at`, `created_at`, `updated_at`, `company_id`, `owner_id`, `contact_phone`, `contact_email`, `show_stock_level_status`, `show_out_of_stock_status`, `delivery_policy`, `returns_policy`, `terms`, `name`) VALUES (2, NULL, '2024-04-08 09:58:05', '2024-04-08 09:58:59', 2, 2, '+0824721073', 'jade@pockittv.com', 0, 0, '', '', '', 'Tade');")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetStoreOrdersDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_orders")

	db.Exec("INSERT INTO `store_orders` (`id`, `deleted_at`, `created_at`, `updated_at`, `customer_id`, `status`, `billing_address_id`, `shipping_address_id`) VALUES (1, NULL, '2024-04-26 16:29:54', '2024-04-29 09:16:01', 1, 0, 1, 1);")
	db.Exec("INSERT INTO `store_orders` (`id`, `deleted_at`, `created_at`, `updated_at`, `customer_id`, `status`, `billing_address_id`, `shipping_address_id`) VALUES (2, NULL, '2024-04-29 09:15:58', '2024-04-29 09:16:02', 2, 0, 2, 2);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetStoreCustomerAddressesDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_orders")

	db.Exec("INSERT INTO `store_customer_addresses` (`id`, `deleted_at`, `created_at`, `updated_at`, `customer_id`, `country_id`, `region_id`, `city_id`, `postal_code`, `address_line1`, `address_line2`, `suburb`, `active`) VALUES (1, NULL, '2024-04-26 16:29:38', '2024-04-26 16:29:46', 1, 1, 1, 1, NULL, NULL, NULL, NULL, 1);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetProductsDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_products")

	db.Exec("INSERT INTO `store_products` (`id`, `deleted_at`, `created_at`, `updated_at`, `store_id`, `title`, `short_description`, `long_description`, `image_urls`, `minimum_stock_level`, `status`, `sku`, `type`, `shipping_class`) VALUES (1, NULL, '2024-04-08 09:59:09', '2024-04-08 10:13:28', 1, 'Gochujang - Korean Chilli Pepper Paste', 'Gochujang is a Sticky, Sweet, Savoury &amp; SPICY Chilli Paste.', 'Our AMAZING range of products are available nation wide in South Africa at select health stores.', 'https://www.chegourmet.co.za/wp-content/uploads/2019/09/Gochujang-Front-scaled.jpg', 0.000000, 0, '44', 1, 'Courier Refrigerated');")
	db.Exec("INSERT INTO `store_products` (`id`, `deleted_at`, `created_at`, `updated_at`, `store_id`, `title`, `short_description`, `long_description`, `image_urls`, `minimum_stock_level`, `status`, `sku`, `type`, `shipping_class`) VALUES (2, NULL, '2024-04-08 09:59:15', '2024-04-08 10:13:50', 2, 'Kimchi Probiotic Tonic - 200ML', 'This Probiotic Rich Tincture packs a sour spicy Punch!', 'Pour over salads, add to savoury juices, smoothies &amp; cocktails or on its own as a healthy shot.', 'https://www.chegourmet.co.za/wp-content/uploads/2019/09/Kimchi-Tonic-Front-2-scaled.jpg', 0.000000, 0, '13', 1, 'Courier Refrigerated');")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetProductReviewDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_product_reviews")

	db.Exec("INSERT INTO `store_product_reviews` (`id`, `deleted_at`, `created_at`, `updated_at`, `product_id`, `customer_id`, `comment`, `rate`, `status`) VALUES (1, NULL, '2024-04-29 00:53:22', '2024-04-29 00:53:22', 1, 1, 'comment1', 0.000000, 0);")
	db.Exec("INSERT INTO `store_product_reviews` (`id`, `deleted_at`, `created_at`, `updated_at`, `product_id`, `customer_id`, `comment`, `rate`, `status`) VALUES (2, NULL, '2024-04-29 00:56:05', '2024-04-29 00:56:13', 1, 1, 'comment2', 0.000000, 0);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetVariationsDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_product_variations")

	db.Exec("INSERT INTO `store_product_variations` (`id`, `deleted_at`, `created_at`, `updated_at`, `product_id`, `sku`, `price`, `stock_level`, `discount_amount`, `discount_type`, `image_urls`, `description`, `title`, `back_order_status`) VALUES (1, NULL, '2024-04-08 09:59:09.000', '2024-04-08 10:00:43.000', 1, '44-125G', 96.000000, 10.000000, 20.000000, 1, '[]', 'Full cream milk kefir made using live kefir grains/cultures, fermented traditionally for maximum probiotic diversity.', 'Gochujang - Korean Chilli Pepper Paste - 125G', 0);")
	db.Exec("INSERT INTO `store_product_variations` (`id`, `deleted_at`, `created_at`, `updated_at`, `product_id`, `sku`, `price`, `stock_level`, `discount_amount`, `discount_type`, `image_urls`, `description`, `title`, `back_order_status`) VALUES (2, NULL, '2024-04-08 09:59:15.000', '2024-04-08 10:00:49.000', 2, '13-200ML', 45.000000, 10.000000, 0.000000, 0, '[]', 'NuMe Kombucha - 350ml, Buchu, Hibiscus & Hawthorne', 'Kimchi Probiotic Tonic - 200ML - 200ML', 0);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetAddressesDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_customer_addresses")

	db.Exec("INSERT INTO `store_customer_addresses` (`id`, `deleted_at`, `created_at`, `updated_at`, `customer_id`, `country_id`, `region_id`, `city_id`, `postal_code`, `address_line1`, `address_line2`, `suburb`, `active`) VALUES (1, NULL, '2024-04-10 19:40:55', '2024-04-10 20:12:33', 1, 1, 1, 1, '11-111', 'Andorra, Andorra', NULL, 'Andorra la Vella', 1);")
	db.Exec("INSERT INTO `store_customer_addresses` (`id`, `deleted_at`, `created_at`, `updated_at`, `customer_id`, `country_id`, `region_id`, `city_id`, `postal_code`, `address_line1`, `address_line2`, `suburb`, `active`) VALUES (2, NULL, '2024-04-10 19:41:18', '2024-04-10 20:13:01', 2, 2, 2, 2, '22-222', 'Badiya, United Arab Emirates', NULL, 'Fujairah', 1);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetCategoriesDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_categories")

	db.Exec("INSERT INTO `store_categories` (`id`, `deleted_at`, `created_at`, `updated_at`, `name`, `store_id`, `parent_id`) VALUES (1, NULL, '2024-04-10 19:27:40', '2024-04-10 19:29:37', 'Kefir', 1, 0);")
	db.Exec("INSERT INTO `store_categories` (`id`, `deleted_at`, `created_at`, `updated_at`, `name`, `store_id`, `parent_id`) VALUES (2, NULL, '2024-04-10 19:27:54', '2024-04-10 19:29:40', 'Kimchi', 2, 0);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetProductCategoriesDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_product_categories")

	db.Exec("INSERT INTO `store_product_categories` (`id`, `deleted_at`, `created_at`, `updated_at`, `product_id`, `category_id`) VALUES (1, NULL, '2024-04-10 19:32:13', '2024-04-10 19:32:13', 1, 1);")
	db.Exec("INSERT INTO `store_product_categories` (`id`, `deleted_at`, `created_at`, `updated_at`, `product_id`, `category_id`) VALUES (2, NULL, '2024-04-10 19:32:18', '2024-04-10 19:32:18', 2, 2);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetVariationDetailsDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_product_variation_details")

	db.Exec("INSERT INTO `store_product_variation_details` (`id`, `deleted_at`, `created_at`, `updated_at`, `variation_id`, `attribute_value_id`) VALUES (1, NULL, '2024-04-10 19:26:29', '2024-04-10 19:26:29', 1, 1);")
	db.Exec("INSERT INTO `store_product_variation_details` (`id`, `deleted_at`, `created_at`, `updated_at`, `variation_id`, `attribute_value_id`) VALUES (2, NULL, '2024-04-10 19:26:33', '2024-04-10 19:26:33', 2, 2);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetAttributesDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_product_attributes")

	db.Exec("INSERT INTO `store_product_attributes` (`id`, `deleted_at`, `created_at`, `updated_at`, `product_id`, `attribute_name`) VALUES (1, NULL, '2024-04-10 19:25:41', '2024-04-10 19:31:45', 1, 'weight');")
	db.Exec("INSERT INTO `store_product_attributes` (`id`, `deleted_at`, `created_at`, `updated_at`, `product_id`, `attribute_name`) VALUES (2, NULL, '2024-04-10 19:25:45', '2024-04-10 19:31:57', 2, 'volume');")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetAttributeValuesDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_product_attribute_values")

	db.Exec("INSERT INTO `store_product_attribute_values` (`id`, `deleted_at`, `created_at`, `updated_at`, `attribute_id`, `attribute_value`) VALUES (1, NULL, '2024-04-10 19:25:57', '2024-04-10 19:31:34', 1, '125g');")
	db.Exec("INSERT INTO `store_product_attribute_values` (`id`, `deleted_at`, `created_at`, `updated_at`, `attribute_id`, `attribute_value`) VALUES (2, NULL, '2024-04-10 19:26:01', '2024-04-10 19:31:31', 2, '200ml');")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetCombosDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_combos")

	db.Exec("INSERT INTO `store_combos` (`id`, `created_at`, `updated_at`, `deleted_at`, `store_id`, `discount_amount`, `discount_type`, `image_urls`, `description`, `title`, `status`) VALUES (1, '2024-04-10 19:33:25.000', '2024-04-10 19:33:22.000', NULL, 1, 10.000000, 1, 'https://www.chegourmet.co.za/wp-content/uploads/2019/09/Gochujang-Front-scaled.jpg', 'Combo of Kefir', 'Kefir Combo', '2');")
	db.Exec("INSERT INTO `store_combos` (`id`, `created_at`, `updated_at`, `deleted_at`, `store_id`, `discount_amount`, `discount_type`, `image_urls`, `description`, `title`, `status`) VALUES (2, '2024-04-10 19:33:14.000', '2024-04-10 19:33:18.000', NULL, 2, 15.000000, 0, 'https://www.chegourmet.co.za/wp-content/uploads/2019/09/Kimchi-Tonic-Front-2-scaled.jpg', 'Combo of Kimchi', 'Kimchi Combo', '1');")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetComboItemsDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_combo_items")

	db.Exec("INSERT INTO `store_combo_items` (`id`, `created_at`, `updated_at`, `deleted_at`, `combo_id`, `variation_id`, `quantity`) VALUES (1, '2024-04-10 19:27:40.000', '2024-04-10 19:29:37.000', NULL, 1, 1, 2.000000);")
	db.Exec("INSERT INTO `store_combo_items` (`id`, `created_at`, `updated_at`, `deleted_at`, `combo_id`, `variation_id`, `quantity`) VALUES (2, '2024-04-10 19:27:54.000', '2024-04-10 19:29:40.000', NULL, 2, 2, 3.000000);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetTagsDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_tags")

	db.Exec("INSERT INTO `store_tags` (`id`, `deleted_at`, `created_at`, `updated_at`, `name`, `store_id`) VALUES (1, NULL, '2024-04-10 20:21:32', '2024-04-10 20:21:32', 'kefir', 1);")
	db.Exec("INSERT INTO `store_tags` (`id`, `deleted_at`, `created_at`, `updated_at`, `name`, `store_id`) VALUES (2, NULL, '2024-04-10 20:21:42', '2024-04-10 20:21:42', 'kimchi', 2);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetProductTagsDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_product_tags")

	db.Exec("INSERT INTO `store_product_tags` (`id`, `deleted_at`, `created_at`, `updated_at`, `tag_id`, `product_id`) VALUES (1, NULL, '2024-04-10 20:22:09', '2024-04-10 20:22:09', 1, 1);")
	db.Exec("INSERT INTO `store_product_tags` (`id`, `deleted_at`, `created_at`, `updated_at`, `tag_id`, `product_id`) VALUES (2, NULL, '2024-04-10 20:22:12', '2024-04-10 20:22:12', 2, 2);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetCouponsDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_coupons")

	db.Exec("INSERT INTO `store_coupons` (`id`, `created_at`, `updated_at`, `deleted_at`, `store_id`, `coupon_code`, `discount_type`, `coupon_amount`, `allow_free_shipping`, `expiry_date`, `minimum_spend`, `maximum_spend`) VALUES (1, '2024-04-10 19:27:40.000', '2024-04-10 19:29:37.000', NULL, 1, '30PERCENTMIN', 0, 30.000000, 0, '2025-04-10 19:37:59', 30.000000, 100.000000);")
	db.Exec("INSERT INTO `store_coupons` (`id`, `created_at`, `updated_at`, `deleted_at`, `store_id`, `coupon_code`, `discount_type`, `coupon_amount`, `allow_free_shipping`, `expiry_date`, `minimum_spend`, `maximum_spend`) VALUES (2, '2024-04-10 19:27:54.000', '2024-04-10 19:29:40.000', NULL, 2, '20FIXMIN', 1, 20.000000, 1, '2025-04-10 19:37:59', 40.000000, 120.000000);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetVisitorsDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_visitors")

	db.Exec("INSERT INTO `store_visitors` (`id`, `deleted_at`, `created_at`, `updated_at`, `store_id`, `product_id`, `ip_address`, `page`, `bounce`) VALUES (1, NULL, '2024-04-24 20:36:14', '2024-04-24 20:36:17', 1, 1, '111.111.111.111', 1, 1);")
	db.Exec("INSERT INTO `store_visitors` (`id`, `deleted_at`, `created_at`, `updated_at`, `store_id`, `product_id`, `ip_address`, `page`, `bounce`) VALUES (2, NULL, '2024-04-24 20:36:35', '2024-04-24 20:36:35', 2, 2, '111.111.111.112', 2, 2);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetStoreCartItemDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_cart_items")

	db.Exec("INSERT INTO `store_cart_items` (`id`, `created_at`, `updated_at`, `deleted_at`, `customer_id`, `variation_id`, `quantity`) VALUES (1, NULL, NULL, NULL, 1, 1, NULL);")
	db.Exec("INSERT INTO `store_cart_items` (`id`, `created_at`, `updated_at`, `deleted_at`, `customer_id`, `variation_id`, `quantity`) VALUES (2, NULL, NULL, NULL, 2, 2, NULL);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetStockTrailsDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_stock_trails")

	db.Exec("INSERT INTO `store_stock_trails` (`id`, `created_at`, `updated_at`, `deleted_at`, `product_id`, `variation_id`, `change`, `event`) VALUES (1, '2024-04-10 20:21:32.000', '2024-04-10 20:21:32.000', NULL, 1, 1, 10.000000, 0);")
	db.Exec("INSERT INTO `store_stock_trails` (`id`, `created_at`, `updated_at`, `deleted_at`, `product_id`, `variation_id`, `change`, `event`) VALUES (2, '2024-04-10 20:21:42.000', '2024-04-10 20:21:42.000', NULL, 2, 2, 20.000000, 1);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetShippingData(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_shipping_data")

	db.Exec("INSERT INTO `store_shipping_data` (`id`, `deleted_at`, `created_at`, `updated_at`, `variation_id`, `weight`, `width`, `height`, `length`) VALUES (1, NULL, '2024-04-27 02:13:42', '2024-04-27 02:13:42', 1, 5.100000, 11.000000, 27.500000, 16.500000);")
	db.Exec("INSERT INTO `store_shipping_data` (`id`, `deleted_at`, `created_at`, `updated_at`, `variation_id`, `weight`, `width`, `height`, `length`) VALUES (2, NULL, '2024-04-27 02:13:42', '2024-04-27 02:13:42', 2, 2.000000, 8.500000, 26.000000, 12.500000);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetRelatedChannelsDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_product_related_channels")

	db.Exec("INSERT INTO `store_product_related_channels` (`id`, `deleted_at`, `created_at`, `updated_at`, `product_id`, `channel_id`) VALUES (1, NULL, '2024-04-29 03:26:49', '2024-04-29 03:26:49', 1, 1);")
	db.Exec("INSERT INTO `store_product_related_channels` (`id`, `deleted_at`, `created_at`, `updated_at`, `product_id`, `channel_id`) VALUES (2, NULL, '2024-04-29 03:26:51', '2024-04-29 03:26:51', 2, 2);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetRelatedContentsDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_product_related_contents")

	db.Exec("INSERT INTO `store_product_related_contents` (`id`, `deleted_at`, `created_at`, `updated_at`, `product_id`, `content_id`) VALUES (1, NULL, '2024-04-29 03:28:39', '2024-04-29 08:40:15', 1, 1);")
	db.Exec("INSERT INTO `store_product_related_contents` (`id`, `deleted_at`, `created_at`, `updated_at`, `product_id`, `content_id`) VALUES (2, NULL, '2024-04-29 03:28:50', '2024-04-29 03:28:50', 2, 2);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetStoreOrderItemsDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_order_items")

	db.Exec("INSERT INTO `store_order_items` (`id`, `deleted_at`, `created_at`, `updated_at`, `order_id`, `store_id`, `variation_id`, `price`, `quantity`, `sub_total_price`, `tax_rate`, `tax_amount`, `shipping_method_id`, `shipping_price`, `total_price`, `status`) VALUES (1, NULL, '2024-04-29 09:02:26', '2024-04-29 23:48:01', 1, 1, 1, 76.000000, 1.000000, 76.000000, 10.000000, 7.600000, 1, 0.000000, 83.600000, 0);")
	db.Exec("INSERT INTO `store_order_items` (`id`, `deleted_at`, `created_at`, `updated_at`, `order_id`, `store_id`, `variation_id`, `price`, `quantity`, `sub_total_price`, `tax_rate`, `tax_amount`, `shipping_method_id`, `shipping_price`, `total_price`, `status`) VALUES (2, NULL, '2024-04-29 09:02:40', '2024-04-29 09:02:40', 2, 2, 2, 0.000000, 0.000000, 0.000000, 0.000000, 0.000000, 2, 0.000000, 0.000000, 0);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetProductLinksDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_product_links")

	db.Exec("INSERT INTO `store_product_links` (`id`, `deleted_at`, `created_at`, `updated_at`, `product_id`, `link_id`, `is_up_cross`) VALUES (1, NULL, '2024-04-29 04:49:34', '2024-04-29 04:49:34', 1, 1, 0);")
	db.Exec("INSERT INTO `store_product_links` (`id`, `deleted_at`, `created_at`, `updated_at`, `product_id`, `link_id`, `is_up_cross`) VALUES (2, NULL, '2024-04-29 04:49:38', '2024-04-29 04:49:38', 2, 2, 1);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetTemplatesDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_email_templates")

	db.Exec("INSERT INTO `store_email_templates` (`id`, `deleted_at`, `created_at`, `updated_at`, `store_id`, `order_status`, `company_name`, `company_link`, `company_logo_url`, `company_primary_color`, `email_pretext`, `header_poster_slogan_title`, `header_poster_slogan_subtitle`, `body_greeting`, `first_name`, `body_message`, `body_cta_btn_link`, `body_cta_btn_label`, `body_secondary_message`, `unsubscribe_link`, `unsubscribe_safe_link`) VALUES (1, NULL, '2024-05-03 04:29:10', '2024-05-03 04:29:14', 1, 1, 'company1', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);")
	db.Exec("INSERT INTO `store_email_templates` (`id`, `deleted_at`, `created_at`, `updated_at`, `store_id`, `order_status`, `company_name`, `company_link`, `company_logo_url`, `company_primary_color`, `email_pretext`, `header_poster_slogan_title`, `header_poster_slogan_subtitle`, `body_greeting`, `first_name`, `body_message`, `body_cta_btn_link`, `body_cta_btn_label`, `body_secondary_message`, `unsubscribe_link`, `unsubscribe_safe_link`) VALUES (2, NULL, '2024-05-03 04:29:24', '2024-05-03 04:29:29', 2, 2, 'company2', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetShippingMethodsDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_shipping_methods")

	db.Exec("INSERT INTO `store_shipping_methods` (`id`, `created_at`, `updated_at`, `deleted_at`, `zone_id`, `store_id`, `method`, `requirement`, `minimum_order_amount`, `tax_status`, `cost`, `tax_included`, `handling_fee`, `maximum_shipping_cost`, `calculation_type`, `handling_fee_per_class`, `minimum_cost_per_class`, `maximum_cost_per_class`, `discount_in_min_max`, `tax_in_min_max`, `title`, `description`) VALUES (1, NULL, NULL, NULL, NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);")
	db.Exec("INSERT INTO `store_shipping_methods` (`id`, `created_at`, `updated_at`, `deleted_at`, `zone_id`, `store_id`, `method`, `requirement`, `minimum_order_amount`, `tax_status`, `cost`, `tax_included`, `handling_fee`, `maximum_shipping_cost`, `calculation_type`, `handling_fee_per_class`, `minimum_cost_per_class`, `maximum_cost_per_class`, `discount_in_min_max`, `tax_in_min_max`, `title`, `description`) VALUES (2, NULL, NULL, NULL, NULL, 2, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetShippingTableRatesDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table store_shipping_table_rates")

	db.Exec("INSERT INTO `store_shipping_table_rates` (`id`, `created_at`, `updated_at`, `deleted_at`, `method_id`, `class_id`, `condition`, `min`, `max`, `break`, `abort`, `row_cost`, `item_cost`, `cost_per_kg`, `percent_cost`) VALUES (1, NULL, NULL, NULL, 1, NULL, 0, NULL, NULL, NULL, NULL, 2.000000, 1.000000, 0.200000, 1.000000);")
	db.Exec("INSERT INTO `store_shipping_table_rates` (`id`, `created_at`, `updated_at`, `deleted_at`, `method_id`, `class_id`, `condition`, `min`, `max`, `break`, `abort`, `row_cost`, `item_cost`, `cost_per_kg`, `percent_cost`) VALUES (2, NULL, NULL, NULL, 2, NULL, 1, NULL, NULL, NULL, NULL, 3.000000, 2.000000, 0.100000, 1.100000);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func ResetCountriesDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	db.Exec("TRUNCATE Table countries")

	db.Exec("INSERT INTO `countries` (`id`, `deleted_at`, `created_at`, `updated_at`, `name`, `code`, `active`, `dial_code`, `currency_code`, `usd_ad_view_fee`, `usd_ad_click_fee`, `affordability_index_id`, `sim_ip`, `tax_rate`, `apply_tax`) VALUES (1, NULL, '2023-11-28 22:46:32', '2024-04-26 06:58:43', 'Andorra', 'ad', 1, '+376', 'EUR', 1.00, 2.00, 1, '46.172.224.0', 4.500000, 0);")
	db.Exec("INSERT INTO `countries` (`id`, `deleted_at`, `created_at`, `updated_at`, `name`, `code`, `active`, `dial_code`, `currency_code`, `usd_ad_view_fee`, `usd_ad_click_fee`, `affordability_index_id`, `sim_ip`, `tax_rate`, `apply_tax`) VALUES (2, NULL, '2023-11-28 22:46:32', '2024-05-02 15:02:26', 'Polska', 'pl', 1, '+48', 'PLN', 1.00, 2.00, 2, '5.57.128.0', 23.000000, 0);")

	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}
