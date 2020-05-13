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
	// Check a CommonName against the black list (AlwaysOnSSL only!)
	//
	res, _ := certcenter.ValidateName(&certcenter.ValidateNameRequest{
		CommonName:         "ValidateName.go",
	})
	fmt.Println(res)
	return
}
