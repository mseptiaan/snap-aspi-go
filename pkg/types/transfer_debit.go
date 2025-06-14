package types

// DirectDebitPaymentPayload represents the payload for direct debit payment
type DirectDebitPaymentPayload struct {
	PartnerServiceId     string                `json:"partnerServiceId"`
	CustomerNo           string                `json:"customerNo"`
	PartnerReferenceNo   string                `json:"partnerReferenceNo"`
	BankCardToken        string                `json:"bankCardToken"`
	ChargeToken          string                `json:"chargeToken"`
	OTP                  string                `json:"otp"`
	OTPTrxCode           string                `json:"otpTrxCode"`
	MerchantId           string                `json:"merchantId"`
	TerminalId           string                `json:"terminalId"`
	JourneyId            string                `json:"journeyId"`
	SubMerchantId        string                `json:"subMerchantId"`
	Amount               *Amount               `json:"amount"`
	URLParams            []URLParam            `json:"urlParams"`
	ExternalStoreId      string                `json:"externalStoreId"`
	ValidUpTo            string                `json:"validUpTo"`
	PointOfInitiation    string                `json:"pointOfInitiation"`
	FeeType              string                `json:"feeType"`
	DisabledPayMethods   string                `json:"disabledPayMethods"`
	PayOptionDetails     []PayOptionDetail     `json:"payOptionDetails"`
	AdditionalInfo       *AdditionalInfo       `json:"additionalInfo"`
}

// URLParam represents URL parameter
type URLParam struct {
	URL        string `json:"url"`
	Type       string `json:"type"`
	IsDeeplink string `json:"isDeeplink"`
}

// PayOptionDetail represents payment option detail
type PayOptionDetail struct {
	PayMethod      string          `json:"payMethod"`
	PayOption      string          `json:"payOption"`
	TransAmount    *Amount         `json:"transAmount"`
	FeeAmount      *Amount         `json:"feeAmount"`
	CardToken      string          `json:"cardToken"`
	MerchantToken  string          `json:"merchantToken"`
	AdditionalInfo *AdditionalInfo `json:"additionalInfo"`
}

// DirectDebitStatusPayload represents the payload for direct debit status
type DirectDebitStatusPayload struct {
	PartnerServiceId         string          `json:"partnerServiceId"`
	CustomerNo               string          `json:"customerNo"`
	OriginalPartnerRefNo     string          `json:"originalPartnerReferenceNo"`
	OriginalReferenceNo      string          `json:"originalReferenceNo"`
	OriginalExternalId       string          `json:"originalExternalId"`
	ServiceCode              string          `json:"serviceCode"`
	TransactionDate          string          `json:"transactionDate"`
	Amount                   *Amount         `json:"amount"`
	MerchantId               string          `json:"merchantId"`
	SubMerchantId            string          `json:"subMerchantId"`
	ExternalStoreId          string          `json:"externalStoreId"`
	AdditionalInfo           *AdditionalInfo `json:"additionalInfo"`
}

// DirectDebitCancelPayload represents the payload for direct debit cancel
type DirectDebitCancelPayload struct {
	PartnerServiceId         string          `json:"partnerServiceId"`
	CustomerNo               string          `json:"customerNo"`
	OriginalPartnerRefNo     string          `json:"originalPartnerReferenceNo"`
	OriginalReferenceNo      string          `json:"originalReferenceNo"`
	ApprovalCode             string          `json:"approvalCode"`
	OriginalExternalId       string          `json:"originalExternalId"`
	MerchantId               string          `json:"merchantId"`
	SubMerchantId            string          `json:"subMerchantId"`
	Reason                   string          `json:"reason"`
	ExternalStoreId          string          `json:"externalStoreId"`
	Amount                   *Amount         `json:"amount"`
	AdditionalInfo           *AdditionalInfo `json:"additionalInfo"`
}

// DirectDebitRefundPayload represents the payload for direct debit refund
type DirectDebitRefundPayload struct {
	PartnerServiceId         string          `json:"partnerServiceId"`
	CustomerNo               string          `json:"customerNo"`
	MerchantId               string          `json:"merchantId"`
	SubMerchantId            string          `json:"subMerchantId"`
	OriginalPartnerRefNo     string          `json:"originalPartnerReferenceNo"`
	OriginalReferenceNo      string          `json:"originalReferenceNo"`
	OriginalExternalId       string          `json:"originalExternalId"`
	PartnerRefundNo          string          `json:"partnerRefundNo"`
	RefundAmount             *Amount         `json:"refundAmount"`
	ExternalStoreId          string          `json:"externalStoreId"`
	Reason                   string          `json:"reason"`
	AdditionalInfo           *AdditionalInfo `json:"additionalInfo"`
}

// CPMGenerateQRPayload represents the payload for CPM QR generation
type CPMGenerateQRPayload struct {
	PartnerServiceId   string          `json:"partnerServiceId"`
	CustomerNo         string          `json:"customerNo"`
	PartnerReferenceNo string          `json:"partnerReferenceNo"`
	UserAccessToken    string          `json:"userAccessToken"`
	MerchantId         string          `json:"merchantId"`
	SubMerchantId      string          `json:"subMerchantId"`
	PartnerTrxDate     string          `json:"partnerTrxDate"`
	AdditionalInfo     *AdditionalInfo `json:"additionalInfo"`
}

// CPMPaymentPayload represents the payload for CPM payment
type CPMPaymentPayload struct {
	PartnerServiceId  string          `json:"partnerServiceId"`
	CustomerNo        string          `json:"customerNo"`
	PartnerReferenceNo string         `json:"partnerReferenceNo"`
	QRContent         string          `json:"qrContent"`
	Amount            *Amount         `json:"amount"`
	FeeAmount         *Amount         `json:"feeAmount"`
	MerchantId        string          `json:"merchantId"`
	SubMerchantId     string          `json:"subMerchantId"`
	Title             string          `json:"title"`
	ExpiryTime        string          `json:"expiryTime"`
	Items             *ItemInfo       `json:"items"`
	ExternalStoreId   string          `json:"externalStoreId"`
	MerchantName      string          `json:"merchantName"`
	MerchantLocation  string          `json:"merchantLocation"`
	AcquirerName      string          `json:"acquirerName"`
	TerminalId        string          `json:"terminalId"`
	ScannerInfo       *ScannerInfo    `json:"scannerInfo"`
	AdditionalInfo    *AdditionalInfo `json:"additionalInfo"`
}

// ItemInfo represents item information
type ItemInfo struct {
	ProductId   string `json:"productId"`
	ProductName string `json:"productName"`
	Qty         string `json:"qty"`
	Desc        string `json:"desc"`
}

// ScannerInfo represents scanner information
type ScannerInfo struct {
	DeviceId      string `json:"deviceId"`
	DeviceVersion string `json:"deviceVersion"`
	DeviceModel   string `json:"deviceModel"`
	DeviceIP      string `json:"deviceIp"`
}

// AuthPaymentPayload represents the payload for auth payment
type AuthPaymentPayload struct {
	PartnerServiceId     string            `json:"partnerServiceId"`
	CustomerNo           string            `json:"customerNo"`
	PartnerReferenceNo   string            `json:"partnerReferenceNo"`
	BankCardToken        string            `json:"bankCardToken"`
	ChargeToken          string            `json:"chargeToken"`
	OTP                  string            `json:"otp"`
	OTPTrxCode           string            `json:"otpTrxCode"`
	MerchantId           string            `json:"merchantId"`
	TerminalId           string            `json:"terminalId"`
	JourneyId            string            `json:"journeyId"`
	SubMerchantId        string            `json:"subMerchantId"`
	Amount               *Amount           `json:"amount"`
	URLParams            []URLParam        `json:"urlParams"`
	ExternalStoreId      string            `json:"externalStoreId"`
	ValidUpTo            string            `json:"validUpTo"`
	PointOfInitiation    string            `json:"pointOfInitiation"`
	FeeType              string            `json:"feeType"`
	DisabledPayMethods   string            `json:"disabledPayMethods"`
	PayOptionDetails     []PayOptionDetail `json:"payOptionDetails"`
	AdditionalInfo       *AdditionalInfo   `json:"additionalInfo"`
}

// AuthCapturePayload represents the payload for auth capture
type AuthCapturePayload struct {
	PartnerServiceId         string          `json:"partnerServiceId"`
	CustomerNo               string          `json:"customerNo"`
	OriginalReferenceNo      string          `json:"originalReferenceNo"`
	OriginalPartnerRefNo     string          `json:"originalPartnerReferenceNo"`
	MerchantId               string          `json:"merchantId"`
	SubMerchantId            string          `json:"subMerchantId"`
	PartnerCaptureNo         string          `json:"partnerCaptureNo"`
	CaptureAmount            *Amount         `json:"captureAmount"`
	Title                    string          `json:"title"`
	LastCapture              string          `json:"lastCapture"`
	AdditionalInfo           *AdditionalInfo `json:"additionalInfo"`
}

// AuthVoidPayload represents the payload for auth void
type AuthVoidPayload struct {
	PartnerServiceId         string          `json:"partnerServiceId"`
	CustomerNo               string          `json:"customerNo"`
	OriginalReferenceNo      string          `json:"originalReferenceNo"`
	OriginalPartnerRefNo     string          `json:"originalPartnerReferenceNo"`
	MerchantId               string          `json:"merchantId"`
	SubMerchantId            string          `json:"subMerchantId"`
	VoidAmount               *Amount         `json:"voidAmount"`
	PartnerVoidNo            string          `json:"partnerVoidNo"`
	VoidRemainingAmount      string          `json:"voidRemainingAmount"`
	Reason                   string          `json:"reason"`
	AdditionalInfo           *AdditionalInfo `json:"additionalInfo"`
}

// DirectDebitBIFASTPayload represents the payload for direct debit BI-FAST
type DirectDebitBIFASTPayload struct {
	PartnerServiceId      string          `json:"partnerServiceId"`
	CustomerNo            string          `json:"customerNo"`
	PartnerReferenceNo    string          `json:"partnerReferenceNo"`
	BankCode              string          `json:"bankCode"`
	SourceAccountNo       string          `json:"sourceAccountNo"`
	SourceAccountName     string          `json:"sourceAccountName"`
	MaxAmount             *Amount         `json:"maxAmount"`
	BillerId              string          `json:"billerId"`
	BillerName            string          `json:"billerName"`
	CustomerId            string          `json:"customerId"`
	ExpiredDatetime       string          `json:"expiredDatetime"`
	AdditionalInfo        *AdditionalInfo `json:"additionalInfo"`
}

// DirectDebitBIFASTPaymentPayload represents the payload for BI-FAST payment
type DirectDebitBIFASTPaymentPayload struct {
	PartnerServiceId        string          `json:"partnerServiceId"`
	CustomerNo              string          `json:"customerNo"`
	PartnerReferenceNo      string          `json:"partnerReferenceNo"`
	Currency                string          `json:"currency"`
	CustomerReference       string          `json:"customerReference"`
	FeeType                 string          `json:"feeType"`
	Remark                  string          `json:"remark"`
	BeneficiaryAccountNo    string          `json:"beneficiaryAccountNo"`
	BeneficiaryAccountName  string          `json:"beneficiaryAccountName"`
	TransactionDate         string          `json:"transactionDate"`
	BankCode                string          `json:"bankCode"`
	SourceAccountNo         string          `json:"sourceAccountNo"`
	SourceAccountName       string          `json:"sourceAccountName"`
	Amount                  *Amount         `json:"amount"`
	EMandateReffId          string          `json:"eMandateReffId"`
	AdditionalInfo          *AdditionalInfo `json:"additionalInfo"`
}