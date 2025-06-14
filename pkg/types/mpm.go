package types

// MPMTransferPayload represents the payload for MPM Transfer Credit
// Based on ASPI Transfer Credit MPM API documentation
type MPMTransferPayload struct {
	PartnerServiceId      string          `json:"partnerServiceId"`
	CustomerNo            string          `json:"customerNo"`
	PartnerReferenceNo    string          `json:"partnerReferenceNo"`
	MerchantId            string          `json:"merchantId"`
	SubMerchantId         string          `json:"subMerchantId"`
	ExternalStoreId       string          `json:"externalStoreId"`
	Amount                *Amount         `json:"amount"`
	BeneficiaryAccountNo  string          `json:"beneficiaryAccountNo"`
	BeneficiaryName       string          `json:"beneficiaryName"`
	BeneficiaryCustomerNo string          `json:"beneficiaryCustomerNo"`
	BeneficiaryEmail      string          `json:"beneficiaryEmail"`
	Currency              string          `json:"currency"`
	CustomerReference     string          `json:"customerReference"`
	FeeType               string          `json:"feeType"`
	Remark                string          `json:"remark"`
	SourceAccountNo       string          `json:"sourceAccountNo"`
	TransactionDate       string          `json:"transactionDate"`
	AdditionalInfo        *AdditionalInfo `json:"additionalInfo"`
}

// MPMInquiryPayload represents the payload for MPM Transfer inquiry
type MPMInquiryPayload struct {
	PartnerServiceId      string          `json:"partnerServiceId"`
	CustomerNo            string          `json:"customerNo"`
	PartnerReferenceNo    string          `json:"partnerReferenceNo"`
	BeneficiaryAccountNo  string          `json:"beneficiaryAccountNo"`
	BeneficiaryCustomerNo string          `json:"beneficiaryCustomerNo"`
	Amount                *Amount         `json:"amount"`
	Currency              string          `json:"currency"`
	AdditionalInfo        *AdditionalInfo `json:"additionalInfo"`
}

// MPMStatusPayload represents the payload for checking MPM transfer status
type MPMStatusPayload struct {
	PartnerServiceId   string          `json:"partnerServiceId"`
	CustomerNo         string          `json:"customerNo"`
	PartnerReferenceNo string          `json:"partnerReferenceNo"`
	ServiceCode        string          `json:"serviceCode"`
	TransactionDate    string          `json:"transactionDate"`
	ExternalStoreId    string          `json:"externalStoreId"`
	AdditionalInfo     *AdditionalInfo `json:"additionalInfo"`
}

// MPMRefundPayload represents the payload for MPM transfer refund
type MPMRefundPayload struct {
	PartnerServiceId    string          `json:"partnerServiceId"`
	CustomerNo          string          `json:"customerNo"`
	OriginalReferenceNo string          `json:"originalReferenceNo"`
	PartnerRefundNo     string          `json:"partnerRefundNo"`
	RefundAmount        *Amount         `json:"refundAmount"`
	Reason              string          `json:"reason"`
	RefundType          string          `json:"refundType"`
	AdditionalInfo      *AdditionalInfo `json:"additionalInfo"`
}

// MPMBalanceInquiryPayload represents the payload for MPM balance inquiry
type MPMBalanceInquiryPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	AccountNo        string          `json:"accountNo"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
}

// MPMAccountInquiryPayload represents the payload for MPM account inquiry
type MPMAccountInquiryPayload struct {
	PartnerServiceId      string          `json:"partnerServiceId"`
	CustomerNo            string          `json:"customerNo"`
	BeneficiaryAccountNo  string          `json:"beneficiaryAccountNo"`
	BeneficiaryCustomerNo string          `json:"beneficiaryCustomerNo"`
	AdditionalInfo        *AdditionalInfo `json:"additionalInfo"`
}

// MPMHistoryPayload represents the payload for MPM transaction history
type MPMHistoryPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	AccountNo        string          `json:"accountNo"`
	FromDateTime     string          `json:"fromDateTime"`
	ToDateTime       string          `json:"toDateTime"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
}

// MPMGenerateQRPayload represents the payload for generating QR MPM
// Based on ASPI Generate QR MPM API documentation
type MPMGenerateQRPayload struct {
	PartnerServiceId   string          `json:"partnerServiceId"`
	CustomerNo         string          `json:"customerNo"`
	PartnerReferenceNo string          `json:"partnerReferenceNo"`
	MerchantId         string          `json:"merchantId"`
	SubMerchantId      string          `json:"subMerchantId"`
	ExternalStoreId    string          `json:"externalStoreId"`
	Amount             *Amount         `json:"amount"`
	MerchantName       string          `json:"merchantName"`
	MerchantCity       string          `json:"merchantCity"`
	MerchantPostalCode string          `json:"merchantPostalCode"`
	TerminalLabel      string          `json:"terminalLabel"`
	PurposeOfPayment   string          `json:"purposeOfPayment"`
	QRType             string          `json:"qrType"`
	ValidityPeriod     string          `json:"validityPeriod"`
	AdditionalInfo     *AdditionalInfo `json:"additionalInfo"`
}

// MPMNotifyQRPayload represents the payload for QR MPM notification
// Based on ASPI Notify QR MPM API documentation
type MPMNotifyQRPayload struct {
	PartnerServiceId        string          `json:"partnerServiceId"`
	CustomerNo              string          `json:"customerNo"`
	PartnerReferenceNo      string          `json:"partnerReferenceNo"`
	MerchantId              string          `json:"merchantId"`
	SubMerchantId           string          `json:"subMerchantId"`
	ExternalStoreId         string          `json:"externalStoreId"`
	OriginalReferenceNo     string          `json:"originalReferenceNo"`
	OriginalPartnerRefNo    string          `json:"originalPartnerRefNo"`
	ServiceCode             string          `json:"serviceCode"`
	TransactionDate         string          `json:"transactionDate"`
	Amount                  *Amount         `json:"amount"`
	FeeAmount               *Amount         `json:"feeAmount"`
	PaidAmount              *Amount         `json:"paidAmount"`
	CurrencyCode            string          `json:"currencyCode"`
	TransactionStatus       string          `json:"transactionStatus"`
	TransactionStatusDesc   string          `json:"transactionStatusDesc"`
	LatestTransactionStatus string          `json:"latestTransactionStatus"`
	PaymentMethod           string          `json:"paymentMethod"`
	GatewayReferenceNo      string          `json:"gatewayReferenceNo"`
	AdditionalInfo          *AdditionalInfo `json:"additionalInfo"`
}

// MPMQRResponse represents the response for QR generation
type MPMQRResponse struct {
	ResponseCode    string          `json:"responseCode"`
	ResponseMessage string          `json:"responseMessage"`
	ReferenceNo     string          `json:"referenceNo"`
	PartnerRefNo    string          `json:"partnerRefNo"`
	QRContent       string          `json:"qrContent"`
	QRString        string          `json:"qrString"`
	ValidUntil      string          `json:"validUntil"`
	AdditionalInfo  *AdditionalInfo `json:"additionalInfo"`
}

// MPMNotifyResponse represents the response for QR notification
type MPMNotifyResponse struct {
	ResponseCode    string          `json:"responseCode"`
	ResponseMessage string          `json:"responseMessage"`
	ReferenceNo     string          `json:"referenceNo"`
	AdditionalInfo  *AdditionalInfo `json:"additionalInfo"`
}
