package picqer

import (
	"encoding/json"
	"strconv"

	"github.com/arjanborst/core/httpconn"
	Model "github.com/arjanborst/picqer/model"
)

type PicqerHttpConnection struct {
	httpconn.HttpConnection
}

func (c PicqerHttpConnection) GetPicklists() (Model.Picklists, error) {
	return c.GetPicklistsByOffset(0)
}

func (c PicqerHttpConnection) GetPicklistsByOffset(offset int) (Model.Picklists, error) {

	_url := url + "/api/v1/picklists"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	c.HttpConnection.CreateNewRequest(_url)
	if c.HttpConnection.Error() != nil {
		return nil, c.HttpConnection.Error()
	}

	picklists := Model.Picklists{}
	json.Unmarshal(c.HttpConnection.Result(), &picklists)

	return picklists, c.HttpConnection.Error()
}

func (c PicqerHttpConnection) GetShipments(idpicklist int) (Model.Shipments, error) {

	_url := url + "/api/v1/picklists/" + strconv.Itoa(idpicklist) + "/shipments"

	c.HttpConnection.CreateNewRequest(_url)
	if c.HttpConnection.Error() != nil {
		return nil, c.HttpConnection.Error()
	}

	shipments := Model.Shipments{}
	json.Unmarshal(c.HttpConnection.Result(), &shipments)

	return shipments, c.HttpConnection.Error()
}

func (c PicqerHttpConnection) GetPurchaseOrders() (Model.PurchaseOrders, error) {
	return c.GetPurchaseOrdersByOffset(0)
}

func (c PicqerHttpConnection) GetPurchaseOrdersByOffset(offset int) (Model.PurchaseOrders, error) {

	_url := url + "/api/v1/purchaseorders"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	c.HttpConnection.CreateNewRequest(_url)
	if c.HttpConnection.Error() != nil {
		return nil, c.HttpConnection.Error()
	}

	purchaseOrders := Model.PurchaseOrders{}
	json.Unmarshal(c.HttpConnection.Result(), &purchaseOrders)

	return purchaseOrders, nil
}

func (c PicqerHttpConnection) GetPurchaseOrder(idpurchaseorder int) (Model.PurchaseOrder, error) {
	_url := url + "/api/v1/purchaseorders/" + strconv.Itoa(idpurchaseorder)

	c.HttpConnection.CreateNewRequest(_url)
	if c.HttpConnection.Error() != nil {
		return Model.PurchaseOrder{}, c.HttpConnection.Error()
	}

	purchaseOrder := Model.PurchaseOrder{}
	json.Unmarshal(c.HttpConnection.Result(), &purchaseOrder)

	return purchaseOrder, nil
}

func (c PicqerHttpConnection) GetSuppliers() (Model.Suppliers, error) {
	return c.GetSuppliersByOffset(0)
}

func (c PicqerHttpConnection) GetSuppliersByOffset(offset int) (Model.Suppliers, error) {

	_url := url + "/api/v1/suppliers"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	c.HttpConnection.CreateNewRequest(_url)
	if c.HttpConnection.Error() != nil {
		return nil, c.HttpConnection.Error()
	}

	suppliers := Model.Suppliers{}
	json.Unmarshal(c.HttpConnection.Result(), &suppliers)

	return suppliers, nil
}

func (c PicqerHttpConnection) GetProducts() (Model.Products, error) {
	return c.GetProductsByOffset(0)
}

func (c PicqerHttpConnection) GetProductsByOffset(offset int) (Model.Products, error) {

	_url := url + "/api/v1/products"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	if err := c.HttpConnection.CreateNewRequest(_url); err != nil {
		return nil, err
	}

	products := Model.Products{}
	json.Unmarshal(c.HttpConnection.Result(), &products)

	return products, nil
}

func (c PicqerHttpConnection) GetPicqerOrders(offset ...int) ([]Model.PicqerOrder, error) {

	_url := url + "/api/v1/orders/"

	if len(offset) > 0 && offset[0] > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset[0]*100)
	}

	if err := c.HttpConnection.CreateNewRequest(_url); err != nil {
		return nil, err
	}

	var picqerOrder []Model.PicqerOrder
	if err := json.Unmarshal(c.HttpConnection.Result(), &picqerOrder); err != nil {
		return nil, err
	}

	return picqerOrder, nil
}
