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
	// RedeemVoucher let you redeem a previously generated voucher code
	// https://developers.certcenter.com/v1/reference#redeemvoucher
	//
	csr, _ := ioutil.ReadFile("csr")
	fmt.Println(string(csr))
	res, _ := certcenter.RedeemVoucher(&certcenter.RedeemVoucherRequest{
			VoucherCode: "JDX1UBDC6AA1",
			// You don't need OrganizationInfo on DV orders, except for SSL123.
			OrganizationInfo: &certcenter.OrganizationInfo{
				OrganizationName: "Acme LCC",
				OrganizationAddress: &certcenter.OrganizationAddress{
					AddressLine1: "40 5th Ave",
					Region: "NY",
					PostalCode: "12012",
					Country: "US",
					City: "NY",
					Phone: "+1 121 444444",
				},
			},
			OrderParameters: &certcenter.OrderParameters{
				ProductCode: "Thawte.SSL123",
				CSR: string(csr),
				ValidityPeriod: 12,
				DVAuthMethod: "EMAIL",
				ServerCount: 1,
				// Needs to be a valid approver email address.
				// Inquire valid addresses via /ApproverEmail
				ApproverEmail:"postmaster@example.com",
		    SignatureHashAlgorithm: "SHA256-FULL-CHAIN",
			},
			AdminContact: &certcenter.Contact{
				FirstName: "John",
				LastName: "Doe",
				Title: "CEO",
				Phone: "+1 212 999 999",
				Email: "john.doe@example.com",
			},
			TechContact: &certcenter.Contact{
				FirstName: "John",
				LastName: "Doe",
				Title: "CEO",
				Phone: "+1 212 999 999",
				Email: "john.doe@example.com",
			},
		},
	)
	fmt.Println(res)
	return
}
