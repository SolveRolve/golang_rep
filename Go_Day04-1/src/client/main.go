package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/lizrice/secure-connections/utils"
	"io"
	"log"
	"net/http"
	"os"
	"src/api"
	"strings"
)

func initReqvest() []byte {

	candyArg := flag.String("k", "", "what type candy you want")
	manyArg := flag.Int("c", 0, "how many candy you have")
	moneyArg := flag.Int("m", 0, "how many money you have")
	flag.Parse()

	reqBody, err := json.Marshal(&api.BuyCandyJSONBody{
		CandyCount: *manyArg,
		CandyType:  *candyArg,
		Money:      *moneyArg,
	})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return reqBody
}
func must(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)

	}
}
func getClient() *http.Client {
	data, err := os.ReadFile("../ca/minica.pem")
	must(err)

	cp, err := x509.SystemCertPool()
	must(err)

	cp.AppendCertsFromPEM(data)

	config := &tls.Config{
		//InsecureSkipVerify: true,
		RootCAs: cp,
		GetCertificate: utils.CertReqFunc(
			"",
			""),
		VerifyPeerCertificate: utils.CertificateChains,
	}
	a := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return a
}

// -k AA -c 2 -m 50
func main() {

	client, err := api.NewClientWithResponses("https://127.0.0.1:3333", api.WithHTTPClient(getClient()))
	must(err)

	readerReq := strings.NewReader(string(initReqvest()))
	resp, er := client.BuyCandyWithBody(context.Background(), "application/json", readerReq)
	must(er)

	defer resp.Body.Close()
	req := api.BuyCandy201JSONResponse{}
	bodyBytes, err := io.ReadAll(resp.Body)
	must(err)
	must(json.Unmarshal(bodyBytes, &req))
	fmt.Printf("%s\nYour change is %d", *req.Thanks, *req.Change)

}
