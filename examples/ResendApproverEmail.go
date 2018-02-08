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
	// Re-send the approver email to the approvers address
	res, _ := certcenter.ResendApproverEmail(&certcenter.ResendApproverEmailRequest{
		CertCenterOrderID: 123456789,
	})
	fmt.Println(res)
	return
}
