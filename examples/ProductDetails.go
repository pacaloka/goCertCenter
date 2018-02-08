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
	// Inquire details about a product
	//
	res, _ := certcenter.ProductDetails("GeoTrust.QuickSSLPremium")
	fmt.Println(res)
	return
}
