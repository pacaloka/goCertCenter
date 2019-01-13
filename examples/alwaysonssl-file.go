package main

import (
	certcenter "certcenter.com/go"
	"fmt"
	"os"
)

func init() {
	certcenter.Bearer = "AValidToken.oauth2.certcenter.com"
	certcenter.KvStoreAuthorizationKey = "aValidTokenAuthKey"
}

func order(CommonName string, Period int) {
	resValidateName, err := certcenter.ValidateName(&certcenter.ValidateNameRequest{
		CommonName:         CommonName,
	})
	if err != nil {
		panic("..")
	}

	if resValidateName.IsQualified == false {
		panic("CommonName is not qualified (blacklisted)")
	}

	resFileData, err := certcenter.FileData(&certcenter.FileDataRequest{
		CSR:         "#CSR#",
		ProductCode: "AlwaysOnSSL.AlwaysOnSSL",
	})
	if err != nil {
		panic("..")
	}

	_, err = certcenter.KvStore(&certcenter.KeyValueStoreRequest{
		Key:   CommonName,
		Value: resFileData.FileAuthDetails.FileContents,
	})
	if err != nil {
		panic("..")
	}

	resOrder, _ := certcenter.Order(&certcenter.OrderRequest{
		OrderParameters: &certcenter.OrderParameters{
			ProductCode:    "AlwaysOnSSL.AlwaysOnSSL",
			CSR:            "#CSR#",
			DVAuthMethod:   "FILE",
			ValidityPeriod: Period,
		},
	})

	fmt.Println("Certificate fulfillment:")
	fmt.Println(resOrder)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("\nUsage:\n\t%s hostname\n\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Println(os.Args)
	order(os.Args[1], 180)
}
