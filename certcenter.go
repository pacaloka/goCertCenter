package certcenter

func Profile() (*ProfileResult, error) {
	req := new(apiRequest)
	req.result = new(ProfileResult)
	err:=req.do("Profile")
	checkErr(err)
	return req.result.(*ProfileResult), err
}

func Limit() (*LimitResult, error) {
	req := new(apiRequest)
	req.result = new(LimitResult)
	err:=req.do("Limit")
	checkErr(err)
	return req.result.(*LimitResult), err
}

func Products() (*ProductsResult, error) {
	req := new(apiRequest)
	req.result = new(ProductsResult)
	err:=req.do("Products")
	checkErr(err)
	return req.result.(*ProductsResult), err
}

func ProductDetails(ProductCode string) (*ProductDetailsResult, error) {
	req := new(apiRequest)
	req.result = new(ProductDetailsResult)
	req.request = new(ProductDetailsRequest)
	req.request = &ProductDetailsRequest{
		ProductCode: ProductCode,
	}
	err:=req.do("ProductDetails", CC_PARAM_TYPE_QS)
	checkErr(err)
	return req.result.(*ProductDetailsResult), err
}

func Quote(request *QuoteRequest) (*QuoteResult, error) {
	req := new(apiRequest)
	req.result = new(QuoteResult)
	req.request = request
	err:=req.do("Quote", CC_PARAM_TYPE_QS)
	checkErr(err)
	return req.result.(*QuoteResult), err
}




