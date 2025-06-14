package types

// BillDetail represents a bill detail structure for Virtual Account
type BillDetail struct {
	BillCode        string            `json:"billCode"`
	BillNo          string            `json:"billNo"`
	BillName        string            `json:"billName"`
	BillShortName   string            `json:"billShortName"`
	BillDescription map[string]string `json:"billDescription"`
	BillSubCompany  string            `json:"billSubCompany"`
	BillAmount      *Amount           `json:"billAmount"`
	AdditionalInfo  map[string]any    `json:"additionalInfo"`
}

// FreeText represents multi-language free text
type FreeText struct {
	English   string `json:"english"`
	Indonesia string `json:"indonesia"`
}

// AdditionalInfo represents additional information for requests
type AdditionalInfo struct {
	DeviceId string `json:"deviceId"`
	Channel  string `json:"channel"`
}

// CreateVAPayload represents the payload for creating a Virtual Account
// Equivalent to PHP CreateVAPayload class
type CreateVAPayload struct {
	PartnerServiceId      string          `json:"partnerServiceId"`
	CustomerNo            string          `json:"customerNo"`
	VirtualAccountNo      string          `json:"virtualAccountNo"`
	VirtualAccountName    string          `json:"virtualAccountName"`
	VirtualAccountEmail   string          `json:"virtualAccountEmail"`
	VirtualAccountPhone   string          `json:"virtualAccountPhone"`
	TrxId                 string          `json:"trxId"`
	TotalAmount           *Amount         `json:"totalAmount"`
	BillDetails           []BillDetail    `json:"billDetails"`
	FreeTexts             []FreeText      `json:"freeTexts"`
	VirtualAccountTrxType string          `json:"virtualAccountTrxType"`
	FeeAmount             *Amount         `json:"feeAmount"`
	ExpiredDate           string          `json:"expiredDate"`
	AdditionalInfo        *AdditionalInfo `json:"additionalInfo"`
}

// InquiryVAPayload represents the payload for Virtual Account inquiry
type InquiryVAPayload struct {
	PartnerServiceId string `json:"partnerServiceId"`
	CustomerNo       string `json:"customerNo"`
	VirtualAccountNo string `json:"virtualAccountNo"`
	TrxDateInit      string `json:"trxDateInit"`
}

// InquiryPayload represents a general inquiry payload
type InquiryPayload struct {
	PartnerServiceId      string          `json:"partnerServiceId"`
	CustomerNo            string          `json:"customerNo"`
	VirtualAccountNo      string          `json:"virtualAccountNo"`
	TrxDateInit           string          `json:"trxDateInit"`
	ChannelCode           string          `json:"channelCode"`
	Language              string          `json:"language"`
	Amount                *Amount         `json:"amount"`
	HashedSourceAccountNo string          `json:"hashedSourceAccountNo"`
	SourceAccountNo       string          `json:"sourceAccountNo"`
	AdditionalInfo        *AdditionalInfo `json:"additionalInfo"`
}

// PaymentPayload represents the payload for Virtual Account payment
type PaymentPayload struct {
	PartnerServiceId        string          `json:"partnerServiceId"`
	CustomerNo              string          `json:"customerNo"`
	VirtualAccountNo        string          `json:"virtualAccountNo"`
	VirtualAccountName      string          `json:"virtualAccountName"`
	VirtualAccountEmail     string          `json:"virtualAccountEmail"`
	VirtualAccountPhone     string          `json:"virtualAccountPhone"`
	TrxId                   string          `json:"trxId"`
	PaymentRequestId        string          `json:"paymentRequestId"`
	ChannelCode             string          `json:"channelCode"`
	SourceAccountNo         string          `json:"sourceAccountNo"`
	PaidAmount              *Amount         `json:"paidAmount"`
	CumulativePaymentAmount *Amount         `json:"cumulativePaymentAmount"`
	PaidBills               string          `json:"paidBills"`
	TotalAmount             *Amount         `json:"totalAmount"`
	TrxDateTime             string          `json:"trxDateTime"`
	ReferenceNo             string          `json:"referenceNo"`
	JournalNum              string          `json:"journalNum"`
	PaymentType             string          `json:"paymentType"`
	FlagAdvise              string          `json:"flagAdvise"`
	AdditionalInfo          *AdditionalInfo `json:"additionalInfo"`
}

// StatusPayload represents the payload for checking status
type StatusPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	VirtualAccountNo string          `json:"virtualAccountNo"`
	InquiryRequestId string          `json:"inquiryRequestId"`
	PaymentRequestId string          `json:"paymentRequestId"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
}

// ReportPayload represents the payload for generating reports
type ReportPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	VirtualAccountNo string          `json:"virtualAccountNo"`
	TrxDateInit      string          `json:"trxDateInit"`
	TrxDateEnd       string          `json:"trxDateEnd"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
}
