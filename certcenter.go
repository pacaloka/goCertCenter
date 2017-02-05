package certcenter

// Profile fetches basic informations about your profile
func Profile() (*ProfileResult, error) {
	req := new(apiRequest)
	req.result = new(ProfileResult)
	err := req.do("Profile")
	checkErr(err)
	return req.result.(*ProfileResult), err
}

// Limit inquires information about your current limit and used amount
func Limit() (*LimitResult, error) {
	req := new(apiRequest)
	req.result = new(LimitResult)
	err := req.do("Limit")
	checkErr(err)
	return req.result.(*LimitResult), err
}

// Products allows you to fetch a list of valid ProductCodes 
func Products() (*ProductsResult, error) {
	req := new(apiRequest)
	req.result = new(ProductsResult)
	err := req.do("Products")
	checkErr(err)
	return req.result.(*ProductsResult), err
}

// ProductDetails inquires detailed information on a particular ProductCode
func ProductDetails(ProductCode string) (*ProductDetailsResult, error) {
	req := new(apiRequest)
	req.result = new(ProductDetailsResult)
	req.request = &ProductDetailsRequest{
		ProductCode: ProductCode,
	}
	err := req.do("ProductDetails", CC_PARAM_TYPE_QS)
	checkErr(err)
	return req.result.(*ProductDetailsResult), err
}

// Quote allows you to generate an individual real-time quotation
func Quote(request *QuoteRequest) (*QuoteResult, error) {
	req := new(apiRequest)
	req.result = new(QuoteResult)
	req.request = request
	err := req.do("Quote", CC_PARAM_TYPE_QS)
	checkErr(err)
	return req.result.(*QuoteResult), err
}

// ValidateCSR allows you to parse and validate a PEM-encoded PKCS#10
func ValidateCSR(request *ValidateCSRRequest) (*ValidateCSRResult, error) {
	req := new(apiRequest)
	req.result = new(ValidateCSRResult)
	req.request = request
	err := req.do("ValidateCSR", CC_PARAM_TYPE_BODY)
	checkErr(err)
	return req.result.(*ValidateCSRResult), err
}

// UserAgreement fetches the latest subscriber agreement from the CA
func UserAgreement(ProductCode string) (*UserAgreementResult, error) {
	req := new(apiRequest)
	req.result = new(UserAgreementResult)
	req.request = &UserAgreementRequest{
		ProductCode: ProductCode,
	}
	err := req.do("UserAgreement", CC_PARAM_TYPE_QS)
	checkErr(err)
	return req.result.(*UserAgreementResult), err
}

// ApproverList will fetch a list of valid email addresses for a particular CommonName and ProductCode
func ApproverList(request *ApproverListRequest) (*ApproverListResult, error) {
	req := new(apiRequest)
	req.result = new(ApproverListResult)
	req.request = request
	err := req.do("ApproverList", CC_PARAM_TYPE_QS)
	checkErr(err)
	return req.result.(*ApproverListResult), err
}

// Order allows you to submit orders for regular certificates as well as S/MIME and AlwaysOnSSL certificates
func Order(request *OrderRequest) (*OrderResult, error) {
	req := new(apiRequest)
	req.result = new(OrderResult)
	req.request = request
	err := req.do("Order", CC_PARAM_TYPE_BODY)
	checkErr(err)
	return req.result.(*OrderResult), err
}
