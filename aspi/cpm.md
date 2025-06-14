CPM - SNAP Developer Site

[![](/img/icon/snap-header-white-100.png)](/)

- [**Beranda**](/)
- [**Info**](/info)
- [**API**](/api-services)
- [**Direktori Publikasi**](/direktori-publikasi)
  - [**Sign In**](/sign-in)

[![](/img/icon/bi-logo-full-100.png)](/)

![](/img/card/transfer-debit-100.png)

### [Transfer Debit](/api-services/transfer-debit)

##### CPM

Overview


Guides


Code Snippets


Response Code


Aplikasi Pengujian


##### Transfer Debit

- [Direct Debit](/api-services/transfer-debit/direct-debit)
- [CPM](/api-services/transfer-debit/cpm)
- [Auth Payment](/api-services/transfer-debit/auth-payment)
- [Direct Debit BI-FAST](/api-services/transfer-debit/direct-debit-bi-fast)

#### OVERVIEW

#### API QR CPM

![](/img/docs/5_6_sequence_diagram_api_qr_cpm.png)

Sequence Diagram API QR CPM

##### API Generate QR CPM

**Informasi Umum**

Service Code59NameAPI Generate QR CPMVersion1.0HTTP MethodPOSTPath.../{version}/qr/qr-cpm-generate

##### API CPM Payment

**Informasi Umum**

Service Code60NameAPI CPM PaymentVersion1.0HTTP MethodPOSTPath.../{version}/qr/qr-cpm-payment

##### API Query Payment

**Informasi Umum**

Service Code61NameAPI Query PaymentVersion1.0HTTP MethodPOSTPath.../{version}/qr/qr-cpm-query

##### API Cancel Payment

**Informasi Umum**

Service Code62NameAPI Cancel PaymentVersion1.0HTTP MethodPOSTPath.../{version}/qr/qr-cpm-cancel

##### API Payment Notification

**Informasi Umum**

Service Code79NameAPI Payment NotificationVersion1.0HTTP MethodPOSTPath.../{version}/qr/qr-cpm-notify

##### API Refund Payment

**Informasi Umum**

Service Code80NameAPI Refund PaymentVersion1.0HTTP MethodPOSTPath.../{version}/qr/qr-cpm-refund

#### GUIDES

#### Spesifikasi Parameter Header dan Body API QR CPM

##### API Generate QR CPM

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringO64Transaction identifier on service consumer systemuserAccessTokenStringO64User Access TokenmerchantIdStringO64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDpartnerTrxDatestringM25Partner transaction Date : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringO64Transaction identifier on service provider systempartnerReferenceNoStringO64Transaction identifier on service consumer systemqrContentStringO512QR String CPMqrUrlStringO255URL to show QR in webexpiryTimeStringM25Format expiry time : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API CPM Payment

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringM64Transaction identifier on service consumer systemqrContentStringM512QR String CPMamountObjectOvalueStringM16,2Transaction AmountcurrencyStringM3Currency (ISO4217)feeAmountObjectOvalueStringM16,2Transaction fee.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)merchantIdStringM64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDtitleStringO256Brief descriptionexpiryTimeStringO25Expiry time. Format expiry time: (ISO 8601) YYYY-MM-DDThh:mm:ssitemsObjectOFor itemsexternalStoreIdStringO64unique shop id in merchant side.merchantNameStringO64Merchant namemerchantLocationStringO64Merchant locationacquirerNameStringO64Acquire NameterminalIdStringO32Terminal IdscannerInfoObjectOdeviceIdStringO64The unique device id.deviceVersionStringO128The device firmware version, updatable.deviceModelStringO128The device model, not updateable.deviceIpStringO64The Public IP address when device sends out the request.additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionpartnerReferenceNoStringO64Transaction identifier on service consumer systemtransactionDateStringO25Format transaction date : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Query Payment

**Request Body**

ParameterData TypeMandatoryLengthDescriptionoriginalReferenceNoStringO64Original transaction identifier on service provider systemoriginalPartnerReferenceNoStringO64Original transaction identifier on service consumer systemoriginalExternalIdStringO32Original CustomerReference NumbermerchantIdStringO64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDexternalStoreIdStringO64External Store IDadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionoriginalReferenceNoStringC64Original transaction identifier on service provider systemoriginalPartnerReferenceNoStringO64Original transaction identifier on service consumer systemoriginalExternalIdStringO32Original CustomerReference NumbertitlestringO256Title querylatestTransactionStatusstringM200 - Success01 - Initiated02 - Paying03 - Pending04 - Refunded05 - Canceled06 - Failed07 - Not foundtransactionStatusDescStringO50Description status transactionpaidTimeStringM25Format paid time : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Cancel Payment

**Request Body**

ParameterData TypeMandatoryLengthDescriptionoriginalPartnerReferenceNoStringM64Original transaction identifier on service consumer systemoriginalReferenceNoStringO64Original transaction identifier on service provider systemoriginalExternalIdStringO32Original External-ID on header messagemerchantIdStringO64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDexternalStoreIdStringO64External Store IDamountObjectOvalueStringM16,2Transaction AmountcurrencyStringM3Currency (ISO4217)reasonstringO256Reason cancellationadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionoriginalPartnerReferenceNoStringO64Original transaction identifier on service consumer systemoriginalReferenceNoStringC64Original transaction identifier on service provider systemoriginalExternalIdStringO32Original External-ID on header messagecancelTimeStringC25Cancel time Must be filled if cancelled transaction successFormat Cancel time : (ISO 8601) YYYY-MM-DDThh:mm:sstransactionDateStringO25Format transaction date : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Payment Notification

**Request Body**

ParameterData TypeMandatoryLengthDescriptionoriginalPartnerReferenceNoStringO64Original transaction identifier on service consumer systemoriginalReferenceNoStringO64Original transaction identifier on service provider systemmerchantIdStringM64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDexternalStoreIdStringO64External Store IDamountObjectOvalueStringM16,2Net amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)latestTransactionStatusstringM200 - Success01 - Initiated02 - Paying03 - Pending04 - Refunded05 - Canceled06 - Failed07 - Not foundtransactionStatusDescStringO50Description status transactioncustomerNumberStringO64Customer Account NumberaccountTypestringO25Account typedestinationNumberstringO25Destination account numberdestinationAccountNamestringO25Destination account namesessionIdstringO25Session idbankCodestringO8Bank codeadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Refund Payment

**Request Body**

ParameterData TypeMandatoryLengthDescriptionmerchantIdStringO64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDexternalStoreIdStringO64External Store IDoriginalPartnerReferenceNoStringM64Original transaction identifier on service consumer systemoriginalReferenceNoStringO64Original transaction identifier on service provider systemoriginalExternalIdStringO32Original CustomerReference NumberpartnerRefundNoStringM64ReferenceNumber from PJP AIS for the refund.refundAmountObjectOvalueStringM16,2Net amount of the refund.currencyStringM3Currency (ISO4217)reasonStringO256Refund reason.additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionoriginalPartnerReferenceNoStringO64Original transaction identifier on service consumer systemoriginalReferenceNoStringO64Original transaction identifier on service provider systemoriginalExternalIdStringO32Original CustomerReference NumberrefundNoStringM64referenceNumberpartnerRefundNoStringO64ReferenceNumber from PJP AIS for the refund.refundAmountObjectOvalueStringM16,2Net amount of the refund.currencyStringM3Currency (ISO4217)refundTimeStringM25Refund time.ISO 8601additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

#### CODE SNIPPETS

#### Code Snippets API QR CPM

##### API Generate QR CPM

**Sample Request**

```http
POST .../v1.0/qr/qr-cpm-generate HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T07:50:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.108841
X-LONGITUDE: 106.7782137
CHANNEL-ID: 95221

{
   "partnerReferenceNo":"2020102900000000000001",
   "userAccessToken":"xxxxxx",
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "partnerTrxDate":"2020-12-23T07:50:11+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-23T07:50:19+07:00

{
   "responseCode":"2005900",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "qrContent":"hduvY...",
   "qrUrl":"PL0001",
   "expiryTime":"2020-12-23T07:50:11+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API CPM Payment

**Sample Request**

```http
POST .../v1.0/qr/qr-cpm-payment HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T07:54:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.108841
X-LONGITUDE: 106.7782137
CHANNEL-ID: 95221

{
   "partnerReferenceNo":"2020102900000000000001",
   "qrContent":"hduvY...",
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "feeAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "title":"example",
   "expiryTime":"2020-12-21T17:21:41+07:00",
   "items":{
      "productId":"12345",
      "productName":"goods A",
      "qty":"6",
      "desc":"barang"
   },
   "externalStoreId":"abcd",
   "merchantName":"Merchant01",
   "merchantLocation":"Jakarta Indonesia",
   "acquirerName":"LinkAja",
   "terminalId":"6476378",
   "scannerInfo":{
      "deviceId":"46252",
      "deviceVersion":"v.1.0",
      "deviceModel":"Scanner",
      "deviceIp":"172.24.281.24"
   },
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-23T07:54:19+07:00

{
   "responseCode":"2006000",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "transactionDate":"2020-12-23T07:54:11+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Query Payment

**Sample Request**

```http
POST .../v1.0/qr/qr-cpm-query HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T08:04:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.108841
X-LONGITUDE: 106.7782137
CHANNEL-ID: 95221

{
   "originalReferenceNo":"2020102977770000000009",
   "originalPartnerReferenceNo":"2020102900000000000001",
   "originalExternalId":"10052019",
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "externalStoreId":"124928924949487",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-23T08:04:18+07:00

{
   "responseCode":"2006100",
   "responseMessage":"Request has been processed successfully",
   "originalReferenceNo":"2020102977770000000009",
   "originalPartnerReferenceNo":"2020102900000000000001",
   "originalExternalId":"10052019",
   "title":"example",
   "latestTransactionStatus":"00",
   "transactionStatusDesc":"success",
   "paidTime":"2020-12-23T08:04:11+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Cancel Payment

**Sample Request**

```http
POST .../v1.0/qr/qr-cpm-cancel HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T08:10:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.108841
X-LONGITUDE: 106.7782137
CHANNEL-ID: 95221

{
   "originalPartnerReferenceNo":"2020102900000000000001",
   "originalReferenceNo":"2020102977770000000009",
   "originalExternalId":"30443786930722726463280097920912",
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "externalStoreId":"124928924949487",
   "amount":{
      "value":"10000.00",
      "currency":"IDR"
   },
   "reason":"cancel reason",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-23T08:10:18+07:00

{
   "responseCode":"2006200",
   "responseMessage":"Request has been processed successfully",
   "originalPartnerReferenceNo":"2020102900000000000001",
   "originalReferenceNo":"2020102977770000000009",
   "originalExternalId":"30443786930722726463280097920912",
   "cancelTime":"2020-12-21T17:07:25+07:00",
   "transactionDate":"2020-12-21T17:55:11+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Payment Notification

**Sample Request**

```http
POST .../v1.0/qr/qr-cpm-notify HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T08:46:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.108841
X-LONGITUDE: 106.7782137
CHANNEL-ID: 95221

{
   "originalPartnerReferenceNo":"2020102900000000000001",
   "originalReferenceNo":"2020102977770000000009",
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "externalStoreId":"124928924949487",
   "amount":{
      "value":"50000.00",
      "currency":"IDR"
   },
   "latestTransactionStatus":"00",
   "transactionStatusDesc":"success",
   "customerNumber":"6281388370001",
   "accountType":"tabungan",
   "destinationNumber":"8377388292",
   "destinationAccountName":"John Doe",
   "sessionId":"0UYEB77329002HY",
   "bankCode":"002",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-23T08:46:16+07:00

{
   "responseCode":"2007900",
   "responseMessage":"Request has been processed successfully",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Refund Payment

**Sample Request**

```http
POST .../v1.0/qr/qr-cpm-refund HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T08:10:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.108841
X-LONGITUDE: 106.7782137
CHANNEL-ID: 95221

{
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "externalStoreId":"124928924949487",
   "originalPartnerReferenceNo":"2020102900000000000001",
   "originalReferenceNo":"2020102977770000000009",
   "originalExternalId":"10052019",
   "partnerRefundNo":"239850918204981205970",
   "refundAmount":{
      "value":"10000.00",
      "currency":"IDR"
   },
   "reason":"Customer complain",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-23T08:10:18+07:00

{
   "responseCode":"2008000",
   "responseMessage":"Request has been processed successfully",
   "originalPartnerReferenceNo":"2020102900000000000001",
   "originalReferenceNo":"2020102977770000000009",
   "originalExternalId":"10052019",
   "refundNo":"REF993883",
   "partnerRefundNo":"239850918204981205970",
   "refundAmount":{
      "value":"10000.00",
      "currency":"IDR"
   },
   "refundTime":"2020-12-21T17:21:41+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

#### RESPONSES CODE

Response status merupakan informasi yang diberikan oleh service provider kepada service consumer pada response body, sebagai indikasi hasil dari pemrosesan request yang diterima.

Response status terdiri dari 2 komponen, yaitu kode (response code) dan deskripsinya (response message).

KomponenTipe DataLengthKeteranganresponseCodeString7response code = HTTP status code + service code + case coderesponseMessageString150

###### **Daftar** _**Response Code**_

CategoryHTTP CodeService CodeCase CodeResponse MessageDescriptionSuccess200any00SuccessfulSuccessfulSuccess202any00Request In ProgressTransaction still on processSystem400any00Bad RequestGeneral request failed error, including message parsing failed.Message400any01Invalid Field Format {field name}Invalid formatMessage400any02Invalid Mandatory Field {field name}Missing or invalid format on mandatory fieldSystem401any00Unauthorized. \[reason\]General unauthorized error (No Interface Def, API is Invalid, Oauth Failed, Verify Client Secret Fail, Client Forbidden Access API, Unknown Client, Key not Found)System401any01Invalid Token (B2B)Token found in request is invalid (Access Token Not Exist, Access Token Expiry)System401any02Invalid Customer TokenToken found in request is invalid (Access Token Not Exist, Access Token Expiry)System401any03Token Not Found (B2B)Token not found in the system. This occurs on any API that requires token as input parameterSystem401any04Customer Token Not FoundToken not found in the system. This occurs on any API that requires token as input parameterBusiness403any00Transaction ExpiredTransaction expiredSystem403any01Feature Not Allowed \[Reason\]This merchant is not allowed to call Direct Debit APIsBusiness403any02Exceeds Transaction Amount LimitExceeds Transaction Amount LimitBusiness403any03Suspected FraudSuspected FraudBusiness403any04Activity Count Limit ExceededToo many request, Exceeds Transaction Frequency LimitBusiness403any05Do Not HonorAccount or User status is abnormalSystem403any06Feature Not Allowed At This Time. \[reason\]Cut off In ProgressBusiness403any07Card BlockedThe payment card is blockedBusiness403any08Card ExpiredThe payment card is expiredBusiness403any09Dormant AccountThe account is dormantBusiness403any10Need To Set Token LimitNeed to set token limitSystem403any11OTP BlockedOTP has been blockedSystem403any12OTP Lifetime ExpiredOTP has been expiredSystem403any13OTP Sent To Cardholerinitiates request OTP to the issuerBusiness403any14Insufficient FundsInsufficient FundsBusiness403any15Transaction Not Permitted.\[reason\]Transaction Not PermittedBusiness403any16Suspend TransactionSuspend TransactionBusiness403any17Token Limit ExceededPurchase amount exceeds the token limit set priorBusiness403any18Inactive Card/Account/CustomerIndicates inactive accountBusiness403any19Merchant BlacklistedMerchant is suspended from calling any APIsBusiness403any20Merchant Limit ExceedMerchant aggregated purchase amount on that day exceeds the agreed limitBusiness403any21Set Limit Not AllowedSet limit not allowed on particular tokenBusiness403any22Token Limit InvalidThe token limit desired by the merchant is not within the agreed range between the merchant and the IssuerBusiness403any23Account Limit ExceedAccount aggregated purchase amount on that day exceeds the agreed limitBusiness404any00Invalid Transaction StatusInvalid transaction statusBusiness404any01Transaction Not FoundTransaction not foundSystem404any02Invalid RoutingInvalid RoutingSystem404any03Bank Not Supported By SwitchBank not supported by switchBusiness404any04Transaction CancelledTransaction is cancelled by customerBusiness404any05Merchant Is Not Registered For Card Registration ServicesMerchant is not registered for Card Registration servicesSystem404any06Need To Request OTPNeed to request OTPSystem404any07Journey Not FoundThe journeyId cannot be found in the systemBusiness404any08Invalid MerchantMerchant does not exist or status abnormalBusiness404any09No IssuerNo issuerSystem404any10Invalid API TransitionInvalid API transition within a journeyBusiness404any11Invalid Card/Account/Customer \[info\]/Virtual AccountCard information may be invalid, or the card account may be blacklisted, or Virtual Account number maybe invalid.Business404any12Invalid Bill/Virtual Account \[Reason\]The bill is blocked/ suspended/not found.Virtual account is suspend/not found.Business404any13Invalid AmountThe amount doesn't match with what supposed toBusiness404any14Paid BillThe bill has been paidSystem404any15Invalid OTPOTP is incorrectBusiness404any16Partner Not FoundPartner number can't be foundBusiness404any17Invalid TerminalTerminal does not exist in the systemBusiness404any18Inconsistent RequestInconsistent request parameter found for the same partner reference number/transaction idIt can be considered as failed in transfer debit, but it should be considered as success in transfer credit.Considered as success:\- Transfer credit = (i) Intrabank transfer; (ii) Interbank transfer; (iii) RTGS transfer; (iv) SKNBI transfer;\- Virtual account = (i) Payment VA; (ii) Payment to VA;\- Transfer debit = (i) Refund payment; (ii) Void;Considered as failed:\- Transfer credit = (i) Transfer to OTC;\- Transfer debit = (i) Direct debit payment; (ii) QR CPM payment; (iii) Auth payment; (iv) Capture;Business404any19Invalid Bill/Virtual AccountThe bill is expired.Virtual account is expired.System405any00Requested Function Is Not SupportedRequested function is not supportedBusiness405any01Requested Opearation Is Not AllowedRequested operation to cancel/refund transaction Is not allowed at this time.System409any00ConflictCannot use same X-EXTERNAL-ID in same daySystem409any01Duplicate partnerReferenceNoTransaction has previously been processed indicates the same partnerReferenceNo already successSystem429any00Too Many RequestsMaximum transaction limit exceededSystem500any00General ErrorGeneral ErrorSystem500Any01Internal Server ErrorUnknown Internal Server Failure, Please retry the process againSystem500Any02External Server ErrorBackend system failure, etcSystem504any00Timeouttimeout from the issuer

#### APLIKASI PENGUJIAN

×

Akses Terbatas, Mohon Sign Up untuk Dapat Mengakses Halaman Ini

© 2025 - SNAP Developer Site