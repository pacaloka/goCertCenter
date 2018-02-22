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
	res, err := certcenter.BaseDomain(&certcenter.BaseDomainRequest{FQDN: "www.certcenter.co.uk"});
	if err==nil {
		fmt.Printf("FQDN's base domain is: %s\n", res.Domain)
	}

	return
}
