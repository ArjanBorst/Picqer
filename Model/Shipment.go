package Model

// https://mholt.github.io/json-to-go/
type Shipment struct {
	Idshipment                       int         `json:"idshipment"`
	Idpicklist                       int         `json:"idpicklist"`
	Idorder                          int         `json:"idorder"`
	Idreturn                         interface{} `json:"idreturn"`
	Idshippingprovider               int         `json:"idshippingprovider"`
	IdcompanyShippingprovider        int         `json:"idcompany_shippingprovider"`
	IdcompanyShippingproviderProfile int         `json:"idcompany_shippingprovider_profile"`
	Provider                         string      `json:"provider"`
	Providername                     string      `json:"providername"`
	PublicProvidername               string      `json:"public_providername"`
	ProfileName                      string      `json:"profile_name"`
	CarrierKey                       string      `json:"carrier_key"`
	Labelurl                         string      `json:"labelurl"`
	LabelurlPdf                      string      `json:"labelurl_pdf"`
	LabelurlZpl                      interface{} `json:"labelurl_zpl"`
	Trackingcode                     string      `json:"trackingcode"`
	Trackingurl                      string      `json:"trackingurl"`
	Tracktraceurl                    string      `json:"tracktraceurl"`
	CreatedByIduser                  int         `json:"created_by_iduser"`
	Cancelled                        bool        `json:"cancelled"`
	Created                          string      `json:"created"`
	Updated                          string      `json:"updated"`
}
