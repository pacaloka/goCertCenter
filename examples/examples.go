package main

import (
	"fmt"
	//"io/ioutil"
	//"time"
	certcenter "certcenter.com/go"
)

// Set your valid OAuth2 Bearer
// (see https://developers.certcenter.com/docs/authentication)
func init() {
	certcenter.Bearer = "AValidToken.oauth2.certcenter.com"
}

func main() {
	/*
			//////////////////////////////////////////////////////

			// Get my profile information
			res, _ := certcenter.Profile()
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Inquire limit informations
			res, _ := certcenter.Limit()
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Get all valid ProductCodes
			res, _ := certcenter.Products()
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Inquire details about a product
			res, _ := certcenter.ProductDetails("GeoTrust.QuickSSLPremium")
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Get a Quote
			res, _ := certcenter.Quote(&certcenter.QuoteRequest{
				ProductCode: "Symantec.SecureSiteEV",
				SubjectAltNameCount: 0,
				ValidityPeriod: 24,
				ServerCount: 1,
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Validate a CSR
			csr, _ := ioutil.ReadFile("csr")
			res, _ := certcenter.ValidateCSR(&certcenter.ValidateCSRRequest{
				CSR: string(csr),
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Get the CA's User Agreement
			res, _ := certcenter.UserAgreement("Symantec.SecureSite")
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Get valid email addresses for email based DV validation
			res, _ := certcenter.ApproverList(&certcenter.ApproverListRequest{
				CommonName: "www.example.com",
				ProductCode: "GeoTrust.QuickSSLPremium",
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Order a certificate
			csr, _ := ioutil.ReadFile("csr")
			res, _ := certcenter.Order(&certcenter.OrderRequest{
					OrderParameters: &certcenter.OrderParameters{
						ProductCode: "RapidSSL.RapidSSL",
						CSR: string(csr),
						ValidityPeriod: 24,
						ApproverEmail:"domains@certcenter.com",
					},
					AdminContact: &certcenter.Contact{
						FirstName: "John",
						LastName: "Doe",
						Phone: "+1 212 999 999",
						Email: "john.doe@example.com",
					},
					TechContact: &certcenter.Contact{
						FirstName: "John",
						LastName: "Doe",
						Phone: "+1 212 999 999",
						Email: "john.doe@example.com",
					},
				},
			)
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Re-set the approvers email address
			res, _ := certcenter.PutApproverEmail(&certcenter.PutApproverEmailRequest{
				CertCenterOrderID: 123456789,
				ApproverEmail: "valid-approver@example.com",
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Re-send the approver email to the approvers address
			res, _ := certcenter.ResendApproverEmail(&certcenter.ResendApproverEmailRequest{
				CertCenterOrderID: 123456789,
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Get filtered orders
			res, _ := certcenter.GetOrders(&certcenter.GetOrdersRequest{
				Status: "COMPLETE", // COMPLETE, PENDING, CANCELLED, REVOKED
				ProductType: "SSL", // SSL, CODESIGN, SMIME, TRUSTSEAL
				CommonName: "%",
				IncludeFulfillment: true,
				IncludeOrderParameters: true,
				IncludeBillingDetails: true,
				IncludeContacts: true,
				IncludeOrganizationInfos: true,
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Get all modfied orders in timespan
			res, _ := certcenter.GetModifiedOrders(&certcenter.GetModifiedOrdersRequest{
				FromDate: time.Now().Add(-10 * time.Minutes),
				ToDate: time.Now(),
				IncludeFulfillment: true,
				IncludeOrderParameters: true,
				IncludeBillingDetails: true,
				IncludeContacts: true,
				IncludeOrganizationInfos: true,
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Get a particular order by CertCenterOrderID
			res, _ := certcenter.GetOrder(&certcenter.GetOrderRequest{
				CertCenterOrderID: 123456789,
				IncludeFulfillment: true,
				IncludeOrderParameters: true,
				IncludeBillingDetails: true,
				IncludeContacts: true,
				IncludeOrganizationInfos: true,
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Delete a particular order
			res, _ := certcenter.DeleteOrder(&certcenter.DeleteOrderRequest{
				CertCenterOrderID: 123456789,
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Reissue a particular order
			res, _ := certcenter.Reissue(&certcenter.ReissueRequest{
				CertCenterOrderID: 123456789,
				OrderParameters:certcenter.ReissueOrderParameters{
					CSR: "#CSR#",
					DVAuthMethod: "EMAIL",
					SignatureHashAlgorithm: "SHA256-FULL-CHAIN",
				},
				ReissueEmail: "valid-approver@example.com",
			})
			fmt.Println(res)


			//////////////////////////////////////////////////////

			// Revoke a certificate
			res, _ := certcenter.Revoke(&certcenter.RevokeRequest{
				CertCenterOrderID: 123456789,
				RevokeReason: "Key compromised",
				Certificate: "#PEM-encoded-X.509-Certificate#",
			})
			fmt.Println(res)


			//////////////////////////////////////////////////////

			// Check a CommonName against the black list (AlwaysOnSSL only!)
			// plus lets you generate a private key and PEM-encoded CSR
			//
			res, _ := certcenter.ValidateName(&certcenter.ValidateNameRequest{
				CommonName: "www.example.com",
				GeneratePrivateKey: true,
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Retrieve appropriate data for DNS-based validation (AlwaysOnSSL only!)
			//
			csr, _ := ioutil.ReadFile("csr")
			res, _ := certcenter.DNSData(&certcenter.DNSDataRequest{
				CSR: string(csr),
				ProductCode: "AlwaysOnSSL.AlwaysOnSSL",
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// Retrieve appropriate data for FILE-based validation (AlwaysOnSSL only!)
			//
			csr, _ := ioutil.ReadFile("csr")
			res, _ := certcenter.FileData(&certcenter.FileDataRequest{
				CSR: string(csr),
				ProductCode: "AlwaysOnSSL.AlwaysOnSSL",
			})
			fmt.Println(res)


			// VulnerabilityAssessment allows you to change an orders initial assessment settings.
			// https://developers.certcenter.com/v1/reference#vulnerabilityassessment
			//
			res, _ := certcenter.VulnerabilityAssessment(&certcenter.VulnerabilityAssessmentRequest{
				CertCenterOrderID: 123456789,
				ServiceStatus: "Active",
				EmailNotificationLevel: "CRITICAL",
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// VulnerabilityAssessmentRescan allows you to initiate a immediate re-assessment
			// https://developers.certcenter.com/v1/reference#vulnerabilityassessmentrescan
			//
			res, _ := certcenter.VulnerabilityAssessmentRescan(&certcenter.VulnerabilityAssessmentRescanRequest{
				CertCenterOrderID: 123456789,
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// CreateUser allows you to create a new UI user for your organizations account
			// https://developers.certcenter.com/v1/reference#createuser
			//
			res, _ := certcenter.CreateUser(&certcenter.CreateUserRequest{
				certcenter.UserData{
					FullName: "John Doe",
					Email: "john@example.org",
					Username: "johndoes",
					Password: "cOmpL3xx/",
					Mobile: "",
					Roles: []string{"PROCUREMENT"},
					Locale: "en_US",
					Timezone: "US/Pacific",
				},
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// UpdateUser allows you to update an user
			// https://developers.certcenter.com/v1/reference#updateuser
			//
			res, _ := certcenter.UpdateUser(&certcenter.UpdateUserRequest{
			certcenter.UserData{
					UsernameOrUserId: "1234567",
					FullName: "John Doe",
					Email: "john@example.org",
				}})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// GetUser allows you to inquire information about one or all users
			// https://developers.certcenter.com/v1/reference#getuser
			//
			res, _ := certcenter.GetUser(&certcenter.GetUserRequest{
			certcenter.UserData{
					UsernameOrUserId: "1234567",
				}})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// DeleteUser allows you to remove a specific user
			// https://developers.certcenter.com/v1/reference#deleteuser
			//
			res, _ := certcenter.DeleteUser(&certcenter.DeleteUserRequest{
				UsernameOrUserId: "1234567",
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// KvStore allows you to use CertCenter's free key-value storage
			// for DV FILE authentication with mod_fauth. Details:
			// https://developers.certcenter.com/v1/docs/howto-order-alwaysonssl-symantec-ee-certificates#section-4-order-procedure
			//
			certcenter.KvStoreAuthorizationKey = ""
			res, err := certcenter.KvStore(&certcenter.KeyValueStoreRequest{
					Key: "test.example.com",
					Value: "201701260800495t3djr2zqhqfvgg1cpjmgs5zx4kd7w51w3cuge90sokdavg6li",
			}); if err!=nil {
			  panic("..")
			}

			//////////////////////////////////////////////////////

			// CreateVoucher creates a coupon code which can later be redeemded.
			// https://developers.certcenter.com/v1/reference#createvoucher
			//
			res, _ := certcenter.CreateVoucher(&certcenter.CreateVoucherRequest{
				OrderParameters:certcenter.OrderParameters{
					ProductCode: "Thawte.SSL123",
					PartnerOrderID: "My voucher order id (optional)",
					ServerCount: 1,
					SubjectAltNameCount: 0,
					ValidityPeriod: 12,
				},
			})
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// RedeemVoucher let you redeem a previously generated voucher code
			// https://developers.certcenter.com/v1/reference#redeemvoucher
			//
			csr, _ := ioutil.ReadFile("csr")
			fmt.Println(string(csr))
			res, _ := certcenter.RedeemVoucher(&certcenter.RedeemVoucherRequest{
					VoucherCode: "JDX1UBDC6AA1",
					// You don't need OrganizationInfo on DV orders, except for SSL123.
					OrganizationInfo: &certcenter.OrganizationInfo{
						OrganizationName: "Acme LCC",
						OrganizationAddress: &certcenter.OrganizationAddress{
							AddressLine1: "40 5th Ave",
							Region: "NY",
							PostalCode: "12012",
							Country: "US",
							City: "NY",
							Phone: "+1 121 444444",
						},
					},
					OrderParameters: &certcenter.OrderParameters{
						ProductCode: "Thawte.SSL123",
						CSR: string(csr),
						ValidityPeriod: 12,
						DVAuthMethod: "EMAIL",
						ServerCount: 1,
						// Needs to be a valid approver email address.
						// Inquire valid addresses via /ApproverEmail
						ApproverEmail:"postmaster@example.com",
				    SignatureHashAlgorithm: "SHA256-FULL-CHAIN",
					},
					AdminContact: &certcenter.Contact{
						FirstName: "John",
						LastName: "Doe",
						Title: "CEO",
						Phone: "+1 212 999 999",
						Email: "john.doe@example.com",
					},
					TechContact: &certcenter.Contact{
						FirstName: "John",
						LastName: "Doe",
						Title: "CEO",
						Phone: "+1 212 999 999",
						Email: "john.doe@example.com",
					},
				},
			)
			fmt.Println(res)

			//////////////////////////////////////////////////////

			// GetVouchers allows you to inquire information about all your vouchers
			// https://developers.certcenter.com/v1/reference#getvouchers
			//
			res, _ := certcenter.GetVouchers()
			fmt.Println(res.Vouchers[0].OrderParameters)

			//////////////////////////////////////////////////////

			// GetVoucher allows you to inquire information about particular voucher
			// https://developers.certcenter.com/v1/reference#getvoucher
			//
			res, _ := certcenter.GetVoucher(&certcenter.GetVoucherRequest{VoucherCode: "JDX1UBHC6UA3"})
			fmt.Println(res.Vouchers[0].OrderParameters)

			//////////////////////////////////////////////////////

			// DeleteVoucher allows you to invalidate a particular voucher
			// https://developers.certcenter.com/v1/reference#deletevoucher
			//
		  res, _ := certcenter.DeleteVoucher(&certcenter.DeleteVoucherRequest{VoucherCode: "JDX1UBHC6UA3"})
		  fmt.Println(res)

			//////////////////////////////////////////////////////

			// GetVoucherAnonymously allows a 3rd party to get information about a
			// particular VoucherCode w/o authentication
			// https://developers.certcenter.com/v1/reference#getvoucheranonymously
			//
			res, _ := certcenter.GetVoucherAnonymously(&certcenter.GetVoucherRequest{VoucherCode: "JDX1UBHC6UA3"})
			fmt.Println(res.Vouchers[0].OrderParameters)

			//////////////////////////////////////////////////////

			// GetVoucherOrderAnonymously allows a 3rd party to get information about a
			// particular Order (based on a voucher) w/o authentication
			// https://developers.certcenter.com/v1/reference#getvoucherorderanonymously
			//
			res, _ := certcenter.GetVoucherOrderAnonymously(&certcenter.GetVoucherRequest{VoucherCode: "JDX1UBHC6UA3"})
			fmt.Println(res.Vouchers[0].OrderParameters)

	*/

	fmt.Println("")
	return
}
