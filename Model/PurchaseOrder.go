package model

// https://mholt.github.io/json-to-go/
type PurchaseOrder struct {
	Idpurchaseorder int         `json:"idpurchaseorder"`
	Idsupplier      int         `json:"idsupplier"`
	Idwarehouse     int         `json:"idwarehouse"`
	Purchaseorderid string      `json:"purchaseorderid"`
	SupplierName    interface{} `json:"supplier_name"`
	SupplierOrderid interface{} `json:"supplier_orderid"`
	Status          string      `json:"status"`
	Remarks         string      `json:"remarks"`
	DeliveryDate    string      `json:"delivery_date"`
	Created         string      `json:"created"`
	Updated         string      `json:"updated"`
	Products        []struct {
		Idproduct           int     `json:"idproduct"`
		Idvatgroup          int     `json:"idvatgroup"`
		Productcode         string  `json:"productcode"`
		ProductcodeSupplier string  `json:"productcode_supplier"`
		Name                string  `json:"name"`
		Price               float64 `json:"price"`
		Amount              int     `json:"amount"`
		Amountreceived      int     `json:"amountreceived"`
		Weight              int     `json:"weight"`
	} `json:"products"`
}
