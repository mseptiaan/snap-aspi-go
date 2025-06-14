Direct Debit - SNAP Developer Site

[![](/img/icon/snap-header-white-100.png)](/)

- [**Beranda**](/)
- [**Info**](/info)
- [**API**](/api-services)
- [**Direktori Publikasi**](/direktori-publikasi)
  - [**Sign In**](/sign-in)

[![](/img/icon/bi-logo-full-100.png)](/)

![](/img/card/transfer-debit-100.png)

### [Transfer Debit](/api-services/transfer-debit)

##### Direct Debit

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

#### API Direct Debit

![](/img/docs/5_2_sequence_diagram_redirect_(web_check_out_with_otp).png)

Sequence Diagram Redirect (web check out with otp)

![](/img/docs/5_3_sequence_diagram_redirect_(web_check_out_without_otp).png)

Sequence Diagram Redirect (web check out without otp)

![](/img/docs/5_4_sequence_diagram_api_direct_debit-payment_host_to_host.jpg)

Sequence Diagram API Direct Debit-Payment Host to Host

![](/img/docs/5_5_sequence_diagram_api_direct_debit-app_link.jpg)

Sequence Diagram API Direct Debit-App Link

##### API Direct Debit Payment

**Informasi Umum**

Service Code54NameAPI Direct Debit PaymentVersion1.0HTTP MethodPOSTPath.../{version}/debit/payment-host-to-host

##### API Direct Debit Payment Status

**Informasi Umum**

Service Code55NameAPI Direct Debit Payment StatusVersion1.0HTTP MethodPOSTPath.../{version}/debit/status

##### API Direct Debit Payment Notification

**Informasi Umum**

Service Code56NameAPI Direct Debit Payment NotificationVersion1.0HTTP MethodPOSTPath.../{version}/debit/notify

##### API Direct Debit Payment Cancel

**Informasi Umum**

Service Code57NameAPI Direct Debit Payment CancelVersion1.0HTTP MethodPOSTPath.../{version}/debit/cancel

##### API Direct Debit Payment Refund

**Informasi Umum**

Service Code58NameAPI Direct Debit Payment RefundVersion1.0HTTP MethodPOSTPath.../{version}/debit/refund

#### GUIDES

#### Spesifikasi Parameter Header dan Body API Direct Debit

##### API Direct Debit Payment

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringM64Transaction identifier on service consumer systembankCardTokenStringO128Card token for paymentchargeTokenStringO40string code for verification OTPotpStringO8OTP Code / PasscodeotpTrxCodeStringO2Purpose of otpmerchantIdStringO64Merchant identifier that isunique per each merchantterminalIdStringO64Terminal IDjourneyIdStringO64Merchant IDsubMerchantIdStringO32Sub merchant IDamountObjectOvalueStringM16,2Net amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)urlParamsArray of ObjectsOurlStringM512The URLtypeStringM32URL TypePAY\_RETURN/PAY\_NOTIFYisDeeplinkStringM1Whther the URL is a deeplink URL or notY/NexternalStoreIdStringO64Store ID to indicate to which storethis payment belongs to.validUpToStringO25The time when the payment will be automatically expired.ISO 8601pointOfInitiationStringO20used for getting more info regarding source of request of the userfeeTypeStringO25to whom the fee will be chargeddisabledPayMethodsStringO64Payment method(s) that cannot be used for this paymentpayOptionDetailsArray ofOPayment option that will be used for this paymentObjectpayMethodStringM64Payment Method. e.g. CREDIT\_CARDpayOptionStringM64Payment option which shows the provider of this paymente.g. CREDIT\_CARD\_VISAtransAmountObjectOvalueStringM16,2Transaction amount that will be paid using this payment method If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)feeAmountObjectOvalueStringM16,2Fee amount that will be paid using this payment method If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)cardTokenStringO64Card token used for this payment.merchantTokenStringO64Merchant token used for this payment.additionalInfoObjectOAdditional information for custom use that are not provided by SNAP for pay Option DetailsadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionpartnerReferenceNoStringO64Transaction identifier on service consumer systemapprovalCodeStringO20appRedirectUrlStringO2048Returns an URL scheme to PJP AIS payment page in native app.webRedirectUrlStringO2048Returns a universal link to PJP AIS payment page. This link is recommended when the Client is unable to implement acheck for whether PJP AIS app is installed on the user's device before redirect.additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Direct Debit Payment Status

**Request Body**

KeyValueMandatoryLengthDescriptionoriginalPartnerReferenceNoStringO64Original transaction identifier on service consumer systemoriginalReferenceNoStringO64Original transaction identifier on service provider systemoriginalExternalIdStringO32Original External-ID on header messageserviceCodeStringM2Transaction type indicator. Service code of the transaction that needs to be checked (original transaction).transactionDateStringO25Format transaction date : (ISO 8601) YYYY-MM-DDThh:mm:ssamountObjectOvalueStringM16,2Net amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)merchantIdStringO64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDexternalStoreIdStringO64External Store ID for merchantadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionoriginalReferenceNoStringC64Original transaction identifier on service provider system. Must be filled upon successful transactionoriginalPartnerReferenceNoStringO64Original transaction identifier on service consumer systemapprovalCodeStringO20originalExternalIdStringO32Original External-ID on header messageserviceCodeStringM2Transaction type indicator. Service code of the transaction that needs to be checked (original transaction).latestTransactionStatusStringM200 - Success01 -Initiated02 - Paying03 - Pending04 - Refunded05 - Canceled06 - Failed07 - Not foundtransactionStatusDescStringO50Description status transactionoriginalResponseCodeStringO7Response codeoriginalResponseMessageStringO150Response descriptionsessionIdStringO25Transaction invoice IDrequestIdStringO25Transaction request IDrefundHistoryArray of ObjectsOrefundNoStringC64Transaction Identifier on Service Provider SystempartnerRefundNoStringM64ReferenceNumber from PJP AIS for the refund.refundAmountObjectOvalueStringM16,2Net amount of the refund.currencyStringM3Currency (ISO4217)refundStatusStringM200 - Success03 - Pending06 - FailedrefundDateStringC25(ISO 8601)transaction date :dd-MM-yyyy (Mandatory)HH:mm:ss(Optional)reasonStringO256Refund reason.transAmountObjectOvalueStringM16,2Transaction amount that will be paid using this payment method If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)feeAmountObjectOvalueStringM16,2Fee amount that will be paid using this payment method If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)paidTimeStringC25Format transaction date : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Direct Debit Payment Notification

**Request Body**

KeyValueMandatoryLengthDescriptionoriginalPartnerReferenceNoStringO64Original transaction identifier on service consumer systemoriginalReferenceNoStringM64Original transaction identifier on service provider systemoriginalExternalIdStringO32Original CustomerReference NumbermerchantIdStringO64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDamountObjectOvalueStringM16,2Net amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)latestTransactionStatusStringM200 - Success01 - Initiated02 - Paying03 - Pending04 - Refunded05 - Canceled06 - Failed07 - Not foundtransactionStatusDescStringO50Description status transactioncreatedTimestringO25Transaction created time.finishedTimestringO25Transaction finished time.externalStoreIdStringO64External Store IdadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response codeapprovalCodeStringO20Approval CoderesponseMessageStringM150Response description

##### API Direct Debit Payment Cancel

**Request Body**

KeyValueMandatoryLengthDescriptionoriginalPartnerReferenceNoStringM64Transaction identifier on service consumer systemoriginalReferenceNoStringO64Original transaction identifier on service provider systemapprovalCodeStringO20originalExternalIdStringO32Original External-ID on header messagemerchantIdStringO64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDreasonStringO256Cancellation reason.externalStoreIdStringO64External Store IDamountObjectOvalueStringM16,2Transaction AmountcurrencyStringM3Currency (ISO4217)additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionoriginalPartnerReferenceNoStringO64Original transaction identifier on service consumer systemoriginalReferenceNoStringC64Original transaction identifier on service provider systemoriginalExternalIdStringO32Original External-ID on header messagecancelTimeStringC25Cancel timeISO-8601. Must be filled if cancelled transaction successtransactionDateStringO25Format transaction date : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Direct Debit Payment Refund

**Request Body**

KeyValueMandatoryLengthDescriptionmerchantIdStringO64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDoriginalPartnerReferenceNoStringM64Original transaction identifier on service consumer systemoriginalReferenceNoStringO64Original transaction identifier on service provider systemoriginalExternalIdStringO32Original CustomerReference NumberpartnerRefundNoStringM64ReferenceNumber from PJP AIS for the refund.refundAmountObjectOvalueStringM16,2Net amount of the refund.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)externalStoreIdStringO64External Store IDreasonStringO256Refund reason.additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionoriginalPartnerReferenceNoStringO64Original transaction identifier on service consumer systemoriginalReferenceNoStringC64Original transaction identifier on service provider systemoriginalExternalIdStringO32Original CustomerReference NumberpartnerTrxIdStringO32Partner Transaction IDrefundNoStringM64referenceNumberpartnerRefundNoStringM64ReferenceNumber from PJP AIS for the refund.refundAmountObjectOvalueStringM16,2Net amount of the refund.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)refundTimeStringM25Refund time.ISO 8601additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

#### CODE SNIPPETS

#### Code Snippets API Direct Debit

##### API Direct Debit Payment

**Sample Request**

```http
POST .../v1.0/debit/payment-host-to-host HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T07:44:11+07:00
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
   "bankCardToken":"00007100010926",
   "chargeToken":"00007100010926",
   "otp":"12345678",
   "merchantId":"23489182303312",
   "terminalId":"489182303312",
   "journeyId":"861023713017210",
   "subMerchantId":"310928924949487",
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "urlParams":[
      {
         "url":"https://test1.bi.go.id/v1/test",
         "type":"PAY_RETURN",
         "isDeeplink":"Y"
      },
      {
         "url":"https://test2.bi.go.id/v1/test",
         "type":"PAY_RETURN",
         "isDeeplink":"Y"
      }
   ],
   "externalStoreId":"239840198240795109",
   "validUpTo":"2020-12-23T07:44:11+07:00",
   "pointOfInitiation":"Mobile App",
   "feeType":"SHA|1000",
   "disabledPayMethods":"CREDIT_CARD",
   "payOptionDetails":[
      {
         "payMethod":"CREDIT_CARD",
         "payOption":"CREDIT_CARD_VISA",
         "transAmount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "feeAmount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "cardToken":"d89ca90ua90sd80as9809",
         "merchantToken":"a90ua90sd80d89cas9809",
         "additionalInfo":{
            "deviceId":"12345679237",
            "channel":"mobilephone"
         }
      },
      {
         "payMethod":"CREDIT_CARD",
         "payOption":"CREDIT_CARD_MASTER",
         "transAmount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "feeAmount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "cardToken":"d89ca90ua90sd80as9809",
         "merchantToken":"a90ua90sd80d89cas9809",
         "additionalInfo":{
            "deviceId":"12345679237",
            "channel":"mobilephone"
         }
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
X-TIMESTAMP: 2020-12-23T07:44:16+07:00

{
   "responseCode":"20054000",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "approvalCode":"201039000200",
   "appRedirectUrl":"https://pjsp.com/app?bizNo=REF993883&...",
   "webRedirectUrl":"https://pjsp.com/universal?bizNo=REF993883&...",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Direct Debit Payment Status

**Sample Request**

```http
POST .../v1.0/debit/status HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T07:44:11+07:00
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
   "originalPartnerReferenceNo":"2020102900000000000001",
   "originalReferenceNo":"2020102977770000000009",
   "originalExternalId":"30443786930722726463280097920912",
   "serviceCode":"54",
   "transactionDate":"2020-12-21T14:56:11+07:00",
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "merchantId":"23489182303312",
   "subMerchantId":"23489182303312",
   "externalStoreId":"183908924912387",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-23T07:44:16+07:00

{
   "responseCode":"2005500",
   "responseMessage":"Request has been processed successfully",
   "originalPartnerReferenceNo":"2020102900000000000001",
   "originalReferenceNo":"2020102977770000000009",
   "approvalCode":"201039000200",
   "originalExternalId":"30443786930722726463280097920912",
   "serviceCode":"54",
   "latestTransactionStatus":"00",
   "transactionStatusDesc":"success",
   "originalResponseCode":"2005400",
   "originalResponseMessage":"Request has been processed successfully",
   "sessionId":"883737GHY8839",
   "requestId":"3763773",
   "refundHistory":[
      {
         "refundNo":"96194816941239812",
         "partnerRefundNo":"239850918204981205970",
         "refundAmount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "refundStatus":"00",
         "refundDate":"2020-12-23T07:44:16+07:00",
         "reason":"Customer Complain"
      },
      {
         "refundNo":"96194123981251341",
         "partnerRefundNo":"2398509123131981205970",
         "refundAmount":{
            "value":"112345678.00",
            "currency":"IDR"
         },
         "refundStatus":"00",
         "refundDate":"2020-12-23T07:54:16+07:00",
         "reason":"Customer Complain"
      }
   ],
   "transAmount":{
      "value":"112345678.00",
      "currency":"IDR"
   },
   "feeAmount":{
      "value":"112345678.00",
      "currency":"IDR"
   },
   "paidTime":"2020-12-21T14:56:11+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Direct Debit Payment Notification

**Sample Request**

```http
POST .../v1.0/debit/notify HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T07:44:11+07:00
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
   "originalPartnerReferenceNo":"2020102900000000000001",
   "originalReferenceNo":"2020102977770000000009",
   "originalExternalId":"30443786930722726463280097920912",
   "merchantId":"23489182303312",
   "subMerchantId":"23489182303312",
   "amount":{
      "value":"10000.00",
      "currency":"IDR"
   },
   "latestTransactionStatus":"00",
   "transactionStatusDesc":"success",
   "createdTime":"2020-12-21T17:07:18+07:00",
   "finishedTime":"2020-12-21T17:07:20+07:00",
   "externalStoreId":"91863213913112",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-23T07:44:16+07:00

{
   "responseCode":"2005600",
   "approvalCode":"201039000200",
   "responseMessage":"Request has been processed successfully"
}

```

##### API Direct Debit Payment Cancel

**Sample Request**

```http
POST .../v1.0/debit/cancel HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T07:44:11+07:00
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
   "originalPartnerReferenceNo":"2020102900000000000001",
   "originalReferenceNo":"2020102977770000000009",
   "approvalCode":"201039000200",
   "originalExternalId":"30443786930722726463280097920912",
   "merchantId":"23489182303312",
   "subMerchantId":"23489182303312",
   "reason":"Network timeout",
   "externalStoreId":"124928924949487",
   "amount":{
      "value":"10000.00",
      "currency":"IDR"
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
X-TIMESTAMP: 2020-12-23T07:44:16+07:00

{
   "responseCode":"2005700",
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

##### API Direct Debit Payment Refund

**Sample Request**

```http
POST .../v1.0/debit/refund HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T07:44:11+07:00
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
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "originalPartnerReferenceNo":"2020102900000000000001",
   "originalReferenceNo":"2020102977770000000009",
   "originalExternalId":"10052019",
   "partnerRefundNo":"239850918204981205970",
   "refundAmount":{
      "value":"10000.00",
      "currency":"IDR"
   },
   "externalStoreId":"124928924949487",
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
X-TIMESTAMP: 2020-12-23T07:44:16+07:00

{
   "responseCode":"2005800",
   "responseMessage":"Request has been processed successfully",
   "originalPartnerReferenceNo":"2020102900000000000001",
   "originalReferenceNo":"2020102977770000000009",
   "originalExternalId":"10052019",
   "partnerTrxId":"LA001",
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