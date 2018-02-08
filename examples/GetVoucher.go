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
	// GetVoucher allows you to inquire information about particular voucher
	// https://developers.certcenter.com/v1/reference#getvoucher
	//
	res, _ := certcenter.GetVoucher(&certcenter.GetVoucherRequest{VoucherCode: "JDX1UBHC6UA3"})
	fmt.Println(res.Vouchers[0].OrderParameters)
	return
}
