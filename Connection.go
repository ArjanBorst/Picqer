package picqer

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	Model "github.com/arjanborst/picqer/model"
)

type PicqerHttpConnection struct {
	HttpConn.httpconnection
}

func createNewRequest(url string) (*http.Response, error) {
	var req *http.Request
	var resp *http.Response
	var err error

	for retries := 0; retries < maxRetries; retries++ {

		req, err = http.NewRequest("GET", url, nil)
		if err == nil {
			req.Header.Set("Accept", "application/json")
			req.SetBasicAuth(username, password)
			client := &http.Client{}
			resp, err = client.Do(req)

			if err == nil && resp.StatusCode == 200 {
				return resp, nil
			}
		}

		time.Sleep(delay)
	}

	log.Println("Retries exhausted, unable to retrieve picklists: " + url + " res:  " + resp.Status)

	if err != nil {
		return nil, err
	}

	if resp != nil {
		return nil, errors.New(resp.Status)
	}

	return nil, errors.New("unknown error occurred")
}

func processRequest(resp *http.Response) ([]byte, error) {

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		println("Error while retrieving data from picqer")
	}

	return body, err
}

func (conn PicqerHttpConnection) GetPicklists() (Model.Picklists, error) {
	return conn.GetPicklistsByOffset(0)
}

func (conn PicqerHttpConnection) GetPicklistsByOffset(offset int) (Model.Picklists, error) {

	_url := url + "/api/v1/picklists"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	conn.HttpConnection.CreateNewRequest(_url)
	if conn.HttpConnection.Error() != nil {
		return nil, conn.HttpConnection.Error()
	}

	picklists := Model.Picklists{}
	json.Unmarshal(conn.Result(), &picklists)

	return picklists, conn.HttpConnection.Error()
}

func GetShipments(idpicklist int) (Model.Shipments, error) {
	resp, err := createNewRequest(url + "/api/v1/picklists/" + strconv.Itoa(idpicklist) + "/shipments")
	if err != nil {
		return nil, err
	}

	body, err := processRequest(resp)

	shipments := Model.Shipments{}
	json.Unmarshal(body, &shipments)

	return shipments, err
}

func GetPurchaseOrders() (Model.PurchaseOrders, error) {
	return GetPurchaseOrdersByOffset(0)
}

func GetPurchaseOrdersByOffset(offset int) (Model.PurchaseOrders, error) {

	_url := url + "/api/v1/purchaseorders"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	resp, err := createNewRequest(_url)
	if err != nil {
		return nil, err
	}

	body, err := processRequest(resp)
	if err != nil {
		return nil, err
	}

	purchaseOrders := Model.PurchaseOrders{}
	json.Unmarshal(body, &purchaseOrders)

	return purchaseOrders, nil
}

func GetPurchaseOrder(idpurchaseorder int) (Model.PurchaseOrder, error) {
	_url := url + "/api/v1/purchaseorders/" + strconv.Itoa(idpurchaseorder)

	resp, err := createNewRequest(_url)
	if err != nil {
		return Model.PurchaseOrder{}, err
	}

	body, err := processRequest(resp)
	if err != nil {
		return Model.PurchaseOrder{}, err
	}

	purchaseOrder := Model.PurchaseOrder{}
	json.Unmarshal(body, &purchaseOrder)

	return purchaseOrder, nil
}

func GetSuppliers() (Model.Suppliers, error) {
	return GetSuppliersByOffset(0)
}

func GetSuppliersByOffset(offset int) (Model.Suppliers, error) {

	_url := url + "/api/v1/suppliers"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	resp, err := createNewRequest(_url)
	if err != nil {
		return nil, err
	}

	body, err := processRequest(resp)
	if err != nil {
		return nil, err
	}

	suppliers := Model.Suppliers{}
	json.Unmarshal(body, &suppliers)

	return suppliers, nil
}

func GetProducts() (Model.Products, error) {
	return GetProductsByOffset(0)
}

func GetProductsByOffset(offset int) (Model.Products, error) {

	_url := url + "/api/v1/products"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset)
	}

	resp, err := createNewRequest(_url)
	if err != nil {
		return nil, err
	}

	body, err := processRequest(resp)
	if err != nil {
		return nil, err
	}

	products := Model.Products{}
	json.Unmarshal(body, &products)

	return products, nil
}

func GetPicqerOrders(offset int) ([]Model.PicqerOrder, error) {

	_url := url + "/api/v1/orders/"

	if offset > 0 {
		_url = _url + offsetUrl + strconv.Itoa(offset*100)
	}

	resp, err := createNewRequest(_url)
	if err != nil {
		return nil, err
	}

	body, err := processRequest(resp)
	if err != nil {
		return nil, err
	}

	var picqerOrder []Model.PicqerOrder
	json.Unmarshal(body, &picqerOrder)

	return picqerOrder, err
}
