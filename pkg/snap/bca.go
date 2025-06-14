package snap

import (
	"context"
	"fmt"

	"github.com/mseptiaan/snap-aspi-go/internal/auth"
	"github.com/mseptiaan/snap-aspi-go/internal/config"
	"github.com/mseptiaan/snap-aspi-go/internal/logging"
	"github.com/mseptiaan/snap-aspi-go/pkg/client"
	"github.com/mseptiaan/snap-aspi-go/pkg/contracts"
	"github.com/mseptiaan/snap-aspi-go/pkg/services"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

// BCAClient represents a specialized client for BCA SNAP integration
type BCAClient struct {
	*Client
	bcaEndpoints *BCAEndpoints
}

// BCAEndpoints contains BCA-specific endpoints
type BCAEndpoints struct {
	VirtualAccount *BCAVirtualAccountEndpoints
	Transfer       *BCATransferEndpoints
	Balance        *BCABalanceEndpoints
	QR             *BCAQREndpoints
}

// BCAVirtualAccountEndpoints contains BCA-specific VA endpoints
type BCAVirtualAccountEndpoints struct {
	Create  string
	Inquiry string
	Update  string
	Payment string
}

// BCATransferEndpoints contains BCA-specific transfer endpoints
type BCATransferEndpoints struct {
	AccountInquiry  string
	TriggerTransfer string
	StatusInquiry   string
}

// BCABalanceEndpoints contains BCA-specific balance endpoints
type BCABalanceEndpoints struct {
	BalanceInquiry string
}

// BCAQREndpoints contains BCA-specific QR endpoints
type BCAQREndpoints struct {
	GenerateQR string
	DecodeQR   string
}

// NewBCAClient creates a new BCA-specific client
func NewBCAClient(cfg Config) (*BCAClient, error) {
	// Set BCA-specific endpoints
	if cfg.CustomEndpoints == nil {
		cfg.CustomEndpoints = &CustomEndpoints{}
	}

	// Apply BCA-specific endpoint overrides
	bcaEndpoints := getBCAEndpoints()
	applyBCAEndpoints(cfg.CustomEndpoints, bcaEndpoints)

	// Create base client
	baseClient, err := NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create base client: %w", err)
	}

	return &BCAClient{
		Client:       baseClient,
		bcaEndpoints: bcaEndpoints,
	}, nil
}

// getBCAEndpoints returns BCA-specific endpoints
func getBCAEndpoints() *BCAEndpoints {
	return &BCAEndpoints{
		VirtualAccount: &BCAVirtualAccountEndpoints{
			Create:   "/api/v1.0/bca/transfer-va/create-va",
			Inquiry:  "/api/v1.0/bca/transfer-va/inquiry-va",
			Update:   "/api/v1.0/bca/transfer-va/update-va",
			Payment:  "/api/v1.0/bca/transfer-va/payment",
		},
		Transfer: &BCATransferEndpoints{
			AccountInquiry:  "/api/v1.0/bca/account-inquiry-external",
			TriggerTransfer: "/api/v1.0/bca/trigger-transfer",
			StatusInquiry:   "/api/v1.0/bca/transfer/status",
		},
		Balance: &BCABalanceEndpoints{
			BalanceInquiry: "/api/v1.0/bca/balance-inquiry",
		},
		QR: &BCAQREndpoints{
			GenerateQR: "/api/v1.0/bca/qr/qr-mpm-generate",
			DecodeQR:   "/api/v1.0/bca/qr/qr-mpm-decode",
		},
	}
}

// applyBCAEndpoints applies BCA-specific endpoints to custom endpoints
func applyBCAEndpoints(customEndpoints *CustomEndpoints, bcaEndpoints *BCAEndpoints) {
	// Apply Virtual Account endpoints
	if customEndpoints.VirtualAccount == nil {
		customEndpoints.VirtualAccount = &VirtualAccountEndpoints{}
	}
	if bcaEndpoints.VirtualAccount != nil {
		if customEndpoints.VirtualAccount.CreateVA == "" {
			customEndpoints.VirtualAccount.CreateVA = bcaEndpoints.VirtualAccount.Create
		}
		if customEndpoints.VirtualAccount.InquiryVA == "" {
			customEndpoints.VirtualAccount.InquiryVA = bcaEndpoints.VirtualAccount.Inquiry
		}
		if customEndpoints.VirtualAccount.UpdateVA == "" {
			customEndpoints.VirtualAccount.UpdateVA = bcaEndpoints.VirtualAccount.Update
		}
		if customEndpoints.VirtualAccount.Payment == "" {
			customEndpoints.VirtualAccount.Payment = bcaEndpoints.VirtualAccount.Payment
		}
	}

	// Apply Transfer endpoints
	if customEndpoints.TransferCredit == nil {
		customEndpoints.TransferCredit = &TransferCreditEndpoints{}
	}
	if bcaEndpoints.Transfer != nil {
		if customEndpoints.TransferCredit.AccountInquiry == "" {
			customEndpoints.TransferCredit.AccountInquiry = bcaEndpoints.Transfer.AccountInquiry
		}
		if customEndpoints.TransferCredit.TriggerTransfer == "" {
			customEndpoints.TransferCredit.TriggerTransfer = bcaEndpoints.Transfer.TriggerTransfer
		}
		if customEndpoints.TransferCredit.TransferStatus == "" {
			customEndpoints.TransferCredit.TransferStatus = bcaEndpoints.Transfer.StatusInquiry
		}
	}

	// Apply Balance endpoints
	if customEndpoints.BalanceInquiry == nil {
		customEndpoints.BalanceInquiry = &BalanceInquiryEndpoints{}
	}
	if bcaEndpoints.Balance != nil {
		if customEndpoints.BalanceInquiry.BalanceInquiry == "" {
			customEndpoints.BalanceInquiry.BalanceInquiry = bcaEndpoints.Balance.BalanceInquiry
		}
	}

	// Apply QR endpoints
	if customEndpoints.MPM == nil {
		customEndpoints.MPM = &MPMEndpoints{}
	}
	if bcaEndpoints.QR != nil {
		if customEndpoints.MPM.GenerateQR == "" {
			customEndpoints.MPM.GenerateQR = bcaEndpoints.QR.GenerateQR
		}
		if customEndpoints.MPM.DecodeQR == "" {
			customEndpoints.MPM.DecodeQR = bcaEndpoints.QR.DecodeQR
		}
	}
}

// CreateBCAVirtualAccount creates a BCA-specific virtual account
func (c *BCAClient) CreateBCAVirtualAccount(ctx context.Context, payload *types.BCACreateVAPayload) (*types.BCACreateVAResponse, error) {
	// Convert BCA-specific payload to standard payload
	standardPayload := &types.CreateVAPayload{
		PartnerServiceId:    payload.PartnerServiceId,
		CustomerNo:          payload.CustomerNo,
		VirtualAccountNo:    payload.VirtualAccountNo,
		VirtualAccountName:  payload.VirtualAccountName,
		VirtualAccountEmail: payload.VirtualAccountEmail,
		VirtualAccountPhone: payload.VirtualAccountPhone,
		TrxId:               payload.TrxId,
		TotalAmount:         payload.TotalAmount,
		ExpiredDate:         payload.ExpiredDate,
		AdditionalInfo:      payload.AdditionalInfo,
	}

	// Call standard CreateVA method
	result, err := c.VirtualAccount().CreateVA(ctx, standardPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to create BCA virtual account: %w", err)
	}

	// Convert standard response to BCA-specific response
	response := &types.BCACreateVAResponse{
		ResponseCode:    result["responseCode"].(string),
		ResponseMessage: result["responseMessage"].(string),
		ReferenceNo:     result["referenceNo"].(string),
		VirtualAccountNo: result["virtualAccountNo"].(string),
	}

	// Add BCA-specific fields if present
	if val, ok := result["bcaReferenceNo"]; ok {
		response.BcaReferenceNo = val.(string)
	}

	return response, nil
}

// InquiryBCAVirtualAccount performs a BCA-specific virtual account inquiry
func (c *BCAClient) InquiryBCAVirtualAccount(ctx context.Context, payload *types.BCAInquiryVAPayload) (*types.BCAInquiryVAResponse, error) {
	// Convert BCA-specific payload to standard payload
	standardPayload := &types.InquiryVAPayload{
		PartnerServiceId: payload.PartnerServiceId,
		CustomerNo:       payload.CustomerNo,
		VirtualAccountNo: payload.VirtualAccountNo,
		TrxDateInit:      payload.TrxDateInit,
	}

	// Call standard InquiryVA method
	result, err := c.VirtualAccount().InquiryVA(ctx, standardPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to inquiry BCA virtual account: %w", err)
	}

	// Convert standard response to BCA-specific response
	response := &types.BCAInquiryVAResponse{
		ResponseCode:       result["responseCode"].(string),
		ResponseMessage:    result["responseMessage"].(string),
		VirtualAccountNo:   result["virtualAccountNo"].(string),
		VirtualAccountName: result["virtualAccountName"].(string),
		TotalAmount:        result["totalAmount"].(map[string]interface{}),
		ExpiredDate:        result["expiredDate"].(string),
	}

	// Add BCA-specific fields if present
	if val, ok := result["bcaReferenceNo"]; ok {
		response.BcaReferenceNo = val.(string)
	}

	return response, nil
}

// BCABalanceInquiry performs a BCA-specific balance inquiry
func (c *BCAClient) BCABalanceInquiry(ctx context.Context, payload *types.BCABalanceInquiryPayload) (*types.BCABalanceInquiryResponse, error) {
	// Convert BCA-specific payload to standard payload
	standardPayload := &types.BalanceInquiryPayload{
		PartnerServiceId: payload.PartnerServiceId,
		CustomerNo:       payload.CustomerNo,
		AccountNo:        payload.AccountNo,
		BalanceTypes:     payload.BalanceTypes,
		AdditionalInfo:   payload.AdditionalInfo,
	}

	// Call standard BalanceInquiry method
	result, err := c.BalanceInquiry().BalanceInquiry(ctx, standardPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to perform BCA balance inquiry: %w", err)
	}

	// Convert to BCA-specific response
	response := &types.BCABalanceInquiryResponse{
		ResponseCode:    result.ResponseCode,
		ResponseMessage: result.ResponseMessage,
		ReferenceNo:     result.ReferenceNo,
		AccountNo:       result.AccountNo,
		Name:            result.Name,
		AccountInfos:    result.AccountInfos,
	}

	return response, nil
}

// GenerateBCAQR generates a BCA-specific QR code
func (c *BCAClient) GenerateBCAQR(ctx context.Context, payload *types.BCAGenerateQRPayload) (*types.BCAQRResponse, error) {
	// Convert BCA-specific payload to standard payload
	standardPayload := &types.MPMGenerateQRPayload{
		PartnerServiceId:   payload.PartnerServiceId,
		CustomerNo:         payload.CustomerNo,
		PartnerReferenceNo: payload.PartnerReferenceNo,
		Amount:             payload.Amount,
		MerchantId:         payload.MerchantId,
		MerchantName:       payload.MerchantName,
		ValidityPeriod:     payload.ValidityPeriod,
		AdditionalInfo:     payload.AdditionalInfo,
	}

	// Call standard GenerateQR method
	result, err := c.MPM().GenerateQR(ctx, standardPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to generate BCA QR: %w", err)
	}

	// Convert to BCA-specific response
	response := &types.BCAQRResponse{
		ResponseCode:    result.ResponseCode,
		ResponseMessage: result.ResponseMessage,
		ReferenceNo:     result.ReferenceNo,
		QRContent:       result.QRContent,
		QRUrl:           result.QRUrl,
		ValidUntil:      result.ValidUntil,
	}

	// Add BCA-specific fields if present
	if result.QRImage != "" {
		response.QRImage = result.QRImage
	}

	return response, nil
}

// BCATransfer performs a BCA-specific transfer
func (c *BCAClient) BCATransfer(ctx context.Context, payload *types.BCATransferPayload) (*types.BCATransferResponse, error) {
	// Convert BCA-specific payload to standard payload
	standardPayload := &types.TriggerTransferPayload{
		PartnerServiceId:       payload.PartnerServiceId,
		CustomerNo:             payload.CustomerNo,
		PartnerReferenceNo:     payload.PartnerReferenceNo,
		Amount:                 payload.Amount,
		BeneficiaryAccountNo:   payload.BeneficiaryAccountNo,
		BeneficiaryAccountName: payload.BeneficiaryAccountName,
		BeneficiaryBankCode:    payload.BeneficiaryBankCode,
		SourceAccountNo:        payload.SourceAccountNo,
		TransactionDate:        payload.TransactionDate,
		FeeType:                payload.FeeType,
		Remark:                 payload.Remark,
		AdditionalInfo:         payload.AdditionalInfo,
	}

	// Call standard TriggerTransfer method
	result, err := c.TransferCredit().TriggerTransfer(ctx, standardPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to perform BCA transfer: %w", err)
	}

	// Convert to BCA-specific response
	response := &types.BCATransferResponse{
		ResponseCode:    result["responseCode"].(string),
		ResponseMessage: result["responseMessage"].(string),
		ReferenceNo:     result["referenceNo"].(string),
	}

	// Add BCA-specific fields if present
	if val, ok := result["bcaReferenceNo"]; ok {
		response.BcaReferenceNo = val.(string)
	}
	if val, ok := result["transactionDate"]; ok {
		response.TransactionDate = val.(string)
	}

	return response, nil
}