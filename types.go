package certcenter

import (
	"fmt"
	"net/http"
)

var Bearer string

const (
	CC_PARAM_TYPE_QS = 1<<iota
	CC_PARAM_TYPE_PATH
	CC_PARAM_TYPE_BODY
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	} 
}

/* Represents a API request
*/
type apiRequest struct {
	method string
	result interface{}
	request interface{}
	client *http.Client
	statusCode int
}

/* Represents a GET /Profile response
*/
type ProfileResult struct {
	AuthType string
	AuthorizationID int64
	Country string
	Currency string
	CustomerID int64
	Locale string
	OAuth2_Token string
	Scope string
	Timezone string
}

/* Represents a GET /Limit response
*/
type LimitResult struct {
	Success bool`json:"success"`
	LimitInfo struct {
		Limit float64
		Used float64
	}
}

/* Represents a GET /Products response
*/
type ProductsResult struct {
	Success bool`json:"success"`
	Products []string
}

/* Represents a GET /ProductDetails response
*/
type ProductDetailsResult struct {
	Success bool`json:"success"`
	ProductDetails struct {
		CA string
		Currency string
		Features []string
		Licenses int
		MaxValidityPeriod int
		Price float64
		ProductCode string
		ProductName string
		RefundPeriod int
		RenewPeriod int
		SANFeatures []string
		SANHostPrice float64
		SANMaxHosts int
		SANPackagePrice float64
		SANPackageSize int
	}
}

/* Represents a GET /ProductDetails request
*/
type ProductDetailsRequest struct {
	ProductCode string
}

/* Represents a GET /Quote response
*/
type QuoteResult struct {
	Success bool`json:"success"`
	Currency string
	OrderParameters struct {
		ProductCode string
		ServerCount int
		SubjectAltNameCount int
		ValidityPeriod int
	}
	Price float64
}

/* Represents a GET /Quote request
*/
type QuoteRequest struct {
	ProductCode string
	SubjectAltNameCount int
	ValidityPeriod int
	ServerCount int
}





