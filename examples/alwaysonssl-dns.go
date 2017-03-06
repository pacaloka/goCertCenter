package main

//
// Usage:
//
//	$ ./alwaysonssl-dns fetch your.host.name 				# to get the DNS information
//	$ ./alwaysonssl-dns order your.host.name 180		# to submit the order, 180 = Period of validity in days
//

import (
	certcenter "certcenter.com/go"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func init() {
	certcenter.Bearer = "AValidToken.oauth2.certcenter.com"
}

func fetch(CommonName string) {
	resValidateName, err := certcenter.ValidateName(&certcenter.ValidateNameRequest{
		CommonName:         CommonName,
		GeneratePrivateKey: true,
	})
	if err != nil {
		panic(err)
	}

	if resValidateName.IsQualified == false {
		panic("CommonName is not qualified (blacklisted)")
	}

	err = ioutil.WriteFile("crt.csr", []byte(resValidateName.CSR), 0400)
	if err != nil {
		panic(err)
	}
	fmt.Println("Your CSR has been saved (crt.csr)")

	err = ioutil.WriteFile("crt.key", []byte(resValidateName.PrivateKey), 0400)
	if err != nil {
		panic(err)
	}
	fmt.Println("Your PrivateKey has been saved (crt.key)")

	resDNSData, err := certcenter.DNSData(&certcenter.DNSDataRequest{
		CSR:         resValidateName.CSR,
		ProductCode: "AlwaysOnSSL.AlwaysOnSSL",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nPlease create this %s-record before you submit your order:\n\t%s IN %s \"%s\"\n\n",
		resDNSData.DNSAuthDetails.PointerType,
		resDNSData.DNSAuthDetails.DNSEntry,
		resDNSData.DNSAuthDetails.PointerType,
		resDNSData.DNSAuthDetails.DNSValue)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("\nUsage:\n\t1. $ %s fetch hostname (to get the DNS information)\n", os.Args[0])
		fmt.Printf("\t2. $ %s order hostname period (to submit the order)\n\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Println(os.Args)

	switch os.Args[1] {
	case "fetch":
		fetch(os.Args[2])
		break
	case "order":
		if len(os.Args) < 3 {
			fmt.Printf("\nUsage:\n\t%s order hostname periodInDays\n\n", os.Args[0])
			os.Exit(0)
		}
		csr, err := ioutil.ReadFile("crt.csr")
		if err != nil {
			panic(err)
		}
		Period, err := strconv.Atoi(os.Args[3])
		if err != nil {
			panic(err)
		}
		resOrder, _ := certcenter.Order(&certcenter.OrderRequest{
			OrderParameters: &certcenter.OrderParameters{
				ProductCode:    "AlwaysOnSSL.AlwaysOnSSL",
				CSR:            string(csr),
				DVAuthMethod:   "DNS",
				ValidityPeriod: Period,
			},
		})

		fmt.Println("Certificate fulfillment:")
		fmt.Println("======================================================================")
		fmt.Println("Fulfillment.Certificate:")
		fmt.Println(resOrder.Fulfillment.Certificate)
		fmt.Println("")
		fmt.Println("Fulfillment.Intermediate:")
		fmt.Println(resOrder.Fulfillment.Intermediate)
		fmt.Println("")
		fmt.Println("Fulfillment.PKCS7:")
		fmt.Println(resOrder.Fulfillment.PKCS7)
		fmt.Println("")

		break
	}
}
