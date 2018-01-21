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
	// Get a particular order by CertCenterOrderID
  //
	res, _ := certcenter.GetOrder(&certcenter.GetOrderRequest{
		CertCenterOrderID: 123456789,
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
