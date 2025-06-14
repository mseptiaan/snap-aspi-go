package types

// BCA-specific types for requests and responses

// BCACreateVAPayload represents the payload for creating a BCA Virtual Account
type BCACreateVAPayload struct {
	PartnerServiceId    string          `json:"partnerServiceId"`
	CustomerNo          string          `json:"customerNo"`
	VirtualAccountNo    string          `json:"virtualAccountNo"`
	VirtualAccountName  string          `json:"virtualAccountName"`
	VirtualAccountEmail string          `json:"virtualAccountEmail"`
	VirtualAccountPhone string          `json:"virtualAccountPhone"`
	TrxId               string          `json:"trxId"`
	TotalAmount         *Amount         `json:"totalAmount"`
	ExpiredDate         string          `json:"expiredDate"`
	AdditionalInfo      *AdditionalInfo `json:"additionalInfo"`
	// BCA-specific fields
	BcaCustomerNo      string `json:"bcaCustomerNo,omitempty"`
	BcaSubCompany      string `json:"bcaSubCompany,omitempty"`
	BcaTransactionType string `json:"bcaTransactionType,omitempty"`
}

// BCACreateVAResponse represents the response from creating a BCA Virtual Account
type BCACreateVAResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ReferenceNo     string `json:"referenceNo"`
	VirtualAccountNo string `json:"virtualAccountNo"`
	// BCA-specific fields
	BcaReferenceNo  string `json:"bcaReferenceNo,omitempty"`
}

// BCAInquiryVAPayload represents the payload for inquiring a BCA Virtual Account
type BCAInquiryVAPayload struct {
	PartnerServiceId string `json:"partnerServiceId"`
	CustomerNo       string `json:"customerNo"`
	VirtualAccountNo string `json:"virtualAccountNo"`
	TrxDateInit      string `json:"trxDateInit"`
	// BCA-specific fields
	BcaCustomerNo string `json:"bcaCustomerNo,omitempty"`
}

// BCAInquiryVAResponse represents the response from inquiring a BCA Virtual Account
type BCAInquiryVAResponse struct {
	ResponseCode       string                 `json:"responseCode"`
	ResponseMessage    string                 `json:"responseMessage"`
	VirtualAccountNo   string                 `json:"virtualAccountNo"`
	VirtualAccountName string                 `json:"virtualAccountName"`
	TotalAmount        map[string]interface{} `json:"totalAmount"`
	ExpiredDate        string                 `json:"expiredDate"`
	// BCA-specific fields
	BcaReferenceNo     string `json:"bcaReferenceNo,omitempty"`
	BcaCustomerNo      string `json:"bcaCustomerNo,omitempty"`
	BcaSubCompany      string `json:"bcaSubCompany,omitempty"`
	BcaTransactionType string `json:"bcaTransactionType,omitempty"`
}

// BCABalanceInquiryPayload represents the payload for BCA balance inquiry
type BCABalanceInquiryPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	AccountNo        string          `json:"accountNo"`
	BalanceTypes     []string        `json:"balanceTypes"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
	// BCA-specific fields
	BcaCustomerNo string `json:"bcaCustomerNo,omitempty"`
}

// BCABalanceInquiryResponse represents the response from BCA balance inquiry
type BCABalanceInquiryResponse struct {
	ResponseCode    string        `json:"responseCode"`
	ResponseMessage string        `json:"responseMessage"`
	ReferenceNo     string        `json:"referenceNo"`
	AccountNo       string        `json:"accountNo"`
	Name            string        `json:"name"`
	AccountInfos    []AccountInfo `json:"accountInfos"`
	// BCA-specific fields
	BcaAccountType string `json:"bcaAccountType,omitempty"`
}

// BCAGenerateQRPayload represents the payload for generating a BCA QR code
type BCAGenerateQRPayload struct {
	PartnerServiceId   string          `json:"partnerServiceId"`
	CustomerNo         string          `json:"customerNo"`
	PartnerReferenceNo string          `json:"partnerReferenceNo"`
	Amount             *Amount         `json:"amount"`
	MerchantId         string          `json:"merchantId"`
	MerchantName       string          `json:"merchantName"`
	ValidityPeriod     string          `json:"validityPeriod"`
	AdditionalInfo     *AdditionalInfo `json:"additionalInfo"`
	// BCA-specific fields
	BcaQRType        string `json:"bcaQRType,omitempty"`
	BcaMerchantCity  string `json:"bcaMerchantCity,omitempty"`
	BcaTerminalLabel string `json:"bcaTerminalLabel,omitempty"`
}

// BCAQRResponse represents the response from generating a BCA QR code
type BCAQRResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ReferenceNo     string `json:"referenceNo"`
	QRContent       string `json:"qrContent"`
	QRUrl           string `json:"qrUrl"`
	QRImage         string `json:"qrImage,omitempty"`
	ValidUntil      string `json:"validUntil"`
	// BCA-specific fields
	BcaQRId string `json:"bcaQRId,omitempty"`
}

// BCATransferPayload represents the payload for BCA transfer
type BCATransferPayload struct {
	PartnerServiceId       string          `json:"partnerServiceId"`
	CustomerNo             string          `json:"customerNo"`
	PartnerReferenceNo     string          `json:"partnerReferenceNo"`
	Amount                 *Amount         `json:"amount"`
	BeneficiaryAccountNo   string          `json:"beneficiaryAccountNo"`
	BeneficiaryAccountName string          `json:"beneficiaryAccountName"`
	BeneficiaryBankCode    string          `json:"beneficiaryBankCode"`
	SourceAccountNo        string          `json:"sourceAccountNo"`
	TransactionDate        string          `json:"transactionDate"`
	FeeType                string          `json:"feeType"`
	Remark                 string          `json:"remark"`
	AdditionalInfo         *AdditionalInfo `json:"additionalInfo"`
	// BCA-specific fields
	BcaTransferType string `json:"bcaTransferType,omitempty"`
	BcaPriority     string `json:"bcaPriority,omitempty"`
}

// BCATransferResponse represents the response from BCA transfer
type BCATransferResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ReferenceNo     string `json:"referenceNo"`
	// BCA-specific fields
	BcaReferenceNo  string `json:"bcaReferenceNo,omitempty"`
	TransactionDate string `json:"transactionDate,omitempty"`
}

// BCAAdditionalInfo represents BCA-specific additional information
// This type ensures consistent structure for BCA-specific additional info
type BCAAdditionalInfo struct {
	DeviceId        string `json:"deviceId"`
	Channel         string `json:"channel"`
	BcaCustomField1 string `json:"bcaCustomField1,omitempty"`
	BcaCustomField2 string `json:"bcaCustomField2,omitempty"`
	BcaCustomField3 string `json:"bcaCustomField3,omitempty"`
}