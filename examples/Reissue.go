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
	// Reissue a particular order
  //
	res, _ := certcenter.Reissue(&certcenter.ReissueRequest{
		CertCenterOrderID: 123456789,
		OrderParameters:certcenter.ReissueOrderParameters{
			CSR: "#CSR#",
			DVAuthMethod: "EMAIL",
			SignatureHashAlgorithm: "SHA256-FULL-CHAIN",
			DomainApprovers: &certcenter.DomainApprovers{
				DomainApprover: []certcenter.DomainApproverItem{certcenter.DomainApproverItem{
					Domain: "certcenter.com",
					Approvers: []certcenter.Approver{certcenter.Approver{
						ApproverEmail: "domains@certcenter.com",
					}},
				}},
			},
		},
		ReissueEmail: "valid-approver@example.com",
	})
	fmt.Println(res)
	return
}
