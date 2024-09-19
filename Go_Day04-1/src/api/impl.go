package api

import (
	"encoding/json"
	"io"
	"net/http"
	"src/cowcode"
)

var candys = []struct {
	name string
	cost int
}{
	{"CE", 10},
	{"AA", 15},
	{"NT", 17},
	{"DE", 21},
	{"YR", 23},
}

type Server struct{}

var _ ServerInterface = (*Server)(nil)
var err400 = "some error in input data"
var err402 = "not enough money"
var resp400 = &BuyCandy400JSONResponse{Error: &err400}
var resp402 = &BuyCandy402JSONResponse{&err402}

func NewServer() Server {
	return Server{}
}

func (b BuyCandyJSONBody) validCandy() int {

	for _, candy := range candys {
		if candy.name == b.CandyType {
			if b.Money < candy.cost*b.CandyCount {
				return 402
			} else {
				return b.Money - candy.cost*b.CandyCount
			}
		}
	}
	return 400
}

func (Server) BuyCandy(w http.ResponseWriter, r *http.Request) {
	var reqBodyJSON BuyCandyJSONBody
	var sucs BuyCandy201JSONResponse

	bodyBytes, err := io.ReadAll(r.Body)
	err = json.Unmarshal(bodyBytes, &reqBodyJSON)

	if err != nil {
		resp400.VisitBuyCandyResponse(w)
		return
	}
	resp := reqBodyJSON.validCandy()
	if resp == 400 {
		resp400.VisitBuyCandyResponse(w)
	} else if resp == 402 {
		resp402.VisitBuyCandyResponse(w)
	} else {
		thx := cowcode.CowTalk("Thank you!")
		sucs.Thanks = &thx
		sucs.Change = &resp
		sucs.VisitBuyCandyResponse(w)
	}
}
