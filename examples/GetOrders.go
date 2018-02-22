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
		IncludeFulfillment:       false,
		IncludeOrderParameters:   true,
		IncludeBillingDetails:    false,
		IncludeContacts:          false,
		IncludeOrganizationInfos: false,
		IncludeDCVStatus:         true,
	})

	if len(res.OrderInfos) > 0 {
		for _, OrderInfo := range res.OrderInfos {
			fmt.Println("----")
			fmt.Printf("CertCenterOrderID: %d\n"+
				"Status: %s\n"+
				"Common Name: %s\n"+
				"DNSnames: %s\n"+
				"Validity Period: %d\n",
				OrderInfo.CertCenterOrderID,
				OrderInfo.OrderStatus.MajorStatus,
				OrderInfo.CommonName,
				OrderInfo.OrderParameters.SubjectAltNames,
				OrderInfo.OrderParameters.ValidityPeriod)
			if len(OrderInfo.DCVStatus) > 0 {
				for _, DCVDomainStatus := range OrderInfo.DCVStatus {
					fmt.Printf("DCV Information:\n\n"+
						"  Authentication Domain: %s\n"+
						"  Current status: %s\n"+
						"  Last Check: %s\n"+
						"  Last Update: %s\n\n",
						DCVDomainStatus.Domain,
						DCVDomainStatus.Status,
						DCVDomainStatus.LastCheckDate,
						DCVDomainStatus.LastUpdateDate,
					)
				}
			}
		}
	}

	return
}
