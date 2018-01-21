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
	// Revoke a certificate
  //
	res, _ := certcenter.Revoke(&certcenter.RevokeRequest{
		CertCenterOrderID: 123456789,
		// Optional parameters
		RevokeReason: "Key compromised",
		Certificate: "#PEM-encoded-X.509-Certificate#",
	})
	fmt.Println(res)
	return
}
