[![Go Report Card](https://goreportcard.com/badge/github.com/CertCenter/goCertCenter)](https://goreportcard.com/report/github.com/CertCenter/goCertCenter)

It's quite easy to use CertCenter's API in your Go projects:

```go
package main

import (
	"fmt"
	"io/ioutil"
	certcenter "certcenter.com/go"
)

/* Set your Authorization Token
 */
func init() {
	certcenter.Bearer = "aValidToken.oauth2.certcenter.com"
}

func main() {

	// Get a Quote
	res, _ := certcenter.Quote(&certcenter.QuoteRequest{
		ProductCode: "Symantec.SecureSiteEV",
		SubjectAltNameCount: 0,
		ValidityPeriod: 24,
		ServerCount: 1,
	})
	fmt.Println(res)

	// Validate a CSR (PEM-encoded PKCS#10)
	csr, _ := ioutil.ReadFile("csr")
	res, _ := certcenter.ValidateCSR(&certcenter.ValidateCSRRequest{CSR: string(csr)})
	fmt.Println(res)

	return
}
```

Find more examples and detailed information:
https://developers.certcenter.com/v1/reference

Have fun!
