Registrasi - SNAP Developer Site

[![](/img/icon/snap-header-white-100.png)](/)

- [**Beranda**](/)
- [**Info**](/info)
- [**API**](/api-services)
- [**Direktori Publikasi**](/direktori-publikasi)
  - [**Sign In**](/sign-in)

[![](/img/icon/bi-logo-full-100.png)](/)

![](/img/card/registrasi-100.png)

### Registrasi

Overview


Guides


Code Snippets


Response Code


Aplikasi Pengujian


#### OVERVIEW

#### API Registrasi

API Registration (Registrasi) diperlukan agar Konsumen dapat melakukan pengaitan datanya untuk melakukan layanan transaksi pembayaran atau mengakses data miliknya. Data-data yang dikaitkan ini dapat beragam tergantung kebutuhan, seperti data kartu debit, kartu kredit, atau rekening.

##### **SKENARIO PENGGUNAAN API REGISTRATION**

![](/img/docs/1_1_use_case_diagram_api_card_registration_(via_pjp_pias).png)

_Use Case Diagram_ API Card Registration (via PJP PIAS)

![](/img/docs/1_2_use_case_diagram_api_card_registration_(direct_integration).png)

_Use Case Diagram_ API Card Registration (Direct Integration)

![](/img/docs/1_3_use_case_diagram_api_account_registration.png)

_Use Case Diagram_ API Account Registration

Merujuk pada _use case diagram_, Konsumen sebagai pemilik rekening atau kartu, dapat melakukan penautan rekening atau kartunya pada Non-PJP Pengguna Layanan, PJP AIns, dan/atau PJP PIAS untuk digunakan sebagai sumber dana dalam bertransaksi dan/atau untuk mengakses layanan lainnya seperti pengecekan saldo dan/atau pengecekan histori transaksi.

API Registrasi dapat digunakan dalam sejumlah skema sebagai berikut:

**Skema 1:** _**Indirect Integration**_ **(Konsumen – Non-PJP Pengguna Layanan / PJP AIns / PJP PIAS – PJP AIS)**

Pada skema _indirect integration_, proses _card registration_ dilakukan dengan cara PJP PIAS menginisiasi pembayaran sejumlah nominal tertentu. Apabila proses otorisasi yang dilakukan PJP AIS kepada Konsumen berhasil maka dapat diyakini kebenaran informasi detail kartu dan pemiliknya, sehingga dapat dilakukan _card registration_ pada Non-PJP Pengguna Layanan, PJP AIns, dan/atau PJP PIAS. Selanjutnya, PJP PIAS menginisiasi pengembalian dana atas pembayaran ( _refund_).

**Skema 2:** _**Direct Integration**_ **(Konsumen – Non-PJP Pengguna Layanan / PJP AIns / PJP PIAS – PJP AIS)**

Pada skema _direct integration_, proses _card registration_ dan _account registration_ dilakukan dengan cara Non-PJP Pengguna Layanan, PJP AIns, atau PJP PIAS melakukan inisiasi validasi atas informasi kartu/rekening dan pemiliknya ke PJP AIS selaku penerbit kartu dan/atau rekening. Apabila proses validasi dan otorisasi yang dilakukan PJP AIS pada Konsumen berhasil, maka dapat diyakini kebenaran informasi detail kartu/rekening dan pemiliknya, sehingga dapat dilakukan _card registration_ atau _account registration_ pada Non-PJP Pengguna Layanan, PJP AIns, atau PJP PIAS.

Sejumlah API yang digunakan dalam skema-skema tersebut antara lain:

- API _Card Registration_ adalah layanan yang digunakan untuk mendaftarkan data kartu milik Konsumen pada Non-PJP Pengguna Layanan, PJP AIns, atau PJP PIAS. Kartu milik Konsumen tersebut diterbitkan oleh PJP AIS dan menjadi sumber dana Konsumen dalam bertransaksi dan/atau untuk mengakses layanan lainnya. API ini dapat disediakan oleh PJP yang melakukan registrasi kartu. Dalam hubungan bisnisB2Cdiperlukan proses verifikasi untuk memastikan kebenaran data konsumen.

Dalam rangka melakukan validasi, dapat menggunakan mekanisme 3D Secure atau mekanisme validasi lainnya yang ditetapkan oleh penerbit kartu.

- API _Card Registration Inquiry_ dapat digunakan untuk menampilkan data kartu yang sudah diregistrasikan pada Non-PJP Pengguna Layanan, PJP AIns, atau PJP PIAS. _Inquiry_ dapat dilakukan dengan memberikan informasi pengenal Konsumen.

- API _Verify OTP_ merupakan API yang digunakan untuk melakukan verifikasi OTP baik untuk registrasi kartu. Verifikasi OTP ini dilakukan khususnya untuk hubungan yang melibatkan Konsumen untuk memastikan kebenaran data Konsumen yang melakukan registrasi. API ini dapat disediakan oleh PJP AIS yang menerbitkan kartu. API Verify OTP juga bisa digunakan untuk API Pembayaran yang menggunakan _sequence_ OTP.

- API _Card Registration Unbinding_ digunakan untuk melakukan penghapusan data kartu milik Konsumen yang telah diregistrasikan pada Non-PJP Pengguna Layanan, PJP AIns, atau PJP PIAS. Penghapusan data kartu dilakukan dengan memasukkan data kartu yang akan dihapuskan penautannya.

- API _Account Creation_ dapat digunakan oleh Konsumen untuk pembuatan akun pada PJP AIS melalui PJP AIns, PJP PIAS, dan/atau Non-PJP Pengguna Layanan. API ini digunakan dalam hubungan _B2C (Business to Customer)._

- API _Account Binding_ adalah layanan yang digunakan untuk mendaftarkan data akun yang diterbitkan oleh PJP AIS yang menjadi sumber dana transaksi atau dalam rangka mengakses layanan lainnya oleh Konsumen pada PJP AIns dan/atau PJP PIAS. API ini digunakan dalam _B2C (Business to Customer)._ Dalam hubungan bisnisB2Cdiperlukan proses verifikasi/otentikasi untuk memastikan kebenaran data Konsumen.

- API _Account Binding Inquiry_ dapat digunakan untuk menampilkan data rekening yang sudah diregistrasikan pada Non-PJP Pengguna Layanan, PJP AIns, dan/atau PJP PIAS. _Inquiry_ dapat dilakukan dengan memberikan informasi pengenal konsumen.

- API _Account Unbinding_ digunakan untuk melakukan penghapusan data rekening milik Konsumen yang telah diregistrasikan pada Non-PJP Pengguna Layanan, PJP AIns, atau PJP PIAS. Penghapusan data rekening dilakukan dengan memasukkan data rekening yang akan dihapuskan penautannya.


Penyelenggaraan layanan pembayaran berbasis kartu termasuk pengelolaan datanya tunduk pada ketentuan yang diberlakukan oleh Penerbit/Prinsipal Kartu (seperti PCIDSS).

##### API Card Registration

![](/img/docs/1_4_sequence_diagram_api_card_registration_(via_pjp_pias).png)

Sequence Diagram API Card Registration (via PJP PIAS)

![](/img/docs/1_5_sequence_diagram_api_card_registration_(direct_integration).png)

Sequence Diagram API Card Registration (Direct Integration)

**Informasi Umum**

Service Code01NameAPI Card RegistrationVersion1.0HTTP MethodPOSTPath…/{version}/registration-card-bind

##### API Card Registration - Set Limit

**Informasi Umum**

Service Code02NameAPI Card Registration – Set LimitVersion1.0HTTP MethodPOSTPath…/{version}/registration/card-bind-limit

##### API Card Registration Inquiry

![](/img/docs/1_6_sequence_diagram_api_card_registration_inquiry.png)

Sequence Diagram API Card Registration Inquiry

**Informasi Umum**

Service Code03NameAPI Card Registration InquiryVersion1.0HTTP MethodGETPath…/{version}/registration-card-inquiry

##### API Verify OTP (Direct Integration)

![](/img/docs/1_7_sequence_diagram_api_otp_validation_(direct_integration).png)

Sequence Diagram API OTP Validation (Direct Integration)

**Informasi Umum**

Service Code04NameAPI Verify OTP (Direct Integration)Version1.0HTTP MethodPOSTPath…/{version}/otp-verification

##### API Card Registration Unbinding

![](/img/docs/1_8_sequence_diagram_api_card_registration_unbinding.png)

Sequence Diagram API Card Registration Unbinding

**Informasi Umum**

Service Code05NameAPI Card Registration UnbindingVersion1.0HTTP MethodPOSTPath…/{version}/registration-card-unbind

##### API Account Creation

![](/img/docs/1_9_sequence_diagram_api_account_creation.png)

Sequence Diagram API Account Creation

**Informasi Umum**

Service Code06NameAPI Account CreationVersion1.0HTTP MethodPOSTPath.../{version}/registration-account-creation

##### API Account Binding

![](/img/docs/1_10_sequence_diagram_api_account_binding.png)

Sequence Diagram API Account Binding

**Informasi Umum**

Service Code07NameAPI Account BindingVersion1.0HTTP MethodPOSTPath.../{version}/registration-account-binding

##### API Account Binding Inquiry

![](/img/docs/1_11_sequence_diagram_api_account_binding_inquiry.png)

Sequence Diagram API Account binding inquiry

**Informasi Umum**

Service Code08NameAPI Account Binding InquiryVersion1.0HTTP MethodPOSTPath.../{version}/registration-account-inquiry

##### API Account Unbinding

![](/img/docs/1_12_sequence_diagram_api_account_unbinding.png)

Sequence Diagram API Account Unbinding

**Informasi Umum**

Service Code09NameAPI Account UnbindingVersion1.0HTTP MethodPOSTPath.../{version}/registration-account-unbinding

##### API Get OAuth URL

**Informasi Umum**

Service Code10NameAPI Get Oauth URLVersion1.0HTTP MethodGETPath../{version}/get-auth-code

##### API OTP

![](/img/docs/1_13_sequence_diagram_api_otp.png)

Sequence Diagram API OTP

**Informasi Umum**

Service Code81NameAPI OTPVersion1.0HTTP MethodPOSTPath.../{version}/otp

#### GUIDES

#### Spesifikasi Parameter Header dan Body API Registrasi

##### API Card Registration

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringO64Transaction identifier on partner systemaccountNameStringO128Customer’sfull name.cardDataEncrypted ObjectRefer to Standard Symmetric Encryption on Security Standard Document section 2.1.9bankAccountNoStringO32Account numberbankCardNoStringM19Card numberbankCardTypeStringO2Type of thecard (D –Debit, C –Credit, UE – Uang Electronik).dateOfBirthStringO8Customer’sdate of birth(YYYYMMDD).emailStringO254RFC 3696 Length 254 after “<” and “>”expiredDatetimeStringO25Expired date time ,Format expired date : (ISO 8601) YYYY-MM-DDThh:mm:ssexpiryDateStringO4Card expiry date.Format: MMYYidentificationNoStringO64Customer’sID numberidentificationTypeStringO2Type of ID(01 -passport, 02 –eKTP&KTP,03-TKTP, 04-SIM (DriverLicense), 99 –Others)custIdMerchantStringM18Merchant’s customer IDisBindAndPayStringO1Landing Page this value is always ‘N’merchantIdStringO64Merchant identifier that isunique per each merchantterminalIdStringO16identifier that is unique pereach merchantjourneyIdStringO32An identifier to uniquly identify a journey. On the first request of the journey, this must be equal to the X-EXTERNAL-ID.subMerchantIdStringO32Sub merchant IDexternalStoreIdStringO64Exkternal Store IDlimitdecimalO17,2Daily transaction limitmerchantLogoUrlStringO300Merchant’s logo URL for webviewphoneNoStringO16Customer’s phone numberFormat: 62xxxxxxxxxxxxxsendOtpFlagStringO3"YES" or "NO" to use OTP from PJSPtypeStringO20Action typeadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringO64Transaction identifier on service provider system.partnerReferenceNoStringO64Transaction identifier on service consumer systembankCardTokenStringM128Card tokenfor payment.chargeTokenStringO40string code for verification OTPrandomStringStringO16Random String to generate validation for webviewtokenExpiryTimeStringO25Time whenthe token will be expired. Time whenthe token willbe expired.Format: ISO-8601additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Card Registration - Set Limit

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringO64Transaction identifier on partner systembankAccountNoStringO32Account numberbankCardNoStringO19Card numberlimitdecimalO17,2Daily transaction limitbankCardTokenStringM128Card tokenfor payment.otpStringO8OTP Code / PasscodeadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringO64Transaction identifier on service provider system.partnerReferenceNoStringO64Transaction identifier on service consumer systemadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Card Registration Inquiry

**Request Body**

ParameterParameter TypeData TypeMandatoryLengthDescriptioncustIdMerchantPath paramStringM18Merchant’s customer ID

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionaccountListArray of Objects-accountDataObject-accountIdStringO16Account IDcreatedDateStringO26Creation datecredentialNoStringO16Credential numbercredentialTypeStringO2Credential typemaxLimitStringO6Maximum limitstatusStringO4statusadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Verify OTP (Direct Integration)

**Request Body**

ParameterData TypeMandatoryLengthDescriptionoriginalPartnerReferenceNoStringO64Transaction identifier on service consumer systemoriginalReferenceNoStringO64Transaction identifier on service provider systemactionStringO10actionmerchantIdStringO64Merchant IDotpStringO8OTP Code / PasscodechargeTokenStringO40OTP string code that is to be verified with the passcode obtained by the usertypeStringO20Action typeadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionoriginalReferenceNoStringO64Transaction identifier on service provider system that echo from request.originalPartnerReferenceNoStringO64Transaction identifier on service consumer systemaccountNoStringO11Customer’s account numberbankCardTokenStringO128Card tokenfor paymentcardPanStringO16Card numbercustomerIdStringO45Customer’s identificationemailStringO254RFC 3696 Length 254 after “<” and “>”expiredDatetimeStringO25Expired timeexpiryDateStringO4Card expiry date.Format: MMYYidentificationNoStringO64registered national id number on bank accountlinkageTokenStringO204Token used forPIN redirectionAPIphoneNoStringO16Customer’s phone number.Format: 62xxxxxxxxxxxxxqParamsURLStringO100Webview URL to set PINprocessqParamsObjectO-Params used toset PINidentificationactionStringO10Action type OTPsendOtpFlagStringO3Flag of using PJSP’s OTP or not. (“YES” or “NO”)subscribeDatetimeStringO25Subscription timetokenExpiryTimeStringO25Time whenthe token willbe expired.transactionTimestampStringO25Random String to generate validation for webviewadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Card Registration Unbinding

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringO64Transaction identifier on service consumer systemtokenStringM128This is an alphanumeric field which contains thepayment token used in a transaction. This field is used during setting token daily limit, purchase, and delete token.This token represent token number of card or token.bankCardNoStringO19Card number of the cardholdertypeStringO20Request type.Example: subscribe, unsubscribe, check OTPpartStringO64Merchant identifier that isunique per each merchantmerchantIdStringO64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDterminalIdStringO16identifier that is unique pereach merchanttokenRequestorIdStringO15An identifier to uniquely identify the token requestor.journeyIdStringO32An identifier to uniquly identify a journey. On the first request of the journey, this must be equal to the X-EXTERNAL-ID.transactionDateStringO25date oftransaction.Format: YYYY-MM-DDThh:mm:ssISO 8601additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringO64Transaction identifier on service provider system.partnerReferenceNoStringO64Transaction identifier on service consumer systemmessageStringO255Response messagecustomerIdStringO45Customer IDunsubscribeDateDatetimeOUnsubscribe date.Format: ISO 8601additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Account Creation

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringO64Transaction identifier on service consumer systemcountryCodeStringO2Requestor’s country codecustomerIdStringO45account ID of the customerdeviceInfoObjectO-osStringO40Device’s OSosVersionStringO40Device’s OS versionmodelStringO40Device’s modelmanufacturerStringO40Device’s manufactureremailStringO254RFC 3696 Length 254 after “<” and “>”langStringO8language support parameterlocaleStringO5Locale and language that customer selected in appnameStringO128User’s nameonboardingPartnerStringO8Onboarding partner of the customerphoneNoStringO16User’s phone number.Format: 62xxxxxxxxxxxxxredirectUrlStringO2048Merchant call back URLscopesStringO256The scopes of the authorizationseamlessDataStringO512the structure for the mobile and verification information，the value should be URLencoded.seamlessSignStringO512the signature data for the seamleassData, the value should be URLencodedstateStringO32statemerchantIdStringO64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDterminalTypeObjectO32Indicates the source terminal type and how the redirect will happenadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionpartnerReferenceNoStringO64Transaction identifier on service consumer systemauthCodeStringO64the auth code used to get accessToken and agreementTokenapiKeyNumericON/ACustomer IDaccountIdStringO32User’s account idstateStringO32stateadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Account Binding

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringO64Transaction identifier on service consumer systemactionStringO10Action type for OTPadditionalDataObjectO-userIdStringO20User ID of the card holderemailStringO254RFC 3696 Length 254 after “<” and “>”postalAddressStringO99Postal address of the userauthCodeStringO64An authorization code which the caller can used to obtain an access token.grantTypeStringO64The accessToken could be granted by authCode or refreshToken.isBindAndPayStringO1Landing Page this value is always ‘N’langStringO8Initiate languagelocaleStringO5Locale and language that customer selected in appmerchantIdStringM64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDmsisdnStringO15Phone number to be bindedotpStringO8OTP ValuephoneNoStringO16User’s phone number.Format: 62xxxxxxxxxxxxxplatformTypeStringO4Merchantplatform type(App, PC, ormobile web)redirectUrlStringO2048Redirect URLFor Agreementpage or PINpagereferenceIdStringO36Reference Idfrom previousGenerate OTPrefreshTokenStringO64Refresh token, which is used to refresh the access token.successParamsObjectO-accountIdstringO36account ID given to the userterminalIdStringO16Terminal identifiertokenRequestorIdStringO15Token Requestor IDadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionpartnerReferenceNoStringO64Transaction identifier on service consumer systemaccountTokenStringO128Account Token for PaymentaccessTokenInfoObjectO-accessTokenStringO32Access tokenexpiresInStringO25Datetime of token expiration.Format: ISO 8601refreshTokenStringO64Refresh token, which is used to refresh the access token.reExpiresInStringO25Datetime of refresh token expiration.Format: ISO 8601tokenStatusStringO25Status of tokenlinkIdStringO24Identifier provided at the time of linkingnextActionStringO255redirect the user to this url forauthenticationlinkageTokenStringO204Token used forPIN redirectionAPIparamsObjectO-Params used forset PINidentificationactionStringO10Action type OTPpinWebViewUrlStringO100URL for set PINprocessredirectToDeeplinkStringO255redirect the user to this deeplink in app for authentication. This is only valid in case of App Redirection FlowredirectUrlStringO2048Redirect URLforAgreementpage or PINpageuserInfoObjectO-publicUserIdStringO20User IDadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Account Binding Inquiry

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringO64Transaction identifier on service consumer systemadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringC64Transaction identifier on service provider system. Must be filled upon successful transactionpartnerReferenceNoStringO64Transaction identifier on service consumer systemaccountCurrencyStringO3Currency of registered AccountaccountNameStringO50Registered account nameaccountNoStringO13Registered account numberaccountTransactionLimitNumericO19,2Max debit amountendDatePeriodStringO10Binding end period.Format: YYYY-MM-DDstartDatePeriodStringO10Binding start period.Format: YYYY-MM-DDadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Account Unbinding

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringO64Transaction identifier on service consumer systemlinkIdStringO24Identifier provided at the time of linkingmerchantIdStringM64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDtokenIdStringO128Access TokenIDadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringO64Transaction identifier on service provider systempartnerReferenceNoStringO64Transaction identifier on service consumer systemmerchantIdStringO64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDlinkIdStringO24Identifier provided at the time of linkingunlinkResultStringO64Result of unlinking processadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Get OAuth URL

**Request Body**

ParameterData TypeMandatoryLengthDescriptionredirectUrlStringM256URL yang digunakan sebagai callback setelah proses getAuthCode berhasil.scopesListM256Scope akses dari authorization yang di inginkan.stateStringM32Random string untuk keperluan perlindungan terhadap CSRFmerchantIdStringO64Merchant identifier that isunique per each merchantsubMerchantIdStringO32Sub merchant IDlangStringO2Kode Bahasa yang layanan. ISO 639-1allowRegistrationBooleanOIf value equals true, provider may enable registration process during binding.Default true.seamlessDataStringO512Data yang diperuntukan untuk mempercepat proses otentikasimobileNumberStringO18Nomor telpon pengguna, apabila field ini terisi maka user wajin login dengan nomor yang sudah disertakanverifiedTimeStringO25Value yang menyatakan bahwa nomor ponsel yang sudah disertakan dalam seamless data sudah diverifikasi kepemilikannya dan tidak memerlukan verifikasi OTP oleh pihak penyedia. Validitas dari verifikasi ini hanya 10 menit.externalUidStringO32ID milik user pada aplikasi partnerdeviceIdStringO32Device ID milik UserseamlessSignStringC512Signature dari seamless data yang disertakan. Harus diisi jika seemles data tersedia.

Cara menyertakan seamless data:

`seamlessData = URLEncode({“mobileNumber”=”62822999999”})`

Cara menyertakan seamlessSign

`seamlessSign = URLEncode(sign(seamlessData)) `

- Partner menggukan private key untuk membuat seamless sign, sedangakan pihak penyedia akan menggukan publick key milik partner untuk proses verifikasi.
- Jika proses verifikasi gagal, maka seamless data akan diabaikan.
- Charset dari URL Encode menggunakan UTF-8

Konstruksi URL:

`../{version}/get-auth-code?state=<RANDOM_UNIQUE>&scopes=QUERY_BALANCE,PUBLIC_ID&redirectUrl=<MERCHANT_OAUTH_CALLBACK_URL>&seamlessData=<SEAMLESS_DATA>&seamlessSign=<SIGNATURE>`

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionauthCodeStringM256Authcode yang dapat ditukarkan dengan access token pada API account bindingstateStringM32Random string untuk keperluan perlindungan terhadap CSRF

##### API OTP

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerReferenceNoStringO64Transaction identifier on partner systemjourneyIdStringM32An identifier to uniquly identify a journey. On the first request of the journey, this must be equal to the X-EXTERNAL-ID.merchantIdStringO64Merchant IDsubMerchantStringO32Sub merchant IDexternalStoreIdStringO64External Store IDtrxDateTimeDateO25PJP internal system datetime with timezone, which follows the ISO-8601 standardbankCardTokenStringC128Card tokenfor payment.otpTrxCodeStringO2This will be used to retry OTP with service code of the original transaction request.otpReasonCodeStringO2otpReasonMessageStringO30additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response coderesponseMessageStringM150Response descriptionreferenceNoStringO64Transaction identifier on service provider system.partnerReferenceNoStringO64Transaction identifier on service consumer systemchargeTokenStringM40string code for verification OTPadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

#### CODE SNIPPETS

#### Code Snippets API Registrasi

##### API Card Registration

**Sample Request**

```http
POST …/v1.0/registration-card-bind HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-17T10:55:00+07:00
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
   "accountName":"John Doe",
   "cardData":"UIdFgZi9BhWx9Scbz/YK+JiMsPiKyhAFDaIRQe0wxMZYhu1rn73cQfgRGllEJjZlsoevALAzajD6qm2mG47kwsjN9tcArzlVAYS/jVflUn+zkiaMKuIxWQYJbB6MpET8Y0PjYB4aEkCvgxjP1wBcQMTVisCt5J7NCHetYAY64ucgxhBNCpCC6cG+3nCIuGqriq8/7EwhjpZ11YcDWecJ8glecZSv4HfjYZIFVXlwZD9rROd6xPgdKuVYVAAH8Y0pMH/x45FouvcuTuNIYOG26/btaUIRRnpkZsfzB4LPAk6CIQ/xia0rrWBwy479iXcV58q90u1ic1j0tuultFRCPmobf4N6AF5XXoERr3TOb7bJyjwodTpeQy+myzCDDidQmkKwNWOliQJdSjI+vHSi37ZfY1jlygnmaD1vQmblIj4= [SA(1] ",
   "custIdMerchant":"0012345679504",
   "isBindAndPay":"N",
   "merchantId":"00007100010926 ",
   "terminalId":"72001126",
   "journeyId":"20190329175623MTISTORE",
   "subMerchantId":"310928924949487",
   "externalStoreId":"000183004658",
   "limit":"1000000",
   "merchantLogoUrl":"https://bilba.test.com/dist/img/merchant-logo.png",
   "phoneNo":"08238748728423",
   "sendOtpFlag":"YES",
   "type":"subcribe",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-17T10:55:06+07:00

{
   "responseCode":"2000100",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "bankCardToken":"6d7963617264746f6b656e",
   "chargeToken":"abcd63617264746f6b656e",
   "randomString":"g4BoEz43jfjVvAvN",
   "tokenExpiryTime":"2020-12-17T11:00:00+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Card Registration - Set Limit

**Sample Request**

```http
POST …/v1.0/registration-card-inquiry HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-17T10:55:00+07:00
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
   "bankAccountNo":"93802938408123",
   "bankCardNo":"3984029384023984",
   "limit":"1000000",
   "bankCardToken":"6d7963617264746f6b656e",
   "otp":"12345678",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-17T10:55:06+07:00

{
   "responseCode":"2000200",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Card Registration Inquiry

**Sample Request**

```http
GET …/v1.0/registration-card-inquiry/custIdMerchant/8a95f0026d2860f301 HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-17T11:43:00+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.1617169
X-LONGITUDE: 106.6643946
CHANNEL-ID: 95221

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-17T11:43:03+07:00

{
   "responseCode":"2000300",
   "responseMessage":"Request has been processed successfully",
   "accountList":[
      {
         "accountData":{
            "accountId":"F8FP2WQWEATXFP8K",
            "createdDate":"2018-12-17T11:59:06+07:00",
            "credentialNo":"************0750",
            "credentialType":"DC",
            "maxLimit":"800000",
            "status":"ACT"
         }
      }
   ],
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Verify OTP (Direct Integration)

**Sample Request**

```http
POST …/v1.0/otp-verification HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-17T13:20:00+07:00
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
   "action":"otpLinkage",
   "merchantId":"00001",
   "otp":"12345678",
   "chargeToken":"TOK_TKNCPPPHUVL3IJVAXZI5GG4WBEC77YZ6::ADVQ",
   "type":"Subscribe",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone",
      "phoneNo":"081275647382"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-17T13:20:04+07:00

{
   "responseCode":"2000400",
   "responseMessage":"Request has been processed successfully",
   "originalReferenceNo":"2020102977770000000009",
   "originalPartnerReferenceNo":"2020102900000000000001",
   "accountNo":"12345678910",
   "bankCardToken":"6d7963617264746f6b656e",
   "cardPan":"2123123123125356",
   "customerId":"afhw6d7963617264746f6b656e963617264746f6b656e",
   "email":"john.doe@email.com",
   "expiredDatetime":"2019-02-24T14:12:25.871+07:00",
   "expiryDate":"1219",
   "identificationNo":"2020102020202000011001",
   "linkageToken":"xswe56",
   "phoneNo":"0899345678864332",
   "qParamsURL":"https://setPin",
   "qParams":{
      "action":"otpLinkage"
   },
   "sendOtpFlag":"YES",
   "subscribeDatetime":"2017-02-24T14:12:25.871+07:00",
   "tokenExpiryTime":"2017-02-24T14:12:25.871+07:00",
   "transactionTimestamp":"g4BoEz43jfjVvAvN",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Card Registration Unbinding

**Sample Request**

```http
POST …/v1.0/registration-card-unbind HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-17T13:50:00+07:00
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
   "token":"g4JeIz43jfjVvAvNxswe56g4JeIz43jfjVvAvNxswe56g4JeIz43jfjVvAvNxswe56g4JeIz43jfjVvAvNxswe56g4JeIz43jfjVvAvNxswe56g4Jg4JeIz43jfdsEga",
   "bankCardNo":"2123123123125356",
   "type":"Unsubscribe",
   "part":"00007100010926",
   "merchantId":"00007100010926",
   "subMerchantId":"23489182303312",
   "terminalId":"310928924949487",
   "tokenRequestorId":"7127425327776087324915228",
   "journeyId":"20190329175623MTISTORE",
   "transactionDate":"2020-12-17T13:50:00+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-17T13:50:04+07:00

{
   "responseCode":"2000500",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "customerId":"ae75e364134cdb2c7a4159106e38ca6b761983859dbv1",
   "unsubscribeDate":"2020-12-17T13:50:04+07:00",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Account Creation

**Sample Request**

```http
POST …/v1.0/registration-account-creation HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-17T14:49:00+07:00
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
   "countryCode":"ID",
   "customerId":"00-abcdefghijklmnopqrstuvwxyz0123456789-11",
   "deviceInfo":{
      "os":"IOS",
      "osVersion":"1",
      "model":"Iphone",
      "manufacture":"Apple"
   },
   "email":"john.doe@email.com",
   "lang":"ID",
   "locale":"en_ID",
   "name":"John doe",
   "onboardingPartner":"GOJEKXXX",
   "phoneNo":"0899345678864332",
   "redirectUrl":"https://merchant.site.com/apptoken",
   "scopes":"QUERY_BALANCE,QUERY_PROFILE",
   "seamlessData": "{\"mobile\":\"62-882345678\",\"verifiedTime\":\"2001-07-04T12:08:56+05:30\",\"externalUid\":\"TIXxxxxxUID\",\"reqTime\":\"2001-07-04T12:08:56+05:30\",\"riskData\":{\"fuzzyDeviceId\":\"k+OrCqw7QMNxlrT3qU2m0TRYTucd+nrMH2izjtltJgLqNI2XZgEAAA\",\"terminalType\":\"APP\",\"riskFlag\":\"00110\",\"realIp\":\"123.23.12.111\"}}",
   "seamlessSign" "URLEncode(BASE64(sigin( "{\"mobile\":\"62-882345678\",\"verifiedTime\":\"2001-07-04T12:08:56+05:30\",\"externalUid\":\"TIXxxxxxUID\",\"reqTime\":\"2001-07-04T12:08:56+05:30\",\"riskData\":{\"fuzzyDeviceId\":\"k+OrCqw7QMNxlrT3qU2m0TRYTucd+nrMH2izjtltJgLqNI2XZgEAAA::\",\"terminalType\":\"APP\",\"riskFlag\":\"00110\",\" realIp\":\"123.23.12.111\"}}" )))",
   "state":"12345556666",
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "terminalType":"SYSTEM",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-17T13:50:04+07:00

{
   "responseCode":"2000600",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "authCode":"g4JeIz43jfjVvAvNxswe56g4JeIz43jfjVvAvNxswe56g4JeIz43jfjVvAvNxswe56g4JeIz43jfjVvAvNxswe56g4JeIz43jfjVvAvNxswe56g4Jg4JeIz43jfdsEga",
   "apiKey":"AB12-CD34-EFGHIJ567890",
   "accountId":"ABCD1234-EF56-GH78-IJ90-KLMNOP123456",
   "state":"12345556666",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Account Binding

**Sample Request**

```http
POST …/v1.0/registration-account-binding HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-18T13:43:31+07:00
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
   "action":"otpLinkage",
   "additionalData":{
      "userId":"John Doe",
      "email":"john.doe@email.com",
      "postalAddress":"134346"
   },
   "authCode":"4b203fe6c11548bcabd8da5bb087a83b",
   "grantType":"AUTHORIZATION_CODE",
   "isBindAndPay":"N",
   "lang":"EN",
   "locale":"en_ID",
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "msisdn":"+62812345678901",
   "otp":"34564367",
   "phoneNo":"0899345678864332",
   "platformType":"app",
   "redirectUrl":"merchantapp://main_page",
   "referenceId":"08400000814-08400000814",
   "refreshToken":"201208134b203fe6c11548bcabd8da5bb087a83b ",
   "successParams":{
      "accountId":"ABCD1234-EF56-GH78-IJ90-KLMNOP123456"
   },
   "terminalId":"ID",
   "tokenRequestorId":"e-commerceA",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-18T13:43:37+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5

{
   "responseCode":"2000700",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "accountToken":"1Fhas7281han862=",
   "accessTokenInfo":{
      "accessToken":"ublicpBa869cad0990e4e17a57ecf7c5469a4b2",
      "expiresIn":"2021-07-04T12:08:56+05:30",
      "refreshToken":"201208134b203fe6c11548bcabd8da5bb087a83b",
      "reExpiresIn":"2051-07-04T12:08:56+05:30",
      "tokenStatus":"ACTIVE"
   },
   "linkId":"abcd1234efgh5678ijkl9012",
   "nextAction":"https://gopayapi.com/validate/otp/abcd123456789",
   "linkageToken":"xswe56",
   "params":{
      "action":"otpLinkage",
      "pinWebViewUrl":"https://setPin",
      "redirectToDeeplink":"https://gopayapi.com/redirect/gopay/abcd123456789"
   },
   "redirectUrl":" https://www.merchantapi.com/redirect/abcd123456789",
   "userInfo":{
      "publicUserId":"20180626111215830192DANAW3ID965200060630"
   },
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Account Binding Inquiry

**Sample Request**

```http
POST …/v1.0/registration-account-inquiry HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-18T14:39:21+07:00
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
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-18T14:39:30+07:00

{
   "responseCode":"2000800",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "accountCurrency":"IDR",
   "accountName":"Alen Miucic",
   "accountNo":"11231271284140",
   "accountTransactionLimit":"1000000",
   "endDatePeriod":"2022-05-21",
   "startDatePeriod":"2020-05-21",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Account Unbinding

**Sample Request**

```http
POST …/v1.0/registration-account-unbinding HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CupvIJbU1Mmu4a11MNDZ7Sg5u9a”
X-TIMESTAMP: 2020-12-18T14:48:11+07:00
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
   "linkId":"abcd1234efgh5678ijkl9012",
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "tokenId":"Aeox320xvijwefop10",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-18T14:48:30+07:00

{
   "responseCode":"2000900",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "merchantId":"00007100010926",
   "subMerchantId":"310928924949487",
   "linkId":"abcd1234efgh5678ijkl9012",
   "unlinkResult":"success",
   "additionalInfo":{
      "deviceId":"12345679237",
      "channel":"mobilephone"
   }
}

```

##### API Get OAuth URL

**Sample Request**

```http
GET …/v1.0/get-auth-code?state=WodkkwijSDs &scopes=QUERY_BALANCE,PUBLIC_ID&redirectUrl=https://domain.com/authSuccess.htm&seamlessData=%7B%22mobileNumber%22%3A%2262822999999999%22%7D&seamlessSign=gsfIUuC%2Bzs101rRFUhzz9753s9Dj4wg0EtwLwr8fMhZmCFybaCcwvAXGZ0RDxqzb9fJuFre%2Bmsi9JcwHICVx%2FB1onruQNldI4Y%2BUZqVQLgUVz1ynAa1qyyaTKliXOfy3t%2FbOhXd0QfZ3e1zbQT5
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CupvIJbU1Mmu4a11MNDZ7Sg5u9a”
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CupvIJbU1Mmu4a11MNDZ7Sg5u9a”
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

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-23T09:10:18+07:00

{
   "responseCode":"2001000",
   "responseMessage":"Request has been processed successfully",
   "authCode":"a4sd5a4fsaf5d5f4df66ad85f4",
   "state":"WodkkwijSDs"
}

```

##### API OTP

**Sample Request**

```http
POST …/v1.0/otp HTTP/1.2
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
   "partnerReferenceNo":"2020102900000000000001",
   "journeyId":"20190329175623MTISTORE",
   "merchantId":"00001",
   "subMerchant":"310928924949487",
   "externalStoreId":"124928924949487",
   "trxDateTime":"2020-12-21T14:56:11+07:00",
   "bankCardToken":"6d7963617264746f6b656e",
   "otpTrxCode":"54",
   "otpReasonCode":"",
   "otpReasonMessage":"",
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
   "responseCode":"2008100",
   "responseMessage":"Request has been processed successfully",
   "referenceNo":"2020102977770000000009",
   "partnerReferenceNo":"2020102900000000000001",
   "chargeToken":"abcd63617264746f6b656e",
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