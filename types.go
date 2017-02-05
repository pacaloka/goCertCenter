package certcenter

import (
	"fmt"
	"net/http"
	"time"
)

// Bearer represents the authentication token you're going to use
var Bearer string

const (
	// CC_PARAM_TYPE_QS is QueryString (eg. ?CertCenterOrderId=123)
	CC_PARAM_TYPE_QS = 1 << iota
	// CC_PARAM_TYPE_PATH is Path (eg. /:CertCenterOrderId/)
	CC_PARAM_TYPE_PATH
	// CC_PARAM_TYPE_BODY is Body (JSON POST)
	CC_PARAM_TYPE_BODY
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// Represents an API request
type apiRequest struct {
	method     string
	httpMethod string
	url        string
	result     interface{}
	request    interface{}
	client     *http.Client
	statusCode int
}

// ProfileResult represents a GET /Profile response
type ProfileResult struct {
	AuthType        string
	AuthorizationID int64
	Country         string
	Currency        string
	CustomerID      int64
	Locale          string
	OAuth2Token     string `json:"OAuth2Token"`
	Scope           string
	Timezone        string
}

// LimitResult represents a GET /Limit response
type LimitResult struct {
	Success   bool `json:"success"`
	LimitInfo struct {
		Limit float64
		Used  float64
	}
}

// ProductsResult represents a GET /Products response
type ProductsResult struct {
	Success  bool `json:"success"`
	Products []string
}

// ProductDetailsResult represents a GET /ProductDetails response
type ProductDetailsResult struct {
	Success        bool `json:"success"`
	ProductDetails struct {
		CA                string
		Currency          string
		Features          []string
		Licenses          int
		MaxValidityPeriod int
		Price             float64
		ProductCode       string
		ProductName       string
		RefundPeriod      int
		RenewPeriod       int
		SANFeatures       []string
		SANHostPrice      float64
		SANMaxHosts       int
		SANPackagePrice   float64
		SANPackageSize    int
	}
}

// ProductDetailsRequest represents a GET /ProductDetails request
type ProductDetailsRequest struct {
	ProductCode string
}

// QuoteResult represents a GET /Quote response
type QuoteResult struct {
	Success         bool `json:"success"`
	Currency        string
	OrderParameters struct {
		ProductCode         string
		ServerCount         int
		SubjectAltNameCount int
		ValidityPeriod      int
	}
	Price float64
}

// QuoteRequest represents a GET /Quote request
type QuoteRequest struct {
	ProductCode         string
	SubjectAltNameCount int
	ValidityPeriod      int
	ServerCount         int
}

// ValidateCSRResult represents a POST /ValidateCSR response
type ValidateCSRResult struct {
	Success   bool `json:"success"`
	ParsedCSR struct {
		CommonName             string
		Organization           string
		OrganizationUnit       string
		Email                  string
		State                  string
		Locality               string
		Country                string
		KeyLength              int
		SignaturAlgorithm      string
		KeyEncryptionAlgorithm string
	}
}

// ValidateCSRRequest represents a POST /ValidateCSR request
type ValidateCSRRequest struct {
	CSR string // PEM-encoded PKCS#10
}

// UserAgreementRequest represents a GET /ProductDetails response
type UserAgreementRequest struct {
	ProductCode string
}

// UserAgreementResult represents a GET /ProductDetails request
type UserAgreementResult struct {
	Success       bool `json:"success"`
	ProductCode   string
	UserAgreement string
}

// ApproverListRequest represents a GET /ApproverList response
type ApproverListRequest struct {
	CommonName  string
	ProductCode string
}

// ApproverListResult represents a GET /ApproverList request
type ApproverListResult struct {
	Success      bool `json:"success"`
	ApproverList []struct {
		ApproverEmail string
		ApproverType  string // Domain, Generic
	}
}

// OrderResult represents a POST /Order response
type OrderResult struct {
	Success           bool `json:"success"`
	Timestamp         time.Time
	CertCenterOrderID int
	OrderParameters   struct {
		CSR                    string // PEM-encoded PKCS#10
		IsCompetitiveUpgrade   bool
		IsRenewal              bool
		PartnerOrderID         string
		ProductCode            string
		ServerCount            int
		SignatureHashAlgorithm string
		SubjectAltNameCount    int
		SubjectAltNames        []string
		ValidityPeriod         int // 12 or 24 month (days for AlwaysOnSSL, min. 180, max. 365)
		// AlwaysOnSSL (Symantec Encryption Everywhere) only:
		DVAuthMethod string // DNS, EMAIL

	}
	// AlwaysOnSSL (Symantec Encryption Everywhere) only:
	Fulfillment struct {
		Certificate       string
		Certificate_PKCS7 string
		Intermediate      string
	}
}

// OrderParameters represents generic Order Parameters
type OrderParameters struct {
	CSR                    string   // PEM-encoded PKCS#10
	IsCompetitiveUpgrade   bool     `json:",omitempty"`
	IsRenewal              bool     `json:",omitempty"`
	PartnerOrderID         string   `json:",omitempty"`
	ProductCode            string   `json:",omitempty"`
	ServerCount            int      `json:",omitempty"`
	SignatureHashAlgorithm string   `json:",omitempty"`
	SubjectAltNameCount    int      `json:",omitempty"`
	SubjectAltNames        []string `json:",omitempty"`
	ValidityPeriod         int      `json:",omitempty"` // 12 or 24 month (days for AlwaysOnSSL, min. 180, max. 365)
	// AlwaysOnSSL (Symantec Encryption Everywhere) only:
	DVAuthMethod  string `json:",omitempty"` // DNS, EMAIL
	ApproverEmail string `json:",omitempty"`
}

// Contact represents a generic Contact type (for AdminContact and TechContact)
type Contact struct {
	Title               string              `json:",omitempty"`
	FirstName           string              `json:",omitempty"`
	LastName            string              `json:",omitempty"`
	OrganizationName    string              `json:",omitempty"`
	OrganizationAddress OrganizationAddress `json:",omitempty"`
	Phone               string              `json:",omitempty"`
	Fax                 string              `json:",omitempty"`
	Email               string              `json:",omitempty"`
}

// OrganizationAddress holds general information about a organization
type OrganizationAddress struct {
	AddressLine1 string `json:",omitempty"`
	PostalCode   string `json:",omitempty"`
	City         string `json:",omitempty"`
	Region       string `json:",omitempty"`
	Country      string `json:",omitempty"`
	Phone        string `json:",omitempty"`
	Fax          string `json:",omitempty"`
}

// OrderRequest represents a POST /Order request
type OrderRequest struct {
	OrganizationInfo struct {
		OrganizationName    string              `json:",omitempty"`
		OrganizationAddress OrganizationAddress `json:",omitempty"`
	} `json:",omitempty"`
	OrderParameters OrderParameters `json:",omitempty"`
	AdminContact    Contact         `json:",omitempty"`
	TechContact     Contact         `json:",omitempty"`
}
