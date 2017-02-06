package certcenter

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/certcenter/goCertCenter/query"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func (req *apiRequest) do(apiMethod string, ParamType ...int) error {

	var postData io.Reader
	paramType := CC_PARAM_TYPE_QS
	req.httpMethod = "GET"
	req.method = apiMethod
	req.url = "https://api.certcenter.com/rest/v1/" + req.method
	req.client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				PreferServerCipherSuites: true,
				MinVersion:               tls.VersionTLS12,
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
				req.url = fmt.Sprintf("%s/%d",
					req.url,
					req.request.(*ResendApproverEmailRequest).CertCenterOrderID)
			}
		case CC_PARAM_TYPE_QS | CC_PARAM_TYPE_PATH:
			if apiMethod == "ApproverEmail" {
				req.httpMethod = "PUT"
				req.url = fmt.Sprintf("%s/%d?ApproverEmail=%s",
					req.url,
					req.request.(*PutApproverEmailRequest).CertCenterOrderID,
					req.request.(*PutApproverEmailRequest).ApproverEmail)
			} else if apiMethod == "Order" {
				v, err := query.Values(req.request)
				if err != nil {
					return err
				}
				req.url += fmt.Sprintf("/%d?", req.request.(*GetOrderRequest).CertCenterOrderID)
				x := v.Encode()
				fmt.Println(x)
				req.url += x
			}
		}
	}

	fmt.Println(req.url)

	request, err := http.NewRequest(req.httpMethod, req.url, postData)
	if err != nil {
		return err
	}

	request.Header.Add("Authorization", "Bearer "+Bearer)
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

	fmt.Println(string(data))

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
