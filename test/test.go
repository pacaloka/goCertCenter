package main

import (
	"fmt"
	certcenter "github.com/certcenter/goCertCenter"
)

/* Set your Authorization Token
 * 
 */
func init() {
	certcenter.Bearer = "XYZ.oauth2.certcenter.com"
}

func main() {
	res, _ := certcenter.Profile()
	fmt.Println(res)
/*
	res, _ := certcenter.Limit()
	fmt.Println(res)

	res, _ := certcenter.Products()
	fmt.Println(res)

	res, _ := certcenter.ProductDetails("GeoTrust.QuickSSLPremium")
	fmt.Println(res)

	res, _ := certcenter.Quote(&certcenter.QuoteRequest{
		ProductCode: "Symantec.SecureSiteEV",
		SubjectAltNameCount: 0,
		ValidityPeriod: 24,
		ServerCount: 1,
	})
	fmt.Println(res)
*/
	return 
}
