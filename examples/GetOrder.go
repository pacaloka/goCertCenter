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
	// Get a particular order by CertCenterOrderID
	//
	res, _ := certcenter.GetOrder(&certcenter.GetOrderRequest{
		CertCenterOrderID:        123456789,
		IncludeFulfillment:       true,
		IncludeOrderParameters:   true,
		IncludeBillingDetails:    true,
		IncludeContacts:          true,
		IncludeOrganizationInfos: true,
		IncludeDCVStatus:         true,
	})

	if len(res.OrderInfo.DCVStatus)>0 {
		for _, DCVDomainStatus := range res.OrderInfo.DCVStatus {
			fmt.Printf("Authentication Domain: %s\n"+
					"Current status: %s\n"+
					"Last Check: %s\n"+
					"Last Update: %s\n\n",
					DCVDomainStatus.Domain,
					DCVDomainStatus.Status,
					DCVDomainStatus.LastCheckDate,
					DCVDomainStatus.LastUpdateDate,
				 )
		}
	}

	return
}
