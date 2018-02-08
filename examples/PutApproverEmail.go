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
	// Re-set the approvers email address
	res, _ := certcenter.PutApproverEmail(&certcenter.PutApproverEmailRequest{
		CertCenterOrderID: 123456789,
		ApproverEmail:     "valid-approver@example.com",
	})
	fmt.Println(res)
	return
}
