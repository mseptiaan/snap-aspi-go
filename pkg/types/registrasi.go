package types

// RegistrationPayload represents the payload for user registration
type RegistrationPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	PhoneNo          string          `json:"phoneNo"`
	Email            string          `json:"email"`
	Name             string          `json:"name"`
	IdNumber         string          `json:"idNumber"`
	IdType           string          `json:"idType"`
	Address          string          `json:"address"`
	DateOfBirth      string          `json:"dateOfBirth"`
	PlaceOfBirth     string          `json:"placeOfBirth"`
	Nationality      string          `json:"nationality"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
}

// RegistrationResponse represents the response from registration
type RegistrationResponse struct {
	ResponseCode    string          `json:"responseCode"`
	ResponseMessage string          `json:"responseMessage"`
	ReferenceNo     string          `json:"referenceNo"`
	CustomerNo      string          `json:"customerNo"`
	RegistrationId  string          `json:"registrationId"`
	Status          string          `json:"status"`
	AdditionalInfo  *AdditionalInfo `json:"additionalInfo"`
}

// CardRegistrationPayload represents the payload for card registration
type CardRegistrationPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	CardNumber       string          `json:"cardNumber"`
	CardType         string          `json:"cardType"`
	ExpiryDate       string          `json:"expiryDate"`
	CVV              string          `json:"cvv"`
	CardHolderName   string          `json:"cardHolderName"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
}

// AccountBindingPayload represents the payload for account binding
type AccountBindingPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	AccountNo        string          `json:"accountNo"`
	BankCode         string          `json:"bankCode"`
	AccountName      string          `json:"accountName"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
}

// OTPVerificationPayload represents the payload for OTP verification
type OTPVerificationPayload struct {
	PartnerServiceId string          `json:"partnerServiceId"`
	CustomerNo       string          `json:"customerNo"`
	OTPCode          string          `json:"otpCode"`
	ReferenceNo      string          `json:"referenceNo"`
	AdditionalInfo   *AdditionalInfo `json:"additionalInfo"`
}