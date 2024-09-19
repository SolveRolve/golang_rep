package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"src/api"

	"github.com/lizrice/secure-connections/utils"
)

func main() {

	server := api.NewServer()
	r := http.NewServeMux()
	h := api.HandlerFromMux(server, r)

	t := &tls.Config{
		GetCertificate: utils.CertReqFunc(
			"../ca/127.0.0.1/cert.pem",
			"../ca/127.0.0.1/key.pem"),
		VerifyPeerCertificate: utils.CertificateChains,
	}

	s := &http.Server{
		Handler:   h,
		Addr:      "127.0.0.1:3333",
		TLSConfig: t,
	}

	http.HandleFunc("/buy_candy", server.BuyCandy)

	log.Fatal(s.ListenAndServeTLS("", ""))
}
