package main

import (
	"fmt"
	"io/ioutil"
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
				ProductCode: "Comodo.InstantSSL",
				CSR: string(csr),
				ValidityPeriod: 24,
				ServerCount: 1,
				DomainApprovers: &certcenter.DomainApprovers{
					DomainApprover: []certcenter.DomainApproverItem{certcenter.DomainApproverItem{
						Domain: "certcenter.com",
						Approvers: []certcenter.Approver{certcenter.Approver{
							ApproverEmail: "domains@certcenter.com",
						}},
					}},
				},
			},
			OrganizationInfo: &certcenter.OrganizationInfo{
				OrganizationName: "CertCenter AG",
				OrganizationAddress: &certcenter.OrganizationAddress{
					AddressLine1: "Bleichstraße 8a",
					PostalCode: "35390",
					City: "Gießen",
					Region: "Hessen",
					Country: "DE",
					Phone: "+49 641 80899520",
					Fax: "+49 641 80899521",
				},
			},
			AdminContact: &certcenter.Contact{
				FirstName: "John",
				LastName: "Doe",
				Title: "CTO",
				Phone: "+49 641 80899520",
				Fax: "+49 641 80899521",
				Email: "john.doe@example.com",
				OrganizationName: "CertCenter AG",
				OrganizationAddress: &certcenter.OrganizationAddress{
					AddressLine1: "Bleichstraße 8a",
					PostalCode: "35390",
					City: "Gießen",
					Region: "Hessen",
					Country: "DE",
					Phone: "+49 641 80899520",
					Fax: "+49 641 80899521",
				},
			},
			TechContact: &certcenter.Contact{
				FirstName: "John",
				LastName: "Doe",
				Title: "CTO",
				Phone: "+49 641 80899520",
				Fax: "+49 641 80899521",
				Email: "john.doe@example.com",
				OrganizationName: "CertCenter AG",
				OrganizationAddress: &certcenter.OrganizationAddress{
					AddressLine1: "Bleichstraße 8a",
					PostalCode: "35390",
					City: "Gießen",
					Region: "Hessen",
					Country: "DE",
					Phone: "+49 641 80899520",
					Fax: "+49 641 80899521",
				},
			},
		},
	)
	fmt.Println(res)
	return
}
