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
	// Check a CommonName against the black list (AlwaysOnSSL only!)
	// plus lets you generate a private key and PEM-encoded CSR
	//
	res, _ := certcenter.ValidateName(&certcenter.ValidateNameRequest{
		CommonName: "www.example.com",
		GeneratePrivateKey: true,
	})
	fmt.Println(res)
	return
}
