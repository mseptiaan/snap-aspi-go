package types

// BalanceInquiryPayload represents the payload for balance inquiry
type BalanceInquiryPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	BankCardToken    string          `json:"bankCardToken"`
	AccountNo        string          `json:"accountNo"`
	BalanceTypes     []string        `json:"balanceTypes"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
}

// BalanceInquiryResponse represents the response from balance inquiry
type BalanceInquiryResponse struct {
	ResponseCode    string          `json:"responseCode"`
	ResponseMessage string          `json:"responseMessage"`
	ReferenceNo     string          `json:"referenceNo"`
	PartnerRefNo    string          `json:"partnerReferenceNo"`
	AccountNo       string          `json:"accountNo"`
	Name            string          `json:"name"`
	AccountInfos    []AccountInfo   `json:"accountInfos"`
	AdditionalInfo  *AdditionalInfo `json:"additionalInfo"`
}

// AccountInfo represents account balance information
type AccountInfo struct {
	BalanceType              string  `json:"balanceType"`
	Amount                   *Amount `json:"amount"`
	FloatAmount              *Amount `json:"floatAmount"`
	HoldAmount               *Amount `json:"holdAmount"`
	AvailableBalance         *Amount `json:"availableBalance"`
	LedgerBalance            *Amount `json:"ledgerBalance"`
	CurrentMultilateralLimit *Amount `json:"currentMultilateralLimit"`
	RegistrationStatusCode   string  `json:"registrationStatusCode"`
	Status                   string  `json:"status"`
}