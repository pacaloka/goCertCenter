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
	// UpdateUser allows you to update an user
	// https://developers.certcenter.com/v1/reference#updateuser
	//
	res, _ := certcenter.UpdateUser(&certcenter.UpdateUserRequest{
		certcenter.UserData{
			UsernameOrUserId: "1234567",
			FullName:         "John Doe",
			Email:            "john@example.org",
		}})
	fmt.Println(res)
	return
}
