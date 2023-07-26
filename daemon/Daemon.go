package daemon

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/arjanborst/core/helpers"
	"github.com/arjanborst/picqer"
	model "github.com/arjanborst/picqer/Model"
)

type shipmentTracker struct {
	LastChecked time.Time `json:"lastchecked"`
	LastUpdated time.Time `json:"lastupdated"`
	Created     time.Time `json:"created"`
	IdOrder     int       `json:"idorder"`
	IdPicklist  int       `json:"idpicklist"`
	IdShipment  int       `json:"idshipment"`
}

type PicqerDaemon struct {
	Pages           int
	FirstLoadPages  int
	DataPath        string
	Orders          map[int]model.PicqerOrder
	Shipments       map[int]model.Shipment
	Debug           bool
	API             *picqer.PicqerHttpConnection
	ShipmentTracker map[int]shipmentTracker
}

var mux = &sync.RWMutex{}

func (c PicqerDaemon) Start() {
	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})

	c.ShipmentTracker = make(map[int]shipmentTracker)

	mux.Lock()
	_ = helpers.LoadDataFromFile(c.DataPath, "orders.json", &c.Orders)
	_ = helpers.LoadDataFromFile(c.DataPath, "shipments.json", &c.Shipments)
	_ = helpers.LoadDataFromFile(c.DataPath, "shipmenttracker.json", &c.ShipmentTracker)
	c.GetOrders(c.FirstLoadPages)
	c.GetShipments()
	mux.Unlock()

	helpers.SaveDataToFile(c.DataPath, "orders.json", &c.Orders)
	helpers.SaveDataToFile(c.DataPath, "shipments.json", &c.Shipments)
	helpers.SaveDataToFile(c.DataPath, "shipmenttracker.json", &c.ShipmentTracker)

	for {
		select {
		case <-ticker.C:

			mux.Lock()
			c.GetOrders(c.Pages)
			c.GetShipments()
			mux.Unlock()

			helpers.SaveDataToFile(c.DataPath, "orders.json", &c.Orders)
			helpers.SaveDataToFile(c.DataPath, "shipments.json", &c.Shipments)
			helpers.SaveDataToFile(c.DataPath, "shipmenttracker.json", &c.ShipmentTracker)

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

			if t, trackerExist := c.ShipmentTracker[picklist.Idpicklist]; trackerExist {

				currentTime := time.Now()
				duration := currentTime.Sub(t.LastUpdated)

				if c.Debug {
					fmt.Println(duration.Minutes())
				}

				if t.IdShipment == 0 && duration.Minutes() >= 4 {

					shipments, _ := c.API.GetShipments(picklist.Idpicklist)
					for _, shipment := range shipments {

						c.Shipments[shipment.Idshipment] = shipment

						c.ShipmentTracker[shipment.Idpicklist] = shipmentTracker{
							LastChecked: time.Now(),
							LastUpdated: time.Now(),
							IdShipment:  shipment.Idshipment,
						}
					}
				}

			} else {

				if c.Debug {
					log.Println("Add entry to ShipmentTracker.")
				}

				c.ShipmentTracker[picklist.Idpicklist] = shipmentTracker{
					LastChecked: time.Now(),
					LastUpdated: time.Now(),
					Created:     time.Now(),
					IdOrder:     picklist.Idorder,
					IdPicklist:  picklist.Idpicklist,
					IdShipment:  0,
				}
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
