Bulk Cashin - SNAP Developer Site

[![](/img/icon/snap-header-white-100.png)](/)

- [**Beranda**](/)
- [**Info**](/info)
- [**API**](/api-services)
- [**Direktori Publikasi**](/direktori-publikasi)
  - [**Sign In**](/sign-in)

[![](/img/icon/bi-logo-full-100.png)](/)

![](/img/card/transfer-credit-100.png)

### [Transfer Kredit](/api-services/transfer-kredit)

##### Bulk Cashin

Overview


Guides


Code Snippets


Response Code


Aplikasi Pengujian


##### Transfer Credit PJP AIS Bank

- [Account Inquiry](/api-services/transfer-kredit/account-inquiry)
- [Trigger Transfer](/api-services/transfer-kredit/trigger-transfer)
- [Virtual Account](/api-services/transfer-kredit/virtual-account)
- [Transfer Status Inquiry Bank](/api-services/transfer-kredit/transaction-status-inquiry-bank)

##### Transfer Credit PJP AIS Non Bank

- [Customer Top Up](/api-services/transfer-kredit/customer-top-up)
- [Bulk Cashin](/api-services/transfer-kredit/bulk-cashin)
- [Transfer To Bank](/api-services/transfer-kredit/transfer-to-bank)
- [Transfer To OTC](/api-services/transfer-kredit/transfer-to-otc)
- [MPM](/api-services/transfer-kredit/mpm)
- [Transfer Status Inquiry](/api-services/transfer-kredit/transaction-status-inquiry)

#### OVERVIEW

#### API Bulk Cashin

##### API Submit Bulk Cash In

![](/img/docs/4_18_sequence_diagram_api_submit_bulk_cash_in.png)

Sequence Diagram API Submit Bulk Cash In

**Informasi Umum**

Service Code40NameAPI Submit Bulk Cash InVersion1.0HTTP MethodPOSTPath.../{version}/emoney/bulk-cashin-payment

##### API Notify Bulk Cash In

**Informasi Umum**

Service Code41NameAPI Notify Bulk Cash InVersion1.0HTTP MethodPOSTPath.../{version}/emoney/bulk-cashin-notify

#### GUIDES

#### Spesifikasi Parameter Header dan Body API Bulk Cashin

##### API Submit Bulk Cash In

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerBulkIdStringO64Bulk transaction identifier on service consumer (bulk transaction sender) systemtransactionDateStringM25transaction date : ISO 8601currencystringO3Currency (ISO4217)bulkObjectArray of ObjectsaccountNumberStringM64Customer Account NumberaccountNameStringO30Customer NameamountObjectOvalueStringM16,2Net amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)partnerReferenceNoStringM64Individual transaction identifier on service consumer (bulk transaction sender) systemadditionalInfoObjectOAdditional informationfeeTypeStringO25to whom the fee will be chargedOUR = Fee is charged to the sender (default), BEN = Fee is charged to the recipient,SHA or 1000 = Fee is shared between sender and recipient, with sender is charged Rp 1.000,00and the recipient will be charged the restadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionbulkidStringM64Transaction identifier on service provider (bulk transaction processing) systempartnerBulkIdStringO64Transaction identifier on service consumer (bulk transaction sender) systemadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Notify Bulk Cash In

**Request Body**

ParameterData TypeMandatoryLengthDescriptionbulkIdStringM64Transaction identifier on bulk transaction processing systempartnerBulkIdStringM64Transaction identifier on bulk transaction sender systembulkObjectArray of ObjectscustomerNumberStringM64Customer Account NumbercustomerNameStringO30Customer NameamountObjectOvalueStringM16,2Net amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)referenceNoStringM64Individual transaction identifier on bulk transaction processing systempartnerReferenceNoStringM64Individual transaction identifier on bulk transaction sender systemresponseCodeStringM7Response CoderesponseMessageStringM150Response descriptionadditionalInfoObjectOAdditional information for custom use on each bulk objectadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionbulkIdStringM64Transaction identifier on bulk transaction processing systempartnerBulkIdStringM64Transaction identifier on bulk transaction sender systemadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

#### CODE SNIPPETS

#### Code Snippets API Bulk Cashin

##### API Submit Bulk Cash In

**Sample Request**

```http
POST .../v1.0/emoney/bulk-cashin-payment HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T17:01:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.1617169
X-LONGITUDE: 106.6643946
CHANNEL-ID: 95221

{
   "partnerBulkId":"2020102900000000000001",
   "transactionDate":"2019-07-03T12:08:56+07:00",
   "currency":"IDR",
   "bulkObject":[
      {
         "accountNumber":"6281388370001",
         "accountName":"Kades",
         "amount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "partnerReferenceNo":"2020102900000000000990",
         "additionalInfo":{
            "deviceId":"12345679237",
            "channel":"mobilephone"
         }
      }
   ],
   "feeType":"OUR",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-21T17:01:17+07:00

{
   "responseCode":"2004000",
   "responseMessage":"Request has been processed successfully",
   "bulkId":" 2020102977770000000009",
   "partnerBulkId ":"2020102900000000000001",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Notify Bulk Cash In

**Sample Request**

```http
POST .../v1.0/emoney/bulk-cashin-notify HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T17:07:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.1617169
X-LONGITUDE: 106.6643946
CHANNEL-ID: 95221

{
   "bulkId":"2020102977770000000009",
   "partnerBulkId":"2020102900000000000001",
   "bulkObject":[
      {
         "customerNumber":"6281388370001",
         "customerName":"Kades",
         "amount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "additionalInfo":{
            "deviceId":"12345679237",
            "channel":"mobilephone"
         },
         "referenceNo":"2020102977770000000009",
         "partnerReferenceNo":"2020102900000000000990",
         "responseCode":"2004000",
         "responseMessage":"Request has been processed successfully"
      }
   ],
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-21T17:07:18+07:00

{
   "responseCode":"2004100",
   "responseMessage":"Request has been processed successfully",
   "bulkId":"2020102977770000000009",
   "partnerBulkId":"2020102900000000000001",
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