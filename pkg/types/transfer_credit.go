package types

// AccountInquiryPayload represents the payload for account inquiry
type AccountInquiryPayload struct {
	PartnerServiceId      string          `json:"partnerServiceId"`
	CustomerNo            string          `json:"customerNo"`
	PartnerReferenceNo    string          `json:"partnerReferenceNo"`
	BeneficiaryAccountNo  string          `json:"beneficiaryAccountNo"`
	BeneficiaryBankCode   string          `json:"beneficiaryBankCode"`
	AdditionalInfo        *AdditionalInfo `json:"additionalInfo"`
}

// TriggerTransferPayload represents the payload for trigger transfer
type TriggerTransferPayload struct {
	PartnerServiceId      string          `json:"partnerServiceId"`
	CustomerNo            string          `json:"customerNo"`
	PartnerReferenceNo    string          `json:"partnerReferenceNo"`
	Amount                *Amount         `json:"amount"`
	BeneficiaryAccountNo  string          `json:"beneficiaryAccountNo"`
	BeneficiaryAccountName string         `json:"beneficiaryAccountName"`
	BeneficiaryBankCode   string          `json:"beneficiaryBankCode"`
	BeneficiaryBankName   string          `json:"beneficiaryBankName"`
	SourceAccountNo       string          `json:"sourceAccountNo"`
	TransactionDate       string          `json:"transactionDate"`
	FeeType               string          `json:"feeType"`
	Remark                string          `json:"remark"`
	Currency              string          `json:"currency"`
	CustomerReference     string          `json:"customerReference"`
	AdditionalInfo        *AdditionalInfo `json:"additionalInfo"`
}

// TransferStatusInquiryPayload represents the payload for transfer status inquiry
type TransferStatusInquiryPayload struct {
	PartnerServiceId        string          `json:"partnerServiceId"`
	CustomerNo              string          `json:"customerNo"`
	OriginalPartnerRefNo    string          `json:"originalPartnerReferenceNo"`
	OriginalReferenceNo     string          `json:"originalReferenceNo"`
	OriginalExternalId      string          `json:"originalExternalId"`
	ServiceCode             string          `json:"serviceCode"`
	TransactionDate         string          `json:"transactionDate"`
	Amount                  *Amount         `json:"amount"`
	AdditionalInfo          *AdditionalInfo `json:"additionalInfo"`
}

// CustomerTopUpPayload represents the payload for customer top up
type CustomerTopUpPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	CustomerName     string          `json:"customerName"`
	Amount           *Amount         `json:"amount"`
	FeeAmount        *Amount         `json:"feeAmount"`
	TransactionDate  string          `json:"transactionDate"`
	SessionId        string          `json:"sessionId"`
	CategoryId       int             `json:"categoryId"`
	Notes            string          `json:"notes"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
}

// BulkCashinPayload represents the payload for bulk cashin
type BulkCashinPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	PartnerBulkId    string          `json:"partnerBulkId"`
	TransactionDate  string          `json:"transactionDate"`
	Currency         string          `json:"currency"`
	BulkObject       []BulkItem      `json:"bulkObject"`
	FeeType          string          `json:"feeType"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
}

// BulkItem represents a single item in bulk transaction
type BulkItem struct {
	AccountNumber      string          `json:"accountNumber"`
	AccountName        string          `json:"accountName"`
	Amount             *Amount         `json:"amount"`
	PartnerReferenceNo string          `json:"partnerReferenceNo"`
	AdditionalInfo     *AdditionalInfo `json:"additionalInfo"`
}

// TransferToBankPayload represents the payload for transfer to bank
type TransferToBankPayload struct {
	PartnerServiceId         string          `json:"partnerServiceId"`
	CustomerNo               string          `json:"customerNo"`
	AccountType              string          `json:"accountType"`
	BeneficiaryAccountNumber string          `json:"beneficiaryAccountNumber"`
	BeneficiaryBankCode      string          `json:"beneficiaryBankCode"`
	Amount                   *Amount         `json:"amount"`
	SessionId                string          `json:"sessionId"`
	FeeType                  string          `json:"feeType"`
	AdditionalInfo           *AdditionalInfo `json:"additionalInfo"`
}

// TransferToOTCPayload represents the payload for transfer to OTC
type TransferToOTCPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	OTP              string          `json:"otp"`
	Amount           *Amount         `json:"amount"`
	FeeType          string          `json:"feeType"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
}