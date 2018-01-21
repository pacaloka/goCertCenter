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
	// Get all modified orders in timespan
  //
	res, _ := certcenter.GetModifiedOrders(&certcenter.GetModifiedOrdersRequest{
		FromDate: time.Now().Add(-10 * time.Minutes),
		ToDate: time.Now(),
		IncludeFulfillment: true,
		IncludeOrderParameters: true,
		IncludeBillingDetails: true,
		IncludeContacts: true,
		IncludeOrganizationInfos: true,
		IncludeDCVStatus: true,
	})
	fmt.Println(res)
	return
}
