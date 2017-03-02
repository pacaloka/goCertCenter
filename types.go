package certcenter

import (
	"fmt"
	"net/http"
	"time"
)

// Bearer represents the authentication token you're going to use
var Bearer string

// KvStoreAuthorizationKey need to be set if you want to use
// CertCenter's free key-value database, please ask your partner
// manager or our customer support team to send you an
// "AlwaysOnSSL KV-Storage Authorization-Key"
var KvStoreAuthorizationKey string

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

// SchemeValidationErrors provides basic fields for scheme validation errors
type SchemeValidationErrors struct {
	Errors []struct {
		Msg    string `json:"msg"`
		Status string `json:"status"`
		Key    string `json:"key"`
		SchemeValidationErrors
	}
}

// BasicResultInfo represents the default values included in each resultset
type BasicResultInfo struct {
	Success bool `json:"success"`
	Message string
	// if !Success, ErrorId and/or ErrorField may be provided
	ErrorId    int
	ErrorField string
	// Scheme validation results
	Msg string `json:"msg"`
	SchemeValidationErrors
}

// ProfileResult represents a GET /Profile response
type ProfileResult struct {
	AuthType        string
	AuthorizationID int64
	Country         string
	Currency        string
	CustomerID      int64
	Locale          string
	OAuth2Token     string `json:"OAuth2_Token"`
	Scope           string
	Timezone        string
}

// LimitResult represents a GET /Limit response
type LimitResult struct {
	BasicResultInfo
	LimitInfo struct {
		Limit float64
		Used  float64
	}
}

// ProductsResult represents a GET /Products response
type ProductsResult struct {
	BasicResultInfo
	Products []string
}

// ProductDetailsResult represents a GET /ProductDetails response
type ProductDetailsResult struct {
	BasicResultInfo
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
	BasicResultInfo
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
	BasicResultInfo
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
	BasicResultInfo
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
	BasicResultInfo
	ApproverList []struct {
		ApproverEmail string
		ApproverType  string // Domain, Generic
	}
}

// OrderResult represents a POST /Order response
type OrderResult struct {
	BasicResultInfo
	Timestamp         time.Time
	CertCenterOrderID int64
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
		ValidityPeriod         int    // 12 or 24 month (days for AlwaysOnSSL, min. 180, max. 365)
		DVAuthMethod           string // DNS, EMAIL
	}
	// AlwaysOnSSL (Symantec Encryption Everywhere) only:
	Fulfillment struct {
		Certificate  string
		PKCS7        string `json:"Certificate_PKCS7"`
		Intermediate string
	}
}

// OrderParameters represents generic Order Parameters
type OrderParameters struct {
	CSR                    string   `json:",omitempty"` // PEM-encoded PKCS#10
	IsCompetitiveUpgrade   bool     `json:",omitempty"`
	IsRenewal              bool     `json:",omitempty"`
	PartnerOrderID         string   `json:",omitempty"`
	ProductCode            string   `json:",omitempty"`
	ServerCount            int      `json:",omitempty"`
	SignatureHashAlgorithm string   `json:",omitempty"`
	SubjectAltNameCount    int      `json:",omitempty"`
	SubjectAltNames        []string `json:",omitempty"`
	ValidityPeriod         int      `json:",omitempty"` // 12 or 24 month (days for AlwaysOnSSL, min. 180, max. 365)
	DVAuthMethod           string   `json:",omitempty"` // DNS, EMAIL
	ApproverEmail          string   `json:",omitempty"`
}

// OrganizationInfo represents organizational information
type OrganizationInfo struct {
	OrganizationName    string               `json:",omitempty"`
	OrganizationAddress *OrganizationAddress `json:",omitempty"`
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

// Contact represents a generic Contact type (for AdminContact and TechContact)
type Contact struct {
	Title               string               `json:",omitempty"`
	FirstName           string               `json:",omitempty"`
	LastName            string               `json:",omitempty"`
	OrganizationName    string               `json:",omitempty"`
	OrganizationAddress *OrganizationAddress `json:",omitempty"`
	Phone               string               `json:",omitempty"`
	Fax                 string               `json:",omitempty"`
	Email               string               `json:",omitempty"`
}

// OrderRequest represents a POST /Order request
type OrderRequest struct {
	OrganizationInfo *OrganizationInfo `json:",omitempty"`
	OrderParameters  *OrderParameters  `json:",omitempty"`
	AdminContact     *Contact          `json:",omitempty"`
	TechContact      *Contact          `json:",omitempty"`
}

// PutApproverEmailResult represents a PUT /ApproverEmail response
type PutApproverEmailResult struct {
	BasicResultInfo
}

// PutApproverEmailRequest represents a PUT /ApproverEmail request
type PutApproverEmailRequest struct {
	CertCenterOrderID int64
	ApproverEmail     string
}

// ResendApproverEmailResult represents a POST /ApproverEmail response
type ResendApproverEmailResult struct {
	BasicResultInfo
}

// ResendApproverEmailRequest represents a POST /ApproverEmail request
type ResendApproverEmailRequest struct {
	CertCenterOrderID int64
}

// OrderInfo contains all information about a certain order
type OrderInfo struct {
	CertCenterOrderID int64
	CommonName        string
	OrderStatus       struct {
		MajorStatus string
		MinorStatus string
		OrderDate   time.Time
		UpdateDate  time.Time
		StartDate   time.Time
		EndDate     time.Time
		Progress    int
	}
	ConfigurationAssessment struct { // done by ssllabs.com
		Engine          string
		Ranking         string
		Effective       time.Time
		CriteriaVersion string
	}
	BillingInfo struct {
		Price      float32
		Currency   string
		Status     string
		InvoiceRef string // if available (Status == cleared)
	}
	OrderParameters struct {
		PartnerOrderID  string
		ValidityPeriod  int
		ServerCount     int32
		ProductCode     string
		DVAuthMethod    string // DV certificates only
		SubjectAltNames []string
	}
	ContactInfo struct { // if includeContacts
		AdminContact Contact
		TechContact  Contact
	}
	OrganizationInfo struct { // if !DV
		OrganizationName    string
		OrganizationAddress OrganizationAddress
	}
	Fulfillment struct {
		StartDate     time.Time
		EndDate       time.Time
		CSR           string
		Certificate   string
		Intermediate  string
		DownloadLinks struct { // cert.sh download links
			Certificate  string
			Intermediate string
			IconScript   string
			PKCS7        string
		}
	}
	DNSAuthDetails struct { // for DV orders with DNS auth and includeOrderParameters
		DNSEntry string
		DNSValue string
	}
	FileAuthDetails struct { // for DV orders with FILE auth and includeOrderParameters
		FileContents string
		FileName     string
		PollStatus   string
		LastPollDate time.Time
	}
	MetaAuthDetails struct { // for GlobalSign DV orders with META auth and includeOrderParameters
		MetaTag string
	}
	EmailAuthDetails struct { // for DV orders with EMAIL auth and includeOrderParameters
		ApproverEmail       string
		ApproverNotifyDate  time.Time
		ApproverConfirmDate time.Time
	}
}

// GetOrdersResult represents a GET /Orders response
type GetOrdersResult struct {
	BasicResultInfo
	OrderInfos []OrderInfo
	Meta       struct {
		ItemsAvailable int64
		ItemsPerPage   int64
		Page           int64
		OrderBy        string
		OrderDir       string
		Status         string
		ProductType    string
		CommonName     string
	} `json:"_meta"`
}

// GetOrdersRequest represents a GET /Orders request
type GetOrdersRequest struct {
	Status                   string
	ProductType              string
	CommonName               string
	IncludeFulfillment       bool `url:"includeFulfillment"`
	IncludeOrderParameters   bool `url:"includeOrderParameters"`
	IncludeBillingDetails    bool `url:"includeBillingDetails"`
	IncludeContacts          bool `url:"includeContacts"`
	IncludeOrganizationInfos bool `url:"includeOrganizationInfos"`
}

// GetModifiedOrdersResult represents a GET /ModifiedOrders response
type GetModifiedOrdersResult struct {
	OrderInfos []OrderInfo
	BasicResultInfo
}

// GetModifiedOrdersRequest represents a GET /ModifiedOrders request
type GetModifiedOrdersRequest struct {
	FromDate                 time.Time
	ToDate                   time.Time
	IncludeFulfillment       bool `url:"includeFulfillment"`
	IncludeOrderParameters   bool `url:"includeOrderParameters"`
	IncludeBillingDetails    bool `url:"includeBillingDetails"`
	IncludeContacts          bool `url:"includeContacts"`
	IncludeOrganizationInfos bool `url:"includeOrganizationInfos"`
}

// GetOrderResult represents a GET /Order/:CertCenterOrderID response
type GetOrderResult struct {
	BasicResultInfo
	OrderInfo OrderInfo
}

// GetOrderRequest represents a GET /Order/:CertCenterOrderID request
type GetOrderRequest struct {
	CertCenterOrderID        int64
	IncludeFulfillment       bool `url:"includeFulfillment"`
	IncludeOrderParameters   bool `url:"includeOrderParameters"`
	IncludeBillingDetails    bool `url:"includeBillingDetails"`
	IncludeContacts          bool `url:"includeContacts"`
	IncludeOrganizationInfos bool `url:"includeOrganizationInfos"`
}

// DeleteOrderResult represents a DELETE /Order/:CertCenterOrderID response
type DeleteOrderResult struct {
	BasicResultInfo
}

// DeleteOrderRequest represents a DELETE /Order/:CertCenterOrderID request
type DeleteOrderRequest struct {
	CertCenterOrderID int64
}

// ReissueResult represents a POST /Reissue response
type ReissueResult struct {
	BasicResultInfo
}

// ReissueOrderParameters represents the required OrderParameters for POST /Reissue
type ReissueOrderParameters struct {
	CSR                    string
	DVAuthMethod           string // for DV certificates only
	SignatureHashAlgorithm string
}

// ReissueRequest represents a POST /Reissue request
// Description:
// 	https://developers.certcenter.com/reference#reissue
type ReissueRequest struct {
	CertCenterOrderID int64
	OrderParameters   ReissueOrderParameters
	ReissueEmail      string
}

// RevokeResult represents a DELETE /Revoke response
type RevokeResult struct {
	BasicResultInfo
}

// RevokeRequest represents a DELETE /Revoke request
type RevokeRequest struct {
	CertCenterOrderID int64
	RevokeReason      string
	Certificate       string // PEM encoded X.509 certificate
}

// ValidateNameResult represents a POST /ValidateName response
type ValidateNameResult struct {
	BasicResultInfo
	IsQualified bool
	// If ValidateNameRequest contained a GeneratePrivateKey=true
	// this two values are included in the result:
	CSR        string
	PrivateKey string
}

// ValidateNameRequest represents a POST /ValidateName request
// https://developers.certcenter.com/v1/reference#validatename
type ValidateNameRequest struct {
	CommonName         string
	GeneratePrivateKey bool // Generates a private key for you.
	// If true the response will also include the CSR and PrivateKey values.
}

// DNSDataResult represents a POST /DNSData response
type DNSDataResult struct {
	BasicResultInfo
	DNSAuthDetails struct {
		PointerType string // =CNAME
		DNSEntry    string
		DNSValue    string
	}
}

// DNSDataRequest represents a POST /DNSData request
// https://dash.readme.io/project/certcenter/v1/refs/dnsdata
type DNSDataRequest struct {
	ProductCode string
	CSR         string
}

// FileDataResult represents a POST /FileData response
type FileDataResult struct {
	BasicResultInfo
	FileAuthDetails struct {
		FileContents string
		FileName     string
	}
}

// FileDataRequest represents a POST /FileData request
// https://dash.readme.io/project/certcenter/v1/refs/filedata
type FileDataRequest struct {
	ProductCode string
	CSR         string
}

// VulnerabilityAssessmentResult represents a POST /VulnerabilityAssessment response
type VulnerabilityAssessmentResult struct {
	BasicResultInfo
}

// VulnerabilityAssessmentRequest represents a POST /VulnerabilityAssessment request
// https://developers.certcenter.com/v1/reference#vulnerabilityassessment
type VulnerabilityAssessmentRequest struct {
	CertCenterOrderID      int64
	ServiceStatus          string
	EmailNotificationLevel string
}

// VulnerabilityAssessmentRescanResult represents a GET /VulnerabilityAssessment/:CertCenterOrderID response
type VulnerabilityAssessmentRescanResult struct {
	BasicResultInfo
}

// VulnerabilityAssessmentRescanRequest represents a GET /VulnerabilityAssessmen/:CertCenterOrderID request
// https://developers.certcenter.com/v1/reference#vulnerabilityassessmentrescan
type VulnerabilityAssessmentRescanRequest struct {
	CertCenterOrderID int64
}

// UserData represents a basic field-set for /User transactions
type UserData struct {
	UsernameOrUserId string   `json:",omitempty"`
	FullName         string   `json:",omitempty"`
	Email            string   `json:",omitempty"`
	Username         string   `json:",omitempty"`
	Password         string   `json:",omitempty"`
	Roles            []string `json:",omitempty"`
	Mobile           string   `json:",omitempty"`
	Timezone         string   `json:",omitempty"`
	Locale           string   `json:",omitempty"`

	// Available on user data retrieval
	SpecialProductAvailability bool   `json:",omitempty"`
	Scope                      string `json:",omitempty"`
	Active                     bool   `json:",omitempty"`
	TwoFactorEnabled           bool   `json:",omitempty"`
	InsertData                 int64  `json:",omitempty"` // Unix time
	LastUpdateData             int64  `json:",omitempty"` // Unix time
	LastPasswordChangeDate     int64  `json:",omitempty"` // Unix time
}

// CreateUserResult represents a POST /User response
type CreateUserResult struct {
	BasicResultInfo
	Id       int64
	FullName string
	Username string
	Roles    []string
}

// CreateUserRequest represents a POST /User request
// https://developers.certcenter.com/v1/reference#createuser
type CreateUserRequest struct {
	UserData
}

// UpdateUserResult represents a POST /User/:UsernameOrUserId response
type UpdateUserResult struct {
	BasicResultInfo
	Id int64
}

// UpdateUserRequest represents a POST /User/:UsernameOrUserId request
// https://developers.certcenter.com/v1/reference#updateuser
type UpdateUserRequest struct {
	UserData
}

// GetUserResult represents a GET /User/:UsernameOrUserId response
type GetUserResult struct {
	BasicResultInfo
	Id int64
}

// GetUserRequest represents a GET /User/:UsernameOrUserId request
// https://developers.certcenter.com/v1/reference#getuser
type GetUserRequest struct {
	UserData
}

// DeleteUserResult represents a DELETE /User/:UsernameOrUserId response
type DeleteUserResult struct {
	BasicResultInfo
	Id int64
}

// DeleteUserRequest represents a GET /User/:UsernameOrUserId request
// https://developers.certcenter.com/v1/reference#deleteuser
type DeleteUserRequest struct {
	UsernameOrUserId string
}

// KeyValueStoreResult represents a basic kv-storage response
type KeyValueStoreResult struct {
	Message string `json:"message"`
}

// KeyValueStoreRequest represents a basic kv-storage request
type KeyValueStoreRequest struct {
	Key   string `json:"filename,omitempty"`
	Value string `json:"hash"`
}

// CreateVoucherResult represents a POST /Voucher response
type CreateVoucherResult struct {
	BasicResultInfo
	VoucherCode     string
	OrderParameters OrderParameters
}

// CreateVoucherRequest represents a POST /Voucher request
// https://developers.certcenter.com/v1/reference#createvoucher
type CreateVoucherRequest struct {
	OrderParameters OrderParameters
}

// RedeemVoucherResult represents a POST /Redeem response
type RedeemVoucherResult struct {
	OrderResult
}

// RedeemVoucherRequest represents a POST /Redeem request
// https://developers.certcenter.com/v1/reference#redeemvoucher
type RedeemVoucherRequest struct {
	VoucherCode      string
	OrganizationInfo *OrganizationInfo `json:",omitempty"`
	OrderParameters  *OrderParameters  `json:",omitempty"`
	AdminContact     *Contact          `json:",omitempty"`
	TechContact      *Contact          `json:",omitempty"`
}

// GetVouchersResult represents a GET /Vouchers and a GET /Voucher/:VoucherCode response
// https://developers.certcenter.com/v1/reference#getvouchers
// https://developers.certcenter.com/v1/reference#getvoucher
type GetVouchersResult struct {
	BasicResultInfo
	Vouchers []struct {
		RedeemInfo struct {
			RedeemDate        time.Time
			CertCenterOrderID int64
		}
		CreationDate    time.Time
		OrderParameters OrderParameters
		VoucherCode     string
		Redeemed        bool
	}
}

// GetVoucherRequest represents a GET /Voucher/:VoucherCode request
// https://developers.certcenter.com/v1/reference#getvoucher
type GetVoucherRequest struct {
	VoucherCode string
}

// DeleteVoucherResult represents a DELETE /Voucher/:VoucherCode response
type DeleteVoucherResult struct {
	OrderResult
}

// DeleteVoucherRequest represents a DELETE /Voucher/:VoucherCode request
// https://developers.certcenter.com/v1/reference#deletevoucher
type DeleteVoucherRequest struct {
	VoucherCode string
}
