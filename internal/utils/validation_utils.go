package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/errors"
)

// ValidationUtils provides utility functions for validating request payloads
type ValidationUtils struct{}

// NewValidationUtils creates a new ValidationUtils instance
func NewValidationUtils() *ValidationUtils {
	return &ValidationUtils{}
}

// ValidateRequired checks if required fields are present and not empty
func (v *ValidationUtils) ValidateRequired(obj interface{}, fields ...string) error {
	val := reflect.ValueOf(obj)
	
	// If pointer, get the underlying value
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return errors.NewValidationError("object is nil")
		}
		val = val.Elem()
	}
	
	// Must be a struct
	if val.Kind() != reflect.Struct {
		return errors.NewValidationError("object must be a struct")
	}
	
	// Check each required field
	for _, field := range fields {
		f := val.FieldByName(field)
		if !f.IsValid() {
			return errors.NewValidationError(fmt.Sprintf("field '%s' does not exist", field))
		}
		
		// Check if field is empty based on its type
		isEmpty := false
		switch f.Kind() {
		case reflect.String:
			isEmpty = f.String() == ""
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			isEmpty = f.Int() == 0
		case reflect.Float32, reflect.Float64:
			isEmpty = f.Float() == 0
		case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map:
			isEmpty = f.IsNil()
		}
		
		if isEmpty {
			return errors.NewValidationError(fmt.Sprintf("field '%s' is required", field))
		}
	}
	
	return nil
}

// ValidateEmail validates email format
func (v *ValidationUtils) ValidateEmail(email string) error {
	if email == "" {
		return nil // Empty email is allowed (not required)
	}
	
	// Simple email validation regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.NewValidationError(fmt.Sprintf("invalid email format: %s", email))
	}
	
	return nil
}

// ValidatePhone validates phone number format
func (v *ValidationUtils) ValidatePhone(phone string) error {
	if phone == "" {
		return nil // Empty phone is allowed (not required)
	}
	
	// Simple phone validation regex (allows +, spaces, and digits)
	phoneRegex := regexp.MustCompile(`^[+]?[\d\s]+$`)
	if !phoneRegex.MatchString(phone) {
		return errors.NewValidationError(fmt.Sprintf("invalid phone format: %s", phone))
	}
	
	return nil
}

// ValidateDate validates date format (YYYY-MM-DD)
func (v *ValidationUtils) ValidateDate(date string) error {
	if date == "" {
		return nil // Empty date is allowed (not required)
	}
	
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return errors.NewValidationError(fmt.Sprintf("invalid date format: %s, expected YYYY-MM-DD", date))
	}
	
	return nil
}

// ValidateDateTime validates datetime format (ISO 8601)
func (v *ValidationUtils) ValidateDateTime(datetime string) error {
	if datetime == "" {
		return nil // Empty datetime is allowed (not required)
	}
	
	_, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		return errors.NewValidationError(fmt.Sprintf("invalid datetime format: %s, expected ISO 8601 (YYYY-MM-DDThh:mm:ss+07:00)", datetime))
	}
	
	return nil
}

// ValidateAmount validates amount format and range
func (v *ValidationUtils) ValidateAmount(amount string) error {
	if amount == "" {
		return nil // Empty amount is allowed (not required)
	}
	
	// Amount should be in format 12345678.00
	amountRegex := regexp.MustCompile(`^\d+\.\d{2}$`)
	if !amountRegex.MatchString(amount) {
		return errors.NewValidationError(fmt.Sprintf("invalid amount format: %s, expected format with 2 decimal places (e.g., 12345678.00)", amount))
	}
	
	return nil
}

// ValidateCurrency validates currency code (ISO 4217)
func (v *ValidationUtils) ValidateCurrency(currency string) error {
	if currency == "" {
		return nil // Empty currency is allowed (not required)
	}
	
	// Simple validation for ISO 4217 currency codes (3 uppercase letters)
	if len(currency) != 3 || !regexp.MustCompile(`^[A-Z]{3}$`).MatchString(currency) {
		return errors.NewValidationError(fmt.Sprintf("invalid currency code: %s, expected ISO 4217 format (e.g., IDR, USD)", currency))
	}
	
	return nil
}

// ValidateAccountNumber validates account number format
func (v *ValidationUtils) ValidateAccountNumber(accountNo string) error {
	if accountNo == "" {
		return nil // Empty account number is allowed (not required)
	}
	
	// Account number should be digits only
	if !regexp.MustCompile(`^\d+$`).MatchString(accountNo) {
		return errors.NewValidationError(fmt.Sprintf("invalid account number format: %s, expected digits only", accountNo))
	}
	
	return nil
}

// ValidateBankCode validates bank code format
func (v *ValidationUtils) ValidateBankCode(bankCode string) error {
	if bankCode == "" {
		return nil // Empty bank code is allowed (not required)
	}
	
	// Bank code should be 3 digits for most Indonesian banks
	if !regexp.MustCompile(`^\d{3}$`).MatchString(bankCode) {
		return errors.NewValidationError(fmt.Sprintf("invalid bank code format: %s, expected 3 digits", bankCode))
	}
	
	return nil
}

// ValidatePayload performs comprehensive validation on a request payload
func (v *ValidationUtils) ValidatePayload(payload interface{}) error {
	// Get payload type and value
	val := reflect.ValueOf(payload)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return errors.NewValidationError("payload is nil")
		}
		val = val.Elem()
	}
	
	// Must be a struct
	if val.Kind() != reflect.Struct {
		return errors.NewValidationError("payload must be a struct")
	}
	
	// Get struct type
	typ := val.Type()
	
	// Check each field for validation tags
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		
		// Skip unexported fields
		if !fieldType.IsExported() {
			continue
		}
		
		// Get validation tag
		tag := fieldType.Tag.Get("validate")
		if tag == "" {
			continue
		}
		
		// Process validation rules
		rules := strings.Split(tag, ",")
		for _, rule := range rules {
			// Check if field is required
			if rule == "required" && isEmptyValue(field) {
				return errors.NewValidationError(fmt.Sprintf("field '%s' is required", fieldType.Name))
			}
			
			// Skip other validations if field is empty and not required
			if isEmptyValue(field) {
				continue
			}
			
			// Apply other validation rules based on field type
			switch field.Kind() {
			case reflect.String:
				strValue := field.String()
				
				// Email validation
				if rule == "email" {
					if err := v.ValidateEmail(strValue); err != nil {
						return err
					}
				}
				
				// Phone validation
				if rule == "phone" {
					if err := v.ValidatePhone(strValue); err != nil {
						return err
					}
				}
				
				// Date validation
				if rule == "date" {
					if err := v.ValidateDate(strValue); err != nil {
						return err
					}
				}
				
				// DateTime validation
				if rule == "datetime" {
					if err := v.ValidateDateTime(strValue); err != nil {
						return err
					}
				}
				
				// Amount validation
				if rule == "amount" {
					if err := v.ValidateAmount(strValue); err != nil {
						return err
					}
				}
				
				// Currency validation
				if rule == "currency" {
					if err := v.ValidateCurrency(strValue); err != nil {
						return err
					}
				}
				
				// Account number validation
				if rule == "accountno" {
					if err := v.ValidateAccountNumber(strValue); err != nil {
						return err
					}
				}
				
				// Bank code validation
				if rule == "bankcode" {
					if err := v.ValidateBankCode(strValue); err != nil {
						return err
					}
				}
			}
		}
	}
	
	return nil
}

// isEmptyValue checks if a value is empty based on its type
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map:
		return v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	default:
		return false
	}
}