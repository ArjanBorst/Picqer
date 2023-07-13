package Model

// https://mholt.github.io/json-to-go/
type Picklists []struct {
	Idpicklist                int         `json:"idpicklist"`
	Picklistid                string      `json:"picklistid"`
	Idcustomer                int         `json:"idcustomer"`
	Idorder                   int         `json:"idorder"`
	Idreturn                  interface{} `json:"idreturn"`
	Idwarehouse               int         `json:"idwarehouse"`
	Idtemplate                int         `json:"idtemplate"`
	IdshippingproviderProfile int         `json:"idshippingprovider_profile"`
	Deliveryname              string      `json:"deliveryname"`
	Deliverycontact           interface{} `json:"deliverycontact"`
	Deliveryaddress           string      `json:"deliveryaddress"`
	Deliveryaddress2          interface{} `json:"deliveryaddress2"`
	Deliveryzipcode           string      `json:"deliveryzipcode"`
	Deliverycity              string      `json:"deliverycity"`
	Deliveryregion            interface{} `json:"deliveryregion"`
	Deliverycountry           string      `json:"deliverycountry"`
	Telephone                 string      `json:"telephone"`
	Emailaddress              string      `json:"emailaddress"`
	Reference                 string      `json:"reference"`
	AssignedToIduser          interface{} `json:"assigned_to_iduser"`
	Invoiced                  bool        `json:"invoiced"`
	Urgent                    bool        `json:"urgent"`
	PreferredDeliveryDate     interface{} `json:"preferred_delivery_date"`
	Status                    string      `json:"status"`
	Totalproducts             int         `json:"totalproducts"`
	Totalpicked               int         `json:"totalpicked"`
	SnoozedUntil              interface{} `json:"snoozed_until"`
	ClosedByIduser            interface{} `json:"closed_by_iduser"`
	ClosedAt                  interface{} `json:"closed_at"`
	Created                   string      `json:"created"`
	Updated                   string      `json:"updated"`
	Products                  []struct {
		IdpicklistProduct          int         `json:"idpicklist_product"`
		Idproduct                  int         `json:"idproduct"`
		IdorderProduct             int         `json:"idorder_product"`
		IdreturnProductReplacement interface{} `json:"idreturn_product_replacement"`
		Idvatgroup                 int         `json:"idvatgroup"`
		Productcode                string      `json:"productcode"`
		Name                       string      `json:"name"`
		Remarks                    interface{} `json:"remarks"`
		Amount                     int         `json:"amount"`
		Amountpicked               int         `json:"amountpicked"`
		AmountPicked               int         `json:"amount_picked"`
		Price                      float64     `json:"price"`
		Weight                     int         `json:"weight"`
		Stocklocation              interface{} `json:"stocklocation"`
		StockLocation              interface{} `json:"stock_location"`
		PartofIdpicklistProduct    interface{} `json:"partof_idpicklist_product"`
		HasParts                   bool        `json:"has_parts"`
	} `json:"products"`
	CommentCount int `json:"comment_count"`
}