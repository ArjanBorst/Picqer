package Model

// https://mholt.github.io/json-to-go/
type Products []struct {
	Idproduct               int         `json:"idproduct"`
	Idvatgroup              int         `json:"idvatgroup"`
	Idsupplier              interface{} `json:"idsupplier"`
	Productcode             string      `json:"productcode"`
	Name                    string      `json:"name"`
	Price                   float64     `json:"price"`
	Fixedstockprice         float64     `json:"fixedstockprice"`
	ProductcodeSupplier     string      `json:"productcode_supplier"`
	Deliverytime            interface{} `json:"deliverytime"`
	Description             interface{} `json:"description"`
	Barcode                 interface{} `json:"barcode"`
	Unlimitedstock          bool        `json:"unlimitedstock"`
	Assembled               bool        `json:"assembled"`
	Type                    string      `json:"type"`
	Weight                  int         `json:"weight"`
	Length                  int         `json:"length"`
	Width                   int         `json:"width"`
	Height                  int         `json:"height"`
	MinimumPurchaseQuantity int         `json:"minimum_purchase_quantity"`
	PurchaseInQuantitiesOf  int         `json:"purchase_in_quantities_of"`
	HsCode                  interface{} `json:"hs_code"`
	CountryOfOrigin         interface{} `json:"country_of_origin"`
	Active                  bool        `json:"active"`
	Tags                    struct {
		TopWebshop struct {
			Idtag     int    `json:"idtag"`
			Title     string `json:"title"`
			Color     string `json:"color"`
			Inherit   bool   `json:"inherit"`
			TextColor string `json:"textColor"`
		} `json:"TopWebshop"`
		SummerProducts struct {
			Idtag     int    `json:"idtag"`
			Title     string `json:"title"`
			Color     string `json:"color"`
			Inherit   bool   `json:"inherit"`
			TextColor string `json:"textColor"`
		} `json:"SummerProducts"`
	} `json:"tags,omitempty"`
	Productfields []struct {
		Idproductfield int    `json:"idproductfield"`
		Title          string `json:"title"`
		Value          string `json:"value"`
	} `json:"productfields"`
	Images []string `json:"images"`
	Stock  []struct {
		Idwarehouse         int `json:"idwarehouse"`
		Stock               int `json:"stock"`
		Reserved            int `json:"reserved"`
		Reservedbackorders  int `json:"reservedbackorders"`
		Reservedpicklists   int `json:"reservedpicklists"`
		Reservedallocations int `json:"reservedallocations"`
		Freestock           int `json:"freestock"`
	} `json:"stock"`
	Pricelists []interface{} `json:"pricelists"`
	Tags0      struct {
	} `json:"tags,omitempty"`
}