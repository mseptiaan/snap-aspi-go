Trigger Transfer - SNAP Developer Site

[![](/img/icon/snap-header-white-100.png)](/)

- [**Beranda**](/)
- [**Info**](/info)
- [**API**](/api-services)
- [**Direktori Publikasi**](/direktori-publikasi)
  - [**Sign In**](/sign-in)

[![](/img/icon/bi-logo-full-100.png)](/)

![](/img/card/transfer-credit-100.png)

### [Transfer Kredit](/api-services/transfer-kredit)

##### Trigger Transfer

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

#### API Transfer

##### API Intrabank Transfer

![](/img/docs/4_4_sequence_diagram_api_trigger_intrabank_transfer.png)

Sequence Diagram API Trigger Intrabank Transfer

**Informasi Umum**

Service Code17NameAPI Intrabank TransferVersion1.0HTTP MethodPOSTPath.../{version}/transfer-intrabank

##### API Interbank Transfer

![](/img/docs/4_5_sequence_diagram_api_trigger_interbank_transfer.png)

Sequence Diagram API Trigger Interbank Transfer

**Informasi Umum**

Service Code18NameAPI Interbank TransferVersion1.0HTTP MethodPOSTPath.../{version}/transfer-interbank

##### API Request for Payment

![](/img/docs/4_6_sequence_diagram_api_request_for_payment.png)

Sequence Diagram API Request for Payment

**Informasi Umum**

Service Code19NameAPI Request for PaymentVersion1.0HTTP MethodPOSTPath.../{version}/transfer-request-for-payment

##### API Interbank Bulk Transfer

![](/img/docs/4_7_sequence_diagram_api_trigger_interbank_bulk_transfer.png)

Sequence Diagram API Trigger Interbank Bulk Transfer

**Informasi Umum**

Service Code20NameAPI Interbank Bulk TransferVersion1.0HTTP MethodPOSTPath.../{version}/transfer-interbank-bulk

##### API Interbank Bulk Transfer - Notification

![](/img/docs/4_8_sequence_diagram_api_trigger_interbank_bulk_transfer_-_notify.png)

Sequence Diagram API Trigger Interbank Bulk Transfer - Notify

**Informasi Umum**

Service Code21NameAPI Interbank Bulk Transfer - NotificationVersion1.0HTTP MethodPOSTPath.../{version}/transfer-interbank-bulk/notify

##### API Transfer RTGS

![](/img/docs/4_9_sequence_diagram_api_transfer_rtgs.png)

Sequence Diagram API Transfer RTGS

**Informasi Umum**

Service Code22NameAPI Transfer RTGSVersion1.0HTTP MethodPOSTPath.../{version}/transfer-rtgs

##### API RTGS - Notification

**Informasi Umum**

Service Code76NameAPI RTGS - NotificationVersion1.0HTTP MethodPOSTPath.../{version}/transfer-rtgs/notify

##### API Transfer SKNBI

![](/img/docs/4_10_sequence_diagram_api_transfer_sknbi.png)

Sequence Diagram API Transfer SKNBI

**Informasi Umum**

Service Code23NameAPI Transfer SKNBIVersion1.0HTTP MethodPOSTPath.../{version}/transfer-skn

##### API SKNBI - Notification

**Informasi Umum**

Service Code75NameAPI SKNBI - NotificationVersion1.0HTTP MethodPOSTPath.../{version}/transfer-skn/notify

#### GUIDES

#### Spesifikasi Parameter Header dan Body API Transfer

##### API Intrabank Transfer

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringM64Transaction identifier on service consumer systemamountObjectMvalueStringM16,2Net amount of the transaction.If it’s IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)beneficiaryAccountNoStringM34Beneficiary Account NumberbeneficiaryEmailStringO50Beneficiary EmailcurrencyStringO3Currency (ISO4217)customerReferenceStringO30Reference Number / No Referral / Transaction IDfeeTypeStringO25to whom the fee will be chargedOUR = Fee is charged to the sender (default), BEN = Fee is charged to the recipient,SHA or 1000 = Fee is shared between sender and recipient, with sender is charged Rp 1.000,00and the recipient will be charged the restremarkStringO50Remark/transaction descriptionsourceAccountNoStringM19Source Account NumbertransactionDateStringM25Format transaction date : (ISO 8601) YYYY-MM-DDThh:mm:ssoriginatorInfosArray of ObjectsCThe Originator Customer Account DetailsTo be filled if there is a request from the sender or if the consent from sender has been granted.This is subject to Article 8 paragraph 5 of Law No. 3 of 2011 concerning Fund Transfers.Also check for other provisions, such as the PPATK regulation.originatorCustomerNoStringM34Originator customer account numberoriginatorCustomerNameStringM100Originatorcustomer account nameoriginatorBankCodeStringM11Originator Bank CodeadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionpartnerReferenceNoStringO64Transaction identifier on service consumer systemamountObjectOvalueStringM16,2Net amount of the transaction.If it’s IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)beneficiaryAccountNoStringM34Beneficiary AccountcurrencyStringO3Currency (ISO4217)customerReferenceStringO30Reference Number / No ReferralsourceAccountNoStringO19Source AccounttransactionDateStringM25Format transaction date : (ISO 8601) YYYY-MM-DDThh:mm:ssoriginatorInfosArray of ObjectsCThe Originator Customer Account DetailsTo be filled if there is a request from the sender or if the consent from sender has been granted.This is subject to Article 8 paragraph 5 of Law No. 3 of 2011 concerning Fund Transfers.Also check for other provisions, such as the PPATK regulation.originatorCustomerNoStringM34Originator customer account numberoriginatorCustomerNameStringM100Originatorcustomer account nameoriginatorBankCodeStringM11Originator Bank CodeadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Interbank Transfer

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringM64Transaction identifier on service consumer systemamountObjectMvalueStringM16,2Net amount of the transaction.If it’s IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)beneficiaryAccountNameStringM100Beneficiary Account NamebeneficiaryAccountNoStringM34Beneficiary AccountbeneficiaryAddressStringO100Beneficiary AddressbeneficiaryBankCodeStringM11Beneficiary Bank CodebeneficiaryBankNameStringO50Beneficiary Bank NamebeneficiaryEmailStringO50Beneficiary EmailcurrencyStringO3Currency (ISO4217)customerReferenceStringO30Reference Number / No Referral / Transaction IDsourceAccountNoStringM19Source Account NumbertransactionDateStringM25Format transaction date : (ISO 8601) YYYY-MM-DDThh:mm:ssfeeTypeStringO25to whom the fee will be chargedOUR = Fee is charged to the sender (default), BEN = Fee is charged to the recipient,SHA or 1000 = Fee is shared between sender and recipient, with sender is charged Rp 1.000,00.and tSHAhe recipient will be charged the rest.originatorInfosArray of ObjectsCThe Originator Customer Account Details.To be filled if there is a request from the sender or if the consent from sender has been granted.This is subject to Article 8 paragraph 5 of Law No. 3 of 2011 concerning Fund Transfers.Also check for other provisions, such as the PPATK regulation.originatorCustomerNoStringM34Originator customer account numberoriginatorCustomerNameStringM100Originatorcustomer account nameoriginatorBankCodeStringM11Originator Bank CodeadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionpartnerReferenceNoStringO64Transaction identifier on service consumer systemamountObjectOvalueStringM16,2Net amount of the transaction.If it’s IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)beneficiaryAccountNoStringM19Beneficiary Account NumberbeneficiaryBankCodeStringO11Beneficiary Bank CodesourceAccountNoStringO19Source Account NumbertraceNoStringO16Number for tracking to destination bankoriginatorInfosArray of ObjectsCThe Originator Customer Account DetailsTo be filled if there is a request from the sender or if the consent from sender has been granted.This is subject to Article 8 paragraph 5 of Law No. 3 of 2011 concerning Fund Transfers.Also check for other provisions, such as the PPATK regulation.originatorCustomerNoStringM34Originator customer account numberoriginatorCustomerNameStringM100Originatorcustomer account nameoriginatorBankCodeStringM11Originator Bank CodeadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Request for Payment

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringM64Transaction identifier on service consumer systembankCodeStringM11Beneficiary bank code based on Bank Indonesia CodebeneficiaryAccountNoStringM34Beneficiary Account NumberbeneficiaryAccountNameStringM100Beneficiary Account NameremarkStringO50Remark/transaction descriptionexpiredDatetimeStringM25RFP expiration date and time , Format expired date time : (ISO 8601) YYYY-MM-DDThh:mm:sssourceAccountNoStringM19Source Account NumbersourceAccountNameStringM100Source Account NamecurrencyStringO3Currency TypeamountObjectOvalueStringM16,2Net amount of the transaction.If it’s IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)feeTypeStringO25to whom the fee will be chargedOUR = Fee is charged to the sender (default), BEN = Fee is charged to the recipient,SHA or 1000 = Fee is shared between sender and recipient, with sender is charged Rp 1.000,00and the recipient will be charged the restadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionpartnerReferenceNoStringO64Transaction identifier on service consumer systemadditionalInfoObjectO…Additional information for custom use that are not provided by SNAP

##### API Interbank Bulk Transfer

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerBulkIdStringO64Bulk transaction identifier on service consumer (bulk transaction sender) systemcurrencyStringO3Currency (ISO4217)customerReferenceStringM30Reference Number / No Referral / Transaction IDfeeTypeStringO25to whom the fee will be chargedOUR = Fee is charged to the sender (default), BEN = Fee is charged to the recipient,SHA or 1000 = Fee is shared between sender and recipient, with sender is charged Rp 1.000,00and the recipient will be charged the restremarkStringO50Remark/transaction descriptionsourceAccountNoStringM19Beneficiary Bank CodetransactionDateStringM25Format transaction date : (ISO 8601) YYYY-MM-DDThh:mm:ssbulkObjectArray of ObjectspartnerReferenceNoStringM64Individual transaction identifier on service consumer (bulk transaction sender) systembankCodeStringM11Bank code based on Bank Indonesia CodebeneficiaryAccountNoStringM34Beneficiary Account NumberbeneficiaryAccountNameStringM100Beneficiary Account NameamountObjectMvalueStringM16,2Net amount of the transaction.If it’s IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)originatorInfosArray of ObjectsCThe Originator Customer Account DetailsTo be filled if there is a request from the sender or if the consent from sender has been granted.This is subject to Article 8 paragraph 5 of Law No. 3 of 2011 concerning Fund Transfers.Also check for other provisions, such as the PPATK regulation.originatorCustomerNoStringM34Originator customer account numberoriginatorCustomerNameStringM100Originatorcustomer account nameoriginatorBankCodeStringM11Originator Bank CodeadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionbulkIdStringC64Transaction identifier on service provider (bulk transaction processing) system. Must be filled upon successful transactionpartnerBulkIdStringO64Transaction identifier on service consumer (bulk transaction sender) systemadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Interbank Bulk Transfer - Notification

**Request Body**

ParameterData TypeMandatoryLengthDescriptionbulkIdStringM64Transaction identifier on bulk transaction processing systempartnerBulkIdStringM64Transaction identifier on bulk transaction sender systembulkObjectArray of ObjectsoriginalReferenceNoStringC64Individual transaction identifier on bulk transaction processing system. Must be filled upon successful transactionoriginalPartnerReferenceNoStringO64Individual transaction identifier on bulk transaction sender systemresponseCodeStringM7Response CoderesponseMessageStringM150Response descriptionamountObjectOvalueStringM16,2Net amount of the transaction.If it’s IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)beneficiaryAccountNoStringM19Beneficiary AccountbeneficiaryBankCodeStringO11Beneficiary Bank CodesourceAccountNoStringM19Source Account NumbertraceNoStringO16Number for tracking to destination bankadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionbulkIdStringM64Transaction identifier on bulk transaction processing systempartnerBulkIdStringM64Transaction identifier on bulk transaction sender systemadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Transfer RTGS

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringM64Transaction identifier on service consumer systemamountObjectMvalueStringM16,2Net amount of the transaction.If it’s IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)beneficiaryAccountNameStringM100Beneficiary Account NamebeneficiaryAccountNoStringM34Beneficiary AccountbeneficiaryAddressStringO100Beneficiary AddressbeneficiaryBankCodeStringM11Beneficiary Bank CodebeneficiaryBankNameStringO50Beneficiary Bank NamebeneficiaryCustomerResidenceStringM1Beneficiary Customer Residence 1.Indonesia 2. Non IndonesiabeneficiaryCustomerTypeStringM1Beneficiary Customer Type1\. Individual2\. corporation3\. GovernmentbeneficiaryEmailStringO50Beneficiary EmailcurrencyStringO3Currency (ISO4217)customerReferenceStringO30Reference Number / No Referral / Transaction IDfeeTypeStringO25to whom the fee will be chargedkodeposStringO10Sender City (kodepos)receiverPhoneStringO20Beneficiary Customer PhoneremarkStringO50Remark/transaction descriptionsenderCustomerResidenceStringO1Beneficiary Customer Residence 1.Indonesia 2. Non IndonesiasenderCustomerTypeStringO1Beneficiary Customer Type1\. Individual2\. corporation3\. GovernmentsenderPhoneStringO20Source Customer PhonesourceAccountNoStringM19Beneficiary Bank CodeoriginatorInfosArray of ObjectsCThe Originator Customer Account DetailsTo be filled if there is a request from the sender or if the consent from sender has been granted.This is subject to Article 8 paragraph 5 of Law No. 3 of 2011 concerning Fund Transfers.Also check for other provisions, such as the PPATK regulation.originatorCustomerNoStringM34Originator customer account numberoriginatorCustomerNameStringM100Originatorcustomer account nameoriginatorBankCodeStringM11Originator Bank CodetransactionDateStringM25Format transaction date : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionpartnerReferenceNoStringO64Transaction identifier on service consumer systemamountObjectOvalueStringM16,2Net amount of the transaction.If it’s IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)beneficiaryAccountNameStringM100Beneficiary Account NamebeneficiaryAccountNoStringM34Beneficiary AccountbeneficiaryAccountTypeStringO1Beneficiary Account Type“D” for Current Account“S” for Saving AccountbeneficiaryBankCodeStringO11Beneficiary Bank CodecurrencyStringO3Currency (ISO4217)customerReferenceStringO30Reference Number / No Referral / Transaction IDsourceAccountNoStringM19Beneficiary Bank CodeoriginatorInfosArray of ObjectsCThe Originator Customer Account DetailsTo be filled if there is a request from the sender or if the consent from sender has been granted.This is subject to Article 8 paragraph 5 of Law No. 3 of 2011 concerning Fund Transfers.Also check for other provisions, such as the PPATK regulation.originatorCustomerNoStringM34Originator customer account numberoriginatorCustomerNameStringM100Originatorcustomer account nameoriginatorBankCodeStringM11Originator Bank CodetraceNoStringO16Number for tracking to destination banktransactionDateStringM25Format transaction date : (ISO 8601) YYYY-MM-DDThh:mm:sstransactionStatusStringM200 - Success01 - Initiated02 - Paying03 - Pending04 - Refunded05 - Canceled06 - Failed07 - Not foundtransactionStatusDescStringO50Description status transactionadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API RTGS - Notification

**Request Body**

ParameterData TypeMandatoryLengthDescriptionoriginalPartnerReferenceNoStringO64Original transaction identifier on service consumer systemoriginalReferenceNoStringO64Original transaction identifier on service provider systemoriginalExternalIdStringO19Original CustomerReference NumberlatestTransactionStatusStringM200 - Success01 - Initiated02 - Paying03 - Pending04 - Refunded05 - Canceled06 - Failed07 - Not foundamountObjectOvalueStringM16,2Net amount of the transaction.If it’s IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency(ISO 4217)beneficiaryAccountNameStringM100Beneficiary Account NamebeneficiaryAccountNoStringM34Beneficiary AccountbeneficiaryBankCodeStringM11Beneficiary Bank CodesourceAccountNoStringM19Beneficiary Bank CodetransactionDateStringM25Format transaction date : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response description

##### API Transfer SKNBI

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringM64Transaction identifier on service consumer systemamountObjectOvalueStringM16,2Net amount of the transaction.If it’s IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO 4217)beneficiaryAccountNameStringM100Beneficiary Account NamebeneficiaryAccountNoStringM34Beneficiary AccountbeneficiaryAddressStringO100Beneficiary AddressbeneficiaryBankCodeStringM11Beneficiary Bank CodebeneficiaryBankNameStringO50Beneficiary Bank NamebeneficiaryCustomerResidenceStringM1Beneficiary Customer Residence 1.Indonesia 2. Non IndonesiabeneficiaryCustomerTypeStringM1Beneficiary Customer Type1\. Individual2\. corporation3\. GovernmentbeneficiaryEmailStringO50Beneficiary EmailcurrencyStringO3Currency (ISO4217)customerReferenceStringO30Reference Number / No Referral / Transaction IDfeeTypeStringO25to whom the fee will be chargedkodeposStringO10Sender City (kodepos)receiverPhoneStringO20Beneficiary Customer PhoneremarkStringO50Remark/transaction descriptionsenderCustomerResidenceStringO1Beneficiary Customer Residence 1.Indonesia 2. Non IndonesiasenderCustomerTypeStringO1Beneficiary Customer Type1\. Individual2\. corporation3\. GovernmentsenderPhoneStringO20Source Customer PhonesourceAccountNoStringM19Beneficiary Bank CodeoriginatorInfosArray of ObjectsCThe Originator Customer Account DetailsTo be filled if there is a request from the sender or if the consent from sender has been granted.This is subject to Article 8 paragraph 5 of Law No. 3 of 2011 concerning Fund Transfers.Also check for other provisions, such as the PPATK regulation.originatorCustomerNoStringM34Originator customer account numberoriginatorCustomerNameStringM100Originatorcustomer account nameoriginatorBankCodeStringM11Originator Bank CodetransactionDateStringM25Format transaction date : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionpartnerReferenceNoStringO64Transaction identifier on service consumer systemamountObjectOvalueStringM16,2Net amount of the transaction.If it’s IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)beneficiaryAccountNameStringM100Beneficiary Account NamebeneficiaryAccountNoStringM34Beneficiary AccountbeneficiaryAccountTypeStringO1Beneficiary Account Type“D” for Current Account“S” for Saving AccountbeneficiaryBankCodeStringO11Beneficiary Bank CodecurrencyStringO3Currency (ISO4217)customerReferenceStringO30Reference Number / No Referral / Transaction IDsourceAccountNoStringM19Beneficiary Bank CodeoriginatorInfosArray of ObjectsCThe Originator Customer Account DetailsTo be filled if there is a request from the sender or if the consent from sender has been granted.This is subject to Article 8 paragraph 5 of Law No. 3 of 2011 concerning Fund Transfers.Also check for other provisions, such as the PPATK regulation.originatorCustomerNoStringM34Originator customer account numberoriginatorCustomerNameStringM100Originatorcustomer account nameoriginatorBankCodeStringM11Originator Bank CodetraceNoStringO16Number for tracking to destination banktransactionDateStringM25Format transaction date : (ISO 8601) YYYY-MM-DDThh:mm:sstransactionStatusStringM200 - Success01 - Initiated02 - Paying03 - Pending04 - Refunded05 - Canceled06 - Failed07 - Not foundtransactionStatusDescStringO50Description status transactionadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API SKNBI - Notification

**Request Body**

ParameterData TypeMandatoryLengthDescriptionoriginalPartnerReferenceNoStringO64Original transaction identifier on service consumer systemoriginalReferenceNoStringO64Original transaction identifier on service provider systemoriginalExternalIdStringO19Original CustomerReference NumberlatestTransactionStatusStringM200 - Success01 - Initiated02 - Paying03 - Pending04 - Refunded05 - Canceled06 - Failed07 - Not foundamountObjectOvalueStringM16,2Net amount of the transaction.If it’s IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency(ISO 4217)beneficiaryAccountNameStringM100Beneficiary Account NamebeneficiaryAccountNoStringM34Beneficiary AccountbeneficiaryBankCodeStringM11Beneficiary Bank CodesourceAccountNoStringM19Beneficiary Bank CodetransactionDateStringM25Format transaction date : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response description

#### CODE SNIPPETS

#### Code Snippets API Transfer

##### API Intrabank Transfer

**Sample Request**

```http
POST …/v1.0/transfer-intrabank HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T10:30:24+07:00
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
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "beneficiaryAccountNo":"888801000003301",
   "beneficiaryEmail":"yories.yolanda@work.bri.co.id ",
   "currency":"IDR",
   "customerReference":"10052019",
   "feeType":"BEN",
   "remark":"remark test",
   "sourceAccountNo":"888801000157508",
   "transactionDate":"2019-07-03T12:08:56+07:00",
   "originatorInfos":[{
   "originatorCustomerNo":"999901000003300",
   "originatorCustomerName":"Hafizh",
   "originatorBankCode":"002"},
   {
   "originatorCustomerNo":"999901000003301",
   "originatorCustomerName":"iin",
   "originatorBankCode":"002"
   },
   {
   "originatorCustomerNo":"999901000003302",
   "originatorCustomerName":"Renaldy",
   "originatorBankCode":"002"
   },
   {
   "originatorCustomerNo":"999901000003303",
   "originatorCustomerName":"Aufar",
   "originatorBankCode":"002"
   }],
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-21T10:30:34+07:00

{
   "responseCode":"2001700",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "beneficiaryAccountNo":"888801000003301",
   "currency":"IDR",
   "customerReference":"10052019",
   "sourceAccount":"888801000157508",
   "transactionDate":"2019-07-03T12:08:56+07:00",
   "originatorInfos":[{
   "originatorCustomerNo":"999901000003300",
   "originatorCustomerName":"Hafizh",
   "originatorBankCode":"002"},
   {
   "originatorCustomerNo":"999901000003301",
   "originatorCustomerName":"iin",
   "originatorBankCode":"002"
   },
   {
   "originatorCustomerNo":"999901000003302",
   "originatorCustomerName":"Renaldy",
   "originatorBankCode":"002"
   },
   {
   "originatorCustomerNo":"999901000003303",
   "originatorCustomerName":"Aufar",
   "originatorBankCode":"002"
   }],
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Interbank Transfer

**Sample Request**

```http
POST …/v1.0/transfer-interbank HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T13:59:21+07:00
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
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "beneficiaryAccountName":"Yories Yolanda",
   "beneficiaryAccountNo":"888801000003301",
   "beneficiaryAddress":"Palembang",
   "beneficiaryBankCode":"002",
   "beneficiaryBankName":"Bank BRI",
   "beneficiaryEmail":"yories.yolanda@work.bri.co.id",
   "currency":"IDR",
   "customerReference":"10052019",
   "sourceAccountNo":"888801000157508",
   "transactionDate":"2019-07-03T12:08:56+07:00",
   "feeType":"OUR",
   "originatorInfos":[{
   "originatorCustomerNo":"999901000003300",
   "originatorCustomerName":"Hafizh",
   "originatorBankCode":"002"},
   {
   "originatorCustomerNo":"999901000003301",
   "originatorCustomerName":"iin",
   "originatorBankCode":"002"
   },
   {
   "originatorCustomerNo":"999901000003302",
   "originatorCustomerName":"Renaldy",
   "originatorBankCode":"002"
   },
   {
   "originatorCustomerNo":"999901000003303",
   "originatorCustomerName":"Aufar",
   "originatorBankCode":"002"
   }],
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-21T13:59:40+07:00

{
   "responseCode":"2001800",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "beneficiaryAccountNo":"888801000003301",
   "beneficiaryBankCode":"002",
   "sourceAccountNo":"888801000157508",
   "traceNo":"10052019",
   "originatorInfos":[{
   "originatorCustomerNo":"999901000003300",
   "originatorCustomerName":"Hafizh",
   "originatorBankCode":"002" },
   {
   "originatorCustomerNo":"999901000003301",
   "originatorCustomerName":"iin",
   "originatorBankCode":"002"
   },
   {
   "originatorCustomerNo":"999901000003302",
   "originatorCustomerName":"Renaldy",
   "originatorBankCode":"002"
   },
   {
   "originatorCustomerNo":"999901000003303",
   "originatorCustomerName":"Aufar",
   "originatorBankCode":"002"
   }],
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Request for Payment

**Sample Request**

```http
POST …/v1.0/transfer-request-for-payment HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-22T08:01:16+07:00
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
   "beneficiaryAccountNo":"888801000003301",
   "beneficiaryAccountName":"Yories Yolanda",
   "remark":"remark test",
   "expiredDatetime":"2022-12-22T08:01:16+07:00",
   "sourceAccountNo":"888801000157508",
   "sourceAccountName":"Yories Yolanda",
   "currency":"IDR",
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "feeType":"BEN",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-22T07:45:11+07:00

{
   "responseCode":"2001900",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Interbank Bulk Transfer

**Sample Request**

```http
POST …/v1.0/transfer-interbank-bulk HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-22T07:41:11+07:00
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
   "currency":"IDR",
   "customerReference":"10052019",
   "feeType":"BEN",
   "remark":"remark test",
   "sourceAccountNo":"888801000157508",
   "transactionDate":"2019-07-03T12:08:56+07:00",
   "bulkObject":[
      {
         "partnerReferenceNo":"2020102900000000000990",
         "bankCode":"014",
         "beneficiaryAccountNo":"888801000003301",
         "beneficiaryAccountName":"Yories Yolanda",
         "amount":{
                    "value":"12345678.00",
                    "currency":"IDR"
                  },

         "originatorInfos":[{
                           "originatorCustomerNo":"999901000003300",
                           "originatorCustomerName":"Hafizh",
                           "originatorBankCode":"002" },
                           {
                           "originatorCustomerNo":"999901000003301",
                           "originatorCustomerName":"iin",
                           "originatorBankCode":"002"
                           },
                           {
                           "originatorCustomerNo":"999901000003302",
                           "originatorCustomerName":"Renaldy",
                           "originatorBankCode":"002"
                           },
                           {
                           "originatorCustomerNo":"999901000003303",
                           "originatorCustomerName":"Aufar",
                           "originatorBankCode":"002"
                           }]

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
X-TIMESTAMP: 2020-12-22T07:45:11+07:00
X-SIGNATURE: 85be8171923ac135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5

{
   "responseCode":"2002000",
   "responseMessage":"Request has been processed successfully",
   "bulkId":"2020102977770000000009",
   "partnerBulkId ":"2020102900000000000001",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Interbank Bulk Transfer - Notification

**Sample Request**

```http
POST …/v1.0/transfer-interbank-bulk/notify HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-22T07:53:16+07:00
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
         "originalReferenceNo":"2020102977770000000009",
         "originalPartnerReferenceNo":"2020102900000000000990",
         "responseCode":"2000000",
         "responseMessage":"Request has been processed successfully",
         "amount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "beneficiaryAccountNo":"888801000003301",
         "beneficiaryBankCode":"002",
         "sourceAccountNo":"888801000157508",
         "traceNo":"10052019"
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
X-TIMESTAMP: 2020-12-22T07:53:21+07:00

{
   "responseCode":"2002100",
   "responseMessage":"Request has been processed successfully",
   "bulkId":"2020102977770000000009",
   "partnerBulkId":"2020102900000000000001",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Transfer RTGS

**Sample Request**

```http
POST …/v1.0/transfer-rtgsHTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T14:06:21+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X- EXTERNAL-ID: 41807553358950093184162180797837
X- IP-ADDRESS: 172.24.281.24
X- DEVICE-ID: 09864ADCASA
X- LATITUDE: -6.1617169
X- LONGITUDE: 106.6643946
CHANNEL-ID: 95221

{
   "partnerReferenceNo":"2020102900000000000001",
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "beneficiaryAccountName":"Yories Yolanda",
   "beneficiaryAccountNo":"888801000003301",
   "beneficiaryAddress":"Palembang",
   "beneficiaryBankCode":"002",
   "beneficiaryBankName":"Bank BRI",
   "beneficiaryCustomerResidence":"1",
   "beneficiaryCustomerType":"1",
   "beneficiaryEmail":"yories.yolanda@work.bri.co.id",
   "currency":"IDR",
   "customerReference":"10052019",
   "feeType":"BEN",
   "kodepos":"12250",
   "receiverPhone":"080901020304",
   "remark":"remark test",
   "senderCustomerResidence":"1",
   "senderCustomerType":"1",
   "senderPhone":"080901020304",
   "sourceAccountNo":"888801000157508",
   "originatorInfos":[{
                           "originatorCustomerNo":"999901000003300",
                           "originatorCustomerName":"Hafizh",
                           "originatorBankCode":"002" },
                           {
                           "originatorCustomerNo":"999901000003301",
                           "originatorCustomerName":"iin",
                           "originatorBankCode":"002"
                           },
                           {
                           "originatorCustomerNo":"999901000003302",
                           "originatorCustomerName":"Renaldy",
                           "originatorBankCode":"002"
                           },
                           {
                           "originatorCustomerNo":"999901000003303",
                           "originatorCustomerName":"Aufar",
                           "originatorBankCode":"002"
                           }],
   "transactionDate":"2019-07-03T12:08:56+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-21T14:06:31+07:00
X-SIGNATURE: 85be8171923ac135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5

{
   "responseCode":"2002200",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "beneficiaryAccountName":"Yories Yolanda",
   "beneficiaryAccountNo":"888801000003301",
   "beneficiaryAccountType":"D",
   "beneficiaryBankCode":"002",
   "currency":"IDR",
   "customerReference":"10052019",
   "sourceAccountNo":"888801000157508",
   "originatorInfos":[{
                           "originatorCustomerNo":"999901000003300",
                           "originatorCustomerName":"Hafizh",
                           "originatorBankCode":"002" },
                           {
                           "originatorCustomerNo":"999901000003301",
                           "originatorCustomerName":"iin",
                           "originatorBankCode":"002"
                           },
                           {
                           "originatorCustomerNo":"999901000003302",
                           "originatorCustomerName":"Renaldy",
                           "originatorBankCode":"002"
                           },
                           {
                           "originatorCustomerNo":"999901000003303",
                           "originatorCustomerName":"Aufar",
                           "originatorBankCode":"002"
                           }],
   "traceNo":"10052019",
   "transactionDate":"2019-07-03T12:08:56+07:00",
   "transactionStatus":"00",
   "transactionStatusDesc":"success",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API RTGS - Notification

**Sample Request**

```http
POST …/v1.0/transfer-rtgs/notify HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-22T07:53:16+07:00
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
   "originalExternalId":"10052019",
   "latestTransactionStatus":"00",
   "amount":{
      "value":"10000.00",
      "currency":"IDR"
   },
   "beneficiaryAccountName":"Yories Yolanda",
   "beneficiaryAccountNo":"888801000003301",
   "beneficiaryBankCode":"002",
   "sourceAccountNo":"888801000157508",
   "transactionDate":"2020-12-21T14:06:21+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-22T07:53:21+07:00

{
   "responseCode":"2007600",
   "responseMessage":"Request has been processed successfully"
}

```

##### API Transfer SKNBI

**Sample Request**

```http
POST …/v1.0/transfer-skn HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T14:36:11+07:00
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
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "beneficiaryAccountName":"Yories Yolanda",
   "beneficiaryAccountNo":"888801000157508",
   "beneficiaryAddress":"Palembang",
   "beneficiaryBankCode":"002",
   "beneficiaryBankName":"Bank BRI",
   "beneficiaryCustomerResidence":"1",
   "beneficiaryCustomerType":"1",
   "beneficiaryEmail":"yories.yolanda@work.bri.co.id",
   "currency":"IDR",
   "customerReference":"10052019",
   "feeType":"BEN",
   "kodepos":"12250",
   "receiverPhone":"080901020304",
   "remark":"remark test",
   "senderCustomerResidence":"1",
   "senderCustomerType":"1",
   "senderPhone":"080901020304",
   "sourceAccountNo":"888801000157508",
   "originatorInfos":[{
                           "originatorCustomerNo":"999901000003300",
                           "originatorCustomerName":"Hafizh",
                           "originatorBankCode":"002" },
                           {
                           "originatorCustomerNo":"999901000003301",
                           "originatorCustomerName":"iin",
                           "originatorBankCode":"002"
                           },
                           {
                           "originatorCustomerNo":"999901000003302",
                           "originatorCustomerName":"Renaldy",
                           "originatorBankCode":"002"
                           },
                           {
                           "originatorCustomerNo":"999901000003303",
                           "originatorCustomerName":"Aufar",
                           "originatorBankCode":"002"
                           }],
   "transactionDate":"2019-07-03T12:08:56+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-21T14:36:19+07:00

{
   "responseCode":"2002300",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "beneficiaryAccountName":"Yories Yolanda",
   "beneficiaryAccountNo":"888801000003301",
   "beneficiaryAccountType":"D",
   "beneficiaryBankCode":"002",
   "currency":"IDR",
   "customerReference":"10052019",
   "sourceAccountNo":"888801000157508",
    "originatorInfos":[{
                           "originatorCustomerNo":"999901000003300",
                           "originatorCustomerName":"Hafizh",
                           "originatorBankCode":"002" },
                           {
                           "originatorCustomerNo":"999901000003301",
                           "originatorCustomerName":"iin",
                           "originatorBankCode":"002"
                           },
                           {
                           "originatorCustomerNo":"999901000003302",
                           "originatorCustomerName":"Renaldy",
                           "originatorBankCode":"002"
                           },
                           {
                           "originatorCustomerNo":"999901000003303",
                           "originatorCustomerName":"Aufar",
                           "originatorBankCode":"002"
                           }],
   "traceNo":"10052019",
   "transactionDate":"2019-07-03T12:08:56+07:00",
   "transactionStatus":"00",
   "transactionStatusDesc":"success",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API SKNBI - Notification

**Sample Request**

```http
POST …/v1.0/transfer-skn/notify HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-22T07:53:16+07:00
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
   "originalExternalId":"10052019",
   "latestTransactionStatus":"00",
   "amount":{
      "value":"10000.00",
      "currency":"IDR"
   },
   "beneficiaryAccountName":"Yories Yolanda",
   "beneficiaryAccountNo":"888801000003301",
   "beneficiaryBankCode":"002",
   "sourceAccountNo":"888801000157508",
   "transactionDate":"2020-12-21T14:06:21+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-22T07:53:21+07:00

{
   "responseCode":"2007500",
   "responseMessage":"Request has been processed successfully"
}

```

#### RESPONSES CODE

Response status merupakan informasi yang diberikan oleh service provider kepada service consumer pada response body, sebagai indikasi hasil dari pemrosesan request yang diterima.

Response status terdiri dari 2 komponen, yaitu kode (response code) dan deskripsinya (response message).

KomponenTipe DataLengthKeteranganresponseCodeString7response code = HTTP status code + service code + case coderesponseMessageString150

###### **Daftar** _**Response Code**_

CategoryHTTP CodeService CodeCase CodeResponse MessageDescriptionSuccess200any00SuccessfulSuccessfulSuccess202any00Request In ProgressTransaction still on processSystem400any00Bad RequestGeneral request failed error, including message parsing failed.Message400any01Invalid Field Format {field name}Invalid formatMessage400any02Invalid Mandatory Field {field name}Missing or invalid format on mandatory fieldSystem401any00Unauthorized. \[reason\]General unauthorized error (No Interface Def, API is Invalid, Oauth Failed, Verify Client Secret Fail, Client Forbidden Access API, Unknown Client, Key not Found)System401any01Invalid Token (B2B)Token found in request is invalid (Access Token Not Exist, Access Token Expiry)System401any02Invalid Customer TokenToken found in request is invalid (Access Token Not Exist, Access Token Expiry)System401any03Token Not Found (B2B)Token not found in the system. This occurs on any API that requires token as input parameterSystem401any04Customer Token Not FoundToken not found in the system. This occurs on any API that requires token as input parameterBusiness403any00Transaction ExpiredTransaction expiredSystem403any01Feature Not Allowed \[Reason\]This merchant is not allowed to call Direct Debit APIsBusiness403any02Exceeds Transaction Amount LimitExceeds Transaction Amount LimitBusiness403any03Suspected FraudSuspected FraudBusiness403any04Activity Count Limit ExceededToo many request, Exceeds Transaction Frequency LimitBusiness403any05Do Not HonorAccount or User status is abnormalSystem403any06Feature Not Allowed At This Time. \[reason\]Cut off In ProgressBusiness403any07Card BlockedThe payment card is blockedBusiness403any08Card ExpiredThe payment card is expiredBusiness403any09Dormant AccountThe account is dormantBusiness403any10Need To Set Token LimitNeed to set token limitSystem403any11OTP BlockedOTP has been blockedSystem403any12OTP Lifetime ExpiredOTP has been expiredSystem403any13OTP Sent To Cardholerinitiates request OTP to the issuerBusiness403any14Insufficient FundsInsufficient FundsBusiness403any15Transaction Not Permitted.\[reason\]Transaction Not PermittedBusiness403any16Suspend TransactionSuspend TransactionBusiness403any17Token Limit ExceededPurchase amount exceeds the token limit set priorBusiness403any18Inactive Card/Account/CustomerIndicates inactive accountBusiness403any19Merchant BlacklistedMerchant is suspended from calling any APIsBusiness403any20Merchant Limit ExceedMerchant aggregated purchase amount on that day exceeds the agreed limitBusiness403any21Set Limit Not AllowedSet limit not allowed on particular tokenBusiness403any22Token Limit InvalidThe token limit desired by the merchant is not within the agreed range between the merchant and the IssuerBusiness403any23Account Limit ExceedAccount aggregated purchase amount on that day exceeds the agreed limitBusiness404any00Invalid Transaction StatusInvalid transaction statusBusiness404any01Transaction Not FoundTransaction not foundSystem404any02Invalid RoutingInvalid RoutingSystem404any03Bank Not Supported By SwitchBank not supported by switchBusiness404any04Transaction CancelledTransaction is cancelled by customerBusiness404any05Merchant Is Not Registered For Card Registration ServicesMerchant is not registered for Card Registration servicesSystem404any06Need To Request OTPNeed to request OTPSystem404any07Journey Not FoundThe journeyId cannot be found in the systemBusiness404any08Invalid MerchantMerchant does not exist or status abnormalBusiness404any09No IssuerNo issurSystem404any10Invalid API TransitionInvalid API transition within a journeyBusiness404any11Invalid Card/Account/Customer \[info\]/Virtual AccountCard information may be invalid, or the card account may be blacklisted, or Virtual Account number maybe invalid.Business404any12Invalid Bill/Virtual Account \[Reason\]The bill is blocked/ suspended/not found.Virtual account is suspend/not found.Business404any13Invalid AmountThe amount doesn't match with what supposed toBusiness404any14Paid BillThe bill has been paidSystem404any15Invalid OTPOTP is incorrectBusiness404any16Partner Not FoundPartner number can't be foundBusiness404any17Invalid TerminalTerminal does not exist in the systemBusiness404any18Inconsistent RequestInconsistent request parameter found for the same partner reference number/transaction idIt can be considered as failed in transfer debit, but it should be considered as success in transfer credit.Considered as success:\- Transfer credit = (i) Intrabank transfer; (ii) Interbank transfer; (iii) RTGS transfer; (iv) SKNBI transfer;\- Virtual account = (i) Payment VA; (ii) Payment to VA;\- Transfer debit = (i) Refund payment; (ii) Void;Considered as failed:\- Transfer credit = (i) Transfer to OTC;\- Transfer debit = (i) Direct debit payment; (ii) QR CPM payment; (iii) Auth payment; (iv) Capture;Business404any19Invalid Bill/Virtual AccountThe bill is expired.Virtual account is expired.System405any00Requested Function Is Not SupportedRequested function is not supportedBusiness405any01Requested Opearation Is Not AllowedRequested operation to cancel/refund transaction Is not allowed at this time.System409any00ConflictCannot use same X-EXTERNAL-ID in same daySystem409any01Duplicate partnerReferenceNoTransaction has previously been processed indicates the same partnerReferenceNo already successSystem429any00Too Many RequestsMaximum transaction limit exceededSystem500any00General ErrorGeneral ErrorSystem500Any01Internal Server ErrorUnknown Internal Server Failure, Please retry the process againSystem500Any02External Server ErrorBackend system failure, etcSystem504any00Timeouttimeout from the issuer

#### APLIKASI PENGUJIAN

×

Akses Terbatas, Mohon Sign Up untuk Dapat Mengakses Halaman Ini

© 2025 - SNAP Developer Site