package main

import (
	certcenter "certcenter.com/go"
	"fmt"
	_ "io/ioutil"
	_ "time"
)

// Set your valid OAuth2 Bearer
// (see https://developers.certcenter.com/docs/authentication)

func init() {
	certcenter.Bearer = "AValidToken.oauth2.certcenter.com"
}

func main() {
	// Retrieve appropriate data for DNS-based validation (AlwaysOnSSL only!)
	//
	csr, _ := ioutil.ReadFile("csr")
	res, _ := certcenter.DNSData(&certcenter.DNSDataRequest{
		CSR:         string(csr),
		ProductCode: "AlwaysOnSSL.AlwaysOnSSL",
	})
	fmt.Println(res)
	return
}
