package model

// https://mholt.github.io/json-to-go/
type FreeStockChanged struct {
	Idhook           int    `json:"idhook"`
	Name             string `json:"name"`
	Event            string `json:"event"`
	EventTriggeredAt string `json:"event_triggered_at"`
	Data             struct {
		Idproduct                 int         `json:"idproduct"`
		Idvatgroup                int         `json:"idvatgroup"`
		Idsupplier                int         `json:"idsupplier"`
		Productcode               string      `json:"productcode"`
		Name                      string      `json:"name"`
		Price                     float64     `json:"price"`
		Fixedstockprice           float64     `json:"fixedstockprice"`
		ProductcodeSupplier       string      `json:"productcode_supplier"`
		Deliverytime              interface{} `json:"deliverytime"`
		Description               interface{} `json:"description"`
		Barcode                   string      `json:"barcode"`
		Unlimitedstock            bool        `json:"unlimitedstock"`
		Assembled                 bool        `json:"assembled"`
		Type                      string      `json:"type"`
		Weight                    int         `json:"weight"`
		Length                    interface{} `json:"length"`
		Width                     interface{} `json:"width"`
		Height                    interface{} `json:"height"`
		MinimumPurchaseQuantity   int         `json:"minimum_purchase_quantity"`
		PurchaseInQuantitiesOf    int         `json:"purchase_in_quantities_of"`
		HsCode                    interface{} `json:"hs_code"`
		CountryOfOrigin           interface{} `json:"country_of_origin"`
		Active                    bool        `json:"active"`
		CommentCount              int         `json:"comment_count"`
		AnalysisAbcClassification string      `json:"analysis_abc_classification"`
		AnalysisPickAmountPerDay  string      `json:"analysis_pick_amount_per_day"`
		Tags                      struct {
		} `json:"tags"`
		Productfields []interface{} `json:"productfields"`
		Images        []interface{} `json:"images"`
		Stock         []struct {
			Idwarehouse         int `json:"idwarehouse"`
			Stock               int `json:"stock"`
			Reserved            int `json:"reserved"`
			Reservedbackorders  int `json:"reservedbackorders"`
			Reservedpicklists   int `json:"reservedpicklists"`
			Reservedallocations int `json:"reservedallocations"`
			Freestock           int `json:"freestock"`
		} `json:"stock"`
		Pricelists []interface{} `json:"pricelists"`
	} `json:"data"`
}
