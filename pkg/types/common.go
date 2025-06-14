package types

// AdditionalInfo represents additional information for requests
// Changed from map[string]any to struct to ensure consistent structure
type AdditionalInfo struct {
	DeviceId string `json:"deviceId"`
	Channel  string `json:"channel"`
	// Additional fields can be added here as needed
	CustomField1 string `json:"customField1,omitempty"`
	CustomField2 string `json:"customField2,omitempty"`
	CustomField3 string `json:"customField3,omitempty"`
	// For backward compatibility, we can include a map for other fields
	OtherFields map[string]interface{} `json:"-"`
}

// NewAdditionalInfo creates a new AdditionalInfo instance
func NewAdditionalInfo(deviceId, channel string) *AdditionalInfo {
	return &AdditionalInfo{
		DeviceId: deviceId,
		Channel:  channel,
	}
}

// WithCustomField1 adds CustomField1 to AdditionalInfo
func (a *AdditionalInfo) WithCustomField1(value string) *AdditionalInfo {
	a.CustomField1 = value
	return a
}

// WithCustomField2 adds CustomField2 to AdditionalInfo
func (a *AdditionalInfo) WithCustomField2(value string) *AdditionalInfo {
	a.CustomField2 = value
	return a
}

// WithCustomField3 adds CustomField3 to AdditionalInfo
func (a *AdditionalInfo) WithCustomField3(value string) *AdditionalInfo {
	a.CustomField3 = value
	return a
}

// WithOtherField adds a custom field to OtherFields
func (a *AdditionalInfo) WithOtherField(key string, value interface{}) *AdditionalInfo {
	if a.OtherFields == nil {
		a.OtherFields = make(map[string]interface{})
	}
	a.OtherFields[key] = value
	return a
}