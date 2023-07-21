package daemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	model "github.com/arjanborst/picqer/Model"
)

type PicqerDaemon struct {
	Pages          int
	FirstLoadPages int
	DataPath       string
	Orders         map[string]model.PicqerOrder
}

var mux = &sync.RWMutex{}

func (c PicqerDaemon) Start() {
	ticker := time.NewTicker(120 * time.Second)
	quit := make(chan struct{})
	pages := c.FirstLoadPages

	mux.Lock()
	c.loadDataFromFile()
	mux.Unlock()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Update with newest data from Picqer")

			mux.Lock()
			//ProcessPicqerOrders(pages)
			mux.Unlock()

			// We only load the first time more pages after this we reset the number of pages to 2
			if pages != c.FirstLoadPages {
				pages = c.Pages
			}

			c.saveDataToFile()
			// TODO Save Order Notes to File
			// SaveOrderNotesToFile()

		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func (c PicqerDaemon) loadDataFromFile() {
	file, err := os.Open(c.DataPath + "orders.json")
	if err != nil {
		log.Println("Error while loading Picqer orders from file. Error opening /data/orders.json")
		panic(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error while loading from file. Error while converting JSON structure to object")
		panic(err)
	}

	json.Unmarshal(data, &c.Orders)

	log.Println("Loaded " + strconv.Itoa(len(c.Orders)) + " orders")
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
