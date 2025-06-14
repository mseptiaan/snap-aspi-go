Auth Payment - SNAP Developer Site

[![](/img/icon/snap-header-white-100.png)](/)

- [**Beranda**](/)
- [**Info**](/info)
- [**API**](/api-services)
- [**Direktori Publikasi**](/direktori-publikasi)
  - [**Sign In**](/sign-in)

[![](/img/icon/bi-logo-full-100.png)](/)

![](/img/card/transfer-debit-100.png)

### [Transfer Debit](/api-services/transfer-debit)

##### Auth Payment

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

#### API Auth Payment

![](/img/docs/5_7_sequence_diagram_api_auth_payment.png)

Sequence Diagram API Auth Payment

##### API Auth Payment

**Informasi Umum**

Service Code63NameAPI Auth PaymentVersion1.0HTTP MethodGETPath.../{version}/auth/payment

##### API Payment Query

**Informasi Umum**

Service Code64NameAPI Payment QueryVersion1.0HTTP MethodGETPath.../{version}/auth/query

##### API Capture

**Informasi Umum**

API Capture bisa dijalankan terlebih dahulu lalu sebelum API Void

Service Code65NameAPI CaptureVersion1.0HTTP MethodGETPath.../{version}/auth/capture

##### API Capture Query

**Informasi Umum**

Service Code66NameAPI Capture QueryVersion1.0HTTP MethodGETPath.../{version}/auth/capture-query

##### API Void

**Informasi Umum**

API Void bisa dijalankan terlebih dahulu lalu sebelum API Capture

Service Code67NameAPI VoidVersion1.0HTTP MethodGETPath.../{version}/auth/void

##### API Void Query

**Informasi Umum**

Service Code68NameAPI Void QueryVersion1.0HTTP MethodGETPath.../{version}/auth/void-query

##### API Refund

**Informasi Umum**

Service Code69NameAPI RefundVersion1.0HTTP MethodGETPath.../{version}/auth/refund

#### GUIDES

#### Spesifikasi Parameter Header dan Body API Auth Payment

##### API Auth Payment

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringM64Transaction identifier on service consumer systemmerchantIdStringM64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDamountObjectOvalueStringM16,2Net amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)feeTypeStringO25to whom the fee will be chargedOUR = Fee is charged to the sender (default), BEN = Fee is charged to the recipient,SHA or 1000 = Fee is shared between sender and recipient, with sender is charged Rp 1.000,00and the recipient will be charged the restmccStringO32Merchant's category.productCodeStringO64Product code to identify which product used for this particular payment.titleStringM256Order title.itemsListON/APurchased goods/service information.additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionpartnerReferenceNoStringO64Transaction identifier on service consumer systemamountObjectOvalueStringM16,2Net amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)paidTimeStringM25Transaction paid time. Format paid time : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Payment Query

**Request Body**

ParameterData TypeMandatoryLengthDescriptionoriginalPartnerReferenceNoStringO64Transaction identifier on service consumer systemoriginalReferenceNoStringO64Transaction identifier on service provider systemmerchantIdStringO64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IdexternalStoreIdStringO64External Store IdadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionoriginalPartnerReferenceNoStringO64Transaction identifier on service consumer systemoriginalReferenceNoStringO64Transaction identifier on service provider systemamountObjectOvalueStringM16,2Net amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)paidTimeStringM25Transaction paid time. Format paid time : (ISO 8601) YYYY-MM-DDThh:mm:sslatestTransactionStatusStringM200 - Success01 - Initiated02 - Paying03 - Pending04 - Refunded05 - Canceled06 - Failed07 - Not foundtransactionStatusDescStringO50Description status transactionadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Capture

**Request Body**

ParameterData TypeMandatoryLengthDescriptionoriginalReferenceNoStringM64Transaction identifier on service provider systemoriginalPartnerReferenceNoStringM64Transaction identifier on service consumer systemmerchantIdStringM64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDpartnerCaptureNoStringM64Capture identifier generated by the partnercaptureAmountObjectOvalueStringM16,2Capture amount.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)titleStringM256Capture title.lastCaptureStringO8Flag to determine whether this is the last capture and void the rest of the money if there's any money left.additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionoriginalReferenceNoStringO64Transaction identifier on service provider systemoriginalPartnerReferenceNoStringO64Transaction identifier on service consumer systempartnerCaptureNoStringO64Capture identifier generated by the partnercaptureNoStringC64PJSP's capture identifier. Used to trace the capture when there's any issue occurred. Must be filled upon successful transactioncaptureAmountObjectOvalueStringM16,2Capture amount .If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)captureTimeStringC25Must be filled upon successful transaction . Format Capture time : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Capture Query

**Request Body**

ParameterData TypeMandatoryLengthDescriptionoriginalReferenceNoStringM64Transaction identifier on service provider systemoriginalPartnerReferenceNoStringO64Transaction identifier on service consumer systemmerchantIdStringM64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDcaptureNoStringO64PJSP's capture identifier. Used to trace the payment when there's any issue occurred.partnerCaptureNoStringM64Capture identifier generated by the partneradditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionoriginalReferenceNoStringO64Transaction identifier on service provider systemoriginalPartnerReferenceNoStringO64Transaction identifier on service consumer systemcaptureNoStringO64PJSP's capture identifier. Used to trace the capture when there's any issue occurred.captureAmountObjectOvalueStringM16,2Capture amount.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)captureTimeStringC25Must be filled upon successful transaction . Format Capture time : (ISO 8601) YYYY-MM-DDThh:mm:sslatestCaptureStatusStringC32Capture status. Must be filled upon successful transaction INIT,SUCCESS,FAILED.partnerCaptureNoStringM64Capture identifier generated by the partneradditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Void

**Request Body**

ParameterData TypeMandatoryLengthDescriptionoriginalReferenceNoStringM64Transaction identifier on service provider systemoriginalPartnerReferenceNoStringM64Transaction identifier on service consumer systemmerchantIdStringM64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDvoidAmountObjectOvalueStringM16,2Void amount.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)partnerVoidNoStringM64Void identifier generated by the partnervoidRemainingAmountStringO8Flag to determine whether this is the last void and void the rest of the money.reasonStringO256Capture title.additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionoriginalReferenceNoStringO64Transaction identifier on service provider systemoriginalPartnerReferenceNoStringO64Transaction identifier on service consumer systemvoidNoStringC64PJSP's void identifier. Used to trace the capture when there's any issue occurred. Must be filled upon successful transactionpartnerVoidNoStringM64Void identifier generated by the partnervoidAmountObjectOvalueStringM16,2Void amount.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)voidTimeStringC25Must be filled upon successful transaction . Format void time : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Void Query

**Request Body**

ParameterData TypeMandatoryLengthDescriptionoriginalReferenceNoStringM64Transaction identifier on service provider systemoriginalPartnerReferenceNoStringO64Transaction identifier on service consumer systemmerchantIdStringM64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDvoidNoStringO64PJSP's void identifier. Used to trace the payment when there's any issue occurred.partnerVoidNoStringM64Void identifier generated by the partneradditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionoriginalReferenceNoStringO64Transaction identifier on service provider systemoriginalPartnerReferenceNoStringO64Transaction identifier on service consumer systemvoidNoStringO64PJSP's void identifier. Used to trace the capture when there's any issue occurred.voidAmountObjectOvalueStringM16,2Void amount.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)voidTimeStringC25Must be filled upon successful transaction . Format void time : (ISO 8601) YYYY-MM-DDThh:mm:sslatestVoidStatusStringC32Void status. Must be filled upon successful transaction INIT, SUCCESS,FAILED.partnerVoidNoStringO64Void identifier generated by the partneradditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Refund

**Request Body**

ParameterData TypeMandatoryLengthDescriptionoriginalPartnerReferenceNoStringM64Transaction identifier on service consumer systemoriginalReferenceNoStringO64Transaction identifier on service provider systempartnerRefundNoStringM64ReferenceNumber from PJP AIS for the refund.merchantIdStringO64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDoriginalCaptureNoStringC64PJSP's capture identifier. Used to trace the payment when there's any issue occurred. Must be filled upon unsuccessful transactionrefundAmountObjectOvalueStringM16,2Refund amount.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)externalStoreIdStringO64External Store IDreasonStringO256Refund reasonadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionoriginalCaptureNoStringC64PJSP's capture identifier. Used to trace the payment when there's any issue occurred. Must be filled upon unsuccessful transactionoriginalReferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionoriginalPartnerReferenceNoStringO64Transaction identifier on service consumer systempartnerRefundNoStringO64PJSP's refund identifier. Used to trace the capture when there's any issue occurred.refundNoStringM64referenceNumberrefundAmountObjectOvalueStringM16,2Refund amount.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)refundTimeStringM25Refund time.additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

#### CODE SNIPPETS

#### Code Snippets API Auth Payment

##### API Auth Payment

**Sample Request**

```http
POST .../v1.0/auth/payment HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T08:58:11+07:00
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
   "merchantId":"merch00001",
   "subMerchantId":"310928924949487",
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "feeType":"BEN",
   "mcc":"5743",
   "productCode":"293840918203",
   "title":"Ikan bakar bumbu kuning",
   "items":[
      {
         "goodsId":"908132",
         "price":{
            "value":"10000.00",
            "currency":"IDR"
         },
         "category":"food",
         "unit":"ekor",
         "quantity":"2"
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
X-TIMESTAMP: 2020-12-23T08:58:19+07:00

{
   "responseCode":"2006300",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "amount":{
      "value":"10000.00",
      "currency":"IDR"
   },
   "paidTime":"2009-07-03T12:08:56+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Payment Query

**Sample Request**

```http
POST .../v1.0/auth/query HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T09:10:11+07:00
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
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-23T09:10:18+07:00

{
   "responseCode":"2006400",
   "responseMessage":"Request has been processed successfully",
   "originalpartnerReferenceNo":"2020102900000000000001",
   "originalReferenceNo":"2020102977770000000009",
   "amount":{
      "value":"10000.00",
      "currency":"IDR"
   },
   "paidTime":"2009-07-03T12:08:56+07:00",
   "latestTransactionStatus":"00",
   "transactionStatusDesc":"success",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Capture

**Sample Request**

```http
POST .../v1.0/auth/capture HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T09:12:11+07:00
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
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "partnerCaptureNo":"0000710001012344",
   "captureAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "title":"Confirmed",
   "lastCapture":"TRUE",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-23T09:12:18+07:00

{
   "responseCode":"2006500",
   "responseMessage":"Request has been processed successfully",
   "originalReferenceNo":"2020102977770000000009",
   "originalPartnerReferenceNo":"2020102900000000000001",
   "partnerCaptureNo":"0000710001012344",
   "captureNo":"202010173821904898012234423",
   "captureAmount":{
      "value":"10000.00",
      "currency":"IDR"
   },
   "captureTime":"2009-07-03T12:08:56+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Capture Query

**Sample Request**

```http
POST .../v1.0/auth/capture-query HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T09:15:11+07:00
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
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "captureNo":"202010173821904898012234423",
   "partnerCaptureNo":"0000710001012344",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-23T09:15:17+07:00

{
   "responseCode":"2006600",
   "responseMessage":"Request has been processed successfully",
   "originalReferenceNo":"2020102977770000000009",
   "originalPartnerReferenceNo":"2020102900000000000001",
   "captureNo":"202010173821904898012234423",
   "captureAmount":{
      "value":"10000.00",
      "currency":"IDR"
   },
   "captureTime":"2009-07-03T12:08:56+07:00",
   "latestCaptureStatus":"SUCCESS",
   "partnerCaptureNo":"0000710001012344",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Void

**Sample Request**

```http
POST .../v1.0/auth/void HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T09:19:47+07:00
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
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "voidAmount":{
      "value":"10000.00",
      "currency":"IDR"
   },
   "partnerVoidNo":"310928924945645",
   "voidRemainingAmount":"TRUE",
   "reason":"Shorter period/distance.",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-23T09:19:56+07:00

{
   "responseCode":"2006700",
   "responseMessage":"Request has been processed successfully",
   "originalReferenceNo":"2020102977770000000009",
   "originalPartnerReferenceNo":"2020102900000000000001",
   "voidNo":"202010173821904898012234423",
   "partnerVoidNo":"310928924945645",
   "voidAmount":{
      "value":"10000.00",
      "currency":"IDR"
   },
   "voidTime":"2009-07-03T12:08:56+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Void Query

**Sample Request**

```http
POST .../v1.0/auth/void-query HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T09:22:47+07:00
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
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "voidNo":"202010173821904898012234423",
   "partnerVoidNo":"2020101738445452",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-23T09:22:56+07:00

{
   "responseCode":"2006800",
   "responseMessage":"Request has been processed successfully",
   "originalReferenceNo":"2020102977770000000009",
   "originalPartnerReferenceNo":"2020102900000000000001",
   "voidNo":"202010173821904898012234423",
   "voidAmount":{
      "value":"10000.00",
      "currency":"IDR"
   },
   "voidTime":"2009-07-03T12:08:56+07:00",
   "latestVoidStatus":"SUCCESS",
   "partnerVoidNo":"2020101738445452",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Refund

**Sample Request**

```http
POST .../v1.0/auth/refund HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-23T09:22:47+07:00
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
   "partnerRefundNo":"239850918204981205970",
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "originalCaptureNo":"202010173821904898012234423",
   "refundAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "externalStoreId":"124928924949487",
   "reason":" Customer complain",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-23T09:22:56+07:00

{
   "responseCode":"2006900",
   "responseMessage":"Request has been processed successfully",
   "originalCaptureNo":"202010173821904898012234423",
   "originalReferenceNo":"2020102977770000000009",
   "originalPartnerReferenceNo":"2020102900000000000001",
   "partnerRefundNo":"202010173821904898012234423",
   "refundNo":"REF993883",
   "refundAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "refundTime":"2020-12-23T09:19:47+07:00",
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