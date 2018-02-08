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
	// Get all modified orders in timespan
	//
	res, _ := certcenter.GetModifiedOrders(&certcenter.GetModifiedOrdersRequest{
		FromDate:                 time.Now().Add(-10 * time.Minutes),
		ToDate:                   time.Now(),
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
