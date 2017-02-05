package certcenter

import (
	"io"
	"fmt"
	"errors"
	"strings"
	"net/http"
	"crypto/tls"
	"io/ioutil"
	"encoding/json"
	"github.com/certcenter/goCertCenter/query"
	// original: "github.com/google/go-querystring/query"
)

func (req *apiRequest)do(apiMethod string, ParamType ...int) (error) {

	var postData io.Reader
	paramType := CC_PARAM_TYPE_QS
	req.httpMethod = "GET"
	req.method=apiMethod
	req.url = "https://api.certcenter.com/rest/v1/"+req.method
	req.client=&http.Client{
			Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
							PreferServerCipherSuites: true,
							MinVersion: tls.VersionTLS12,
					},
			},
	}

	if len(ParamType)>0 {
		paramType = ParamType[0]
		switch(paramType) {
			case CC_PARAM_TYPE_QS:
				v, err := query.Values(req.request); if err!=nil {
					return err
				}
				req.url += "?"+v.Encode()
			case CC_PARAM_TYPE_BODY:
				req.httpMethod = "POST"
				d, err := json.Marshal(req.request); if err!=nil {
					return err
				}
				postData = strings.NewReader(string(d))
		}
	}

	request, err := http.NewRequest(req.httpMethod, req.url, postData); if err != nil {
		return err
	}

	request.Header.Add("Authorization", "Bearer "+Bearer)
	request.Header.Set("Content-Type","application/json; charset=utf8")

	response, err := req.client.Do(request)
	defer response.Body.Close()

	req.statusCode=response.StatusCode
	if response.StatusCode != 200 {
		switch(response.StatusCode) {
			case 401:
				return errors.New(fmt.Sprintf("CertCenter API: Autorization failed. Used bearer token is invalid or does not have the proper rights"))
			default:
				return errors.New(fmt.Sprintf("CertCenter API: Returned with Status %d", response.StatusCode))
		}
	}

	if response.ContentLength>1<<16||response.ContentLength==0 {
		return errors.New("CertCenter API: Returned content with wired length")
	}

	data, err := ioutil.ReadAll(response.Body); if err != nil { 
		return err

	}
	fmt.Println(string(data))

	if err := json.Unmarshal(data, &req.result); err != nil {
		return err
	}

	return nil
}