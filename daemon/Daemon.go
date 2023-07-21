package daemon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

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
	_ = c.loadDataFromFile()
	c.GetOrders(c.FirstLoadPages)
	mux.Unlock()

	for {
		select {
		case <-ticker.C:

			mux.Lock()
			c.GetOrders(c.Pages)
			c.GetShipments()
			mux.Unlock()

			c.saveDataToFile()

		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func (c PicqerDaemon) GetShipments() {

	for _, order := range c.Orders {

		for _, picklist := range order.Picklists {

			shipments, _ := c.API.GetShipments(picklist.Idpicklist)
			for _, shipment := range shipments {
				c.Shipments[shipment.Idshipment] = shipment
			}
		}
	}
}

func (c PicqerDaemon) GetOrders(pages int) {
	if c.Debug {
		fmt.Println("Start Checking for new orders.")
	}

	for i := 0; i < pages; i++ {
		orders, _ := c.API.GetPicqerOrders(i)
		for _, order := range orders {
			c.Orders[order.Idorder] = order
		}
	}

	if c.Debug {
		fmt.Println("End Checking for new orders (" + strconv.Itoa(len(c.Orders)) + ").")
	}
}

func (c PicqerDaemon) loadDataFromFile() error {
	if c.Debug {
		log.Println("Start reading orders from file")
	}

	filename := c.DataPath + "orders.json"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if c.Debug {
			log.Println("No existing file. Creating a new one.")
		}

		file, err := os.Create(filename)
		if err != nil {
			return errors.New("ERROR CREATING NEW FILE")
		}
		defer file.Close()

		// Exit function because nothing more to load.
		return nil
	} else if err != nil {
		return errors.New("Error accessing the file")
	}

	file, err := os.Open(filename)
	if err != nil {
		return errors.New("Error loading Data from file")
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.New("Error reading Data from file")
	}

	err = json.Unmarshal(data, &c.Orders)
	if err != nil {
		return errors.New("Error unmarshalling JSON")
	}

	if c.Debug {
		log.Println("Loaded a total of " + strconv.Itoa(len(c.Orders)) + " orders")
	}

	return nil
}

func (c PicqerDaemon) saveDataToFile() {
	jsonData, err := json.Marshal(c.Orders)
	if err != nil {
		println("Error while saving picqer orders to file. Problem while convert object to JSON format")
		panic(err)
	}

	jsonFile, err := os.Create(c.DataPath + "orders.json")
	if err != nil {
		println("Error while saving picqer orders to file. Error while creating file or truncating file")
		panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)

	fmt.Println("Saved orders to file: " + c.DataPath + "orders.json")
}
