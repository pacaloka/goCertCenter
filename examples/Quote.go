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
	// Get a Quote
	//
	res, _ := certcenter.Quote(&certcenter.QuoteRequest{
		ProductCode:         "Symantec.SecureSiteEV",
		SubjectAltNameCount: 0,
		ValidityPeriod:      24,
		ServerCount:         1,
	})
	fmt.Println(res)
	return
}
