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
	// Get the CA's User Subscriber Agreement
	//
	res, _ := certcenter.UserAgreement("Symantec.SecureSite")
	fmt.Println(res)
	return
}
