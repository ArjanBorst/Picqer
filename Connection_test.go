package picqer

import (
	"testing"

	"github.com/arjanborst/core/httpconn"
)

func TestGetPicklist(t *testing.T) {

	picqer := &PicqerHttpConnection{
		HttpConnection: &httpconn.HttpConn{
			Username:   username,
			Password:   password,
			Hostname:   url,
			MaxRetries: maxRetries,
			Delay:      delay,
			Debug:      true,
		},
	}

	Picklists, err := picqer.GetPicklists()

	if err != nil {
		t.Fatal(err)
	}

	if len(Picklists) > 100 {
		t.Fatal("List is greather than 100 should be 100")
	}

	if len(Picklists) < 100 {
		t.Fatal("List is less than 100 should be 100")
	}

}

func TestGetPicklistWithOffset(t *testing.T) {

	picqer := &PicqerHttpConnection{
		HttpConnection: &httpconn.HttpConn{
			Username:   username,
			Password:   password,
			Hostname:   url,
			MaxRetries: maxRetries,
			Delay:      delay,
			Debug:      true,
		},
	}

	Picklists, err := picqer.GetPicklistsByOffset(1)

	if err != nil {
		t.Fatal(err)
	}

	if len(Picklists) > 100 {
		t.Fatal("List is greather than 100 should be 100")
	}

	if len(Picklists) < 100 {
		t.Fatal("List is less than 100 should be 100")
	}
}

func TestGetShipments(t *testing.T) {

	picqer := &PicqerHttpConnection{
		HttpConnection: &httpconn.HttpConn{
			Username:   username,
			Password:   password,
			Hostname:   url,
			MaxRetries: maxRetries,
			Delay:      delay,
			Debug:      true,
		},
	}

	resNoShipmnents, err1 := picqer.GetShipments(91632583)
	resOneShipmnents, err2 := picqer.GetShipments(91522445)

	if err1 != nil {
		t.Fatal(err1)
	}

	if err2 != nil {
		t.Fatal(err2)
	}

	if len(resNoShipmnents) != 0 {
		t.Fatal("Length of array should be 0")
	}

	if len(resOneShipmnents) != 1 {
		t.Fatal("Length of array should be 1")
	}

	if resOneShipmnents[0].Idorder != 131502569 {
		t.Fatal("Idorder should be 131502569")
	}
}

func TestGetPurchaseOrdersByOffset(t *testing.T) {

	picqer := &PicqerHttpConnection{
		HttpConnection: &httpconn.HttpConn{
			Username:   username,
			Password:   password,
			Hostname:   url,
			MaxRetries: maxRetries,
			Delay:      delay,
			Debug:      true,
		},
	}

	Picklists, err := picqer.GetPurchaseOrdersByOffset(1)

	if err != nil {
		t.Fatal(err)
	}

	if len(Picklists) > 100 {
		t.Fatal("List is greather than 100 should be 100")
	}

	if len(Picklists) < 100 {
		t.Fatal("List is less than 100 should be 100")
	}
}

func TestGetPicqerOrders(t *testing.T) {

	picqer := &PicqerHttpConnection{
		HttpConnection: &httpconn.HttpConn{
			Username:   username,
			Password:   password,
			Hostname:   url,
			MaxRetries: maxRetries,
			Delay:      delay,
			Debug:      true,
		},
	}

	res, err := picqer.GetPicqerOrders()

	if err != nil {
		t.Fatal(err)
	}

	if len(res) > 100 {
		t.Fatal("List is greather than 100 should be 100")
	}

	if len(res) < 100 {
		t.Fatal("List is less than 100 should be 100")
	}
}
