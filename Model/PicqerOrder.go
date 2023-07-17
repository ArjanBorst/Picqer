package model

// https://mholt.github.io/json-to-go/

type PicqerOrder struct {
	Idorder                   int    `json:"idorder"`
	Idcustomer                int    `json:"idcustomer"`
	Idtemplate                int    `json:"idtemplate,omitempty"`
	IdshippingproviderProfile any    `json:"idshippingprovider_profile,omitempty"`
	Orderid                   string `json:"orderid"`
	Deliveryname              string `json:"deliveryname"`
	Deliverycontactname       string `json:"deliverycontactname"`
	Deliveryaddress           string `json:"deliveryaddress"`
	Deliveryaddress2          any    `json:"deliveryaddress2"`
	Deliveryzipcode           string `json:"deliveryzipcode"`
	Deliverycity              string `json:"deliverycity"`
	Deliveryregion            any    `json:"deliveryregion"`
	Deliverycountry           string `json:"deliverycountry"`
	FullDeliveryAddress       string `json:"full_delivery_address"`
	Invoicename               string `json:"invoicename"`
	Invoicecontactname        string `json:"invoicecontactname"`
	Invoiceaddress            string `json:"invoiceaddress"`
	Invoiceaddress2           any    `json:"invoiceaddress2"`
	Invoicezipcode            string `json:"invoicezipcode"`
	Invoicecity               string `json:"invoicecity"`
	Invoiceregion             any    `json:"invoiceregion"`
	Invoicecountry            string `json:"invoicecountry"`
	FullInvoiceAddress        string `json:"full_invoice_address"`
	Telephone                 any    `json:"telephone"`
	Emailaddress              string `json:"emailaddress"`
	Reference                 string `json:"reference"`
	CustomerRemarks           any    `json:"customer_remarks"`
	PickupPointData           any    `json:"pickup_point_data"`
	Partialdelivery           bool   `json:"partialdelivery"`
	AutoSplit                 bool   `json:"auto_split"`
	Invoiced                  bool   `json:"invoiced"`
	PreferredDeliveryDate     any    `json:"preferred_delivery_date"`
	Discount                  int    `json:"discount"`
	Calculatevat              bool   `json:"calculatevat"`
	Language                  string `json:"language"`
	Status                    string `json:"status"`
	PublicStatusPage          string `json:"public_status_page"`
	Created                   string `json:"created"`
	Updated                   string `json:"updated"`
	Warehouses                []int  `json:"warehouses"`
	Tags                      []struct {
		Idtag     int    `json:"idtag"`
		Title     string `json:"title"`
		Color     string `json:"color"`
		Inherit   bool   `json:"inherit"`
		TextColor string `json:"textColor"`
	} `json:"Tags"`
	Orderfields []Orderfield `json:"Orderfields"`
	Products    []struct {
		IdorderProduct       int     `json:"idorder_product"`
		Idproduct            int     `json:"idproduct"`
		Idvatgroup           int     `json:"idvatgroup"`
		Productcode          string  `json:"productcode"`
		Name                 string  `json:"name"`
		Remarks              string  `json:"remarks"`
		Price                float64 `json:"price"`
		Amount               int     `json:"amount"`
		AmountCancelled      int     `json:"amount_cancelled"`
		Weight               int     `json:"weight"`
		PartofIdorderProduct any     `json:"partof_idorder_product"`
		HasParts             bool    `json:"has_parts"`
	} `json:"products"`
	Pricelists []any `json:"pricelists"`
	Picklists  []struct {
		AssignedToIduser          int    `json:"assigned_to_iduser"`
		ClosedAt                  string `json:"closed_at"`
		ClosedByIduser            int    `json:"closed_by_iduser"`
		CommentCount              int    `json:"comment_count"`
		Created                   string `json:"created"`
		Deliveryaddress           string `json:"deliveryaddress"`
		Deliveryaddress2          any    `json:"deliveryaddress2"`
		Deliverycity              string `json:"deliverycity"`
		Deliverycontact           any    `json:"deliverycontact"`
		Deliverycountry           string `json:"deliverycountry"`
		Deliveryname              string `json:"deliveryname"`
		Deliveryregion            any    `json:"deliveryregion"`
		Deliveryzipcode           string `json:"deliveryzipcode"`
		Emailaddress              string `json:"emailaddress"`
		Idcustomer                int    `json:"idcustomer"`
		Idorder                   int    `json:"idorder"`
		Idpicklist                int    `json:"idpicklist"`
		Idreturn                  any    `json:"idreturn"`
		IdshippingproviderProfile int    `json:"idshippingprovider_profile"`
		Idtemplate                int    `json:"idtemplate"`
		Idwarehouse               int    `json:"idwarehouse"`
		Invoiced                  bool   `json:"invoiced"`
		Picklistid                string `json:"picklistid"`
		PreferredDeliveryDate     any    `json:"preferred_delivery_date"`
		Reference                 string `json:"reference"`
		SnoozedUntil              any    `json:"snoozed_until"`
		Status                    string `json:"status"`
		Telephone                 string `json:"telephone"`
		Totalpicked               int    `json:"totalpicked"`
		Totalproducts             int    `json:"totalproducts"`
		Updated                   string `json:"updated"`
		Urgent                    bool   `json:"urgent"`
	} `json:"picklists"`
}

type Orderfield struct {
	Idorderfield int    `json:"Idorderfield"`
	Title        string `json:"Title"`
	Value        string `json:"Value"`
}
