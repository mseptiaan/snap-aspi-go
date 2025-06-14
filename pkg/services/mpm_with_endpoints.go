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

// NewMPMServiceWithEndpoints creates a new MPM service with custom endpoints
func NewMPMServiceWithEndpoints(
	httpClient contracts.HTTPClient,
	authManager *auth.AccessTokenManager,
	cfg *config.Config,
	logger logging.Logger,
	customEndpoints map[string]string,
) (*MPMService, error) {
	// Create symmetric signer for API signatures
	symSigner := signature.NewSymmetricSigner(cfg.ASPI.ClientSecret)

	// Use custom endpoints if provided, otherwise use defaults
	endpoints := getMPMEndpoints(customEndpoints)
	methods := getMPMMethods()

	return &MPMService{
		httpClient:  httpClient,
		authManager: authManager,
		symSigner:   symSigner,
		config:      cfg,
		logger:      logger,
		endpoints:   endpoints,
		methods:     methods,
	}, nil
}

// getMPMEndpoints returns MPM endpoints (custom or default)
func getMPMEndpoints(customEndpoints map[string]string) map[string]string {
	// Default endpoints
	defaultEndpoints := map[string]string{
		"transfer":        "/api/v1.0/transfer-kredit/mpm",
		"inquiry":         "/api/v1.0/transfer-kredit/mpm/inquiry",
		"status":          "/api/v1.0/transfer-kredit/mpm/status",
		"refund":          "/api/v1.0/transfer-kredit/mpm/refund",
		"balance-inquiry": "/api/v1.0/transfer-kredit/mpm/balance-inquiry",
		"account-inquiry": "/api/v1.0/transfer-kredit/mpm/account-inquiry",
		"history":         "/api/v1.0/transfer-kredit/mpm/history",
		"generate-qr":     "/api/v1.0/qr/qr-mpm-generate",
		"notify-qr":       "/api/v1.0/qr/qr-mpm-notify",
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

// getMPMMethods returns HTTP methods for MPM endpoints
func getMPMMethods() map[string]string {
	return map[string]string{
		"transfer":        "POST",
		"inquiry":         "POST",
		"status":          "POST",
		"refund":          "POST",
		"balance-inquiry": "POST",
		"account-inquiry": "POST",
		"history":         "POST",
		"generate-qr":     "POST",
		"notify-qr":       "POST",
	}
}