package main

import (
	"fmt"
//	"io/ioutil"
	certcenter "../../goCertCenter"
)

/* Set your Authorization Token
 * 
 */
func init() {
	certcenter.Bearer = "ABCDEFGHIJKLMNOP.oauth2.certcenter.com"
}

func main() {

	// Limit
	res, _ := certcenter.Limit()
	fmt.Println(res)

/*

	// Profile
	res, _ := certcenter.Profile()
	fmt.Println(res)


	// Products
	res, _ := certcenter.Products()
	fmt.Println(res)


	// ProductDetails
	res, _ := certcenter.ProductDetails("GeoTrust.QuickSSLPremium")
	fmt.Println(res)


	// Quote
	res, _ := certcenter.Quote(&certcenter.QuoteRequest{
		ProductCode: "Symantec.SecureSiteEV",
		SubjectAltNameCount: 0,
		ValidityPeriod: 24,
		ServerCount: 1,
	})
	fmt.Println(res)


	// ValidateCSR
	csr, _ := ioutil.ReadFile("csr")
	res, _ := certcenter.ValidateCSR(&certcenter.ValidateCSRRequest{
		CSR: string(csr),
	})
	fmt.Println(res)

*/

	return 
}
