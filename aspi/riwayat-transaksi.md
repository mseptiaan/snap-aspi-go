Riwayat Transaksi - SNAP Developer Site

[![](/img/icon/snap-header-white-100.png)](/)

- [**Beranda**](/)
- [**Info**](/info)
- [**API**](/api-services)
- [**Direktori Publikasi**](/direktori-publikasi)
  - [**Sign In**](/sign-in)

[![](/img/icon/bi-logo-full-100.png)](/)

![](/img/card/transaction-history-100.png)

### Riwayat Transaksi

Overview


Guides


Code Snippets


Response Code


Aplikasi Pengujian


#### OVERVIEW

#### API Transaction History

API Transaction History (Riwayat Transaksi) diperlukan agar Konsumen, Non-PJP Pengguna Layanan, atau PJP PIAS dapat mengakses informasi riwayat transaksi dari rekening yang dimiliki secara _real time_, sesuai dengan layanan yang disediakan oleh PJP AIS. Informasi riwayat transaksi berisi rincian mengenai transaksi kredit maupun debit, saldo rekening, dan beberapa informasi lainnya.

API Bank Statement merupakan API yang digunakan untuk mengakses riwayat transaksi keuangan secara menyeluruh dari sebuah rekening Bank, baik rekening milik individu pemilik rekening. Konsumen dapat memanfaatkan _platform_ dari Non-PJP Pengguna Layanan, atau PJP PIAS yang telah disertifikasi oleh PJP AIS untuk mengakses API ini sehingga dapat dengan mudah melihat rincian keuangannya dalam satu aplikasi.

#### SKENARIO PENGGUNAAN API TRANSACTION HISTORY

![](/img/docs/3_1_use_case_diagram_api_transaction_history.png)

_Use Case Diagram_ API Transaction History

Merujuk pada _use case diagram_, pemilik akun yaitu Konsumen, Non-PJP Pengguna Layanan, PJP PIAS, PJP AIS Lembaga Selain Bank dapat mengakses informasi riwayat transaksi (transaction history list) dan/atau rincian dari riwayat transaksi (transaction history detail) atas rekeningnya menggunakan API Riwayat Transaksi yang disediakan oleh PJP PIAS atau PJP AIS. Konsumen, PJP PIAS, atau PJP AIS Lembaga Selain Bank dapat menggunakan API Bank Statement yang disediakan oleh PJP AIS melalui pihak yang sudah tersertifikasi untuk mengakses API ini.

API Transaction History dapat digunakan dalam sejumlah skema sebagai berikut:

**Skema 1: Konsumen - Non-PJP Pengguna Layanan - PJP AIS Lembaga Selain Bank**

Dalam skema ini, Konsumen mengakses riwayat transaksi yang tercatat pada PJP AIS Lembaga Selain Bank yang dilakukan melalui Non-PJP Pengguna Layanan sebagai kanal pembayaran. Mekanisme ini dapat dilakukan setelah dilakukan proses _card registration/account binding_. Pada saat pengaksesan riwayat transaksi, diperlukan permintaan persetujuan Konsumen (otorisasi/otentikasi) secara eksplisit menggunakan mekanisme OAuth 2.0 untuk memastikan kebenaran informasi Konsumen serta agar tidak ada penyalahgunaan data dan hak dari Konsumen. Setelah mendapatkan persetujuan dari Konsumen, Non-PJP Pengguna Layanan akan mendapatkan access token sebagai credential dari PJP AIS Lembaga Selain Bank untuk mengakses data Konsumen tersebut.

Dalam skema ini, Konsumen dapat mengakses riwayat transaksi yang transaksinya dilakukan melalui PJP AIS Lembaga Selain Bank dan Non-PJP Pengguna Layanan tersebut saja. Dalam hal ini, PJP AIS Lembaga Selain Bank tidak menyertakan riwayat transaksi yang dilakukan melalui kanal pembayaran lain.

API yang dapat digunakan dalam skema ini adalah API _Transaction History List_ dan _Transaction History Detail_. API \_Transaction History List\_mengembalikan daftar riwayat transaksi, sedangkan API _Transaction History Detail_ mengembalikan informasi lengkap dari 1 riwayat transaksi.

**Skema 2: Konsumen - Non-PJP Pengguna Layanan - PJP PIAS**

Dalam skema ini, Konsumen mengakses riwayat transaksi yang tercatat pada PJP PIAS yang dilakukan melalui suatu Non-PJP Pengguna Layanan sebagai kanal pembayaran. PJP PIAS pada umumnya menyediakan beberapa kanal pembayaran seperti kartu kredit/debit, transfer ke _virtual account_, hingga pembayaran _offline_ pada tempat tertentu seperti minimarket.

Konsumen tidak memiliki akun khusus pada PJP PIAS dan Konsumen memasukkan beberapa informasi setiap kali melakukan transaksi. Konsumen dapat mengakses riwayat transaksi yang dilakukan pada Non-PJP Pengguna Layanan tersebut dengan menggunakan API _Transaction History List_ dan _Transaction History Detail_. Non-PJP Pengguna Layanan cukup menggunakan _credentials_ yang telah disediakan oleh PJP PIAS.

**Skema 3: Konsumen - PJP PIAS - PJP AIS Bank**

Dalam skema ini, Konsumen mengakses riwayat transaksi melalui layanan rekening koran pada PJP AIS Bank yang dilakukan melalui PJP PIAS sebagai kanal pembayaran. Konsumen dapat mengakses riwayat transaksi yang dilakukan dengan menggunakan API Bank Statement.

API ini hanya bisa digunakan oleh pihak yang sudah tersertifikasi dan lolos pemeriksaan oleh PJP AIS Bank sebagai penyedia layanan karena mengandung data pribadi dan rahasia.

Pada saat pengaksesan riwayat transaksi, diperlukan permintaan persetujuan Konsumen (otorisasi/otentikasi) secara eksplisit menggunakan mekanisme OAuth 2.0 untuk memastikan kebenaran informasi pengguna serta agar tidak ada penyalahgunaan data.

Setelah mendapatkan persetujuan dari Konsumen, PJP PIAS akan mendapatkan _access token_ sebagai _credential_ dari PJP AIS Bank untuk mengakses data.

**Skema 4: PJP AIS Lembaga Selain Bank/PJP PIAS - PJP AIS Bank**

Dalam skema ini, PJP AIS Lembaga Selain Bank/PJP PIAS dapat mengakses riwayat transaksi rekeningnya melalui layanan rekening koran pada PJP AIS Bank dengan menggunakan API Bank Statement.

##### API Transaction History List

Ketentuan dan Keterbatasan API

PengurutanDESCpageSize Maksimal50Rentang Waktu Maksimal1 bulanRiwayat Transaksi Tertua1 tahun

![](/img/docs/3_2_sequence_diagram_untuk_api_transaction_history_list.png)

Sequence Diagram untuk API Transaction History List

**Informasi Umum**

Service Code12NameAPI Transaction History ListVersion1.0HTTP MethodPOSTPath.../{version}/transaction-history-list

##### API Transaction History Detail

Ketentuan dan Keterbatasan API

Riwayat Transaksi Tertua1 tahun

![](/img/docs/3_3_sequence_diagram_untuk_api_transaction_history_detail.png)

Sequence Diagram untuk API Transaction History Detail

**Informasi Umum**

Service Code13NameAPI Transaction History DetailVersion1.0HTTP MethodPOSTPath.../{version}/transaction-history-detail

##### API Bank Statement

Ketentuan dan Keterbatasan API

PengurutanDESCRentang Waktu Maksimal1 bulanRiwayat Transaksi Tertua1 tahun

![](/img/docs/3_4_sequence_diagram_untuk_api_bank_statement.png)

Sequence Diagram untuk API Bank Statement

**Informasi Umum**

Service Code14NameAPI Bank StatementVersion1.0HTTP MethodPOSTPath.../{version}/bank-statement

#### GUIDES

#### Spesifikasi Parameter Header dan Body API Transaction History

##### API Transaction History List

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringO64Transaction identifier on service consumer systemfromDateTimeStringO25Starting time range.Default:NOW (DESC)or NOW - 1 months (ASC) Format from date time: (ISO 8601) YYYY-MM-DDThh:mm:sstoDateTimeStringO25Ending time range.Default:NOW - 1 months (DESC)or NOW (ASC) Format to date time: (ISO 8601) YYYY-MM-DDThh:mm:sspageSizeIntegerO2Maximum number of transaction returned in one pagination.Default: 10pageNumberIntegerO2Current page number.Default: 0additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringO64Transaction identifier on service provider systempartnerReferenceNoStringO64Transaction identifier on service consumer systemdetailDataArraydateTimeStringO25Format date time: (ISO 8601) YYYY-MM-DDThh:mm:ssamountObjectOvalueStringM16,2Net amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency(ISO 4217)remarkStringO256Transaction remark.sourceOfFundsList< SourceOfFund >ON/ASource of funds used for this transaction.See object definition.statusStringM32Transaction status.INIT, SUCCESS, CLOSED, CANCELLEDtypeStringM32Transaction type.PAYMENT, REFUND, TOP\_UP, SEND\_MONEY, RECEIVE\_MONEY, DISBURSMENT, etcadditionalInfoObjectOAdditional information for custom use from detailDataadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Transaction History Detail

**Request Body**

ParameterData TypeMandatoryLengthDescriptionoriginalPartnerReferenceNoStringM64Transaction identifier on service consumer systemadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionpartnerReferenceNoStringO64Transaction identifier on service consumer systemamountObjectOvalueStringM16,2Net amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency(ISO 4217)cancelledTimeStringO25Transaction cancelled time. Format date: (ISO 8601) YYYY-MM-DDThh:mm:ssdateTimeStringM25Transaction created time. Format date: (ISO 8601) YYYY-MM-DDThh:mm:ssrefundAmountObjectOvalueStringM16,2Refund amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency(ISO 4217)remarkStringO256Transaction remark.sourceOfFundsList< SourceOfFund >ON/ASource of funds used for this transaction.See object definition.statusStringM32Transaction status.INIT, SUCCESS, CLOSED, CANCELLEDtypeStringM32Transaction type.PAYMENT, REFUND, TOP\_UP, SEND\_MONEY, RECEIVE\_MONEY, etcadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Bank Statement

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringO64Transaction identifier on service consumer systembankCardTokenStringC128Card token for payment.Must be fille if accountNo is null.Bank account numberaccountNoStringC16Must be filled if bank CardToken is nullfromDateTimeStringO25Starting time range. Default: NOW (DESC) or NOW – 1 months (ASC)Format from date time :(ISO 8601) YYYY-MMDDThh:mm:sstoDateTimeStringO25Ending time range. Default: NOW – 1 months (DESC) or NOW (ASC)Format to date time: (ISO 8601) YYYYMMDDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionpartnerReferenceNoStringO64Transaction identifier on service consumer systembalanceArray of ObjectON/AStarting and ending balance before the first/last transaction.amountObjectMvalueStringM16,2Transaction Amount.Nominal inputted by Customer with 2 decimalcurrencyStringM3Currency(ISO 4217)dateTimeStringM25Format date time: (ISO 8601) YYYY-MM-DDThh:mm:ssstartingBalanceObjectMvalueStringM16,2Transaction Amount.Nominal inputted by Customer with 2 decimalcurrencyStringM3Currency(ISO 4217)dateTimeStringM25Format date time: (ISO 8601) YYYY-MM-DDThh:mm:ssendingBalanceObjectMvalueStringM16,2Transaction Amount.Nominal inputted by Customer with 2 decimalcurrencyStringM3Currency(ISO 4217)dateTimeStringM25Format date time: (ISO 8601) YYYY-MM-DDThh:mm:sstotalCreditEntriesDebitAndCreditEntriesON/ATotal transaction amount with type = CREDIT. See object definition.totalDebitEntriesDebitAndCreditEntriesON/ATotal transaction amount with type = DEBIT. See object definition.hasMoreStringO1Memberikan informasi apakah masih ada record yang belum ditampilkan dalam rentang fromDateTime toDateTimeY = Yes N = No.lastRecordDateTimeStringO25Format last record date time: (ISO 8601) YYYY-MMDDThh:mm:ssdetailDataArray of objectdetailBalanceObjectON/AStarting and ending balance before and after transaction.startAmountArray of ObjectOendAmountArray of ObjectOamountObjectOvalueStringM16,2Net amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)originAmountObjectOvalueStringM16,2Origin net amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)transactionDateStringM25Timestamp of the Transaction. Format transactiondate: (ISO 8601) YYYY-MM-DDThh:mm:ssremarkStringM256Transaction remark.transactionIdStringO35Internal transaction identifier from publisher perspective.typeStringM6Transaction type.CREDIT/DEBITtransactionDetailStatusStringO20Transaction detail indicator (original transaction or error correction)SUCCESS/ERROR CORECTIONdetailInfoObjectOAdditional information of detail transactionadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### Definisi Tipe

**ResultInfo**

ParameterData TypeMandatoryLengthDescriptionresultCodeIdStringM8Result code unique identifier.resultCodeStringM64Result code.resultStatusStringM1Result status.S/F/UresultMsgStringO256Result message, can be filled with the reason of error.

**SourceOfFund**

ParameterData TypeMandatoryLengthDescriptionsourceStringM32Source of fund.BALANCE/etc.amountObjectOvalueStringM16,2Net amount of the transaction.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)

**CashBalance**

ParameterData TypeMandatoryLengthDescriptionamountObjectOvalueStringM16,2Amount of balance.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)dateTimeISODateTimeM25Timestamp of the balance.ISO-8601

**DetailBalance**

ParameterData TypeMandatoryLengthDescriptionamountObjectOvalueStringM16,2Amount of balance.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency (ISO4217)

**ActiveCurrencyAndAmount**

ParameterData TypeMandatoryLengthDescriptionvalueStringM18Value of the amount.If it's IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency. (ISO4217)

**DebitAndCreditEntries**

ParameterData TypeMandatoryLengthDescriptionnumberOfEntriesintO5Number of entriesamountObjectvalueStringM18Value of the amount.If it�s IDR then value includes 2 decimal digits.e.g. IDR 10.000,- will be placed with 10000.00currencyStringM3Currency. (ISO4217)

#### CODE SNIPPETS

#### Code Snippets API Transaction History

##### API Transaction History List

**Sample Request**

```http
POST .../v1.0/transaction-history-list HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-18T15:34:40+07:00
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
   "fromDateTime":"2019-07-03T12:08:56+07:00",
   "toDateTime":"2019-07-03T12:08:56+07:00",
   "pageSize":"10",
   "pageNumber":"2",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-18T15:34:44+07:00

{
   "responseCode":"2001200",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "detailData":[
      {
         "dateTime":"2019-07-03T12:08:56+07:00",
         "amount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "remark":"Payment to Warung Ikan Bakar",
         "sourceOfFunds":[
            {
               "source":"BALANCE",
               "amount":{
                  "value":"10000.00",
                  "currency":"IDR"
               }
            }
         ],
         "status":"SUCCESS",
         "type":"PAYMENT",
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

##### API Transaction History Detail

**Sample Request**

```http
POST .../v1.0/transaction-history-detail HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-18T15:55:40+07:00
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
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-18T15:55:47+07:00

{
   "responseCode":"2001300",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "cancelledTime":"2009-07-03T12:08:56+07:00",
   "dateTime":"2009-07-03T12:08:56+07:00",
   "refundAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "remark":"Payment to Warung Ikan Bakar",
   "sourceOfFunds":[
      {
         "source":"BALANCE",
         "amount":{
            "value":"10000.00",
            "currency":"IDR"
         }
      }
   ],
   "status":"SUCCESS",
   "type":"PAYMENT",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Bank Statement

**Sample Request**

```http
POST .../v1.0/bank-statement HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-18T16:03:40+07:00
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
   "bankCardToken":"6d7963617264746f6b656e",
   "accountNo":"7382382957893840",
   "fromDateTime":"2019-07-03T12:08:56+07:00",
   "toDateTime":"2019-07-03T12:08:56+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-18T16:03:45+07:00

{
   "responseCode":"2001400",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "balance":[
      {
         "amount":{
            "value":"10000.00",
            "currency":"IDR",
            "dateTime":"2020-12-18T16:03:45+07:00"
         },
         "startingBalance":{
            "value":"12345678.00",
            "currency":"IDR",
            "dateTime":"2020-12-18T16:03:45+07:00"
         },
         "endingBalance":{
            "value":"12345678.00",
            "currency":"IDR",
            "dateTime":"2020-12-19T16:03:45+07:00"
         }
      }
   ],
   "totalCreditEntries":{
      "numberOfEntries":"10",
      "amount":{
         "value":"10000.00",
         "currency":"IDR"
      }
   },
   "totalDebitEntries":{
      "numberOfEntries":"10",
      "amount":{
         "value":"10000.00",
         "currency":"IDR"
      }
   },
   "hasMore":"Y",
   "lastRecordDateTime":"2021-11-24T11:54:56+07:00",
   "detailData":[
      {
         "detailBalance":{
            "startAmount":[
               {
                  "amount":{
                     "value":"10000.00",
                     "currency":"IDR"
                  }
               }
            ],
            "endAmount":[
               {
                  "amount":{
                     "value":"10000.00",
                     "currency":"IDR"
                  }
               }
            ]
         },
         "amount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "originAmount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "transactionDate":"2009-07-03T12:08:56+07:00",
         "remark":"Payment to Warung Ikan Bakar",
         "transactionId":"20200801198230912830091123",
         "type":"CREDIT",
         "transactionDetailStatus":"SUCCESS",
         "detailInfo":{
            "page":"12"
         }
      }
   ],
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