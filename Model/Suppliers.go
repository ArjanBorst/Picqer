package Model

// https://mholt.github.io/json-to-go/
type Suppliers []struct {
	Idsupplier   int         `json:"idsupplier"`
	Name         string      `json:"name"`
	Contactname  interface{} `json:"contactname"`
	Address      string      `json:"address"`
	Address2     interface{} `json:"address2"`
	Zipcode      string      `json:"zipcode"`
	City         string      `json:"city"`
	Region       interface{} `json:"region"`
	Country      string      `json:"country"`
	Telephone    interface{} `json:"telephone"`
	Emailaddress interface{} `json:"emailaddress"`
	Remarks      interface{} `json:"remarks"`
	Language     string      `json:"language"`
}