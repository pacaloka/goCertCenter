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
	// GetVouchers allows you to inquire information about all your vouchers
	// https://developers.certcenter.com/v1/reference#getvouchers
	//
	res, _ := certcenter.GetVouchers()
	fmt.Println(res.Vouchers[0].OrderParameters)
	return
}
