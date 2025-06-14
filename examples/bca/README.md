# BCA SNAP Integration Examples

This directory contains examples for integrating with BCA (Bank Central Asia) using the SNAP ASPI Go SDK. These examples demonstrate how to use the BCA-specific client and types to interact with BCA's implementation of the ASPI standard.

## Prerequisites

Before running these examples, make sure you have:

1. BCA API credentials (Client ID, Client Secret)
2. RSA key pair for signing requests
3. Go 1.23 or later installed

## Environment Setup

Set up the following environment variables:

```bash
export ASPI_BASE_URL="https://sandbox.aspi-indonesia.or.id"
export BCA_CLIENT_ID="your-bca-client-id"
export BCA_CLIENT_SECRET="your-bca-client-secret"
export ASPI_PRIVATE_KEY_PATH="path/to/private_key.pem"
export ASPI_PUBLIC_KEY_PATH="path/to/public_key.pem"
```

## Running the Examples

```bash
go run main.go
```

## Examples Included

The `main.go` file demonstrates the following BCA-specific operations:

1. **Create BCA Virtual Account**: Creates a virtual account with BCA-specific parameters
2. **Inquiry BCA Virtual Account**: Retrieves information about an existing BCA virtual account
3. **BCA Balance Inquiry**: Checks the balance of a BCA account
4. **Generate BCA QR Code**: Generates a QR code for BCA payments
5. **BCA Transfer**: Performs a transfer using BCA's transfer service

## BCA-Specific Features

The BCA integration includes several BCA-specific features:

- **BCA-specific endpoints**: Uses BCA's specific API endpoints
- **BCA-specific request parameters**: Includes BCA-specific fields in requests
- **BCA-specific response handling**: Processes BCA-specific response fields
- **BCA-specific headers**: Adds BCA-specific headers to requests

## Additional Resources

- [BCA API Documentation](https://developer.bca.co.id/)
- [ASPI Documentation](https://aspi-indonesia.or.id/documentation)
- [SNAP ASPI Go SDK Documentation](https://github.com/mseptiaan/snap-aspi-go)

## Notes

- These examples use the sandbox environment. For production use, change the `ASPI_BASE_URL` to the production URL.
- The examples use hardcoded values for demonstration purposes. In a real application, you should use dynamic values.
- Error handling is minimal in these examples for clarity. In a production environment, implement comprehensive error handling.