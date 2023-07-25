package daemon

import (
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/arjanborst/core/helpers"
	"github.com/arjanborst/picqer"
	model "github.com/arjanborst/picqer/Model"
)

type PicqerDaemon struct {
	Pages          int
	FirstLoadPages int
	DataPath       string
	Orders         map[int]model.PicqerOrder
	Shipments      map[int]model.Shipment
	Debug          bool
	API            *picqer.PicqerHttpConnection
}

var mux = &sync.RWMutex{}

func (c PicqerDaemon) Start() {
	ticker := time.NewTicker(120 * time.Second)
	quit := make(chan struct{})

	mux.Lock()
	_ = helpers.LoadDataFromFile(c.DataPath, "orders.json", &c.Orders)
	_ = helpers.LoadDataFromFile(c.DataPath, "shipments.json", &c.Shipments)
	c.GetOrders(c.FirstLoadPages)
	c.GetShipments()
	mux.Unlock()

	helpers.SaveDataToFile(c.DataPath, "orders.json", &c.Orders)
	helpers.SaveDataToFile(c.DataPath, "shipments.json", &c.Shipments)

	for {
		select {
		case <-ticker.C:

			mux.Lock()
			c.GetOrders(c.Pages)
			c.GetShipments()
			mux.Unlock()

			helpers.SaveDataToFile(c.DataPath, "orders.json", &c.Orders)
			helpers.SaveDataToFile(c.DataPath, "shipments.json", &c.Shipments)

		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func (c PicqerDaemon) GetShipments() {
	if c.Debug {
		log.Println("Start Checking for new Shipments.")
	}

	for _, order := range c.Orders {

		for _, picklist := range order.Picklists {

			shipments, _ := c.API.GetShipments(picklist.Idpicklist)
			for _, shipment := range shipments {
				c.Shipments[shipment.Idshipment] = shipment
			}
		}
	}

	if c.Debug {
		log.Println("End Checking for new Shipments.")
	}
}

func (c PicqerDaemon) GetOrders(pages int) {
	if c.Debug {
		log.Println("Start Checking for new orders.")
	}

	for i := 0; i < pages; i++ {
		orders, _ := c.API.GetPicqerOrders(i)
		for _, order := range orders {
			c.Orders[order.Idorder] = order
		}
	}

	if c.Debug {
		log.Println("End Checking for new orders (" + strconv.Itoa(len(c.Orders)) + ").")
	}
}
