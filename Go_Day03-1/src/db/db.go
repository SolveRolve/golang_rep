package db

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type restoran struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Location struct {
		Latitude  string `json:"lat"`
		Longitude string `json:"lon"`
	} `json:"location"`
}

func makeNewClient(s string) *elasticsearch.Client {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	res, err := es.Info()

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	if res, err = es.Indices.Delete([]string{s}); err != nil {
		log.Fatalf("Cannot delete index: %s", err)
	}

	res, err = es.Indices.Create(
		s,
		es.Indices.Create.WithBody(strings.NewReader(mapping)),
	)

	if err != nil {
		log.Fatal(err)
	}

	if res.IsError() {
		log.Fatalf("Cannot create index: %s", res)
	}
	return es
}
func getPlaces() []*restoran {
	var restorans []*restoran
	file, err := os.Open("../materials/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	newScan := bufio.NewScanner(file)
	newScan.Split(bufio.ScanLines)

	for newScan.Scan() {
		line := newScan.Text()
		arrData := strings.Split(line, "\t")
		if len(arrData) < 6 {
			continue
		}
		restorans = append(restorans, &restoran{
			Name:    arrData[1],
			Address: arrData[2],
			Phone:   arrData[3],
			Location: struct {
				Latitude  string `json:"lat"`
				Longitude string `json:"lon"`
			}{arrData[5], arrData[4]},
		})

	}
	return restorans
}
func setPlaces(restorans []*restoran, es *elasticsearch.Client) {

	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:      "places",
		Client:     es,
		NumWorkers: 5,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer bulkIndexer.Close(context.Background())

	for i, place := range restorans {

		jsonPlace, _ := json.Marshal(place)

		err = bulkIndexer.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: strconv.Itoa(i),
				Body:       bytes.NewReader(jsonPlace),
			})

	}
}
func IndexPlaces(NameIndex string) {

	restorans := getPlaces()
	es := makeNewClient(NameIndex)
	setPlaces(restorans, es)

}

const mapping = `
    {
      "settings": {
        "number_of_shards": 1,
		"max_result_window" : 20000
      },
      "mappings": {
        "properties": {
    "name": {
        "type":  "text"
    },
    "address": {
        "type":  "text"
    },
    "phone": {
        "type":  "text"
    },
    "location": {
      "type": "geo_point"
    }
  }
        }
      }
    }`
