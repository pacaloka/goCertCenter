It's quite easy to use CertCenter's API in your Go projects:

```go
package main

import (
	"fmt"
	"io/ioutil"
	certcenter "github.com/CertCenter/goCertCenter"
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

There are more examples available in test/test.go.

Have fun!
