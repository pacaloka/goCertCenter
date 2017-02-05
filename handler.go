package certcenter

import (
	"fmt"
	"errors"
	"net/http"
	"crypto/tls"
	"io/ioutil"
	"encoding/json"
	"github.com/certcenter/goCertCenter/query"
	// original: "github.com/google/go-querystring/query"
)

func (req *apiRequest)do(apiMethod string, ParamType ...int) (error) {

	req.method=apiMethod
	req.client=&http.Client{
			Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
							PreferServerCipherSuites: true,
							MinVersion: tls.VersionTLS12,
					},
			},
	}

	url := "https://api.certcenter.com/rest/v1/"+req.method
	if len(ParamType)>0 {
		if ParamType[0] == CC_PARAM_TYPE_QS {
			v, _ := query.Values(req.request)
			url += "?"+v.Encode()
			fmt.Println(url)
		}
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	request.Header.Add("Authorization", "Bearer "+Bearer)
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

	if err := json.Unmarshal(data, &req.result); err != nil {
		return err
	}

	return nil
}