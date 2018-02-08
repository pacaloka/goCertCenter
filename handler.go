package certcenter

import (
	"certcenter.com/go/query"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// do is the central API communication handler
//
func (req *apiRequest) do(apiMethod string, ParamType ...int) error {

	var postData io.Reader
	paramType := CC_PARAM_TYPE_QS
	anonymously := false
	req.httpMethod = "GET"
	req.method = apiMethod
	rawURL := "https://api.certcenter.com/rest/v1/"
	req.url = rawURL + req.method
	req.client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				PreferServerCipherSuites: true,
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	if len(ParamType) > 0 {
		paramType = ParamType[0]
		switch paramType {
		case CC_PARAM_TYPE_QS:
			v, err := query.Values(req.request)
			if err != nil {
				return err
			}
			req.url += "?" + v.Encode()
		case CC_PARAM_TYPE_BODY:
			req.httpMethod = "POST"
			d, err := json.Marshal(req.request)
			if err != nil {
				return err
			}
			postData = strings.NewReader(string(d))
		case CC_PARAM_TYPE_PATH:
			if apiMethod == "ApproverEmail" {
				req.httpMethod = "POST"
				req.url = fmt.Sprintf("%s/%d", req.url,
					req.request.(*ResendApproverEmailRequest).CertCenterOrderID)
			} else if apiMethod == "Order" {
				req.httpMethod = "DELETE"
				req.url = fmt.Sprintf("%s/%d", req.url,
					req.request.(*DeleteOrderRequest).CertCenterOrderID)
			} else if apiMethod == "VulnerabilityAssessment" {
				req.url = fmt.Sprintf("%s/%d", req.url,
					req.request.(*VulnerabilityAssessmentRequest).CertCenterOrderID)
			} else if apiMethod == "User" {
				req.url = fmt.Sprintf("%s/%s", req.url, req.request.(*GetUserRequest).UsernameOrUserId)
			} else if apiMethod == "DeleteUser" {
				apiMethod = "User"
				req.httpMethod = "DELETE"
				fmt.Println(req.request.(*DeleteUserRequest).UsernameOrUserId)
				req.url = fmt.Sprintf("%sUser/%s", rawURL, req.request.(*DeleteUserRequest).UsernameOrUserId)
			} else if apiMethod == "GetVoucher" {
				apiMethod = "Voucher"
				req.url = fmt.Sprintf("%sVoucher/%s", rawURL, req.request.(*GetVoucherRequest).VoucherCode)
			} else if apiMethod == "GetVoucherAnonymously" {
				apiMethod = "Voucher"
				req.url = fmt.Sprintf("%sVoucher/*/%s", rawURL, req.request.(*GetVoucherRequest).VoucherCode)
			} else if apiMethod == "DeleteVoucher" {
				apiMethod = "Voucher"
				req.httpMethod = "DELETE"
				req.url = fmt.Sprintf("%sVoucher/%s", rawURL, req.request.(*DeleteVoucherRequest).VoucherCode)
			} else if apiMethod == "GetVoucherOrderAnonymously" {
				apiMethod = "Order"
				req.url = fmt.Sprintf("%sOrder/*/%s", req.url, req.request.(*GetVoucherRequest).VoucherCode)
			}
		case CC_PARAM_TYPE_QS | CC_PARAM_TYPE_PATH:
			if apiMethod == "ApproverEmail" {
				req.httpMethod = "PUT"
				req.url = fmt.Sprintf("%s/%d?ApproverEmail=%s", req.url,
					req.request.(*PutApproverEmailRequest).CertCenterOrderID,
					req.request.(*PutApproverEmailRequest).ApproverEmail)
			} else if apiMethod == "Order" {
				v, err := query.Values(req.request)
				if err != nil {
					return err
				}
				req.url += fmt.Sprintf("/%d?", req.request.(*GetOrderRequest).CertCenterOrderID)
				x := v.Encode()
				req.url += x
			}
		case CC_PARAM_TYPE_BODY | CC_PARAM_TYPE_PATH:
			if apiMethod == "Revoke" {
				req.httpMethod = "DELETE"
				req.url = fmt.Sprintf("%s/%d", req.url,
					req.request.(*RevokeRequest).CertCenterOrderID)
			} else if apiMethod == "User" {
				req.httpMethod = "POST"
				req.url = fmt.Sprintf("%s/%s", req.url, req.request.(*UpdateUserRequest).UsernameOrUserId)
				req.request.(*UpdateUserRequest).UsernameOrUserId = ""
			}
			d, err := json.Marshal(req.request)
			if err != nil {
				return err
			}
			postData = strings.NewReader(string(d))
		}
	}

	request, err := http.NewRequest(req.httpMethod, req.url, postData)
	if err != nil {
		return err
	}

	if !anonymously {
		request.Header.Add("Authorization", "Bearer "+Bearer)
	}
	request.Header.Set("Content-Type", "application/json; charset=utf8")

	response, err := req.client.Do(request)
	defer response.Body.Close()

	if response.ContentLength > 1<<24 || response.ContentLength == 0 {
		return errors.New("CertCenter API: Returned content with wired length")
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	req.statusCode = response.StatusCode
	if response.StatusCode != 200 {
		switch response.StatusCode {
		default:
			return fmt.Errorf("CertCenter API: Returned with Status %d", response.StatusCode)
		case 401:
			return fmt.Errorf("CertCenter API: Autorization failed. Used bearer token is invalid or does not have the proper rights")
		case 417: // Invalid Request Data
		case 406: // No Changes Made
		}
	}

	if err := json.Unmarshal(data, &req.result); err != nil {
		return err
	}

	return nil
}

// kv allows you to use CertCenter's free key-value storage in conjunction
// with AlwaysOnSSL (aka DigiCert Encryption Everywhere) certificates as
// described at https://developers.certcenter.com/docs/tutorial-integrate-alwaysonssl
//
func (req *apiRequest) kv() error {

	if KvStoreAuthorizationKey == "" {
		return errors.New("KvStoreAuthorizationKey not set. See https://developers.certcenter.com/v1/docs/file-validation-mod-fauth for more details.")
	}

	req.url = "https://fauth-db.eu.certcenter.com/" + req.request.(*KeyValueStoreRequest).Key
	req.request.(*KeyValueStoreRequest).Key = ""
	req.client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				PreferServerCipherSuites: true,
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	d, err := json.Marshal(req.request)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", req.url, strings.NewReader(string(d)))
	if err != nil {
		return err
	}

	request.Header.Add("x-api-key", KvStoreAuthorizationKey)
	request.Header.Add("Content-Type", "application/json")

	response, err := req.client.Do(request)
	defer response.Body.Close()

	if response.ContentLength > 1<<12 || response.ContentLength == 0 {
		return errors.New("CertCenter KV-API: Returned content with wired length")
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	req.statusCode = response.StatusCode
	if response.StatusCode != 200 {
		return errors.New("CertCenter KV-API: Invalid status code")
	}

	if err := json.Unmarshal(data, &req.result); err != nil {
		return err
	}

	return nil
}
