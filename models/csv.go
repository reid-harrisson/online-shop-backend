package models

type CSVs struct {
	ID                    string
	Type                  string `json:"Type"`
	Sku                   string `json:"SKU"`
	Name                  string `json:"Name"`
	Published             string `json:"Published"`
	IsFeatured            string `json:"Is featured?"`
	VisibilityInCatalogue string `json:"Visibility in catalogue"`
	ShortDescription      string `json:"Short description"`
	Description           string `json:"Description"`
	DateSalePriceStarts   string `json:"Date sale price starts"`
	DateSalePriceEnds     string `json:"Date sale price ends"`
	TaxStatus             string `json:"Tax status"`
	TaxClass              string `json:"Tax class"`
	InStock               string `json:"In stock?"`
	Stock                 string `json:"Stock"`
	LowStockAmount        string `json:"Low stock amount"`
	BackordersAllowed     string `json:"Backorders allowed?"`
	SoldIndividually      string `json:"Sold individually?"`
	Weight                string `json:"Weight (kg)"`
	Length                string `json:"Length (cm)"`
	Width                 string `json:"Width (cm)"`
	Height                string `json:"Height (cm)"`
	AllowCustomerReviews  string `json:"Allow customer reviews?"`
	PurchaseNote          string `json:"Purchase note"`
	SalePrice             string `json:"Sale price"`
	RegularPrice          string `json:"Regular price"`
	Categories            string `json:"Categories"`
	Tags                  string `json:"Tags"`
	ShippingClass         string `json:"Shipping class"`
	Images                string `json:"Images"`
	DownloadLimit         string `json:"Download limit"`
	DownloadExpiryDays    string `json:"Download expiry days"`
	Parent                string `json:"Parent"`
	GroupedProducts       string `json:"Grouped products"`
	Upsells               string `json:"Upsells"`
	CrossSells            string `json:"Cross-sells"`
	ExternalUrl           string `json:"External URL"`
	ButtonText            string `json:"Button text"`
	Position              string `json:"Position"`
	Attribute1Name        string `json:"Attribute 1 name"`
	Attribute1Values      string `json:"Attribute 1 value(s)"`
	Attribute1Visible     string `json:"Attribute 1 visible"`
	Attribute1Global      string `json:"Attribute 1 global"`
	Attribute2Name        string `json:"Attribute 2 name"`
	Attribute2Values      string `json:"Attribute 2 value(s)"`
	Attribute2Visible     string `json:"Attribute 2 visible"`
	Attribute2Global      string `json:"Attribute 2 global"`
}
