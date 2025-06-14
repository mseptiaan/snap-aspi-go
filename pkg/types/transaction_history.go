package types

// TransactionHistoryListPayload represents the payload for transaction history list
type TransactionHistoryListPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	FromDateTime     string          `json:"fromDateTime"`
	ToDateTime       string          `json:"toDateTime"`
	PageSize         int             `json:"pageSize"`
	PageNumber       int             `json:"pageNumber"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
}

// TransactionHistoryDetailPayload represents the payload for transaction history detail
type TransactionHistoryDetailPayload struct {
	PartnerServiceId         string          `json:"partnerServiceId"`
	CustomerNo               string          `json:"customerNo"`
	OriginalPartnerRefNo     string          `json:"originalPartnerReferenceNo"`
	AdditionalInfo           *AdditionalInfo `json:"additionalInfo"`
}

// BankStatementPayload represents the payload for bank statement
type BankStatementPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	BankCardToken    string          `json:"bankCardToken"`
	AccountNo        string          `json:"accountNo"`
	FromDateTime     string          `json:"fromDateTime"`
	ToDateTime       string          `json:"toDateTime"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
}

// TransactionHistoryResponse represents the response from transaction history
type TransactionHistoryResponse struct {
	ResponseCode    string              `json:"responseCode"`
	ResponseMessage string              `json:"responseMessage"`
	ReferenceNo     string              `json:"referenceNo"`
	PartnerRefNo    string              `json:"partnerReferenceNo"`
	DetailData      []TransactionDetail `json:"detailData"`
	AdditionalInfo  *AdditionalInfo     `json:"additionalInfo"`
}

// TransactionDetail represents a single transaction detail
type TransactionDetail struct {
	DateTime       string            `json:"dateTime"`
	Amount         *Amount           `json:"amount"`
	Remark         string            `json:"remark"`
	SourceOfFunds  []SourceOfFund    `json:"sourceOfFunds"`
	Status         string            `json:"status"`
	Type           string            `json:"type"`
	AdditionalInfo *AdditionalInfo   `json:"additionalInfo"`
}

// SourceOfFund represents source of fund information
type SourceOfFund struct {
	Source string  `json:"source"`
	Amount *Amount `json:"amount"`
}

// BankStatementResponse represents the response from bank statement
type BankStatementResponse struct {
	ResponseCode         string                   `json:"responseCode"`
	ResponseMessage      string                   `json:"responseMessage"`
	ReferenceNo          string                   `json:"referenceNo"`
	PartnerRefNo         string                   `json:"partnerReferenceNo"`
	Balance              []BalanceInfo            `json:"balance"`
	TotalCreditEntries   *DebitCreditEntries      `json:"totalCreditEntries"`
	TotalDebitEntries    *DebitCreditEntries      `json:"totalDebitEntries"`
	HasMore              string                   `json:"hasMore"`
	LastRecordDateTime   string                   `json:"lastRecordDateTime"`
	DetailData           []BankStatementDetail    `json:"detailData"`
	AdditionalInfo       *AdditionalInfo          `json:"additionalInfo"`
}

// BalanceInfo represents balance information
type BalanceInfo struct {
	Amount          *Amount `json:"amount"`
	StartingBalance *Amount `json:"startingBalance"`
	EndingBalance   *Amount `json:"endingBalance"`
	DateTime        string  `json:"dateTime"`
}

// DebitCreditEntries represents debit/credit entries summary
type DebitCreditEntries struct {
	NumberOfEntries int     `json:"numberOfEntries"`
	Amount          *Amount `json:"amount"`
}

// BankStatementDetail represents bank statement detail
type BankStatementDetail struct {
	DetailBalance             *DetailBalance  `json:"detailBalance"`
	Amount                    *Amount         `json:"amount"`
	OriginAmount              *Amount         `json:"originAmount"`
	TransactionDate           string          `json:"transactionDate"`
	Remark                    string          `json:"remark"`
	TransactionId             string          `json:"transactionId"`
	Type                      string          `json:"type"`
	TransactionDetailStatus   string          `json:"transactionDetailStatus"`
	DetailInfo                *AdditionalInfo `json:"detailInfo"`
}

// DetailBalance represents detailed balance information
type DetailBalance struct {
	StartAmount []BalanceAmount `json:"startAmount"`
	EndAmount   []BalanceAmount `json:"endAmount"`
}

// BalanceAmount represents balance amount
type BalanceAmount struct {
	Amount *Amount `json:"amount"`
}