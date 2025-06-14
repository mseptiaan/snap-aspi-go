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
	Amount             *Amount         `json:"amount"`
	FeeAmount          *Amount         `json:"feeAmount"`
	MerchantId         string          `json:"merchantId"`
	SubMerchantId      string          `json:"subMerchantId"`
	StoreId            string          `json:"storeId"`
	TerminalId         string          `json:"terminalId"`
	ValidityPeriod     string          `json:"validityPeriod"`
	MerchantName       string          `json:"merchantName"`
	MerchantCity       string          `json:"merchantCity"`
	MerchantPostalCode string          `json:"merchantPostalCode"`
	TerminalLabel      string          `json:"terminalLabel"`
	PurposeOfPayment   string          `json:"purposeOfPayment"`
	QRType             string          `json:"qrType"`
	AdditionalInfo     *AdditionalInfo `json:"additionalInfo"`
}

// MPMDecodeQRPayload represents the payload for decoding QR MPM
type MPMDecodeQRPayload struct {
	PartnerServiceId   string          `json:"partnerServiceId"`
	CustomerNo         string          `json:"customerNo"`
	PartnerReferenceNo string          `json:"partnerReferenceNo"`
	QRContent          string          `json:"qrContent"`
	Amount             *Amount         `json:"amount"`
	MerchantId         string          `json:"merchantId"`
	SubMerchantId      string          `json:"subMerchantId"`
	ScanTime           string          `json:"scanTime"`
	AdditionalInfo     *AdditionalInfo `json:"additionalInfo"`
}

// MPMApplyOTTPayload represents the payload for applying OTT
type MPMApplyOTTPayload struct {
	UserResources []string        `json:"userResources"`
	AdditionalInfo *AdditionalInfo `json:"additionalInfo"`
}

// MPMPaymentPayload represents the payload for MPM payment
type MPMPaymentPayload struct {
	PartnerServiceId   string          `json:"partnerServiceId"`
	CustomerNo         string          `json:"customerNo"`
	PartnerReferenceNo string          `json:"partnerReferenceNo"`
	MerchantId         string          `json:"merchantId"`
	SubMerchantId      string          `json:"subMerchantId"`
	Amount             *Amount         `json:"amount"`
	FeeAmount          *Amount         `json:"feeAmount"`
	OTP                string          `json:"otp"`
	VerificationId     string          `json:"verificationId"`
	AdditionalInfo     *AdditionalInfo `json:"additionalInfo"`
}

// MPMQueryPaymentPayload represents the payload for querying MPM payment
type MPMQueryPaymentPayload struct {
	PartnerServiceId      string          `json:"partnerServiceId"`
	CustomerNo            string          `json:"customerNo"`
	OriginalReferenceNo   string          `json:"originalReferenceNo"`
	OriginalPartnerRefNo  string          `json:"originalPartnerReferenceNo"`
	OriginalExternalId    string          `json:"originalExternalId"`
	ServiceCode           string          `json:"serviceCode"`
	MerchantId            string          `json:"merchantId"`
	SubMerchantId         string          `json:"subMerchantId"`
	ExternalStoreId       string          `json:"externalStoreId"`
	AdditionalInfo        *AdditionalInfo `json:"additionalInfo"`
}

// MPMCancelPaymentPayload represents the payload for cancelling MPM payment
type MPMCancelPaymentPayload struct {
	PartnerServiceId      string          `json:"partnerServiceId"`
	CustomerNo            string          `json:"customerNo"`
	OriginalPartnerRefNo  string          `json:"originalPartnerReferenceNo"`
	OriginalReferenceNo   string          `json:"originalReferenceNo"`
	OriginalExternalId    string          `json:"originalExternalId"`
	MerchantId            string          `json:"merchantId"`
	SubMerchantId         string          `json:"subMerchantId"`
	ExternalStoreId       string          `json:"externalStoreId"`
	Reason                string          `json:"reason"`
	Amount                *Amount         `json:"amount"`
	AdditionalInfo        *AdditionalInfo `json:"additionalInfo"`
}

// MPMNotifyQRPayload represents the payload for QR MPM notification
// Based on ASPI Notify QR MPM API documentation
type MPMNotifyQRPayload struct {
	OriginalReferenceNo     string          `json:"originalReferenceNo"`
	OriginalPartnerRefNo    string          `json:"originalPartnerReferenceNo"`
	LatestTransactionStatus string          `json:"latestTransactionStatus"`
	TransactionStatusDesc   string          `json:"transactionStatusDesc"`
	CustomerNumber          string          `json:"customerNumber"`
	AccountType             string          `json:"accountType"`
	DestinationNumber       string          `json:"destinationNumber"`
	DestinationAccountName  string          `json:"destinationAccountName"`
	Amount                  *Amount         `json:"amount"`
	SessionId               string          `json:"sessionId"`
	BankCode                string          `json:"bankCode"`
	ExternalStoreId         string          `json:"externalStoreId"`
	AdditionalInfo          *AdditionalInfo `json:"additionalInfo"`
}

// MPMQRResponse represents the response for QR generation
// Updated to match ASPI documentation
type MPMQRResponse struct {
	ResponseCode    string          `json:"responseCode"`
	ResponseMessage string          `json:"responseMessage"`
	ReferenceNo     string          `json:"referenceNo"`
	PartnerRefNo    string          `json:"partnerReferenceNo"`
	QRContent       string          `json:"qrContent"`
	QRUrl           string          `json:"qrUrl"`
	QRImage         string          `json:"qrImage"`
	RedirectUrl     string          `json:"redirectUrl"`
	MerchantName    string          `json:"merchantName"`
	StoreId         string          `json:"storeId"`
	TerminalId      string          `json:"terminalId"`
	ValidUntil      string          `json:"expiryTime"`
	AdditionalInfo  *AdditionalInfo `json:"additionalInfo"`
}

// MPMNotifyResponse represents the response for QR notification
type MPMNotifyResponse struct {
	ResponseCode    string          `json:"responseCode"`
	ResponseMessage string          `json:"responseMessage"`
	AdditionalInfo  *AdditionalInfo `json:"additionalInfo"`
}

// MPMDecodeQRResponse represents the response for QR decoding
type MPMDecodeQRResponse struct {
	ResponseCode       string          `json:"responseCode"`
	ResponseMessage    string          `json:"responseMessage"`
	ReferenceNo        string          `json:"referenceNo"`
	PartnerRefNo       string          `json:"partnerReferenceNo"`
	RedirectUrl        string          `json:"redirectUrl"`
	MerchantName       string          `json:"merchantName"`
	MerchantCategory   string          `json:"merchantCategory"`
	MerchantLocation   string          `json:"merchantLocation"`
	MerchantInfos      []MerchantInfo  `json:"merchantInfos"`
	TransactionAmount  *Amount         `json:"transactionAmount"`
	FeeAmount          *Amount         `json:"feeAmount"`
	AdditionalInfo     *AdditionalInfo `json:"additionalInfo"`
}

// MerchantInfo represents merchant information
type MerchantInfo struct {
	MerchantPAN   string `json:"merchantPAN"`
	AcquirerName  string `json:"acquirerName"`
}

// MPMApplyOTTResponse represents the response for applying OTT
type MPMApplyOTTResponse struct {
	ResponseCode    string          `json:"responseCode"`
	ResponseMessage string          `json:"responseMessage"`
	UserResources   []UserResource  `json:"userResources"`
	AdditionalInfo  *AdditionalInfo `json:"additionalInfo"`
}

// UserResource represents a user resource
type UserResource struct {
	ResourceType string `json:"resourceType"`
	Value        string `json:"value"`
}

// MPMPaymentResponse represents the response for MPM payment
type MPMPaymentResponse struct {
	ResponseCode    string          `json:"responseCode"`
	ResponseMessage string          `json:"responseMessage"`
	ReferenceNo     string          `json:"referenceNo"`
	PartnerRefNo    string          `json:"partnerReferenceNo"`
	TransactionDate string          `json:"transactionDate"`
	Amount          *Amount         `json:"amount"`
	FeeAmount       *Amount         `json:"feeAmount"`
	VerificationId  string          `json:"verificationId"`
	AdditionalInfo  *AdditionalInfo `json:"additionalInfo"`
}

// MPMQueryPaymentResponse represents the response for querying MPM payment
type MPMQueryPaymentResponse struct {
	ResponseCode            string          `json:"responseCode"`
	ResponseMessage         string          `json:"responseMessage"`
	OriginalReferenceNo     string          `json:"originalReferenceNo"`
	OriginalPartnerRefNo    string          `json:"originalPartnerReferenceNo"`
	OriginalExternalId      string          `json:"originalExternalId"`
	ServiceCode             string          `json:"serviceCode"`
	LatestTransactionStatus string          `json:"latestTransactionStatus"`
	TransactionStatusDesc   string          `json:"transactionStatusDesc"`
	PaidTime                string          `json:"paidTime"`
	Amount                  *Amount         `json:"amount"`
	FeeAmount               *Amount         `json:"feeAmount"`
	TerminalId              string          `json:"terminalId"`
	AdditionalInfo          *AdditionalInfo `json:"additionalInfo"`
}