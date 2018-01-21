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
	// DeleteVoucher allows you to invalidate a particular voucher
	// https://developers.certcenter.com/v1/reference#deletevoucher
	//
  res, _ := certcenter.DeleteVoucher(&certcenter.DeleteVoucherRequest{VoucherCode: "JDX1UBHC6UA3"})
  fmt.Println(res)
	return
}
