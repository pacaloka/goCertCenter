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
	// Get filtered orders
	//
	res, _ := certcenter.GetOrders(&certcenter.GetOrdersRequest{
		Status:                   "COMPLETE", // COMPLETE, PENDING, CANCELLED, REVOKED
		ProductType:              "SSL",      // SSL, CODESIGN, SMIME, TRUSTSEAL
		CommonName:               "%",
		IncludeFulfillment:       true,
		IncludeOrderParameters:   true,
		IncludeBillingDetails:    true,
		IncludeContacts:          true,
		IncludeOrganizationInfos: true,
		IncludeDCVStatus:         true,
	})
	fmt.Println(res)
	return
}
