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
	// CreateUser allows you to create a new UI user for your organizations account
	// https://developers.certcenter.com/v1/reference#createuser
	//
	res, _ := certcenter.CreateUser(&certcenter.CreateUserRequest{
		certcenter.UserData{
			FullName: "John Doe",
			Email:    "john@example.org",
			Username: "johndoes",
			Password: "cOmpL3xx/",
			Mobile:   "",
			Roles:    []string{"PROCUREMENT"},
			Locale:   "en_US",
			Timezone: "US/Pacific",
		},
	})
	fmt.Println(res)
	return
}
