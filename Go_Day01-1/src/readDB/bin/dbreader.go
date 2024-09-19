package bin

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type DBReader interface {
	ParseFile(NameFile string) ([]byte, error)
	PrettyPrint() error
	CopyToData() *Data
	CountCake() int
}
type DataJSON struct {
	Cake []struct {
		Name        string `json:"name"`
		Time        string `json:"time"`
		Ingredients []struct {
			IngredientName  string `json:"ingredient_name"`
			IngredientCount string `json:"ingredient_count"`
			Ingredient_unit string `json:"ingredient_unit"`
		} `json:"ingredients"`
	} `json:"cake"`
}
type DataXML struct {
	Cake []struct {
		Name        string `xml:"name"`
		Stovetime   string `xml:"stovetime"`
		Ingredients struct {
			Item []struct {
				Itemname  string `xml:"itemname"`
				Itemcount string `xml:"itemcount"`
				Itemunit  string `xml:"itemunit"`
			} `xml:"item"`
		} `xml:"ingredients"`
	} `xml:"cake"`
}
type Data struct {
	Cake [10]struct {
		Name        string
		Stovetime   string
		Ingredients struct {
			Item [10]struct {
				Name  string
				Count string
				Unit  string
			}
		}
	}
}

func (d *DataJSON) ParseFile(NameFile string) ([]byte, error) {

	jsonFile, err := os.Open(NameFile)

	if err != nil {
		return nil, err
	}

	fileData, err := io.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileData, &d)

	if err != nil {
		return nil, err
	}

	return fileData, nil
}
func (d DataJSON) PrettyPrint() error {
	fmt.Println(d)
	byteForPrint, err := xml.MarshalIndent(d, "", "    ")

	if err != nil {

		return err
	}
	fmt.Println(string(byteForPrint))
	return nil
}
func (d DataXML) PrettyPrint() error {
	byteForPrint, err := json.MarshalIndent(d, "", "    ")

	if err != nil {
		return err
	}
	fmt.Println(string(byteForPrint))
	return nil
}
func (d *DataXML) ParseFile(NameFile string) ([]byte, error) {

	if len(NameFile) == 0 {
		return nil, errors.New("XML file dosent exist")
	}
	var xmlFile, err = os.Open(NameFile)

	if err != nil {
		return nil, err
	}

	fileData, err := io.ReadAll(xmlFile)

	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(fileData, &d)

	if err != nil {
		return nil, err
	}

	return fileData, nil
}
func InitReader(f string) (inter DBReader) {
	oldData := DataXML{}
	newData := DataJSON{}
	if strings.Contains(f, ".json") {
		inter = &newData
	} else {
		inter = &oldData
	}
	return
}
func (d *DataXML) PrintData() string {
	return fmt.Sprint(d)
}

func (d *DataXML) CopyToData() *Data {
	res := Data{}
	for index, data := range d.Cake {
		res.Cake[index].Name = data.Name
		res.Cake[index].Stovetime = data.Stovetime
		for idx, prod := range d.Cake[index].Ingredients.Item {
			res.Cake[index].Ingredients.Item[idx].Unit = prod.Itemunit
			res.Cake[index].Ingredients.Item[idx].Name = prod.Itemname
			res.Cake[index].Ingredients.Item[idx].Count = prod.Itemcount
		}
	}
	return &res
}
func (d *DataJSON) CopyToData() *Data {
	res := Data{}

	for index, data := range d.Cake {
		res.Cake[index].Name = data.Name
		res.Cake[index].Stovetime = data.Time
		for idx, prod := range d.Cake[index].Ingredients {
			res.Cake[index].Ingredients.Item[idx].Unit = prod.Ingredient_unit
			res.Cake[index].Ingredients.Item[idx].Name = prod.IngredientName
			res.Cake[index].Ingredients.Item[idx].Count = prod.IngredientCount
		}
	}
	return &res
}
func (d *DataJSON) CountCake() int {
	return len(d.Cake)
}
func (d *DataXML) CountCake() int {
	return len(d.Cake)
}
