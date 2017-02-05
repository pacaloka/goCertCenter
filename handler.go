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
	// original: "github.com/google/go-querystring/query"
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
		}
	}

	request, err := http.NewRequest(req.httpMethod, req.url, postData)
	if err != nil {
		return err
	}

	request.Header.Add("Authorization", "Bearer "+Bearer)
	request.Header.Set("Content-Type", "application/json; charset=utf8")

	response, err := req.client.Do(request)
	defer response.Body.Close()

	if response.ContentLength > 1<<16 || response.ContentLength == 0 {
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
		case 401:
			return errors.New(fmt.Errorf("CertCenter API: Autorization failed. Used bearer token is invalid or does not have the proper rights"))
		case 417:
			return errors.New(fmt.Errorf("CertCenter API: Please double check your request data: %s", string(data)))
		default:
			return errors.New(fmt.Errorf("CertCenter API: Returned with Status %d", response.StatusCode))
		}
	}

	if err := json.Unmarshal(data, &req.result); err != nil {
		return err
	}

	return nil
}
