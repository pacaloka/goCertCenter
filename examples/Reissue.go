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
	// Reissue a particular order
  //
	res, _ := certcenter.Reissue(&certcenter.ReissueRequest{
		CertCenterOrderID: 123456789,
		OrderParameters:certcenter.ReissueOrderParameters{
			CSR: "#CSR#",
			DVAuthMethod: "EMAIL",
			SignatureHashAlgorithm: "SHA256-FULL-CHAIN",
		},
		ReissueEmail: "valid-approver@example.com",
	})
	fmt.Println(res)
	return
}
