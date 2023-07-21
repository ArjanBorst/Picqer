package picqer

import (
	"encoding/json"
	"strconv"

	"github.com/arjanborst/core/httpconn"
	model "github.com/arjanborst/picqer/Model"
)

type PicqerHttpConnection struct {
	httpconn.HttpConnection
}

func (c PicqerHttpConnection) GetPicklists() (model.Picklists, error) {
	return c.GetPicklistsByOffset(0)
}

func (c PicqerHttpConnection) GetPicklistsByOffset(offset int) (model.Picklists, error) {

	_url := url + "/api/v1/picklists"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	c.HttpConnection.CreateNewRequest(_url)
	if c.HttpConnection.Error() != nil {
		return nil, c.HttpConnection.Error()
	}

	picklists := model.Picklists{}
	json.Unmarshal(c.HttpConnection.Result(), &picklists)

	return picklists, c.HttpConnection.Error()
}

func (c PicqerHttpConnection) GetShipments(idpicklist int) (model.Shipments, error) {

	_url := url + "/api/v1/picklists/" + strconv.Itoa(idpicklist) + "/shipments"

	c.HttpConnection.CreateNewRequest(_url)
	if c.HttpConnection.Error() != nil {
		return nil, c.HttpConnection.Error()
	}

	shipments := model.Shipments{}
	json.Unmarshal(c.HttpConnection.Result(), &shipments)

	return shipments, c.HttpConnection.Error()
}

func (c PicqerHttpConnection) GetPurchaseOrders() (model.PurchaseOrders, error) {
	return c.GetPurchaseOrdersByOffset(0)
}

func (c PicqerHttpConnection) GetPurchaseOrdersByOffset(offset int) (model.PurchaseOrders, error) {

	_url := url + "/api/v1/purchaseorders"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	c.HttpConnection.CreateNewRequest(_url)
	if c.HttpConnection.Error() != nil {
		return nil, c.HttpConnection.Error()
	}

	purchaseOrders := model.PurchaseOrders{}
	json.Unmarshal(c.HttpConnection.Result(), &purchaseOrders)

	return purchaseOrders, nil
}

func (c PicqerHttpConnection) GetPurchaseOrder(idpurchaseorder int) (model.PurchaseOrder, error) {
	_url := url + "/api/v1/purchaseorders/" + strconv.Itoa(idpurchaseorder)

	c.HttpConnection.CreateNewRequest(_url)
	if c.HttpConnection.Error() != nil {
		return model.PurchaseOrder{}, c.HttpConnection.Error()
	}

	purchaseOrder := model.PurchaseOrder{}
	json.Unmarshal(c.HttpConnection.Result(), &purchaseOrder)

	return purchaseOrder, nil
}

func (c PicqerHttpConnection) GetSuppliers() (model.Suppliers, error) {
	return c.GetSuppliersByOffset(0)
}

func (c PicqerHttpConnection) GetSuppliersByOffset(offset int) (model.Suppliers, error) {

	_url := url + "/api/v1/suppliers"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	c.HttpConnection.CreateNewRequest(_url)
	if c.HttpConnection.Error() != nil {
		return nil, c.HttpConnection.Error()
	}

	suppliers := model.Suppliers{}
	json.Unmarshal(c.HttpConnection.Result(), &suppliers)

	return suppliers, nil
}

func (c PicqerHttpConnection) GetProducts() (model.Products, error) {
	return c.GetProductsByOffset(0)
}

func (c PicqerHttpConnection) GetProductsByOffset(offset int) (model.Products, error) {

	_url := url + "/api/v1/products"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	if err := c.HttpConnection.CreateNewRequest(_url); err != nil {
		return nil, err
	}

	products := model.Products{}
	json.Unmarshal(c.HttpConnection.Result(), &products)

	return products, nil
}

func (c PicqerHttpConnection) GetPicqerOrders(offset ...int) ([]model.PicqerOrder, error) {

	_url := url + "/api/v1/orders/"

	if len(offset) > 0 && offset[0] > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset[0]*100)
	}

	if err := c.HttpConnection.CreateNewRequest(_url); err != nil {
		return nil, err
	}

	var picqerOrder []model.PicqerOrder
	if err := json.Unmarshal(c.HttpConnection.Result(), &picqerOrder); err != nil {
		return nil, err
	}

	return picqerOrder, nil
}
