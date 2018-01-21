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
	// DeleteUser allows you to remove a specific user
	// https://developers.certcenter.com/v1/reference#deleteuser
	//
	res, _ := certcenter.DeleteUser(&certcenter.DeleteUserRequest{
		UsernameOrUserId: "1234567",
	})
	fmt.Println(res)
	return
}
