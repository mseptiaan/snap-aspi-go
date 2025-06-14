package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/mseptiaan/snap-aspi-go/internal/auth"
	"github.com/mseptiaan/snap-aspi-go/internal/config"
	"github.com/mseptiaan/snap-aspi-go/internal/errors"
	"github.com/mseptiaan/snap-aspi-go/internal/logging"
	"github.com/mseptiaan/snap-aspi-go/internal/signature"
	"github.com/mseptiaan/snap-aspi-go/pkg/contracts"
	"github.com/mseptiaan/snap-aspi-go/pkg/types"
)

// NewVirtualAccountServiceWithEndpoints creates a new VirtualAccount service with custom endpoints
func NewVirtualAccountServiceWithEndpoints(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
	customEndpoints map[string]string,
) (*VirtualAccountService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	// Use custom endpoints if provided, otherwise use defaults
	endpoints := getVAEndpoints(customEndpoints)
	methods := getVAMethods()

	return &VirtualAccountService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		endpoints:   endpoints,
		methods:     methods,
	}, nil
}

// getVAEndpoints returns VA endpoints (custom or default)
func getVAEndpoints(customEndpoints map[string]string) map[string]string {
	// Default endpoints
	defaultEndpoints := map[string]string{
		"inquiry":       "/api/v1.0/transfer-va/inquiry",
		"inquiry-va":    "/api/v1.0/transfer-va/inquiry-va",
		"create-va":     "/api/v1.0/transfer-va/create-va",
		"update-va":     "/api/v1.0/transfer-va/update-va",
		"delete-va":     "/api/v1.0/transfer-va/delete-va",
		"payment":       "/api/v1.0/transfer-va/payment",
		"status":        "/api/v1.0/transfer-va/status",
		"report":        "/api/v1.0/transfer-va/report",
		"update-status": "/api/v1.0/transfer-va/update-status",
	}

	// If custom endpoints provided, merge them
	if customEndpoints != nil {
		for key, endpoint := range customEndpoints {
			if endpoint != "" {
				defaultEndpoints[key] = endpoint
			}
		}
	}

	return defaultEndpoints
}

// getVAMethods returns HTTP methods for VA endpoints
func getVAMethods() map[string]string {
	return map[string]string{
		"inquiry":       "POST",
		"inquiry-va":    "POST",
		"create-va":     "POST",
		"update-va":     "PUT",
		"delete-va":     "DELETE",
		"payment":       "POST",
		"status":        "POST",
		"report":        "POST",
		"update-status": "PUT",
	}
}