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
	// GetVoucherOrderAnonymously allows a 3rd party to get information about a
	// particular Order (based on a voucher) w/o authentication
	// https://developers.certcenter.com/v1/reference#getvoucherorderanonymously
	//
	res, _ := certcenter.GetVoucherOrderAnonymously(&certcenter.GetVoucherRequest{VoucherCode: "JDX1UBHC6UA3"})
	fmt.Println(res.Vouchers[0].OrderParameters)
	return
}
