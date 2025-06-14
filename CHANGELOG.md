# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2024-07-15

### Added
- Added comprehensive support for all ASPI endpoints
- Added Registration service for user registration, card registration, and account binding
- Added Balance Inquiry service for account balance information
- Added Transaction History service for transaction history and bank statements
- Added Transfer Credit service for account inquiry, transfers, and top-ups
- Added Transfer Debit service for direct debit, CPM, auth payment, and BI-FAST
- Added detailed documentation for all services and endpoints
- Added examples for all services in the examples directory

### Changed
- Updated MPM service to align with ASPI documentation
- Fixed endpoint paths to match ASPI specifications
- Improved request/response structures to match ASPI documentation
- Enhanced error handling with more specific error messages
- Updated bank presets with comprehensive endpoint configurations

### Fixed
- Fixed MPM QR response structure to match ASPI documentation
- Corrected refund endpoint path in MPM service
- Fixed missing endpoints in MPM service (decode-qr, apply-ott, payment, query-payment, cancel-payment)
- Improved field naming consistency across all request/response structures

## [1.0.0] - 2024-06-30

### Added
- Initial release with basic Virtual Account and MPM functionality
- Support for authentication with B2B and B2B2C tokens
- Connection pooling and request optimization
- Retry logic with exponential backoff
- Multi-bank support with bank presets
- Custom endpoint configuration