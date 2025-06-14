package types

// BCA-specific response types

// BCAErrorResponse represents a BCA-specific error response
type BCAErrorResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ErrorCode       string `json:"errorCode,omitempty"`
	ErrorMessage    string `json:"errorMessage,omitempty"`
	BcaErrorCode    string `json:"bcaErrorCode,omitempty"`
	BcaErrorMessage string `json:"bcaErrorMessage,omitempty"`
}

// BCAPaymentResponse represents a BCA-specific payment response
type BCAPaymentResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ReferenceNo     string `json:"referenceNo"`
	PaymentNo       string `json:"paymentNo"`
	TransactionDate string `json:"transactionDate"`
	// BCA-specific fields
	BcaReferenceNo  string `json:"bcaReferenceNo,omitempty"`
	BcaPaymentId    string `json:"bcaPaymentId,omitempty"`
	BcaReceiptNo    string `json:"bcaReceiptNo,omitempty"`
}

// BCAStatusResponse represents a BCA-specific status response
type BCAStatusResponse struct {
	ResponseCode            string `json:"responseCode"`
	ResponseMessage         string `json:"responseMessage"`
	OriginalReferenceNo     string `json:"originalReferenceNo"`
	OriginalPartnerRefNo    string `json:"originalPartnerReferenceNo"`
	LatestTransactionStatus string `json:"latestTransactionStatus"`
	TransactionStatusDesc   string `json:"transactionStatusDesc"`
	// BCA-specific fields
	BcaTransactionId string `json:"bcaTransactionId,omitempty"`
	BcaStatusCode    string `json:"bcaStatusCode,omitempty"`
}

// BCAAccountInquiryResponse represents a BCA-specific account inquiry response
type BCAAccountInquiryResponse struct {
	ResponseCode          string `json:"responseCode"`
	ResponseMessage       string `json:"responseMessage"`
	ReferenceNo           string `json:"referenceNo"`
	BeneficiaryAccountNo  string `json:"beneficiaryAccountNo"`
	BeneficiaryAccountName string `json:"beneficiaryAccountName"`
	BeneficiaryBankCode   string `json:"beneficiaryBankCode"`
	BeneficiaryBankName   string `json:"beneficiaryBankName"`
	// BCA-specific fields
	BcaAccountType    string `json:"bcaAccountType,omitempty"`
	BcaAccountStatus  string `json:"bcaAccountStatus,omitempty"`
	BcaCustomerNumber string `json:"bcaCustomerNumber,omitempty"`
}

// BCATransactionHistoryResponse represents a BCA-specific transaction history response
type BCATransactionHistoryResponse struct {
	ResponseCode    string              `json:"responseCode"`
	ResponseMessage string              `json:"responseMessage"`
	ReferenceNo     string              `json:"referenceNo"`
	DetailData      []BCATransactionDetail `json:"detailData"`
	// BCA-specific fields
	BcaStatementId  string `json:"bcaStatementId,omitempty"`
	BcaAccountType  string `json:"bcaAccountType,omitempty"`
}

// BCATransactionDetail represents a BCA-specific transaction detail
type BCATransactionDetail struct {
	DateTime       string            `json:"dateTime"`
	Amount         *Amount           `json:"amount"`
	Remark         string            `json:"remark"`
	Status         string            `json:"status"`
	Type           string            `json:"type"`
	// BCA-specific fields
	BcaTransactionId   string `json:"bcaTransactionId,omitempty"`
	BcaTransactionType string `json:"bcaTransactionType,omitempty"`
	BcaDescription     string `json:"bcaDescription,omitempty"`
}