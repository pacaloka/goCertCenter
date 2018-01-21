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

	// Order a certificate
	csr, _ := ioutil.ReadFile("csr")
	res, _ := certcenter.Order(&certcenter.OrderRequest{
			OrderParameters: &certcenter.OrderParameters{
				ProductCode: "RapidSSL.RapidSSL",
				CSR: string(csr),
				ValidityPeriod: 24,
				ApproverEmail:"domains@certcenter.com",
			},
			AdminContact: &certcenter.Contact{
				FirstName: "John",
				LastName: "Doe",
				Phone: "+1 212 999 999",
				Email: "john.doe@example.com",
			},
			TechContact: &certcenter.Contact{
				FirstName: "John",
				LastName: "Doe",
				Phone: "+1 212 999 999",
				Email: "john.doe@example.com",
			},
		},
	)
	fmt.Println(res)
	return
}
