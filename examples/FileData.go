package main

import (
	"fmt"
	_ "io/ioutil"
	_ "time"
	certcenter "certcenter.com/go"
)

// Set your valid OAuth2 Bearer
// (see https://developers.certcenter.com/docs/authentication)

func init() {
	certcenter.Bearer = "AValidToken.oauth2.certcenter.com"
}

func main() {
	// Retrieve appropriate data for FILE-based validation (AlwaysOnSSL only!)
	//
	csr, _ := ioutil.ReadFile("csr")
	res, _ := certcenter.FileData(&certcenter.FileDataRequest{
		CSR: string(csr),
		ProductCode: "AlwaysOnSSL.AlwaysOnSSL",
	})
	fmt.Println(res)
	return
}
