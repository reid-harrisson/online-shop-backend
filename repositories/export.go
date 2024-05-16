package repositories

import (
	"OnlineStoreBackend/models"
	"OnlineStoreBackend/pkgs/utils"
	"encoding/json"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type RepositoryExport struct {
	DB *gorm.DB
}

func NewRepositoryExport(db *gorm.DB) *RepositoryExport {
	return &RepositoryExport{DB: db}
}

func (repository *RepositoryExport) ReadAll(result *[][]string, storeID uint64) error {
	columns := []string{"ID", "Type", "SKU", "Name", "Published", "Is featured?", "Visibility in catalogue", "Short description", "Description", "Date sale price starts", "Date sale price ends", "Tax status", "Tax class", "In stock?", "Stock", "Low stock amount", "Backorders allowed?", "Sold individually?", "Weight (kg)", "Length (cm)", "Width (cm)", "Height (cm)", "Allow customer reviews?", "Purchase note", "Sale price", "Regular price", "Categories", "Tags", "Shipping class", "Images", "Download limit", "Download expiry days", "Parent", "Grouped products", "Upsells", "Cross-sells", "External URL", "Button text", "Position", "Meta: fusion_builder_status", "Meta: _yoast_wpseo_primary_product_cat", "Meta: _yoast_wpseo_content_score", "Meta: _yoast_wpseo_metadesc", "Meta: pyre_main_top_padding", "Meta: pyre_main_bottom_padding", "Meta: pyre_hundredp_padding", "Meta: pyre_display_header", "Meta: pyre_header_100_width", "Meta: pyre_header_bg_color", "Meta: pyre_header_bg_opacity", "Meta: pyre_header_bg", "Meta: pyre_header_bg_full", "Meta: pyre_header_bg_repeat", "Meta: pyre_displayed_menu", "Meta: pyre_display_footer", "Meta: pyre_display_copyright", "Meta: pyre_footer_100_width", "Meta: pyre_sidebar_position", "Meta: pyre_sidebar_bg_color", "Meta: pyre_slider_type", "Meta: pyre_slider", "Meta: pyre_wooslider", "Meta: pyre_revslider", "Meta: pyre_elasticslider", "Meta: pyre_slider_position", "Meta: pyre_avada_rev_styles", "Meta: pyre_fallback", "Meta: pyre_demo_slider", "Meta: pyre_page_bg_layout", "Meta: pyre_page_bg_color", "Meta: pyre_page_bg", "Meta: pyre_page_bg_full", "Meta: pyre_page_bg_repeat", "Meta: pyre_wide_page_bg_color", "Meta: pyre_wide_page_bg", "Meta: pyre_wide_page_bg_full", "Meta: pyre_wide_page_bg_repeat", "Meta: pyre_page_title", "Meta: pyre_page_title_breadcrumbs_search_bar", "Meta: pyre_page_title_text", "Meta: pyre_page_title_text_alignment", "Meta: pyre_page_title_custom_text", "Meta: pyre_page_title_text_size", "Meta: pyre_page_title_custom_subheader", "Meta: pyre_page_title_custom_subheader_text_size", "Meta: pyre_page_title_font_color", "Meta: pyre_page_title_100_width", "Meta: pyre_page_title_height", "Meta: pyre_page_title_mobile_height", "Meta: pyre_page_title_bar_bg_color", "Meta: pyre_page_title_bar_borders_color", "Meta: pyre_page_title_bar_bg", "Meta: pyre_page_title_bar_bg_retina", "Meta: pyre_page_title_bar_bg_full", "Meta: pyre_page_title_bg_parallax", "Meta: site-sidebar-layout", "Meta: site-content-layout", "Meta: theme-transparent-header-meta", "Meta: _yoast_wpseo_focuskw", "Meta: _yoast_wpseo_linkdex", "Meta: product_quantity_per_parcel", "Meta: _ywsbs_subscription", "Meta: _ywsbs_price_is_per", "Meta: _ywsbs_price_time_option", "Meta: _ywsbs_max_length", "Meta: mep_list_thumbnail", "Meta: mep_event_cc_email_text", "Meta: fb_product_group_id", "Meta: _ywsbs_enable_limit", "Meta: _ywsbs_enable_max_length", "Meta: _ywsbs_limit", "Meta: _wc_facebook_sync_enabled", "Meta: fb_visibility", "Meta: fb_product_description", "Meta: _wc_facebook_product_image_source", "Meta: _wc_facebook_commerce_enabled", "Meta: fb_product_item_id", "Meta: product_free_shipping", "Meta: product_single_parcel", "Meta: product_prohibit_tcg", "Meta: mep_event_faq", "Meta: mep_event_day", "Meta: _yoast_wpseo_estimated-reading-time-minutes", "Meta: _wc_gla_mc_status", "Meta: _wc_gla_synced_at", "Meta: _wc_gla_sync_status", "Meta: _wc_gla_visibility", "Meta: _yoast_wpseo_wordproof_timestamp", "Meta: _dp_original", "Meta: wpiudacb_disable_add_to_cart", "Meta: wpiudacb_inqure_us_link", "Attribute 1 name", "Attribute 1 value(s)", "Attribute 1 visible", "Attribute 1 global", "Meta: product_length_per_parcel", "Meta: product_width_per_parcel", "Meta: product_height_per_parcel", "Meta: _wc_facebook_enhanced_catalog_attributes_color", "Meta: _wc_facebook_enhanced_catalog_attributes_size", "Meta: _wc_facebook_enhanced_catalog_attributes_gender", "Meta: _wc_facebook_enhanced_catalog_attributes_brand", "Meta: _wc_facebook_enhanced_catalog_attributes_pattern", "Meta: _wc_facebook_enhanced_catalog_attributes_material", "Meta: _wc_facebook_enhanced_catalog_attributes_age_group", "Meta: product_free_shipping_pudo", "Meta: product_single_parcel_pudo", "Meta: product_prohibit_pudo", "Meta: _yoast_wpseo_primary_fb_product_set", "Meta: _wp_page_template", "Meta: _wp_old_date", "Meta: rs_page_bg_color", "Meta: _wpb_vc_js_status", "Meta: _wds_meta-robots-adv", "Meta: _wds_trimmed_excerpt", "Meta: _wc_facebook_google_product_category", "Meta: _wds_meta-robots-nofollow", "Meta: _wds_meta-robots-noindex", "Meta: fb_product_image", "Meta: fb_product_price", "Meta: _wds_metadesc", "Meta: _wds_title", "Meta: _wds_focus-keywords", "Meta: _wds_canonical", "Meta: _wc_facebook_enhanced_catalog_attributes_power_type", "Meta: _wc_facebook_enhanced_catalog_attributes_product_form", "Meta: _wc_facebook_enhanced_catalog_attributes_result_time", "Meta: _wc_facebook_enhanced_catalog_attributes_scent", "Meta: _wc_facebook_enhanced_catalog_attributes_absorbency", "Meta: _wc_facebook_enhanced_catalog_attributes_serving_size", "Meta: _wc_facebook_enhanced_catalog_attributes_skin_care_concern", "Meta: _wc_facebook_enhanced_catalog_attributes_skin_type", "Meta: _wc_facebook_enhanced_catalog_attributes_spf_value", "Meta: _wc_facebook_enhanced_catalog_attributes_standard_features", "Meta: _wc_facebook_enhanced_catalog_attributes_stop_use_indications", "Meta: _wc_facebook_enhanced_catalog_attributes_package_quantity", "Meta: _wc_facebook_enhanced_catalog_attributes_lens_material", "Meta: _wc_facebook_enhanced_catalog_attributes_lens_type", "Meta: _wc_facebook_enhanced_catalog_attributes_lens_tint", "Meta: _wc_facebook_enhanced_catalog_attributes_keywords", "Meta: _wc_facebook_enhanced_catalog_attributes_is_powered", "Meta: _wc_facebook_enhanced_catalog_attributes_instructions", "Meta: _wc_facebook_enhanced_catalog_attributes_ingredients_composition", "Meta: _wc_facebook_enhanced_catalog_attributes_inactive_ingredients", "Meta: _wc_facebook_enhanced_catalog_attributes_flavor", "Meta: _wc_facebook_enhanced_catalog_attributes_eyewear_rim", "Meta: _wc_facebook_enhanced_catalog_attributes_dosage", "Meta: _wc_facebook_enhanced_catalog_attributes_body_part", "Meta: _wc_facebook_enhanced_catalog_attributes_batteries_required", "Meta: _wc_facebook_enhanced_catalog_attributes_u_v_rating", "Meta: _wc_facebook_enhanced_catalog_attributes_capacity", "Meta: _wc_facebook_enhanced_catalog_attributes_health_concern", "Meta: _wc_facebook_enhanced_catalog_attributes_ingredients", "Attribute 2 name", "Attribute 2 value(s)", "Attribute 2 visible", "Attribute 2 global"}
	*result = append(*result, columns)
	exports := []map[string]string{}

	modelCates := []models.Categories{}
	catePaths := map[uint64]string{}
	cateIndices := map[uint64]int{}
	cateFlags := []bool{}

	err := repository.DB.Where("store_id = ?", storeID).Find(&modelCates).Error
	if err != nil {
		return err
	}

	for index, modelCate := range modelCates {
		catePaths[uint64(modelCate.ID)] = modelCate.Name
		cateIndices[uint64(modelCate.ID)] = index
		cateFlags = append(cateFlags, true)
	}

	for id := range catePaths {
		parentID := modelCates[cateIndices[id]].ParentID
		for parentID != 0 {
			catePaths[id] = catePaths[parentID] + " > " + catePaths[id]
			parentID = modelCates[cateIndices[parentID]].ParentID
			cateFlags[cateIndices[parentID]] = false
		}
	}

	modelTags := []models.Tags{}
	tagNames := map[uint64]string{}
	err = repository.DB.Where("store_id = ?", storeID).Find(&modelTags).Error
	if err != nil {
		return err
	}

	for _, modelTag := range modelTags {
		tagNames[uint64(modelTag.ID)] = modelTag.Name
	}

	modelProds := []models.Products{}
	prodIndices := map[uint64]int{}
	prodIDs := []uint64{}
	err = repository.DB.
		Table("store_products As prods").
		Select(`
			prods.*,
			(	Case ( Select Count( vars.id ) From store_product_variations As vars Where vars.product_id = prods.id And vars.deleted_at Is Null )
				When 1 Then	1 When 0 Then 0	Else 2 End ) AS type
		`).
		Where("prods.store_id = ? And prods.deleted_at Is Null", storeID).
		Scan(&modelProds).
		Error
	if err != nil {
		return err
	}

	for index, modelProd := range modelProds {
		prodIndices[uint64(modelProd.ID)] = index
		prodIDs = append(prodIDs, uint64(modelProd.ID))
	}

	modelProdCates := []models.ProductCategories{}
	prodcateIndices := map[uint64][]int{}
	err = repository.DB.Where("product_id In (?)", prodIDs).Find(&modelProdCates).Error
	if err != nil {
		return err
	}
	for index, modelProdCate := range modelProdCates {
		prodcateIndices[modelProdCate.ProductID] = append(prodcateIndices[modelProdCate.ProductID], index)
	}

	modelProdTags := []models.ProductTags{}
	prodtagIndices := map[uint64][]int{}
	err = repository.DB.Where("product_id In (?)", prodIDs).Find(&modelProdTags).Error
	if err != nil {
		return err
	}
	for index, modelProdTag := range modelProdTags {

		prodtagIndices[modelProdTag.ProductID] = append(prodtagIndices[modelProdTag.ProductID], index)
	}

	modelAttrs := []models.Attributes{}
	attrIndices := map[uint64][]int{}
	attrMaps := map[uint64]int{}
	attrIDs := []uint64{}
	err = repository.DB.Where("product_id In (?)", prodIDs).Find(&modelAttrs).Error
	if err != nil {
		return err
	}
	for index, modelAttr := range modelAttrs {
		attrMaps[uint64(modelAttr.ID)] = index
		attrIndices[modelAttr.ProductID] = append(attrIndices[modelAttr.ProductID], index)
		attrIDs = append(attrIDs, uint64(modelAttr.ID))
	}

	modelVals := []models.AttributeValues{}
	valIndices := map[uint64][]int{}
	valMaps := map[uint64]int{}
	err = repository.DB.Where("attribute_id In (?)", attrIDs).Find(&modelVals).Error
	if err != nil {
		return err
	}
	for index, modelVal := range modelVals {
		valMaps[uint64(modelVal.ID)] = index
		valIndices[uint64(modelVal.AttributeID)] = append(valIndices[uint64(modelVal.AttributeID)], index)
	}

	modelVars := []models.Variations{}
	varIDs := []uint64{}
	err = repository.DB.Where("product_id In (?)", prodIDs).Find(&modelVars).Error
	if err != nil {
		return err
	}
	for _, modelVar := range modelVars {
		varIDs = append(varIDs, uint64(modelVar.ID))
	}

	modelLinks := []models.Links{}
	upSells := map[uint64][]string{}
	crossSells := map[uint64][]string{}
	err = repository.DB.Where("product_id In (?)", prodIDs).Find(&modelLinks).Error
	if err != nil {
		return err
	}
	for _, modelLink := range modelLinks {
		if modelLink.IsUpCross == utils.UpSell {
			upSells[modelLink.ProductID] = append(upSells[modelLink.ProductID], modelProds[prodIndices[modelLink.LinkID]].Sku)
		} else {
			crossSells[modelLink.ProductID] = append(crossSells[modelLink.ProductID], modelProds[prodIndices[modelLink.LinkID]].Sku)
		}
	}

	modelShips := []models.ShippingData{}
	shipIndices := map[uint64]int{}
	err = repository.DB.Where("variation_id In (?)", varIDs).Find(&modelShips).Error
	if err != nil {
		return err
	}
	for index, modelShip := range modelShips {
		shipIndices[modelShip.VariationID] = index
	}

	modelDets := []models.VariationDetails{}
	detIndices := map[uint64][]uint64{}
	err = repository.DB.Where("variation_id In (?)", varIDs).Find(&modelDets).Error
	if err != nil {
		return err
	}
	for _, modelDet := range modelDets {
		detIndices[modelDet.VariationID] = append(detIndices[modelDet.VariationID], modelDet.AttributeValueID)
	}

	for _, modelProd := range modelProds {
		prodID := uint64(modelProd.ID)
		published := "0"
		if modelProd.Status == utils.Approved {
			published = "1"
		}
		images := []string{}
		json.Unmarshal([]byte(modelProd.ImageUrls), &images)
		if modelProd.Type == 2 {
			attr1, attr2, val1, val2 := "", "", "", ""
			if len(attrIndices[prodID]) > 0 {
				modelAttr := modelAttrs[attrIndices[prodID][0]]
				attr1 = modelAttr.AttributeName
				vals := []string{}
				for _, index := range valIndices[uint64(modelAttr.ID)] {
					vals = append(vals, modelVals[index].AttributeValue)
				}
				val1 = strings.Join(vals, ", ")
			} else if len(attrIndices[uint64(modelProd.ID)]) > 1 {
				modelAttr := modelAttrs[attrIndices[prodID][0]]
				attr2 = modelAttr.AttributeName
				vals := []string{}
				for _, index := range valIndices[uint64(modelAttr.ID)] {
					vals = append(vals, modelVals[index].AttributeValue)
				}
				val2 = strings.Join(vals, ", ")
			}
			cates := []string{}
			for _, index := range prodcateIndices[prodID] {
				cateID := modelProdCates[index].CategoryID
				if cateFlags[cateIndices[cateID]] {
					cates = append(cates, catePaths[cateID])
				}
			}
			tags := []string{}
			for _, index := range prodtagIndices[prodID] {
				tags = append(tags, tagNames[modelProdTags[index].TagID])
			}
			exports = append(exports, map[string]string{
				"ID":                      strconv.Itoa(len(exports) + 1),
				"Type":                    "variable",
				"SKU":                     modelProd.Sku,
				"Name":                    modelProd.Title,
				"Published":               published,
				"Is featured?":            "1",
				"Visibility in catalogue": "visible",
				"Short description":       modelProd.ShortDescription,
				"Description":             modelProd.LongDescription,
				"Tax status":              "taxable",
				"In stock?":               "1",
				"Backorders allowed?":     "0",
				"Sold individually?":      "0",
				"Allow customer reviews?": "1",
				"Categories":              strings.Join(cates, ", "),
				"Tags":                    strings.Join(tags, ", "),
				"Attribute 1 name":        attr1,
				"Attribute 1 value(s)":    val1,
				"Attribute 2 name":        attr2,
				"Attribute 2 value(s)":    val2,
				"Shipping class":          modelProd.ShippingClass,
				"Images":                  strings.Join(images, ", "),
				"Upsells":                 strings.Join(upSells[prodID], ", "),
				"Cross-sells":             strings.Join(crossSells[prodID], ", "),
			})
		}
	}

	for _, modelVar := range modelVars {
		varID := uint64(modelVar.ID)
		prodID := modelVar.ProductID
		modelProd := modelProds[prodIndices[prodID]]
		published := "0"
		if modelProd.Status == utils.Approved {
			published = "1"
		}
		images := []string{}
		json.Unmarshal([]byte(modelProd.ImageUrls), &images)
		cates := []string{}
		for _, index := range prodcateIndices[prodID] {
			cateID := modelProdCates[index].CategoryID
			if cateFlags[cateIndices[cateID]] {
				cates = append(cates, catePaths[cateID])
			}
		}
		tags := []string{}
		for _, index := range prodtagIndices[prodID] {
			tags = append(tags, tagNames[modelProdTags[index].TagID])
		}
		price := modelVar.Price
		if modelVar.DiscountType == utils.FixedAmountOff {
			price -= modelVar.DiscountAmount
		} else if modelVar.DiscountType == utils.PercentageOff {
			price -= modelVar.DiscountAmount * modelVar.Price / 100.0
		}
		salePrice := ""
		if modelVar.Price != price {
			salePrice = strconv.FormatFloat(price, 'f', 2, 64)
		}
		attr1, attr2, val1, val2 := "", "", "", ""
		if len(detIndices[varID]) > 0 {
			modelVal := modelVals[valMaps[detIndices[varID][0]]]
			val1 = modelVal.AttributeValue
			attr1 = modelAttrs[attrMaps[modelVal.AttributeID]].AttributeName
		} else if len(attrIndices[uint64(modelProd.ID)]) > 1 {
			modelVal := modelVals[valMaps[detIndices[varID][1]]]
			val1 = modelVal.AttributeValue
			attr1 = modelAttrs[attrMaps[modelVal.AttributeID]].AttributeName
		}
		if modelProd.Type == 1 {
			exports = append(exports, map[string]string{
				"ID":                      strconv.Itoa(len(exports) + 1),
				"Type":                    "simple",
				"SKU":                     modelProd.Sku,
				"Name":                    modelProd.Title,
				"Published":               published,
				"Is featured?":            "1",
				"Visibility in catalogue": "visible",
				"Short description":       modelProd.ShortDescription,
				"Description":             modelProd.LongDescription,
				"Tax status":              "taxable",
				"In stock?":               "1",
				"Backorders allowed?":     "0",
				"Sold individually?":      "0",
				"Allow customer reviews?": "1",
				"Weight (kg)":             strconv.FormatFloat(modelShips[shipIndices[varID]].Weight, 'f', 2, 64),
				"Length (cm)":             strconv.FormatFloat(modelShips[shipIndices[varID]].Length, 'f', 2, 64),
				"Width (cm)":              strconv.FormatFloat(modelShips[shipIndices[varID]].Width, 'f', 2, 64),
				"Height (cm)":             strconv.FormatFloat(modelShips[shipIndices[varID]].Height, 'f', 2, 64),
				"Categories":              strings.Join(cates, ", "),
				"Tags":                    strings.Join(tags, ", "),
				"Shipping class":          modelProd.ShippingClass,
				"Images":                  strings.Join(images, ", "),
				"Upsells":                 strings.Join(upSells[prodID], ", "),
				"Cross-sells":             strings.Join(crossSells[prodID], ", "),
				"Sale price":              salePrice,
				"Regular price":           strconv.FormatFloat(modelVar.Price, 'f', 2, 64),
			})
		} else if modelProd.Type == 2 {
			exports = append(exports, map[string]string{
				"ID":                      strconv.Itoa(len(exports) + 1),
				"Type":                    "variation",
				"SKU":                     modelVar.Sku,
				"Name":                    modelVar.Title,
				"Published":               published,
				"Is featured?":            "1",
				"Visibility in catalogue": "visible",
				"Short description":       "",
				"Description":             modelVar.Description,
				"Tax status":              "taxable",
				"In stock?":               "1",
				"Backorders allowed?":     "0",
				"Sold individually?":      "0",
				"Allow customer reviews?": "1",
				"Attribute 1 name":        attr1,
				"Attribute 1 value(s)":    val1,
				"Attribute 2 name":        attr2,
				"Attribute 2 value(s)":    val2,
				"Weight (kg)":             strconv.FormatFloat(modelShips[shipIndices[varID]].Weight, 'f', 2, 64),
				"Length (cm)":             strconv.FormatFloat(modelShips[shipIndices[varID]].Length, 'f', 2, 64),
				"Width (cm)":              strconv.FormatFloat(modelShips[shipIndices[varID]].Width, 'f', 2, 64),
				"Height (cm)":             strconv.FormatFloat(modelShips[shipIndices[varID]].Height, 'f', 2, 64),
				"Shipping class":          modelProd.ShippingClass,
				"Images":                  strings.Join(images, ", "),
				"Upsells":                 strings.Join(upSells[prodID], ", "),
				"Cross-sells":             strings.Join(crossSells[prodID], ", "),
				"Sale price":              salePrice,
				"Regular price":           strconv.FormatFloat(modelVar.Price, 'f', 2, 64),
			})
		}
	}

	for _, export := range exports {
		values := []string{}
		for _, column := range columns {
			values = append(values, export[column])
		}
		*result = append(*result, values)
	}

	return nil
}
