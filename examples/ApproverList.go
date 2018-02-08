package main

import (
	certcenter "certcenter.com/go"
	"fmt"
)

// Set your valid OAuth2 Bearer
// (see https://developers.certcenter.com/docs/authentication)
func init() {
	certcenter.Bearer = "AValidToken.oauth2.certcenter.com"
}

func main() {
	// Get valid email addresses for email based DV validation
	res, _ := certcenter.ApproverList(&certcenter.ApproverListRequest{
		CommonName:  "www.example.com",
		ProductCode: "GeoTrust.QuickSSLPremium",
		// optional value (comma seperated list of hosts)
		//DNSNames: "www.example.org, www.example.co.uk",
	})

	if res.DomainApprovers != (&certcenter.DomainApprovers{}) {
		// Make use of the new DomainApprovers structure
		for i := 0; i < len(res.DomainApprovers.DomainApprover); i++ {
			fmt.Println(res.DomainApprovers.DomainApprover[i])
		}
	} else {
		// Use the legacy structure (provided by non-compatible CAs)
		for i := 0; i < len(res.ApproverList); i++ {
			fmt.Println(res.ApproverList[i])
		}
	}

	return
}
