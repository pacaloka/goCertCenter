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
	// KvStore allows you to use CertCenter's free key-value storage
	// for DV FILE authentication with mod_fauth. Details:
	// https://developers.certcenter.com/v1/docs/howto-order-alwaysonssl-symantec-ee-certificates#section-4-order-procedure
	//
	certcenter.KvStoreAuthorizationKey = ""
	res, err := certcenter.KvStore(&certcenter.KeyValueStoreRequest{
			Key: "test.example.com",
			Value: "201701260800495t3djr2zqhqfvgg1cpjmgs5zx4kd7w51w3cuge90sokdavg6li",
	}); if err!=nil {
	  panic("..")
	}
	return
}
