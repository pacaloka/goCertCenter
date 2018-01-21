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
	// UpdateUser allows you to update an user
	// https://developers.certcenter.com/v1/reference#updateuser
	//
	res, _ := certcenter.UpdateUser(&certcenter.UpdateUserRequest{
	certcenter.UserData{
			UsernameOrUserId: "1234567",
			FullName: "John Doe",
			Email: "john@example.org",
		}})
	fmt.Println(res)
	return
}
