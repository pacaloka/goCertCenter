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
	// Re-send the approver email to the approvers address
	res, _ := certcenter.ResendApproverEmail(&certcenter.ResendApproverEmailRequest{
		CertCenterOrderID: 123456789,
	})
	fmt.Println(res)
	return
}
