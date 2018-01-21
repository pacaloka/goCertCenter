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
	// GetUser allows you to inquire information about one or all users
	// https://developers.certcenter.com/v1/reference#getuser
	//
	res, _ := certcenter.GetUser(&certcenter.GetUserRequest{
	certcenter.UserData{
			UsernameOrUserId: "1234567",
		}})
	fmt.Println(res)
	return
}
