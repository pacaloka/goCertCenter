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
	// Get my profile information
	//
	res, _ := certcenter.Profile()
	fmt.Println(res)
}
