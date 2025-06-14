Direct Debit BI-FAST - SNAP Developer Site

[![](/img/icon/snap-header-white-100.png)](/)

- [**Beranda**](/)
- [**Info**](/info)
- [**API**](/api-services)
- [**Direktori Publikasi**](/direktori-publikasi)
  - [**Sign In**](/sign-in)

[![](/img/icon/bi-logo-full-100.png)](/)

![](/img/card/transfer-debit-100.png)

### [Transfer Debit](/api-services/transfer-debit)

##### Direct Debit BI-FAST

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

#### API Direct Debit BI-FAST

##### API Registrasi e-Mandate

![](/img/docs/5_8_sequence_diagram_api_registrasi_e-mandate.png)

Sequence Diagram API Registrasi e-Mandate

**Informasi Umum**

Service Code70NameAPI Registrasi e-MandateVersion1.0HTTP MethodPOSTPath.../{version}/debit/fast-emandate

##### API Payment

![](/img/docs/5_9_sequence_diagram_api_trigger_direct_debit_transfer.png)

Sequence Diagram API Trigger Direct Debit Transfer

**Informasi Umum**

Service Code71NameAPI Trigger Direct Debit TransferVersion1.0HTTP MethodPOSTPath.../{version}/debit/fast-payment

##### API Notify

![](/img/docs/5_10_sequence_diagram_api_notify_-_direct_debit.png)

Sequence Diagram API Notify - Direct Debit

**Informasi Umum**

Service Code72NameAPI NotifyVersion1.0HTTP MethodPOSTPath.../{version}/debit/fast-notify

#### GUIDES

#### Spesifikasi Parameter Header dan Body API Direct Debit BI-FAST

##### API Registrasi e-Mandate

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringO64Transaction identifier on service consumer systembankCodeStringM3Biller bank code based on Bank Indonesia CodesourceAccountNoStringM19Account number to be billedsourceAccountNameStringM100Source Account NamemaxAmountObjectOvalueStringM16,2Maximum amount to be paid.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)billerIdStringM30Unique ID to identify billerbillerNameStringM50Biller namecustomerIdStringM45Customer ID registered on billerexpiredDatetimeStringM25e-Mandate expiration date and time, Format expiration date time : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionpartnerReferenceNoStringO64Transaction identifier on service consumer systemeMandateReffIdStringM30Unique key to identify an e-Mandate registrationadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Payment

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringM64Transaction identifier on service consumer systemcurrencyStringO3Currency (ISO4217)customerReferenceStringM30Reference Number / No Referral / Transaction IDfeeTypeStringO25to whom the fee will be chargedremarkStringO50Remark/transaction descriptionbeneficiaryAccountNoStringM19Biller accountbeneficiaryAccountNameStringM100Biller account Account NametransactionDateStringM25transaction date : ISO 8601bankCodeStringM3Bank code based on Bank Indonesia CodesourceAccountNoStringM34Debited customer account numbersourceAccountNameStringM100Debited customer account NameamountObjectOvalueStringM16,2Debit amount.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)eMandateReffIdStringM30Unique key to identify an e-Mandate registrationadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionpartnerReferenceNoStringO64Transaction identifier on service consumer systemadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Notify

**Request Body**

ParameterData TypeMandatoryLengthDescriptionoriginalReferenceNoStringM64Original transaction identifier on service provider systemoriginalPartnerReferenceNoStringO64Original transaction identifier on service consumer systemoriginalExternalIdStringO19Original CustomerReference NumbertransactionStatusStringM200 - Success01 - Initiated02 - Paying03 - Pending04 - Refunded05 - Canceled06 - Failed07 - Not foundtransactionStatusDescStringO50Description status transactioneMandateReffIdStringM30Unique key to identify an e-Mandate registrationsourceAccountNoStringM34Debited customer account numbersourceAccountNameStringM100Debited customer account NameamountObjectOvalueStringM16,2Net amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217traceNoStringO16Number for tracking to destination bankadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

#### CODE SNIPPETS

#### Code Snippets API Direct Debit BI-FAST

##### API Registrasi e-Mandate

**Sample Request**

```http
POST .../v1.0/debit/fast-emandate HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-22T08:12:16+07:00
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
   "partnerReferenceNo":"2020102900000000000001",
   "bankCode":"014",
   "sourceAccountNo":"888801000157508",
   "sourceAccountName":"Saving",
   "maxAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "billerId":"315134",
   "billerName":"biller X",
   "customerId":"96891414",
   "expiredDatetime":"2022-12-22T08:01:16+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-22T08:12:22+07:00

{
   "responseCode":"2007000",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "eMandateReffId":"9a8fau6d81had833bas7184",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Payment

**Sample Request**

```http
POST .../v1.0/debit/fast-payment HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-22T08:26:16+07:00
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
   "partnerReferenceNo":"2020102900000000000001",
   "currency":"IDR",
   "customerReference":"10052019",
   "feeType":"BEN",
   "remark":"remark test",
   "beneficiaryAccountNo":"888801000157508",
   "beneficiaryAccountName":"Yories Yolanda",
   "transactionDate":"2019-07-03T12:08:56+07:00",
   "bankCode":"014",
   "sourceAccountNo":"888801000003301",
   "sourceAccountName":"Saving",
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "eMandateReffId":"9a8fau6d81had833bas7184",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-22T08:26:30+07:00

{
   "responseCode":"2007100",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Notify

**Sample Request**

```http
POST .../v1.0/debit/fast-notify HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-22T08:26:16+07:00
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
   "originalReferenceNo":"2020102977770000000009",
   "originalPartnerReferenceNo":"2020102900000000000001",
   "originalExternalId":"10052019",
   "transactionStatus":"00",
   "transactionStatusDesc":"success",
   "eMandateReffId":"9a8fau6d81had833bas7184",
   "sourceAccountNo":"888801000003301",
   "sourceAccountName":"Saving",
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "traceNo":"13415141",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-22T08:26:30+07:00

{
   "responseCode":"2007200",
   "responseMessage":"Request has been processed successfully",
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