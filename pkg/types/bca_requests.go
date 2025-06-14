package types

// BCA-specific request types

// BCAPaymentPayload represents the payload for BCA payment
type BCAPaymentPayload struct {
	PartnerServiceId        string          `json:"partnerServiceId"`
	CustomerNo              string          `json:"customerNo"`
	VirtualAccountNo        string          `json:"virtualAccountNo"`
	VirtualAccountName      string          `json:"virtualAccountName"`
	TrxId                   string          `json:"trxId"`
	PaymentRequestId        string          `json:"paymentRequestId"`
	ChannelCode             string          `json:"channelCode"`
	SourceAccountNo         string          `json:"sourceAccountNo"`
	PaidAmount              *Amount         `json:"paidAmount"`
	TotalAmount             *Amount         `json:"totalAmount"`
	TrxDateTime             string          `json:"trxDateTime"`
	ReferenceNo             string          `json:"referenceNo"`
	PaymentType             string          `json:"paymentType"`
	AdditionalInfo          *AdditionalInfo `json:"additionalInfo"`
	// BCA-specific fields
	BcaCustomerNo      string `json:"bcaCustomerNo,omitempty"`
	BcaPaymentMethod   string `json:"bcaPaymentMethod,omitempty"`
	BcaTransactionType string `json:"bcaTransactionType,omitempty"`
}

// BCAStatusPayload represents the payload for checking BCA transaction status
type BCAStatusPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	VirtualAccountNo string          `json:"virtualAccountNo"`
	ReferenceNo      string          `json:"referenceNo"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
	// BCA-specific fields
	BcaCustomerNo    string `json:"bcaCustomerNo,omitempty"`
	BcaTransactionId string `json:"bcaTransactionId,omitempty"`
}

// BCAAccountInquiryPayload represents the payload for BCA account inquiry
type BCAAccountInquiryPayload struct {
	PartnerServiceId     string          `json:"partnerServiceId"`
	CustomerNo           string          `json:"customerNo"`
	BeneficiaryAccountNo string          `json:"beneficiaryAccountNo"`
	BeneficiaryBankCode  string          `json:"beneficiaryBankCode"`
	AdditionalInfo       *AdditionalInfo `json:"additionalInfo"`
	// BCA-specific fields
	BcaCustomerNo    string `json:"bcaCustomerNo,omitempty"`
	BcaInquiryType   string `json:"bcaInquiryType,omitempty"`
}

// BCATransactionHistoryPayload represents the payload for BCA transaction history
type BCATransactionHistoryPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	AccountNo        string          `json:"accountNo"`
	FromDateTime     string          `json:"fromDateTime"`
	ToDateTime       string          `json:"toDateTime"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
	// BCA-specific fields
	BcaCustomerNo    string `json:"bcaCustomerNo,omitempty"`
	BcaStatementType string `json:"bcaStatementType,omitempty"`
}

// BCAUpdateVAPayload represents the payload for updating a BCA Virtual Account
type BCAUpdateVAPayload struct {
	PartnerServiceId    string          `json:"partnerServiceId"`
	CustomerNo          string          `json:"customerNo"`
	VirtualAccountNo    string          `json:"virtualAccountNo"`
	VirtualAccountName  string          `json:"virtualAccountName"`
	TrxId               string          `json:"trxId"`
	TotalAmount         *Amount         `json:"totalAmount"`
	ExpiredDate         string          `json:"expiredDate"`
	AdditionalInfo      *AdditionalInfo `json:"additionalInfo"`
	// BCA-specific fields
	BcaCustomerNo      string `json:"bcaCustomerNo,omitempty"`
	BcaSubCompany      string `json:"bcaSubCompany,omitempty"`
	BcaTransactionType string `json:"bcaTransactionType,omitempty"`
	BcaUpdateType      string `json:"bcaUpdateType,omitempty"`
}

// BCADeleteVAPayload represents the payload for deleting a BCA Virtual Account
type BCADeleteVAPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	VirtualAccountNo string          `json:"virtualAccountNo"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
	// BCA-specific fields
	BcaCustomerNo string `json:"bcaCustomerNo,omitempty"`
	BcaReason     string `json:"bcaReason,omitempty"`
}