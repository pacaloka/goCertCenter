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
	// CreateVoucher creates a coupon code which can later be redeemded.
	// https://developers.certcenter.com/v1/reference#createvoucher
	//
	res, _ := certcenter.CreateVoucher(&certcenter.CreateVoucherRequest{
		OrderParameters:certcenter.OrderParameters{
			ProductCode: "Thawte.SSL123",
			PartnerOrderID: "My voucher order id (optional)",
			ServerCount: 1,
			SubjectAltNameCount: 0,
			ValidityPeriod: 12,
		},
	})
	fmt.Println(res)
	return
}
