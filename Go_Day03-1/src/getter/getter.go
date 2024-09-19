package getter

import (
	"encoding/json"
	"fmt"
	"log"
	"src/types"
)

func RestoransHTTP(rest []types.Place) []string {

	var body []string
	for _, restor := range rest {
		str := "<li>\n"
		str += fmt.Sprintf("<div> %s</div>\n", restor.Name)
		str += fmt.Sprintf("<div> %s</div>\n", restor.Address)
		str += fmt.Sprintf("<div> %s</div>\n", restor.Phone)
		str += "</li>\n"
		body = append(body, str)
	}

	return body
}
func RestorasJSON(rest []types.Place, p int) string {
	var resultPage types.ApiStruct
	resultPage.Places = rest
	resultPage.Name = "Places"
	resultPage.Total = 13649
	resultPage.LastPage = 1365
	resultPage.NextPage = p + 1
	resultPage.PrevPage = p - 1

	res, err := json.MarshalIndent(resultPage, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	return string(res)
}
func RestoransNear(rest []types.Place) string {
	var resultPage types.ApiNear

	resultPage.Places = rest
	resultPage.Name = "Recomended"
	res, err := json.MarshalIndent(resultPage, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	return string(res)
}
