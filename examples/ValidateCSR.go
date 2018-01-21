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
	// Validate a CSR
  //
	csr, _ := ioutil.ReadFile("csr")
	res, _ := certcenter.ValidateCSR(&certcenter.ValidateCSRRequest{
		CSR: string(csr),
	})
	fmt.Println(res)
	return
}
