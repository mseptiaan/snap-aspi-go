Developer APIDokumentasi

![](/content/reskindeveloperapibcacoid/images/icons/arrow-left-small.png)Kunjungi bca.co.id

[![Developer BCA](/-/media/Project/DeveloperApiBcaCoId/Images/logobca2x.png?h=56&w=344&hash=1BE872950DC5DBD5F50638586EBA3326)](/id)

![](/content/reskindeveloperapibcacoid/images/icons/menu.png)

- [Home](/id)
- [Fitur API](/id/Fitur-API)
- [Developer Center](/id/Developer)

- [Login](/id/Login)
- [Daftar API](https://webform.bca.co.id/api/registrasi-api)
- ![](/content/reskindeveloperapibcacoid/images/icons/menu.png)



[![](/-/media/Project/DeveloperApiBcaCoId/Images/Icon/Menu/cara-penggunaan.png?h=20&w=20&hash=A3032986CE92B49D2A318DF6E40D0B8E)\
Cara Penggunaan](/id/Cara-Penggunaan)



[![](/-/media/Project/DeveloperApiBcaCoId/Images/Icon/Menu/fitur-api.png?h=20&w=20&hash=B36061478C7709900401C76326DCBD8F)\
Fitur API](/id/Fitur-API)



[![](/-/media/Project/DeveloperApiBcaCoId/Images/Icon/Menu/dokumentasi.png?h=20&w=20&hash=FF6C3470F46208BEAC3A0C8673C72354)\
Dokumentasi](/id/Dokumentasi)



[![](/-/media/Project/DeveloperApiBcaCoId/Images/Icon/Menu/sandbox.png?h=20&w=20&hash=D38067FE2B8E757ED7BE2B5B2DF24F7B)\
Sandbox](/id/Sandbox)



[![](/-/media/Project/DeveloperApiBcaCoId/Images/Icon/Menu/faq.png?h=20&w=20&hash=246B99551002C5D861B44816BC7AFDD2)\
FAQ](/id/FAQ)

- [ID](/id/dokumentasi/) [EN](/en/dokumentasi/)

![](/content/reskindeveloperapibcacoid/images/icons/clear.png)
- [Home](/id)

- [Fitur API](/id/Fitur-API)

- [Developer Center](/id/Developer)

- [![](/-/media/Project/DeveloperApiBcaCoId/Images/Icon/Menu/cara-penggunaan.png?h=20&w=20&hash=A3032986CE92B49D2A318DF6E40D0B8E)\
  Cara Penggunaan](/id/Cara-Penggunaan)

- [![](/-/media/Project/DeveloperApiBcaCoId/Images/Icon/Menu/fitur-api.png?h=20&w=20&hash=B36061478C7709900401C76326DCBD8F)\
  Fitur API](/id/Fitur-API)

- [![](/-/media/Project/DeveloperApiBcaCoId/Images/Icon/Menu/dokumentasi.png?h=20&w=20&hash=FF6C3470F46208BEAC3A0C8673C72354)\
  Dokumentasi](/id/Dokumentasi)

- [![](/-/media/Project/DeveloperApiBcaCoId/Images/Icon/Menu/sandbox.png?h=20&w=20&hash=D38067FE2B8E757ED7BE2B5B2DF24F7B)\
  Sandbox](/id/Sandbox)

- [![](/-/media/Project/DeveloperApiBcaCoId/Images/Icon/Menu/faq.png?h=20&w=20&hash=246B99551002C5D861B44816BC7AFDD2)\
  FAQ](/id/FAQ)

- [ID](/id/dokumentasi/) [EN](/en/dokumentasi/)

- [Masuk](/id/Login)
- [Pengajuan Kerja Sama](https://webform.bca.co.id/api/registrasi-api)

![](/content/ReskinDeveloperApiBcaCoId/images/icons/clear.png)

[Introduction](#introduction)

AUTHENTICATION

[OAuth2.0](#oauth20)

[Signature](#signature)

[Headers](#headers)

AUTHENTICATION SNAP

[OAuth2.0 (SNAP)](#oauth20snap)

[Headers (SNAP)](#headerssnap)

[Signature (SNAP)](#signaturesnap)

API Services

[Business Debit Card](#businessdebitcard)

[Informasi Rekening](#informasirekening)

[Pembukaan Rekening Online](#pembukaanrekeningonline)

[Virtual Account untuk Biller](#virtualaccountuntukbiller)

[Transfer ke Rekening Virtual Account BCA](#transferkerekeningvirtualaccountbca)

[ECR EDC](#ecredc)

[QRIS MPM](#qrismpm)

[QRIS CPM](#qriscpm)

[Shared Biller BCA](#sharedbillerbca)

[Transfer Dana](#transferdana)

[Collection](#collection)

[OneKlik](#oneklik)

[Valuta Asing](#valutaasing)

[Informasi Umum](#informasiumum)

[Financing](#financing)

[FIRE](#fire)

[Top Up Flazz](#topupflazz)

Others

[Download List Bank Code](/-/media/Project/DeveloperApiBcaCoId/Files/beneficiary_bank_list_for_API_BCA.pdf)

Give us some feedback

[Panduan Sandbox API](/-/media/Project/DeveloperApiBcaCoId/Files/sandboxTutorial.pdf)

[Introduction](#introduction)

AUTHENTICATION

[OAuth2.0](#oauth20) [Signature](#signature) [Headers](#headers)

AUTHENTICATION SNAP

[OAuth2.0 (SNAP)](#oauth20snap) [Headers (SNAP)](#headerssnap) [Signature (SNAP)](#signaturesnap)

API Services

[Business Debit Card](#businessdebitcard) [Informasi Rekening](#informasirekening) [Pembukaan Rekening Online](#pembukaanrekeningonline) [Virtual Account untuk Biller](#virtualaccountuntukbiller) [Transfer ke Rekening Virtual Account BCA](#transferkerekeningvirtualaccountbca) [ECR EDC](#ecredc) [QRIS MPM](#qrismpm) [QRIS CPM](#qriscpm) [Shared Biller BCA](#sharedbillerbca) [Transfer Dana](#transferdana) [Collection](#collection) [OneKlik](#oneklik) [Valuta Asing](#valutaasing) [Informasi Umum](#informasiumum) [Financing](#financing) [FIRE](#fire) [Top Up Flazz](#topupflazz)

Others

[Download List Bank Code](/-/media/Project/DeveloperApiBcaCoId/Files/beneficiary_bank_list_for_API_BCA.pdf)Give us some feedback [Panduan Sandbox API](/-/media/Project/DeveloperApiBcaCoId/Files/sandboxTutorial.pdf)

## Search Result

#### Introduction

API Updated: March 17th, 2025

- New API: QRIS CPM

Try our APIs using Sandbox. We provide sandbox with dummy and static datas. All the parameter value that can be used to try our sandbox are written on the blue box in this Documentation.

###### AUTHENTICATION

#### OAuth2.0

BCA APIs is using OAuth 2.0 as the authorization framework. To get the access token, you need to be authorized by `client_id` and `client_secret`. To learn more about the OAuth 2.0 authorization framework, you can read the [RFC6749 Documentation](https://tools.ietf.org/html/rfc6749).

client\_secret dan client\_id are used for authentication using OAuth 2.0. You can generate client\_id and client\_secret after sign in and create application.

Do not share your client\_secret!  This token act like password, keep it secret and secure, should anyone obtain this information, immediately reset or revoke your **client\_secret**.

**Access Token**

access\_token is an opaque string token that identify the user of the API. This token is required each time an application call API. There are several way to obtain an access\_token, which will be described bellow.

Access token must be stored in a secure storage!

Since access\_token is portable which mean that once its obtained any request with valid credentials will be considered valid, any agent (mobile device, web browser, or server) could call API requests.

**Obtaining Access Token**

Access token can be obtained in many way, depend on the grant\_type of the application. To access all the services in this sandbox, you will need the access token with grant\_type = client\_credentials.

**Client Credentials Grant**

client\_credentials grant will provide application access to API without requiring any user credential. Any call requested using access\_token obtained using this method are made on behalf of the application instead of the user.

This grant type is designed to be used by server to server call. In order to obtain access\_token a request must be made with following specification

##### Request

SettingValueHTTP MethodPOSTPath/api/oauth/tokenHostsandbox.bca.co.id

##### Request Headers

NameFormat MandatoryDescription AuthorizationBasic `base64(client_id:client_secret)`YesContent-Typeapplication/x-www-form-urlencodedYes

##### Request

FieldData Type Mandatory Descriptiongrant\_type StringYes value = `client_credentials`

Result of the request will contains following information:

##### Response

FieldData TypeDescriptionaccess\_tokenStringyour access\_tokentoken\_typeStringdefault is `Bearer`expires\_inString`access_token` validity, in seconds scopeStringapplication scope/permission granted to application

#### Signature

Signature is used by BCA to verify that your request is not altered by attackers.

The outline of the HMAC validation process is as follows:

1. Retrieve Timestamp from HTTP Header (X-BCA-Timestamp)
2. Retrieve the API Key form HTTP Header (X-BCA-Key)
3. Lookup the API Secret corresponding to the received key in internal store
4. Retrieve client HMAC from HTTP Header lowercase hexadecimal format (X-BCA-Signature)
5. Calculate HMAC using the API Secret as the HMAC secret key
6. Compare client HMAC with calculated HMAC

If HMAC hash comparison is invalid API Gateway will return a HTTP 400 error code with "HMAC Mismatch" message.

API Key and API Secret are used for hashing HMAC. You can generate API Key and API Secret after sign in and create application.

**Generate Signature**

SHA-256 HMAC is used to generate the signature with your API secret as the key.

Signature = HMAC-SHA256(apiSecret, StringToSign)

The StringToSign will be a colon-separated list derived from some request data as below:

StringToSign = HTTPMethod+":"+RelativeUrl+":"+AccessToken+":"+Lowercase(HexEncode(SHA-256(RequestBody)))+":"+Timestamp

For GET request (with no RequestBody), you still need to calculate SHA-256 of an empty string.

Details about the data used to derived The StringToSign is explained in the next sections.

**HTTP Method**

- HTTP Method is HTTP Method such as GET, POST, PUT, PATCH, DELETE.
- HTTP Method must be given in upper case.

**Relative URL**

- Relative URL is the URL after the hostname & port number.
- Relative URL alse includes the query string and must begin with a slash character. Example

**Timestamp**

The timestamp must be presented in ISO8601 format (yyyy-MM-ddTHH:mm:ss.SSSTZD)

FormatDescriptionyyyyfour-digit yearMMtwo-digit month (01=January, etc.)ddtwo-digit day of month (01 through 31) Tliteral ’T’ as date and time separator HHtwo digits of hour (00 through 23) (am/pm NOT allowed) mmtwo digits of minute (00 through 59) stwo digits of second (00 through 59) SSSthree digits representing millisecond (000 through 999) TZDtime zone designator (+hh:mm or -hh:mm)

EXAMPLE

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST EXAMPLE

Data :
  - Method : POST
  - Relative URL : /banking/corporates/transfers
  - Access token : lIWOt2p29grUo59bedBUrBY3pnzqQX544LzYPohcGHOuwn8AUEdUKS
  - Request body :
    {
    "CorporateID" : "BCAAPI2016",
        "SourceAccountNumber" : "0201245680",
        "TransactionID" : "00000001",
        "TransactionDate" : "2016-01-30",
        "ReferenceID" : "12345/PO/2016",
        "CurrencyCode" : "IDR",
        "Amount" : "100000.00",
        "BeneficiaryAccountNumber" : "0201245681",
        "Remark1" : "Transfer Test",
        "Remark2" : "Online Transfer"
    }
  - Timestamp : 2016-02-03T10:00:00.000+07:00

API Secret : 22a2d25e-765d-41e1-8d29-da68dcb5698b

Lowercase(HexEncode(SHA-256(RequestBody))) : e3cf5797ac4ac02f7dad89ed2c5f5615c9884b2d802a504e4aebb76f45b8bdfb

StringToSign : POST:/banking/corporates/transfers:lIWOt2p29grUo59bedBUrBY3pnzqQX544LzYPohcGHOuwn8AUEdUKS:e3cf5797ac4ac02f7dad89ed2c5f5615c9884b2d802a504e4aebb76f45b8bdfb:2016-02-03T10:00:00.000+07:00

Signature : 69ad66589ade078a30922a0848725cf153aecfcca82eba94e3270285b4a9c604

GET Example:

Data:
  - Method : GET
  - Relative URL : /banking/v3/corporates/BCAAPI2016/accounts/0201245680/statements?StartDate=2016-09-01&EndDate=2016-09-01
  - Access token : lIWOt2p29grUo59bedBUrBY3pnzqQX544LzYPohcGHOuwn8AUEdUKS
  - Request body :
  - Timestamp : 2016-02-03T10:00:00.000+07:00

API Secret : 22a2d25e-765d-41e1-8d29-da68dcb5698b

Lowercase(HexEncode(SHA-256(RequestBody))) : e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855

StringToSign : GET:/banking/v2/corporates/BCAAPI2016/accounts/0201245680/statements?EndDate=2016-09-01&StartDate=2016-09-
01:lIWOt2p29grUo59bedBUrBY3pnzqQX544LzYPohcGHOuwn8AUEdUKS:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855:2016-02-03T10:00:00.000+07:00

Signature : 3ac124303746d222387d4398dddf33201a384aa22137aa08f4d9843c6f467a48

```

EXAMPLE

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs

```

#### Headers

To successfully communicate with BCA Banking API, you **must provide** the following headers **in every API request**:

NameTypeDescriptionAuthorizationAlphanumericOAuth2.0 Token Format value: Bearer {access\_token}Content-TypeAlphanumericContent of your request body e.g. application/json Origin urlOrigin domain e.g. yourdomain.comX-BCA-KeyAlphanumericYour API Key generated by BCAX-BCA-Timestampyyyy-MM-ddTHH:mm:ss.SSSTZD (ISO 8601)Timestamp generated by merchant in ISO 8601 e.g. “2016-02-03T10:00:00.000+07:00”X-BCA-SignatureAlphanumericSignature, please refer to section above

Sample request with complete headers:

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
GET /general/rate/forex HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer a0daAPsGwPDHGMv6MpWzoIPgZvN9YvrSi7xdVI7Jb98638ilM7ila7
Content-Type: application/json
Origin: yourdomain.com
X-BCA-Key: 41138489-1057-4e7e-ab93-9bc97b511cf6
X-BCA-Timestamp:  2016-02-03T10:00:00.000+07:00
X-BCA-Signature: ce20470484c27c0441d7c9dedc7d3e187e5e908df1780c6c389941b779033519
```

Sample request with complete headers:

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs

```

###### AUTHENTICATION SNAP

#### OAuth2.0 (SNAP)

##### 1\. OAuth2.0 (SNAP)

The BCA Corporate Banking API is using OAuth 2 as the authorization framework. To access all the services you’ll need the access token with grantType=client\_credentials. To get the access token, you need to be authorized by client\_id. To learn more about the OAuth 2 authorization framework you can read the rfc6749 documentation (https://tools.ietf.org/html/rfc6749).

POST /openapi/v1.0/access-token/b2b HTTP/1.1

Host: server.example.com

X-TIMESTAMP : DateTime with timezone, which follows the ISO-8601 standard

X-CLIENT-KEY : client\_id

X-SIGNATURE : Signature Asymmetric

Content-Type : application/json

{

           "grantType": "client\_credentials"

}

##### Errors

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message (Indonesian)4004007300Invalid field format \[clientId/clientSecret/grantType\]4004007300Unauthorized. \[Connection not allowed\]4004007301Invalid field format \[X-TIMESTAMP\]4004007302Invalid mandatory field \[X-CLIENT-KEY\]4014017300Unauthorized. \[Signature\]4014017300Unauthorized. \[Unknown client\]5045047300Timeout

OAuth2.0 (SNAP)

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
curl https://sandbox.bca.co.id/openapi/v1.0/access-token/b2b \
-H "X-TIMESTAMP: 2022-04-21T17:34:52+07:00" \
-H "X-CLIENT-KEY: 123 \
-H "X-SIGNATURE:
cjCDSRIf8FIXGs28uVFhxBuRhgV/xEnMOxargNqkTwvgMqDMdyEfwqQ2PIFz891K0LE81dgN
b2CEsJrGAG7tRTur68A0YEHpx8iXXWHD40+iSCgBa8o3QsVBB1ryP1d8FV+dMEeixUKw08fh
QCsN1vqIJhPyZFkg2r3Sr8h+dzG4t7ldr7fTkO7wvEFT4LG9kTAEDqPNPW1P8bZqdHZUe6VD
zqlUWO12Zu2wHofuzRjvXXVYDm1/lwhDmdJqVrvi9g8L5/PmLnzGkjcniJlpBGYYkGE0eA25
bOdJEWdaklQjQDrFg+G8Pl4EnY3PXW1/u7ZwGU131mVy1kvUw801Yw== \
-H "Content-Type: application/json \
 -d "{
 "grantType": "client_credentials"
 }"
```

OAuth2.0 (SNAP)

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
 "responseCode": "2007300",
 "responseMessage": "Successful",
"accessToken": "7t4tbXnlyn4NABRn0FAhB69CRhxghlPPfPK2l9quE29l4D4k5DLH57,
"tokenType": "bearer",
"expiresIn": "900"
}
```

#### Headers (SNAP)

##### 1\. Headers

To successfully communicate with BCA Banking Open API, you must provide the following headers in every API request :

##### Request

NameType**Length****Mandatory**DescriptionAuthorizationANN/AYRepresents access\_token of a request, String starts with keyword “Bearer” followed by accessToken Content-TypeAN16 (Fixed Length)YContent of you request body e.g. application/json

X-TIMESTAMP

yyyy-MMddThh:mi:ssTZD (ISO 8601)

25 (Fixed Length)YClient’s current local time in yyyy-MM-ddTHH:mm:ssTZD format X-SIGNATURE

ANN/AYSignature Symmetric, please refer to Signature section ORIGIN

N/ANOrigin Domain www.yourdomain.com X-EXTERNAL-ID

36 (Max Length)YNumeric String Reference number that should be unique in the same day

#### Signature (SNAP)

#### Signature Asymmetric

Signature asymmetric is used by BCA to verify that your access token request is not altered by attackers.

The outline of the HMAC validation process is as follows:

1. Retrieve Timestamp from HTTP Header (X-TIMESTAMP)
2. Retrieve the Client Key form HTTP Header (X-CLIENT-KEY)
3. Retrieve client HMAC from HTTP Header lowercase base64 format (X-SIGNATURE)
4. Calculate HMAC using the Private Key as the HMAC secret key
5. Compare client HMAC with calculated HMAC

If HMAC hash comparison is invalid API Gateway will return a HTTP 401 error code together with the following error message on JSON format:

{

        "responseCode" : "4017300",

        "responseMessage" : "Unathorized. \[Signature\]"

}

If the HMAC calculation is successful and the calculated value matches the value received from the client, the signature is considered valid.

##### Generate Signature Asymmetric

SHA256withRSA is used to generate the signature with your Private Key as the key

X-SIGNATURE = SHA256withRSA(PrivateKey, StringToSign)

Note = X-SIGNATURE should be encoded by Base64

The StringToSign will be a colon-separated list derived from some request data as below :

StringToSign = client\_ID+"\|"+X-TIMESTAMP

Details about the data used to derived The StringToSign is explained in the next sections.

Note : **Partner need to send their public key in x.509 format for BCA to use when verifying signature**

Sample public key in x.509 format :

-----BEGIN PUBLIC KEY-----

MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAomV+Vm1xlRXanmh108Kusls7SSKec0oCejtc9QG Obpd4RnQ+7gihm2k6etnSNP7b+XrpY+fBkiQNaBInii9M10kW9Bhf/M9GH/edL3IqnzDNSi7tcoQgnO7h8x mzLNWHTjtR6bkrsdBS5dry6htotaF5KXomuoYgztCdGDOa0W20aeLzYSXIoW7s/Ay5yIXt0xaXTll3/bmez leguFPnwQZq5EqZFWlUZvutDi+f2l9rTRY0Fb64y+VAf+mnIbEovGqsPEeF/p97YWxcY7CWm8NsT0lwBVOt kmEl967Brz5yvEObF5bJgVodi6mNVsN1ki0MCitIhYO8shcE7eUilQIDAQAB

 -----END PUBLIC KEY-----

#### Signature Symmetric

Signature is used by BCA to verify that your open api service request is not altered by attackers.

The outline of the HMAC validation process is as follows:

1. Retrieve Timestamp from HTTP Header (X-TIMESTAMP)
2. Lookup the Client Secret corresponding to the received key in internal store
3. Retrieve client HMAC from HTTP Header lowercase base64 format (X-SIGNATURE)
4. Calculate HMAC using the Client Secret as the HMAC secret key
5. Compare client HMAC with calculated HMAC

If HMAC hash comparison is invalid API Gateway will return a HTTP 401 error code together with the following error message on JSON format:

{

        "responseCode" : "401xx00",

        "responseMessage" : " Unauthorized. \[Signature\]",

        "data" : {}

}

**Note : xx -> customize to each service code**

If the HMAC calculation is successful and the calculated value matches the value received from the client, the signature is considered valid.

##### Generate Signature Symmetric

SHA-512 HMAC is used to generate the signature with your Client Secret as the key.

X-SIGNATURE = HMAC-SHA512(ClientSecret, StringToSign)

**Note = X-SIGNATURE should be encoded by Base64**

The StringToSign will be a colon-separated list derived from some request data as below :

StringToSign = HTTPMethod+":"+RelativeUrl+":"+AccessToken+":"+ Lowercase(HexEncode(SHA-256(MinifyJson(RequestBody))))+":"+Timestamp

HexEncode are optional to use, use it if the SHA-256 returns a binary stream.

Details about the data used to derived The StringToSign is explained in the next sections.

##### HTTP Method

- HTTP Method is HTTP Method such as GET, POST, PUT, PATCH, DELETE.
- HTTP Method must be given in upper case.

##### Relative URL

- Relative URL is the URL after the hostname & port number.
- Relative URL also includes the query string and must begin with a slash character. Example :
Full URLRelative URLhttps://example.com/api/v2/sample?param1=value/api/v2/sample?param1=value1&pa ram2=value2 https://example.com or https://example.com//- The Relative URL must be URI-encoded according to the following rules:
1. Do not URI-encode forward slash ( / ) if it was used as path component.
2. Do not URI-encode question mark ( ? ), equals sign ( = ), and ampersand ( & ) if they were used as query string component: as separator between the path and query string, between query parameter and its value, and between each query parameter and value pairs.
3. Do not URI-encode these characters: A-Z, a-z, 0-9, hyphen ( - ), underscore ( \_ ), period ( . ), and tilde ( ~ ) which are defined as unreserved characters in RFC 3986. As for RFC 3986, means that comma that appear in the value of query params or path params should be encoded too when generating Signature.
4. Percent-encode all other characters not meeting the above conditions using the format: %XY, where X and Y are hexadecimal characters (0-9 and uppercase A-F). For example, the space character must be encoded as %20 (not using '+', as some encoding schemes do) and extended UTF-8 characters must be in the form %XY%ZA%BC.

- The query string parameters must be re-ordered according to the following rules:


1. Sorted by parameter name lexicographically
2. If there are two or more parameters with the same name, sort them by parameter values.
Relative URLSorted Relative URL/api/v2/sample?A-param=value1&Zparam=value2&B-param=value3 /api/v2/sample?A-param=value1&Bparam=value3&Z-param=value2

##### AccessToken

- AccessToken is an OAuth 2 access token retrieved from the HTTP “Authorization” header.
- AccessToken has a validity time of 900 seconds

##### RequestBody

- RequestBody need to be hashed with SHA-256.
- Before hashed with SHA-256, RequestBody must to convert to MinifyJSON (remove whitespace except for the key or value json)
- If the RequestBody is empty, set it to empty string. .

##### Timestamp

The timestamp must be presented in ISO8601 format (YYYY-MM-DDThh:mm:ssTZD)

YYYY = four-digit year

MM = two-digit month (01=January, etc.)

DD = two-digit day of month (01 through 31)

T = literal 'T' as date and time separator

hh = two digits of hour (00 through 23) (am/pm NOT allowed)

mm = two digits of minute (00 through 59)

ss = two digits of second (00 through 59)

TZD = time zone designator (Z or +hh:mm or –hh:mm)

###### API Services

#### Business Debit Card

##### 1\. Change Business Debit Card Status API

This service is used to update status BDC card.

This feature is not yet available to be accessed via Sandbox

**Additional Headers and URI Params:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)FixedY Channel's Identifier using WSID = 95411X-PARTNER-IDHeaderString(36)MaxY Partner ID (company code)X-EXTERNAL-IDHeaderString(36)MaxYNumeric String. Reference number that should be unique in the same day

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptioncardNumberString (16)FixedY NumericFilled with card number of BDCblockTypeString (1)FixedYAlphanumericBlock Type for the BDC

O: Open half block

H: Half Block

B: Full Block

Fill with capital lettersreasonCodeString (1)FixedYAlphanumericReason of doing card status change

H: Hilang (Lost)

R: Tertelan (Swallowed)

O: Others

Fill with capital letters

Result of the request will contains following information:

**Response:**

**Field****Data Type****Length Type****Mandatory****Format****Description**responseCodeString (7)FixedYAlphanumericResponse code to identify transaction status (success or failed)

Format: AAABBCC

AAA: HTTP Code

BB: Service Code

CC: Case Code

Service Code API BDC Update Status: US

Example: 200US00 – SuccessfulresponseMessageString (150)MaxYAlphanumericResponse description

Example: SuccessfulresponseDataObjectY   cardNumberString (16)FixedYNumericCard Number for the BDC

If error will be filled with '-'   cardStatusString (1)FixedYAlphanumericStatus of BDC

If error will be filled with '-'

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message (Indonesian)**Error Message (English)****200**200US00SuccessfulTransaction successful.**400**400US00Bad RequestGeneral request failed error**400**400US01Invalid Field Format {Field}Field does not match the format.**401**401US00Unauthorized. \[Reason\]Invalid Signature/Unknown Client/Connection Not Allowed.**401**401US01Invalid token (B2B)Access Token Not Exist/Access Token Expired/Token is Invalid.**403**403US01Feature Not Allowed The feature is not allowed to use. Please make sure you registered this feature in your partnership data.**403**403US07Card BlockedCard Status = H (Half Blocked / B (Blocked) / C (Closed)**403**403US15Transaction Not Permitted. \[Reason\]Card is already closed or active.**404**404US11Invalid CardCard number is not found.**409**409US00ConflictX-EXTERNAL-ID duplicate.**500**500US01Internal Server ErrorUnknown internal server failure, please retry the process again.**504**504IC00TimeoutYour transaction timeout.

curl –X POST https://sandbox.bca.co.id/openapi/bdc/v1.0/update-status

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'Content-type:application/json'
-H 'Authorization:Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-TIMESTAMP:2023-12-25T13:59:21+08:00'
-H 'X-SIGNATURE:85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'ORIGIN:www.hostname.com'
-H 'Content-type:application/json'
-H 'CHANNEL-ID:95411'
-H 'X-PARTNER-ID:012345'
-H 'X-EXTERNAL-ID:01234506022024001'
-d '
{
	"cardNumber": "1234567890123456",
	"blockType": "B",
	"reasonCode": "H"
}'
```

curl –X POST https://sandbox.bca.co.id/openapi/bdc/v1.0/update-status

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
Response Sample
-H 'Content-type:application/json'
-H 'X-TIMESTAMP:2020-12-18T14:48:30+07:00'
{
	"responseCode": "200US00",
	"responseMessage": "Successful",
	"responseData":
	{
		"cardNumber": "1234567890123456",
		"cardStatus": "B"
	}
}
```

##### 2\. Data Business Debit Card Inquiry API

This service is used to inquiry data BDC card.

This feature is not yet available to be accessed via Sandbox

**Additional Headers and URI Params:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)FixedY Channel's Identifier using WSID = 95411X-PARTNER-IDHeaderString(36)MaxY Partner ID (company code)X-EXTERNAL-IDHeaderString(36)MaxYNumeric String. Reference number that should be unique in the same day

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionbankCardTokenString(16)FixedY NumericFilled with card number of BDC

Result of the request will contains following information:

**Response:**

**Field****Data Type****Length Type****Mandatory****Format****Description**responseCodeString (7)FixedYAlphanumericResponse code to identify transaction status (success or failed)

Format: AAABBCC

AAA: HTTP Code

BB: Service Code

CC: Case Code

Service Code API BDC Setting Limit: SL

Example: 200SL00 – SuccessfulresponseMessageString (150)MaxYAlphanumericResponse description

Example: SuccessfulresponseDataObjectY   accountNoString (32)FixedNNumericAccount Number for the BDC

If error will be filled with '-'   nameString (140)MaxNAlphanumericName of the card holder

If error will be filled with '-'   accountInfosArrayN      balanceTypeString (70)MaxNAlphanumericThe balance type of BDC

If error will be filled with '-'      amountObjectN         valueString (16,2)MaxYISO4217The amount value of limit

If error will be filled with '-'         currencyString (3)FixedYAlphanumericThe currency used for the limit If error will be filled with '-'   additionalInfoObjectNAdditional Information      cardNumberString (16)FixedNNumericCard Number of BDC

If error will be filled with '-'       factTypeString (2)FixedNAlphanumericType of account (Primary/Secondary)

PF: Primary

SE: Secondary

If error will be filled with '-'      totalAccountString (2)FixedNNumericTotal number of Account

01 - 05

BDC only 01

If error will be filled with '-'      lanIdString (1)FixedNNumericLanguage Code

01: English

02: Indonesia

If error will be filled with '-'      openDateString (8)FixedNNumericCard Open Date

If error will be filled with '-'      prodDateString (8)FixedNNumericCard Production Date

If error will be filled with '-'      closeDateString (8)FixedNNumericCard Close Date

If error will be filled with '-'      expiryDateString (4)FixedNNumericCard Expiry Date

If error will be filled with '-'      embNameString (28)MaxNAlphanumericCompany Name

If error will be filled with '-'      addr1String (30)MaxNAlphanumericAddress

If error will be filled with '-'      addr2String (30)MaxNAlphanumericAddress

If error will be filled with '-'      addr3String (30)MaxNAlphanumericAddress

If error will be filled with '-'      birthDateString (8)FixedNNumericBirth Date of Card Holder

If error will be filled with '-'      branchCodeString (4)FixedNNumericCoordinator Branch Code

If error will be filled with '-'      empCodeString (1)FixedNAlphanumericCode to specify.

N: Normal

F: Free

If error will be filled with '-'      chargeFlagString (1)FixedNAlphanumericFlag to determine if the card have been billed.

U: Unpaid

P: Paid

If error will be filled with '-'      typeString (1)FixedNNumericFilled with the type of debit card.

0 : MAGNETIC STRIPE

1 : NSICCS/GPN FISIK

2 : MASTERCARD FISIK

3 : GPN DIGITAL

4 : MASTERCARD DIGITAL

5 : GPN CONTACTLESS

6 : MASTERCARD CONTACTLESS

If error will be filled with '-'      deptCodeString (5)FixedNNumericCard co-brand code

BDC = 0 (Conventional Card)

If error will be filled with '-'      stockBranchCodeString (4)FixedNNumericFilled with the card stock branch code

If error will be filled with '-'      customerCodeString (1)FixedNAlphanumericFlag to determine PRIO/SOLI card, conventional

debit card, or BDC

BDC = T

If error will be filled with '-'      customFieldString (22)MaxNAlphanumericName embossed on BDC

If error will be filled with '-'      cardStatusString (1)FixedNAlphanumericCard current status

O: Open

B: Blocked

H: Half-Blocked

C: Closed

If error will be filled with '-'      corpIdString (6)FixedNNumericCompany Code

If error will be filled with '-'      outstandingString (1)FixedNNumericOutstanding status of debit card

For BDC will be filled with '0'

If error will be filled with '-'      sendFlagString (1)FixedNAlphanumericFlag to determine where the card will be sent.

B: Branch

H: Home

BDC will be filled with 'B'

If error will be filled with '-      idNumberString (16)FixedNNumericCustomer ID number

If error will be filled with '-'      customerIdString (12)FixedNNumericField is not used

If error will be filled with '-'      maidenNameString (28)MaxNAlphanumericMother's maiden name

If error will be filled with '-'      corpTypeString (1)FixedNAlphanumericBDC card type

P: Petty Cash

D: Deposit

L: Loyalty

If error will be filled with '-'      uniqueKeyString (20)FixedNAlphanumericCustomer Unique Key

If error will be filled with '-'      limitUsageString (16,2)MaxNISO 4217Amount of limit usage of BDC

If error will be filled with '-'      limitNextDateString (8)FixedNNumericTime of next limit reset

If error will be filled with '-'      limitExpiryDateString (8)FixedNNumericLimit expiry date

If error will be filled with '-'      phoneString (15)MaxNNumericPhone Number

If error will be filled with '-'

      emailString (50)MaxNAlphanumericEmail Address

If error will be filled with '-'      descString (25)MaxNAlphanumericNotes

If error will be filled with '-'      logonIdString (8)MaxNAlphanumericFilled with user ID

If error will be filled with '-'      reasonCodeString (1)FixedNAlphanumericCard status change reason

H: Hilang (Lost)

R: Tertelan (Swallowed)

O: Lain-lain (Others)

If error will be filled with '-'      resetTypeString (1)FixedNAlphanumericLimit reset type

N: No Reset

D: Reset by Date

X: Reset by Day

If error will be filled with '-'      resetPeriodString (3)FixedNNumericLimit reset period

If error will be filled with '-'      phoneEcommString (15)MaxNAlphanumericEcommerce Phone Number

If error will be filled with '-'

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message (Indonesian)**Error Message (English)****200**200IC00SuccessfulTransaction successful.**400**400IC00Bad RequestGeneral request failed error**400**400IC01Invalid Field Format {Field}Field does not match the format.**401**401IC00Unauthorized. \[Reason\]Invalid Signature/Unknown Client/Connection Not Allowed.**401**401IC01Invalid token (B2B)Access Token Not Exist/Access Token Expired/Token is Invalid.**403**40ICL01Feature Not Allowed The feature is not allowed to use. Please make sure you registered this feature in your partnership data.**404**404IC11Invalid CardCard number is not found.**409**409IC00ConflictX-EXTERNAL-ID duplicate.**500**500IC01Internal Server ErrorUnknown internal server failure, please retry the process again.**504**504IC00TimeoutYour transaction timeout.

curl –X POST https://sandbox.bca.co.id/openapi/bdc/v1.0/inquiry-card

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'Content-type:application/json'
-H 'Authorization:Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-TIMESTAMP:2023-12-25T13:59:21+08:00'
-H 'X-SIGNATURE:85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'ORIGIN:www.hostname.com'
-H 'CHANNEL-ID:95411'
-H 'X-PARTNER-ID:012345'
-H 'X-EXTERNAL-ID:01234506022024001'
-d '
{
	"bankCardToken": "1234567890123456"
}
```

curl –X POST https://sandbox.bca.co.id/openapi/bdc/v1.0/inquiry-card

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
-H 'Content-type:application/json'
-H 'X-TIMESTAMP:2020-12-18T14:48:30+07:00'
{
	"responseCode": "200IC00",
	"responseMessage": "Successful",
	"responseData":
	{
		"accountNo": "",
		"name": "John Doe",
		"accountInfos":
		[{
			"balanceType": "Limit",
			"amount":
			{
				"value": "20000.00",
				"currency": "IDR"
			}
		}],
		"additionalInfo":
		{
			"cardNumber": "1234567890123456",
			"facType": "1",
			"totalAccount": "01",
			"lanId": "1",
			"openDate": "20080206",
			"prodDate": "20080206",
			"closeDate": "00000000",
			"expiryDate": "4912",
			"embName": "st Code Corp Regresi 5",
			"addr1": "JKT",
			"addr2": "JKT",
			"addr3": "THAMRIN",
			"birthDate": "00000000",
			"branchCode": "0437",
			"empCode": "F",
			"chargeFlag": "P",
			"type": "0",
			"deptCode": "00000",
			"stockBranchCode": "0437",
			"customerCode": "T",
			"customField": "REG CORP ENG 13",
			"cardStatus": "O",
			"corpId": "000097",
			"outstanding": "0",
			"sendFlag": "B",
			"idNumber": "2100001000001103",
			"customerId": "",
			"maidenName": "",
			"corpType": "L",
			"uniqueKey": "R000052",
			"limitUsage": "0",
			"limitNextDate": "20991231",
			"limitExpDate": "20180110",
			"phone": "08561589203",
			"email": "REG-CORP13@regresi.COM",
			"desc": "REG CORP LOYALITY 13",
			"logonId": "userkbb",
			"reasonCode": "",
			"resetType": "N",
			"resetPeriod": "000",
			"phoneEcomm": ""
		}
	}
}
```

##### 3\. Set Business Debit Card Limit API

Your request must contain following information:

This feature is not yet available to be accessed via Sandbox

**Additional Headers and URI Params:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)FixedY Channel's Identifier using WSID = 95411X-PARTNER-IDHeaderString(36)MaxY Partner ID (company code)X-EXTERNAL-IDHeaderString(36)MaxYNumeric String. Reference number that should be unique in the same day

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptiontransactionIdString(14)FixedY NumericTransaction identifiereffectiveDateString(8)FixedY NumericEffective date of limit setting

Guide: Fill with today's dateaccountNumberString(10)FixedYNumericAccount Number of the BDCcardTypeString (1)FixedYAlphanumericBDC type:

L: Loyalty

P: Petty Cash

D: Deposit

Fill with capital lettersrecurringTypeString (1)FixedYAlphanumericRecurring type of BDC limit

N: No Reset (can be used by card type 'P' & 'L')

D: Reset by Date (can be used by card type 'P')

X: Reset by Day (can be used by card type 'P' & 'D')

Fill with capital lettersrecurringPeriodString (3)MaxYNumericPeriodic time for BDC limit

card type D = '1'

card type L = '000'

card type P =

'000' if recurring type 'N'

'1' - '31' if recurring type 'D'

free number if recurring type 'X' cardNumberString (16)FixedYNumericCard Number for the BDCexpiryLimitString (8)FixedYNumericLimit Expiration Date

Guide:

Fill with today's date or biggerlimitAmountString (13)MaxYAlphanumericLimit Amount of BDC

Guide:

1\. Fill with numerical value -> change the amount limit of BDC

2\. Fill with space -> does not change the amount limit of BDC

3\. Fill with other value (alphabetical letters/special characters) -> ErroruniqueKeyString (20)MaxYAlphanumericCustomer Unique Key

Guide:

For new BDC card :

1\. Fill with different unique key from other BDC cards

For existing BDC card, existing unique key will be replaced with the inserted value and may be filled with:

1\. Same unique key as previous card limit setting, or

2\. Different unique key from previous card limit settingcardholderString (35)MaxNAlphanumericName of Card Holder of BDCidNumberString (16)FixedNAlphanumericID Number of the Card Holder of BDCphoneNumberString (15)MaxNAlphanumericPhone Number of the Card Holder of BDCemailAddressString (50)MaxNAlphanumeric Email Address of the Card Holder of BDCdescriptionString (25)MaxNAlphanumericDetail informationprocessTypeString (1)FixedYAlphanumericLimit Setting Process Type

S: Single Process

B: Bulk Process

Can only use 'S'

Fill with capital letters

Result of the request will contains following information:

**Response:**

**Field****Data Type****Length Type****Mandatory****Format****Description**responseCodeString (7)FixedYAlphanumericResponse code to identify transaction status (success or failed)

Format: AAABBCC

AAA: HTTP Code

BB: Service Code

CC: Case Code

Service Code API BDC Setting Limit: SL

Example: 200SL00 – SuccessfulresponseMessageString (150)MaxYAlphanumericResponse description

Example: SuccessfulresponseDataObjectY   transactionIdString (14)FixedYNumericTransaction identifier on service provider system.

Must be filled upon successful transaction.

If error will be filled with '-'   updateDateTimeString (14)FixedYNumericDate and Time of the transaction

If error will be filled with '-'   cardNumberString (16)FixedYNumericCard Number for the BDC

If error will be filled with '-'

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message (Indonesian)**Error Message (English)****200**200SL00SuccessfulTransaction successful.**400**400SL00Bad RequestGeneral request failed error**400**400SL01Invalid Field Format {Field}Field does not match the format.**401**401SL00Unauthorized. \[Reason\]Invalid Signature/Unknown Client/Connection Not Allowed.**401**401SL01Invalid token (B2B)Access Token Not Exist/Access Token Expired/Token is Invalid.**403**403SL01Feature Not Allowed The feature is not allowed to use. Please make sure you registered this feature in your partnership data.**403**403SL02Exceeds Transaction Amount LimitTransaction limit exceeds BDC Card limit.**403**403SL07Card BlockedCard Status = H (Half Blocked / B (Blocked) / C (Closed) **403**403SL08Card ExpiredCard is expired. **403**403SL15Transaction Not Permitted. \[Reason\]recurringType or recurringPeriod does not match the format. **404**404SL11Invalid CardCard number is not found.**404**404SL11Invalid Customer uniqueKeyuniqueKey not unique.**404**404SL11Invalid AccountAccount number is not found. **409**409SL00ConflictX-EXTERNAL-ID duplicate.**500**500SL01Internal Server ErrorUnknown internal server failure, please retry the process again.**504**504SL00TimeoutYour transaction timeout.

curl –X POST https://sandbox.bca.co.id/openapi/bdc/v1.0/setting-limit

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'Content-type:application/json'
-H 'Authorization:Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-TIMESTAMP:2023-12-25T13:59:21+08:00'
-H 'X-SIGNATURE:85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'ORIGIN:www.hostname.com'
-H 'CHANNEL-ID:95411'
-H 'X-PARTNER-ID:012345'
-H 'X-EXTERNAL-ID:01234506022024001'
-d '
{
	"transactionId": "12345600098765",
	"effectiveDate": "01022024",
	"accountNumber": "1234567890",
	"cardType": "D",
	"recurringType": "X",
	"recurringPeriod": "007"
	"cardNumber": "5307902332213091",
	"expiryLimit": "10022024",
	"limitAmount": "99999999999",
	"uniqueKey": "12345678901234567890"
	"cardHolder": "Joni",
	"idNumber": "1234567890123456",
	"phoneNumber": "123456789012",
	"emailAddress": "xxxxxxxxxxxxx@mail.com",
	"description": "",
	"processType": "S"
}'
```

curl –X POST https://sandbox.bca.co.id/openapi/bdc/v1.0/setting-limit

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
-H 'Content-type:application/json'
-H 'X-TIMESTAMP:2020-12-18T14:48:30+07:00'
{
	"responseCode": "200SL00",
	"responseMessage": "Successful",
	"responseData": {
		"transactionId": "12345600098765",
		"updateDateTime": "10022024",
		"cardNumber": "5307902332213091"
	}
}
```

#### Informasi Rekening

##### 1\. Notification Account Statement Offline

Service for BCA to send notification that File Statement is ready to download.

This feature is not yet available to be accessed via Sandbox

**Request:**

FieldDataType**Mandatory**DescriptionRequestIDString(21)YID for every requestAccountNumberString(10)YRequested Account NumberStartDateString(10)Y Statement start date. Format: yyyy-MM-ddEndDateString(10)YStatement end date. Format: yyyy-MM-dd

**Response:**

FieldDataTypeDescriptionRequestIDStringStatement request IDResponseWSStringResponse from Co-Partner

NOTIFICATION ACCOUNT STATEMENT OFFLINE

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST https://www.copartners.com/bca/async/statement

-H "X-Pass-Signature:b9125ab10816f6929d214c96ffca77dfd5a1ea13856362b85eeaeb70155"
-d "{
"RequestID": "h2hauto014_0000000551",
"AccountNumber": "0613004197",
"StartDate": "2019-02-20",
"EndDate": "2019-02-20",
}"
```

NOTIFICATION ACCOUNT STATEMENT OFFLINE

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "RequestID" : "h2hauto014_0000000551",
    "ResponseWS" : "1",
}
```

##### 2\. Notifikasi Rekening Dana Nasabah

Retrieve information of Account Statement of Investor Account for being debited or credited.

This feature is not yet available to be accessed via Sandbox

**Request:**

FieldDataType**Mandatory**DescriptionExternalReferenceString(32)YAlphanumeric. Unique reference number for the request sent.SeqNumberString(14)YNumeric. Identifier for the transaction.AccountNumberString(10)Y Numeric. The investor account number.CurrencyString(3)YCurrently always set to IDRTxnDateString(15)YNumeric. In format ”mmddyyyy hhmmss”.TxnTypeString(4)YRefer to table Transaction TypeTxnCodeString(1)YD/C indicate debit or credit transaction.AccountDebitString(23)YNumeric. The debited account (source of fund).AccountCreditString(23)YNumeric. The credited account (transfer destination account).TxnAmountString(13.2)YNumeric. Money nominal of the transfer.OpenBalanceString(13.2)YNumeric. Possible signed amount with - / + symbol.CloseBalanceString(13.2)YNumeric. Possible signed amount with - / + symbol.TxnDescString(90)YAny description.

Result of the request will contains following information:

**No.**Transaction TypeTransaction CodeDescription 1.NTRF DFund Transfer from KlikBCA Bisnis

MFTS – Auto collection from KlikBCA Bisnis CFund Transfer from KlikBCA Bisnis, KlikBCA Individu, m-BCA, BCA by Phone, ATM, ATM Kiosk

MFTS – Auto credit from KlikBCA Bisnis

Cash deposit, Fund Transfer, Credit Note from Branch

Deposit Clearing 2.NINTCInterest

Giro Service 3.NREVDCredit Reversal 4.NKORDDebit CorrectionC Credit Correction 5.NTAXDTax 6.NCHGDAdmin Fee

When Transaction Type (TxnType) value is NREV, the value of External Reference (ExternalReference) will be the same with the reversed transaction (the original transaction with same ExternalReference number but TxnType is NTRF).

**Response:**

FieldDataTypeDescriptionResponseWSStringFlag received status.

0 : request successfully accepted and no internal error

1 : request unsuccessfully accepted and or some internal error occurred

NOTIFICATION FOR SECURITIES

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST https://www.somesecurities.com/bca/investor-account/statement
{
"ExternalReference":"AAABBBCCCDDDEEEE1234567890123456",
"SeqNumber":"10023",
"AccountNumber":"5421036985",
"Currency":"IDR",
"TxnDate":"08252017 143801",
"TxnType":"NTRF",
"TxnCode":"C",
"AccountDebit":"1704000420",
"AccountCredit":"5421036985",
"TxnAmount":"200000.00",
"OpenBalance":"-150000.00",
"CloseBalance":"+50000.00",
"TxnDesc":"-"
}
```

NOTIFICATION FOR SECURITIES

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "ResponseWS":"0"
}
```

##### 3\. Request Mutasi Rekening Bulk

You can get your bulk statement in form of file for a period up to 7 days. Your Request must contain following information:

This feature is not yet available to be accessed via Sandbox

##### Headers

FieldDataTypeMandatoryDescriptionChannelIDString(5)YChannel Identification Number (Ex: 95051 for KlikBCA Bisnis) CredentialIDString(10)YYour Channel Identity (ex: Your KlikBCA Bisnis CorporateID)

##### Payload

FieldDataType MandatoryDescriptionAccountNumberString(10)YAccount NumberStartDateString(10)YStart Date of the account statement that you wants to get. Format: yyyy-MM-ddEndDateString(10)YEnd Date of the account statement that you wants to get. Format: yyyy-MM-dd

##### Response

FieldDataTypeDescriptionRequestIDString(21)Statement request IDTransactionDateString(10)Transcation date. Format: yyyy-MM-dd

##### Error

HTTP CodeError CodeError Message (Indonesian)Error Message (English)404 ESB-07-271 Tidak ada transaksi No transaction400 ESB-07-279 Tanggal awal lebih kecil dari tanggal buka rekening Start date less than account open date400 ESB-14-001 HMAC tidak cocok HMAC mismatch400 ESB-14-002 Permintaan tidak valid Invalid request400 ESB-14-003 Timestamp tidak valid Invalid timestamp400 ESB-14-004 parameter input tidak valid Invalid input parameter500 ESB-14-005 Sistem sedang dalam maintenance System under maintenance504 ESB-14-006 Timeout, silakan periksa mutasi rekening Timeout, please check your account statement504 ESB-14-007 Timeout Timeout401 ESB-14-008 client\_id/client\_secret/grant\_type tidak valid Invalid client\_id/client\_secret/grant\_type401 ESB-14-009 Tidak berhak Unauthorized429 ESB-14-010 Jumlah permintaan melebihi limit Limit request exceeded400 ESB-14-015 Content Type tidak valid Invalid Content Type400 ESB-14-016 Format JSON tidak valid Invalid JSON format403 ESB-14-019 Koneksi tidak diperbolehkan Connection not allowed400 ESB-14-020 Request tidak valid Invalid Request400 ESB-14-021 API Key Tidak Valid API Key Invalid400 ESB-14-022 API Key/API Secret tidak cocok API Key/API Secret mismatch400 ESB-82-001 Field \[FieldName\] harus diisi Missing mandatory field \[FieldName\]400 ESB-82-002 Corporate Id tidak valid Invalid CorporateID400 ESB-82-003 TransactionID tidak unik TransactionID not unique400 ESB-82-004 TransactionDate tidak valid Invalid TransactionDate403 ESB-82-006 Nominal transaksi melebihi batas maksimum Max amount transaction is exceeded400 ESB-82-008 Rekening perusahaan tidak valid Company account invalid400 ESB-82-009 CurrencyCode tidak valid Invalid CurrencyCode400 ESB-82-010 Format TransactionID tidak valid Invalid TransactionID format400 ESB-82-011 Amount tidak valid Invalid Amount400 ESB-82-012 Panjang karakter data input tidak valid Invalid input length400 ESB-82-013 Tipe data input tidak valid Invalid data type400 ESB-82-014 AccountNumber tidak valid Invalid AccountNumber403 ESB-82-015 Min Max Amount Execeeded Min Max Amount Execeeded403 ESB-82-018 Rekening Dormant Account Dormant402 ESB-82-019 Saldo tidak cukup Insufficient fund400 ESB-82-020 Tipe rekening tidak valid Invalid account type403 ESB-82-021 Rekening tidak dapat melakukan transaksi Account cannot do transaction404 ESB-82-022 Rekening tutup Account closed500 ESB-82-023 Transaksi gagal Transaction failed504 ESB-82-024 Timeout, silakan periksa mutasi rekening Timeout, please check your account statement500 ESB-82-025 Sedang diproses, silakan periksa mutasi rekening In progress, please check your account statement500 ESB-82-026 Lewat batas waktu cut off Cut off time is exceeded400 ESB-99-009 Field \[FieldName\] harus diisi Missing mandatory field \[FieldName\]403 ESB-99-075 KeyID tidak ditemukan KeyID is not found403 ESB-99-089 Transaksi berhasil sebagian Transaction success partially400 ESB-99-112 Input string JSON tidak valid Invalid JSON string input400 ESB-99-113 Transaksi ditolak Transaction rejected400 ESB-99-128 Panjang field \[FieldName\] melebihi ketentuan Field \[FieldName\] exceed limit 400 ESB-99-128 Total input \[FieldName\] melebihi ketentuan Total \[FieldName\] input exceed limit400 ESB-99-156 ChannelType tidak valid ChannelType is not valid400 ESB-99-156 AccountNumber Tidak Valid Invalid AccountNumber400 ESB-99-157 Transaksi ditolak Transaction rejected400 ESB-99-158 Tanggal berakhir lebih besar dari tanggal hari ini End date more than today400 ESB-99-158 Tanggal mulai lebih besar dari tanggal hari ini Start date more than today400 ESB-99-158 Maksimal mutasi rekening yang dipilih 31 hari yang lalu Maximal account statement 31 days ago400 ESB-99-158 Tanggal mulai lebih besar dari tanggal berakhir Start date more than end date404 ESB-99-127 Transaksi tidak ditemukan Transaction not found400 ESB-99-197 Nama atau nomor rekening tidak sesuai Account name/number does not match400 ESB-99-284 \[FieldName\] tidak valid Invalid \[FieldName\]500 ESB-99-999 Sistem sedang tidak tersedia System unavaliable

GET /banking/offline/corporates/accounts/{AccountNumber}/filestatements

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
GET /banking/offline/corporates/accounts/0201245680/filestatements?StartDate=2016-06-24&EndDate=2016-06-24 HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
```

GET /banking/offline/corporates/accounts/{AccountNumber}/filestatements

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "RequestID" : "ABCDEFGHIJ_0000000752",
    "ResponseWS" : "Request Account Statement Berhasil Diterima dan Sedang
Diproses",
}
```

##### 4\. SNAP Banking Balance Inquiry

This service is used to inquiry account balance.

##### Additional Headers and URI Params

FieldParams Type**Data Type****Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString (5)FixedYChannel’s identifier using WSID KlikBCA Bisnis (95051)X-PARTNER-IDHeaderString (32)MaxYCorporate ID

##### Payload

FieldDataType**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString(64)Max Y-Transaction identifier on service customer system accountNoString(16)MaxY NumericRegistered account number in KlikBCA Bisnis. For BCA, length account number is 10 digit
Result of the request will contains following information:

##### Response

FieldDataType**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)FixedY-Response coderesponseMessageString (150)MaxY-Response description referenceNoString (64)MaxY (When the transaction is successful)NumericTransaction identifier on service provider system. Unique each day from BCA

Must be filled upon successful transaction

With format as follows : yymmddHH (from timestamp, 8 digit) xx (prefix menu, 2 digit) sequence number (8 digit) partnerReferenceNoString (64)MaxY-Transaction identifier on service customer system

\*according to the request inputaccountNoString (32)MaxNNumericSend by BCA according to the request input.

Registered account number in KlikBCA Bisnis. For BCA, length account number is 10 digit.nameString (140)MaxNAlphanumericSend by BCA. Customer account nameaccountInfosArray of objectStarting and ending balance before the first/last transaction.   amountObjectYSend by BCA. Net amount of the transaction      valueString(16.2)MaxYNumericIf it’s IDR then value includes 2 decimal digits.

BCA will send response amount with format value Numeric (13.2)      currencyString(3)FixedYAlphanumericCurrency ISO 4217   floatAmountObjectNSend by BCA.

Amount of deposit that is not effective yet (due to holiday, etc)      valueString(16.2)MaxYNumericIf it’s IDR then value includes 2 decimal digits.

BCA will send response amount with format value Numeric (13.2)      currencyString(3)FixedYAlphanumericCurrency ISO 4217   holdAmountObjectNSend by BCA.

Hold amount that cannot be used      valueString(16.2)MaxYNumericIf it’s IDR then value includes 2 decimal digits.

BCA will send response amount with format value Numeric (13.2)      currencyString(3)FixedYAlphanumericCurrency ISO 4217availableBalanceObjectYSend by BCA.

Account balance that can be used for financial transaction      valueString(16.2)MaxYNumericIf it’s IDR then value includes 2 decimal digits.

BCA will send response amount with format value Numeric (13.2)      currencyString(3)FixedYAlphanumericCurrency ISO 4217

##### Error

HTTP CodeError CodeError Message (Indonesian)4004001100Bad request400 4001101Invalid Mandatory Field {Field}400 4001102Invalid Field Format {Field}4014011100Unauthorized. \[Reason\]4014011101Invalid token (B2B) 4034031101Feature Not Allowed4034031118Inactive Account4044041111Invalid Account4094091100Conflict500 5001100General Error5045041100Timeout

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/balance-inquiry

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95051'
-H 'X-PARTNER-ID': 'KBBABCINDO'
-d '{
 "partnerReferenceNo": "2020102900000000000001",
 "accountNo": "1234567890"
}
```

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/balance-inquiry

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
 "responseCode": "2001100",
 "responseMessage": "Successful",
 "referenceNo": "2020102977770000000009",
 "partnerReferenceNo": "2020102900000000000001",
 "accountNo": "1234567890",
 "name": "ANDHIKA",
 "accountInfos": {
 "amount": {
 "value": "100000.00",
 "currency": "IDR"
 },
 "floatAmount": {
 "value": "500000.00",
 "currency": "IDR"
 },
 "holdAmount": {
 "value": "200000.00",
 "currency": "IDR"
 },
 "availableBalance": {
 "value": "200000.00",
 "currency": "IDR"
 }
}
}

```

##### 5\. SNAP Banking Bank Statement

This service is used to inquiry account statement.

##### Additional Headers and URI Params

FieldParams Type**Data Type****Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString (5)FixedYChannel’s identifier using WSID KlikBCA Bisnis (95051)X-PARTNER-IDHeaderString (32)MaxYCorporate ID

##### Payload

FieldDataType**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString(64)Max Y-Transaction identifier on service customer system accountNoString(16)MaxY NumericRegistered account number in KlikBCA Bisnis. For BCA, length account number is 10 digitfromDateTimeDate (25)MaxYISODateTime ISO-8601Starting time range

Default : NOW (DESCENDING) or NOW -1 year (ASCENDING)

(must be filled YYYYMMDDT00:00:00+07:00)

For BCA, only provide data for NOW - 31 daystoDateTimeDate (25)MaxYISODateTime ISO-8601Ending time range

Default : NOW (DESCENDING) or NOW – 1 year (ASCENDING)

(must be filled YYYYMMDDT00:00:00+07:00)

For BCA, only provide data for NOW – 31 days

Account mutations appear in the order of the newest transaction to the oldest transaction.

The maximum number of mutations that can be withdrawn in 1x request is 9000.

If fromDateTime and toDateTime are inputted on the request for a holiday date, then the mutation that will be displayed is the mutation from the request date to the date of the next working day.

Result of the request will contains following information:

##### Response

FieldDataType**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)FixedY-Response coderesponseMessageString (150)MaxY-Response description referenceNoString (64)MaxY (When the transaction is successful)NumericTransaction identifier on service provider system. Unique each day from BCA

Must be filled upon successful transaction

With format as follows : yymmddHH (from timestamp, 8 digit) xx (prefix menu, 2 digit) sequence number (8 digit) partnerReferenceNoString (64)MaxY-Transaction identifier on service customer system

\*according to the request inputbalanceArray of objectStarting and ending balance before the first/last transaction.   amountObjectYAmount of balance      valueString(16.2)MaxYNumericIf it’s IDR then value includes 2 decimal digits.

BCA will send response amount with format value Numeric (13.2)      currencyString(3)FixedYAlphanumericCurrency ISO 4217      dateTimeString(25)MaxYISODateTime

8601Timestamp of the balance   startingBalanceObjectNStarting and ending balance before the first/last transaction       valueString(16.2)MaxYNumericIf it’s IDR then value includes 2 decimal digits.

BCA will send response amount with format value Numeric (13.2)      currencyString(3)FixedYAlphanumericCurrency ISO 4217      dateTimeString(25)MaxYISODateTime

8601Timestamp of the balance   endingBalanceObjectNStarting and ending balance before the first/last transaction       valueString(16.2)MaxYNumericIf it’s IDR then value includes 2 decimal digits.

BCA will send response amount with format value Numeric (13.2)      currencyString(3)FixedYAlphanumericCurrency ISO 4217      dateTimeString(25)MaxYISODateTime

8601Timestamp of the balancetotalCreditEntriesObjectNTotal transaction amount with type = CREDIT   numberOfEntriesString (5)MaxYNumericNumber of entries for credit transaction   amountObjectYAmount of balance      valueString(16.2)MaxYNumericIf it’s IDR then value includes 2 decimal digits.

BCA will send response amount with format value Numeric (13.2)      currencyString(3)FixedYAlphanumericCurrency ISO 4217totalDebitEntriesObjectNTotal transaction amount with type =DEBIT   numberOfEntriesString (5)MaxYNumericNumber of entries for credit transaction   amountObjectYAmount of balance      valueString(16.2)MaxYNumericIf it’s IDR then value includes 2 decimal digits.

BCA will send response amount with format value Numeric (13.2)      currencyString(3)FixedYAlphanumericCurrency ISO 4217detailDataArray of object   endAmountObjectYNet amount of the transaction       valueString(16.2)MaxYNumericIf it’s IDR then value includes 2 decimal digits.

BCA will send response amount with format value Numeric (13.2)      currencyString(3)FixedYAlphanumericCurrency ISO 4217   amountObjectY      valueString(16.2)MaxYNumericIf it’s IDR then value includes 2 decimal digits.

BCA will send response amount with format value Numeric (13.2)      currencyString(3)FixedYAlphanumericCurrency ISO 4217   transactionDateString (25)FixedYISO8601Transaction date   remarkString (256)MaxY-Remark / transaction description. For BCA, remark is 36 digit divided in 2 field (remark 1 with 18 digit and remark 2 with 18 digit).

Filled with data from txn\_name and trailer 1-6   typeString (6)FixedYAlphanumericTransaction type CREDIT/DEBIT

##### Error

HTTP CodeError CodeError Message (Indonesian)400 4001402Invalid Mandatory Field {Field}400 4001402Invalid Field Format {Field}4014011400Unauthorized. \[Reason\]4014011401Invalid token (B2B) 4034031401Feature Not Allowed4044041411Invalid Account4094091400Conflict4294291400Too Many Requests500 5001400General Error5045041400Timeout

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/bank-statement

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95051'
-H 'X-PARTNER-ID': 'KBBABCINDO'
-d '
{
	"partnerReferenceNo": "2020102900000000000001",
	"accountNo": "1234567890",
	"fromDateTime": "2021-04-21T00:00:00+07:00",
	"toDateTime": "2021-04-21T00:00:00+07:00"
}'
```

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/bank-statement

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
	"responseCode": "2001400",
	"responseMessage": "Successful",
	"referenceNo": "2020102977770000000009",
	"partnerReferenceNo": "2020102900000000000001",
	"balance": [
	{
		"amount": {
			"value": "500000.00",
			"currency": "IDR",
			"dateTime": "2020-12-19T16:03:45+07:00"
		},
		"startingBalance": {
			"value": "500000.00",
			"currency": "IDR?",
			"dateTime": "2020-12-19T16:03:45+07:00"
		},
		"endingBalance": {
			"value": "500000.00",
			"currency": "IDR?",
			"dateTime": "2020-12-19T16:03:45+07:00"
		}
	}],
	"totalCreditEntries":
	{
		"numberOfEntries": "10",
		"amount": {
			"value": "500000.00",
			"currency": "IDR"
		}
	},
	"totalDebitEntries":
	{
		"numberOfEntries": "10",
		"amount": {
			"value": "500000.00",
			"currency": "IDR"
		}
	},
	"detailData": [
	{
		"amount":
		{
			"value": "10000.00",
			"currency": "IDR"
		},
		"transactionDate": "2020-12-18T16:03:45+07:00",
		"remark": "TRSF E-BANKING DB 1111/ATRTG/WS95051 09112016123456pencairan merchant SITI KHADIJAH",
		"type": "DEBIT",
	},
	{
		"amount":
		{
			"value": "10000.00",
			"currency": "IDR"
		},
		"transactionDate": "2020-12-18T16:03:45+07:00",
		"remark": "BA JASA E-BANKING 0709/FTCHG/WS95321DOM210907115626680PPU_CGAMBIAYA TRANSAKSI",
		"type": "DEBIT",
	}]
}

```

#### Pembukaan Rekening Online

##### 1\. BCA Account Opening Notification

This service is used for BCA account opening approval notification.

This feature is not yet available to be accessed via Sandbox

**Recommended Specification:**

Type**Recommendation**AndroidOS version 8.0+ dan Google Chrome version 75+iOSOS version 12.0+ dan Safari version 12+

**Request Headers:**

FieldData Type**Mandatory**DescriptionX-Pass-SignatureStringY This is a header field, use to replace the “Pass” field on everybody payload that contains signature.

This field occurred on every BCA request to Copartner’s API.

Notes: the field name “X-Pass-Signature” can be replaced with any other name depends on BCA’s company partner’s specification.

**Payload (Form Data):**

FieldData Type**Mandatory**Descriptionapplication\_idString (36)YID Submission Form created by prospective customersuser\_idString (50)YProspective customer ID of each merchant

This user ID is sent from the merchant when they want to generate a tokenstatus\_dateString (19)Ydate and time when the poll was success / rejected

Format : YYYY-MM-DD HH:MM:SScreated\_dateString (19)YThe date when the initial form submission processmodified\_dateString (19)YThe date in case of data change such as updating of eform numbers or other things will automatically update the datestatus\_idintYOpen account status code

Code : 0, 1status\_nameString (15)Ytype of account opening status. depends on the status\_id

0 = REJECTED

1 = SUCCESSnomor\_rekeningString (10)Nif the customer agrees to the provision of the account to a third party, the customer's account number will be sent

Result of the request will contains following information:

**Response:**

FieldData TypeDescriptionStatusString (7)The push notification process is success or failedMessageString (8)Response from Co-Partner (“Berhasil”)

https://www.copartners.com/bca/open-account/notification

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST https://www.copartners.com/bca/open-account/notification
-H "X-Pass-Signature: b9125ab10816f6929d214c96ffca77dfd5a1ea13856362b85eeaeb70155"
-d " {
	"user_id":"DD6DC0C5-17ED-4731-8600-B63EEDB126D9",
	"application_id":"7C73063B-93EA-4F57-A3bE-2A1F1773DB4F",
	"status_date":"2020-08-05 11:03:56",
	"created_date":"2020-03-18 09:03:57",
	"modified_date":"2020-08-11 11:16:01",
	"status_id":"1",
	"status_name":"SUCCESS",
	"nomor_rekening":"1234567890"
}"
```

https://www.copartners.com/bca/open-account/notification

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
	"Status": "Success",
	"Message": "Berhasil"
}
```

##### 2\. BCA Account Opening Status

This service is used for BCA account opening status update for Co-Partners.

This feature is not yet available to be accessed via Sandbox

**Recommended Specification:**

Type**Recommendation**AndroidOS version 8.0+ dan Google Chrome version 75+iOSOS version 12.0+ dan Safari version 12+

**Request Headers and URI Params:**

Field**Params Type**Data Type**Mandatory**Descriptionchannel-idHeaderStringY WSID bcacoid (95171)credential-idHeaderString(36)Y Registered Merchant ID at bca.co.idapplication-idHeaderString(36)YApplication ID that sent from bca.co.id to merchant when successfully submitted form (before uploading document and video call)

Result of the request will contains following information:

**Response:**

FieldData TypeDescriptionmerchant\_idStringMerchant ID application\_idString (36)Application IDstatus\_nameString (15)Transaction Statusstatus\_idintStatus ID

0 = Rejected

1 = Success

2 = Retry

3 = On-Lockcreated\_dateString (19)The date when the initial form submission processmodified\_dateString (19)The date in case of data change such as updating of eform numbers or other things will automatically update the date status\_dateString (19)Status Datenomor\_rekeningString (10)Account Number (If the Customer does not check Account Number Approval, then the Account\_Number field will display an empty string (“”))

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message (Indonesian)**Error Message (English)****Description****400**ESB-14-001HMAC tidak cocokHMAC mismatch400**400**ESB-14-002Permintaan tidak validInvalid request400**400**ESB-14-003Timestamp tidak validInvalid timestamp400**400**ESB-14-004parameter input tidak valid Invalid input parameter400**500**ESB-14-005Sistem sedang dalam maintenanceSystem under maintenance500**504**ESB-14-007TimeoutTimeout504**401**ESB-14-008Client\_Id/Client\_Secret /Grant\_Type tidak validInvalid Client\_Id/Client\_Secret /Grant\_Type401**401**ESB-14-009Tidak berhakUnauthorized401**404**ESB-14-011Service tidak tersediaService doesn't exist404**401**ESB-14-012Tidak berhak mengakses service iniNot allowed to access this service401**401**ESB-14-014ChannelID/CredentialID tidak validInvalid ChannelID/CredentialID 401**403**ESB-14-019Koneksi tidak diperbolehkanConnection not allowed 403**400**ESB-14-021API Key tidak validInvalid API Key400**400**ESB-14-022API Key/API Secret tidak cocokAPI Key/API Secret mismatch 400**400**WBF-04-001Header Application Id tidak adaHeader Application Id tidak adaGet Open Account Approval Status**400**WBF-04-002Header Credential ID tidak adaHeader Credential ID tidak adaGet Open Account Approval Status**400**WBF-04-003Credential ID / Application Id Value tidak sesuaiCredential ID / Application Id Value tidak sesuaiGet Open Account Approval Status**400**WBF-04-004Data dengan Credential ID dan Application Id tersebut tidak ditemukanData dengan Credential ID dan Application Id tersebut tidak ditemukanGet Open Account Approval Status**400**WBF-04-005Header Client Id tidak adaHeader Client Id tidak adaGet Open Account Approval Status**400**WBF-04-006Data dengan Client ID tersebut tidak ditemukanData dengan Client ID tersebut tidak ditemukanGet Open Account Approval Status**400**WBF-99-991ex.Message ex.Message Pesan global dari sistem**400**WBF-99-992messagesmessagesPesan global dari sistem

/open-account/marketplace/status

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
curl –X GET https://sandbox.bca.co.id/open-account/marketplace/status
-H
'credentialid': F90338EA-853B-4512-AE58-0E72BE2E88ED',
'application-id':'C5276A67-711E-468EBB14-01861D2746C6',
‘channel-id’:’95171’
```

/open-account/marketplace/status

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
	"merchant_id": "F90338EA-853B-4512-AE58-0E72BE2E88ED",
	"application_id": "5F7D49A6-E3F7-493A-B255-5906D4E65FD2",
	"status_name": "SUCCESS",
	"status_id": "1",
	"created_date": "2020-08-05 10:28:21",
	"modified_date": "2020-08-05 11:03:56",
	"status_date": "2020-08-05 11:03:56",
	"nomor_rekening": "0842378425"
}

Response Berhasil dan tanpa menampilkan Nomor Rekening
{
	"merchant_id": "F90338EA-853B-4512-AE58-0E72BE2E88ED",
	"application_id": "5F7D49A6-E3F7-493A-B255-5906D4E65FD2",
	"status_name": "SUCCESS",
	"status_id": "1",
	"created_date": "2020-08-05 10:28:21",
	"modified_date": "2020-08-05 11:03:56",
	"status_date": "2020-08-05 11:03:56",
	"nomor_rekening": ""
}
```

##### 3\. BCA Account Opening Webform

This service is used for Open Account Marketplace.

This feature is not yet available to be accessed via Sandbox

**Recommended Specification:**

Type**Recommendation**AndroidOS version 8.0+ dan Google Chrome version 75+iOSOS version 12.0+ dan Safari version 12+

**Request Headers and URI Params:**

Field**Params Type**Data Type**Mandatory**Descriptionchannel-idHeaderStringY WSID bcacoid (95171)credential-idHeaderString(36)Y Registered Merchant ID at bca.co.id

**Payload (Form Data):**

FieldData Type**Mandatory**Descriptionuser\_idString (50)YRegistered User Merchant IDapplication\_idString (36)ConditionalApplication ID

Not mandatory if activity\_code = ‘01’, application\_id = ‘’

Mandatory if activity\_code = ‘02’,

application\_id sent from web-view to merchant when successfully submitted form activity\_codeString (2)YActivity Code

‘01’ = web-view access (consumer profile)

‘02’ = mobile-browser access (open account document and video call capture) emailString (100)NMerchant Emailmobile\_phoneString (15)NMerchant Mobile Phonefull\_nameString (100)NMerchant Full Name

Result of the request will contains following information:

**Response:**

FieldData TypeDescriptionverification\_idString (36)Verification IDapplicant\_idString (36)Applicant IDurl\_pageviewStringURL Pageviewapplication\_idString (40)Hanya dikirim jika Activity code = 02

Kode unik khusus yang berfungsi sebagai reference code sesi pemrek yang dimasking untuk diinfokan ke copartner

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message (Indonesian)**Error Message (English)****Description****400**ESB-14-001HMAC tidak cocokHMAC mismatch400**400**ESB-14-002Permintaan tidak validInvalid request400**400**ESB-14-003Timestamp tidak validInvalid timestamp400**400**ESB-14-004parameter input tidak valid Invalid input parameter400**500**ESB-14-005Sistem sedang dalam maintenanceSystem under maintenance500**504**ESB-14-007TimeoutTimeout504**401**ESB-14-008Client\_Id/Client\_Secret /Grant\_Type tidak validInvalid Client\_Id/Client\_Secret /Grant\_Type401**401**ESB-14-009Tidak berhakUnauthorized401**404**ESB-14-011Service tidak tersediaService doesn't exist404**401**ESB-14-012Tidak berhak mengakses service iniNot allowed to access this service401**401**ESB-14-014ChannelID/CredentialID tidak validInvalid ChannelID/CredentialID 401**403**ESB-14-019Koneksi tidak diperbolehkanConnection not allowed 403**400**ESB-14-021API Key tidak validInvalid API Key400**400**ESB-14-022API Key/API Secret tidak cocokAPI Key/API Secret mismatch 400**400**WBF-03-001The {field-name} field is required.The {field-name} field is required.Generate Merchant Token**400**WBF-03-002Parameter header Credential ID tidak ditemukanParameter header Credential ID tidak ditemukanGenerate Merchant Token**400**WBF-03-003Credential ID Value tidak sesuaiCredential ID Value tidak sesuaiGenerate Merchant Token**400**WBF-03-004Credential ID tidak terdaftarCredential ID tidak terdaftarGenerate Merchant Token**400**WBF-03-005Application Id tidak sesuaiApplication Id tidak sesuaiGenerate Merchant Token**400**WBF-03-006Application Id tidak ditemukanApplication Id tidak ditemukanGenerate Merchant Token**400**WBF-03-007Application Id ditolak / berhasil / sedang proses Application Id ditolak / berhasil / sedang proses Generate Merchant Token**400**WBF-03-008Merchant Id dan App Data Id tidak sesuaiMerchant Id dan App Data Id tidak sesuaiGenerate Merchant Token**400**WBF-03-009Telah mencapai batas akses maksimum halaman ini.Telah mencapai batas akses maksimum halaman ini.Generate Merchant Token**400**WBF-99-991ex.Message ex.Message Pesan global dari sistem**400**WBF-99-992messagesmessagesPesan global dari sistem

/open-account/marketplace/merchants/tokens/generator

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
curl –X POST https://sandbox.bca.co.id/open-account/marketplace/merchants/tokens/generator
-H 'credential-id':'F90338EA-853B4512-AE58-0E72BE2E88ED', ‘channel-id’:’95171’
–d
'user_id=123456789&application_id=&activity_code=01&email=maiuw@gmail.com&mobile_ph
one=08128084423&full_name=8196047056'
```

/open-account/marketplace/merchants/tokens/generator

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
Response (if ActivityCode = 01)
{
	"verification_id": "cdec3c69-f549-4b91-bccf-25a3dad2de5a",
	"applicant_id": "117f6c7e-ceec-4fde-ab4d-6b628230ed59",
	"url_pageview": "http://10.20.212.217:2007/applyonline/bukarekening/form1?MerchantId=F90338EA-853B-4512-AE58-0E72BE2E88ED&ApplicantId=117f6c7e-ceec-4fde-ab4d6b628230ed59&VerificationId=cdec3c69-f549-4b91-bccf-25a3dad2de5a"
	"application_id": null
}

Response (if ActivityCode = 02)
{
	"verification_id": "cdec3c69-f549-4b91-bccf-25a3dad2de5a",
	"applicant_id": "117f6c7e-ceec-4fde-ab4d-6b628230ed59",
	"url_pageview": http://10.20.212.217:2007/applyonline/bukarekening/form1?MerchantId=F90338EA-853B-4512-AE58-0E72BE2E88ED&ApplicantId=117f6c7e-ceec-4fde-ab4d6b628230ed59&VerificationId=cdec3c69-f549-4b91-bccf-25a3dad2de5a,
	"application_id": "06B440A2-72E7-47E1-B389-1EDFDC9A636D"
}
```

#### Virtual Account untuk Biller

##### 1\. SNAP Virtual Account Inquiry

This service is used to VA BillPresentment.

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)Y Channel’ Identifier using WSID BCA Virtual Account (95231) X-PARTNER-IDHeaderString(32)Y Partner ID using Company Code VAX-EXTERNAL-IDHeaderString (36)YNumeric String. Reference number that should be unique in the same day

**Payload:**

FieldData Type**Mandatory**DescriptionpartnerServiceIdString (8) YDerivative of XPARTNER-ID, similar to company code

8 digit left padding space “ “customerNoString (20)YUnique number (up to 20 digits)virtualAccountNoString (28)YpartnerServiceId (8 digit left padding space “ “) + customerNo (up to 20 digit)trxDateInitDate (25)NBCA internal system datetime with timezone, which follows the ISO-8601 standardchannelCodeNumeric (4)NChannel code based on ISO 18245inquiryRequestIdString (128) YUnique identifier for this inquiry. Generated by BCA. additionalInfoObjectNOptional. Additional Information for custom use. Refer to Appendix B for complete list of additionalInfo for Virtual Account

Result of the request will contains following information:

**Response:**

FieldData Type**Mandatory**DescriptionresponseCodeString (7)YMandatory in BCA. Response code from Partner, refer to Appendix A for list of response coderesponseMessageString (150)YMandatory in BCA. Response message from Partner, refer to Appendix A for list of response codevirtualAccountDataObjectY   inquiryStatusString (2)NMandatory in BCA. Status of inquiry, refer to Notes for status values    inquiryReasonObjectNMandatory in BCA. Reason for inquiry Status multi language      englishString (64)NReason for inquiry Status in English      indonesiaString (64)NReason for inquiry Status in Bahasa   partnerServiceIdString (8)YMandatory in BCA. Derivative of X-PARTNER-ID, similar to company code, 8 digit left padding space “ ”   customerNoString (20)YMandatory in BCA. Unique number (up to 20 digits).    virtualAccountNoString (28)YMandatory in BCA. partnerServiceId (8 digit left padding space “ ”) + customerNo (up to 20 digit)    virtualAccountNameString (255)YMandatory in BCA. Customer name    inquiryRequestIdString (128)YMandatory in BCA. From Inquiry Request    totalAmountObjectNMandatory in BCA.      valueString (16.2)

ISO 4217YTransaction Amount. Total Amount with 2 decimal      currencyString (3)YCurrency   subCompanyString (5)NMandatory in BCA. Partner's product code (sub company code). Mandatory for non-multibills transaction.billDetailsArray of ObjectsNArray with maximum 24 Objects   billNoString (18)NCustomer bill number generated by Partner   billDescriptionObjectNBill Description      englishString (18)NBill Description in English      indonesiaString (18)NBill Description in Bahasa   billSubCompanyString (5)CBill sub company code   billAmountObjectNAmount specific to bill numbervalueString (16.2)

ISO 4217YTransaction AmountcurrencyString (3) YadditionalInfoObjectN Optional. Additional information for custom use. Refer to Appendix B for complete list of additionalInfo for Virtual AccountfreeTextsArray of ObjectsNOptional. Array with maximum 9 objects. freeTexts field in inquiry bill should not be greater than 5englishString (32) N Will be shown in ChannelindonesiaString (32)N  Will be shown in ChanneladditionalInfoObjectNOptional. Additional information for custom use. Refer to Appendix B for complete list of additionalInfo for Virtual Account

Note :

- The final status of inquiry response is not determined by responseCode and responseMessage.
- Status of inquiry will be defined in inquiryStatus and inquiryReason field to be displayed to channel's screen. This field may vary with these values:


00 = Success inquiry, detail bill returned


01 = Failed inquiry, partner must defined reason for this status in InquiryReason field.
- BCA will continue to check the inquiryStatus and inquiryReason if the responseCode received is 2002400 and 2022400.
- If the response code is not either 2002400 or 2022400, BCA will on consider the response as failed transaction and will be rejected.
- If the inquiryStatus and inquiryReason field are empty, BCA will consider it a failed transaction and will be rejected.
- If virtualAccountName length > 30, channel will display first 30 characters.
- If any amount field value > 13.2 digit length, BCA will consider it a failed inquiry and error message will be displayed in channel.
- currency of totalAmount and billAmount must be the same. currency field may vary with these values: IDR, USD, SGD
- For multi bills transaction, totalAmount field value can be 0, but for non-multi bills transaction totalAmount field value should be total amount of all the detail bills provided.
- For non multi bils transaction with single settlement, billDetails is optional. If billDetails available then billSubCompany field in billDetails must match with subCompany field value.
- For multi bills and non multi bills transaction with multi settlements,
  - billSubCompany & billAmount field is mandatory
  - billSubCompany field value must be different for each bill number, used for multiple settlements needs. Product Name and Admin Fee from subCompany 00000 would be used and shown in channel.
- billNo and billDescription field value are defined by co-partner and will be displayed in BCA channel’s screen
- subCompany and billSubCompany value must be registered in BCA, else transaction will be rejected.
- freeTexts field defined by partner and will be displayed in channel's screen. This field must be defined in two languages, The occurences for freeTexts field in inquiry bill should not be greater than 5.
- feeAmount field will be ignored. feeAmount will refer to company profile data registered in BCA's system.
- virtualAccountTrxType will be ignored. will refer to company profile data (bill type: YV/NV/YF) registered in BCA's system.
- billDetails should not be greater than 5
- channelCode field may vary with these values based on ISO 18245:
  - 6010 = Teller (Financial institutions - manual cash disbursements)
  - 6011 = eBanking (Financial institutions - automated cash disbursements)
- inquiryRequestId is unique from BCA for each transaction.
- the usage of additional info is to be discussed with BCA first.

Expected Response from Partner.

Scenario**responseCode****responseMessage**inquiryStatusAccess Token Invalid4012401"Invalid Token (B2B)"-Unauthorized . Signature4012400"Unauthorized. \[Signature\]" -Unauthorized . stringToSign4012400"Unauthorized. \[Signature\]" -Unauthorized . Unknown Client4012400"Unauthorized. \[Unknown client\]"-Missing Mandatory Field {} …, etc4002402"Invalid Mandatory Field {........}" 01Invalid Field Format {} .., etc4002401"Invalid Field Format {........}" 01Cannot use the same X-EXTERNAL-ID4092400"Conflict" 01Input no Virtual Account Valid2002400"Success" 01Input no Virtual Account Valid sudah lunas 4042414"Paid Bill" 01Input no Virtual Account Valid kadaluarsa4042419"Invalid Bill/Virtual Account" 01Input no Virtual Account tidak terdaftar4042412"Invalid Bill/Virtual Account \[Reason\]" 01Request Parsing Error4002400"Bad Request"01Response Parsing Error4002400"Bad Request"01

curl –X POST https://copartners.com/openapi/v1.0/transfer-va/inquiry

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95231'
-H 'X-PARTNER-ID': '12345'
-H 'X-EXTERNAL-ID': '123456789012345678901234567890123456'
-d '
{
	"partnerServiceId": " 12345",
	"customerNo": "123456789012345678",
	"virtualAccountNo": " 12345123456789012345678",
	"trxDateInit": "2022-02-12T17:29:57+07:00",
	"channelCode": 6011,
	"additionalInfo": {},
	"inquiryRequestId": "202202110909311234500001136962"
}'
```

curl –X POST https://copartners.com/openapi/v1.0/transfer-va/inquiry

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
Sample for single settlement (non multi bills) transaction:

{
	"responseCode": "2002400",
	"responseMessage": "Successful",
	"virtualAccountData":
	{
		"inquiryStatus": "00",
		"inquiryReason": {
			"english": "Success",
			"indonesia": "Sukses"
		},
		"partnerServiceId": " 12345",
		"customerNo": "123456789012345678",
		"virtualAccountNo": " 12345123456789012345678",
		"virtualAccountName": "Jokul Doe",
		"inquiryRequestId": "202202110909314440200001136962",
		"totalAmount": {
			"value": "100000.00",
			"currency": "IDR"
		},
		"subCompany": "00000",
		"billDetails": [
		{
			"billNo": "123456789012345678",
			"billDescription": {
				"english": "Maintenance",
				"indonesia": "Pemeliharaan"
			},
			"billSubCompany": "00000",
			"billAmount": {
				"value": "100000.00",
				"currency": "IDR"
			}
		}],
		"freeTexts": [
		{
			"english": "Free text",
			"indonesia": "Tulisan bebas"
		}],
		"virtualAccountTrxType": "C",
		"feeAmount": {
			"value": "",
			"currency": ""
		},
		"additionalInfo": {
			"additionalInfo1": {
				"label": {
					"indonesia": "Unit",
					"english": "Unit"
				},
				"value": {
					"indonesia": "10C",
					"english": "10C"
				}
			},
			"additionalInfo2": {
				"label": {
					"indonesia": "Bulan",
					"english": "Month"
				},
				"value": {
					"indonesia": "Januari",
					"english": "January"
				}
			}
		}
	}
}

Sample for multi settlement (non multi bills/multi bills) transaction:

{
	"responseCode": "2002400",
	"responseMessage": "Successful",
	"virtualAccountData":
	{
		"inquiryStatus": "00",
		"inquiryReason":
		{
			"english": "Success",
			"indonesia": "Sukses"
		},
		"partnerServiceId": " 12345",
		"customerNo": "123456789012345678",
		"virtualAccountNo": " 12345123456789012345678",
		"virtualAccountName": "Jokul Doe",
		"inquiryRequestId": "202202110909314440200001136962",
		"totalAmount":
		{
			"value": "100000.00",
			"currency": "IDR"
		},
		"subCompany": "",
		"billDetails": [
		{
			"billNo": "123456789012345678",
			"billDescription":
			{
				"english": "Maintenance",
				"indonesia": "Pemeliharaan"
			},
			"billSubCompany": "00000",
			"billAmount":
			{
				"value": "50000.00",
				"currency": "IDR"
			}
		},
		{
			"billNo": "223456789012345678",
			"billDescription":
			{
				"english": "Electricity",
				"indonesia": "Listrik"
			},
			"billSubCompany": "00001",
			"billAmount":
			{
				"value": "20000.00",
				"currency": "IDR"
			}
		},
		{
			"billNo": "323456789012345678",
			"billDescription": {
				"english": "Water",
				"indonesia": "Air"
			},
			"billSubCompany": "00002",
			"billAmount": {
				"value": "30000.00",
				"currency": "IDR"
			}
		}
		],
		"freeTexts": [
		{
			"english": "Free text",
			"indonesia": "Tulisan bebas"
		}],
		"virtualAccountTrxType": "C",
		"feeAmount": {
			"value": "",
			"currency": ""
		},
		"additionalInfo": {
			"additionalInfo1": {
				"label": {
					"indonesia": "Unit",
					"english": "Unit"
				},
				"value": {
					"indonesia": "10C",
					"english": "10C"
				}
			},
			"additionalInfo2": {
				"label": {
					"indonesia": "Bulan",
					"english": "Month"
				},
				"value": {
					"indonesia": "Januari",
					"english": "January"
				}
			}
		}
	}
}
```

##### 2\. SNAP Virtual Account Inquiry Status

This service is used to VA Payment Status.

**Additional Headers:**

Field**Params Type**Data Type**Mandatory**DescriptionCHANNEL-IDHeaderString(5)Y Channel’ Identifier using WSID Virtual Account (95231)X-PARTNER-IDHeaderString(32)Y Partner ID using Company Code VAX- EXTERNAL -IDHeaderString(36)YReference number that should be unique in the same day

**Payload:**

FieldData Type**Mandatory**DescriptionpartnerServiceIdString (8) YDerivative of X-PARTNER-ID, similar to company code, 8 digit left padding space “ “customerNoString (20)YUnique number (up to 20 digits)virtualAccountNoString (28)YpartnerServiceId (8 digit left padding space “ “) + customerNo (up to 20 digit)paymentRequestIdString (128)NUnique identifier from Payment generated by BCA. additionalInfoObjectNOptional. Additional Information for custom use. Refer to Appendix B for complete list of additionalInfo for Virtual Account

Result of the request will contains following information:

**Response:**

FieldData Type**Mandatory**DescriptionresponseCodeString (7)YMandatory In BCA. Response code from Partner, refer to Appendix A for list of response coderesponseMessageString (150)YMandatory In BCA. Response message from Partner, refer to Appendix A for list of response messagevirtualAccountDataObjectY   paymentFlagReasonObjectNReason for Payment Status multi language      englishString (200)NReason for inquiry status in English      indonesiaString (200)NReason for inquiry status in Bahasa    partnerServiceIdString (8)YDerivative of XPARTNER-ID, similar to company code, 8 digit left padding space “ “   customerNoString (20)YUnique number (up to 20 digit)   virtualAccountNoString (28)YpartnerServiceId (8 digit left padding space “ “) + customerNo (up to 20 digit)    inquiryRequestIdString (128)YUnique identifier from Inquiry   paymentRequestIdString (128)CUnique identifier for this Payment from BCA. Mandatory if payment happened.    paidAmountObjectN      valueString (16.2)

ISO 4217 YPaid amount with 2 decimal      currencyString (3)YCurrency    totalAmount ObjectN-      valueString (16.2)

ISO 4217 YTransaction Amount. Total Amount from Inquiry with 2 decimal      currencyString (3)YCurrency   transactionDateDate (25)NPayment datetime when the payment happened   referenceNoString (15)NPayment auth code generated by BCA   paymentFlagStatusString (2)NStatus for Payment Flag   billDetailsArray of ObjectsNArray with maximum 24 Objects       billNoString (18)NBill number from Partner      billDescriptionObjectNBill Description         englishString (18)NBill Description in English         indonesiaString (18)NBill Description in Bahasa      billSubCompanyString (5) NPartner’s product code       billAmountObjectNNominal inputted by Customer with 2 decimal         valueString (16.2) ISO 4217YTransaction Amount. Total Amount from Inquiry with 2 decimal         currencyString (3) YCurrency     additionalInfo ObjectNOptional. Additional Information for customer use for each bill. Refer to Appendix B for complete list of additionalInfo for Virtual AccountbillReferenceNoNumeric (15)NBill auth code generated by BCAstatusString (2)NPayment status for specific BillreasonObjectNReason for Payment Status for specific Bill multi languageenglishString (64)NReason for Payment Status for specific Bill in EnglishindonesiaString (64)NReason for Payment Status for specific Bill in Bahasa additionalInfoObjectNOptional. Additional Information for customer use for each bill. Refer to Appendix B for complete list of additionalInfo for Virtual Account

Note :

- Inquiry payment status can only be requested if the payment flag transaction has been made.
- The available payment (transaction) status are D-day and D-1 transaction.
- The key for inquiry payment status can either be
  - partnerServiceId + customerNo
  - virtualAccountNo
  - paymentRequestId
- If all keys are filled in, then the payment status will be enquired based on the paymentRequestID.
- If using partnerServiceId + customerNo or virtualAccountNo as the key, then the payment status will be enquired based on the most recent transaction using that partnerServiceId + customerNo or virtualAccountNo.
- The final status of payment flag is not determined by responseCode and responseMessage.
- Status for payment flag will be defined in paymentFlagStatus and paymentFlagReason field.

Here is the list of error codes that can be returned.

HTTP Code **Error Code**Error Message4004002600Bad Request4004002601Invalid Field Format {field name}4004002602Invalid Mandatory Field {field name}4014012600Unauthorized. \[Reason\]4014012601Invalid token (B2B)4044042601Transaction Not Found5005002600Internal Server Error 5045042600Timeout

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/transfer-va/status

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95231'
-H 'X-PARTNER-ID': '12345’
-H 'X-EXTERNAL-ID': '123456789012345678901234567890123456’
-d '
{
	"partnerServiceId":" 00012345",
	"customerNo":123456789012345678,
	"virtualAccountNo":" 00012345123456789012345678",
    "paymentRequestId":"202202111031031234500001136962",
	"additionalInfo":{}
}'
```

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/transfer-va/status

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
"Sample for single settlement (non multi bills) transaction:"

{
	"responseCode": "2002600",
	"responseMessage": "Success",
	"virtualAccountData": {
		"paymentFlagStatus": "00",
		"paymentFlagReason": {
			"indonesia": "BERHASIL",
			"english": "SUCCESS"
		},
		"partnerServiceId": " 12345",
		"customerNo": "123456789012345678",
		"virtualAccountNo": " 12345123456789012345678",
		"inquiryRequestId": "202202111031031234500001136962",
		"paymentRequestId": "202202111031031234500001136962",
		"paidAmount": {
			"value": "100000.00",
			"currency": "IDR"
		},
		"totalAmount": {
			"value": "100000.00",
			"currency": "IDR"
		},
		"transactionDate": "2022-02-11T10:16:04+07:00",
		"referenceNo": "00113696201",
		"billDetails": [
		{
			"billNo": "123456789012345678",
			"billDescription": {
				"english": "Maintenance",
				"indonesia": "Pemeliharaan"
			},
			"billSubCompany": "00000",
			"billAmount": {
				"value": "100000.00",
				"currency": "IDR"
			},
			"additionalInfo": {},
			"billReferenceNo": "00113696201",
			"status": "00",
			"reason": {
				"english": "Success",
				"indonesia": "Sukses"
			}
		}
		],
	}
}

Sample for multi settlement (non multi bills/multi bills) transaction:
{
	"responseCode": "2002600",
	"responseMessage": "Success",
	"virtualAccountData": {
		"paymentFlagStatus": "00",
		"paymentFlagReason": {
			"indonesia": "BERHASIL",
			"english": "SUCCESS"
		},
		"partnerServiceId": " 12345",
		"customerNo": "123456789012345678",
		"virtualAccountNo": " 12345123456789012345678",
		"inquiryRequestId": "202202111031031234500001136962",
		"paymentRequestId": "202202111031031234500001136962",
		"paidAmount": {
			"value": "100000.00",
			"currency": "IDR"
		},
		"totalAmount": {
			"value": "100000.00",
			"currency": "IDR"
		},
		"transactionDate": "2022-02-11T10:16:04+07:00",
		"referenceNo": "",
		"billDetails": [
		{
			"billNo": "123456789012345678",
			"billDescription": {
				"english": "Maintenance",
				"indonesia": "Pemeliharaan"
			},
			"billSubCompany": "00000",
			"billAmount": {
				"value": "50000.00",
				"currency": "IDR"
			},
			"additionalInfo": {},
			"billReferenceNo": "00113696201",
			"status": "00",
			"reason": {
				"english": "Success",
				"indonesia": "Sukses"
			}
		},
		{
			"billNo": "223456789012345678",
			"billDescription": {
				"english": "Electricity",
				"indonesia": "Listrik"
			},
			"billSubCompany": "00001",
			"billAmount": {
				"value": "20000.00",
				"currency": "IDR"
			},
			"additionalInfo": {},
			"billReferenceNo": "00213696201",
			"status": "00",
			"reason": {
				"english": "Success",
				"indonesia": "Sukses"
			}
		},
		{
			"billNo": "323456789012345678",
			"billDescription": {
				"english": "Water",
				"indonesia": "Air"
			},
			"billSubCompany": "00002",
			"billAmount": {
				"value": "30000.00",
				"currency": "IDR"
			},
			"additionalInfo": {},
			"billReferenceNo": "00313696201",
			"status": "00",
			"reason": {
				"english": "Success",
				"indonesia": "Sukses"
			}
		}
		],
	}
}
```

##### 3\. SNAP Virtual Account Payment

This service is used to VA Payment Flag.

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)Y Channel’ Identifier using WSID BCA Virtual Account (95231)X-PARTNER-IDHeaderString(32)Y Partner ID using Company Code VAX-EXTERNAL-IDHeaderString (36) YNumeric String. Reference number that should be unique in the same day

**Payload:**

FieldData Type**Mandatory**DescriptionpartnerServiceIdString (8) YDerivative of X-PARTNER-ID, similar to company code, 8 digit left padding space “ “ customerNoString (20)YUnique number (up to 20 digits)virtualAccountNoString (28)YpartnerServiceId (8 digit left padding space “ “) + customerNo (up to 20 digit)virtualAccountNameString (255) NCustomer namepaymentRequestIdString (128)YUnique identifier generated by BCA. If payment comes from the Inquiry process, this value must be the same with inquiryRequestId.channelCodeNumeric (4) NChannel code based on ISO 18245hashedSourceAccountNoString (32)NOptional. Source account number in hashsourceBankCodeString (3)NOptional. Source account bank codepaidAmountObjectY-   valueString (16.2)

ISO 4217YPaid Amount with 2 decimal   currencyString (3)YCurrencyreferenceNoString (64)NPayment auth code generated by BCAflagAdviseString (1)NStatus is this a retry notification. (Default value: N)subCompanyString (5)NSub Company code generated by PartnerbillDetailsArray of ObjectsNArray with maximum 24 Objects   billNoString (18)NFrom Inquiry Response   billDescriptionJSON ObjectNFrom Inquiry Response      englishString (18)NFrom Inquiry Response      indonesiaString (18)NFrom Inquiry Response   billSubCompanyString (5)NFrom Inquiry Response   billAmountObjectsN      valueString (16.2) ISO 4217 YTransaction Amount. From Inquiry Response      currencyString (3)YCurrency   billReferenceNoNumeric (15)NFrom Inquiry ResponseadditionalInfoObjectNOptional. Additional information for custom use. Refer to Appendix B for complete list of additionalInfo for Virtual Account

Note :

- totalAmount field value will be same with totalAmount field value returned from partner when BCA do inquiry list of bils.
- paymentRequestID corresponds to inquiryRequestID for each transaction (one pair of inquiry and payment)
- paidAmount field value will be total amount paid by customer through BCA.
- For multi bills transaction, paidAmount field value might be different with totalAmount field value. For non multi bills transaction, paidAmount field value will be same with totalAmount field value.
- currency must be same for totalAmount, billAmount, and paidAmount. currency field may vary with these values: IDR, SGD, USD.
- For multi bills and non multi bills transaction with multi settlements,
  - billSubCompany field is mandatory.
  - billSubCompany field value must be different for each bill number and billReferenceNo field value will be different for each bill number. referenceNo and subCompany field value will be blank for multi settlement transaction. Transaction would be settled according to each billSubCompany field value.
- flagAdvice field may vary with these values: N= new request for payment flag; Y= advice request (retry flag) for payment flag.
- billDetails should not be more than 5. BCA will returns only bills that customer choose to pay in payment flag request to partner.
- channelCode field may vary with these values :
  - 6010 = Teller (Financial institutions - manual cash disbursements)
  - 6011 = eBanking (Financial institutions - automated cash disbursements)

Result of the request will contains following information:

**Response:**

FieldData Type**Mandatory**DescriptionresponseCodeString (7)YMandatory In BCA. Response code from Partner, refer to Appendix A for list of response coderesponseMessageString (150)YMandatory in BCA. Response message from Partner, refer to Appendix A for list of response code virtualAccountDataObjectY   paymentFlagReasonObjectNMandatory in BCA. Reason for Payment Status multi language      englishString (200)NReason for inquiry status in English      indonesiaString (200)NReason for inquiry status in Bahasa    partnerServiceIdString (8)YMandatory in BCA. Derivative of XPARTNER-ID, similar to company code, 8 digit left padding space “ “   customerNoString (20)YMandatory in BCA. Unique number (up to 20 digit)   virtualAccountNoString(28)YMandatory in BCA. partnerServiceId (8 digit left padding space “ “) + customerNo (up to 20 digit)   virtualAccountNameString (255)YMandatory in BCA. Customer name   paymentRequestIdString (128)YMandatory in BCA. Payment request id from request payload   paidAmountObjectNMandatory in BCA.      valueString (16.2)

ISO 4217 YTransaction Amount. From Payment Request       currencyString (3)YCurrency    totalAmount ObjectNMandatory in BCA.      valueString (16.2)

ISO 4217 YTransaction Amount. From Payment Request      currencyString (3)YCurrency   trxDateTimeDate (25)NMandatory in BCA. From Payment Request   referenceNoString (15)NFrom Payment Request   paymentFlagStatusString (2)NMandatory in BCA. Status for Payment Flag from Partner   billDetailsArray of ObjectsNArray with maximum 24 Objects       billerReferenceIdString (64)NFrom Inquiry Response      billNoString (18)NFrom Inquiry Response      billDescriptionObjectNFrom Inquiry Response         englishString (18)NFrom Inquiry Response         indonesiaString (18)NFrom Inquiry Response      billSubCompanyString (5) NFrom Inquiry Response      billAmountObjectNFrom Inquiry Response         valueString (16.2) ISO 4217YTransaction Amount. From Payment Request         currencyString (3) YCurrency      additionalInfoObjectNAdditional Information for custom use in each bill. Refer to Appendix B for complete list of additionalInfo for Virtual Account   statusString (2)NPayment status for specific Bill   reasonObjectNReason for Payment Status for specific Bill multi language      englishString (64)NReason for Payment Status for specific Bill in English      indonesiaString (64)NReason for Payment Status for specific Bill in Bahasa freeTextsArray of Objects NOptional. Array with maximum 9 Objects   englishString (32)NWill be shown in Channel   indonesiaString (32)NWill be shown in ChanneladditionalInfoObjectNAdditional Information for custom use. Refer to Appendix B for complete list of additionalInfo for Virtual Account

Note :

\- The final status of payment flag response is not determined by responseCode and responseMessage.

\- Status for payment flag will be defined in paymentFlagStatus and paymentFlagReason field to be displayed to channel's screen. This field may vary with these values:

00 = success payment flag

01 = Reject payment flag by partner, partner must defined reason for this status in paymentFlagReason field values in two languages.

02 = Timeout payment flag between switcher to partner. If company's reconciliation type are reversal or force settle, transaction with status 02 will be considered as success transaction. (payment flag status other than 00,01,02 will be considered as 01)

\- BCA will continue to check the paymentFlagStatus and paymentFlagReason field if the responseCode received is 2002500 and 2022500.

\- If the response code is not either 2002500 or 2022500, BCA will consider the response as failed transaction and will be rejected.

\- For responseCode, paymentFlagStatus and status inside billdetail, this is condition for timeout :

responseCode**paymentFlagStatus**status inside billdetailFinal Status NNNTimeoutNYYTimeoutY (Success)NNTimeoutY (Failed)NNFailed

\- For multi bills and non multi bills transaction with multi settlement, billSubCompany field value must be different for each bill number. paymentFlagStatus field shows the final status for all bills. Status and reason defined in billDetails must be the same for all bills, and also contain the same status and reason as paymentFlagStatus and paymentFlagReason field.

\- If virtualAccountName length is greater than 30, then only first 30 character will be displayed on channel's screen

\- totalAmount field value will be same with totalAmount field value returned from partner when BCA do inquiry list of bills.

\- paymentRequestID is the same value with inquiryRequestID from BCA for each transaction.

\- paidAmount field value will be total amount paid by customer through BCA.

\- For multi bills transaction, paidAmount field value might be different with totalAmount field value. For non multi bills transaction, paidAmount field value will be same with totalAmount field value.

\- currency must be same for totalAmount, billAmount, and paidAmount. currency field may vary with these values: IDR, SGD, USD.

\- freeTexts field defined by partner and will be displayed in channel's screen. This field must be defined in two languages. The occurences of freeTexts field in payment flag should not be greater than 9

\- If a system error occurs and causing BCA to send double flagging request with the same X-EXTERNAL-ID and paymentRequestId, then partner can send responseCode

"4042518" and responseMessage "Inconsistent Request", with paymentFlagStatus and paymentFlagReason according to the results of the first request

\- If BCA sends a flagging request with the same X-EXTERNAL-ID but a different paymentRequestId, then partner can send responseCode "4092500" and responseMessage "Conflict".

\- The usage of additional info is to be discussed with BCA first.

Expected Response from Partner.

Scenario**responseCode**responseMessagepaymentFlagStatusAccess Token Invalid4012501"Invalid Token (B2B)"-Unauthorized . Signature4012500"Unauthorized. \[Signature\]"-Unauthorized . stringToSign4012500"Unauthorized. \[Signature\]"-Unauthorized . Unknown Client 4012500"Unauthorized. \[Unknown client\]"-Missing Mandatory Field {} …, etc4002502"Invalid Mandatory Field {........}" 01Invalid Field Format {} .., etc4002501"Invalid Field Format {........}" 01Cannot use the same X-EXTERNAL-ID4092500"Conflict" 01Input no Virtual Account Valid2002500"Success" 00Input no Virtual Account Valid sudah lunas4042514"Paid Bill" 01Duplicate XEXTERNAL-ID and paymentRequestId4042518"Inconsistent Request"00/01Input no Virtual Account Valid kadaluarsa4042519"Invalid Bill/Virtual Account"01Input no Virtual Account tidak terdaftar4042512"Invalid Bill/Virtual Account \[Reason\]" 01Request Parsing Error4002500"Bad Request" 01Response Parsing Error4002500"Bad Request" 01

curl –X POST https://copartners.com/openapi/v1.0/transfer-va/payment

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95231'
-H 'X-PARTNER-ID': '12345'
-H 'X-EXTERNAL-ID': '123456789012345678901234567890123456'

-d '
{
	"partnerServiceId": " 12345",
	"customerNo": "123456789012345678",
	"virtualAccountNo": " 12345123456789012345678",
	"virtualAccountName": "Jokul Doe",
	"paymentRequestId": "202202111031031234500001136962",
	"channelCode": 6011,
	"paidAmount": {
		"value": "100000.00",
		"currency": "IDR"
	},
	"referenceNo": "00113696201",
	"flagAdvise": "N",
	"subCompany": "00000",
	"billDetails": [
	{
		"billNo": "123456789012345678",
		"billDescription": {
			"english": "Maintenance",
			"indonesia": "Pemeliharaan"
		},
		"billSubCompany": "00000",
		"billAmount": {
			"value": "100000.00",
			"currency": "IDR"
		},
		"additionalInfo": {},
		"billReferenceNo": "00113696201"
	}
	],
	"additionalInfo": {}
}'

Sample for multi settlement (non multi bills/multi bills) transaction:

-H 'CHANNEL-ID':'95231'
-H 'X-PARTNER-ID': '12345'
-H 'X-EXTERNAL-ID': '123456789012345678901234567890123456'

-d '
{
	"partnerServiceId": " 12345",
	"customerNo": "123456789012345678",
	"virtualAccountNo": " 12345123456789012345678",
	"virtualAccountName": "Jokul Doe",
	"paymentRequestId": "202202111031031234500001136962",
	"channelCode": 6011,
	"hashedSourceAccountNo": "",
	"sourceBankCode": "014",
	"paidAmount": {
		"value": "100000.00",
		"currency": "IDR"
	},
	"cumulativePaymentAmount": null,
	"paidBills": "",
		"totalAmount": {
		"value": "100000.00",
		"currency": "IDR"
	},
	"trxDateTime": "2022-02-12T17:29:57+07:00",
	"referenceNo": "",
	"flagAdvise": "N",
	"subCompany": "",
	"billDetails": [
	{
		"billNo": "123456789012345678",
		"billDescription": {
			"english": "Maintenance",
			"indonesia": "Pemeliharaan"
		},
		"billSubCompany": "00000",
		"billAmount": {
			"value": "50000.00",
			"currency": "IDR"
		},
		"additionalInfo": {},
		"billReferenceNo": "00113696201"
	},
	{
		"billNo": "223456789012345678",
		"billDescription": {
			"english": "Electricity",
			"indonesia": "Listrik"
		},
		"billSubCompany": "00001",
		"billAmount": {
			"value": "20000.00",
			"currency": "IDR"
		},
		"additionalInfo": {},
		"billReferenceNo": "00213696201"
	},
	{
		"billNo": "323456789012345678",
		"billDescription": {
			"english": "Water",
			"indonesia": "Air"
		},
		"billSubCompany": "00002",
		"billAmount": {
			"value": "30000.00",
			"currency": "IDR"
		},
		"additionalInfo": {},
		"billReferenceNo": "00313696201"
	}
	],
	"additionalInfo": {}
}'
```

curl –X POST https://copartners.com/openapi/v1.0/transfer-va/payment

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
	"responseCode": "2002500",
	"responseMessage": "Successful",
	"virtualAccountData": {
		"paymentFlagReason": {
			"english": "Success",
			"indonesia": "Sukses"
		},
		"partnerServiceId": " 12345",
		"customerNo": "123456789012345678",
		"virtualAccountName": "Jokul Doe",
		"paymentRequestId": "202202111031031234500001136962",
		"paidAmount": {
			"value": "100000.00",
			"currency": "IDR"
		},
		"totalAmount": {
			"value": "100000.00",
			"currency": "IDR"
		},
		"trxDateTime": "2022-02-12T17:29:57+07:00",
		"referenceNo": "00113696201",
		"paymentFlagStatus": "00",
		"billDetails": [
		{
			"billerReferenceId": "00113696201",
			"billNo": "123456789012345678",
			"billDescription": {
				"english": "Maintenance",
				"indonesia": "Pemeliharaan"
			},
			"billSubCompany": "00000",
			"billAmount": {
				"value": "100000.00",
				"currency": "IDR"
			},
			"additionalInfo": {},
			"status": "00",
			"reason": {
				"english": "Success",
				"indonesia": "Sukses"
			}
		}
		],
	}
	"additionalInfo":
	{
		"additionalInfo1":
		{
			"label":
			{
				"indonesia": "Unit",
				"english": "Unit"
			},
			"value":
			{
				"indonesia": "10C",
				"english": "10C"
			}
		},
		"additionalInfo2":
		{
			"label":
			{
				"indonesia": "Bulan",
				"english": "Month"
			},
			"value":
			{
				"indonesia": "Januari",
				"english": "January"
			}
		}
	}
}

Sample for multi settlement (non multi bills/multi bills) transaction:

{
	"responseCode": "2002500",
	"responseMessage": "Successful",
	"virtualAccountData": {
		"paymentFlagReason": {
			"english": "Success",
			"indonesia": "Sukses"
		},
		"partnerServiceId": " 12345",
		"virtualAccountNo": " 12345123456789012345678",
		"virtualAccountName": "Jokul Doe",
		"paymentRequestId": "202202111031031234500001136962",
		"paidAmount": {
			"value": "100000.00",
			"currency": "IDR"
		},
		"totalAmount": {
			"value": "100000.00",
			"currency": "IDR"
		},
		"trxDateTime": "2022-02-12T17:29:57+07:00",
		"referenceNo": "",
		"paymentFlagStatus": "00",
		"billDetails": [
		{
			"billerReferenceId": "00113696201",
			"billNo": "123456789012345678",
			"billDescription": {
				"english": "Maintenance",
				"indonesia": "Pemeliharaan"
			},
			"billSubCompany": "00000",
			"billAmount": {
				"value": "50000.00",
				"currency": "IDR"
			},
			"additionalInfo": {},
			"status": "00",
			"reason": {
				"english": "Success",
				"indonesia": "Sukses"
			}
		},
		{
			"billerReferenceId": "00213696201",
			"billNo": "223456789012345678",
			"billDescription": {
				"english": "Electricity",
				"indonesia": "Listrik"
			},
			"billSubCompany": "00001",
			"billAmount": {
				"value": "20000.00",
				"currency": "IDR"
			},
			"additionalInfo": {},
			"status": "00",
			"reason": {
				"english": "Success",
				"indonesia": "Sukses"
			}
		},
		{
			"billerReferenceId": "00313696201",
			"billNo": "323456789012345678",
			"billDescription": {
				"english": "Water",
				"indonesia": "Air"
			},
			"billSubCompany": "00002",
			"billAmount": {
				"value": "30000.00",
				"currency": "IDR"
			},
			"additionalInfo": {},
			"status": "00",
			"reason": {
				"english": "Success",
				"indonesia": "Sukses"
			}
		}
		],
	}
	"additionalInfo":
	{
		"additionalInfo1":
		{
			"label":
			{
				"indonesia": "Unit",
				"english": "Unit"
			},
			"value":
			{
				"indonesia": "10C",
				"english": "10C"
			}
		},
		"additionalInfo2":
		{
			"label":
			{
				"indonesia": "Bulan",
				"english": "Month"
			},
			"value":
			{
				"indonesia": "Januari",
				"english": "January"
			}
		}
	}
}
```

#### Transfer ke Rekening Virtual Account BCA

##### 1\. SNAP Virtual Account Inquiry Payment to VA from Intrabank

This service is used to VA transfer BillPresentment.

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)Y Channel’ Identifier using WSID KlikBCA Bisnis (95051)X-PARTNER-IDHeaderString(32)Y Partner ID using Corporate ID KlikBCA BisnisX-EXTERNAL-IDHeaderString (32)YNumeric String.

Reference number that should be unique in the same day

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionvirtualAccountNoString (28)MaxYAlphanumericMandatory in BCA.

partnerServiceId (8 digit left padding space) + customerNo (up to 20 digit)

In BCA, virtualAccountNo max length is 26 (partnerServiceId(8 digits)+customerNo(18 digits)).

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)FixedY-Response coderesponseMessageString (150)MaxY-Response message descriptionvirtualAccountDataObjectY   virtualAccountNoString (28)MaxYAlphanumericpartnerServiceId (8 digit left padding space) + customerNo (up to 20 digits)   virtualAccountNameString (255)MaxYAlphanumericCustomer name   totalAmountArray of ObjectsN      valueString (16.2)MaxYSend by BCA.

Total amount with 2 decimal

In BCA, totalAmount max length is 13.2 digits.      currencyString (3)FixedYISO-4217Send by BCA.

Currency of amount based on ISO 4217    billDetailsArray of ObjectsNArray with maximum 24 Objects       billDescriptionObjectNSend by BCA.

Bill Description         englishString (18)MaxNAlphanumericSend by BCA.

Bill Description in English         indonesiaString (18)MaxNAlphanumericSend by BCA

Bill Description in Bahasa      billAmountArray of ObjectsN         valueString (16.2) NSend by BCA.

Transaction amount. Nominal inputted by Customer with 2 decimal

In BCA, billAmount max length is 13.2 digits.          currencyString (3) MaxYISO-4217Currency      additionalInfoString MaxNAlphanumeric Additional information for custom use for each bill   freeTextsStringNSend by BCA.

Array with maximum 25 objects      englishString (32)MaxNAlphanumericWill be shown in Channel       indonesiaString (32)MaxNAlphanumericWill be shown in Channel    virtualAccountTrxTypeString (1)FixedNType of Virtual Account that send by BCA is the highlighted transaction type.

1\. Closed Payment (C): Tagihan muncul, harus dibayar sesuai tagihan (Fixed Bill)

2\. Open Payment (O): Tagihan tidak muncul (No Bill)

3\. Partial (I): Sisa tagihan

4\. Minimum (M): Dibayar hanya sekali dengan nominal minimum namun tidak boleh lebih kecil dari tagihan.

5\. Maximum (L): Dibayar hanya sekali, tidak boleh lebih besar dari tagihan, dan tagihan muncul

6\. Open Minimum (N): Dibayar berkali2 dengan nominal minimum namun tidak boleh lebih kecil dari tagihan dan muncul tagihannya.

7\. Open Maximum (X): Dibayar berkali2 dengan nominal maksimum (cumulative) namun tidak boleh lebih besar dari tagihan dan muncul tagihannya.

8\. Bill Variable (V): Bisa dibayar sekail, bisa lebih besar / lebih kecil dari tagihan, dan tagihan muncul.

9\. Multi Bill Variable (W)   feeAmountArray of ObjectN      valueString (16.2)MaxYSend by BCA.

Nominal inputted by customer with 2 decimal.

In BCA, totalAmount max length is 13.2 digits.       currencyString (3)FixedYISO-4217Currency   productNameString (30) MaxNAlphanumericSend by BCA.

Product category

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message4004003200Bad request4004003201Invalid Field Format {field name} 4004003202Invalid Mandatory Field {field name}4014013200Invalid token (B2B)4014013201Unauthorized. \[Reason\] 4034033201Feature Not Allowed4044043212Invalid Bill4044043214Paid Bills4094093211Conflict4294293200Too Many Requests 5005003200General Error5005003200Internal Server Error5045043200Timeout

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/transfer-va/inquiry-intrabank

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95051'
-H 'X-PARTNER-ID': '1234567890'
-H 'X-EXTERNAL-ID': '41807553358950093184162180797837'
-d '
{
	"virtualAccountNo" : " 0008899912345678901234567890",
	"amount" : {
		"value" : "50000.00",
		"currency" : "IDR"
	}
}'
```

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/transfer-va/inquiry-intrabank

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
	"responseCode": "2003200",
	"responseMessage": "Successful",
	"virtualAccountData": {
		"virtualAccountNo": "0008889912345678901234567890",
		"virtualAccountName": "Jokul Doe",
		"totalAmount": {
			"value" : "88000.00",
			"currency" : "IDR"
		},
		"billDetails": [
		{
			"billDescription": {
				"english": "Maintenance",
				"indonesian": "Pemeliharaan"
			},
			"billAmount": {
				"value" : "88000.00",
				"currency" : "IDR"
			}
		}
		],
		freeTexts": [
		{
			"english": "Free text",
			"indonesia": "Tulisan bebas"
		}
		],
		"virtualAccountTrxType": "C",
		"feeAmount": {
			"value" : "1000.00",
			"currency" : "IDR"
		},
		"productName": "OVO"
	}
}

```

##### 2\. SNAP Virtual Account Notification Payment to VA from Intrabank

This service is used to Notification VA transfer.

**Additional Headers:**

Field**Params Type**Data Type**Mandatory****Format**DescriptionCHANNEL-IDHeaderString(5)Y 95051Channel’ Identifier using WSID KlikBCA Bisnis (95051)X-PARTNER-IDHeaderString(32)Y -Partner client id using KlikBCA Bisnis’s Corporate IDX-EXTERNAL-IDHeaderString(32)YAlphanumericNumeric String. Reference number that should be unique in the same day

**Payload:**

FieldData Type**Mandatory****Format**DescriptionvirtualAccountNoString (28)YAlphanumericSend by BCA. partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits)

In BCA, virtualAccountNo max length is 26 (partnerServiceId (8 digits) + customerNo (18 digits)).partnerReferenceNoString (128)Y-Send by BCA.

Unique identifier for this Payment.

Generated by Partner. In BCA, partnerReferenceNo max length is 64 digits.trxDateTimeDate (25)N-Send by BCA.

PJP internal system datetime with timezone, which follows the ISO-8601 standard paymentStatusString (20)NAlphanumericSend by BCA.

Status of payment requestpaymentFlagReasonObjectsN-Reason for Payment Status multi language    englishString (200)NAlphanumericSend by BCA.

Reason for Payment Status in English   indonesiaString (200)NAlphanumericSend by BCA.

Reason for Payment Status in Bahasa

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)FixedY-Response coderesponseMessageString (150)MaxY-Response descriptionvirtualAccountDataObjectY   virtualAccountNoString (28)MaxYAlphanumericMandatory in BCA according to request input.

partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits).

In BCA, virtualAccountNo max length is 26 (partnerServiceId (8 digits) + customerNo (18 digits)).   partnerReferenceNoString (128)MaxY-Mandatory in BCA according to request input.

From Payment request

The following structure should be returned only when not-success with additional related HTTP status of the response if applicable.

The copartners define the responseCode by themselves. For the multi-languageresponseMessage, it allowed to use same value if multi-language are not possible

{

"responseCode" : "Error Code",

"responseMessage" : "Error Message",

"virtualAccountData":

{

"virtualAccountNo": "",

"partnerReferenceNo": ""

}

}

curl –X POST https://copartners.com/openapi/v1.0/transfer-va/notify-payment-intrabank

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95051'
-H 'X-PARTNER-ID': 'KBCABCINDO'
-H 'X-EXTERNAL-ID': '41807553358950093184162180797837'
-d '
{
	"virtualAccountNo": " 8889912345678901234567890",
	"partnerReferenceNo": "12345678",
	"trxDateTime": "2020-12-21T10:30:24+07:00",
	"paymentStatus": "Success",
	"paymentFlagReason": {
		"english": "Success",
		"indonesia": "Sukses"
	}
}'
```

curl –X POST https://copartners.com/openapi/v1.0/transfer-va/notify-payment-intrabank

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
	"responseCode": 2003400,
	"responseMessage": "Successful",
	"virtualAccountData": {
		"virtualAccountNo": " 8889912345678901234567890",
		"partnerReferenceNo": "12345678"
	}
}

```

##### 3\. SNAP Virtual Account Payment to VA from Intrabank

This service is used to VA transfer.

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory****Format**DescriptionCHANNEL-IDHeaderString(5)FixedY 95051Channel’ Identifier using WSID KlikBCA Bisnis (95051)X-PARTNER-IDHeaderString(32)MaxY Partner client id using KlikBCA Bisnis’s WSID (95051)

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionvirtualAccountNoString (28)MaxYAlphanumericMandatory in BCA.

partnerServiceId (8 digit left padding space) + customerNo (up to 20 digit)

In BCA, virtualAccountNo max length is 26 (partnerServiceId(8 digits)+customerNo(18 digits)).virtualAccountEmailString (255)MaxN-Optional BCA

Customer email with max length is 255 digits.sourceAccountNoString (32)MaxNNumericSource account numberpartnerReferenceNoString (128)MaxY-Mandatory in BCA.

Unique identifier for this Payment. Generated by PJPpaidAmountObjectY   valueString (16.2)MaxYMandatory in BCA.

Paid amount with 2 decimal In BCA, partnerReferen ceNumber max length is 13.2 digits   currencyString (3)FixedYISO-4217Mandatory in BCA.

Currency of amount based on ISO 4217trxDateTimeDate (25)FixedN-Mandatory in BCA.

PJP internal system datetime with timezone, which follows the ISO-8601 standard

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)FixedY-Response coderesponseMessageString (150)MaxY-Response message descriptionvirtualAccountDataObjectY   virtualAccountNoString (28)MaxYAlphanumericpartnerServiceId (8 digit left padding space) + customerNo (up to 20 digits)   virtualAccountNameString (255)MaxYAlphanumericCustomer name   virtualAccountEmailString (255)MaxNCustomer email   sourceAccountNoString (32)MaxN-Source account number   partnerReferenceNoString (128)MaxY-From Payment request   referenceNoString (128)MaxNNumericGenerated by PJP   paidAmountObjectN-      valueString (16.2)MaxYSend by BCA according to the request input.

From Payment request with max length (13.2).       currencyString (3)FixedYISO-4217Send by BCA according to the request input.

Currency from Payment request.paidAmountObjectNvalueString (16.2)MaxYSend by BCA according to the request input.

From Payment request with max length (13.2).currencyString (3)FixedYISO-4217Send by BCA.

Currency   trxDateTimeDate (25)MaxNFrom Payment Request    billDetailsArray of ObjectsNArray with maximum 24 Objects       billDescriptionJSON ObjectN
Bill Description         englishString (18)MaxNAlphanumeric
Bill Description in English         indonesiaString (18)MaxNAlphanumeric
Bill Description in Bahasa      billAmountArray of ObjectsN         valueString (16.2) NSend by BCA.

Nominal inputted by Customer with 2 decimal

In BCA, billAmount max length is (13.2) digits.         currencyString (3) MaxYISO-4217Currency   freeTextsArray of ObjectsN-From Inquiry response      englishString (32)MaxNAlphanumericFrom Inquiry response      indonesiaString (32)MaxNAlphanumericFrom Inquiry response   feeAmountObjectN      valueString (16.2)MaxYSend by BCA.

Nominal inputted by customer with 2 decimal.

In BCA, totalAmount max length is 13.2 digits.       currencyString (3)FixedYISO-4217Send by BCA.

Currency   productNameString (30) MaxNAlphanumericSend by BCA.

Product category

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message4004003300Bad request4004003301Invalid Field Format {field name} 4004003302Invalid Mandatory Field {field name}4014013300Unauthorized. \[Reason\]4014013301Invalid token (B2B)4034033301Feature Not Allowed4034033302Exceeds Transaction Amount Limit4034033304Activity Count Limit Exceeded4044043312Invalid Bill4044043314Paid Bills4094093311Conflict4294293300Too Many Requests 5005003301Internal Server Error5005003300General Error5045043300Timeout

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/transfer-va/payment-intrabank

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95051'
-H 'X-PARTNER-ID': 'KBCABCINDO'
-H 'X-EXTERNAL-ID': '41807553358950093184162180797837'
-d '
{
	"virtualAccountNo": " 00088999123456789012345678900008899912345678901234567890",
	"virtualAccountEmail": "john@email.com",
	"sourceAccountNo": "1234567890",
	"partnerReferenceNo": "12345678",
	"paidAmount": {
		"value": "50000.00",
		"currency": "IDR"
	}
	 "trxDateTime": "2020-12-21T10:30:24+07:00"
}'
```

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/transfer-va/payment-intrabank

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
	"responseCode": 2023300,
	"responseMessage": "Request In Progress",
	"virtualAccountData":
	{
		"virtualAccountNo": "0008889912345678901234567890",
		"virtualAccountName": "Jokul Doe",
		"virtualAccountEmail": "john@email.com",
		"sourceAccountNo": "1234567890",
		"partnerReferenceNo": "12345678",
		"referenceNo": "12345678901234",
		"paidAmount":
		{
			"value": "55000.00",
			"currency": "IDR"
		},
		"totalAmount":
		{
			"value": "55000.00",
			"currency": "IDR"
		},
		"trxDateTime": "2020-12-21T10:30:24+07:00",
		"billDetails": [
		{
			"billDescription":
			{
				"english": "Maintenance",
				"indonesia": "Pemeliharaan"
			},
			"billAmount": [
			{
				"value": "55000.00",
				"currency": "IDR"
			}]
		}],
		"freeTexts": [
		{
			"english": "Free text",
			"Indonesia": "Tulisan bebas"
		}
		],
		"feeAmount":
		{
			"value": "5000.00",
			"currency": "IDR"
		},
		"productName": "OVO"
	}
}

```

#### ECR EDC

##### 1\. ECR EDC BCA Inquiry Status

With ECR EDC BCA Inquiry Status, users can record and keep users' payment status when doing transaction with EDC BCA.

These are things users can do with ECR EDC BCA Inquiry Status:

• Inquiry Purchase / Sale / Void / Settlement

• Inquiry QRIS Payments

ECR EDC BCA Inquiry Status is used to check the current transaction on EDC BCA

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString (5) FixedY Channel’s identifier using EDC’s WSID (95251)X-PARTNER-IDHeaderString (36) MaxY Unique ID for partner (Group Code/MID)X-EXTERNAL-ID

HeaderString (36) MaxYNumeric String.

Transaction ID that should be unique per transaction and generated by merchant

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionmerchantIdString (9)MaxYAlphanumeric MID merchantterminalIdString (8)MaxYAlphanumeric Terminal IdentificationtransTypeString (2)Fixed YNumeric Transaction Type that are requested:

\- 42 Inquiry Status Purchase / Sale / Void / Settlement.

\- 32 Inquiry QRIS.partnerTypeString (2)FixedNAlphanumericFor general/special cooperation indicators. will be used later

originalPartnerReferenceNoString(64)Max YAlphanumericoriginalPartnerReferenceNo is a unique transaction ID generated by the merchant when Request Payment.

Result of the request will contains following information:

**Response:** **Inquiry Status Purchase / Sale / Void / Settlement / QRIS (Pending)**

FieldDataType**Length Type****Mandatory****Format**DescriptionresponseCode

String (7) FixedYAlphanumeric

Response code

responseMessage

String (150)

MaxYAlphanumericResponse description

merchantIdString (9)MaxY AlphanumericMID MerchantterminalIdString (8)MaxYAlphanumericTerminal IdentificationtransTypeString (2)FixedYNumericTransaction Type that are requested:

\- 42 Inquiry Status Purchase / Sale / Void / Settlement.

\- 32 Inquiry QRIS Payment.

originalPartnerReferenceNoString (64)MaxYAlphanumericoriginalPartnerReferenceNo is a unique transaction ID generated by the merchant when Request PaymentcreatedDateString (25) Fixed Y Date Format(YYYY-MM-DDTHH:MM:SS+07:00)Time Send Response/ Time create DB if successful

**Inquiry Status Purchase / Sale / Void (Success)**

FieldDataType**Length Type****Mandatory****Format**DescriptionresponseCode

String (7) FixedYAlphanumeric

Response code

responseMessage

String (150)

MaxYAlphanumericResponse description

midTransString (15)MaxYAlphanumericMID Transaction MerchanttidTransString (8)MaxYAlphanumericTerminal Id TransactiontransTypeString (2)FixedYNumericTransaction Type that are requested:

\- 42 Inquiry Status Purchase / Sale / Void / Settlement

originalPartnerReferenceNoString (64)MaxYAlphanumericoriginalPartnerReferenceNo is a unique transaction ID generated by the merchant when Request PaymentcreatedDateString (25) Fixed Y Date Format(YYYY-MM-DDTHH:MM:SS+07:00)Time Send Response/ Time create DB if successfultransDateString (25)FixedYDate Format(YYYY-MM-DDTHH:MM:SS+07:00)Get from BCA’s host during transaction

additionalInfoObject    batchNumberString  (6)FixedYNumericBatch transaction   traceNo

String (16)MaxYNumericNumber for tracking to destination bank   referenceNo

String (64)MaxYNumericReference number transaction   approvalCodeString (6)FixedYNumericWhen an input is made and the EDC prints the receipt but the status is still pending, the host changes the status to successful and replies to the merchant

issuerId

String (2)FixedYNumericIssuer ID   cardType

String (255)MaxYAlphanumericType of Card   cardholderName String (150)MaxYAlphanumericCardholder Name on Card Number   amount1

String (9.2)MaxYNumericValue amount transaction   amount2

String (9.2)

Max

N
Numeric

Value tips / tunai

**Inquiry Status Settlement**

FieldDataType**Length Type****Mandatory****Format**DescriptionresponseCode

String (7) FixedYAlphanumeric

Response code

responseMessage

String (150)

MaxYAlphanumericResponse description

merchantIdString (9)MaxYAlphanumericMID MerchantterminalIdString (8)MaxYNumericTerminal IdentificationtransTypeString (2)FixedYNumericTransaction Type that are requested:

\- 42 Inquiry Status Purchase / Sale / Void / Settlement

originalPartnerReferenceNoString (64)MaxYAlphanumericoriginalPartnerReferenceNo is a unique transaction ID generated by the merchant when Request PaymentcreatedDateString (25) Fixed Y Date Format(YYYY-MM-DDTHH:MM:SS+07:00)Time Send Response/ Time create DB if successfultransDateString (25)FixedYDate Format(YYYY-MM-DDTHH:MM:SS+07:00)Get from BCA’s host during transaction

additionalInfoArray of Object   tidTransString  (8)MaxYAlphanumericTerminal Id Transaction   midTrans

String (15)MaxYAlphanumericMID Transaction Merchant   settleTrans

String (10)MaxYAlphanumericDEBIT, QR, CREDIT   batchNumberString (6)FixedYNumericTransaction batch number recorded in BCA's host   totalCount

String (3)MaxYNumericValue count settlement   totalAmount

String (10.2)MaxYNumericValue amount settlement

**Inquiry QRIS**

FieldDataType**Length Type****Mandatory****Format**DescriptionresponseCode

String (7) FixedYAlphanumeric

Response code

responseMessage

String (150)

MaxYAlphanumericResponse description

midTransString (15)MaxY AlphanumericMID Transaction MerchanttidTransString (8)MaxYAlphanumericTerminal Id TransactiontransTypeString (2)FixedYNumericTransaction Type that are requested:

\- 32 Inquiry QRIS.

originalPartnerReferenceNoString (64)MaxYAlphanumericoriginalPartnerReferenceNo is a unique transaction ID generated by the merchant when Request PaymentcreatedDateString (25) Fixed Y Date Format(YYYY-MM-DDTHH:MM:SS+07:00)Time Send Response/ Time create DB if successfultransDateString (25)FixedYDate Format(YYYY-MM-DDTHH:MM:SS+07:00)

Get from BCA’s host during transaction

additionalInfoObject     batchNumberString (6)FixedYNumericBatch Transaction     referenceNo String (64)MaxYNumericReference number transaction     rrnString (12)MaxYAlphanumericQRIS ref number obtained from the issuer at BCA's host     traceNoString (16)MaxYNumericNumber of tracking to destination bank     approvalCode

String (6)FixedYNumericWhen an input is made and the EDC prints the receipt but the status is still pending, the host changes the status to successful and replies to the merchant     issuerId

String (2)FixedYNumericIssuer ID     issuerName String (12)MaxYAlphanumericIssuer Name     customerPan

String (19)FixedYNumericBuyer card number (customer\_pan), will be used to calculate transaction limit. Mandatory if payment requires limit validation (i.e. payment from mBCA)

     merchantPan

String (19)FixedYNumericMerchant Number. example: 9360001413243523431     personalNumber

String (16)

MaxYNumericWill be filled if BCA's host receive buyer phone number. Format: first 4 digits and last 2 digits are not masked, else masked     cardholderName

String (150)MaxYAlphanumericCardholder Name on Card Number     amount1 String (9.2)Max

Y

Numeric

Value amount transaction

     amount2

String (9.2)

Max

NNumeric

Value amount transaction

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message**Condition**202202IS00Request In ProgressTransaction request in progress400400IS00Bad requestGeneral request failed error400400IS01Invalid Field Format {field}The data you entered does not match the format requirements400400IS02Invalid Mandatory Field {field}Mandatory field should be fulfilled401401IS00Unauthorized. \[Reason\]Invalid signature/Unknown client/Connection not allowed401401IS01Invalid token (B2B)Access Token Not Exist/Access Token Expired/Token is Invalid403403IS01Feature Not AllowedThe feature is not allowed to use. Please make sure you registered this feature in your partnership data.403403IS05Do Not Honor Transaction Error From Host403403IS15Transaction Not Permitted.

transType Not FoundtransType not available in the list404404IS01Transaction Not FoundoriginalPartnerReferenceNo Not Found404404IS08Invalid MerchantInvalid Merchant409409IS00ConflictX-EXTERNAL-ID duplicate429429IS00Too Many RequestsMaximum transaction limit exceeded500500IS00General ErrorWrong request field input500500IS01Internal Server ErrorUnknown internal server failure, please retry the process again504504IS00Timeout

Your transaction timeout

POST /openapi/v1.0/ecr/inquiry-status

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
curl –X POST https://sandbox.bca.co.id/openapi/v1.0/ecr/inquiry-status
-H 'Content-type: application/json'
-H 'Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-TIMESTAMP: 2024-03-25T13:50:04+07:00'
-H 'X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'ORIGIN: www.hostname.com'
-H 'CHANNEL-ID: 95251'
-H 'X-PARTNER-ID: 82150823919040774621823174737537'
-H 'X-EXTERNAL-ID: 41807553358950093184162180797837'
-d '
{
 "merchantId": "000002872",
 "terminalId": "C2019345",
 "transType": "42",
 "partnerType": "",
 "originalPartnerReferenceNo": "41807553358950093145162180797837"
}'
```

POST /openapi/v1.0/ecr/inquiry-status

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
Inquiry Status Purchase / Sale / Void / Settlement / QRIS (Pending)
-H 'Content-Type: application/json'
-H 'X-TIMESTAMP: 2024-03-25T13:50:06+07:00'
{
 "responseCode": "202IS00",
 "responseMessage": "Request In Progress",
 "merchantId": "123455622",
 "terminalId": "C2825423",
 "transType": "42",
 "originalPartnerReferenceNo": "41807553358950093184162180797837",
 "createdDate": "2024-03-25T14:00:00+07:00"
}

Inquiry Status Purchase / Sale / Void (Success)
-H 'Content-Type: application/json'
-H 'X-TIMESTAMP: 2024-03-25T13:50:06+07:00'
{
 "responseCode": "200IS00",
 "responseMessage": "Successful",
 "midTrans": "000885123455622",
 "tidTrans": "D2825423",
 "transType": "42",
 "originalPartnerReferenceNo": "41807553358950093184162180797837",
 "createdDate": "2020-12-23T08:43:11+07:00",
 "transDate": "2024-03-25T14:00:08+07:00",
 "additionalInfo": {
 "batchNumber": "000004",
 "traceNo": "534728342",
 "referenceNo": "63813282934192893",
 "approvalCode": "256473",
 "issuerId": "13",
 "cardType": "Debit BCA",
 "cardholderName": "Felicia Miranda",
 "amount1": "200000000.00",
 "amount2": "100000000.00"
 }
}

Inquiry Status Settlement
-H 'Content-Type: application/json'
-H 'X-TIMESTAMP: 2024-03-25T13:50:06+07:00'
{
 "responseCode": "200IS00",
 "responseMessage": "Successful",
  "merchantId": "123455622",
 "terminalId": "C2825423",
 "transType": "42",
 "originalPartnerReferenceNo": "41807553358950093184162180797837",
 "createdDate": "2020-12-23T08:43:11+07:00",
 "transDate": "2024-03-25T14:00:08+07:00",
 "additionalInfo": [
 {
 "tidTrans": "D2825423",
 "midTrans": "000885123455622",
 "settleTrans":"DEBIT",
 "batchNumber": "000004",
 "totalCount": "4",
 "totalAmount": "1000000000.00"
 },
 {
 "tidTrans": "A2825423",
 "midTrans": "000885123455622",
 "settleTrans":"QRIS",
 "batchNumber": "000003",
 "totalCount": "8",
 "totalAmount": "1000000000.00"
 }
 ]
}

Inquiry Status QRIS
-H 'Content-Type: application/json'
-H 'X-TIMESTAMP: 2024-03-25T13:50:06+07:00'
{
 "responseCode": "200IS00",
 "responseMessage": "Successful",
 "midTrans": "000885123455622",
 "tidTrans": "A2825423",
 "transType": "32",
 "originalPartnerReferenceNo": "41807553358950093184162180797837",
 "createdDate": "2020-12-23T08:43:11+07:00",
 "transDate": "2024-03-25T14:00:08+07:00",
 "additionalInfo": {
 "batchNumber": "000004",
 "referenceNo": "63813282934192893",
 "rrn": "9F23AD582M",
 "traceNo": "534728342",
 "approvalCode": "256473",
 "issuerId": "00",
 "issuerName": "BCA",
 "customerPan": "9360001413243523432",
 "merchantPan": "93600014323452342",
 "personalNumber": "0851******92",
 "carholderName": "Felicia Miranda",
 "amount1": "200000000.00",
 "amount2": "100000000.00"
 }
}

Error Sample
H 'Content-type: application/json'
-H 'X-TIMESTAMP: 2024-03-25T13:50:04+07:00'
{
 "responseCode": "403IS05",
 "responseMessage": "Do Not Honor",
 "originalPartnerReferenceNo": "21132312100000001"
}
```

##### 2\. ECR EDC BCA Request Payment

With ECR EDC BCA Request Payment, users can request users' payment method through EDC BCA

ECR EDC BCA Request Payment can only request payment to EDC BCA for the transaction below:

• Purchase

• Sale

• QRIS

•Void

• Settlement

To use ECR EDC BCA Request Payment, please generate new transaction id because transaction id cannot be duplicate.

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString (5) FixedY Channel’s identifier using EDC’s WSID (95251)X-PARTNER-IDHeaderString (36) MaxY Unique ID for partner (Group Code/MID)X-EXTERNAL-ID

HeaderString (36) MaxYNumeric String.

Transaction ID that should be unique per transaction and generated by merchant

**Payload:**

**Request Purchase**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNo String (64)MaxYAlphanumericTransaction identifier on service consumer system.

Format: include character “–“.

Example:

1\. EA-012432342asea30 for partnerReferenceNo that use character "-"

2\. PEG3482093840820091 for partnerReferenceNo that don't use character "-" merchantIdString (9)MaxYAlphanumeric MID merchantterminalIdString (8)MaxYAlphanumeric Terminal IdentificationtransTypeString (2)Fixed YNumeric Transaction Type that are requested:

\- 01 = Purchase

\- 02 = Sale (Debit/Credit)

\- 31 = QRIS

\- 08 = Void

\- 10 = SettlementpartnerTypeString (2)FixedNAlphanumericFor general/special cooperation indicators. will be used later

additionalInfoObjectCMandatory for transType 01,02,08,31    amount1String (9.2)MaxYNumericValue amount transaction.

Format: 9 digit of transaction nominal, 2 digit for (.00)

**Request Payment Sale (Debit/Kredit)**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNo String (64)MaxYAlphanumericTransaction identifier on service consumer system.

Format: include character “–“.

Example:

1\. EA-012432342asea30 for partnerReferenceNo that use character "-"

2\. PEG3482093840820091 for partnerReferenceNo that don't use character "-" merchantIdString (9)MaxYAlphanumeric MID merchantterminalIdString (8)MaxYAlphanumeric Terminal IdentificationtransTypeString (2)Fixed YNumeric Transaction Type that are requested:

\- 01 = Purchase

\- 02 = Sale (Debit/Credit)

\- 31 = QRIS

\- 08 = Void

\- 10 = SettlementpartnerTypeString (2)FixedNAlphanumericFor general/special cooperation indicators. will be used later

additionalInfoObjectCMandatory for transType 01,02,08,31    amount1

String (9.2)MaxYNumeric Value Amount Transaction (Amount Sale + amount 2).

Format: 9 digit of transaction nominal, 2 digit for (.00)

    amount2

String (9.2)

MaxN NumericValue Amount Transaction (Amount Sale + amount 2).

Format: 9 digit of transaction nominal, 2 digit for (.00)

**Request Payment QRIS**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNo String (64)MaxYAlphanumericTransaction identifier on service consumer system.

Format: include character “–“.

Example:

1\. EA-012432342asea30 for partnerReferenceNo that use character "-"

2\. PEG3482093840820091 for partnerReferenceNo that don't use character "-" merchantIdString (9)MaxYAlphanumeric MID merchantterminalIdString (8)MaxYAlphanumeric Terminal IdentificationtransTypeString (2)Fixed YNumeric Transaction Type that are requested:

\- 01 = Purchase

\- 02 = Sale (Debit/Credit)

\- 31 = QRIS

\- 08 = Void

\- 10 = SettlementpartnerTypeString (2)FixedNAlphanumericFor general/special cooperation indicators. will be used later

additionalInfoObjectCMandatory for transType 01,02,08,31    amount1

String (9.2)MaxYNumeric Value Amount Transaction (Amount Sale + amount 2).

Format: 9 digit of transaction nominal, 2 digit for (.00)

    amount2

String (9.2)

MaxN NumericValue Amount Transaction (Amount Sale + amount 2).

Format: 9 digit of transaction nominal, 2 digit for (.00)

**Request Payment Void**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNo String (64)MaxYAlphanumericTransaction identifier on service consumer system.

Format: include character “–“.

Example:

1\. EA-012432342asea30 for partnerReferenceNo that use character "-"

2\. PEG3482093840820091 for partnerReferenceNo that don't use character "-" merchantIdString (9)MaxYAlphanumeric MID merchantterminalIdString (8)MaxYAlphanumeric Terminal IdentificationtransTypeString (2)Fixed YNumeric Transaction Type that are requested:

\- 01 = Purchase

\- 02 = Sale (Debit/Credit)

\- 31 = QRIS

\- 08 = Void

\- 10 = SettlementpartnerTypeString (2)FixedNAlphanumericFor general/special cooperation indicators. will be used later

additionalInfoObjectCMandatory for transType 01,02,08,31originalPartnerReferenceNo

String (64)MaxYAlphanumeric originalPartnerReferenceNo is a unique transaction ID generated by the merchant when Request Payment


**Request Payment Settlement**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNo String (64)MaxYAlphanumericTransaction identifier on service consumer system.

Format: include character “–“.

Example:

1\. EA-012432342asea30 for partnerReferenceNo that use character "-"

2\. PEG3482093840820091 for partnerReferenceNo that don't use character "-" merchantIdString (9)MaxYAlphanumeric MID merchantterminalIdString (8)MaxYAlphanumeric Terminal IdentificationtransTypeString (2)Fixed YNumeric Transaction Type that are requested:

\- 01 = Purchase

\- 02 = Sale (Debit/Credit)

\- 31 = QRIS

\- 08 = Void

\- 10 = SettlementpartnerTypeString (2)FixedNAlphanumericFor general/special cooperation indicators. will be used later

Result of the request will contains following information:

**Response:**

FieldDataType**Length Type****Mandatory****Format**DescriptionresponseCode

String (7) FixedYAlphanumeric

Response code

responseMessage

String (150)

MaxYAlphanumericResponse description

merchantIdString (9)MaxY AlphanumericMID MerchantterminalIdString (8)MaxYAlphanumericTerminal IdentificationtransTypeString (2)FixedYNumericTransaction Type that are requested:

\- 01 = Purchase

\- 02 = Sale (Debit/Credit)

\- 31 = QRIS

\- 08 = Void

\- 10 = Settlement

PartnerReferenceNoString (64)MaxYAlphanumericPartnerReferenceNo is a unique transaction ID generated by the merchant when Request PaymentcreatedDateString (25) Fixed Y Date Format(YYYY-MM-DDTHH:MM:SS+07:00)Time Send Response/ Time create DB if successful

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message**Condition**400400RP00Bad requestGeneral request failed error400400RP01Invalid Field Format {field}The data you entered does not match the format requirements400400RP02Invalid Mandatory Field {field}Mandatory field should be fulfilled401401RP00Unauthorized. \[Reason\]Invalid signature/Unknown client/Connection not allowed401401RP01Invalid token (B2B)Access Token Not Exist/Access Token Expired/Token is Invalid403403RP01Feature Not AllowedThe feature is not allowed to use. Please make sure you registered this feature in your partnership data.403403RP15Transaction Not Permitted.

transType Not FoundtransType not available in the list404404RP01Transaction Not FoundoriginalPartnerReferenceNo Not Found404404RP08Invalid MerchantInvalid Merchant409409RP00ConflictX-EXTERNAL-ID duplicate409409RP01Duplicate partnerReferenceNoDuplicate partnerReferenceNo429429RP00Too Many RequestsMaximum transaction limit exceeded500500RP00General ErrorWrong request field input500500RP01Internal Server ErrorUnknown internal server failure, please retry the process again504504RP00Timeout

Your transaction timeout

POST /openapi/v1.0/ecr/payment

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
Request Purchase
curl –X POST https://sandbox.bca.co.id/openapi/v1.0/ecr/payment
-H 'Content-type: application/json'
-H 'Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-TIMESTAMP: 2024-03-25T13:50:04+07:00'
-H 'X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'ORIGIN: www.hostname.com'
-H 'CHANNEL-ID: 95251'
-H 'X-PARTNER-ID: 82150823919040774621823174737537'
-H 'X-EXTERNAL-ID: 41807553358950093184162180797837'
-d '
{
 "partnerReferenceNo": "2020102900000000000001",
 "merchantId": "000002872",
 "terminalId": "C2012345",
 "transType": "01",
 "partnerType": "",
 "additionalInfo": {
 "amount1":"100000000.00"
 }
}'

Request Payment Sale (Debit/Credit)
curl –X POST https://sandbox.bca.co.id/openapi/v1.0/ecr/payment
-H 'Content-type: application/json'
-H 'Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-TIMESTAMP: 2024-03-25T13:50:04+07:00'
-H 'X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'ORIGIN: www.hostname.com'
-H 'CHANNEL-ID: 95251'
-H 'X-PARTNER-ID: 82150823919040774621823174737537'
-H 'X-EXTERNAL-ID: 41807553358950093184162180797837'
-d '
{
 "partnerReferenceNo": "2020102900000000000001",
 "merchantId": "000002872",
 "terminalId": "C2012345",
 "transType": "02",
 "partnerType": "",
 "additionalInfo": {
 "amount1":"200000000.00",
 "amount2":"100000000.00"
 }
}'

Request Payment QRIS
curl –X POST https://sandbox.bca.co.id/openapi/v1.0/ecr/payment
-H 'Content-type: application/json'
-H 'Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-TIMESTAMP: 2024-03-25T13:50:04+07:00'
-H 'X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'ORIGIN: www.hostname.com'
-H 'CHANNEL-ID: 95251'
-H 'X-PARTNER-ID: 82150823919040774621823174737537'
-H 'X-EXTERNAL-ID: 41807553358950093184162180797837'
-d '
{
 "partnerReferenceNo": "2020102900000000000001",
 "merchantId": "000002872",
 "terminalId": "C2012345",
 "transType": "31",
 "partnerType": "",
 "additionalInfo": {
 "amount1":"200000000.00",
 "amount2":"100000000.00"
 }
}'

Request Payment Void
curl –X POST https://sandbox.bca.co.id/openapi/v1.0/ecr/payment
-H 'Content-type: application/json'
-H 'Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-TIMESTAMP: 2024-03-25T13:50:04+07:00'
-H 'X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'ORIGIN: www.hostname.com'
-H 'CHANNEL-ID: 95251'
-H 'X-PARTNER-ID: 82150823919040774621823174737537'
-H 'X-EXTERNAL-ID: 41807553358950093184162180797837'
-d '
{
 "partnerReferenceNo": "2020102900000000000001",
 "merchantId": "000002872",
 "terminalId": "C2012345",
 "transType": "08",
 "partnerType": "",
 "additionalInfo": {
 "originalPartnerReferenceNo":"41807553358950093184162180797837"
 }
}'

Request Payment Settlement
curl –X POST https://sandbox.bca.co.id/openapi/v1.0/ecr/payment
-H 'Content-type: application/json'
-H 'Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-TIMESTAMP: 2024-03-25T13:50:04+07:00'
-H 'X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'ORIGIN: www.hostname.com'
-H 'CHANNEL-ID: 95251'
-H 'X-PARTNER-ID: 82150823919040774621823174737537'
-H 'X-EXTERNAL-ID: 41807553358950093184162180797837'
-d '
{
 "partnerReferenceNo": "2020102900000000000001",
 "merchantId": "000002872",
 "terminalId": "C2012345",
 "transType": "10",
 "partnerType": ""
}'
```

POST /openapi/v1.0/ecr/payment

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
Success Sample
-H 'Content-Type: application/json'
-H 'X-TIMESTAMP: 2024-03-25T13:50:06+07:00'
{
 "responseCode": "200RP00",
 "responseMessage": "Successful",
 "merchantId": "000002872",
 "terminalId": "C2825423",
 "transType": "02",
 "partnerReferenceNo": "41807553358950093184162180797837",
 "createdDate": "2024-12-23T08:43:11+07:00"
}

Error Sample
-H 'Content-type: application/json'
-H 'X-TIMESTAMP: 2024-03-25T13:50:04+07:00'
{
 "responseCode": "400RP02",
 "responseMessage": "Invalid Mandatory Field merchantId",
 "partnerReferenceNo": "21132312100000001"
}
```

#### QRIS MPM

##### 1\. SNAP QR MPM Generate QR

This service is used to generate QRIS.

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)FixedY Channel identifier. Using EDC’s WSID (95251)X-PARTNER-IDHeaderString(36)MaxY Unique ID for merchant

\- Merchant Reguler: MerchantID (9)

\- Merchant Partnership: Partner Name (36)

X-EXTERNAL-IDHeaderString(36)MaxYID that should be unique per partner per service in the same day

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString(64) Max YAlphanumeric Transaction ID provided by the partneramountObjectY

     valueString (13.2)

Max Y Numeric Net amount of the transaction (exclude convenience fee)

Example : IDR 10.000,- will be placed with 10000.00      currencyString (3)

FixedYAlphanumeric

Currency code (IDR)

merchantId

String (9)

MaxYNumericMerchant identifier that is unique per each merchant.

For merchant aggregator, filled with aggregator's merchantID

subMerchantId

String (9)

MaxCAlphanumericMandatory applied to merchant aggregator, filled with submerchantID

terminalId

String (8)

FixedYNumericTerminal ID

Can be :

\- TID provided by BCA (merchant reguler)

\- TID provided by partner (merchant partnership)
validityPeriodString (25)MaxNYYYY-MM-DD Thh:mm:ssTZ DThe time when the QRIS valid

Minimum : created time + 5 mins

Maximum : created time + 120 mins

Default value 60 minutes will be applied if validityPeriod not providedadditionalInfo

ObjectNAdditional Information

     convenienceFeeString (10.2)

MaxNNumericConvenience fee or tips

     partnerMerchantTypeString (4)

MaxCAlphanumericMandatory applied to merchant partnership, filled with partner identifier

     terminalLocationNameString (25)

MaxCAlphanumericMandatory applied to merchant aggregator, filled with submerchant name

     qrOption

String (1)FixedNAlphabetChoice of QR content or QR image format as response

Value : (case non sensitive)

C = QR Content

I = QR Image

A = QR Content & QR Image

Default QR Image format will be applied if not stated in request / other than the above values

Result of the request will contains following information:

**Response:**

FieldDataType**Length Type****Mandatory****Format**DescriptionresponseCode

String(7) FixedYNumericResponse code to identify generate QRIS request status, refer to the list of response code

Format: AAABBCC

AAA : HTTP Code

BB : Service Code

CC : Case Code

Service Code Generate QRIS : 47

Example : 2004700

responseMessage

String (150)

MaxYAlphanumericError message in English

Example : successful

referenceNoString (36)

MaxNAlphanumericTransaction ID provided by BCA

Mandatory applied only if the request from merchant is validpartnerReferenceNo

String (64) MaxYAlphanumericTransaction ID provided by the partner

qrContentString (512) MaxN AlphanumericQR String MPMqrUrlString (256)Max NAlphanumericURL for download QR Image

Not Used in BCAqrImage

StringMaxNAlphanumericLength = unlimited

Encode string (Base64) from QR image

QR String MPM

Mandatory applied only if QR successfully generatedredirectUrlString (512)MaxNAlphanumericRedirect URL to PJSP page to process the payment

Not used in BCAmerchantName

String (25)

MaxNAlphanumericMerchant name for print receipt

For merchant aggregator, filled with the submerchant's outlet name

Mandatory applied only if QR successfully generated

storeIdString (64)MaxNAlphanumericUnique shop id

Not used in BCAterminalId

String (8)

MaxN

Alphanumeric

Terminal ID

Can be:

\- TID provided by BCA (merchant reguler)

\- TID provided by partner (merchant partnership)

Mandatory applied only if QR successfully generatedadditionalInfoObjectNAdditional Information

**Error**

HTTP CodeError CodeError Message**Condition**400 4004700Bad requestGenerate request failed error4004004701Invalid Field Format {Field}The data you entered does not match the format requirements4004004702Invalid Mandatory Field {Field}Mandatory field should be fulfilled

4014014700Unauthorized. \[Reason\]Invalid Signature/Unknown Client/Connection Not Allowed/Invalid API/Company is not authorized4014014701Invalid token (B2B)Access Token Not Exist/Access Token Expired/Token is Invalid4034034701Feature Not AllowedThe feature is not allowed to use. Please make sure you registered this feature in your partnership data403

4034706

Feature Not Allowed At This Time

The feature is not allowed at this time. Please try again

4044044701Transaction Not FoundTransaction not found.

404

4044708

Invalid Merchant

Invalid merchant

4044044718Inconsistent Request \[X-PARTNER-ID & merchantID not match\]merchantID and X-PARTNER-ID value are not match. These fields should consist of same value

409 4094700ConflictX-EXTERNAL-ID duplicate409 4094701 Duplicate partnerReferenceNo partnerReferenceNo duplicate

4294294700

Too Many Requests

Maximum request limit exceeded

5005004700General ErrorGeneral Error

5045044700TimeoutYour request timeout

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/qr/qr-mpm-generate

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
Merchant Reguler

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/qr/qr-mpm-generate
-H 'Authorization:Bearer iNrWtmeHe7t3h16x4i6o1kV9yj575H77iB9daG3kKXcE7MOHU2ubaS'
-H 'X-TIMESTAMP: 2024-09-01T10:05:08+07:00'
-H 'X-SIGNATURE:HTcmQDWBXwBvzkIPCYYxOvV+rzFjEEgpnMBP28Jydfi+HMJZsmek70Ix7VDibQbAveswQAr6JoNzpH2w6b7bAWsfg1hdwJeMK…'
-H 'Content-type:application/json'
-H 'CHANNEL-ID: 95251'
-H 'X-PARTNER-ID: 123456789'
-H 'X-EXTERNAL-ID: 41807553358950093184162180797837'
-d '
{
    "partnerReferenceNo": "2024090100000000000001",
    "amount": {
        "value": "50000.00",
        "currency": "IDR"
    },
    "merchantId": "123456789",
    "subMerchantId": "",
    "terminalId": "A1234567",
    "validityPeriod": "2024-09-01T10:10:10+07:00",
    "additionalInfo": {
        "convenienceFee": "0.00",
        "partnerMerchantType": "",
        "terminalLocationName": "",
        "qrOption": "C"
    }
}'

Merchant Partnership
curl –X POST https://sandbox.bca.co.id/openapi/v1.0/qr/qr-mpm-generate
-H 'Authorization:Bearer iNrWtmeHe7t3h16x4i6o1kV9yj575H77iB9daG3kKXcE7MOHU2ubaS'
-H 'X-TIMESTAMP: 2024-09-01T10:05:08+07:00'
-H 'X-SIGNATURE:HTcmQDWBXwBvzkIPCYYxOvV+rzFjEEgpnMBP28Jydfi+HMJZsmek70Ix7VDibQbAveswQAr6JoNzpH2w6b7bAWsfg1hdwJeMK…'
-H 'Content-type:application/json'
-H 'CHANNEL-ID: 95251'
-H 'X-PARTNER-ID:partnerABC '
-H 'X-EXTERNAL-ID: 41807553358950093184162180797837'
-d '
{
    "partnerReferenceNo": "2024090100000000000001",
    "amount": {
        "value": "50000.00",
        "currency": "IDR"
    },
    "merchantId": "123456789",
    "subMerchantId": "234891823",
    "terminalId": "A1234567",
    "validityPeriod": "2024-09-01T10:10:10+07:00",
    "additionalInfo": {
        "convenienceFee": "0.00",
        "partnerMerchantType": "XX",
        "terminalLocationName": "Merchant XYZ",
        "qrOption": ""
    }
}'
```

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/qr/qr-mpm-generate

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
-H 'Content-type:application/json'
-H 'X-TIMESTAMP: 2024-09-01T10:05:10+07:00'
{
    "responseCode": "2004700",
    "responseMessage": "Successful",
    "referenceNo": "43057a74-e179-49ed-b119-b004a3d522ec",
    "partnerReferenceNo": "2024090100000000000001",
    "qrContent":"0002010102122654000200011893600014300001450702150008850000145070303UMI51240008ID.GPNQR020100303UMI
52045814530336054061000055802ID5915AGGREGOBISA16013JAKARTAPUSAT6105102306259010680394905120242998039490708A1AAC0929
9170002000107DINAMIS6304FADF",
"qrUrl": null,
    "qrImage": "iVBORw0KGgoAAAANSUhEUgAAASwAAAF8CAYAAAB184aCAAAW2klEQVR42u3dv67tJJkS41kXgboE2Xt0DwGCmSBsEDpE9x...",
    "redirectUrl": null,
    "merchantName": "Merchant XYZ",
    "storeId": null,
    "terminalId": "A1234567",
    "additionalInfo": null
}
```

##### 2\. SNAP QR MPM Inquiry

This service is used to inquiry transaction status.

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString (5)FixedY Channel’s identifier using EDC’s WSID (95251)X-PARTNER-IDHeaderString (36)MaxY Unique ID for a merchant

\- Merchant Reguler : MerchantID (9)

\- Merchant Partnership: Partner Name (36) X-EXTERNAL-ID

HeaderString (36)MaxYID that should be unique per partner per service in the same day

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionoriginalReferenceNoString (36) Max YAlphanumeric Transaction ID provided by the partner

Taken from response API Generate QR 'referenceNo' FieldoriginalPartnerReferenceNoString (64)Max YAlphanumeric Transaction ID provided by the partnerserviceCodeString (2)FixedYNumeric Transaction type indicator (service code of the original transaction request)

Filled with serviceCode generate QR = 47merchantIdString (9)MaxYNumericMerchant identifier that is unique per each merchant

For merchant aggregator, filled with aggregator's merchantIDsubMerchantIdString (9)MaxCAlphanumericMandatory applied to merchant aggregator, filled with submerchantIDadditionalInfoObjectNAdditional Information     partnerMerchantTypeString (4)MaxCAlphanumericMandatory applied to merchant aggregator, filled with partner identifier     terminalIdString (8) MaxYAlphanumericTerminal ID

Can be :

\- TID provided by BCA (merchant reguler)

\- TID provided by partner (merchant partnership)

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)FixedYNumericResponse code to identify inquiry request status, refer to the list of response code

Response code is different from transaction status ('latestTransactionStatus' field)

format : AAABBCC

AAA : HTTP Code

BB : Service Code

CC : Case Code

Service Code Inquiry Payment QRIS : 51

Example: 2005100responseMessageString (150)MaxYAlphanumericResponse message in English, refer to the list of response code

Example: successfuloriginalPartnerReferenceNo

String (64) Max

YAlphanumericTransaction ID provided by the partneroriginalReferenceNoString (36) Max YAlphanumeric Transaction ID provided by BCAoriginalExternalIdString (36)FixedYNumeric Original External-ID on header message requestserviceCodeString (2)FixedYNumericTransaction type indicator (service code of the original transaction request)

Filled with serviceCode generate QR = 47latestTransactionStatusString (2)FixedYNumericTransaction Status

00 - Success : transaction is paid

03 - Pending : transaction is recorded, but still unpaid

04 - Refunded : transaction is successfully refunded

06 - Failed : transaction failed, because QR Expired transactionStatusDescString (50) MaxYAlphanumericShow description of transaction status

Example:

\- Success (transaction Success)

\- Pending (Unpaid Transaction)

\- Failed (QR Expired)

\- Refunded (Partial refund, available amount : Rp 5000)

\- Refunded (Fully refund)

\- Refunded (Max count limit)paidTimeString (25)MaxNYYY-MM-DDThh:mm:ssTZD

Transaction date

Example: 2020-12-21T10:30:24+07:00

Mandatory applied only if the transaction is successfulamountObject

NMandatory applied only if the transaction is successful

     valueString (13.2) MaxYNumericNet amount of the transaction (exclude convenience fee)

Example: IDR 10.000,- will be placed with 10000.00      currencyString (3)FixedYAlphanumericCurrency code (IDR)feeAmountObjectN     valueString (16.2

MaxYNumericDefault in BCA = 0.00     currencyString (3)FixedYAlphanumericCurrency code (IDR)terminalIdString (8)MaxNAlphanumericTerminal ID. Can be:

\- TID provided by BCA (merchant reguler)

\- TID provided by partner (merchant partnership)

Mandatory applied only if the transaction is successfuladditionalInfoObjectNMandatory applied only if the transaction is successful     referenceNumberString (12)FixedYAlphanumericUnique ID for tracing     approvalCodeString (6)FixedNAlphanumericApproval code from host for success transaction

Mandatory applied only if the transaction is successful     payerPhoneNumberString (13)MaxNAlphanumericMasked buyer phone number     batchNumberString (6)FixedNNumericTransaction batch number

Mandatory applied only if the transaction is successful     convenienceFeeString (10.2)Max NNumericConvenience fee or tips

     customerPanString (19)MaxNNumericBuyer card number (customer\_pan)     issuerReferenceNumberString (12)MaxNNumericRRN QRIS

Mandatory applied only if the transaction is successful     payerNameString (30)MaxNAlphanumericBuyer name

Mandatory applied only if the transaction is successful     issuerNameString (26)MaxNAlphanumericIssuer name

Mandatory applied only if the transaction is successful     acquirerNameString (3)FixedNAlphanumericFilled with 'BCA'

Mandatory applied only if the transaction is successful     merchantInfoObjectN          merchantIdString (15) MaxNNumericMerchant identifier that is unique per each merchant

Format: 000 885 + MID (9 digit)

For merchant aggregator, filled with subMerchantId

Mandatory applied only if the transaction is successful.          merchantPanString (19)MaxNNumericMerchant PAN

Mandatory applied only if the transaction is successful          nameString (25) Max NAlphanumericMerchant name          cityString (13) MaxNAlphanumericMerchant city

Mandatory applied only if the transaction is successful          postalCodeString (10)MaxNNumericMerchant postal code

Mandatory applied only if the transaction is successful          countryString (2)FixedNAlphanumericMerchant location country

Mandatory applied only if the transaction is successful          emailString (30)MaxNAlphanumericMerchant e-mail address          paymentChannelNameString (10) MaxNAlphanumeric\- Sakuku

\- Debit

\- Switching

\- Paylater

\- Credit

Mandatory applied only if the transaction is successful

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message**Condition**4004005100Bad requestGeneral request failed error

4004005101Invalid Field Format {Field} The data you entered does not match the format requirements

4004005102Invalid Mandatory Field {Field}Mandatory field should be fulfilled

4014015100Unauthorized. \[Reason\]Invalid Signature/Unknown Client/Connection Not Allowed/Invalid API/Company is not authorized4014015101Invalid token (B2B)Access Token Not Exist/Access Token Expired/Token is Invalid

4034035100Transaction ExpiredQR has expired

4034035101Feature Not AllowedThe feature is not allowed to use. Please make sure you registered this feature in your partnership data4034035106Feature Not Allowed At This TimeThe feature is not allowed at this time. Please try again

404

4045101

Transaction Not Found

Transaction not found

4044045108Invalid MerchantInvalid merchant

4044045118Inconsistent Request \[X-PARTNER-ID & merchantID not match\]merchantID and X-PARTNER-ID value are not match. These fields should consist of same value

4094095100 ConflictX-EXTERNAL-ID duplicate

4294295100Too Many RequestsMaximum request limit exceeded

5005005100General ErrorGeneral Error

5005005101Internal Server ErrorUnknown internal server failure, please retry the process again

5045045100TimeoutYour request timeout

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/qr/qr-mpm-query

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
Merchant Reguler

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/qr/qr-mpm-query
-H 'Authorization:Bearer iNrWtmeHe7t3h16x4i6o1kV9yj575H77iB9daG3kKXcE7MOHU2ubaS'
-H 'X-TIMESTAMP:2024-09-01T08:45:00+07:00'
-H 'X-SIGNATURE:HTcmQDWBXwBvzkIPCYYxOvV+rzFjEEgpnMBP28Jydfi+HMJZsmek70Ix7VDibQbAveswQAr6JoNzpH2w6b7bAWsfg1hdwJeMK…'
-H 'Content-type:application/json'
-H 'CHANNEL-ID: 95251'
-H 'X-PARTNER-ID: 123456789'
-H 'X-EXTERNAL-ID: 41807553358950093184162180797839
-d '
{
    "originalPartnerReferenceNo": "2024090100000000000001",
    "originalReferenceNo": "43057a74-e179-49ed-b119-b004a3d522ec",
    "serviceCode": "47",
    "merchantId": "123456789",
    "subMerchantId": "",
    "additionalInfo": {
        "terminalId": "A1234567",
        "partnerMerchantType": ""
    }
}'

Merchant Partnership

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/qr/qr-mpm-query
-H 'Authorization:Bearer iNrWtmeHe7t3h16x4i6o1kV9yj575H77iB9daG3kKXcE7MOHU2ubaS'
-H 'X-TIMESTAMP:2024-09-01T08:45:00+07:00'
-H 'X-SIGNATURE:HTcmQDWBXwBvzkIPCYYxOvV+rzFjEEgpnMBP28Jydfi+HMJZsmek70Ix7VDibQbAveswQAr6JoNzpH2w6b7bAWsfg1hdwJeMK…'
-H 'Content-type:application/json'
-H 'CHANNEL-ID: 95251'
-H 'X-PARTNER-ID:partnerABC'
-H 'X-EXTERNAL-ID: 41807553358950093184162180797839
-d '
{
    "originalPartnerReferenceNo": "2024090100000000000001",
    "originalReferenceNo": "43057a74-e179-49ed-b119-b004a3d522ec",
    "serviceCode": "47",
    "merchantId": "123456789",
    "subMerchantId": "234891823",
    "additionalInfo": {
        "terminalId": "A1234567",
        "partnerMerchantType": "XX"
    }
}'
```

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/qr/qr-mpm-query

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
Merchant Reguler - Transaction status success

-H 'Content-type:application/json'
-H 'X-TIMESTAMP:2024-09-01T08:45:05+07:00'
{
    "responseCode": "2005100",
    "responseMessage": "Success",
    "originalPartnerReferenceNo": "2024090100000000000001",
    "originalReferenceNo": "43057a74-e179-49ed-b119-b004a3d522ec",
    "originalExternalId": "41807553358950093184162180797839",
    "serviceCode": "47",
    "latestTransactionStatus": "00",
    "transactionStatusDesc": "Success (transaction Success)",
    "paidTime": "2024-09-01T08:43:11+07:00",
    "amount": {
        "value": "50000.00",
        "currency": "IDR"
    },
    "feeAmount": null,
    "terminalId": "A1234567",
    "additionalInfo": {
        "referenceNumber": "022411000106",
        "approvalCode": "223356",
        "payerPhoneNumber": "*******4450",
        "batchNumber": "000004",
        "convenienceFee": "0.00",
        "issuerReferenceNumber": "000013101103",
        "payerName": "Customer A",
        "customerPan": "9360001430000131018",
        "issuerName": "BCA",
        "acquirerName": "BCA",
        "merchantInfo": {
            "merchantId": "000885123456789",
            "merchantPan": "9360001431234567899",
            "name": "Merchant XYZ",
            "city": "JAKARTA PUSAT",
            "postalCode": "10310",
            "country": "ID",
            "email": null,
            "paymentChannelName": "Sakuku"
        }
    }
}'

Merchant Partnership - Transaction status success

-H 'Content-type:application/json'
-H 'X-TIMESTAMP:2024-09-01T08: 45: 05+07: 00'
{
    "responseCode": "2005100",
    "responseMessage": "Success",
    "originalPartnerReferenceNo": "2024090100000000000001",
    "originalReferenceNo": "43057a74-e179-49ed-b119-b004a3d522ec",
    "originalExternalId": "41807553358950093184162180797839",
    "serviceCode": "47",
    "latestTransactionStatus": "00",
    "transactionStatusDesc": "Success (transaction Success)",
    "paidTime": "2024-09-01T08:43:11+07:00",
    "amount": {
        "value": "50000.00",
        "currency": "IDR"
    },
    "feeAmount": null,
    "terminalId": "A1234567",
    "additionalInfo": {
        "referenceNumber": "022411000106",
        "approvalCode": "223356",
        "payerPhoneNumber": "*******4450",
        "batchNumber": "000004",
        "convenienceFee": "0.00",
        "issuerReferenceNumber": "000013101103",
        "payerName": "Customer A",
        "customerPan": "9360001430000131018",
        "issuerName": "BCA",
        "acquirerName": "BCA",
        "merchantInfo": {
            "merchantId": "000885234891823",
            "merchantPan": "9360001432348918239",
            "name": "Merchant XYZ",
            "city": "JAKARTA PUSAT",
            "postalCode": "10310",
            "country": "ID",
            "email": null,
            "paymentChannelName": "Sakuku"
        }
    }
}'

Merchant Reguler / Partnership - Transaction status failed

-H 'Content-type:application/json'
-H 'X-TIMESTAMP:2024-09-01T08:45:05+07:00'
{
    "responseCode": "2005100",
    "responseMessage": "Success",
    "originalReferenceNo": "43057a74-e179-49ed-b119-b004a3d522ec",
    "originalPartnerReferenceNo": "2024090100000000000001",
    "originalExternalId": "41807553358950093184162180797839",
    "serviceCode": "47",
    "latestTransactionStatus": "06",
    "transactionStatusDesc": "Failed (QR Expired)",
    "paidTime": null,
    "amount": null,
    "feeAmount": null,
    "terminalId": null,
    "additionalInfo": {
        "referenceNumber": "022411000106",
        "approvalCode": null,
        "payerPhoneNumber": null,
        "batchNumber": null,
        "convenienceFee": null,
        "issuerReferenceNumber": null,
        "payerName": null,
        "customerPan": null,
        "issuerName": null,
        "acquirerName": null,
        "merchantInfo": null
    }
}'
```

##### 3\. SNAP QR MPM Refund

This service is used to request refund.

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString (5)FixedY Channel’s identifier. Using EDC’s WSID (95251)X-PARTNER-IDHeaderString (36)MaxY Unique ID for a merchant

\- Merchant Reguler : MerchantID (9)

\- Merchant Partnership: Partner Name (36) X-EXTERNAL-ID

HeaderString (36)MaxYID that should be unique per partner per service in the same day

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionmerchantIdString (9)MaxYNumericMerchant identifier that is unique per each merchant

For merchant aggregator, filled with aggregator's merchantIDsubMerchantIdString (9)MaxCAlphanumericMandatory applied to merchant aggregator, filled with submerchantIDoriginalPartnerReferenceNoString (64)MaxYAlphanumericTransaction ID provided by the partneroriginalReferenceNoString (36)MaxYAlphanumericTransaction ID provided by BCA

Taken from response API Generate QR 'referenceNo' field

partnerRefundNoString (12)MaxYAlphanumericReference Number issuer

Taken from Service API Inquiry / Payment Notification 'issuerReferenceNumber' field

refundAmountObjectN     valueString (13.2)MaxYNumericRefund amount

Example: IDR 10.000,- will be placed with 10000.00

     currencyString (3)FixedYAlphanumericCurrency code (IDR)additionalInfoObjectNAdditional Information     terminalIdString (8) MaxYAlphanumericTerminal ID     transactionDateString (25)FixedYYYYY-MM-DD Thh:mm:ssTZ D

Transaction date

Example: 2020-12-21T10:30:24+07:00

Taken from Service API Inquiry 'paidTime' field / Payment Notification 'transactionDate' field

     partnerMerchantTypeString (4)MaxCAlphanumericMandatory applied to merchant partnership, filled with partner identifier     issuerName

String (26)

MaxNAlphanumericIssues name

Taken from Service API Inquiry / Payment

Notification 'issuerName' field

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)FixedYNumericResponse code to identify refund request status, refer to the list of response code

format : AAABBCC

AAA : HTTP Code

BB : Service Code

CC : Case Code

Service Code Inquiry Payment QRIS : 78

Example: 2007800responseMessageString (150)MaxYAlphanumericResponse message in English

Example: successfuloriginalPartnerReferenceNo

String (64) Max

YAlphanumericTransaction ID provided by the partneroriginalReferenceNoString (36) Max YAlphanumeric Transaction ID provided by BCAoriginalExternalIdString (36)FixedYNumeric Original External-ID on header message requestrefundNoString (12) MaxNAlphanumericRefund ID provided by BCA

Mandatory applied only if the refund transaction is successfulpartnerRefundNoString (12)

Max

Y AlphanumericReference Number issues taken from Service API

Inquiry / Payment Notification

'issuerReferenceNumber' fieldrefundAmountObjectNMandatory applied only if the refund transaction is succesful     valueString (13.2) MaxYNumericRefund amount

Example: IDR 10.000,- will be placed with 10000.00      currencyString (3)FixedYAlphanumericCurrency code (IDR)refundTimeString (25)FixedNYYYY-MM-DD Thh:mm:ssTZ D

Refund Time

Example: 2020-10-20T17:56:57+07:00

Mandatory applied only if the refund transaction is successfuladditionalInfoObjectNAdditional Information     merchantId

String (15)MaxNNumericMerchant identifier that is unique per each merchant

Format: 000 885 + MID (9 digit)

For merchant aggregator, filled with submerchantID

Mandatory applied only if the transaction is successful

     terminalIdString (8)MaxNAlphanumericTerminal ID

Mandatory applied only if the refund transaction is successful

     referenceNumberString (12)FixedYAlphanumericUnique ID for tracing     availableAmountObjectNMandatory applied only if the refund transaction is successful         valueString (13.2)MaxYNumericAvailable amount for refund         currencyString (3)FixedYAlphanumericCurrency code (IDR)     refundCounterString (1)FixedNNumericMandatory applied only if the transaction is successful

counter refund recorded for this transaction

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message**Condition**4004007800Bad requestGeneral request failed error

4004007801Invalid Field Format {Field} The data you entered does not match the format requirements

4004007802Invalid Mandatory Field {Field}Mandatory field should be fulfilled

4014017800Unauthorized. \[Reason\]Invalid Signature/Unknown Client/Connection Not Allowed/Invalid API/Company is not authorized4014017801Invalid token (B2B)Access Token Not Exist/Access Token Expired/Token is Invalid

4034037801Feature Not AllowedThe feature is not allowed to use. Please make sure you registered this feature in your partnership data4034037805Do not honorDo not honor4034037806Feature Not Allowed At This TimeThe feature is not allowed at this time. Please try again

4034037815Transaction Not Permitted \[Reason\]Transaction data does not meet the refund's criteria/Refund payment is failed because Payment transaction status has been settled404

4047801

Transaction Not Found

Transaction not found

4044047808Invalid MerchantInvalid merchant

4044047813Invalid AmountRefund nominal is greater than the remaining transaction nominal

4044047818Inconsistent Request \[X-PARTNER-ID & merchantID not match\]merchantID and X-PARTNER-ID value are not match. These fields should consist of same value

4054057801Requested operation to refund transaction is not allowed at this time

The issuer used for refund does not support refund

4094097800 ConflictX-EXTERNAL-ID duplicate

4294297800Too Many RequestsMaximum request limit exceeded

5005007800General ErrorGeneral Error

5005007801Internal Server ErrorUnknown internal server failure, please retry the process again

5045047800TimeoutYour request timeout

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/qr/qr-mpm-refund

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
curl –X POST https://sandbox.bca.co.id/openapi/v1.0/qr/qr-mpm-refund
-H 'Authorization:Bearer iNrWtmeHe7t3h16x4i6o1kV9yj575H77iB9daG3kKXcE7MOHU2ubaS'
-H 'X-TIMESTAMP:2024-09-03T14:48:25+07:00'
-H 'X-SIGNATURE:HTcmQDWBXwBvzkIPCYYxOvV+rzFjEEgpnMBP28Jydfi+HMJZsmek70Ix7VDibQbAveswQAr6JoNzpH2w6b7bAWsfg1hdwJeMK…'
-H 'Content-type:application/json'
-H 'CHANNEL-ID: 95251'
-H 'X-PARTNER-ID: 123456789'
-H 'X-EXTERNAL-ID: 41807553358950093184162180797840'
-d '
{
    "merchantId": "123456789",
    "originalPartnerReferenceNo": "2024090100000000000001",
    "originalReferenceNo": "43057a74-e179-49ed-b119-b004a3d522ec",
    "partnerRefundNo": "000013101103",
    "refundAmount": {
        "value": "10000.00",
        "currency": "IDR"
    },
    "additionalInfo": {
        "terminalId": "A1234567",
        "transactionDate": "2024-09-01T08:43:11+07:00",
        "partnerMerchantType": "",
        "issuerName": "BCA"
    }
}'

```

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/qr/qr-mpm-refund

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
Refund payment status success

-H 'Content-type:application/json'
-H 'X-TIMESTAMP:2024-09-03T14:48:30+07:00'
{
    "responseCode": "2007800",
    "responseMessage": "Successful",
    "originalPartnerReferenceNo": "2024090100000000000001",
    "originalReferenceNo": "43057a74-e179-49ed-b119-b004a3d522ec",
    "originalExternalId": "41807553358950093184162180797840",
    "refundNo": "21E4172823",
    "partnerRefundNo": "000013101103",
    "refundAmount": {
        "value": "10000.00",
        "currency": "IDR"
    },
    "refundTime": "2024-09-03T14:48:27+07:00",
    "additionalInfo": {
        "merchantId": "123456789",
        "terminalId": "A1234567",
        "referenceNumber": "022411000106",
        "availableAmount": {
            "value": "40000.00",
            "currency": "IDR"
        },
        "refundCounter": "1"
    }
}

Refund payment status failed

-H 'Content-type:application/json'
-H 'X-TIMESTAMP:2024-09-03T14:48:30+07:00'
{
    "responseCode": "4037815",
    "responseMessage": "Transaction Not Permitted",
    "originalPartnerReferenceNo": "2024090100000000000001",
    "originalReferenceNo": "43057a74-e179-49ed-b119-b004a3d522ec",
    "originalExternalId": "41807553358950093184162180797840",
    "refundNo": "",
    "partnerRefundNo": "000013101103",
    "refundAmount": {
        "value": "10000.00",
        "currency": "IDR"
    },
    "refundTime": "",
    "additionalInfo": {
        "merchantId": "123456789",
        "terminalId": "A1234567",
        "referenceNumber": "022411000106",
        "availableAmount": {
            "value": "",
            "currency": ""
        },
        "refundCounter": ""
    }
}
```

##### 4\. SNAP QRIS Notification

This service is used to Notification QRIS.

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

Field**Params Type**Data Type**Length Type**

**Mandatory**DescriptionCHANNEL-IDHeaderString (5)FixedY Channel’s identifier using EDC’s WSID (95251)X-PARTNER-IDHeaderString (36)MaxY Unique ID for a partner

\- Merchant Reguler: MerchantID (15)

\- Merchant Partnership : Partner Name (36)X-EXTERNAL-ID

HeaderString (36)MaxYID that should be unique per co-partner per service in the same day

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionoriginalReferenceNoString (36) Max YAlphanumeric TransactionID provided by BCAoriginalPartnerReferenceNoString (64)Max NAlphanumeric Transaction ID provided by the partnerlatestTransactionStatusString (2)FixedYNumericTransaction Status

00 - Success : transaction is paid

03 - Pending : transaction is recorded, but still unpaid

06 - Failed : transaction failed, because QR Expired transactionStatusDescString (50)MaxYAlphanumericShow description of transaction status

Example:

\- Success (transaction Success)

\- Pending (Unpaid Transaction)

\- Failed (QR Expired) customerNumberString (19)MaxNNumericBuyer card number (customer\_pan)

Mandatory applied only if the transaction is successfulaccountTypeString (25)MaxNAlphanumericAccount type

Not used in BCAdestinationNumberString (19)MaxNNumericMerchant PAN

Mandatory applied only if the transaction is successfuldestinationAccountNameString(25)MaxNAlphanumericMerchant name

Mandatory applied only if the transaction is successfulamountObject

N

Mandatory applied only if the transaction is successful

     valueString (13.2)

Max Y NumericNet amount of the transaction (exclude convenience fee)

Example : IDR 10.000,- will be placed with 10000.00      currencyString (3)

FixedYAlphanumeric

Currency code (IDR)

bankCode

String (8)

NNumericBank Code

Not used in BCAadditionalInfoObject

N

Additional Information     referenceNumberString (12)FixedYAlphanumeric

Unique ID for tracing     transactionDateString (25)FixedNYYYY-MM:DDThh:mm:ssTZDTransaction date

Example: 2020-12-21T10:30:24+07:00

Mandatory applied only if the transaction is successful

     approvalCodeString (6)

FixedNAlphanumericApproval code from host for success transaction

Mandatory applied only if the transaction is successful     payerPhoneNumberString (13)

MaxNAlphanumericMasked Buyer phone number     batchNumberString(6)FixedNNumericTransaction batch number

Mandatory applied only if the transaction is successful     convenienceFeeString(10.2)MaxNNumericConvenience fee or tips      issuerReferenceNumberString(12)MaxNNumericRRN QRIS

Mandatory applied only if the transaction is successful     payerNameString(30)MaxNAlphanumericBuyer name

Mandatory applied only if the transaction is successful     issuerNameString(26)MaxNAlphanumericIssuer name

Mandatory applied only if the transaction is successful     acquirerNameString(3)FixedNAlphanumericFilled with 'BCA'

Mandatory applied only if the transaction is successful     merchantInfoArray of ObjectN          terminalIdString(8)MaxNAlphanumericTerminal ID

Can be:

\- TID provided by BCA (merchant reguler)

\- TID provided by partner (merchant partnership)

Mandatory applied only if the transaction is successful          merchantIdString(15)MaxNNumericMerchant identifier that is unique per each merchant

Format: 000 885 + MID (9 digit)

For merchant aggregator, filled with submerchantID

Mandatory applied only if the transaction is successful.          cityString(13)MaxNAlphanumericMerchant city

Mandatory applied only if the transaction is successfu.          postalCodeString(10)MaxNNumericMerchant postal code

Mandatory applied only if the transaction is successful          countryString(2)FixedNAlphanumericMerchant location country

Mandatory applied only if the transaction is successful          emailString (30)MaxNAlphanumericMerchant e-mail address          paymentChannelNameString(10)Max NAlphanumeric\- Sakuku

\- Debit

\- Switching

\- Paylater

\- Credit

Mandatory applied only if the transaction is successful

Result of the request will contains following information:

**Response:**

FieldDataType**Length Type****Mandatory****Format**DescriptionresponseCode

String(7) FixedYAlphanumeric

Response code to identify Notification request status

Response code is different from transaction status ('latestTransactionStatus' field)

Format: AAABBCC

AAA : HTTP

BB : Service Code

CC : Case Code

Service Code Notification Payment QRIS : 52

Example: 2005200

Response from copartner

responseMessage

String (150)

MaxYAlphanumericError message in English

Example : successful

**Error List** Here is the list of error codes that can be returned

HTTP CodeError CodeError Message**Condition**4004005200Bad requestGenerate request failed error4004005201Invalid Field Format {Field}The data you entered does not match the format requirements4004005202Invalid Mandatory Field {Field}Mandatory field should be fulfilled

4014015200Unauthorized. \[Reason\]Invalid Signature/Unknown Client/Connection Not Allowed/Invalid API/Company is not authorized4014015201Invalid token (B2B)Access Token Not Exist/Access Token Expired/Token is Invalid4034035200Transaction ExpiredQR has expired4034035201Feature Not AllowedThe feature is not allowed to use. Please make sure you registered this feature in your partnership data403

4035206

Feature Not Allowed At This Time

The feature is not allowed at this time. Please try again

4044045201Transaction Not FoundTransaction not found.

404

4045208

Invalid Merchant

Invalid merchant

4044045218Inconsistent Request \[X-PARTNER-ID & merchantID not match\]merchantID and X-PARTNER-ID value are not match. These fields should consist of same value

4094095200ConflictX-EXTERNAL-ID duplicate4294295200

Too Many Requests

Maximum request limit exceeded

5005005200General ErrorGeneral Error

5045045200TimeoutYour request timeout

curl –X POST https://copartners.com/openapi/v1.0/qr-mpm-notify

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
Merchant Reguler

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/qr-mpm-notify
-H 'Content-type: application/json'
-H 'Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-TIMESTAMP: 2024-09-01T08:43:11+07:00'
-H 'X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'CHANNEL-ID: 95251'
-H 'X-PARTNER-ID: 000885123456789'
-H 'X-EXTERNAL-ID: 41807553358950093184162180797837'
-d '
{
    "originalReferenceNo": "43057a74-e179-49ed-b119-b004a3d522ec",
    "originalPartnerReferenceNo": "2024090100000000000001",
    "latestTransactionStatus": "00",
    "transactionStatusDesc": "success",
    "customerNumber": "9360001430000131018",
    "accountType": null,
    "destinationNumber": "9360001431234567899",
    "destinationAccountName": "Merchant XYZ",
    "amount": {
        "value": "50000.00",
        "currency": "IDR"
    },
    "bankCode": null,
    "additionalInfo": {
        "referenceNumber": "022411000106",
        "transactionDate": "2024-09-01T08:43:11+07:00",
        "approvalCode": "223356",
        "payerPhoneNumber": "*******4450",
        "batchNumber": "000004",
        "convenienceFee": "0.00",
        "issuerReferenceNumber": "000013101103",
        "payerName": "Customer A",
        "issuerName": "BCA",
        "acquirerName": "BCA",
        "merchantInfo": {
            "terminalId": "A1234567",
            "merchantId": "000885123456789",
            "city": "JAKARTA PUSAT",
            "postalCode": "10310",
            "country": "ID",
            "email": null,
            "paymentChannelName": "Sakuku"
        }
    }
}'

Merchant Partnership

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/qr-mpm-notify
-H 'Content-type: application/json'
-H 'Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-TIMESTAMP: 2024-09-01T08:43:11+07:00'
-H 'X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'CHANNEL-ID: 95251'
-H 'X-PARTNER-ID: partnerABC'
-H 'X-EXTERNAL-ID: 41807553358950093184162180797837'
-d '
{
    "originalReferenceNo": "43057a74-e179-49ed-b119-b004a3d522ec",
    "originalPartnerReferenceNo": "2024090100000000000001",
    "latestTransactionStatus": "00",
    "transactionStatusDesc": "success",
    "customerNumber": "9360001430000131018",
    "accountType": null,
    "destinationNumber": "9360001432348918239",
    "destinationAccountName": "Merchant XYZ",
    "amount": {
        "value": "50000.00",
        "currency": "IDR"
    },
    "bankCode": null,
    "additionalInfo": {
        "referenceNumber": "022411000106",
        "transactionDate": "2024-09-01T08:43:11+07:00",
        "approvalCode": "223356",
        "payerPhoneNumber": "*******4450",
        "batchNumber": "000004",
        "convenienceFee": "0.00",
        "issuerReferenceNumber": "000013101103",
        "payerName": "Customer A",
        "issuerName": "BCA",
        "acquirerName": "BCA",
        "merchantInfo": {
            "terminalId": "A1234567",
            "merchantId": "000885234891823",
            "city": "JAKARTA PUSAT",
            "postalCode": "10310",
            "country": "ID",
            "email": null,
            "paymentChannelName": "Sakuku"
        }
    }
}'
```

curl –X POST https://copartners.com/openapi/v1.0/qr-mpm-notify

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
-H 'Content-Type: application/json'
-H 'X-TIMESTAMP: 2024-09-01T08:43:11+07:00'
{
    "responseCode": "2005200",
    "responseMessage": "Successful"
}
```

#### QRIS CPM

##### 1\. SNAP QR CPM Cancel Payment

This service is used to process cancel payment for CPM QR Codes.

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString (5) FixedY Channel’s identifier using EDC’s WSID (95251)X-PARTNER-IDHeaderString (36) MaxY Unique ID for merchant

\- Merchant Reguler: MerchantID (9)

\- Merchant Partnership : Partner Name (36)X-EXTERNAL-ID

HeaderString (36) MaxYID that should be unique per co-partner per service in the same day

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionoriginalPartnerReferenceNoString (64)Max YAlphanumeric Transaction ID provided by the partner

Format : include character ‘-‘ merchantIdString (9)MaxYNumeric Merchant identifier that is unique per each merchantSubMerchantIdString (9)MaxCAlphanumeric Mandatory applied to merchant aggregator, filled with sub merchant IDadditionalInfoObjectNAdditional Information     terminalIdString (8)MaxYAlphanumericTerminal ID

Can be :

\- TID generated by BCA (merchant reguler)

\- TID generated by partner (merchant partnership)     partnerMerchantTypeString(4)Max CAlphanumericMandatory applied to merchant aggregator, filled with partner identifier

Result of the request will contains following information:

**Response:**

FieldDataType**Length Type****Mandatory****Format**DescriptionresponseCode

String (7) FixedYNumeric

Response code to identify cancel payment request status, refer to the list of response code.

Response code is different from transaction status ('latestTransactionStatus' field)

Format: AAABBCC

AAA : HTTP Code

BB : Service Code

CC : Case Code

Service Code Cancel Payment QRIS : 62

Example : 2006200

responseMessage

String (150)

MaxYAlphanumericResponse message in English

Example : successful

originalReferenceNoString (36)MaxY AlphanumericTransaction ID provided by BCA

Taken from response API Payment QRIS CPM 'referenceNo' fieldoriginalPartnerReferenceNoString (64)MaxYAlphanumericTransaction ID provided by the partner

Format : include character ‘-‘originalExternalIdString (36) Fixed Y Numeric Original External-ID on header message requestcancelTimeString (25)Max N yyyy-MM-dd HH:mm:ssTZDCancel time

Example: 2024-09-01T17:56:57+07:00

Mandatory applied only if the cancel transaction is successfultransactionDateString (25)MaxNyyyy-MM-ddHH:mm:ssTZDTransaction date

Example: 2024-09-01T10:30:24+07:00

Mandatory applied only if the request from merchant is valid and the transaction is successful additionalInfoObject N Additional Information   merchantIdString (15)MaxNNumericMerchant identifier that is unique per each merchant

For merchant aggregator, filled with submerchantID

Format : 000 885 + MID (9 digit)

Mandatory applied only if the request from merchant is valid   terminalIdString (8)MaxNAlphanumeric Terminal ID

Mandatory applied only if the request from merchant is valid  referenceNumberString (12)FixedNAlphanumeric Unique ID for tracing

Mandatory applied only if the request from merchant is valid

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message**Condition**4004006200Bad requestGeneral request failed error4004006201Invalid Field Format {field}The data you entered does not match the format and length requirement4004006202Invalid Mandatory Field {field}Mandatory field should be fulfilled4014016200Unauthorized. \[Reason\]Invalid signature/Unknown client/Connection not allowed4014016201Invalid token (B2B)Access Token Not Exist/Access Token Expired/Token is Invalid4034036201Feature Not AllowedThe feature is not allowed to use. Please make sure you registered this feature in your partnership data.4034036206Feature Not Allowed At This TimeThe feature is not allowed at this time. Please try again4034036215 Transaction Not Permitted. \[Reason\]Payment transaction status has failed/Payment transaction status has been settled 4044046201Transaction Not FoundTransaction Not Found4044046208Invalid MerchantInvalid Merchant4044046218Inconsistent Request \[Reason\]merchantID and X-PARTNER-ID value are not match. These fields should consist of same value4094096200ConflictX-EXTERNAL-ID duplicate4294296200Too Many RequestsMaximum transaction limit exceeded5005006200General ErrorWrong request field input5005006201Internal Server ErrorUnknown internal server failure, please retry the process again5045046200Timeout

Your transaction timeout

POST /openapi/v1.0/qr/qr-cpm-cancel

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
curl -X POST https://sandbox.bca.co.id/openapi/v1.0/qr/qr-cpm-cancel
-H 'Content-type: application/json'
-H 'X-TIMESTAMP: 2024-09-01T09:05:23+07:00'
-H 'Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'ORIGIN: www.hostname.com'
-H 'X-PARTNER-ID: 123456789'
-H 'X-EXTERNAL-ID: 41807553358950093184162180797837'
-H 'CHANNEL-ID: 95251'
-d '
{
 "originalPartnerReferenceNo": "2024090100000000000001",
 "merchantId": "123456789",
 "subMerchantId": "",
 "additionalInfo": {
 "terminalId": "A1234567",
 "partnerMerchantType": ""
 }
}'

```

POST /openapi/v1.0/qr/qr-cpm-cancel

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
Merchant Regular - Success

H 'Content-type: application/json'
-H 'X-TIMESTAMP: 2024-09-01T09:05:26+07:00'
{
 "responseCode": "2006200",
 "responseMessage": "Success",
 "originalReferenceNo": "43057a74-e179-49ed-b119-b004a3d522ec",
 "originalPartnerReferenceNo": "2024090100000000000001",
 "originalExternalId": "41807553358950093184162180797838",
 "cancelTime": "2024-09-01T09:05:25+07:00",
 "transactionDate": "2024-09-01T08:43:11+07:00",
 "additionalInfo": {
 "merchantId": "000885123456789",
 "terminalId": "A1234567",
 "referenceNumber": "022411000106",
 }
}

Merchant Regular - Failed

-H 'Content-type: application/json'
-H 'X-TIMESTAMP: 2024-09-01T09:05:26+07:00'
{
 "responseCode": "4036215",
 "responseMessage": "Transaction Not Permitted [Transaction status : settled]",
 "originalReferenceNo": "43057a74-e179-49ed-b119-b004a3d522ec",
 "originalPartnerReferenceNo": "2024090100000000000001",
 "originalExternalId": "41807553358950093184162180797838",
 "cancelTime": "",
 "transactionDate": "2024-09-01T08:43:11+07:00",
 "additionalInfo": {
 "merchantId": "000885123456789",
 "terminalId": "A1234567",
 "referenceNumber": "022411000106"
 }
}

```

##### 2\. SNAP QR CPM Payment

This service is used to process payment for CPM QR Codes.

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString (5) FixedY Channel’s identifier using EDC’s WSID (95251)X-PARTNER-IDHeaderString (36) MaxY Unique ID for merchant

\- Merchant Reguler: MerchantID (9)

\- Merchant Partnership : Partner Name (36)X-EXTERNAL-ID

HeaderString (36) MaxYID that should be unique per co-partner per service in the same day

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString (64)Max YAlphanumeric Transaction ID provided by the partner Format : include character ‘-‘ qrContentString (512)MaxYAlphanumeric QR String CPM Format : include character ‘+‘ , ‘=’, and ‘/’ amountObjectYMandatory applied only if the transaction is successful     valueString (13.2)MaxYNumericNet amount of the transaction (exclude convenience fee)

Example : IDR 10.000,- will be placed with 10000.00      currencyString (3)FixedYAlphanumericCurrency code (IDR)merchantIdString (9)MaxYNumericMerchant identifier that is unique per each merchantsubMerchantIdString (9)MaxCAlphanumeric

For merchant aggregator, filled with submerchantIDacquirerNameString (3)MaxYAlphanumericFilled with BCAterminalIdString (8)

MaxYAlphanumericTerminal ID Can be :

\- TID generated by BCA (merchant reguler)

\- TID generated by partner (merchant partnership) additionalInfoObjectNAdditional Information     convenienceFeeString (10.2)MaxNNumericConvenience fee or tips      partnerMerchantTypeString(4)Max CAlphanumericMandatory applied to merchant aggregator, filled with partner identifier

Result of the request will contains following information:

**Response:**

FieldDataType**Length Type****Mandatory****Format**DescriptionresponseCode

String (7) FixedYNumeric

Response code to identify payment request status, refer to the list of response code.

Response code is different from transaction status ('latestTransactionStatus' field)

Format: AAABBCC

AAA : HTTP Code

BB : Service Code

CC : Case Code

Service Code Notification Payment QRIS CPM : 60

Example: 2006000

responseMessage

String (150)

MaxYAlphanumericResponse message in English

Example : successful

referenceNoString (36)MaxN AlphanumericTransaction ID provided by BCA

Mandatory applied only if the request from merchant is validpartnerReferenceNoString (64)MaxYAlphanumericTransaction ID provided by the partnertransactionDateString (25)MaxNyyyy-MM-ddHH:mm:ssTZDTransaction date

Example: 2024-09-01T10:30:24+07:00

Mandatory applied only if the transaction is successfuladditionalInfoObject N Additional Information   latestTransactionStatusString (2)  FixedY Numeric Transaction Status

00 - Success : transaction is paid

06 -Failed : transaction failed, because QR Expired  transactionStatusDesc

String (50) Max Y Alphanumeric Show description of transaction status

Example:

\- Success (transaction Success)

\- Failed (QR Expired) amountObjectY      valueString (13.2) Max Y Numeric Net amount of the transaction

(exclude convenience fee)

Examlple: IDR 10.000,- will be placed with 10000.00      currencyString (3)Fixed Y Alphanumeric Currency code (IDR)   convenienceFeeString (10.2)MaxNNumericConvenience fee or tips  terminalIdString (8)MaxNAlphanumeric Terminal ID

Mandatory applied only if the request from merchant is valid referenceNumberString (12)FixedNAlphanumeric Unique ID for tracing

Mandatory applied only if the request from merchant is valid approvalCodeString (6) FixedNAlphanumericApproval code from host for success transaction

Mandatory applied only if the transaction is successful payerPhoneNumberString (13)MaxNAlphanumericMasked buyer phone number batchNumberString (6) FixedNNumericTransaction batch number Mandatory applied only if the transaction is successful customerPanString (19) MaxNNumericBuyer card number (customer\_pan)

Mandatory applied only if the request from merchant is valid issuerReferenceNumberString (12)MaxNAlphanumericRRN QRIS

Mandatory applied only if the transaction is successfu payerNameString (30)MaxN AlphanumericBuyer name

Mandatory applied only if the request from merchant is valid issuerNameString (26)MaxNAlphanumericIssuer name Mandatory applied only if the transaction is successful  acquirerNameString (3)FixedNAlphanumericFilled with 'BCA' Mandatory applied only if the request from merchant is valid paymentChannelNameString (10)MaxNAlphanumeric\- Debit

\- Switching

Mandatory applied only if the transaction is successful merchantInfoObjectN      merchantIdString (15) MaxNNumericMerchant identifier that is unique per each merchant

Format : 000 885 + MID (9 digit)

Mandatory applied only if the request from merchant is valid     merchantPanString (19) MaxNNumericMerchant PAN

Mandatory applied only if the request from merchant is valid     nameString (25)MaxN AlphanumericMerchant name     cityString (13)MaxNAlphanumericMerchant city

Mandatory applied only if the request from merchant is valid     postalCodeString (10)MaxNNumericMerchant postal code

Mandatory applied only if the request from merchant is valid     countryString (2)FixedNAlphanumericMerchant location country

Mandatory applied only if the request from merchant is valid     emailString (30)MaxNAlphanumericMerchant e-mail address

The following is the data that must be informed on the proof of the transaction to the customer:

-transactionDate

-latestTransactionStatus

-amount

-convenienceFee

-customerPan

-issuerReferenceNumber

-payerName

-issuerName

-acquirerName

-merchantPan

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message**Condition**4004006000Bad requestGeneral request failed error4004006001Invalid Field Format {field}The data you entered does not match the format and length requirement4004006002Invalid Mandatory Field {field}Mandatory field should be fulfilled4014016000Unauthorized. \[Reason\]Invalid signature/Unknown client/Connection not allowed4014016001Invalid token (B2B)Access Token Not Exist/Access Token Expired/Token is Invalid4034036000Transaction ExpiredQR has expired4034036001Feature Not AllowedThe feature is not allowed to use. Please make sure you registered this feature in your partnership data.4034036002Exceeds Transaction Amount LimitExceeds Transaction Amount Limit4034036006Feature Not Allowed At This TimeThe feature is not allowed at this time. Please try again4034036015 Transaction Not Permitted. \[Reason\]Transaction Not Permitted. Please try again4044046001Transaction Not FoundTransaction Not Found4044046008Invalid MerchantInvalid Merchant4044046014Paid BillQR already used 4044046018Inconsistent Request \[Reason\]merchantID and X-PARTNER-ID value are not match. These fields should consist of same value4094096000ConflictX-EXTERNAL-ID duplicate4094096001Duplicate partnerReferenceNopartnerReferenceNo duplicate4294296000Too Many RequestsMaximum transaction limit exceeded5005006000General ErrorWrong request field input5005006001Internal Server ErrorUnknown internal server failure, please retry the process again5045046000Timeout

Your transaction timeout

POST /openapi/v1.0/qr/qr-cpm-payment

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
curl -X POST https://sandbox.bca.co.id/openapi/v1.0/qr/qr-cpm-payment
-H 'Content-type: application/json'
-H 'Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-TIMESTAMP: 2024-09-01T08:43:10+07:00'
-H 'X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'ORIGIN: www.hostname.com'
-H 'X-PARTNER-ID: 123456789'
-H 'X-EXTERNAL-ID: 41807553358950093184162180797837'
-H 'CHANNEL-ID: 95251'
-d '
{
 "partnerReferenceNo": "2024090100000000000001",
 "qrContent": "hduvY...",
 "amount": {
 "value": "50000.00",
 "currency": "IDR"
 },
 "merchantId": "123456789",
 "subMerchantId": "",
 "acquirerName": "BCA",
 "terminalId": "A1234567",
 "additionalInfo": {
 "convenienceFee": "0.00",
 "partnerMerchantType": ""
 }
}'

```

POST /openapi/v1.0/qr/qr-cpm-payment

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
Transaction Status Success

-H 'Content-Type: application/json'
-H 'X-TIMESTAMP: 2024-09-01T08:43:12+07:00'
{
 "responseCode": "2006000",
 "responseMessage": "Success",
 "referenceNo": "43057a74-e179-49ed-b119-b004a3d522ec",
 "partnerReferenceNo": "2024090100000000000001",
 "transactionDate": "2024-09-01T08:43:11+07:00",
 "additionalInfo": {
 "latestTransactionStatus": "00",
 "transactionStatusDesc": "Success (transaction Success)",
 "amount": {
 "value": "50000.00",
 "currency": "IDR"
 },
 "convenienceFee": "0.00",
 "terminalId": "A1234567",
 "referenceNumber": "022411000106",
 "approvalCode": "223356",
 "payerPhoneNumber": "*******4450",
 "batchNumber": "000004",
 "customerPan": "9360001430000131018",
 "issuerReferenceNumber": "000013101103",
 "payerName": "Customer A",
 "issuerName": "BCA",
"acquirerName": "BCA",
 "paymentChannelName": "Debit",
 "merchantInfo": {
 "merchantId": "000885123456789",
 "merchantPan": "9360001431234567899",
 "name": "Merchant XYZ",
 "city": "JAKARTA PUSAT",
 "postalCode": "10310",
 "country": "ID",
 "email": ""
 }
 }
}

Transaction Status Failed

-H 'Content-Type: application/json'
-H 'X-TIMESTAMP: 2024-09-01T08:43:12+07:00'
{
 "responseCode": "2006000",
 "responseMessage": "Successful",
 "referenceNo": "43057a74-e179-49ed-b119-b004a3d522ec",
 "partnerReferenceNo": "2024090100000000000001",
 "transactionDate": "",
 "additionalInfo": {
 "latestTransactionStatus": "06",
 "transactionStatusDesc": "Failed (Do not honor)",
 "amount": {
  "value": "50000.00",
 "currency": "IDR"
 },
 "convenienceFee": "0.00",
 "terminalId": "A1234567",
 "referenceNumber": "022411000106",
 "approvalCode": "000000",
 "payerPhoneNumber": "",
 "batchNumber": "000000",
 "customerPan": "9360001430000131018",
 "issuerReferenceNumber": "",
 "payerName": "Customer A",
 "issuerName": "",
 "acquirerName": "BCA",
 "paymentChannelName": "",
 "merchantInfo": {
 "merchantId": "000885123456789",
 "merchantPan": "9360001431234567899",
 "name": "Merchant XYZ",
 "city": "JAKARTA PUSAT",
 "postalCode": "10310",
 "country": "ID",
 "email": ""
 }
 }
}
```

#### Shared Biller BCA

##### 1\. Shared Biller Transaction Status Inquiry

This service is used for inquiry payment status.

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)FixedY Mandatory in BCA. Channel’ Identifier using WSID Shared Biller (95391)X-PARTNER-IDHeaderString(32)MaxY Mandatory in BCA. Partner ID (Mitra ID)X-EXTERNAL-IDHeaderString(36)MaxYMandatory in BCA. Reference number that should be unique in the same day

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionvirtualAccountNoString (28)MaxYAlphanumericMandatory in BCA. partnerServiceId (8 digit left padding space) + customerNo (up to 20 digit). In BCA, virtualAccountNo max length is 26 (partnerServiceId(8 digits) + customerNo(18 digits)).inquiryRequestIdString (128)MaxNAlphanumericMandatory in BCA. Unique identifier for inquiry. Generated by BCA from inquiry response.

Note:

\- Transaction can be queried between D-day (hari H) until D-1 day (H-1 / yesterday).

\- paymentRequestID is the same value with inquiryRequestID from BCA for each transaction.

\- For inquiry payment status, can use just one keyword (customerNo or inquiryRequestId/paymentRequestID) as filter criteria.

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)FixedYNumericMandatory in BCA. Response code: HTTP Code + Service Code + Case Code responseMessageString (150)MaxYAlphanumericMandatory in BCA. Response descriptionpaymentRequestIdString (128)MaxNAlphanumericMandatory in BCA. Unique identifier for payment. Generated by BCA and has the same value with inquiryRequestI d.virtualAccountNoString (28)MaxYAlphanumericMandatory in BCA. partnerServiceId (8 digit left padding space) + customerNo (up to 20 digit) In BCA, virtualAccountNo max length is 9 (partnerServiceId(8 digits)+customerNo(18 digits)).paymentFlagReasonObject-N-Mandatory in BCA. Reason for Payment Status from Biller in multi languageenglishString (64) MaxNAlphanumericReason for inquiry status in EnglishindonesiaString (64) MaxNAlphanumericReason for inquiry status in BahasainquiryRequestIdString (128)MaxNAlphanumericMandatory in BCA. Inquiry identifier for inquiry. Generated by BCA from inquiry response. paidAmountObjects-N-Mandatory in BCA.valueString (16,2)MaxYDecimalPaid amount with 2 decimal. In BCA, totalAmount max length is 13.2 digitscurrencyString (3)FixedYISO-4217Currency of amount based on ISO 4217 trxDateTimeDate(25)FixedNISO-8601Mandatory in BCA. BCA internal system datetime with timezone, which follows the ISO-8601 standard transactionDateDate(25)FixedNISO-8601Mandatory in BCA. Payment datetime when the payment happenedpaymentFlagStatusString (2)FixedNNumericMandatory in BCA. Status for payment flag, This field may vary with these values: 00 : success transaction 01 : transaction in progress 99 : failed transactionbillDetailsArray of Objects-N-Array with maximum 24 ObjectsbillNoString (18)MaxNAlphanumericBill number that customer choose to pay. From Payment Request billDescriptionObject-N-Bill DescriptionenglishString (18)MaxNAlphanumericBill Description in EnglishindonesiaString (18)MaxNAlphanumericBill Description in BahasabillSubCompanyString (5)MaxNAlphanumericBiller’s product codebillAmountObjects-N-valueString (16,2)MaxYDecimalTransaction amount for specific bill. Nominal inputted by Customer with 2 decimal In BCA, billAmount max length is 13.2 digits.currencyString (3)FixedYISO-4217CurrencyadditionalInfoObject-N-Additional Information for customer use for each billfreeTextsArray of Objects-N-Array with maximum 9 Objects send by BCA.englishString (32)MaxNAlphanumericWill be shown in Channel indonesiaString (32)MaxNAlphanumericWill be shown in Channel

Note:

\- The final status of payment flag is not determined by responseCode and responseMessage.

\- Status for payment flag will be defined in paymentFlagStatus and paymentFlagReason field. This field may vary with these values:

00 = success transaction

01 = transaction in progress

99 = failed transaction

\- If the paymentFlagStatus is 01, the result will be defined after reconciliation process.

\- If the reconciliation results show that the transaction was successfully settled, the paymentFlagStatus will contains 00.

\- If the reconciliation results show a refund transaction, the paymentFlagStatus will contains 99.

\- ResponseCode 4042601 shows request inquiry status is failed, because the transaction data is not available/not recorded at BCA, the paymentFlagStatus will be empty.

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message2002002600Successful4004002601Invalid Field Format {Field Name}4004002602Invalid Mandatory Field {Field Name}4004002600Bad request 4014012600Unauthorized. \[Reason\]4014012601Invalid token (B2B)4034032601Feature Not Allowed4044042601Transaction Not Found4044042616Partner Not Found4094092600Conflict5005002600General Error5005002601Internal Server Error

/openapi/shared-biller/v1.0/transfer-va/status

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
curl –X POST https://sandbox.bca.co.id/openapi/shared-biller/v1.0/transfer-va/status
-H 'CHANNEL-ID': '95391'
-H 'X-PARTNER-ID': '676767'
-H 'X-EXTERNAL-ID': '123456789'
-d '
{
	"partnerServiceId": "03013",
	"customerNo": "123456789012345678",
	"virtualAccountNo": "00003013123456789012345678",
	"inquiryRequestId": "202209161033070301300001197441"
} '
```

/openapi/shared-biller/v1.0/transfer-va/status

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
	"responseCode": "2002600",
	"responseMessage": "Successful",
	"virtualAccountData": {
		"paymentFlagReason": {
			"english": "Transaction Success",
			"indonesia": "Transaksi Berhasil"
		},
		"partnerServiceId": " 21180",
		"customerNo": "22203",
		"virtualAccountNo": " 2118022203",
		"inquiryRequestId": "202307311953292118000043785288",
		"paymentRequestId": "202307311953292118000043785288",
		"paidAmount": {
			"value": "803000.00",
			"currency": "IDR"
		},
		"paidBills": "",
		"totalAmount": {
			"value": "",
			"currency": ""
		},
		"trxDateTime": "2023-07-31T19:57:18+07:00",
		"transactionDate": "2023-07-31T19:53:29+07:00",
		"referenceNo": "",
		"paymentType": "",
		"flagAdvise": "",
		"paymentFlagStatus": "00",
		"billDetails": [
		{
			"billCode": "",
			"billNo": "00011114",
			"billName": "",
			"billShortName": "",
			"billDescription": {
			"english": "Pencil",
			"indonesia": "Pensil"
			},
			"billSubCompany": "00003",
			"billAmount": {
				"value": "803000.00",
				"currency": "IDR"
			},
			"additionalInfo": "",
			"billReferenceNo": "",
			"status": "",
			"reason": {
				"english": "",
				"indonesia": ""
			}
		}
		],
		"freeTexts": [
		{
			"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 1",
			"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 1"
		},
		{
			"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 1",
			"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 1"
		},
		{
			"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 1",
			"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 1"
		},
		{
			"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 1",
			"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 1"
		},
		{
			"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 1",
			"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 1"
		},
		{
			"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 1",
			"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 1"
		},
		{
			"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 1",
			"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 1"
		},
		{
			"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 1",
			"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 1"
		},
		{
			"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 1",
			"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 1"
		}
		],
		"additionalInfo": ""
	}
}
```

##### 2\. SNAP Shared Biller Inquiry Payment to VA from Intrabank

This service is used for billing inquiry to biller.

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)FixedY Mandatory in BCA. Channel’ Identifier using WSID Shared Biller (95391)X-PARTNER-IDHeaderString(32)MaxY Mandatory in BCA. Partner ID (Mitra ID)X-EXTERNAL-IDHeaderString(36)MaxYMandatory in BCA. Reference number that should be unique in the same day

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString (128)MaxNAlphanumericMandatory in BCA. Unique identifier for this Payment. Generated by Mitra (channel-transactionID)virtualAccountNoString (28)MaxYAlphanumericMandatory in BCA. partnerServiceId (8 digit left padding space) + customerNo (up to 20 digit). In BCA, virtualAccountNo max length is 26 (partnerServiceId(8 digits) + customerNo(18 digits)).trxDateTimeDate(25)MaxNISO-8601Mandatory in BCA. Mitra internal system datetime with timezone, which follows the ISO-8601 standardadditionalInfoObject-N-Optional in BCA. Additional Information for custom use

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)FixedYNumericMandatory in BCA. Response code: HTTP Code + Service Code + Case Code responseMessageString (150)MaxYAlphanumericMandatory in BCA. Response descriptionvirtualAccountDataObject-Y-partnerReferenceNoString (128)MaxNAlphanumericMandatory in BCA. From Inquiry Request (channeltransaction-ID)virtualAccountNoString (28)MaxYAlphanumericMandatory in BCA. partnerServiceI d (8 digit left padding space) + customerNo (up to 20 digit) In BCA, virtualAccountNo max length is 26. 9 (partnerServiceId(8 digits)+customerNo(18 digits)).virtualAccountNameString (255)MaxYAlphanumericMandatory in BCA. Customer nameinquiryRequestIdString (128)MaxNAlphanumericMandatory in BCA. Unique identifier for inquiry. Generated by BCA.totalAmountObjects-N-valueString (16,2)MaxYDecimalMandatory in BCA. Send by BCA. Total amount with 2 decimal In BCA, totalAmount max length is 13.2 digits.currencyString(3)FixedYISO-4217Mandatory in BCA. Currency of amount Send by BCA. Currency of amount based on ISO 4217. Currency only available for IDRbillDetailsArray of Objects-N-Optional in BCA. Array with maximum 24 ObjectsbillNoString (18)MaxNAlphanumericBill number from BillerbillDescriptionObject-N-Send by BCA. Bill DescriptionenglishString (18)MaxNAlphanumericSend by BCA. Bill Description in EnglishindonesiaString(18)MaxNAlphanumericSend by BCA. Bill Description in BahasabillSubCompanyString(5)MaxNAlphanumericBiller’s product code billAmountObject-N-valueString (16,2)MaxYDecimalSend by BCA. Transaction amount for specific bill. Nominal inputted by Customer with 2 decimal In BCA, billAmount max length is 13.2 digitscurrencyString (3)FixedYISO-4217CurrencyadditionalInfoObject-N-Additional Information for custom use for each billfreeTextsArray of Objects-N-Optional in BCA. Array with maximum 5 Objects send by BCA.englishString (32) MaxNAlphanumericWill be shown in Channel indonesiaString (32) MaxNAlphanumericWill be shown in Channel virtualAccountTrxTypeString(1)FixedNAlphanumericMandatory in BCA. Type of Virtual Account.

BCA will response virtualAccountT rxType value only for one of these value: - O - Open Payment - C - Closed Payment - V - Bill Variable -I - Partial -W- Multi Bill VariableproductNameString(30)MaxNAlphanumericMandatory in BCA. Send by BCA. Product Category

Note:

1\. BCA will response virtualAccountTrxType value only for one of these value:

\- O (Open Payment) = (non Multi Bill - No Bill) Bill amount is not displayed,customer may enter paid amount freely

\- C (Closed Payment) = (non Multi Bill - Fixed Bill)

 Bill amount is displayed, customer must pay according to the displayed amount

 If more than one bill is displayed, customer must pay all displayed bills & total amount for all bills

\- V (Bill Variable) = (non Multi Bill - Variable Bill) Bill amount is displayed, customer may enter paid amount freely

\- I (Partial) = (Flexible Multi Bill - Fixed Bill)

 More than one bill may be displayed, customer may choose more than one of the displayed bills to pay

 Bill amount is displayed, customer must pay according to the displayed amount in each bill

 Total of paid amount is equal to the sum of chosen bills

\- W (Variable Multi Bill) = (Multi Bill - Fixed Bill)

 More than one bill may be displayed, customer may choose only one of the displayed bills to pay

 Bill amount is displayed, customer must pay according to the displayed amount of the chosen bill

2\. Currency only available for IDR

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message2002003200Successful4004003201Invalid Field Format {Field Name}4004003202Invalid Mandatory Field {Field Name}4004003200Bad request 4014013200Unauthorized. \[Reason\]4014013201Invalid token (B2B)4034033201Feature Not Allowed4044043208Invalid Merchant4044043212Invalid Bill. \[Reason\] (Reason is defined by Biller) 4044043214Paid Bills4044043216Partner Not Found4044043218Inconsistent Request4094093200Conflict5005003200General Error5005003201Internal Server Error

/openapi/shared-biller/v1.0/transfer-va/inquiry-intrabank

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
curl –X POST https://sandbox.bca.co.id/openapi/shared-biller/v1.0/transfer-va/inquiry-intrabank
-H 'CHANNEL-ID': '95391'
-H 'X-PARTNER-ID': '676767'
-H 'X-EXTERNAL-ID': '123456789'
-d '
{
 "partnerServiceId": " 03013",
"customerNo": "123456789012345678",
"partnerReferenceNo": "1122334455",
"virtualAccountNo": " 00003013123456789012345678",
 "trxDateTime": "2022-09-16T10:33:07+07:00"
}'
```

/openapi/shared-biller/v1.0/transfer-va/inquiry-intrabank

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
	"responseCode": "2003200",
	"responseMessage": "Successful",
	"virtualAccountData": {
		"inquiryStatus": "",
		"inquiryReason": {
			"english": "",
			"indonesia": ""
		},
		"partnerServiceId": " 21180",
		"partnerReferenceNo": "90903200114",
		"customerNo": "22203",
		"virtualAccountNo": " 2118022203",
		"virtualAccountName": "Customer Name",
		"virtualAccountEmail": "",
		"virtualAccountPhone": "",
		"sourceAccountNo": "",
		"sourceAccountType": "",
		"inquiryRequestId": "202307311953292118000043785288",
		"totalAmount": {
			"value": "8045000.00",
			"currency": "IDR"
		},
		"billDetails": [
		{
			"billCode": "",
			"billNo": "00011111",
			"billName": "",
			"billShortName": "",
			"billDescription": {
				"english": "Pencil",
				"indonesia": "Pensil"
			},
			"billSubCompany": "00000",
			"billAmount": {
				"value": "800000.00",
				"currency": "IDR"
			},
			"billAmountLabel": "",
			"billAmountValue": "800000.00",
			"additionalInfo": []
		},
		{
			"billCode": "",
			"billNo": "00011112",
			"billName": "",
			"billShortName": "",
			"billDescription": {
				"english": "Book",
				"indonesia": "Buku"
			},
			"billSubCompany": "00001",
			"billAmount": {
				"value": "801000.00",
				"currency": "IDR"
			},
			"billAmountLabel": "",
			"billAmountValue": "801000.00",
			"additionalInfo": []
		},
		{
			"billCode": "",
			"billNo": "00011113",
			"billName": "",
			"billShortName": "",
			"billDescription": {
				"english": "Bag",
				"indonesia": "Tas"
			},
			"billSubCompany": "00002",
			"billAmount": {
				"value": "802000.00",
				"currency": "IDR"
			},
			"billAmountLabel": "",
			"billAmountValue": "802000.00",
			"additionalInfo": []
		},
		{
			"billCode": "",
			"billNo": "00011114",
			"billName": "",
			"billShortName": "",
			"billDescription": {
				"english": "Pencil",
				"indonesia": "Pensil"
			},
			"billSubCompany": "00003",
			"billAmount": {
				"value": "803000.00",
				"currency": "IDR"
			},
			"billAmountLabel": "",
			"billAmountValue": "803000.00",
			"additionalInfo": []
		},
		{
			"billCode": "",
			"billNo": "00011115",
			"billName": "",
			"billShortName": "",
			"billDescription": {
				"english": "Book",
				"indonesia": "Tas"
			},
			"billSubCompany": "00004",
			"billAmount": {
				"value": "804000.00",
				"currency": "IDR"
			},
			"billAmountLabel": "",
			"billAmountValue": "804000.00",
			"additionalInfo": []
		},
		{
			"billCode": "",
			"billNo": "00011110",
			"billName": "",
			"billShortName": "",
			"billDescription": {
				"english": "Bag",
				"indonesia": "Pensil"
			},
			"billSubCompany": "00005",
			"billAmount": {
				"value": "805000.00",
				"currency": "IDR"
			},
			"billAmountLabel": "",
			"billAmountValue": "805000.00",
			"additionalInfo": []
		},
		{
			"billCode": "",
			"billNo": "00011116",
			"billName": "",
			"billShortName": "",
			"billDescription": {
				"english": "Pencil",
				"indonesia": "Pencil"
			},
			"billSubCompany": "00006",
			"billAmount": {
				"value": "806000.00",
				"currency": "IDR"
			},
			"billAmountLabel": "",
			"billAmountValue": "806000.00",
			"additionalInfo": []
		},
		{
			"billCode": "",
			"billNo": "00011117",
			"billName": "",
			"billShortName": "",
			"billDescription": {
				"english": "Book",
				"indonesia": "Buku"
			},
			"billSubCompany": "00007",
			"billAmount": {
				"value": "807000.00",
				"currency": "IDR"
			},
			"billAmountLabel": "",
			"billAmountValue": "807000.00",
			"additionalInfo": []
		},
		{
			"billCode": "",
			"billNo": "00011118",
			"billName": "",
			"billShortName": "",
			"billDescription": {
				"english": "Bag",
				"indonesia": "Tas"
			},
			"billSubCompany": "00008",
			"billAmount": {
				"value": "808000.00",
				"currency": "IDR"
			},
			"billAmountLabel": "",
			"billAmountValue": "808000.00",
			"additionalInfo": []
		},
		{
			"billCode": "",
			"billNo": "00011119",
			"billName": "",
			"billShortName": "",
			"billDescription": {
				"english": "Pencil",
				"indonesia": "Pensil"
			},
			"billSubCompany": "00009",
			"billAmount": {
				"value": "809000.00",
				"currency": "IDR"
			},
			"billAmountLabel": "",
			"billAmountValue": "809000.00",
			"additionalInfo": []
		}
		],
		"freeTexts": [
		{
			"english": "free text english",
			"indonesia": "free text indonesia"
		}
		],
		"virtualAccountTrxType": "W",
		"feeAmount": {
			"value": "0.00",
			"currency": "IDR"
		},
		"productName": "Limit VA",
		"additionalInfo": ""
	}
}
```

##### 3\. SNAP Shared Biller Payment to VA from Intrabank

This service is to make payment to biller

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)FixedY Mandatory in BCA. Channel’ Identifier using WSID Shared Biller (95391)X-PARTNER-IDHeaderString(32)MaxY Mandatory in BCA. Partner ID (Mitra ID)X-EXTERNAL-IDHeaderString(36)MaxYMandatory in BCA. Reference number that should be unique in the same day

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString (128)MaxNAlphanumericMandatory in BCA. Unique identifier for this Payment. Generated by Mitra (channel-transactionID)virtualAccountNoString (28)MaxYAlphanumericMandatory in BCA. partnerServiceId (8 digit left padding space) + customerNo (up to 20 digit). In BCA, virtualAccountNo max length is 26 (partnerServiceId(8 digits) + customerNo(18 digits)).inquiryRequestIdString (128)MaxNAlphanumericMandatory in BCA. Inquiry identifier for inquiry. Generated by BCA from inquiry response.paidAmountObjects-Y-valueString (16,2)MaxYDecimalMandatory in BCA. Paid amount with 2 decimal. In BCA, totalAmount max length is 13.2 digits.currencyString (3)FixedYISO-4217Mandatory in BCA. Currency of amount based on ISO 4217 Currency only available for IDRtrxDateTimeDate(25)MaxNISO-8601Mandatory in BCA. Mitra internal system datetime with timezone, which follows the ISO-8601 standardbillDetailsArray of Objects-N-Array with maximum 24 ObjectsbillNoString (18)MaxNAlphanumericBill number that customer choose to paybillDescriptionObject-N-From Inquiry ResponseenglishString (18)MaxNAlphanumericFrom Inquiry ResponseindonesiaString (18)MaxNAlphanumericFrom Inquiry ResponsebillSubCompanyString (5)MaxNAlphanumericFrom Inquiry ResponsebillAmountObjects-N-From Inquiry ResponsevalueString (16,2)MaxYDecimalFrom Inquiry ResponsecurrencyString (3)FixedYISO-4217From Inquiry ResponseadditionalInfoObject-N-From Inquiry ResponsefreeTextArray of Objects-N-From Inquiry Response englishString (32)MaxNAlphanumericFrom Inquiry Response indonesiaString (32)MaxNAlphanumericFrom Inquiry Response

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)FixedYNumericMandatory in BCA. Response code: HTTP Code + Service Code + Case Code responseMessageString (150)MaxYAlphanumericMandatory in BCA. Response descriptionpaymentRequestIdString (128)MaxNAlphanumericMandatory in BCA. Unique identifier for payment. Generated by BCA and has the same value with inquiryRequestI d.partnerReferenceNoString (128)MaxNAlphanumericMandatory in BCA. From Inquiry Request (channeltransaction-ID)virtualAccountNoString (28)MaxYAlphanumericMandatory in BCA. partnerServiceI d (8 digit left padding space) + customerNo (up to 20 digit) In BCA, virtualAccountN o max length is 26 Tech. Doc. OpenAPI-Shared-Biller API v1.1\| PT. Bank Central Asia, Tbk. 9 (partnerServiceI d(8 digits)+custome rNo(18 digits)).billDetailsArray of Objects-N-Array with maximum 24 ObjectsbillNoString (18)MaxNAlphanumericBill number that customer choose to paybillDescriptionObject-N-From Inquiry ResponseenglishString (18)MaxNAlphanumericFrom Inquiry ResponseindonesiaString (18)MaxNAlphanumericFrom Inquiry ResponsebillSubCompanyString (5)MaxNAlphanumericFrom Inquiry ResponsebillAmountObjects-N-From Inquiry ResponsevalueString (16,2)MaxYDecimalFrom Inquiry ResponsecurrencyString (3)FixedYISO-4217From Inquiry ResponseadditionalInfoObject-N-From Inquiry ResponsefreeTextsArray of Objects-N-Optional in BCA. Array with maximum 9 Objects send by BCA.englishString (32) MaxNAlphanumericWill be shown in Channel indonesiaString (32) MaxNAlphanumericWill be shown in Channel

Note:

1\. currency only available for IDR

2\. paymentRequestID is the same value with inquiryRequestID from BCA for each transaction.

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message2002003300Successful2022023300Request In Progress4004003301Invalid Field Format {Field Name}4004003302Invalid Mandatory Field {Field Name}4004003300Bad request 4014013300Unauthorized. \[Reason\]4014013301Invalid token (B2B)4034033201Feature Not Allowed4034033302Exceeds Transaction Amount Limit4034033316Suspend Transaction4044043208Invalid Merchant4044043313Invalid Amount4044043212Invalid Bill. \[Reason\] (Reason is defined by Biller) 4044043214Paid Bills4044043216Partner Not Found4044043218Inconsistent Request4094093200Conflict5005003200General Error5005003201Internal Server Error

/openapi/shared-biller/v1.0/transfer-va/payment-intrabank

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
curl –X POST https://sandbox.bca.co.id/openapi/shared-biller/v1.0/transfer-va/payment-intrabank
-H 'CHANNEL-ID': '95391'
-H 'X-PARTNER-ID': '676767'
-H 'X-EXTERNAL-ID': '123456789'
-d '
{
	"partnerServiceId": " 21180",
	"customerNo": "123456789012345678",
	"referenceNo": "12345",
	"virtualAccountNo": " 00003013123456789012345678",
	"inquiryRequestId": "202209161033070301300001197441",
	"partnerReferenceNo": "23123123",
	"paidAmount": {
		"value": "803000.00",
		"currency": "IDR"
	},
	"trxDateTime": "2023-07-31T15:08:24+07:00",
	"billDetails": [
	{
		"billcode": "",
		"billNo": "00011114",
		"billName": "",
		"billShortName": "",
		"billDescription": {
			"english": "Pencil",
			"indonesia": "Pensil"
		},
		"billSubCompany": "00003",
		"billAmount": {
			"value": "803000.00",
			"currency": "IDR"
		},
		"additionalInfo": "",
		"billReferenceNo": ""
	}
	],
	"freeText": [
		{
			"english": "eating",
			"indonesia": "makan"
		}
	],
	"additionalInfo": "alalalala"
}'
```

/openapi/shared-biller/v1.0/transfer-va/payment-intrabank

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
	"responseCode": "2003300",
	"responseMessage": "Successful",
	"virtualAccountData": {
		"paymentFlagReason": {
			"english": "",
			"indonesia": ""
		},
		"partnerServiceId": " 21180",
		"customerNo": "22203",
		"virtualAccountNo": " 2118022203",
		"virtualAccountName": "",
		"virtualAccountEmail": "",
		"virtualAccountPhone": "",
		"sourceAccountNo": "",
		"sourceAccountType": "",
		"inquiryRequestId": "202307311953292118000043785288",
		"paymentRequestId": "202307311953292118000043785288",
		"partnerReferenceNo": "5555614",
		"referenceNo": "",
		"paidAmount": {
			"value": "803000.00",
			"currency": "IDR"
		},
		"paidBills": "",
		"totalAmount": {
			"value": "",
			"currency": ""
		},
		"trxDateTime": "",
		"journalNum": "",
		"paymentType": "",
		"flagAdvise": "",
		"billDetails": [
		{
			"billCode": "",
			"billNo": "00011114",
			"billName": "",
			"billShortName": "",
			"billDescription": {
				"english": "Pencil",
				"indonesia": "Pensil"
			},
			"billSubCompany": "00003",
			"billAmount": {
				"value": "803000.00",
				"currency": "IDR"
			},
			"additionalInfo": "",
			"status": "",
			"reason": {
				"english": "",
				"indonesia": ""
			}
		}
		],
		"freeTexts": [
			{
				"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 1",
				"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 1"
			},
			{
				"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 2",
				"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 2"
			},
			{
				"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 3",
				"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 3"
			},
			{
				"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 4",
				"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 4"
			},
			{
				"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 5",
				"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 5"
			},
			{
				"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 6",
				"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 6"
			},
			{
				"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 7",
				"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 7"
			},
			{
				"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 8",
				"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 8"
			},
			{
				"english": "ABCDEFGHIJKLM ENGLISH FREETEXT 9",
				"indonesia": "ABCDEFGHIJ INDONESIAN FREETEXT 9"
			}
		],
		"feeAmount": {
			"value": "",
			"currency": ""
		},
		"productName": "",
		"additionalInfo": ""
	}
}
```

#### Transfer Dana

##### 1\. SNAP Banking External Account Inquiry

This service is used to inquiry account external.

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)FixedY Channel ID using KlikBCA Bisnis's WSID (95051)X-PARTNER-IDHeaderString(10)FixedY Corporate ID X-EXTERNAL-IDHeaderString(32)MaxYNumeric String.

Reference number that should be unique in the same day

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString (64)MaxY-Mandatory in BCA.

Transaction identifier on service consumer systembeneficiaryBankCodeString (8)MaxYAlphanumericMandatory in BCA. Beneficiary Bank Code.

Beneficiary bank code. In BCA using SWIFT Code or 3 digit bank code (if destination doesn't have SWIFT Code). Beneficiary Bank Code can be checked on KBB menu: Miscellaneous > Information > Bank Code ListbeneficiaryAccountNoString (34)MaxYAlphanumericMandatory in BCA.

Beneficiary Account

If inquiryService = 1 (Switcher) max length = 28

If inquiryService = 2 (BI-FAST) max length = 34 additionalInfoObjectY  sourceAccountNoString (10)FixedCNumericMandatory in BCA if inquiry using BI-FAST

Source Account  inquiryServiceString (1)FixedYNumericFill in according to the transfer transaction that will be used

1: Inquiry via Switcher

2: Inquiry via BI-FASTamountobjectC   valueString (13,2)MaxCDecimalThis field conditional and will mandatory if inquiryService = 2

(BI-FAST)

Net amount of the transaction. if it's IDR then value includes 2

decimal digits. In BCA, max length = 13.2

currencyString(3)FixedCISO-4217This field conditional and will mandatory if inquiryService = 2 (BI-FAST)

Currency (IDR)purposeCodeString (5)MaxCNumericThis field conditional and will mandatory if inquiryService = 2 (BI-FAST)

Purpose Code

01 = Investment / Investasi

02 = Transfer of Wealth / Pemindahan Dana

03 = Purchase / Pembelian

99 = Others / Lainnya

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString (64)MaxY-Send by BCA according to the request input.

Transaction identifier on service consumer systemresponseCodeString (7)FixedYAlphanumericResponse code to identify transaction status.

Format : AAABBCC

AAA : HTTP

BB : Service Code

CC : Case Code

Example: 2001600responseMessageString (150)MaxYAlphanumericResponse descriptionreferenceNoString (64)MaxCNumericSend by BCA.

Transaction identifier on service provider system. Must be filled upon successful transaction

Unique each day from BCA. Must be filled upon successful transaction.benefiaryAcco untNameString (100) MaxY-Send by BCA.

Beneficiary Account NamebeneficiaryAccountNoString (34) MaxY-Send by BCA according to the request input.

Beneficiary account numberbeneficiaryBankCodeString (8)MaxY-Send by BCA according to the request input.

Beneficiary Bank Code. In BCA using SWIFT Code or 3 digit bank code (if destination doesn’t have SWIFT Code)

Beneficiary Bank Code can be checked on KBB menu:

Miscellaneous > Information > Bank Code Lis

Here is the list of error codes that can be returned.

HTTP CodeError CodeError MessageCondition2002001600SuccessfulTransaction successful.4004001601Invalid Field Format {field}The data you entered does not match the format requirements. 4004001602Invalid Mandatory Field {Field}Mandatory field should be fulfilled.4014011600Unauthorized. \[Reason\]Invalid Signature/Unknown Client/Connection Not Allowed. 4014011601Invalid token (B2B)Access Token Not Exist/Access Token Expired/Token is Invalid.4034031601Feature Not AllowedFeature Not Allowed. Please make sure you registered this feature in your partnership data.4044041602Invalid RoutingYour transaction amount exceeds your company limit4044041611Invalid Account The account used in transaction not valid4094091600ConflictCannot use same X-EXTERNAL-ID in the same day. 5005001600General ErrorGeneral Error.5045041600TimeoutTransaction timeout.

curl –X POST https://sandbox.bca.co.id/openapi/v2.0/account-inquiry-external

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
curl –X POST https://sandbox.bca.co.id/openapi/v2.0/account-inquiry-external
-H 'Content-type:application/json'
-H 'Authorization:Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-TIMESTAMP:2020-12-21T10:30:24+07:00'
-H 'X-SIGNATURE:85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'ORIGIN:www.hostname.com'
-H 'X-PARTNER-ID:KBBABCINDO'
-H 'X-EXTERNAL-ID:28910000006578499987546738976812'
-H 'CHANNEL-ID:95051'
-d '
{
	"beneficiaryBankCode": "BRINIDJA",
	"beneficiaryAccountNo": "888801000157508",
	"partnerReferenceNo": "2020102900000000000001",
	"additionalInfo": {
		"inquiryService": "1"
	}

}'

External Account Inquiry V2 using BI-FAST:

curl –X POST https://sandbox.bca.co.id/openapi/v2.0/account-inquiry-external
-H 'Content-type:application/json'
-H 'Authorization:Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-TIMESTAMP:2020-12-21T10:30:24+07:00'
-H 'X-SIGNATURE:85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'ORIGIN:www.hostname.com'
-H 'X-PARTNER-ID:KBBABCINDO'
-H 'X-EXTERNAL-ID:28910000006578499987546738976812'
-H 'CHANNEL-ID:95051'
-d '
{
	"beneficiaryBankCode": "BRINIDJA",
	"beneficiaryAccountNo": "888801000157508",
	"partnerReferenceNo": "2020102900000000000001",
	"additionalInfo": {
		"sourceAccountNo": "0123456789",
		"inquiryService": "2",
		"amount": {
			"value": "10000.00",
			"currency": "IDR"
		},
		"purposeCode": "02"
	}
}'

```

curl –X POST https://sandbox.bca.co.id/openapi/v2.0/account-inquiry-external

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
	"partnerReferenceNo": "2020102900000000000001",
	"responseCode": "2001600",
	"responseMessage": "Successful",
	"referenceNo": "2020102977770000000009",
	"beneficiaryAccountName": "Yories Yolanda",
	"beneficiaryAccountNo": "888801000157508",
	"beneficiaryBankCode": "BRINIDJA"
}

SNAP Banking External Account Inquiry V2 uses following error sample schema:
-H 'Content-type:application/json'
-H 'X-TIMESTAMP:2020-12-21T10:30:24+07:00'
{
	"partnerReferenceNo": "2020102900000000000001",
	"responseCode": "4001601",
	"responseMessage": "Invalid Field Format beneficiaryAccountNo",
	"referenceNo": "",
	"beneficiaryAccountName": "",
	"beneficiaryAccountNo": "888801000157508",
	"beneficiaryBankCode": "BRINIDJA"
}
```

##### 2\. SNAP Banking Interbank Transfer

**Introduction** With SNAP Banking Interbank transfer, you can perform fund transfer to other domestic bank by using Switcher/BI-FAST service. The source of fund transfer must be from your own account on your KlikBCA Bisnis. SNAP Banking Interbank Transfer transaction amount limit will follow KlikBCA Bisnis’s rules about transaction amount limit in the following menu: Miscellaneous > Information > Transaction Limit Usage > Funds Transfer to Domestic BCA thru LLG, ONLINE and BI FAST Maximum Limit per day per corporate

SNAP Banking Interbank Transfer frequency limit will follow KlikBCA Bisnis’s rules about frequency limit in the following menu: Miscellaneous > Information > Transaction Limit Usage > Funds Transfer to Domestic Bank thru LLG, ONLINE and BI FAST Maximum Number of Transaction per day per corporate

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)FixedY Channel’s identifier using WSID KlikBCA Bisnis (95051) X-PARTNER-IDHeaderString(10)MaxY Corporate ID X-EXTERNAL-IDHeaderString(32)MaxYNumeric String. Reference number that should be unique in the same day.

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString (64)MaxY-
Transaction identifier on service consumer systemamountObjectY-   valueString (13.2)MaxYDecimalNet amount of the transaction.

If it’s IDR then value includes 2 decimal digits.    currencyString (3)FixedYISO-4217
CurrencybeneficiaryAccountNameString (100)MaxY-
Beneficiary Account NamebeneficiaryAccountNoString (34)MaxY-
Beneficiary Account beneficiaryBankCodeString(8)MaxYSWIFT Code/3 digit bank codeBeneficiary Bank Code.

In BCA using SWIFT Code or 3 digit bank code (if destination doesn’t have SWIFT Code).

Beneficiary Bank Code can be checked on KBB menu : Miscellaneous > Information > Bank Code ListbeneficiaryEmailString(50)MaxNabc@domain.com
Beneficiary Email. Only one email address can be submitted.sourceAccountNoString (10)FixedYNumericThis response will be sent according to the request input.

Source Account.transactionDateString (25)MaxYISO-8601Mandatory in BCA.

Transaction dateadditionalInfoObjectY   transferTypeString (1)FixedY1 = Transfer via switcher

2 = Transfer via BIFAST    purposeCodeString (5)MaxNMandatory if transferType = 2

Purpose Code as defined in techdoc BI

01 = Investment/Investasi

02 = Transfer of Wealth/Pemindahan Dana

03 = Purchase/Pembelian

99 = Others/Lainnya

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString(64)MaxNNumericThis response will be sent according to the request input. Transaction identifier on sevice consumer system.responseCodeString (7)FixedYNumericResponse coderesponseMessageString(150)MaxYAlphanumericResponse descriptionreferenceNoString (64)MaxCNumericWill be sent when the transaction is successful as transaction identifier on service provider system.

referenceNo = timestamp(6, yymmdd) + prefix menu(2) + sequence(6).

prefix menu=kode fitur=02.

sequence number unique each day.amountObjectY   valueString (13.2)MaxYDecimalThis response will be sent according to the request input. Net amount of the transaction. If it’s IDR then value includes 2 decimal digits.    currencyString (3)FixedYISO-4217This response will be sent according to the request input. Currency.beneficiaryAccountNoString (34) MaxYNumericThis response will be sent according to the request input. Beneficiary account number. beneficiaryBankCodeString (8)MaxNSWIFT Code/3 digit bank codeThis response will be sent according to the request input. Beneficiary Bank Code. In BCA using SWIFT Code or 3 digit bank code (if destination doesn’t have SWIFT Code). Beneficiary Bank Code can be checked on KBB menu : Miscellaneous > Information > Bank Code ListsourceAccountNoString (19)MaxN-Send by BCA according to the request input.

Source accountadditionalInfobifastIdString (18)FixedCAlphanumericWill be sent when the transaction is successful and transferType = 2.

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message**Condition****Transaction Status**4004001800Bad requestGeneral request failed error.4004001801Invalid Field Format {field}The data you entered does not match the format requirements.4004001802Invalid Mandatory Field {Field Name}Mandatory field should be fulfilled. 4014011800Unauthorized. \[Reason\]Invalid Signature/Unknown Client/Connection Not Allowed.4014011801Invalid token (B2B)Access Token Not Exist/Access Token Expired/Token is Invalid.4034031802Exceeds Transaction Amount LimitYour transaction amount exceeds your company limit. Please check your Company Limit at KlikBCA Bisnis.4034031804Activity Count Limit ExceededToo many request, Exceeds Transaction Frequency Limit. Please check your Company Limit at KlikBCA Bisnis.4034031814Insufficient FundsYour balance in source account less than the amount of transfer you perform.4034031801Feature Not Allowed The feature is not allowed to use. Please make sure you registered this feature in your partnership data. 4034031815Transaction not permittedBI-FAST service for selected Beneficiary Bank is currently unavailable. Please choose another transfer service.4044041803Bank Not Supported By SwitchBank not supported by BI-FAST.4044041811Invalid AccountThe account used in transaction not valid (do not exist/closed/cant be used for transactions).4044041813Invalid Amount Wrong amount input.4044041802Invalid Routing Connection problem to Switcher.4094091800ConflictX-EXTERNAL-ID duplicate.4294291800Too Many RequestsMaximum transaction limit exceeded.5005001800General ErrorWrong request field input. Internal system/external BI-FAST service is currently under maintenance. Please use other BCA transfer services.5005001801Internal Server ErrorUnknown internal server failure, please retry the process again.5045041800TimeoutYour transaction timeout

If you get responseCode 5001801 (Unknown Error/Unknown internal server failure, please retry the process again) from BCA, please make sure to check your account statement before retry the transaction.

curl –X POST https://sandbox.bca.co.id/openapi/v2.0/transfer-interbank

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
Request Interbank Transfer via Switcher Sample
-H 'CHANNEL-ID':'95051'
-H 'X-PARTNER-ID': 'KBBABCINDO'
-H 'X-EXTERNAL-ID': '28910000006578499987546738976812'
-d '
{
	"partnerReferenceNo": "2020102900000000000001",
	"amount": {
		"value": "10000.00",
		"currency": "IDR"
	},
	"beneficiaryAccountName": "Yories Yolanda",
	"beneficiaryAccountNo": "888801000157508",
	"beneficiaryBankCode": "BRINIDJA",
	"beneficiaryEmail": "yories.yolanda@work.bri.co.id",
	"sourceAccountNo": "0123456789",
	"transactionDate": "2020-12-21T14:36:11+07:00"
	"additionalInfo": {
        "transferType": "1",
        "purposeCode": ""
    }
}'

Request Interbank Transfer via BI-FAST Sample

-H 'CHANNEL-ID':'95051'
-H 'X-PARTNER-ID': 'KBBABCINDO'
-H 'X-EXTERNAL-ID': '28910000006578499987546738976812'

-d '
{
	"partnerReferenceNo": "2020102900000000000001",
	"amount": {
	"value": "10000.00",
	"currency": "IDR"
	},
	"beneficiaryAccountName": "Yories Yolanda",
	"beneficiaryAccountNo": "888801000157508",
	"beneficiaryBankCode": "BRINDIJA",
	"beneficiaryEmail": "yories.yolanda@work.bri.co.id",
	"sourceAccountNo": "0123456789",
	"transactionDate": "2020-12-21T14:36:11+07:00",
	"additionalInfo": {
		"transferType": "2",
		"purposeCode": "02"
	}
}
```

curl –X POST https://sandbox.bca.co.id/openapi/v2.0/transfer-interbank

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
Response Sample

{
	"responseCode": "2001800",
	"responseMessage": "Successful",
	"referenceNo": "2020102977770000000009",
	"partnerReferenceNo": "2020102900000000000001",
	"amount": {
		"value": "10000.00",
		"currency": "IDR"
	},
	"beneficiaryAccountNo": "888801000157508",
	"beneficiaryBankCode": "BRINDIJA",
	"sourceAccountNo": "0123456789"
}

Error Sample
-H 'Content-type': 'application/json'
-H 'X-TIMESTAMP': '2020-12-17T13:50:04+07:00'
{
 "responseCode": "4001801",
 "responseMessage": "Invalid Field Format beneficiaryAccountNo",
 "referenceNo": "",
 "partnerReferenceNo": "2020102900000000000001",
 "amount": {
 "value": "10000.00",
 "currency": "IDR"
 },
 "beneficiaryAccountNo": "888801000157508",
 "beneficiaryBankCode": " BRINIDJA",
 "sourceAccountNo": "0123456789"
 "additionalInfo": {
	"transferType": "1",
	"purposeCode": ""
 }

}
```

##### 3\. SNAP Banking Internal Account Inquiry

This service is used to inquiry BCA account.

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)FixedY Channel’s identifier using WSID KlikBCA Bisnis (95051) X-PARTNER-IDHeaderString(10)FixedY KlikBCA Bisnis’s Corporate ID. X-EXTERNAL-IDHeaderString(32)MaxYNumeric String. Reference number that should be unique in the same day.

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString (64)MaxYNumericTransaction identifier on service consumer systems. The partnerReferenceNo used in the inquiry transaction must be reused when transferring to the account inquired.beneficiaryAccountNoString (10)FixedYNumericBeneficiary Account

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)FixedYNumericResponse coderesponseMessageString (150)MaxYAlphanumericResponse descriptionreferenceNoString (64)MaxYNumericThis response will be sent as transaction identifier on service provider system. Must be filled upon successful transaction. partnerReferenceNoString (64)MaxNNumericThis response will be sent according to the request input. Transaction identifier on service consumer systembeneficiaryAccountNameString (100) MaxYAlphanumeric & Special CharacterThis response will be sent according to the request input. Beneficiary Account Name.

The name returned is the name displayed on the channel, which have 24 digit characters. There is no special character validation, but the special characters displayed are: \| @ # $ % ¬ & \* ( ) \_ + - = ¢ ! \ { } \] ; : , . / < > ? \` ^beneficiaryAccountNoString (10) FixedYNumericThis response will be sent according to the request input.

Beneficiary account number.

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message**Condition**4004001500Bad requestGeneral request failed error.4004001501Invalid Field Format {field name}The data you entered does not match the format requirements.4004001502Invalid Mandatory Field {field name}Mandatory field should be fulfilled.4014011500Unauthorized. \[Reason\]Invalid Signature/Unknown Client/Connection Not Allowed.4014011501Invalid token (B2B)Access Token Not Exist/Access Token Expired/Token is Invalid.4034031501Feature Not AllowedThe feature is not allowed to use. Please make sure you registered this feature in your partnership data. 4094091500ConflictX-EXTERNAL-ID duplicate.4034031518Inactive AccountIndicates inactive account4044041511Invalid Account The account used in transaction not valid (do not exist/can’t be used for transactions). 5005001500General ErrorWrong request field input. 5005001501Internal Server Error Unknown internal server failure, please retry the process again.5045041500TimeoutYour transaction timeout.

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/account-inquiry-internal

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'Content-type': 'application/json'
-H 'Authorization': 'Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a'
-H 'X-TIMESTAMP': '2020-12-17T13:50:04+07:00'
-H 'X-SIGNATURE': '85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5'
-H 'ORIGIN': 'www.hostname.com'
-H 'CHANNEL-ID':'95051'
-H 'X-PARTNER-ID': 'KBBABCINDO'
-H 'X-EXTERNAL-ID': '28910000006578499987546738976812'
-d '
{
"partnerReferenceNo": "202010290000000001",
"beneficiaryAccountNo": "8010001575"
}

```

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/account-inquiry-internal

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
-H 'Content-type': 'application/json'
-H 'X-TIMESTAMP': '2020-12-17T13:50:04+07:00'
{
"responseCode": "2001500",
"responseMessage": "Successful",
"referenceNo": "202010297777000009",
"partnerReferenceNo": "202010290000000001",
"beneficiaryAccountName": "Yories Yolanda",
"beneficiaryAccountNo": "8010001575"
}

```

##### 4\. SNAP Banking Intrabank Transfer

This service is used to transfer BCA.

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)FixedY Channel’s identifier using WSID KlikBCA Bisnis (95051) X-PARTNER-IDHeaderString(32)MaxY Corporate ID

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString (64)MaxY-Mandatory in BCA.

Transaction identifier on service consumer system

For transfers that begin with beneficiary account number inquiry using the SNAP Banking Internal Account Inquiry, it is required to use the same partnerReferenceNo with SNAP Banking Internal Account Inquiry that has been done.beneficiaryEmailString (50)MaxNabc@domain.comOptional in BCA.

Beneficiary EmailamountObjectYNumericMandatory in BCA.

Net amount of the transaction   valueString (16.2)MaxYDecimalMandatory in BCA.

If it’s IDR then value includes 2 decimal digits. BCA will send response amount with format value Numeric (13.2)   currencyString (3)FixedYISO-4217Mandatory in BCA.

Account currency typebeneficiaryAccountNoString (34)MaxYNumericMandatory in BCA.

Beneficiary Account Number registered in BCA. For BCA, length account number is 10 digit remarkString (50)MaxNOptional in BCA.

Remark/transac tion description. For BCA, remark is 36 digit divided in 2 field (remark 1 with 18 digit and remark 2 with 18 digit) with no special character sourceAccountNoString (19)MaxYNumericsourceAccountNo registered in BCA.

For BCA, length account number is 10 digit transactionDateString (25)MaxYISODateTime

ISO-8601Mandatory in BCA.

Transaction dateadditionalInfoObjectNAdditional informationeconomicActivityString (60)MaxY

(if the

beneficiaryAccountNumber

are WNA citizenship) -For details, please check on the Economic Activity table belowtransactionPurposeString (2)FixedY

(if the currency of beneficiaryAccountNumber

are in SGD or USD)Numeric01 Pembelian barang / Purchase of Goods

02 Pembelian jasa / Purchase of Service

03 Pembayaran hutang / Debt Payment

04 Pembayaran pajak / Tax Payment

05 Impor / Import

06 Biaya Perjalanan Luar Negeri / Business Travel Expenses

07 Biaya remunerasi pegawai / Remuneration Expenses

08 Biaya administrasi / Administration Fee

09 Biaya overhead / Overhead Expenses

10 Simpanan dalam valas / Saving

11 Investasi Pembelian Obligasi Korporasi / Purchase of Corporate Bonds

12 Investasi Pembelian Surat Berharga Negara (SBN) / Purchase of Government Securities

13 Investasi Penyertaan Langsung / Direct Capital Investment

14 Penambahan modal kerja / Working Capital

15 Telah mendapat ijin BI / Approved by BI

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)FixedY-Response coderesponseMessageString (150)MaxY-Response descriptionreferenceNoString (64)MaxY (When the transaction is successful)NumericSend by BCA.

Transaction identifier on service provider system.

Unique each day from BCA

Must be filled upon successful transaction.partnerReferenceNoString (64)MaxY-Send by BCA according to the request input.

Transaction identifier on service customer system (unique each day from partner)amountObjectY   valueString (16.2)

MaxYDecimalSend by BCA according to the request input.

If it’s IDR then value includes 2 decimal digits. BCA will send response amount with format value Numeric (13.2)   currencyString (3)MaxYISO-4217Send by BCA according to the request input. CurrencybeneficiaryAccountNoString (34) MaxYNumericSend by BCA according to the request input.

Beneficiary account numbersourceAccountNoString (19)MaxYNumericSend by BCA according to the request input.

Source account numbertransactionDateString (25)MaxYISODateTime

ISO-8601Send by BCA according to the request input.

Transaction date additionalInfoObjectNAdditional information   economicActivityString (60)MaxY (if the beneficiaryAccountNumber are WNA citizenship) -Send by BCA according to the request input.

For details, please check on the Economic Activity table below   transactionPurposeString (2)FixedY (if the currency of beneficiaryAccountNumber are in SGD or USD)NumericSend by BCA according to the request input.

01 Pembelian barang / Purchase of Goods

02 Pembelian jasa / Purchase of Service

03 Pembayaran hutang / Debt Payment

04 Pembayaran pajak / Tax Payment

05 Impor / Import

06 Biaya Perjalanan Luar Negeri / Business Travel Expenses

07 Biaya remunerasi pegawai / Remuneration Expenses

08 Biaya administrasi / Administration Fee

09 Biaya overhead / Overhead Expenses

10 Simpanan dalam valas / Saving

11 Investasi Pembelian Obligasi Korporasi / Purchase of Corporate Bonds

12 Investasi Pembelian Surat Berharga Negara (SBN) / Purchase of Government Securities

13 Investasi Penyertaan Langsung / Direct Capital Investment

14 Penambahan modal kerja / Working Capital

15 Telah mendapat ijin BI / Approved by BI

##### Economic Activity Table

Source AccountDestination AccountEconomic ActivityWNIWNAA :

- Biaya Hidup Pihak Asing
- Jual beli barang & jasa di Ind
- Pembukaan SKBDN
- Pembukaan LC Impor dlm Negeri
- Pembukaan LC Impor dlm Negeri Utang LN & restrukturisasi
- Surat Berharga (SSB) dalam Rp
- Penyertaan langsung di Ind

WNAWNAB :

- Pengalihan milik penyertaan lsg
- Surat berharga dalam Rp
- Jual beli barang & jasa di Ind
- Biaya hidup pihak asing

Choose one of the economic activities according to the transaction needs.

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message4004001700Bad request4004001702Invalid Mandatory Field {Field} 4004001701Invalid Field Format {Field}4014011700Unauthorized. \[Reason\]4014011701Invalid token (B2B)4034031701Feature Not Allowed4034031718Inactive Account 4034031702Exceeds Transaction Amount Limit4034031704Activity Count Limit Exceeded4034031714Insufficient Funds4044041713Invalid Amount4044041711Invalid Account 4094091700Conflict4294291700Too Many Requests 5005001700General Error5005001701Internal Server Error5045041700Timeout

If you get responseCode 5001701 (Internal Server Error) from BCA, please make sure to check your account statement before retry the transaction.

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/transfer-intrabank

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95051'
-H 'X-PARTNER-ID': 'KBBABCINDO'
-d '
{
	"partnerReferenceNo": "2020102900000000000001",
	"amount": {
		"value": "10000.00",
		"currency": "IDR"
	},
	"beneficiaryAccountNo": "888801000157508",
	"beneficiaryEmail": "yories.yolanda@work.co.id",
	"remark": "remark test",
	"sourceAccountNo": "888801000157508",
	"transactionDate": "2020-12-21T10:30:24+07:00",
	"additionalInfo": {
		"economicActivity": "Biaya Hidup Pihak Asing",
		"transactionPurpose": "01"
	}
}
```

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/transfer-intrabank

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
	"responseCode": "2001700",
	"responseMessage": "Successful",
	"referenceNo": "2020102977770000000009",
	"partnerReferenceNo": "2020102900000000000001",
	"amount": {
		"value": "10000.00",
		"currency": "IDR"
	},
	"beneficiaryAccountNo": "888801000157508",
	"sourceAccountNo": "888801000157508",
	"transactionDate": "2020-12-21T10:30:24+07:00",
	"additionalInfo": {
		"economicActivity": "Biaya Hidup Pihak Asing",
		"transactionPurpose": "01"
	}
}
```

##### 5\. SNAP Banking Notify RTGS

This service is used to Transfer RTGS Notification.

This feature is not yet available to be accessed via Sandbox.

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)FixedY Channel’ IdentifierX-PARTNER-IDHeaderString(32)MaxY Corporate ID

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionoriginalPartnerReferenceNoString (64)MaxN-Sent by BCA

Original transaction identifier on service consumer systemoriginalReferenceNoString (64)MaxN-Sent by BCA

Original transaction identifier on service provider system originalExternalIdString (19)MaxN-Sent by BCA

Original Customer Reference NumberlatestTransactionStatusString (2)FixedYSent by BCA.

BCA only notifies rejection from transactions that have been successfully processed by BCA

00 - Success

01 - Initiated

02 - Paying

03 - Pendin

04 - Refunded

05 - Canceled

06 - Failed

07 - Not foundamountObjectY   valueString (16.2)MaxYDecimalSent by BCA

Net amount of the transaction. If it's IDR then value includes 2 decimal digits.

In BCA, max length = 13.2   currencyString (3)FixedYISO-4217Sent by BCA.

CurrencybeneficiaryAccountNameString (100)MaxY-Sent by BCA

Beneficiary Account NamebeneficiaryAccountNoString (34)MaxY-Sent by BCA

Beneficiary Account beneficiaryBankCodeString(8)MaxY-Sent by BCA.

Beneficiary Bank Code. In BCA using SWIFT Code or 3 digit bank code (if destination doesn't have SWIFT Code)sourceAccountNoString (19)MaxY-Sent by BCA

Source Account. 10 digits for BCA Account.transactionDateString (25)MaxYISO-8601Mandatory in BCA.

Transaction Date, which follows the ISO-8601 standardadditionalInfoObject-Y-Sent by BCA

Additional Info   rejectStatusIDString (100)MaxY-Sent by BCA

Reject Description in Indonesian    rejectStatusENString (100)MaxY-Sent by BCA

Reject Description in EnglishtraceNoString (100)MaxY-Sent by BCA

Number for tracking to destination bank.

In BCA traceNo=ppu\_num ber with prefix “(“

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)FixedY-Response coderesponseMessageString(150)MaxY-Response description

curl –X POST https://copartners.com/openapi/v1.0/transfer-rtgs/notify

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95051'
-H 'X-PARTNER-ID': 'KBCABCINDO'
-d '
{
 "originalPartnerReferenceNo": "2020102900000000000001",
 "originalReferenceNo": "2020102977770000000009",
 "originalExternalId": "10052019",
 "latestTransactionStatus": "00",
 "amount": {
 "value": "10000.00",
 "currency": "IDR"
 },
 "beneficiaryAccountName": "Yories Yolanda",
 "beneficiaryAccountNo": "888801000003301",
 "beneficiaryBankCode": "002",
 "sourceAccountNo": "1234567890",
 "transactionDate": "2020-12-21T14:06:21+07:000",
 "additionalInfo": {
 "rejectStatusID": "Nomor rekening tidak terdaftar",
 "rejectStatusEN": "Account number is not registered",
 "traceNo": "(DAC8"
}
}'

```

curl –X POST https://copartners.com/openapi/v1.0/transfer-rtgs/notify

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
 "responseCode": "2007600",
 "responseMessage": "Successful"
 }

```

##### 6\. SNAP Banking Notify SKNBI

This service is used to Transfer SKN Notification.

This feature is not yet available to be accessed via Sandbox.

**Additional Headers:**

Field**Params Type**Data Type**Mandatory**DescriptionCHANNEL-IDHeaderString(5)Y Channel’ IdentifierX-PARTNER-IDHeaderString(32)Y Corporate ID

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionoriginalPartnerReferenceNoString (64)MaxN-Sent by BCA

Original transaction identifier on service consumer systemoriginalReferenceNoString (64)MaxN-Sent by BCA

Original transaction identifier on service provider system originalExternalIdString (19)MaxN-Sent by BCA

Original Customer Reference NumberlatestTransactionStatusString (2)FixedYSent by BCA.

BCA only notifies rejection from transactions that have been successfully processed by BCA

00 - Success

01 - Initiated

02 - Paying

03 - Pendin

04 - Refunded

05 - Canceled

06 - Failed

07 - Not foundamountObjectY   valueString (16.2)MaxYDecimalSent by BCA

Net amount of the transaction. If it's IDR then value includes 2 decimal digits.

In BCA, max length = 13.2   currencyString (3)

ISO 4217FixedY-Sent by BCA.

CurrencybeneficiaryAccountNameString (100)MaxY-Sent by BCA

Beneficiary Account NamebeneficiaryAccountNoString (34)MaxY-Sent by BCA

Beneficiary Account beneficiaryBankCodeString(8)MaxY-Sent by BCA.

Beneficiary Bank Code. In BCA using SWIFT Code or 3 digit bank code (if destination doesn't have SWIFT Code)sourceAccountNoString (19)MaxY-Sent by BCA

Source Account. 10 digits for BCA Account.transactionDateString (25)MaxYISO-8601Mandatory in BCA.

Transaction Date, which follows the ISO-8601 standardadditionalInfoObject-Y-Sent by BCA

Additional Info   rejectStatusIDString (100)MaxY-Sent by BCA

Reject Description in Indonesian    rejectStatusENString (100)MaxY-Sent by BCA

Reject Description in EnglishtraceNoString (100)MaxY-Sent by BCA

Number for tracking to destination bank.

In BCA traceNo=ppu\_num ber with prefix “(“

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)FixedY-Response coderesponseMessageString(150)MaxY-Response description

curl –X POST https://copartners.com/openapi/v1.0/transfer-skn/notify

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95051'
-H 'X-PARTNER-ID': 'KBCABCINDO'
-d '
{
 "originalPartnerReferenceNo": "2020102900000000000001",
 "originalReferenceNo": "2020102977770000000009",
 "originalExternalId": "10052019",
 "latestTransactionStatus": "00",
 "amount": {
 "value": "10000.00",
 "currency": "IDR"
 },
 "beneficiaryAccountName": "Yories Yolanda",
 "beneficiaryAccountNo": "888801000003301",
 "beneficiaryBankCode": "002",
 "sourceAccountNo": "1234567890",
 "transactionDate": "2020-12-21T14:06:21+07:000",
 "additionalInfo": {
 "rejectStatusID": "Nomor rekening tidak terdaftar",
 "rejectStatusEN": "Account number is not registered",
 "traceNo": "(DAC8"
 }
}'
```

curl –X POST https://copartners.com/openapi/v1.0/transfer-skn/notify

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
 "responseCode": "2007500",
 "responseMessage": "Successful"
 }
```

This service is used to transfer RTGS.

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)FixedY Channel’s identifier using WSID KlikBCA Bisnis (95051) X-PARTNER-IDHeaderString(32)MaxY Corporate ID

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString (64)MaxY-Mandatory in BCA.

Transaction identifier on service consumer systemamountObjectY   valueString (16.2) MaxYDecimalMandatory in BCA. Net amount of the transaction.

If it’s IDR then value includes 2 decimal digits. In BCA, max length = 13.2   currencyString (3)FixedYISO-4217Mandatory in BCA.

CurrencybeneficiaryAccountNameString (100)MaxY-Mandatory in BCA.

Beneficiary Account NamebeneficiaryAccountNoString (34)MaxY-Mandatory in BCA.

Beneficiary Account beneficiaryAddressString (100)MaxN-Mandatory in BCA.

Beneficiary Address.beneficiaryBankCodeString(8)MaxY-Mandatory in BCA.

Beneficiary Bank Code. In BCA using SWIFT CodebeneficiaryCustomerResidenceString (1)FixedYNumericMandatory in BCA.

Beneficiary Customer Residence

1\. Indonesia

2\. Non IndonesiabeneficiaryCustomerTypeString (1)FixedYNumericMandatory in BCA.

Beneficiary Customer Type

1\. Individual

2\. Corporation

3\. GovernmentbeneficiaryEmailString(50)MaxNabc@domain.comOptional in BCA.

Beneficiary EmailremarkString (50) MaxN-Optional in BCA.

Remark/transact ion description.

For BCA, the remark is divided into two parts. The max length of each part is 18. sourceAccountNoString (19)MaxY-Mandatory in BCA. Source Account.

For BCA, length account number is 10 digit.transactionDateString (25)MaxYISO-8601Mandatory in BCA.

Transaction date

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString(64)MaxN-Send by BCA according to the request input.

Transaction identifier on sevice consumer systemresponseCodeString (7)FixedY-Response coderesponseMessageString(150)MaxY-Response descriptionreferenceNoString (64)MaxNNumericSend by BCA.

Transaction identifier on service provider system.

Must be filled upon successful transaction.amountObjectY   valueString (16.2)MaxYDecimalSend by BCA according to the request input.

Net amount of the transaction. If it’s IDR then value includes 2 decimal digits. In BCA, max length = 13.2   currencyString (3)FixedYISO-4217Send by BCA according to the request input.

CurrencybeneficiaryAccountNameString (100)MaxY-Send by BCA according to the request input.

Beneficiary Account Name beneficiaryAccountNoString (34) MaxYNumericSend by BCA according to the request input.

Beneficiary accounbeneficiaryBankCodeString (8)MaxNBCA : SWIFTSend by BCA according to the request input.

Beneficiary Bank CodesourceAccountNoString (19)MaxN-Send by BCA according to the request input.

Source account.

For BCA account 10 digits traceNoString (16)MaxN-Send by BCA.

Number for tracking to destination bank.

In BCA, traceNo = ppu\_number with prefix “(” and the length is 5transactionDateString (25)MaxYISO-8601Send by BCA according to the request input.

Transaction datetransactionStatusString (2)FixedYNumeric
00 - Success

03 – Pending

06 – Failed

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message4004002200Bad request4004002201Invalid Field Format {field}4004002202Invalid Mandatory Field {Field}4014012200Unauthorized. \[Reason\]4014012201Invalid token (B2B)4034032201Feature Not Allowed4034032202Exceeds Transaction Amount Limit4034032204Activity Count Limit Exceeded4034032206Feature Not Allowed At This Time. The transaction has exceeded time limit for today processing4034032206Feature Not Allowed At This Time. Transfer service to other bank is temporarily closed4034032214Insufficient Funds4034032218Inactive Account4044042211Invalid Account4044042213Invalid Amount 4044041802Invalid Routing 4094092200Conflict5005002201Internal Server Error5045042200Timeout

If you get responseCode 5002201 (Unknown Error/Unknown internal server failure, please retry the process again) from BCA, please make sure to check your account statement before retry the transaction.

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/transfer-rtgs

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95051'
-H 'X-PARTNER-ID': 'KBBABCINDO'
-d '
{
	"partnerReferenceNo": "2020102900000000000001",
	"amount": {
		"value": "10000.00",
		"currency": "IDR"
	},
	"beneficiaryAccountName": "Yories Yolanda",
	"beneficiaryAccountNo": "888801000157508",
	"beneficiaryAddress": "Palembang",
	"beneficiaryBankCode": "002",
	"beneficiaryCustomerResidence": "1",
	"beneficiaryCustomerType": "1",
	"beneficiaryEmail": "yories.yolanda@work.bri.co.id",
	"remark": "remark test",
	"sourceAccountNo": "1234567899",
	"transactionDate": "2020-12-21T14:36:11+07:00"
}'
```

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/transfer-rtgs

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
	"responseCode": "2002200",
	"responseMessage": "Request has been processed successfully",
	"referenceNo": "2020102977770000000009",
	"partnerReferenceNo": "2020102900000000000001",
	"amount": {
		"value": "10000.00",
		"currency": "IDR"
	},
	"beneficiaryAccountName": "Yories Yolanda",
	"beneficiaryAccountNo": "888801000157508",
	"beneficiaryBankCode": "002",
	"sourceAccountNo": "1234567899",
	"traceNo": "(DAC8",
	"transactionDate": "2020-12-21T14:36:11+07:00",
	"transactionStatus": "00"
}

```

##### 7\. SNAP Banking SKNBI Transfer

This service is used to transfer SKN.

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)FixedY Channel’s identifier using WSID KlikBCA Bisnis (95051) X-PARTNER-IDHeaderString(32)MaxY Corporate ID

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString (64)MaxY-Mandatory in BCA.

Transaction identifier on service consumer systemamountObjectY   valueString (16.2)  MaxYDecimalMandatory in BCA. Net amount of the transaction.

If it’s IDR then value includes 2 decimal digits. In BCA, max length = 13.2   currencyString (3)FixedYISO-4217Mandatory in BCA.

CurrencybeneficiaryAccountNameString (100)MaxY-Mandatory in BCA.

Beneficiary Account NamebeneficiaryAccountNoString (34)MaxY-Mandatory in BCA.

Beneficiary Account beneficiaryAddressString (100)MaxN-Mandatory in BCA.

Beneficiary Address.beneficiaryBankCodeString(8)MaxY-Mandatory in BCA.

Beneficiary Bank Code. In BCA using SWIFT CodebeneficiaryCustomerResidenceString (1)FixedYNumericMandatory in BCA.

Beneficiary Customer Residence

1\. Indonesia

2\. Non IndonesiabeneficiaryCustomerTypeString (1)FixedYNumericMandatory in BCA.

Beneficiary Customer Type

1\. Individual

2\. Corporation

3\. GovernmentbeneficiaryEmailString(50)MaxNabc@domain.comOptional in BCA.

Beneficiary EmailremarkString (50) MaxN-Optional in BCA.

Remark/transact ion description.

For BCA, the remark is divided into two parts. The max length of each part is 18. sourceAccountNoString (19)MaxY-Mandatory in BCA. Source Account.

For BCA, length account number is 10 digit.transactionDateString (25)MaxYISO-8601Mandatory in BCA.

Transaction date

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString(64)MaxN-Send by BCA according to the request input.

Transaction identifier on service consumer systemresponseCodeString (7)FixedY-Response coderesponseMessageString(150)MaxY-Response descriptionreferenceNoString (64)MaxNNumericSend by BCA.

Transaction identifier on service provider system.

Must be filled upon successful transaction.amountObjectY   valueString (16.2)MaxYDecimalSend by BCA according to the request input.

Net amount of the transaction. If it’s IDR then value includes 2 decimal digits. In BCA, max length = 13.2   currencyString (3)FixedYISO-4217Send by BCA according to the request input.

CurrencybeneficiaryAccountNameString (100)MaxY-Send by BCA according to the request input.

Beneficiary Account Name beneficiaryAccountNoString (34) MaxYNumericSend by BCA according to the request input.

Beneficiary accounbeneficiaryBankCodeString (8)MaxNBCA : SWIFTSend by BCA according to the request input.

Beneficiary Bank CodesourceAccountNoString (19)MaxN-Send by BCA according to the request input.

Source account.

For BCA account 10 digits traceNoString (16)MaxN-Send by BCA.

Number for tracking to destination bank.

In BCA, traceNo = ppu\_number with prefix “(” and the length is 5transactionDateString (25)MaxYISO-8601Send by BCA according to the request input.

Transaction datetransactionStatusString (2)FixedYNumeric
00 - Success

03 – Pending

06 – Failed

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message4004002300Bad request4004002301Invalid Field Format {field}4004002302Invalid Mandatory Field {Field}4014012300Unauthorized. \[Reason\]4014012301Invalid token (B2B)4034032301Feature Not Allowed4034032302Exceeds Transaction Amount Limit4034032304Activity Count Limit Exceeded4034032306Feature Not Allowed At This Time. The transaction has exceeded time limit for today processing4034032306Feature Not Allowed At This Time. Transfer service to other bank is temporarily closed4034032314Insufficient Funds4034032318Inactive Account4044042311Invalid Account4044042313Invalid Amount 4044041802Invalid Routing 4094092300Conflict5005002301Internal Server Error5045042300Timeout

If you get responseCode 5002301 (Unknown Error/Unknown internal server failure, please retry the process again) from BCA, please make sure to check your account statement before retry the transaction.

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/transfer-skn

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95051'
-H 'X-PARTNER-ID': 'KBBABCINDO'
-d '
{
	"partnerReferenceNo": "2020102900000000000001",
	"amount": {
		"value": "10000.00",
		"currency": "IDR"
	},
	"beneficiaryAccountName": "Yories Yolanda",
	"beneficiaryAccountNo": "888801000157508",
	"beneficiaryAddress": "Palembang",
	"beneficiaryBankCode": "002",
	"beneficiaryCustomerResidence": "1",
	"beneficiaryCustomerType": "1",
	"beneficiaryEmail": "yories.yolanda@work.bri.co.id",
	"remark": "remark test",
	"sourceAccountNo": "1234567899",
	"transactionDate": "2020-12-21T14:36:11+07:00"
}'
```

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/transfer-skn

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
	"responseCode": "2002300",
	"responseMessage": "Request has been processed successfully",
	"referenceNo": "2020102977770000000009",
	"partnerReferenceNo": "2020102900000000000001",
	"amount": {
		"value": "10000.00",
		"currency": "IDR"
	},
	"beneficiaryAccountName": "Yories Yolanda",
	"beneficiaryAccountNo": "888801000157508",
	"beneficiaryBankCode": "002",
	"sourceAccountNo": "1234567899",
	"traceNo": "(DAC8",
	"transactionDate": "2020-12-21T14:36:11+07:00",
	"transactionStatus": "00",
}

```

##### 8\. SNAP Banking Transaction Status Inquiry

This service is used to transfer BCA.

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)FixedY Channel’s identifier using WSID KlikBCA Bisnis (95051)X-PARTNER-IDHeaderString(32)MaxY Corporate ID X-EXTERNAL-IDHeaderString(32)MaxYNumeric String.

Reference number that should be unique in the same day.

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionoriginalPartnerReferenceNoString (64)MaxN-
Original transaction identifier on service from partneroriginalExternalIdString (32)MaxYNumeric
Original External-ID on header messageserviceCodeString (2)FixedYNumeric
Transaction type indicator (service code of the original transaction request)

Service code of original transaction that can be requested

17 : Service Intrabank Transfer

18 : Service Interbank Transfer

22 : Service Interbank Transfer RTGS

23 : Service Interbank Transfer SKN

33 : Service Payment to VA from IntrabanktransactionDateString (25)MaxYISODateTime

ISO-8601
Transfer Transaction date

Note :

Transaction Status Inquiry can only make inquiries on the status of the transaction below via Open API.

- Intrabank Transfer

- Interbank Transfer
- RTGS Transfer
- SKNBI Transfer
- Payment to VA from Intrabank

Transaction that can be inquired are only valid for the past 31 days.

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)FixedY-Response coderesponseMessageString(150)MaxY-Response descriptionoriginalReferenceNoString (64)MaxNNumericOriginal transaction identifier on service provider system.

Only mandatory if original transaction is successful.originalPartnerReferenceNoString (64)MaxN-This response will be sent according to the request input.


Original transaction identifier on service customer systemoriginalExternalIdString (32)MaxNNumericThis response will be sent according to the request input.


Original External-ID on header messageserviceCodeString (2)FixedYNumericThis response will be sent according to the request input.


Transaction type indicator (service code of the original transaction request)

Service code of original transaction that can be checked

17 : Service Intrabank Transfer

18 : Service Interbank Transfer

22 : Service Interbank Transfer RTGS

23 : Service Interbank Transfer SKN

33 : Service Payment to VAtransactionDateString (25)MaxNISODateTime

ISO-8601This response will be sent according to the request input.


Transaction dateamountObject-Y-
Net amount of the transactionvalueString (13.2)MaxYNumeric
If it’s IDR then value includes 2 decimal digits. BCA will send response amount with format value Numeric (13.2)   currencyString (3)FixedYISO-4217
Currency of the amount   beneficiaryAccountNoString (34)MaxYNumeric
Beneficiary account numberbeneficiaryBankCodeString (8)MaxNSWIFT Code/

3 digit bank codeBeneficiary Bank Code.

In BCA using SWIFT Code or 3 digit bank code (if destination doesn’t have SWIFT Code).

Beneficiary Bank Code can be checked on KBB menu :

Miscellaneous > Information > Bank Code ListreferenceNumberString (30)MaxNNumeric
Transaction identifier on service provider system.

Unique each day from BCA Must be filled upon successful transaction.sourceAccountNoString (10)MaxYNumericSource Account Number.latestTransactionStatusString (2)FixedYNumeric
Transfer status code that sent is the highlighted status code.

00 - Success

03 - Pending

06 – Failed

transactionStatusDescString(50)MaxNDescription status transaction

00 : Transaction Success

03 : Transaction In Progress

06 : Response message from failed transaction

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message**Condition**4004003600Bad requestGeneral request failed error.4004003601Invalid Field Format {field}The data you entered does not match the format requirements. 4004003602Invalid Mandatory Field {Field}Mandatory field should be fulfilled. 4014013600Unauthorized. \[Reason\]Invalid Signature/Unknown Client/Connection Not Allowed.4014013601Invalid token (B2B)Access Token Not Exist/Access Token Expired/Token is Invalid.4034033601Feature Not AllowedThe feature is not allowed to use. Please make sure you registered this feature in your partnership data. 4044043601Transaction not foundTransaction not found.4094093600ConflictX-EXTERNAL-ID duplicate.4294293600Too Many RequestsMaximum transaction limit exceeded. 5005003600General ErrorWrong request field input. 5005003601Internal Server ErrorUnknown internal server failure, please retry the process again.5045043600TimeoutYour transaction timeout.

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/transfer/status

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95051'
-H 'X-PARTNER-ID': 'KBBABCINDO'
-H 'X-EXTERNAL-ID': '28910000006578499987546738976812'
-d '
{
	"originalPartnerReferenceNo": "2020102900000000000001",
	"originalExternalId": "30443786930722726463280097920912",
	"serviceCode": "17",
	"transactionDate": "2023-02-02T00:00:00+07:00"
}'

```

curl –X POST https://sandbox.bca.co.id/openapi/v1.0/transfer/status

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
Inquiry Successful Status Transaction

-H 'Content-type': 'application/json'
-H 'X-TIMESTAMP': '2020-12-17T13:50:04+07:00'
{
	"responseCode": "2003600",
	"responseMessage": "Successful",
	"originalReferenceNo": "2010297777000009",
	"originalPartnerReferenceNo": "2020102900000000000001",
	"originalExternalId": "30443786930722726463280097920912",
	"serviceCode": "17",
	"transactionDate": "2023-02-02T00:00:00+07:00",
	"amount": {
	"value": "10000.00",
	"currency": "IDR"
	},
	"beneficiaryAccountNo": "888801000157508",
	"beneficiaryBankCode": "BRINIDJA",
	"referenceNumber": "10052019",
	"sourceAccountNo": "888801000157508",
	"latestTransactionStatus": "00",
	"transactionStatusDesc": "Transaction Success"
}

Inquiry Pending/Failed Status Transaction
-H 'Content-type': 'application/json'
-H 'X-TIMESTAMP': '2020-12-17T13:50:04+07:00'
{
	"responseCode": "2003600",
	"responseMessage": "Successful",
	"originalReferenceNo": "",
	"originalPartnerReferenceNo": "2020102900000000000001",
	"originalExternalId": "30443786930722726463280097920912",
	"serviceCode": "17",
	"transactionDate": "2023-02-02T00:00:00+07:00",
	"amount": {
	"value": "10000.00",
	"currency": "IDR"
	},
	"beneficiaryAccountNo": "888801000157508",
	"beneficiaryBankCode": "BRINIDJA",
	"referenceNumber": "10052019",
	"sourceAccountNo": "888801000157508",
	"latestTransactionStatus": "03",
	"transactionStatusDesc": "Transaction in Progress"
}
```

#### Collection

##### 1\. Account Debiting Consent Notification

Service for BCA to notify Account Authorization Registration Status.

This feature is not yet available to be accessed via Sandbox

Your Request must contain following information:

##### Request

FieldData Type**Mandatory****Length Type****Format**Descriptionrequest\_idString(33)Y  Fix- Request IDcustomer\_id\_merchantString(15)YMax Length Number CustomerID Merchantcustomer\_nameString(30)YMax Length- Customer Namedb\_account\_noString(10) YFixNumberMandatory if the status is 01 and 03statusString(2) YFixNumberSKPR Registration Status

“01” = Success Register

“02” = Failed Register

“03” = Success Delete

“04” = Failed Delete ReasonObject -  --  -english String(50) N Max Length- indonesia String(50) N Max Length-
Response

FieldDataType**Mandatory**Descriptionrequest\_idString(33)  Y (if SKPR Digital)Request ID if SKPR Digital response\_wsString(1)  YResponse

0 = Success

1 = Failed

POST https://copartners.com/account-authorization/notification

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
SKPR Digital Success

POST https://copartners.com/account-authorization/notification
-d '{
"request_id" : "A2AAB46BE91573E6E05400144FFED17AB",
 "customer_id_merchant" : "081818181818181",
 "customer_name" : "nama pelanggan",
 "db_account_no" : "0611112221",
 "status" : "01"
 } '

Request for registration SKPR Digital (failed)
POST https://copartners.com/account-authorization/notification
-d '{
 "request_id" : "A2AAB46BE91573E6E05400144FFED17AB",
 "customer_id_merchant" : "081818181818181",
 "customer_name" : "nama pelanggan",
 "status" : "02",
 "reason" : {
 "english" : "Invalid Account",
 "indonesia" : "Rekening tidak valid"
 }
 } '

Request for delete SKPR Digital (success)
POST https://copartners.com/account-authorization/notification
-d '{
 "request_id" : "A2AAB46BE91573E6E05400144FFED17AB",
 "customer_id_merchant" : "081818181818181",
 "customer_name" : "nama pelanggan",
 "db_account_no" : "0611112221",
 "status" : "03"
 } '

Request for delete SKPR Digital (failed)
POST https://copartners.com/account-authorization/notification -d '{
"customer_id_merchant" : "081818181818181",
"db_account_no" : "0611112221",
"status" : "04",
"reason" : {
"english" : "Invalid Account",
"indonesia" : "Rekening tidak valid"
}
} '

Request for delete SKPR Conventional (success)
POST https://copartners.com/account-authorization/notification
-d '{
 "customer_id_merchant" : "081818181818181",
 "db_account_no" : "0611112221",
 "status" : "03"
 } '

Request for delete SKPR Conventional (failed)
POST https://copartners.com/account-authorization/notification
-d '{
 "customer_id_merchant" : "081818181818181",
 "db_account_no" : "0611112221",
 "status" : "04",
 "reason" : {
 "english" : "Invalid Account",
 "indonesia" : "Rekening tidak valid"
 }
 } '

```

POST https://copartners.com/account-authorization/notification

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
Response for SKPR Digital
{
 "request_id" : "A2AAB46BE91573E6E05400144FFED17AB”,
 "response_ws" : "0"
}

Response for SKPR Conventional
{
 "response_ws" : "0"
}
```

##### 2\. Add Fund Collection

This service is used to collect some amount of money from your destination bank account that have been blocked before.

This feature is not yet available to be accessed via Sandbox

Your Request must contain following information:

##### Additional headers

FieldDataType**Mandatory**DescriptionChannelIDString(5)YChannel’s identifier. Fill it with “95051”.CredentialIDString(10)YChannel’s credential (KlikBCA Bisnis Corporate ID)

##### Request

FieldDataType**Mandatory**DescriptionTransactionIDString(18)YUnique data from Corporate. Format: Numeric.ReferenceNumberString(14)Y Sender's Transaction Reference ID. Format: Numeric.RequestTypeString(2)YB = With Blocking, NB = Without Blocking. Format: B or NB.DebitedAccountString(10)OAccount to be debited, please make sure you have authority to debit the account. Mandatory ONLY IF RequestType = NB (Without Blocking). Otherwise must left empty. Format: Numeric.AmountString(16)YMaximum length is 16, include decimal amount. Format: Numeric, 13.2CurrencyString(3)YCurrently is IDR only.CreditedAccountString(10)YAccount to be credited. Must be registered as own account at KlikBCA Bisnis. EffectiveDateString(10)YTransaction Effective Date. Max : 366 days. Format: yyyy-MM-dd.TransactionDateString(10)YThe date of the Transaction. Format: yyyy-MM-dd.Remark1String(18)YTransfer remark for receiver.Remark2String(18)YTransfer remark for receiver.EmailString(254)YEmail address of beneficiary (blocked account’s owner).

##### Response

FieldDataTypeDescriptionTransactionIDString(18)Unique data from Corporate. Format: Numeric.ReferenceNumberString(14)Sender's Transaction Reference ID. Format: Numeric.RequestTypeString(2)B = With Blocking, NB = Without Blocking. Format: B or NB.DebitedAccountString(10)Account to be debited, must registered at KlikBCA Bisnis.Mandatory ONLY IF RequestType = NB (Without Blocking). Otherwise must left empty. Format: Numeric.AmountString(16)Maximum length is 16, include decimal amount. Format: Numeric, 13.2CurrencyString(3)Currently is IDR only.CreditedAccountString(10)Account to be creditedEffectiveDateString(10)Transaction Effective Date. Max : 366 days. Format: yyyy-MM-dd.TransactionDateString(10)The date of the Transaction Performed get by system. Format: yyyy-MM-dd.StatusString(10)Transaction status

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message (Indonesian)Error Message (English)Additional Information400ESB-07-575 Sistem BCA sedang dalam pemeliharaanBCA system on maintenanceSystem under maintenance. 400ESB-14-001HMAC tidak cocokHMAC mismatchPlease check generate signature process.400ESB-14-002Permintaan tidak validInvalid request Invalid request.400ESB-14-003Timestamp tidak validInvalid timestampInvalid timestamp format. 500ESB-14-005Sistem sedang dalam maintenanceSystem under maintenanceSystem under maintenance.401ESB-14-008client\_id/client\_secret/grant \_type tidak validInvalid client\_id/client\_secret/grant\_ty peInvalid client\_id/client\_secret/grant\_ty pe.401ESB-14-009Tidak berhakUnauthorizedInvalid Signature/Unknown Client/Connection Not Allowed.429ESB-14-010Jumlah permintaan melebihi limitLimit request exceededLimit request exceeded.404ESB-14-011Service tidak adaService doesn't existInvalid service, please check service’s URI.401ESB-14-012Tidak berhak mengakses service iniNot allowed to access this serviceThe feature is not allowed to use. Please make sure you registered this feature in your partnership data.401ESB-14-014ChannelID/CredentialID tidak validInvalid ChannelID/CredentialIDChannel ID : 95051 Credential ID : KlikBCA Bisnis Corporate ID.400ESB-14-015Content Type tidak validInvalid Content TypeInvalid Content Type400ESB-14-016Format JSON tidak validInvalid JSON formatInvalid JSON format.504ESB-14-023Koneksi Timeout. Silakan cek transaksi kembali untuk memastikan keberhasilan transaksi Anda.Connection timeout. Please reinquiry your transaction to ensure transaction status. Please check your account statement before retry the transaction.400ESB-82-003TransactionID tidak unikTransactionID not uniquePlease provide unique TransactionID.400ESB-82-006Nominal Transaksi melebihi Limit MaksimumTransaction amount is greater than Maximum Limit AllowedYour transaction amount exceeds your company limit. Please check your Company Limit at KlikBCA Bisnis.400ESB-82-007Tanggal efektif lebih kecil dari tanggal hari iniEffective Date is smaller than todayEffective Date is smaller than today.400ESB-82-015Nominal transaksi kurang dari Limit MinimumTransaction amount is smaller than Minimum Limit AllowedYour transaction amount below your company limit. Please check your Company Limit at KlikBCA Bisnis.400ESB-82-019Saldo tidak cukupInsufficient fundAccount balance insufficient for the transaction.400ESB-82-027Amount harus lebih dari nolAmount must be greater than zeroAmount must be greater than zero400ESB-82-030Untuk Request Type: B, field DebitedAccount tidak diperlukanFor Request Type: B, field DebitedAccount is not necessaryDebitedAccount field only for Request type : NB. 400ESB-82-031 Frekuensi transaksi telah melebihi batas maksimumTransaction frequency has exceeded the maximum numberToo many request, Exceeds Transaction Frequency Limit. Please check your Company Limit at KlikBCA Bisnis.400ESB-82-177Tidak memiliki kuasa debet atas rekening debetNot authorized to debit the debited accountAuthorization to debit the debited account is not found in the system. Please make sure you have authority to debit the account.400ESB-82-178Rekening tidak aktifInactive AccountThe credited or debited account cannot be used for transaction. The account may be closed or blocked.400ESB-99-009Field (0) harus diisiMissing mandatory field (0)Mandatory field must be filled.400ESB-99-0222 Format alamat email tidak sesuaiEmail address format is not validCheck your email address format. Valid email address example : example@email.com.400ESB-99-040Tipe isian field (0) tidak validInput type for field (0) not validPlease refer to request format above. 400ESB-99-112Input string JSON tidak validInvalid JSON string inputPlease check JSON string input400ESB-99-128Panjang field (0) melebihi ketentuanField (0) exceed limitPlease refer to request format above. 400ESB-99-155Mata uang harus IDRCurrency must in IDRThe currency must in IDR.400ESB-99-173Format tanggal tidak sesuaiInvalid date formatValid date format : yyyy-MM-dd400ESB-99-211Panjang field (0) tidak valid Invalid (0) field lengthPlease check field length according to request-response format above.400ESB-99-376Format (0) tidak sesuaiInvalid field (0) formatAmount format is not valid. Please refer to request format above.400ESB-99-407Nomor rekening tidak terdaftar sebagai rekening operasional Account number not registered as operational accountPlease make sure the credited account already registered as operational account in KlikBCA Bisnis.400ESB-99-524ReferenceID sudah adaReferenceID already existsReferenceID already exists.400ESB-99-533Field (0) harus numerik Field (0) should be numericOnly use numeric format.400ESB-99-573TransactionDate harus diisi dengan tanggal hari ini Transaction Date must be filled with today's dateTransaction Date must be filled with today's date.400ESB-99-670Jenis transaksi tidak diijinkanTransaction Type not allowedPlease make sure your KlikBCA Bisnis package has AutoCollection Satuan features.400ESB-99-676Status blokir tidak aktifHold status is inactiveHold status is inactive.400ESB-99-677Rekening InvalidInvalid AccountAccount not found.400ESB-99-999Sistem sedang tidak tersediaSystem unavailableSystem unavailable at the moment.400ESB-99-999Koneksi Timeout. Silakan cek transaksi kembali untuk memastikan keberhasilan transaksi Anda.Connection timeout. Please reinquiry your transaction to ensure transaction status.Please check your account statement before retry the transaction.

POST /fund-collection/v2

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST /fund-collection/v2 HTTP/1.1
Host: sandbox.bca.co.id
-H "ChannelID:95051"
-H "CredentialID:ABCDEFGHIJ"
-d "{
"TransactionID" : "0000000001",
"ReferenceNumber" : "1111111110",
"RequestType" : "NB",
"DebitedAccount" : "1234567890",
"Amount" : "100200300.00",
"Currency" : "IDR",
"CreditedAccount" : "1234567891",
"EffectiveDate" : "2018-02-15",
"TransactionDate" : "2018-02-15",
"Remark1" : "Testing block 1",
"Remark2" : "Testing block 2",
"Email" : "beneficiary@mail.com"
}"
```

POST /fund-collection/v2

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "TransactionID" : "0000000001",
    "ReferenceID"" : "1111111110",
    "DebitedAccount" : "1234567890",
    "Amount" : "100200300.00",
    "Currency" : "IDR",
    "CreditedAccount" : "1234567891",
    "EffectiveDate" : "2018-02-15",
    "TransactionDate" : "2018-02-15"
	"Status"" : "TRAN SUCCESSFUL",
}
```

##### 3\. Inquiry Status Kuasa Pendebetan Rekening

This service is used to get account authorization status based on the customer id.

This feature is not yet available to be accessed via Sandbox

Your Request must contain following information:

##### Additional headers and URI Params

Field**Params Type**DataType**Mandatory**DescriptionChannelID HeaderString(5)YChannel ID (BCA channel’s identifier).Ex: “95051”.CredentialID HeaderString(10)YCompany Code Auto-Collectioncustomer\_id\_merchant PathString(15)Y Customer ID in number. Account authorization status that can be inquired is 7 days back from the inquiry date

##### Response

FieldDataTypeDescriptioncustomer\_id\_merchantString(15)Customer id from inputskpr\_pendingObjectAccount authorization pending list. Shows list of on process/failed account authorization registrationrequest\_idString(33)Request ID. ID when request registrationstatusString(2)Registration status:

02 = Failed

03 = On processfailed\_dateString(10)Account authorization failed date (only when status = 02)

Format: yyyy-MM-ddskpr\_activeObject Account authorization active list. Shows list of active account authorization registrationskpr\_idString(33)SKPR ID. ID when request registrationdb\_account\_noString(10)Account number registeredactive\_dateString(10)SKPR active date

Format: yyyy-MM-dd

Here is the list of error codes that can be returned:

HTTP CodeError Code**Error Message (Indonesian)**Error Message (English)4003-2-100ID pelanggan harus diisiCustomer ID cannot be empty4003-2-100ID pelanggan tidak validInvalid customer ID4003-2-100Nama pelanggan harus diisiCustomer name cannot be empty4003-2-100Nama pelanggan tidak validInvalid customer name4003-2-100Tipe debit harus diisiDebit type cannot be empty4003-2-100Tipe debit tidak validInvalid debit type4003-2-100Jumlah pembayaran harus diisiAmount cannot be empty4003-2-100Jumlah pembayaran tidak validInvalid amount4003-2-101Kode perusahaan tidak validInvalid company code4003-2-300Registrasi Kuasa Debet telah melebihi batas maksimal.

Silahkan hubungi Halo BCA untuk informasi lebih lanjutAccount Authorization registrastion already exceed the limit.

Please contact Halo BCA for the further information4003-2-307ID pelanggan sudah aktifCustomer ID already active4003-2-325Data pelanggan tidak ditemukanCustomer data not found4003-2-300Transaksi tidak dapat diproses. Silahkan ulangi beberapa saat lagiTransaction cannot be completed. Please try again later4003-2-998Transaksi gagal. Silahkan ulangi beberapa saat lagiTransaction failed. Please try again later.4003-2-999Transaksi tidak dapat diproses. Silahkan ulangi beberapa saat lagiTransaction cannot be completed. Please try again later

GET /account-authorization/inquiry/{customer\_id\_merchant}

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
GET /account-authorization/inquiry/081818181818181 HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
ChannelID: [channelID]
CredentialID: [credentialID]

```

GET /account-authorization/inquiry/{customer\_id\_merchant}

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "customer_id_merchant" : "081818181818181",
    "skpr_pending" : [
	{
		"request_id" : "A9DB1D7A9CE804C5E05400144FFED17A"",
		"status" : "03""
	},
	{
		"request_id" : "A9DB1D7A9CE704C5E05400144FFED17A"",
		"status" : "02"",
		"failed_date" : "2020-07-06"
	}
      ],
      "skpr_active" : [
	{
		"skpr_id" : "A7C6791C9ADC6CB9E05400144FFED17A"",
		"db_account_no" : "0250995514"",
		"active_date" : "2020-06-11"
	}
      ]
}
```

##### 4\. Pencabutan Kuasa Pendebetan Rekening

Delete your registration Account Authorization.

This feature is not yet available to be accessed via Sandbox

Your Request must contain following information:

##### Request Headers and URI Params

Field Params TypeDataType**Mandatory**Descriptionchannel-idHeaderString(5)YChannel ID KlikBCA Bisnis 95051credential-idHeader String(8)YCompany Code Auto-Collectioncustomer\_id\_merchantPath String(15) YCustomer ID Merchant db\_account\_noPath String(10)Y Customer Account No

##### Response

FieldDataType**Mandatory**Descriptioncustomer\_id\_merchantString(15)Y Customer ID Merchant db\_account\_noString(10) YCustomer Account No statusString(11) YDefault status “In Progress”

Here is the list of error codes that can be returned:

HTTP CodeError Code**Error Message (Indonesian)**Error Message (English)4003-2-100ID pelanggan harus diisiCustomer ID cannot be empty4003-2-100ID pelanggan tidak validInvalid customer ID4003-2-100Nama pelanggan harus diisiCustomer name cannot be empty4003-2-100Nama pelanggan tidak validInvalid customer name4003-2-100Tipe debit harus diisiDebit type cannot be empty4003-2-100Tipe debit tidak validInvalid debit type4003-2-100Jumlah pembayaran harus diisiAmount cannot be empty4003-2-100Jumlah pembayaran tidak validInvalid amount4003-2-101Kode perusahaan tidak validInvalid company code4003-2-300Registrasi Kuasa Debet telah melebihi batas maksimal.

Silahkan hubungi Halo BCA untuk informasi lebih lanjutAccount Authorization registrastion already exceed the limit.

Please contact Halo BCA for the further information4003-2-307ID pelanggan sudah aktifCustomer ID already active4003-2-325Data pelanggan tidak ditemukanCustomer data not found4003-2-300Transaksi tidak dapat diproses. Silahkan ulangi beberapa saat lagiTransaction cannot be completed. Please try again later4003-2-998Transaksi gagal. Silahkan ulangi beberapa saat lagiTransaction failed. Please try again later.4003-2-999Transaksi tidak dapat diproses. Silahkan ulangi beberapa saat lagiTransaction cannot be completed. Please try again later

DELETE /accountauthorization/customer/832789861532974/account-number/51421329865

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
curl –X DELETE https://sandbox.bca.co.id/accountauthorization/customer/832789861532974/account-number/51421329865
-H 'channel-id':'95051'
-H 'credential-id': '00000001'
```

DELETE /accountauthorization/customer/832789861532974/account-number/51421329865

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
 "customer_id_merchant": "832789861532974",
 "db_account_no": "51421329865",
 "status" : "In Progress"
}
```

##### 5\. Pendaftaran Kuasa Pendebetan rekening

Get your request ID which will be required to access Account Authorization webview.

This feature is not yet available to be accessed via Sandbox

Your Request must contain following information:

##### Additional headers

FieldDataType**Mandatory**DescriptionChannelIDString(5)YChannel ID (BCA channel’s identifier).Ex: “95051”.CredentialIDString(10)YCompany Code Auto-Collection

##### Request

FieldDataType**Mandatory**Descriptioncustomer\_id\_merchantString(15)YCustomer id merchant in number.identificationsObjectY Customer datatypeNumeric(1)YCustomer data type. Type contains numbers 1, 2, and 3.

1 = “Nama Pelanggan”

2 = “Tipe Debit”

3 = “Nominal Pembayaran”identificationString(30)YIf “type” : “1” then “identification ” : “customer name”.

Example : “identification” : “budi”identificationString(1)YIf “type” : “2” then “identification ” : “tipe\_debit” (F : Fix / V : Variable).

Example : “identification” : “V”identificationNumeric(13)YIf “type” : “3” then “identification ” :“nominal pembayaran”.

If tipe\_debit = V then amount “0”.

If tipe\_debit = F the amount “nominal transaksi.

Example : “identification” : “0”merchant\_logo\_urlStringN Merchant logo URL to display on webview

##### Response

FieldDataType**Mandatory**Descriptionrequest\_idString(33)Y Registration token to open a webviewrandom\_stringString(16) YRandom string used to generate String verification to open webviewexpired\_dateLong(13) YExpired date in epochcreated\_dateLong(13) YCreated date in epoch

Here is the list of error codes that can be returned:

HTTP CodeError Code**Error Message (Indonesian)**Error Message (English)4003-2-100ID pelanggan harus diisiCustomer ID cannot be empty4003-2-100ID pelanggan tidak validInvalid customer ID4003-2-100Nama pelanggan harus diisiCustomer name cannot be empty4003-2-100Nama pelanggan tidak validInvalid customer name4003-2-100Tipe debit harus diisiDebit type cannot be empty4003-2-100Tipe debit tidak validInvalid debit type4003-2-100Jumlah pembayaran harus diisiAmount cannot be empty4003-2-100Jumlah pembayaran tidak validInvalid amount4003-2-101Kode perusahaan tidak validInvalid company code4003-2-300Registrasi Kuasa Debet telah melebihi batas maksimal.

Silahkan hubungi Halo BCA untuk informasi lebih lanjutAccount Authorization registrastion already exceed the limit.

Please contact Halo BCA for the further information4003-2-307ID pelanggan sudah aktifCustomer ID already active4003-2-325Data pelanggan tidak ditemukanCustomer data not found4003-2-300Transaksi tidak dapat diproses. Silahkan ulangi beberapa saat lagiTransaction cannot be completed. Please try again later4003-2-998Transaksi gagal. Silahkan ulangi beberapa saat lagiTransaction failed. Please try again later.4003-2-999Transaksi tidak dapat diproses. Silahkan ulangi beberapa saat lagiTransaction cannot be completed. Please try again later

POST /account-authorization/registration

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST /account-authorization/registration HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
ChannelID: [channelID]
CredentialID: [credentialID]

{
    "customer_id_merchant" : "081818181818181",
     "identifications" : [
        {"type":"1","identification" : "nama pelanggan"},
        {"type":"2","identification" : "V"},
	{"type":"3","identification" : "0"},
        ],
    "merchant_logo_url" : "https://logo.com/",
}

```

POST /account-authorization/registration

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "request_id" : "A2AAB46BE91573E6E05400144FFED17A",
    "random_string"" : "L1APowN4R3uM1CiY",
    "expired_date" : "1586226654111",
    "created_date" : "1586224854111"
}
```

#### OneKlik

##### 1\. SNAP OneKlik Account Binding

This service is used to pre-processing OneKlik Registration by Generating OneKlik Registration Web Token.

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)Y Channel’s identifier (95221)X-PARTNER-IDHeaderString(36)Y Partner/Merchant Unique ID X-EXTERNAL-IDHeaderString (36)YReference Number from Partner/MerchantX-DEVICE-IDHeaderString (400)YCustomer Device Info Format: {device-id}\|\|{user-agent}\|\|{OS type}\|\|{OS version} device-id :

\- Android ID (Android)

\- UUID (iOS)

\- "WEB" (browser) user-agent : user agent device / browser (max 300 characters)

OS type :

Operating system type of device (Android,Iphone,Web)

OS version :

Operating system version of device with format (x.y.z or x.y)X-IP-ADDRESSHeaderString (15)YCustomer IP Address (IPv4 Format)

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionmerchantIdString (5)YPartner/Merchant Unique IDadditionalDataObjectY     userIdString (18)YCustomer Id from MerchantadditionalInfoObjectY     additionalInfoString (50)YAdditional information about customer (phone/email/etc)     merchantLogoUrl

String (300)NMerchant logo URL to show in OneKlik Registration Web

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)YResponse coderesponseMessageString (150)YResponse message descriptionreferenceNoString (32)YResponse reference number from OneKlikadditionalInfoObjectCFilled upon transaction with responseCode success     requestIdString (32)YRequest ID to open OneKlik Registration Web     randomStringString (16)YRandom String to generate token to open OneKlik Registration Web

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message2002000700Successful4004000700Bad request4004000701Invalid Field Format {Field} 4004000702Invalid Mandatory Field {Field}4014010701Invalid token (B2B)4014010700Unauthorized. \[Reason\] 4034030701Feature Not Allowed4034030704Activity Count Limit Exceeded4044040708Invalid Merchant5005000700General Error5005000700Internal Server Error5045040700Timeout

curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/registration-account-binding

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
DEVICE WEB

-H 'CHANNEL-ID':'95221'
-H 'X-PARTNER-ID': '61123'
-H 'X-EXTERNAL-ID': 'ef343939477c4f19b01c7600ca487bc8'
-H 'X-IP-ADDRESS': '123.123.123.123'
-H 'X-DEVICE-ID': 'WEB||Mozilla/5.0 (Windows NT 10.0; Win64; x64)
AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36||Web||95.0'
-d '
{
 "merchantId": "61123",
 "additionalData": {
 "userId": "customer01"
 },
 "additionalInfo": {
 "additionalInfo": "customer01@gmail.com",
 "merchantLogoUrl": " https://png.pngtree.com/pngvector/20190223/ourmid/pngtree-vector-picture-icon-png-image_695350.jpg"
 }
}'

DEVICE ANDROID/IOS

-H 'CHANNEL-ID':'95221'
-H 'X-PARTNER-ID': '61123'
-H 'X-EXTERNAL-ID': 'ef343939477c4f19b01c7600ca487bc8'
-H 'X-IP-ADDRESS': '123.123.123.123'
-H 'X-DEVICE-ID': '84723551-f7f0-452b-b9a5-b0f270752761||Mozilla/5.0 (iPhone; CPU
iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko)
Mobile/15E148||iphone||14.8.1'
-d '
{
 "merchantId": "61123",
 "additionalData": {
 "userId": "customer01"
 },
 "additionalInfo": {
 "additionalInfo": "customer01@gmail.com",
 "merchantLogoUrl": " https://png.pngtree.com/pngvector/20190223/ourmid/pngtree-vector-picture-icon-png-image_695350.jpg"
 }
}'
```

curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/registration-account-binding

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
 "responseCode": "2000700",
 "responseMessage": "Successful",
 "referenceNo": "37f89370b5034db0a283130c231dd27b",
 "additionalInfo": {
 "requestId": "2c90c800702d1234017033ab9a060003",
 "randomString": "g4BoEz43jfjVvAvN"
 }
}

```

##### 2\. SNAP OneKlik Account Binding Inquiry

This service is used to inquiry OneKlik Account Binding Status.

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)Y Channel’s identifier (95221)X-PARTNER-IDHeaderString(36)Y Partner/Merchant Unique ID X-EXTERNAL-IDHeaderString (36)YReference Number from Partner/Merchant

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionadditionalInfoObjectYMust only fill one field : requestIds or userId    requestIds>Array of String \[1..10\] @32CAlphanumeric Request ID to search registration status. Can be filled up to 10 strings    userId

String (18)CAlphanumericCustomer id from merchant   merchantId

String (5)

C

Numeric Merchant unique ID. Mandatory for IT Gateway integration only

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)YResponse coderesponseMessageString (150)YResponse message descriptionreferenceNoString (32)CResponse reference number from OneKlik

Filled if response successadditionalInfoObjectCFilled upon transaction with responseCode success  accountDataArray of ObjectCList of account.

Sent if inquiry data successful.

If input filled with requestId then Registration datas. If input filled with userId then Active account datas  requestId String (32)

C

Request ID used when registering account. Only show up when inquiry registration data

xcoIdString (32)CValid OneKlik Identifier when credential successfully registered to BCA.

Only show up when data already active   credentialTypeString (2)CCredential type.

DC = Debit card

CC = Credit card   credentialNumberString (16) CCustomer credential number (masked)   maxLimitString (9)COneKlik max daily limit registered   userIdString (18) CCustomer id from merchant   merchantIdString (5) CPartner/Merchant Unique ID   createdDateStringCDateTimeRequest Id created date.

Format: ISO 8601

Note: registration data can only be inquiried maximum of 7 days   statusString (3)YRegistration Status

REG = Registration in progress

EXP = Expired

ACT = Already activated (\*\* does not define the latest account status)

 -= not found/already past 7 days

Account Status

ACT = Active

BLK = Temporarily Blocked

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message2002000800Successful4004000800Bad request4004000801Invalid Field Format {Field} 4004000802Invalid Mandatory Field {Field}4014010801Invalid token (B2B)4014010800Unauthorized. \[Reason\] 4034030801Feature Not Allowed4044040808Invalid Merchant5005000800General Error5005000800Internal Server Error5045040800Timeout

With userId: curl –X POST https://sandbox.bca.co.idopenapi/oneklik/v1.0/registration-account-inquiry With requestIds: curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/registration-accountbinding

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
WITH userId - with merchantId

curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/registration-account-inquiry
-H 'CHANNEL-ID':'95221'
-H 'X-PARTNER-ID': '61123'
-H 'X-EXTERNAL-ID': 'ef343939477c4f19b01c7600ca487bc8'
-d '
{
 "additionalInfo": {
 "userId": "customer01",
 "merchantId": "61123",
 }
}'

WITH requestIds - without merchantId

curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/registration-account-binding
-H 'CHANNEL-ID':'95221'
-H 'X-PARTNER-ID': '61123'
-H 'X-EXTERNAL-ID': 'ef343939477c4f19b01c7600ca487bc8'
-d '
{
 "additionalInfo": {
 "requestIds": [
 "2c90c800702d1234017033ab9a060003",
 "2c90c800702d1234017033ab9a060004",
 "2c90c800702d1234017033ab9a060005",
 "2c90c800702d1234017033ab9a060006"
 ]
 }
}'

```

With userId: curl –X POST https://sandbox.bca.co.idopenapi/oneklik/v1.0/registration-account-inquiry With requestIds: curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/registration-accountbinding

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
WITH userId filled

{
 "responseCode": "2000800",
 "responseMessage": "Successful",
 "referenceNo": "37f89370b5034db0a283130c231dd27b",
 "additionalInfo": {
 "accountData": [
 {
 "xcoId": "a6fe3266cafd4cddbe3d77e1a9855634",
 "credentialType": "DC",
 "credentialNumber": "************1234",
 "maxLimit": "1000000",
 "userId": "customer01",
 "merchantId": "61123",
 "createdDate": "2022-01-03T01:07:51+07:00",
 "status": "ACT"
 },
 {
"xcoId": "76feef3439394cddbe3d77e1a98556ab",
 "credentialType": "DC",
 "credentialNumber": "************5678",
 "maxLimit": "1000000",
 "userId": "customer01",
 "merchantId": "61123",
 "createdDate": "2022-01-03T01:10:49+07:00",
 "status": "BLK"
 }
 ]
 }
}

WITH requestIds filled

{
 "responseCode": "2000800",
 "responseMessage": "Successful",
 "referenceNo": "37f89370b5034db0a283130c231dd27b",
 "additionalInfo": {
 "accountData": [
 {
 "requestId": "2c90c800702d1234017033ab9a060003",
 "xcoId": "a6fe3266cafd4cddbe3d77e1a9855634",
 "credentialType": "DC",
 "credentialNumber": "************1234",
 "maxLimit": "1000000",
 "userId": "customer01",
 "merchantId": "61123",
 "createdDate": "2022-01-03T01:02:03+07:00",
 "status": "ACT"
 },
 {
 "requestId": "2c90c800702d1234017033ab9a060004",
 "userId": "customer02",
"merchantId": "61123",
 "createdDate": "2022-01-03T01:05:25+07:00",
 "status": "REG"
 },
 {
 "requestId": "2c90c800702d1234017033ab9a060005",
 "userId": "customer03",
 "merchantId": "61123",
 "createdDate": "2022-01-03T01:10:09+07:00",
 "status": "EXP"
 },
 {
 "requestId": "2c90c800702d1234017033ab9a060006",
 "status": "-"
 }
 ]
 }
}

```

##### 3\. SNAP OneKlik Account Unbinding

This service is used to unbinding OneKlik Account.

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)Y Channel’s identifier (95221)X-PARTNER-IDHeaderString(36)Y Partner/Merchant Unique ID X-EXTERNAL-IDHeaderString (36)YReference Number from Partner/Merchant

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionmerchantIdString (5)YPartner/Merchant Unique IDadditionalInfoObjectY     xcoIdString (32)YValid OneKlik Identifier when credential successfully registered to BCA

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)YResponse coderesponseMessageString (150)YResponse message descriptionreferenceNoString (32)CResponse reference number from OneKlik

Filled if response successunlinkResultString (32) CResult of Unlinking process.

“success”/”failed”additionalInfoObjectCFilled upon transaction with responseCode success     xcoIdString (32)YValid OneKlik Identifier when credential successfully registered to BCA.

Only show up when data already active.

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message2002000900Successful4004000900Bad request4004000901Invalid Field Format {Field} 4004000902Invalid Mandatory Field {Field}4014010901Invalid token (B2B)4014010900Unauthorized. \[Reason\] 4034030901Feature Not Allowed4044040911Invalid Card/Account/Customer5005000900General Error5005000900Internal Server Error5045040900Timeout

curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/registration-account-unbinding

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95221'
-H 'X-PARTNER-ID': '61123'
-H 'X-EXTERNAL-ID': 'ef343939477c4f19b01c7600ca487bc8'
-d '
{
 "merchantId": "61123",
 "additionalInfo": {
 "xcoId": "a6fe3266cafd4cddbe3d77e1a9855634"
 }
}'
```

curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/registration-account-unbinding

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
 "responseCode": "2000900",
 "responseMessage": "Successful",
 "referenceNo": "37f89370b5034db0a283130c231dd27b",
 "unlinkResult": "success",
 "additionalInfo": {
 "xcoId": "a6fe3266cafd4cddbe3d77e1a9855634"
 }
}

```

##### 4\. SNAP OneKlik Direct Debit Payment

This service is used to pre-processing OneKlik Registration by Generating OneKlik Registration Web Token.

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)Y Channel’s identifier (95221)X-PARTNER-IDHeaderString(36)Y Partner/Merchant Unique ID X-EXTERNAL-IDHeaderString (36)YReference Number from Partner/MerchantX-IP-ADDRESSHeaderString (15)YCustomer IP Address (IPv4 Format)X-DEVICE-IDHeaderString (400)YCustomer Device Info Format: {device-id}\|\|{user-agent} device-id :

\- Android ID (Android)

\- UUID (iOS)

\- "WEB" (browser) user-agent : user agent device / browser (max 300 characters)

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionpartnerReferenceNoString (18) YTransaction ID from Partner/MerchantchargeTokenString (40)CToken to validate OTP. Received from Generate OTP service (service code 81).

Mandatory if otp field filledotpString (6)NOTP from SMS to validate transaction (used on different device / on merchant’s accord)merchantIdString (5)YPartner/Merchant Unique ID payOptionDetailsArray of Object (1)YPayment options used for this payment     payMethodString (64)YPayment method. Fill with “ONEKLIK”     payOptionString (64)YPayment option. Fill with “ONEKLIK\_BCA”     transAmountObjectYTotal billed amount from customer (transaction amount + fee amount)          valueString (16.2) YNumericTransAmount value          currencyString (3) YTransAmount currency      feeAmountObjectYFee/Admin Fee for BCA          valueString (16.2) YNumericFeeAmount value          currencyString (3)YFeeAmount currencyadditionalInfoObjectY     userIdString (18) YCustomer Id from Merchant     customerNameString (18) YCustomer name from Merchant. If sent more than 18 characters, it will be trimmed to 18 in BCA side     descriptionString (100) YTransaction description     paymentTypeString (7) YPayment Type: FULL = full payment PARTIAL = partial payment     transactionTimeStringYDateTimeTransaction date and time from merchant. Format yyyy-MMddTHH:mm:ssTZD format (ISO 8601 format)     xcoIdString (32) YValid OneKlik Identifier when credential successfully registered to BCA     billDetailsArray of Object (5)YBill distributed to specified details          subBillAmountString (16.2)YNumeric Sub Bill amount that will be credited to the specified subCompanyCode          subCompanyCodeString (5)YSub Company Code. Filled with “00000” if none     transferDetailObject C Mandatory for payment transfer only          beneficiaryNameString(100) C Beneficiary Name. Filled for payment transfer only          beneficiaryAccountString(32) CBeneficiary Account. Filled for payment           bankCodeString(10) C Bank Code. Filled for payment transfer only

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)YResponse coderesponseMessageString (150)YResponse message descriptionreferenceNoString (32)CResponse reference number from OneKlik appear only on success transactionpartnerReferenceNoString (18)YTransaction ID from merchant additionalInfoObjectCFilled upon transaction with responseCode success     statusString (50)YTransaction status.

“SUCCESS” = transaction success

“DEVICE IS DIFFERENT” = OTP verification is required     billAmountString (16.2)CNumericTotal bill amount from merchant.

Appear if status = "SUCCESS"      paidAmountString (16.2)CNumericTotal paid amount from customer.

Appear if status = "SUCCESS"     msisdnListArray of ObjectCList of phone number to send OTP.

Appear if status = “DEVICE IS DIFFERENT”          phoneNumberString (15)YMasked customer phone number          phoneIdString (32)YPhone id     paymentRequestIdString(30)C VA Transaction ID. Appear if status = “SUCCESS”

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message2002005400Successful2022025400Request In Progress4004005400Bad request4004005401Invalid Field Format {Field} 4004005402Invalid Mandatory Field {Field}4014015400Unauthorized. \[Reason\] 401 4015401 Invalid token (B2B)4034035401Feature Not Allowed4034035414Insufficient Funds4034035402Exceeds Transaction Amount Limit4034035412OTP Lifetime Expired403 4035415 Transaction Not Permitted. \[Reason\]  4034035418Inactive Card/Account/Customer4044045411Invalid Card/Account/Customer4044045408Invalid Merchant 4044045406Need to Request OTP4044045412Invalid Bill/Virtual Account total additionalInfo billDetails subBillAmount does not match transAmount.value4044045412Invalid Bill/Virtual Account adminFee.value does not match4044045412Invalid Bill/Virtual Account transAmount.value below transaction amount limit4044045412Invalid Bill/Virtual Account transAmount.value above transaction amount limit 4094095400 Cannot use same X-EXTERNAL-ID in same day 4094095401Duplicate partnerReferenceNo429 4295400 Too Many Requests 5005005400General Error5005005400Internal Server Error5045045400Timeout

curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/debit/payment-host-to-host

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
Without OTP – without Subcompany – without transferDetail – device type WEB

-H 'CHANNEL-ID':'95221'
-H 'X-PARTNER-ID': '61123'
-H 'X-EXTERNAL-ID': 'ef343939477c4f19b01c7600ca487bc8'
-H 'X-IP-ADDRESS': '123.123.123.123'
-H 'X-DEVICE-ID': 'WEB||Mozilla/5.0 (Windows NT 10.0; Win64; x64)
AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36'
-d '
{
 "partnerReferenceNo": "202010290000000001",
 "merchantId": "61123",
 "payOptionDetails": [
 {
 "payMethod": "ONEKLIK",
 "payOption": "ONEKLIK_BCA",
 "transAmount": {
 "value": "100000.00",
 "currency": "IDR"
 },
 "feeAmount": {
 "value": "1000.00",
 "currency": "IDR"
 }
 }
 ],
 "additionalInfo": {
 "userId": "customer01",
 "customerName": "John Smith",
 "description": "Payment for phone billing",
 "paymentType": "FULL",
 "transactionTime": "2020-10-29T00:00:01+07:00",
 "xcoId": " a6fe3266cafd4cddbe3d77e1a9855634",
 "billDetails": [
 {
 "subCompanyCode": "00000",
 "subBillAmount": "99000.00"
 }
 ]
 }
}'

With OTP – with Subcompany – with transferDetail – device type MOBILE (iOS/Android)

-H 'CHANNEL-ID':'95051'
-H 'X-PARTNER-ID': '61123'
-H 'X-EXTERNAL-ID': 'ef343939477c4f19b01c7600ca487bc8'
-H 'X-IP-ADDRESS': '123.123.123.123'
-H 'X-DEVICE-ID': '84723551-f7f0-452b-b9a5-b0f270752761||Mozilla/5.0 (iPhone; CPU
iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko)
Mobile/15E148'
-d '
{
 "partnerReferenceNo": "202010290000000001",
 "merchantId": "61123",
 "chargeToken": "1a6418f60dad46b9bdb2216d1a641186",
 "otp": "123456",
 "payOptionDetails": [
 {
 "payMethod": "ONEKLIK",
 "payOption": "ONEKLIK_BCA",
 "transAmount": {
 "value": "100000.00",
 "currency": "IDR"
 },
 "feeAmount": {
 "value": "1000.00",
 "currency": "IDR"
 }
 }
 ],
 "additionalInfo": {
 "userId": "customer01",
 "customerName": "John Smith",
 "description": "Payment for phone billing",
 "paymentType": "FULL",
 "transactionTime": "2020-10-29T00:00:01+07:00",
"xcoId": " a6fe3266cafd4cddbe3d77e1a9855634",
 "billDetails": [
 {
 "subCompanyCode": "00000",
 "subBillAmount": "50000.00"
 },
 {
 "subCompanyCode": "12345",
 "subBillAmount": "49000.00"
 }
 ]
 "transferDetail": {
 "beneficiaryName": "Steve Doe",
 "beneficiaryAccount": "8740012345"
 "bankCode": "014"
 }
 }
}'

```

curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/debit/payment-host-to-host

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
 "responseCode": "2005400",
 "responseMessage": "Successful",
 "referenceNo": "37f89370b5034db0a283130c231dd27b",
 "partnerReferenceNo": "202010290000000001",
 "additionalInfo": {
 "status": "SUCCESS",
 "billAmount": "100000.00",
 "paidAmount": "100000.00"
 "paymentRequestId": "202409250226156100100921955639"
 }

DEVICE IS DIFFERENT

{
 "responseCode": "4045406",
 "responseMessage": "Need to Request OTP",
 "partnerReferenceNo": "202010290000000001",
"additionalInfo": {
 "status": "DEVICE IS DIFFERENT",
 "msisdnList": [
 {
 "phoneId": "b5da903138c748a38264cea0424a191e",
 "phoneNumber": "0812XXXXX29"
 },
 {
 "phoneId": "2648cd3b308d4037bf4a3e75cc8988c9",
 "phoneNumber": "0817XXXXXX55"
 }
 ]
 }
}

```

##### 5\. SNAP OneKlik Direct Debit Payment Status

This service is used to inquiry OneKlik transaction status.

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)Y Channel’s identifier (95221)X-PARTNER-IDHeaderString(36)Y Partner/Merchant Unique ID X-EXTERNAL-IDHeaderString (36)YReference Number from Partner/Merchant

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionoriginalPartnerReferenceNoString (18) YAlphanumericTransaction ID from Partner/MerchantserviceCodeString (2)YNumericOriginal transaction service code. Fill with “54”merchantId

String (5)CNumericMerchant unique ID. Mandatory for IT Gateway integration only

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)YResponse coderesponseMessageString (150)YResponse message descriptionoriginalReferenceNoString (32)COriginal Response reference number from OneKlik. Filled if response successoriginalPartnerReferenceNoString (18)YTransaction ID from merchant serviceCodeString (2) YOriginal transaction service code. Filled with "54" (Direct Debit OneKlik)latestTransactionStatusString (2)CNumeric00 - Success

01 - Initiated

02 - Paying

06 - Failed

07 - Not foundtransAmountObjectCTotal billed amount from customer (transaction amount + fee amount)     valueString (16.2)YNumericTransAmount value     currencyString (3) YTransAmount currencyfeeAmountObjectCFee/Admin Fee for BCA     valueString (16.2)YNumericFeeAmount value     currencyString (3) YFeeAmount currencypaidTimeStringCDateTimeTimestamp when the payment done. Sent in yyyy-MMddTHH:mm:ssTZD format (ISO 8601 format)additionalInfoObjectCFilled upon transaction with responseCode success      paidAmountString (16.2)CNumericTotal paid amount from customer.     isTimeoutString (1) CY”/”N” Indicates if status 06 is Timeout or not     chargeTokenString (32) CCharge token after Generate OTP for status 01     paymentRequestId

String (30)C VA Transaction ID. Appear if latestTransactionStatus = 00

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message2002005500Successful4004005500Bad request4004005501Invalid Field Format {Field} 4004005502Invalid Mandatory Field {Field}4014015501Invalid token (B2B)4014015500Unauthorized. \[Reason\] 4034035501Feature Not Allowed4034035518Inactive Card/Account/Customer4044045511Invalid Card/Account/Customer4044045508Invalid Merchant4094095500 Cannot Use Same X-EXTERNAL-ID in same day

5005005500General Error5005005500Internal Server Error5045045500Timeout

curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/debit/status

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95221'
-H 'X-PARTNER-ID': '61123'
-H 'X-EXTERNAL-ID': 'ef343939477c4f19b01c7600ca487bc8'
-d '
{
 "originalPartnerReferenceNo": "202010290000000001",
 "serviceCode": "54",
 "merchantId": "61123"
}'

```

curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/debit/status

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
 "responseCode": "2005500",
 "responseMessage": "Successful",
 "originalReferenceNo": "37f89370b5034db0a283130c231dd27b",
 "originalPartnerReferenceNo": "202010290000000001",
 "serviceCode": "54",
 "latestTransactionStatus": "00",
 "transAmount": {
 "value": "100000.00",
 "currency": "IDR"
},
 "feeAmount": {
 "value": "1000.00",
 "currency": "IDR"
 },
 "paidTime": "2020-10-29T00:00:04+07:00",
 "additionalInfo": {
 "paidAmount": "100000.00"
 },
 "paymentRequestId": "202409250226156100100921955639"
}

Success – Timeout Transaction

{
 "responseCode": "2005500",
 "responseMessage": "Successful",
 "originalReferenceNo": "37f89370b5034db0a283130c231dd27b",
 "originalPartnerReferenceNo": "202010290000000001",
 "serviceCode": "54",
 "latestTransactionStatus": "06",
 "transAmount": {
 "value": "100000.00",
 "currency": "IDR"
 },
 "feeAmount": {
 "value": "1000.00",
 "currency": "IDR"
 },
 "paidTime": "2020-10-29T00:00:04+07:00",
 "additionalInfo": {
 "paidAmount": "100000.00",
 "isTimeout": "Y"
 }
}

Success – Initiated Transaction (Need OTP)
{
 "responseCode": "2005500",
 "responseMessage": "Successful",
 "originalReferenceNo": "37f89370b5034db0a283130c231dd27b",
 "originalPartnerReferenceNo": "202010290000000001",
 "serviceCode": "54",
 "latestTransactionStatus": "06",
 "transAmount": {
 "value": "100000.00",
 "currency": "IDR"
 },
 "feeAmount": {
 "value": "1000.00",
 "currency": "IDR"
 },
 "paidTime": "2020-10-29T00:00:04+07:00",
 "additionalInfo": {
 "chargeToken": "1a6418f60dad46b9bdb2216d1a641186"
 }
}

```

##### 6\. SNAP OneKlik Inquiry OTP

This service is used to inquiry list of phone numbers that can be used to send OTP.

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)Y Channel’s identifier (95221)X-PARTNER-IDHeaderString(36)Y Partner/Merchant Unique ID X-EXTERNAL-IDHeaderString (36)YReference Number from Partner/Merchant

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionmerchantIdString (5)YPartner/Merchant Unique IDadditionalInfoObjectY     xcoIdString (32)YValid OneKlik Identifier when credential successfully registered to BCA     userIdString (18YCustomer Id from Merchant

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)YResponse coderesponseMessageString (150)YResponse message descriptionadditionalInfoObjectCFilled upon transaction with responseCode success    msisdnListArray of ObjectYList of phone number to send OTP.          phoneNumberString (15)YMasked customer phone number           phoneIdString (32)YPhone id

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message2002008500Successful4004008500Bad request4004008501Invalid Field Format {Field} 4004008502Invalid Mandatory Field {Field}4014018501Invalid token (B2B)4014018500Unauthorized. \[Reason\] 4034038501Feature Not Allowed4034038518Inactive Card/Account/Customer4044048511Invalid Card/Account/Customer5005008500General Error5005008500Internal Server Error5045048500Timeout

curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/payment/inquiry-phone-number

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95221'
-H 'X-PARTNER-ID': '61123'
-H 'X-EXTERNAL-ID': 'ef343939477c4f19b01c7600ca487bc8'
-d '
{
 "merchantId": "61123",
 "additionalInfo": {
 "xcoId": "a6fe3266cafd4cddbe3d77e1a9855634",
 "userId": "customer01"
 }
}'
```

curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/payment/inquiry-phone-number

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
 "responseCode": "2008500",
 "responseMessage": "Successful",
 "additionalInfo": {
 "msisdnList": [
 {
 "phoneId": "b5da903138c748a38264cea0424a191e",
 "phoneNumber": "0812XXXXX29"
 },
 {
 "phoneId": "2648cd3b308d4037bf4a3e75cc8988c9",
 "phoneNumber": "0817XXXXXX55"
 }
 ]
 }
}
```

##### 7\. SNAP OneKlik Notify Registration/Delete

This service is used to inquiry list of phone numbers that can be used to send OTP.

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

Field**Params Type**Data Type**Format****Mandatory**DescriptionCHANNEL-IDHeaderString(5)NumericY Channel’s identifierX-PARTNER-IDHeaderString(32)AlphanumericY BCA Unique ID X-EXTERNAL-IDHeaderString (32)AlphanumericYReference Number from BCA

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionadditionalInfoObjectY     requestIdString (32)CAlphanumericRequest ID used when registering account. Only show up when notifying activation data     xcoIdString (32)YAlphanumericValid OneKlik Identifier when credential successfully registered to BCA.     credentialTypeString (2)CAlphanumericCredential type.

DC = Debit card

CC = Credit card     credentialNumberString (16)CAlphanumericCustomer credential number (masked)     maxLimitString (9)CNumericOneKlik max daily limit registered      userIdString (18)YAlphanumericCustomer id from merchant     createdDateStringYDateTimeCreated / Deleted date. Format: ISO-8601     statusString (3)YAlphanumericOneKlik Account Nofitication Status

ACT = Activated

DEL = Deleted     merchantId

String (5) YNumericMerchant Unique ID

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)YNumericResponse coderesponseMessageString (150)YAlphanumericResponse message descriptionreferenceNoString (32)CAlphanumericResponse reference from copartner. Filled if the response was successadditionalInfoObjectCFilled upon transaction with responseCode success     xcoIdString (32) YAlphanumericXCO ID     responseWsString (16)YAlphanumericResponse:

0 = Success

1 = Failed

POST https://copartners.com/oneklik/v2.0/notification

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
Request (Activation Notification)

-H 'CHANNEL-ID':'95221'
-H 'X-PARTNER-ID':'61123'
-H 'X-EXTERNAL-ID': 'ef343939477c4f19b01c7600ca487bc8'
-d ' {
 "additionalInfo":{
"requestd":" 2c909de773e578f80173e600b7950018",
"xcoId":" 2c909de773e578f80173e600b7950019",
"createdDate":"2022-01-03T01:02:03+07:00",
"credentialNumber":"************1234",
"credentialType":"DC",
"maxLimit":"100000",
"userId":"customer01",
"status":"ACT",
"merchantId":"61123"
}
}'

Request (Delete Account Notification)

-H 'CHANNEL-ID':'95221'
-H 'X-PARTNER-ID':'61123'
-H 'X-EXTERNAL-ID': 'ef343939477c4f19b01c7600ca487bc8'
-d ' {
 "additionalInfo":{
"xcoId":" 2c909de773e578f80173e600b7950019",
"createdDate":"2022-01-03T01:02:03+07:00",
"userId":"customer01",
"status":"DEL",
"merchantId":"61123"
 }
}'
```

POST https://copartners.com/oneklik/v2.0/notification

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
Response (Success)

{
 "responseCode": "2008600",
 "responseMessage": "Successful",
 "referenceNo": "37f89370b5034db0a283130c231dd27b",
 "additionalInfo": {
 "xcoId": "2c90c800702d1234017033ab9a060003",
 "responseWs": "0"
 }
}

Response (Failed)

{
 "responseCode": "4048611",
 "responseMessage":"Invalid Card/Account/Customer "
}
```

##### 8\. SNAP OneKlik OTP

This service is used to generate and send OTP to customer phone to authorize payment.

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

Field**Params Type**Data Type**Length Type****Mandatory**DescriptionCHANNEL-IDHeaderString(5)Y Channel’s identifier (95221)X-PARTNER-IDHeaderString(36)Y Partner/Merchant Unique ID X-EXTERNAL-IDHeaderString (36)YReference Number from Partner/Merchant

**Payload:**

FieldData Type**Length Type****Mandatory****Format**DescriptionjourneyIdString (32)YJourney Identifier. Must be equal with X-EXTERNAL-ID headerbankCardTokenString (32)YCard Token for Payment. Fill with XCOIDmerchantIdString (5)YPartner/Merchant Unique IDpartnerReferenceNoString (18)YTransaction ID from Partner/MerchanttrxDateTimeStringYDateTimeTransaction date and time from merchant. Format yyyy-MMddTHH:mm:ssTZD format (ISO 8601 format)additionalInfoObjectY     userIdString (18)YCustomer Id from Merchant     transAmountObjectYTotal billed amount from customer (transaction amount + fee amount)     valueString (16.2)YNumericTransAmount value     currencyString (3)YTransAmount currencyphoneIdString (32)YPhone id

Result of the request will contains following information:

**Response:**

FieldData Type**Length Type****Mandatory****Format**DescriptionresponseCodeString (7)YResponse coderesponseMessageString (150)YResponse message descriptionchargeTokenString (32)YString code for OTP verification. Sent when response code success / timeout additionalInfoObjectCFilled upon transaction with responseCode success     xcoIdString (32)YValid Identifier when credential successfully registered to BCA     phoneIdString (32)YPhone id     merchantIdString (5)YPartner/Merchant Unique ID

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message2002008100Successful4004008100Bad request4004008101Invalid Field Format {Field} 4004008102Invalid Mandatory Field {Field}4014018101Invalid token (B2B)4014018100Unauthorized. \[Reason\] 4034038101Feature Not Allowed4034038118Inactive Card/Account/Customer4044048111Invalid Card/Account/Customer4044048108Invalid Merchant4294298100Too Many Requests5005005100General Error5005008500Internal Server Error5045048500Timeout

curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/otp

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
-H 'CHANNEL-ID':'95221'
-H 'X-PARTNER-ID': '61123'
-H 'X-EXTERNAL-ID': 'ef343939477c4f19b01c7600ca487bc8'
-d '
{
 "journeyId": "ef343939477c4f19b01c7600ca487bc8",
 "bankCardToken": "a6fe3266cafd4cddbe3d77e1a9855634",
 "merchantId": "61123",
 "partnerReferenceNo": "202010290000000001",
 "trxDateTime": "2020-10-29T00:00:01+07:00",
 "additionalInfo": {
 "userId": "customer01",
 "transAmount": {
 "value": "100000.00",
 "currency": "IDR"
 },
 "phoneId": "b5da903138c748a38264cea0424a191e"
 }
}'

```

curl –X POST https://sandbox.bca.co.id/openapi/oneklik/v1.0/otp

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
SUCCESS

{
 "responseCode": "2008100",
 "responseMessage": "Successful",
 "chargeToken": "1a6418f60dad46b9bdb2216d1a641186",
 "additionalInfo": {
 "xcoId": "a6fe3266cafd4cddbe3d77e1a9855634",
 "phoneId": "b5da903138c748a38264cea0424a191e",
 "merchantId": "61123"
 }
}

TIMEOUT

{
 "responseCode": "5048100",
 "responseMessage": "Timeout",
 "chargeToken": "1a6418f60dad46b9bdb2216d1a641186",
 "additionalInfo": {
 "xcoId": "a6fe3266cafd4cddbe3d77e1a9855634",
 "phoneId": "b5da903138c748a38264cea0424a191e",
 "merchantId": "61123"
 }
}
```

#### Valuta Asing

##### 1\. Info Kurs

Get information about Foreign Currency Exchange Rate for applicable currencies in BCA. The exchange rate are divide into four types eRate (Electronic Rate), TT (Telegrafic Transfer), TC (Travellers Cheque) and BN (Bank Notes).

Query String Parameters:

FieldDataType**Mandatory**DescriptionCurrencyCodeStringN Known currency code. For multiple currencies use comma separated.

Following are the known foreign currency in BCA.

**AUD - Australia Dollar**

**BND - Bruneian Dollar**

**CAD - Canadian Dollar**

**CHF - Francs**

**CNY - China Yuan**

**DKK - Danish Krone**

**EUR - Euro**

**GBP - Great Britain Poundsterling**

**HKD - Hongkong Dollar**

**JPY - Japan Yen**

**KRW - Korea Won**

**NOK - Norwegian Krone**

**NZD - New Zealand Dollar**

**PHP - Phillipine Peso**

**SAR - Saudi Riyal**

**SEK - Swedish Krona**

**SGD - Singapore Dollar**

**THB - Thailand Baht**

**TWD - Taiwan Dollar**

**USD - US Dollar**

If the field was left empty, it will considered as “all know currencies”.RateTypeString YThe type for the desired rate.

**eRate (electronic Rate)**

**TT (Telegrafic Transfer)**

**TC (Travellers Cheque)**

**BN (Bank Notes)**

If the field was left empty, it will considered as “all rate type”.

Result of the request will contains following information:

**Response:**

FieldDataTypeDescriptionInvalidRateTypeStringIf the request contains valid and invalid rate types in single request, the invalid rate will be list down in this field.InvalidCurrencyCodeStringIf the request contains valid and invalid currency code in single request, the invalid currency code will be list down in this field.**Currencies**CurrencyCodeString (3)The currency code.

e.g. IDR, USD, JPY**RateDetail**RateType String (5)Must be among `erate, tt, tc, bn`.BuyString (13) In format `6.6`.

When the currency has no value for this field, it will be set to `"-"`.SellString (13) In format `6.6`.

When the currency has no value for this field, it will be set to `"-"`.LastUpdateString (19) In format `YYYY-MM-DDTHH24:Mi:SS`.SSSZ.

When the currency has no value for this type, it will be set to `"-"`.

##### Error

HTTP CodeError CodeError Message (Indonesian)Error Message (English)404 ESB-07-271 Tidak ada transaksi No transaction400 ESB-07-279 Tanggal awal lebih kecil dari tanggal buka rekening Start date less than account open date400 ESB-14-001 HMAC tidak cocok HMAC mismatch400 ESB-14-002 Permintaan tidak valid Invalid request400 ESB-14-003 Timestamp tidak valid Invalid timestamp400 ESB-14-004 parameter input tidak valid Invalid input parameter500 ESB-14-005 Sistem sedang dalam maintenance System under maintenance504 ESB-14-006 Timeout, silakan periksa mutasi rekening Timeout, please check your account statement504 ESB-14-007 Timeout Timeout401 ESB-14-008 client\_id/client\_secret/grant\_type tidak valid Invalid client\_id/client\_secret/grant\_type401 ESB-14-009 Tidak berhak Unauthorized429 ESB-14-010 Jumlah permintaan melebihi limit Limit request exceeded404 ESB-14-011 Service tidak ada Service doesn’t exist401 ESB-14-012 Tidak berhak mengakses service ini Not allowed to access this service400 ESB-99-009 Field Latitude harus diisi Missing Mandatory field Latitude400 ESB-99-009 Field Longitude harus diisi Missing Mandatory field Longitude400 ESB-99-039 Isian input field Count tidak valid Input for field Count not valid400 ESB-99-040 Tipe isian field Count tidak valid Input type for field Count not valid400 ESB-99-040 Tipe isian field Latitude tidak valid Input type for field Latitude not valid400 ESB-99-040 Tipe isian field Longitude tidak valid Input type for field Longitude not valid400 ESB-99-040 Tipe isian field Radius tidak valid Input type for field Radius not valid400 ESB-99-132 Data tidak ditemukan No Data Found400 ESB-99-335 Mata uang tidak valid Invalid CurrencyCode400 ESB-99-339 Tipe Kurs tidak valid Invalid RateType400 ESB-99-392 Nilai SearchBy tidak valid Invalid SearchBy value400 ESB-99-460 Panjang input Value tidak valid Input Value minimal length invalid500 ESB-99-999 Sistem sedang tidak tersedia System unavaliable

Info Kurs

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
GET /general/rate/forex?CurrencyCode=USD,JPY,XXX&RateType=erate,tt,yy HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
```

Info Kurs

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
  "Currencies": [{
	"CurrencyCode": "USD",
	"RateDetail": [{
		"RateType": "erate"
		"Buy": 17500.00,
		"Sell": 17000.00,
		"LastUpdate": "2015-09-04T19:45:44.234+07:00",
		},
		{
		"RateType": "tt"
		"Buy": 12000.00,
		"Sell": 12300.00,
		"LastUpdate": "2015-09-04T19:45:44.234+07:00",
		}]
	},
	{
	"CurrencyCode": "JPY",
	"RateDetail": [{
		"RateType": "erate"
		"Buy": 121.00,
		"Sell": 122.00,
		"LastUpdate": "2015-09-04T19:45:44.234+07:00",
		},
		{
		"RateType": "tt"
		"Buy": -,
		"Sell": -,
		"LastUpdate": "-",
		}]
	}],
  "InvalidRateType":"yy"
  "InvalidCurrency":"XXX"
}
```

##### 2\. Inquiry Kurs

Inquiry exchange rate to be used on Forex Transaction API

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

FieldData Type**Mandatory**Descriptionchannel-idString(5)Y Channel Identification Number (Ex: 95057 for JVALAS)credential-idString(10)Y CIN/CIS Example : 01234567890

**Path Parameters:**

FieldDataType**Mandatory**Descriptioncis String(10) Y Customer number currency-codeString(3)Y Foreign currency code.

Following are the known foreign currency in BCA.

**AUD - Australia Dollar**

**BND - Bruneian Dollar**

**CAD - Canadian Dollar**

**CHF - Francs**

**CNY - China Yuan**

**DKK - Danish Krone**

**EUR - Euro**

**GBP - Great Britain Poundsterling**

**HKD - Hongkong Dollar**

**JPY - Japan Yen**

**KRW - Korea Won**

**NOK - Norwegian Krone**

**NZD - New Zealand Dollar**

**PHP - Phillipine Peso**

**SAR - Saudi Riyal**

**SEK - Swedish Krona**

**SGD - Singapore Dollar**

**THB - Thailand Baht**

**TWD - Taiwan Dollar**

**USD - US Dollar**

If the field was left empty, it will considered as “all know currencies”.

##### Query Parameter

FieldData Type MandatoryDescriptionsettlement-typeStringY Param2 for decide Valas or Outward Remmitance (OR)

Result of the request will contains following information:

**Response:**

FieldDataTypeDescriptioncurrency\_codeString(3) The currency code. e.g. IDR, USD, JPYbuyString (13) In format number(6.6). When the currency has no value for this field, it will be set to “-“.sellString (13) In format number(6.6). When the currency has no value for this field, it will be set to “-“.session\_idString (50)

##### Error

HTTP CodeError CodeError Message (Indonesian)Error Message (English)404 ESB-07-271 Tidak ada transaksi No transaction400 ESB-07-279 Tanggal awal lebih kecil dari tanggal buka rekening Start date less than account open date400 ESB-14-001 HMAC tidak cocok HMAC mismatch400 ESB-14-002 Permintaan tidak valid Invalid request400 ESB-14-003 Timestamp tidak valid Invalid timestamp400 ESB-14-004 parameter input tidak valid Invalid input parameter500 ESB-14-005 Sistem sedang dalam maintenance System under maintenance504 ESB-14-006 Timeout, silakan periksa mutasi rekening Timeout, please check your account statement504 ESB-14-007 Timeout Timeout401 ESB-14-008 client\_id/client\_secret/grant\_type tidak valid Invalid client\_id/client\_secret/grant\_type401 ESB-14-009 Tidak berhak Unauthorized429 ESB-14-010 Jumlah permintaan melebihi limit Limit request exceeded404 ESB-14-011 Service tidak ada Service doesn’t exist401 ESB-14-012 Tidak berhak mengakses service ini Not allowed to access this service400 ESB-99-009 Field Latitude harus diisi Missing Mandatory field Latitude400 ESB-99-009 Field Longitude harus diisi Missing Mandatory field Longitude400 ESB-99-039 Isian input field Count tidak valid Input for field Count not valid400 ESB-99-040 Tipe isian field Count tidak valid Input type for field Count not valid400 ESB-99-040 Tipe isian field Latitude tidak valid Input type for field Latitude not valid400 ESB-99-040 Tipe isian field Longitude tidak valid Input type for field Longitude not valid400 ESB-99-040 Tipe isian field Radius tidak valid Input type for field Radius not valid400 ESB-99-132 Data tidak ditemukan No Data Found400 ESB-99-335 Mata uang tidak valid Invalid CurrencyCode400 ESB-99-339 Tipe Kurs tidak valid Invalid RateType400 ESB-99-392 Nilai SearchBy tidak valid Invalid SearchBy value400 ESB-99-460 Panjang input Value tidak valid Input Value minimal length invalid500 ESB-99-999 Sistem sedang tidak tersedia System unavaliable

GET /rate/management/customers/{cis}/currenc ies/{currency\_code}

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
GET
https://sandbox.bca.co.id/rate/management/customers/00000002120/currencies/USD?settle
ment-type=VALAS
-H "channel-id" : "95057"
-H "credential-id" : "00000002120"

```

GET /rate/management/customers/{cis}/currenc ies/{currency\_code}

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
 "currency_code": "USD",
 "buy": "14058",
 "sell": "14068",
"session_id": "rHaaVK4vNlb713GMmdXFXB94MSfhZMwjL+Uqv9EKqwIpA1CGKcOF36IWHaa0DKqaBK6Im1mvzCRg9HMHUZ550w=="
}
```

##### 3\. Transaksi Valas

You can send fund transfer forex instructions to BCA using this service. The source of fund
transfer must be from your corporate's own deposit account. The recipient may be any deposit
account within domestic bank include BCA.

This feature is not yet available to be accessed via Sandbox

**Additional Headers:**

FieldData Type**Mandatory**Descriptionchannel-idString(5)Y Channel Identification Number (Ex: 95057 for JVALAS)credential-idString(10)Y CIN/CIS Example : 01234567890

**Path Parameters:**

FieldDataType**Mandatory**Descriptioncis String(10) Y Customer number currency-codeString(3)Y Foreign currency code.

Following are the known foreign currency in BCA.

**AUD - Australia Dollar**

**BND - Bruneian Dollar**

**CAD - Canadian Dollar**

**CHF - Francs**

**CNY - China Yuan**

**DKK - Danish Krone**

**EUR - Euro**

**GBP - Great Britain Poundsterling**

**HKD - Hongkong Dollar**

**JPY - Japan Yen**

**KRW - Korea Won**

**NOK - Norwegian Krone**

**NZD - New Zealand Dollar**

**PHP - Phillipine Peso**

**SAR - Saudi Riyal**

**SEK - Swedish Krona**

**SGD - Singapore Dollar**

**THB - Thailand Baht**

**TWD - Taiwan Dollar**

**USD - US Dollar**

If the field was left empty, it will considered as “all know currencies”.

##### Query Parameter

FieldData Type MandatoryDescriptionsettlement-typeStringY Param2 for decide Valas or Outward Remmitance (OR)

Result of the request will contains following information:

**Response:**

FieldDataTypeDescriptioncurrency\_codeString(3) The currency code. e.g. IDR, USD, JPYbuyString (13) In format number(6.6). When the currency has no value for this field, it will be set to “-“.sellString (13) In format number(6.6). When the currency has no value for this field, it will be set to “-“.session\_idString (50)

##### Error

HTTP CodeError CodeError Message (Indonesian)Error Message (English)404 ESB-07-271 Tidak ada transaksi No transaction400 ESB-07-279 Tanggal awal lebih kecil dari tanggal buka rekening Start date less than account open date400 ESB-14-001 HMAC tidak cocok HMAC mismatch400 ESB-14-002 Permintaan tidak valid Invalid request400 ESB-14-003 Timestamp tidak valid Invalid timestamp400 ESB-14-004 parameter input tidak valid Invalid input parameter500 ESB-14-005 Sistem sedang dalam maintenance System under maintenance504 ESB-14-006 Timeout, silakan periksa mutasi rekening Timeout, please check your account statement504 ESB-14-007 Timeout Timeout401 ESB-14-008 client\_id/client\_secret/grant\_type tidak valid Invalid client\_id/client\_secret/grant\_type401 ESB-14-009 Tidak berhak Unauthorized429 ESB-14-010 Jumlah permintaan melebihi limit Limit request exceeded404 ESB-14-011 Service tidak ada Service doesn’t exist401 ESB-14-012 Tidak berhak mengakses service ini Not allowed to access this service400 ESB-99-009 Field Latitude harus diisi Missing Mandatory field Latitude400 ESB-99-009 Field Longitude harus diisi Missing Mandatory field Longitude400 ESB-99-039 Isian input field Count tidak valid Input for field Count not valid400 ESB-99-040 Tipe isian field Count tidak valid Input type for field Count not valid400 ESB-99-040 Tipe isian field Latitude tidak valid Input type for field Latitude not valid400 ESB-99-040 Tipe isian field Longitude tidak valid Input type for field Longitude not valid400 ESB-99-040 Tipe isian field Radius tidak valid Input type for field Radius not valid400 ESB-99-132 Data tidak ditemukan No Data Found400 ESB-99-335 Mata uang tidak valid Invalid CurrencyCode400 ESB-99-339 Tipe Kurs tidak valid Invalid RateType400 ESB-99-392 Nilai SearchBy tidak valid Invalid SearchBy value400 ESB-99-460 Panjang input Value tidak valid Input Value minimal length invalid500 ESB-99-999 Sistem sedang tidak tersedia System unavaliable

#### Informasi Umum

##### 1\. Lokasi ATM BCA

Inquiry ATM information, such as, locate the nearest ATM to user or find by ATM Type, Name, Address or City.

This feature is not yet available to be accessed via Sandbox

**Query String Parameters:**

FieldDataType**Mandatory**DescriptionLatitudeStringYIn format number( (-)3.12 ) Latitude of client location, possible for minus signed value.LongitudeStringYIn format number( (-)3.12 ) Longitude of client location, possible for minus signed value.CountStringN Number of branches to search, in Integer. Maximum value = 20.SearchByStringN Possible values: Address/City/Type/NameRadiusStringNBranch distance (in meters) from client location, in Integer. \*default = 500ValueStringNValue should be filled in :

\- String; If SearchBy = Address or City or Name

\- ATM, ANT, or CRM; If SearchBy = Type

Result of the request will contains following information:

**Response:**

FieldDataTypeDescriptionatm\_detailsArray of ObjectThe object that hold information of the ATM.WSIDStringATM wsidTypeStringATM type eg:

ATM Tarik Tunai

ATM Non-Tunai

ATM Setor TarikNameStringATM nameAddressStringATM full addressCityStringATM cityCountryStringATM countryLatitudeStringATM latitudeLongitudeStringATM longitudeDistanceStringATM distance from client location in meters

##### Error

HTTP CodeError CodeError Message (Indonesian)Error Message (English)404 ESB-07-271 Tidak ada transaksi No transaction400 ESB-07-279 Tanggal awal lebih kecil dari tanggal buka rekening Start date less than account open date400 ESB-14-001 HMAC tidak cocok HMAC mismatch400 ESB-14-002 Permintaan tidak valid Invalid request400 ESB-14-003 Timestamp tidak valid Invalid timestamp400 ESB-14-004 parameter input tidak valid Invalid input parameter500 ESB-14-005 Sistem sedang dalam maintenance System under maintenance504 ESB-14-006 Timeout, silakan periksa mutasi rekening Timeout, please check your account statement504 ESB-14-007 Timeout Timeout401 ESB-14-008 client\_id/client\_secret/grant\_type tidak valid Invalid client\_id/client\_secret/grant\_type401 ESB-14-009 Tidak berhak Unauthorized429 ESB-14-010 Jumlah permintaan melebihi limit Limit request exceeded404 ESB-14-011 Service tidak ada Service doesn’t exist401 ESB-14-012 Tidak berhak mengakses service ini Not allowed to access this service400 ESB-99-009 Field Latitude harus diisi Missing Mandatory field Latitude400 ESB-99-009 Field Longitude harus diisi Missing Mandatory field Longitude400 ESB-99-039 Isian input field Count tidak valid Input for field Count not valid400 ESB-99-040 Tipe isian field Count tidak valid Input type for field Count not valid400 ESB-99-040 Tipe isian field Latitude tidak valid Input type for field Latitude not valid400 ESB-99-040 Tipe isian field Longitude tidak valid Input type for field Longitude not valid400 ESB-99-040 Tipe isian field Radius tidak valid Input type for field Radius not valid400 ESB-99-132 Data tidak ditemukan No Data Found400 ESB-99-335 Mata uang tidak valid Invalid CurrencyCode400 ESB-99-339 Tipe Kurs tidak valid Invalid RateType400 ESB-99-392 Nilai SearchBy tidak valid Invalid SearchBy value400 ESB-99-460 Panjang input Value tidak valid Input Value minimal length invalid500 ESB-99-999 Sistem sedang tidak tersedia System unavaliable

GET /general/info-bca/atm?SearchBy= &Latitude=&Longitude=&Value=&Radius=&Count=

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
GET /general/info-bca/atm?Latitude=6.1900718&Longitude=106.797190 HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
```

GET /general/info-bca/atm?SearchBy= &Latitude=&Longitude=&Value=&Radius=&Count=

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "atm_details": [
      {
        "WSID": "69715",
        "Type": "ATM",
        "Name": "221U-Mayapada Tower 2"
        "Address": "Jl. Jend. Sudirman Kav.28, Jakarta 12920",
        "City": "Jakarta Selatan",
        "Country": "INDONESIA",
        "Latitude": "-6.2112194",
        "Longitude": "106.8201026",
        "Distance": ""8243.97"
      },
      {
        "WSID": "84034",
        "Type": "ATM",
        "Name": "221U-Mayapada Tower 2"
        "Address": "Jl. Jend. Sudirman Kav.28, Jakarta 12920",
        "City": "Jakarta Selatan",
        "Country": "INDONESIA",
        "Latitude": "-6.2112194",
        "Longitude": "106.8201026",
        "Distance": ""8243.97"
      },
      {
        "WSID": "69765",
        "Type": "ATM",
        "Name": "2748-Prudential Tower 2"
        "Address": "Jl. Jend. Sudirman Kav.79 Jakarta Selatan",
        "City": "Jakarta Selatan",
      }
    ]
}
```

##### 2\. Lokasi Cabang BCA

Get information about nearest branch location by distance, address, city or branch type.

**Query String Parameters:**

FieldDataType**Mandatory**DescriptionLatitudeStringYIn format number( (-)3.12 ) Latitude of client location, possible for minus signed value.LongitudeStringYIn format number( (-)3.12 ) Longitude of client location, possible for minus signed value.CountStringN Number of branches to search, in Integer. Maximum value = 20.RadiusStringNBranch distance (in meters) from client location, in Integer. \*default = 500SearchByStringNPossible values: Address/City/TypeValueStringNIf SearchBy = Address or City, value should be filled in : String;

If SearchBy = Type, value should be filled with: kcu , kcp, e-branch, weekendbanking

Result of the request will contains following information:

**Response:**

FieldDataTypeDescriptionBranchDetailsArray of ObjectThe object that hold information of the branch.TypeStringBranch type eg. KCU, KCPNameStringBranch nameAddressStringBranch full addressCityStringBranch cityCountryStringBranch countryLatitudeStringBranch latitudeLongitudeStringBranch longitudeDistanceStringBranch distance from client location in metersE-BranchString“Y” if e-branch reservation is available otherwise it will be set to “N”WeekendBankingOperationStringThe name of the day of weekend, “Saturday”, “Sunday” and “Saturday&Sunday”

##### Error

HTTP CodeError CodeError Message (Indonesian)Error Message (English)404 ESB-07-271 Tidak ada transaksi No transaction400 ESB-07-279 Tanggal awal lebih kecil dari tanggal buka rekening Start date less than account open date400 ESB-14-001 HMAC tidak cocok HMAC mismatch400 ESB-14-002 Permintaan tidak valid Invalid request400 ESB-14-003 Timestamp tidak valid Invalid timestamp400 ESB-14-004 parameter input tidak valid Invalid input parameter500 ESB-14-005 Sistem sedang dalam maintenance System under maintenance504 ESB-14-006 Timeout, silakan periksa mutasi rekening Timeout, please check your account statement504 ESB-14-007 Timeout Timeout401 ESB-14-008 client\_id/client\_secret/grant\_type tidak valid Invalid client\_id/client\_secret/grant\_type401 ESB-14-009 Tidak berhak Unauthorized429 ESB-14-010 Jumlah permintaan melebihi limit Limit request exceeded404 ESB-14-011 Service tidak ada Service doesn’t exist401 ESB-14-012 Tidak berhak mengakses service ini Not allowed to access this service400 ESB-99-009 Field Latitude harus diisi Missing Mandatory field Latitude400 ESB-99-009 Field Longitude harus diisi Missing Mandatory field Longitude400 ESB-99-039 Isian input field Count tidak valid Input for field Count not valid400 ESB-99-040 Tipe isian field Count tidak valid Input type for field Count not valid400 ESB-99-040 Tipe isian field Latitude tidak valid Input type for field Latitude not valid400 ESB-99-040 Tipe isian field Longitude tidak valid Input type for field Longitude not valid400 ESB-99-040 Tipe isian field Radius tidak valid Input type for field Radius not valid400 ESB-99-132 Data tidak ditemukan No Data Found400 ESB-99-335 Mata uang tidak valid Invalid CurrencyCode400 ESB-99-339 Tipe Kurs tidak valid Invalid RateType400 ESB-99-392 Nilai SearchBy tidak valid Invalid SearchBy value400 ESB-99-460 Panjang input Value tidak valid Input Value minimal length invalid500 ESB-99-999 Sistem sedang tidak tersedia System unavaliable

GET /general/info-bca/branch

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
GET /general/info-bca/branch?Latitude=-6.137235&Longitude=106.824928 HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
```

GET /general/info-bca/branch

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "BranchDetails": [
      {
        "Type": "KCP",
        "Name": "Mangga Dua Mal",
        "Address": "Mangga Dua Mal No. RM. 49, Jl. Mangga Dua Raya,,,",
        "City": "Jakarta Pusat",
        "Country": "INDONESIA",
        "Latitude": "-6.137235",
        "Longitude": "106.824928",
        "Distance": "0"
        "E-Branch": "N"
        "WeekendBankingOperation": "Saturday&Sunday"
      },
      {
        "Type": "KCP",
        "Name": "Mangga Dua",
        "Address": "Pst. Perd. Grosir Psr Pagi Lt. 2 Bl. KA No. 001A-B,002 & 012, Jl. Arteri M.D.,,",
        "City": "Jakarta Utara",
        "Country": "INDONESIA",
        "Latitude": "-6.133801",
        "Longitude": "106.822994",
        "Distance": "437.63"
        "E-Branch": "N"
        "WeekendBankingOperation": "-"
      }
    ]
}
```

#### Financing

##### 1\. Batal Faktur

You can cancel your invoice data which had been sent to BCA using this service, at least 1(one) day before disburse date.

Cut-off time: 08.00 P.M - 04.00 A.M (using UTC+07 Time Zone)

Your Request must contain following information:

##### Additional Headers

FieldDataTypeMandatoryDescriptionChannelIDString(5)YChannel’s identifier. Fill it with “95051”.CredentialIDString(10)YChannel’s credential (KlikBCA Bisnis Corporate ID)

##### Request

FieldData Type**Mandatory**DescriptionTransactionIDString(20)Y Transaction ID unique per request (using UTC+07 Time Zone format). Must be 20 in lengthTransactionDateString(10)YTransaction date. Format: yyyyMM-dd. Must be today’s date ReferenceIDString(30)YSender’s invoice number. Must be same as previous created sender’s invoice number.DistributorCodeString(20) YDistributor Code. Format: Alphanumeric and Special CharacterAmount String(16) YTransfer Amount. Format: Number, 13.2. Must be same as previous created invoice transfer amountCurrencyCode String(3) YCurrency Code. Format: Alphanumeric PrincipalCode String(11) YPrincipal Code. Must be 11 in length. Format: Numeric SubPrincipalCode String(11) YSub Principal Code. Must be 11 in length. Disbursement to Sub Principal only. Format: Numeric. DisburseDate String(10) YTransaction date. Format: yyyy-MM-dd. Must be same as previous created invoice disburse date

##### Response

FieldDataTypeDescriptionTransactionIDString(20) Transaction ID unique per request (using UTC+07 Time Zone format). Must be 20 in lengthTransactionDateString(10) Transaction date. Format: yyyy-MM-dd. Must be today’s dateReferenceIDString(30) Sender’s invoice number. Must be uniqueDistributorCode String(20)  Distributor Code. Format: Alphanumeric and Special CharacterAmount String(16)  Transfer Amount. Format: Number, 13.2CurrencyCode String(3)  Currency Code. Format: AlphanumericPrincipalCode String(11)  Principal Code. Must be 11 in length. Format: NumericSubPrincipalCode String(11)  Sub Principal Code. Must be 11 in length. Disbursement to Sub Principal only. Format: Numeric.DisburseDate String(10)  Transaction date. Format: yyyy-MM-dd

##### Error

HTTP CodeError CodeError Message (Indonesian)Error Message (English)400 ESB-01-005 Data tidak ditemukan No Data Found400 ESB-11-005 Kode mata uang tidak vaild Currency not valid400 ESB-11-038 Prinsipal tidak aktif Principal is inactive400 ESB-11-040 Distributor tidak aktif Distributor is not active400 ESB-14-001 HMAC tidak cocok HMAC mismatch400 ESB-14-002 Permintaan tidak valid Invalid request400 ESB-14-003 Timestamp tidak valid Invalid timestamp500 ESB-14-005 Sistem sedang dalam maintenance System under maintenance504 ESB-14-007 Timeout Timeout401 ESB-14-008 client\_id/client\_secret/grant\_type tidak valid Invalid client\_id/client\_secret/grant\_type401 ESB-14-009 Tidak berhak Unauthorized429 ESB-14-010 Jumlah permintaan melebihi limit Limit request exceeded404 ESB-14-011 Service tidak ada Service doesn’t exist401 ESB-14-012 Tidak berhak mengakses service ini Not allowed to access this service504 ESB-14-013 Timeout. Silakan melakukan transaksi kembali Timeout. Please retry your transaction400 ESB-14-015 Content Type tidak valid Invalid Content Type400 ESB-14-016 Format JSON tidak valid Invalid JSON format400 ESB-82-002 Corporate ID tidak valid Invalid Corporate ID400 ESB-82-003 TransactionID tidak unik TransactionID not unique400 ESB-82-026 Lewat batas waktu cut off Cut off time is exceeded400 ESB-99-009 Field (0) harus diisi Missing mandatory field (0)400 ESB-99-113 Transaksi ditolak Transaction rejected400 ESB-99-128 Panjang field (0) melebihi ketentuan Field (0) exceed limit400 ESB-99-284 (0) tidak valid Invalid (0)400 ESB-99-524 ReferenceID sudah diproses pada tanggal (entry\_date) ReferenceID already processed by (entry\_date)400 ESB-99-559 Timeout. Silakan melakukan transaksi kembali Timeout. Please retry your transaction400 ESB-99-568 SubprincipalCode tidak terdaftar SubprincipalCode is not registered400 ESB-99-569 Tidak terdaftar untuk fitur Distributor Financing Not registered for Distributor Financing features400 ESB-99-573 TransactionDate harus diisi dengan tanggal hari ini Transaction Date must be filled with today's date400 ESB-99-655 Subprinsipal tidak aktif Subprincipal is inactive

PUT /df/invoice

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
PUT df/invoice HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
ChannelID: [channelID]
CredentialID: [credentialID]

{
    "TransactionID" : "ABCDEFGHIJ0123456789",
    "TransactionDate" : "2018-02-15",
    "ReferenceID" : "INV0000001",
    "DistributorCode"" : "ABCDE-12345",
    "Amount" : "100000000.00",
    "CurrencyCode" : "IDR",
    "PrincipalCode"" : "00000000001",
    "DisburseDate"" : "2018-02-30",
}

```

PUT /df/invoice

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "TransactionID" : "ABCDEFGHIJ0123456789",
    "TransactionDate" : "2018-02-15",
    "ReferenceID"" : "INV0000001",
    "DistributorCode" : "ABCDE-12345",
    "Amount" : "100000000.00",
    "CurrencyCode" : "IDR",
    "PrincipalCode" : "00000000001",
    "SubPrincipalCode" : "",
    "DisburseDate" : "2018-02-30"
}
```

##### 2\. Kirim Faktur

You can send your invoice data to BCA using this service.

Cut-off time: 08.00 P.M - 04.00 A.M (using UTC+07 Time Zone)

Your Request must contain following information:

##### Additional Headers

FieldDataTypeMandatoryDescriptionChannelIDString(5)YChannel’s identifier. Fill it with “95051”.CredentialIDString(10)YChannel’s credential (KlikBCA Bisnis Corporate ID)

##### Payload

FieldData Type**Mandatory**DescriptionTransactionIDString(20)Y Transaction ID unique per request (using UTC+07 Time Zone format). Must be 20 in lengthTransactionDateString(10)YTransaction date. Format: yyyyMM-dd. Must be today’s date ReferenceIDString(30)YSender’s invoice number. Must be unique DistributorCodeString(20) YDistributor Code. Format: Alphanumeric and Special CharacterAmount String(16) YTransfer Amount. Format: Number, 13.2  CurrencyCode String(3) YCurrency Code. Format: Alphanumeric PrincipalCode String(11) YPrincipal Code. Must be 11 in length. Format: Numeric SubPrincipalCode String(11) YSub Principal Code. Must be 11 in length. Disbursement to Sub Principal only. Format: Numeric. DisburseDate String(10) YTransaction date. Format: yyyyMM-dd

##### Response

FieldDataTypeDescriptionTransactionIDString(20) Transaction ID unique per request (using UTC+07 Time Zone format). Must be 20 in lengthTransactionDateString(10) Transaction date. Format: yyyy-MM-dd. Must be today’s dateReferenceIDString(30) Sender’s invoice number. Must be uniqueDistributorCode String(20)  Distributor Code. Format: Alphanumeric and Special CharacterAmount String(16)  Transfer Amount. Format: Number, 13.2CurrencyCode String(3)  Currency Code. Format: AlphanumericPrincipalCode String(11)  Principal Code. Must be 11 in length. Format: NumericSubPrincipalCode String(11)  Sub Principal Code. Must be 11 in length. Disbursement to Sub Principal only. Format: Numeric.DisburseDate String(10)  Transaction date. Format: yyyy-MM-dd

##### Error

HTTP CodeError CodeError Message (Indonesian)Error Message (English)400 ESB-01-005 Data tidak ditemukan No Data Found400 ESB-11-005 Kode mata uang tidak vaild Currency not valid400 ESB-11-038 Prinsipal tidak aktif Principal is inactive400 ESB-11-040 Distributor tidak aktif Distributor is not active400 ESB-14-001 HMAC tidak cocok HMAC mismatch400 ESB-14-002 Permintaan tidak valid Invalid request400 ESB-14-003 Timestamp tidak valid Invalid timestamp500 ESB-14-005 Sistem sedang dalam maintenance System under maintenance504 ESB-14-007 Timeout Timeout401 ESB-14-008 client\_id/client\_secret/grant\_type tidak valid Invalid client\_id/client\_secret/grant\_type401 ESB-14-009 Tidak berhak Unauthorized429 ESB-14-010 Jumlah permintaan melebihi limit Limit request exceeded404 ESB-14-011 Service tidak ada Service doesn’t exist401 ESB-14-012 Tidak berhak mengakses service ini Not allowed to access this service504 ESB-14-013 Timeout. Silakan melakukan transaksi kembali Timeout. Please retry your transaction400 ESB-14-015 Content Type tidak valid Invalid Content Type400 ESB-14-016 Format JSON tidak valid Invalid JSON format400 ESB-82-002 Corporate ID tidak valid Invalid Corporate ID400 ESB-82-003 TransactionID tidak unik TransactionID not unique400 ESB-82-026 Lewat batas waktu cut off Cut off time is exceeded400 ESB-99-009 Field (0) harus diisi Missing mandatory field (0)400 ESB-99-113 Transaksi ditolak Transaction rejected400 ESB-99-128 Panjang field (0) melebihi ketentuan Field (0) exceed limit400 ESB-99-284 (0) tidak valid Invalid (0)400 ESB-99-524 ReferenceID sudah diproses pada tanggal (entry\_date) ReferenceID already processed by (entry\_date)400 ESB-99-559 Timeout. Silakan melakukan transaksi kembali Timeout. Please retry your transaction400 ESB-99-568 SubprincipalCode tidak terdaftar SubprincipalCode is not registered400 ESB-99-569 Tidak terdaftar untuk fitur Distributor Financing Not registered for Distributor Financing features400 ESB-99-573 TransactionDate harus diisi dengan tanggal hari ini Transaction Date must be filled with today's date400 ESB-99-655 Subprinsipal tidak aktif Subprincipal is inactive

POST /df/invoice

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST df/invoice HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
ChannelID: [channelID]
CredentialID: [credentialID]

{
    "TransactionID" : "ABCDEFGHIJ0123456789",
    "TransactionDate" : "2018-02-15",
    "ReferenceID" : "INV0000001",
    "DistributorCode"" : "ABCDE-12345",
    "Amount" : "100000000.00",
    "CurrencyCode" : "IDR",
    "PrincipalCode"" : "00000000001",
    "DisburseDate"" : "2018-02-30",
}
```

POST /df/invoice

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "TransactionID" : "ABCDEFGHIJ0123456789",
    "TransactionDate" : "2018-02-15",
    "ReferenceID"" : "INV0000001",
    "DistributorCode" : "ABCDE-12345",
    "Amount" : "100000000.00",
    "CurrencyCode" : "IDR",
    "PrincipalCode" : "00000000001",
    "SubPrincipalCode" : "",
    "DisburseDate" : "2018-02-30"
}
```

##### 3\. Notifikasi Pembayaran Faktur

Service for BCA to send notification of disbursement.

This feature is not yet available to be accessed via Sandbox

##### Additional Headers

FieldDataTypeMandatoryDescriptionX-Pass-SignatureString(5)YThis is a header field, use to replace the “Pass” field on everybody payload

that contains signature. This field occurred on every BCA request to Copartner’s API.

##### Request

FieldData Type**Mandatory**DescriptionTransactionIDString(20)Y UUID as unique ID for every requestTransactionDateString(10)YTransaction date. Format: yyyyMM-dd. Must be today’s date ReferenceIDString(30)YInvoice number. Must be uniqueDistributorCodeString(20) YDistributor Code. Format: Alphanumeric and Special CharacterAmount String(16) YTransfer Amount. Format: Number, 13.2CurrencyCode String(3) YCurrency Code. Format: Alphanumeric PrincipalCode String(11) YPrincipal Code. Must be 11 in length. Format: Numeric SubPrincipalCode String(11) YSub Principal Code. Format : NumericDisburseDate String(10) YTransaction date. Format: yyyy-MM-dd

##### Response

FieldDataTypeDescriptionTransactionIDString(20) Same TransactionID from request. It’s a UUIDTransactionDateString(10) Transaction date. Format: yyyy-MM-dd. Must be today’s dateReferenceIDString(30) Invoice number. Must be uniqueDistributorCode String(20)  Distributor Code. Format: Alphanumeric and Special CharacterAmount String(16)  Transfer Amount. Format: Number, 13.2CurrencyCode String(3)  Currency Code. Format: AlphanumericPrincipalCode String(11)  Principal Code. Must be 11 in length. Format: NumericSubPrincipalCode String(11)  Sub Principal Code. Must be 11 in length. Disbursement to Sub Principal only. Format: Numeric.DisburseDate String(10)  Transaction date. Format: yyyy-MM-dd

POST https://www.copartners.com/bca-df/disbursement/notification

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST https://www.copartners.com/bca-df/disbursement/notification HTTP/1.1
X-Pass-Signature: b9125ab10816f6929d214c96ffca77dfd5a1ea13856362b85eeaeb70155

{
    "TransactionID" : "65d5dd3474b44c5abbb0ce1cb5ff563f",",
    "TransactionDate" : "2018-02-15",
    "ReferenceID"" : "INV001923",",
    "PrincipalCode"" : "00000000001",
    "SubPrincipalCode"" : "",
    "DistributorCode" : "ABCDE-12345",
    "Amount" : "100000000.00",
    "CurrencyCode" : "IDR",
}

```

POST https://www.copartners.com/bca-df/disbursement/notification

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "TransactionID" : "65d5dd3474b44c5abbb0ce1cb5ff563f",",
    "TransactionDate" : "2018-02-15",
    "ReferenceID"" : "INV001923",",
    "PrincipalCode"" : "00000000001",
    "SubPrincipalCode"" : "00000000001",
    "DistributorCode" : "ABCDE-12345",
    "Amount" : "100000000.00",
    "CurrencyCode" : "IDR",
}
```

##### 4\. Notifikasi Pembayaran Pinjaman

Service for BCA to send notification of transaction limit update.

This feature is not yet available to be accessed via Sandbox

##### Additional Headers

FieldDataTypeMandatoryDescriptionX-Pass-SignatureString(5)YThis is a header field, use to replace the “Pass” field on everybody payload

that contains signature. This field occurred on every BCA request to Copartner’s API.

##### Request

FieldData Type**Mandatory**DescriptionTransactionIDString(32)Y UUID as unique ID for every requestDistributorCodeString(20) YDistributor Code. Format: Alphanumeric and Special CharacterPrincipalCode String(11) YPrincipal Code. Format: NumericAmountString(16) YTransfer Amount. Minus sign ( - ) is applied for corrective payment. Format: Number, 13.2.TransactionDateString(10) YTransaction date. Format: yyyy-MM-dd

##### Response

FieldDataTypeDescriptionTransactionIDString(32) Same TransactionID from request. It’s a UUIDDistributorCode String(20)  Same Distributor Code from request.

POST https://www.copartners.com/bca-df/txn-limit/notification

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST https://www.copartners.com/bca-df/txn-limit/notification HTTP/1.1
X-Pass-Signature: b9125ab10816f6929d214c96ffca77dfd5a1ea13856362b85eeaeb70155

{
    "TransactionID" : "65d5dd3474b44c5abbb0ce1cb5ff563f",",
    "PrincipalCode"" : "00000000001",
    "DistributorCode"" : "ABCDE-12345",
    "Amount" : "100000000.00",
    "TransactionDate" : "2018-02-15",
}
```

POST https://www.copartners.com/bca-df/txn-limit/notification

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "TransactionID" : "65d5dd3474b44c5abbb0ce1cb5ff563f",",
    "DistributorCode" : "ABCDE-12345",
}
```

##### 5\. Notifikasi Status Fasilitas Kredit

Service for BCA to send notification of transaction limit update.

This feature is not yet available to be accessed via Sandbox

##### Additional Headers

FieldDataTypeMandatoryDescriptionX-Pass-SignatureString(5)YThis is a header field, use to replace the “Pass” field on everybody payload

that contains signature. This field occurred on every BCA request to Copartner’s API.

##### Request

FieldData Type**Mandatory**DescriptionTransactionIDString(32)Y UUID as unique ID for every requestDistributorCodeString(20) YDistributor Code. Format: Alphanumeric and Special CharacterPrincipalCode String(11) YPrincipal Code. Must be 11 in length. Format: Numeric StopSupplyFlagString(1) Y0 : Stop (Distributor Non-Aktif/Wanprestasi); 1: Open (Distributor Aktif)TransactionDateString(10) YTransaction date. Format: yyyy-MM-dd

##### Response

FieldDataTypeDescriptionTransactionIDString(32) Same TransactionID from request. It’s a UUIDDistributorCode String(20)  Same Distributor Code from request.

POST https://www.copartners.com/bca-df/stop-supply/notification

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST https://www.copartners.com/bca-df/stop-supply/notification HTTP/1.1
X-Pass-Signature: b9125ab10816f6929d214c96ffca77dfd5a1ea13856362b85eeaeb70155

{
    "TransactionID" : "65d5dd3474b44c5abbb0ce1cb5ff563f",",
    "DistributorCode"" : "ABCDE-12345",
    "PrincipalCode"" : "00000000001",
    "TransactionDate" : "2018-02-15",
    "StopSupplyFlag"" : "1",
}

```

POST https://www.copartners.com/bca-df/stop-supply/notification

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "TransactionID" : "65d5dd3474b44c5abbb0ce1cb5ff563f",",
    "DistributorCode" : "ABCDE-12345",
}
```

#### FIRE

##### 1\. Amandemen Kiriman Uang Tunai

Provides service for Amendment “Cash Transfer” to Non account holder. Yout request must contain following information:

##### Payload

**Field****DataType****Mandatory****Description****Authentication**CorporateIDString(6)YAgen code used to connect to FIRe systemAccessCodeString(20)YPassword that will be read by system for each transaction requestBranchCodeString(8)YFI branch code of API clientUserIDString(12)YUser IDLocalIDString(15)YBranch code from FI sub branch of API client**AmendmentDetails**SenderDetailsFirstNameString(35)YSender’s first nameLastNameString(35)NSender’s last nameDateOfBirthString(8)NSender’s date of birth.

Format : ddmmyyyyAddress1String(35)YSender’s first addressAddress2String(35)NSender’s second addressCityString(18)YSender’s cityStateIDString(2)NState IDPostalCodeString(10)NSender’s postal codeCountryIDString(2)YSender’s country initial code.

Example : Indonesia =IDMobileString(20)NSender’s mobile phone numberIdentificationTypeString(2)NSender’s identity typeIdentificationNumberString(24)NSender’s identity number**BenficiaryDetails**NameString(35)YBeneficiary’s nameDateOfBirthString(8)NBeneficiary’s date of birth.

Format : ddmmyyyyAddress1String(35)NBeneficiary’s first addressAddress2String(35)NBeneficiary’s second addressCityString(35)NBeneficiary’s cityStateIDString(2)State IDPostalCodeString(10)NBeneficiary’s postal codeCountryIDString(2)NBeneficiary’s country initial code.

Example : Indonesia =IDMobileString(20)NBeneficiary’s mobile phone numberIdentificationTypeString(2)NBeneficiary’s identity typeIdentificationNumberString(24)NBeneficiary’s identity numberNationalityIDString(2)NCountry initial code of beneficiary’s nationality.

Example : Indonesia =IDOccupationString(30)NBeneficiary occupation**TransactionDetails**Description1String(35)NFirst DescriptionDescription2String(35)NSecond DescriptionSecretQuestionString(100)NQuestionSecretAnswerString(100)NAnswer**TransactionDetails**FormNumberString(16)YFI Ref

Result of the request will contains following information:

##### Response

**Field****DataType****Mandatory****Description**StatusTransactionString(4)YResponseCode and Sub-ResponseCode of transactionStatusMessageString(100)YDescription of StatusTransaction**SenderDetails**FirstNameString(35)NSender’s first nameLastNameString(35)NSender’s last nameDateOfBirthString(8)NSender’s date of birth. Format : ddmmyyyyAddress1String(35)NSender’s first addressAddress2String(35)NSender’s second addressCityString(18)NSender’s cityStateIDString(2)NState IDPostalCodeString(10)NSender’s postal codeCountryIDString(2)NSender’s country initial code. Example : Indonesia =IDMobileString(20)NSender’s mobile phone numberIdentificationTypeString(2)NSender’s identity typeIdentificationNumberString(24)NSender’s identity number**BeneficiaryDetails**NameString(35)NBeneficiary’s nameDateOfBirthString(8)NBeneficiary’s date of birth. Format : ddmmyyyyAddress1String(35)NBeneficiary’s first addressAddress2String(35)NBeneficiary’s second addressCityString(35)NBeneficiary’s cityStateIDString(2)NState IDPostalCodeString(10)NBeneficiary’s postal codeCountryIDString(2)NBeneficiary’s country initial code.Example : Indonesia =IDMobileString(20)NBeneficiary’s mobile phone numberIdentificationTypeString(2)NBeneficiary’s identity type. (PP = Passport, ID = Identity, DR = Dr. License)IdentificationNumberString(24)NBeneficiary’s identity numberNationalityIDString(2)NCountry initial code of beneficiary’s nationality. Example : Indonesia =IDOccupationString(30)NBeneficiary occupation**TransactionDetails**Description1String(35)NFirst descriptionDescription2String(35)NSecond descriptionSecretQuestionString(100)NQuestionSecretAnswerString(100)NAnswerFormNumberString(16)NFI Ref

##### Error

HTTP CodeError CodeError Message (Indonesian)Error Message (English)400 ESB-14-001 HMAC tidak cocok HMAC mismatch400 ESB-14-002 Permintaan tidak valid Invalid request400 ESB-14-003 Timestamp tidak valid Invalid timestamp400 ESB-14-004 Parameter input tidak valid Invalid input parameter503 ESB-14-005 Sistem sedang dalam maintenance System under maintenance504 ESB-14-006 Timeout, silakan periksa mutasi rekening Timeout, please check your account statement504 ESB-14-007 Timeout Timeout401 ESB-14-008 client\_id/client\_secret/grant\_type tidak valid Invalid client\_id/client\_secret/grant\_type401 ESB-14-009 Tidak berhak Unauthorized429 ESB-14-010 Jumlah permintaan melebihi limit Limit request exceeded404 ESB-14-011 Service tidak ada Service doesn't exist401 ESB-14-012 Tidak berhak mengakses service ini Not allowed to access this service400 0101Field \[fieldName\] is required Field \[fieldName\] is required400 0102Field \[fieldName\] is larger than expected Field \[fieldName\] is larger than expected400 0103Data type is invalid at field \[fieldName\] Data type is invalid at field \[fieldName\]409 0201Ref pengirim sama Same Sender Ref409 0202PIN sudah pernah dipakai PIN already in used400 0203Format mata uang tidak valid Invalid Currency format400 0204BIC Code tidak valid Invalid BIC Code400 0205ABA Code tidak valid Invalid ABA Code400 0206CHIPS Code tidak valid Invalid CHIPS Code400 0207NID Code tidak valid Invalid NID Code400 0208Country Code untuk NID Code tidak valid Invalid Country Code for NID Code401 0302Access Code salah Wrong Access Code400 0303Branch Code tidak ditemukan Unknown Branch Code400 0304User ID tidak ditemukan Unknown User ID401 0306Tidak berhak Not Authorized Privilege400 0401PRODUK/MATA UANG TIDAK DITEMUKAN PRODUCT/CURRENCY NOT FOUND400 0402LIMIT PRODUK TIDAK CUKUP PRODUCT LIMIT NOT ENOUGH400 0403LIMIT RELEASER TIDAK CUKUP LIMIT RELEASER NOT ENOUGH400 0404ERROR PDN LIMIT ERROR PDN LIMIT400 0405GAGAL VALIDASI REKENING PENERIMA FAILED TO VALIDATE BENEFICIARY ACCOUNT400 0406REKENING PENERIMA TIDAK VALID INVALID BENEFICIARY ACCOUNT400 0407SALDO TIDAK CUKUP INSUFFICIENT FUND404 0408TRANSAKSI TIDAK DITEMUKAN TRANSACTION NOT FOUND400 0409REKENING PENERIMA TUTUP CLOSED BENEFICIARY ACCOUNT500 0410TRANSAKSI TIDAK DAPAT DIPROSES TRANSACTION CAN NOT BE PROCESSED400 0412BUKAN MEMBER SWITCHING NON SWITCHING MEMBER400 0413REKENING TIDAK VALID INVALID ACCOUNT NUMBER400 0414DITOLAK KARENA SALDO TIDAK CUKUP REJECT DUE INSUFFICIENT FUND503 0415DITOLAK KARENA CUT OFF SISTEM BCA REJECT DUE TO CUT OFF SYSTEM BCA500 0901Error Database Error Database503 0902Sistem sedang dalam maintenance System Under Maintenance503 0903Awal hari error Begin of Day Error500 0999Error sistem System error

PUT https://sandbox.bca.co.id/fire/transaction/cash-transfer/amend

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST /fire/transactions/cash-transfer/amend HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]

{
    "Authentication":{
        "CorporateID":"ABCDE",
        "AccessCode":"k1oawj8tygkrY9tBVkt",
        "BranchCode":"BCAAPI01",
        "UserID":"BCAUSERID001",
        "LocalID":"BCA01"
    },
    "AmendmentDetails":{
        "SenderDetails":{
            "FirstName":"ERIK HERNANDEZ CORTES",
            "LastName":"",
            "DateOfBirth":"",
            "Address1":"4TH AVENUE MANHATTAN",
            "Address2":"",
            "City":"NEW YORK",
            "StateID":"",
            "PostalCode":"",
            "CountryID":"US",
            "Mobile":"",
            "IdentificationType":"",
            "IdentificationNumber":""
        },
        "BeneficiaryDetails":{
            "Name":"ADRIANA PEREZ ENRRIQUEZ",
            "DateOfBirth":"",
            "Address1":"MAIN STREET 512",
            "Address2":"",
            "City":"",
            "StateID":"",
            "PostalCode":"",
            "CountryID":"ID",
            "Mobile":"",
            "IdentificationType":"",
            "IdentificationNumber":"",
            "NationalityID":"",
            "Occupation":""
        },
        "TransactionDetails":{
            "SecretQuestion":"",
            "SecretAnswer":"",
            "Description1":"",
            "Description2":""
        }
    },
    "TransactionDetails":{
        "FormNumber":"CT293 IDR2D"
    }
}
```

PUT https://sandbox.bca.co.id/fire/transaction/cash-transfer/amend

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "SenderDetails":
    {
        "FirstName":"ERIK HERNANDEZ CORTES",
        "LastName":"",
        "DateOfBirth":"",
        "Address1":"4TH AVENUE MANHATTAN",
        "Address2":"",
        "City":"NEW YORK",
        "StateID":"",
        "PostalCode":"",
        "CountryID":"US",
        "Mobile":"",
        "IdentificationType":"",
        "IdentificationNumber":""
    },
    "BeneficiaryDetails":
    {
        "Name":"ADRIANA PEREZ ENRRIQUEZ",
        "DateOfBirth":"",
        "Address1":"MAIN STREET 512",
        "Address2":"",
        "City":"",
        "StateID":"",
        "PostalCode":"",
        "CountryID":"ID",
        "Mobile":"",
        "IdentificationType":"",
        "IdentificationNumber":"",
        "NationalityID":"",
        "Occupation":""
    },
    "TransactionDetails":
    {
        "Description1":"",
        "Description2":"",
        "SecretQuestion":"",
        "SecretAnswer":"",
        "FormNumber":"CT293_IDR2D"
    },
    "StatusTransaction":"0000",
    "StatusMessage":"Success"
}
```

##### 2\. Informasi Rekening Penerima

Provides service to Inquiry BCA’s Account name or Other Bank Switching’s Account name.

Your request must contain following information:

##### **Payload**

**Field****DataType****Mandatory****Description****Authentication**

CorporateID String(6) YAgent code used to connect to FIRe system AccessCode String(20) YPassword that will be read by system for each transaction request BranchCode String(8) YFI branch code of API client UserID String(12) YUserID  LocalID String(15) YBranch code from FI sub branch of API client**BeneficiaryDetails** BankCodeTypeString(3) YBeneficiary’s SWIFT code type from PY04 table.

(BIC, NID, CHP (CHIPS), ABA, NAM(NAME)) BankCodeValueString(23) Y AccountNumberString(34) Y Beneficiary bank’s account number

Result of the request will contains following information:

##### Response

**Field****DataType****Mandatory****Description****BeneficiaryDetails** ServerBeneAccountName String(35) NName of beneficiary according to application (bank / switcher) StatusTransaction String(4) YResponseCode and Sub-ResponseCode of transaction StatusMessage String(100) YDescription of StatusTransaction

##### Error

HTTP CodeError CodeError Message (Indonesian)Error Message (English)400 ESB-14-001 HMAC tidak cocok HMAC mismatch400 ESB-14-002 Permintaan tidak valid Invalid request400 ESB-14-003 Timestamp tidak valid Invalid timestamp400 ESB-14-004 Parameter input tidak valid Invalid input parameter503 ESB-14-005 Sistem sedang dalam maintenance System under maintenance504 ESB-14-006 Timeout, silakan periksa mutasi rekening Timeout, please check your account statement504 ESB-14-007 Timeout Timeout401 ESB-14-008 client\_id/client\_secret/grant\_type tidak valid Invalid client\_id/client\_secret/grant\_type401 ESB-14-009 Tidak berhak Unauthorized429 ESB-14-010 Jumlah permintaan melebihi limit Limit request exceeded404 ESB-14-011 Service tidak ada Service doesn't exist401 ESB-14-012 Tidak berhak mengakses service ini Not allowed to access this service400 0101Field \[fieldName\] is required Field \[fieldName\] is required400 0102Field \[fieldName\] is larger than expected Field \[fieldName\] is larger than expected400 0103Data type is invalid at field \[fieldName\] Data type is invalid at field \[fieldName\]409 0201Ref pengirim sama Same Sender Ref409 0202PIN sudah pernah dipakai PIN already in used400 0203Format mata uang tidak valid Invalid Currency format400 0204BIC Code tidak valid Invalid BIC Code400 0205ABA Code tidak valid Invalid ABA Code400 0206CHIPS Code tidak valid Invalid CHIPS Code400 0207NID Code tidak valid Invalid NID Code400 0208Country Code untuk NID Code tidak valid Invalid Country Code for NID Code401 0302Access Code salah Wrong Access Code400 0303Branch Code tidak ditemukan Unknown Branch Code400 0304User ID tidak ditemukan Unknown User ID401 0306Tidak berhak Not Authorized Privilege400 0401PRODUK/MATA UANG TIDAK DITEMUKAN PRODUCT/CURRENCY NOT FOUND400 0402LIMIT PRODUK TIDAK CUKUP PRODUCT LIMIT NOT ENOUGH400 0403LIMIT RELEASER TIDAK CUKUP LIMIT RELEASER NOT ENOUGH400 0404ERROR PDN LIMIT ERROR PDN LIMIT400 0405GAGAL VALIDASI REKENING PENERIMA FAILED TO VALIDATE BENEFICIARY ACCOUNT400 0406REKENING PENERIMA TIDAK VALID INVALID BENEFICIARY ACCOUNT400 0407SALDO TIDAK CUKUP INSUFFICIENT FUND404 0408TRANSAKSI TIDAK DITEMUKAN TRANSACTION NOT FOUND400 0409REKENING PENERIMA TUTUP CLOSED BENEFICIARY ACCOUNT500 0410TRANSAKSI TIDAK DAPAT DIPROSES TRANSACTION CAN NOT BE PROCESSED400 0412BUKAN MEMBER SWITCHING NON SWITCHING MEMBER400 0413REKENING TIDAK VALID INVALID ACCOUNT NUMBER400 0414DITOLAK KARENA SALDO TIDAK CUKUP REJECT DUE INSUFFICIENT FUND503 0415DITOLAK KARENA CUT OFF SISTEM BCA REJECT DUE TO CUT OFF SYSTEM BCA500 0901Error Database Error Database503 0902Sistem sedang dalam maintenance System Under Maintenance503 0903Awal hari error Begin of Day Error500 0999Error sistem System error

POST https://sandbox.bca.co.id/fire/accounts

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST /fire/accounts HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]

{
    "Authentication":{
        "CorporateID":"POSTID",
        "AccessCode":"jSvazuvEnz2mavv1wj6L",
        "BranchCode":"POSTID01",
        "UserID":"OPRPOS1",
        "LocalID":"2370000"
    },
    "BeneficiaryDetails":{
        "BankCodeType":"BIC",
        "BankCodeValue":"CENAIDJAXXX",
        "AccountNumber":"0106666011"
    }
}
```

POST https://sandbox.bca.co.id/fire/accounts

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "BeneficiaryDetails":
    {
        "ServerBeneAccountName":"ERWINA TAUFIK"
    },
    "StatusTransaction":"0000",
    "StatusMessage":"Success"
}
```

##### 3\. Informasi Saldo Rekening Mitra

Provides service to Inquiry balance for Vostro’s Account.

Your request must contain following information:

##### Payload

FieldDataTypeMandatory**Description****Authentication**CorporateIDString(6)YAgent code used to connect to FIRe systemAccessCodeString(20)YPassword that will be read by system for each transaction requestBranchCodeString(8) YFI branch code of API client UserId String(12)Y  User ID LocalIdString(15) Y  Branch code from FI sub branch of API client**FIDetails** AccountNumberString(34) Y FI sender's account number thar registered in FIRe

Result of the request will contains following information:

##### Response

FieldDataTypeMandatory**Description**StatusTransactionString(4)YResponseCode and Sub-ResponseCode of transactionStatusMessageString(100)YDescription of StatusTransaction**FIDetails**CurrencyIDString(3) Y FI sender’s currency that registered in FIReAccountBalanceString(20) Y FI sender’s account balance that registered in FIRe. Format : number(17.2)

##### Error

HTTP CodeError CodeError Message (Indonesian)Error Message (English)400 ESB-14-001 HMAC tidak cocok HMAC mismatch400 ESB-14-002 Permintaan tidak valid Invalid request400 ESB-14-003 Timestamp tidak valid Invalid timestamp400 ESB-14-004 Parameter input tidak valid Invalid input parameter503 ESB-14-005 Sistem sedang dalam maintenance System under maintenance504 ESB-14-006 Timeout, silakan periksa mutasi rekening Timeout, please check your account statement504 ESB-14-007 Timeout Timeout401 ESB-14-008 client\_id/client\_secret/grant\_type tidak valid Invalid client\_id/client\_secret/grant\_type401 ESB-14-009 Tidak berhak Unauthorized429 ESB-14-010 Jumlah permintaan melebihi limit Limit request exceeded404 ESB-14-011 Service tidak ada Service doesn't exist401 ESB-14-012 Tidak berhak mengakses service ini Not allowed to access this service400 0101Field \[fieldName\] is required Field \[fieldName\] is required400 0102Field \[fieldName\] is larger than expected Field \[fieldName\] is larger than expected400 0103Data type is invalid at field \[fieldName\] Data type is invalid at field \[fieldName\]409 0201Ref pengirim sama Same Sender Ref409 0202PIN sudah pernah dipakai PIN already in used400 0203Format mata uang tidak valid Invalid Currency format400 0204BIC Code tidak valid Invalid BIC Code400 0205ABA Code tidak valid Invalid ABA Code400 0206CHIPS Code tidak valid Invalid CHIPS Code400 0207NID Code tidak valid Invalid NID Code400 0208Country Code untuk NID Code tidak valid Invalid Country Code for NID Code401 0302Access Code salah Wrong Access Code400 0303Branch Code tidak ditemukan Unknown Branch Code400 0304User ID tidak ditemukan Unknown User ID401 0306Tidak berhak Not Authorized Privilege400 0401PRODUK/MATA UANG TIDAK DITEMUKAN PRODUCT/CURRENCY NOT FOUND400 0402LIMIT PRODUK TIDAK CUKUP PRODUCT LIMIT NOT ENOUGH400 0403LIMIT RELEASER TIDAK CUKUP LIMIT RELEASER NOT ENOUGH400 0404ERROR PDN LIMIT ERROR PDN LIMIT400 0405GAGAL VALIDASI REKENING PENERIMA FAILED TO VALIDATE BENEFICIARY ACCOUNT400 0406REKENING PENERIMA TIDAK VALID INVALID BENEFICIARY ACCOUNT400 0407SALDO TIDAK CUKUP INSUFFICIENT FUND404 0408TRANSAKSI TIDAK DITEMUKAN TRANSACTION NOT FOUND400 0409REKENING PENERIMA TUTUP CLOSED BENEFICIARY ACCOUNT500 0410TRANSAKSI TIDAK DAPAT DIPROSES TRANSACTION CAN NOT BE PROCESSED400 0412BUKAN MEMBER SWITCHING NON SWITCHING MEMBER400 0413REKENING TIDAK VALID INVALID ACCOUNT NUMBER400 0414DITOLAK KARENA SALDO TIDAK CUKUP REJECT DUE INSUFFICIENT FUND503 0415DITOLAK KARENA CUT OFF SISTEM BCA REJECT DUE TO CUT OFF SYSTEM BCA500 0901Error Database Error Database503 0902Sistem sedang dalam maintenance System Under Maintenance503 0903Awal hari error Begin of Day Error500 0999Error sistem System error

POST https://sandbox.bca.co.id/fire/accounts/balance

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST /fire/accounts/balance HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]

{
    "Authentication":{
        "CorporateID":"POSTID",
        "AccessCode":"jSvazuvEnz2mavv1wj6L",
        "BranchCode":"POSTID01",
        "UserID":"OPRPOS1,
        "LocalID":"40115"
    },
    "FIDetails":{
        "AccountNumber":"0012323008"
    }
}
```

POST https://sandbox.bca.co.id/fire/accounts/balance

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "FIDetails":
    {
        "CurrencyID":"USD",
        "AccountBalance":"228277.56"
    },
    "StatusTransaction":"0000",
    "StatusMessage":"Success"
}
```

##### 4\. Informasi Status Kiriman Uang

Provides service to Inquiry Transaction that has been submitted before. Your request must contain following information:

##### Request

FieldDataTypeMandatory**Description****Authentication**CorporateIDString(6)YAgent code used to connect to FIRe system

AccessCodeString(20)YPassword that will be read by system for each transaction request

BranchCodeString(8)Y FI branch code of API clientUserID

String(12)YUser IDLocalID

String(15)YBranch code from FI sub branch of API client**TransactionDetails**InquiryByString(1)Y N : PIN NON

F : Form Number

B : BCARef NumberInquiryValueString(35)Y Value for PIN NON/Form Number/ BCARef Number

Result of the request will contains following information:

**Response**

**Field**

**DataType****Mandatory****Description**StatusTransactionString(4)YResponseCode and Sub-ResponseCode of transctionStatusMessageString(100)YDescription of StatusTransaction**SenderDetails**FirstNameString(35)NSender’s first nameLastNameString(35)NSender’s last name **BeneficiaryDetails**NameString(35)NBeneficiary’s nameBankCodeTypeString(3)NBeneficiary’s SWIFT code type from PY04 table. (BIC, NID, CHP(CHIPS), ABA, NAM(NAME))BankCodeValueString(23)NValue for SWIFT code typeAccountNumberString(34)NBeneficiary bank’s account number**TransactionDetails**AmountPaidString(20)NTransaction value.Format: Number(17.2)CurrencyIDString(3)NCurrency codeReleaseDateTimeString(24)NTransaction request time by FILocalIDString(15)NLocal ID of Last Processed Non-Inquiry TransactionFormNumberString(16)NFI RefReferenceNumberString(25)NRef number FI/BCAPINString(35)NPersonal Identification Number for encashmentDescription1String(35)NFirst DescriptionDescription2String(35)NSecond description

##### Error

HTTP CodeError CodeError Message (Indonesian)Error Message (English)400 ESB-14-001 HMAC tidak cocok HMAC mismatch400 ESB-14-002 Permintaan tidak valid Invalid request400 ESB-14-003 Timestamp tidak valid Invalid timestamp400 ESB-14-004 Parameter input tidak valid Invalid input parameter503 ESB-14-005 Sistem sedang dalam maintenance System under maintenance504 ESB-14-006 Timeout, silakan periksa mutasi rekening Timeout, please check your account statement504 ESB-14-007 Timeout Timeout401 ESB-14-008 client\_id/client\_secret/grant\_type tidak valid Invalid client\_id/client\_secret/grant\_type401 ESB-14-009 Tidak berhak Unauthorized429 ESB-14-010 Jumlah permintaan melebihi limit Limit request exceeded404 ESB-14-011 Service tidak ada Service doesn't exist401 ESB-14-012 Tidak berhak mengakses service ini Not allowed to access this service400 0101Field \[fieldName\] is required Field \[fieldName\] is required400 0102Field \[fieldName\] is larger than expected Field \[fieldName\] is larger than expected400 0103Data type is invalid at field \[fieldName\] Data type is invalid at field \[fieldName\]409 0201Ref pengirim sama Same Sender Ref409 0202PIN sudah pernah dipakai PIN already in used400 0203Format mata uang tidak valid Invalid Currency format400 0204BIC Code tidak valid Invalid BIC Code400 0205ABA Code tidak valid Invalid ABA Code400 0206CHIPS Code tidak valid Invalid CHIPS Code400 0207NID Code tidak valid Invalid NID Code400 0208Country Code untuk NID Code tidak valid Invalid Country Code for NID Code401 0302Access Code salah Wrong Access Code400 0303Branch Code tidak ditemukan Unknown Branch Code400 0304User ID tidak ditemukan Unknown User ID401 0306Tidak berhak Not Authorized Privilege400 0401PRODUK/MATA UANG TIDAK DITEMUKAN PRODUCT/CURRENCY NOT FOUND400 0402LIMIT PRODUK TIDAK CUKUP PRODUCT LIMIT NOT ENOUGH400 0403LIMIT RELEASER TIDAK CUKUP LIMIT RELEASER NOT ENOUGH400 0404ERROR PDN LIMIT ERROR PDN LIMIT400 0405GAGAL VALIDASI REKENING PENERIMA FAILED TO VALIDATE BENEFICIARY ACCOUNT400 0406REKENING PENERIMA TIDAK VALID INVALID BENEFICIARY ACCOUNT400 0407SALDO TIDAK CUKUP INSUFFICIENT FUND404 0408TRANSAKSI TIDAK DITEMUKAN TRANSACTION NOT FOUND400 0409REKENING PENERIMA TUTUP CLOSED BENEFICIARY ACCOUNT500 0410TRANSAKSI TIDAK DAPAT DIPROSES TRANSACTION CAN NOT BE PROCESSED400 0412BUKAN MEMBER SWITCHING NON SWITCHING MEMBER400 0413REKENING TIDAK VALID INVALID ACCOUNT NUMBER400 0414DITOLAK KARENA SALDO TIDAK CUKUP REJECT DUE INSUFFICIENT FUND503 0415DITOLAK KARENA CUT OFF SISTEM BCA REJECT DUE TO CUT OFF SYSTEM BCA500 0901Error Database Error Database503 0902Sistem sedang dalam maintenance System Under Maintenance503 0903Awal hari error Begin of Day Error500 0999Error sistem System error

POST https://sandbox.bca.co.id/fire/transactions

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST /fire/transactions HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]

{
    "Authentication":{
        "CorporateID":"ABCDE",
        "AccessCode":"iAepe8QGv39mTbuFrXN1",
        "BranchCode":"ABCDE01",
        "UserID":"ABCDEO003",
        "LocalID":"ABCDE01"
    },
    "TransactionDetails":{
        "InquiryBy":"N",
        "InquiryValue":"0247986325"
    }
}
```

POST https://sandbox.bca.co.id/fire/transactions

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs

```

##### 5\. Kiriman Uang ke Rekening Bank

Provides service transaction “Transaction to BCA’s Account” and also “Transfer to Other Bank”. Yout request must contain following information:

##### Payload

**Field****DataType****Mandatory****Description****Authentication**CorporateIDString(6)YAgent code used to connect to FIRe systemAccessCodeString(20)YPassword that will be read by system for each transaction requestBranchCodeString(8)YFI branch code of API clientUserIDString(12)YUser IDLocalIDString(15)YBranch code from FI sub branch of API client**SenderDetails**FirstNameString(35)YSender’s first nameLastNameString(35)NSender’s last nameDateOfBirthString(8)NSender’s date of birth. Format : ddmmyyyyAddress1String(35)YSender’s first addressAddress2String(35)NSender’s second addressCityString(18)YSender’s cityStateIDString(2)NState IDPostalCodeString(10)NSender’s postal codeCountryIDString(2)YSender’s country initial code. Example: Indonesia =IDMobileString(20)NSender’s mobile phone numberIdentificationTypeString(2)NSender’s identity type (PP = Passport, ID = identity, DR = Dr. License)IdentificationNumberString(24)NSender’s identity numberAccountNumberString(34)NSender’s account number**BeneficiaryDetails**NameString(35)YBeneficiary’s nameDateOfBirthString(8)NBeneficiary’s date of birth. Format : ddmmyyyyAddress1String(35)NBeneficiary’s first addressAddress2String(35)NBeneficiary’s second addressCityString(35)NBeneficiary’s cityStateIDString(2)NBeneficiary’s state IDPostalCodeString(10)NBeneficiary’s postal codeCountryIDString(2)YBeneficiary’s country initial code. Example : Indonesia =IDMobileString(20)NBeneficiary’s mobile phone numberIdentificationTypeString(2)NBeneficiary’s identity type. (PP = Passport, ID = identity, DR = Dr. License)IdentificationNumberString(24)NBeneficiary’s identity numberNationalityIDString(2)NCountry initial code of beneficiary’s nationality. Example = Indonesia =IDOccupationString(30)NBeneficiary occupationBankCodeTypeString(3)YBeneficiary’s SWIFT code type from PY04 table. (BIC, NID, CHP(CHIPS), ABA, NAM(NAME))BankCodeValueString(23)YValue of SWIFT CodeBankCountryIDString(2)YBeneficiary bank’s country codeBankAddressString(35)NBeneficiary bank’s addressBankCityString(35)NBeneficiary bank’s cityAccountNumberString(34)YBeneficiary bank’s account number**TransactionDetails**CurrencyIDString(3)YTransaction currency between FI and FIRe based on contract (setting FI Country in TPS)AmountString(20)YTransaction nominal. Format: Number, 17.2PurposeCodeString(3)YTransaction purpuse:

011 EXPORT

012 IMPORT

030 BUSINESS/PRIVATE TRIP

040 EDUCATION

150 WORKER’S REMMITANCE

161 TAX AND FINE

162 GRANT

999 OTHERS

Description1String(35)NFirst descriptionDescription2String(35)NSecond descriptionDetailOfChargesString(3)YCharges(SHA / OUR)SourceOfFundString(100)NSource of transaction fundFormNumberString(16)YFI Ref

Result of the request will contains following information:

##### Response

**Field****DataType****Mandatory****Description****BeneficiaryDetails**NameString(35)NBeneficiary’s nameAccountNumberString(34)NBeneficiary bank’s account numberServerBeneAccountNameString(35)NName of beneficiary according to application (bank / switcher)**TransactionDetails**CurrencyIDString(3)NTransaction currency between FI and FIRe based on contract (setting FI Country in TPS)AmountString(20)NTransaction nominal. Format: Number 17.2Description1String(35)NFirst DescriptionDescription2String(35)NSecond DescriptionFormNumberString(16)NFI RefReferenceNumberString(25)NFI RefReleaseDateTimeString(24)NTransaction release time. Format : YYYY-MM-DDThh:mm:ssTZDStatusTransactionString(4)YResponseCode and Sub-ResponseCode of transactionStatusMessageString(100)YDescription of StatusTransaction

##### Error

HTTP CodeError CodeError Message (Indonesian)Error Message (English)400 ESB-14-001 HMAC tidak cocok HMAC mismatch400 ESB-14-002 Permintaan tidak valid Invalid request400 ESB-14-003 Timestamp tidak valid Invalid timestamp400 ESB-14-004 Parameter input tidak valid Invalid input parameter503 ESB-14-005 Sistem sedang dalam maintenance System under maintenance504 ESB-14-006 Timeout, silakan periksa mutasi rekening Timeout, please check your account statement504 ESB-14-007 Timeout Timeout401 ESB-14-008 client\_id/client\_secret/grant\_type tidak valid Invalid client\_id/client\_secret/grant\_type401 ESB-14-009 Tidak berhak Unauthorized429 ESB-14-010 Jumlah permintaan melebihi limit Limit request exceeded404 ESB-14-011 Service tidak ada Service doesn't exist401 ESB-14-012 Tidak berhak mengakses service ini Not allowed to access this service400 0101Field \[fieldName\] is required Field \[fieldName\] is required400 0102Field \[fieldName\] is larger than expected Field \[fieldName\] is larger than expected400 0103Data type is invalid at field \[fieldName\] Data type is invalid at field \[fieldName\]409 0201Ref pengirim sama Same Sender Ref409 0202PIN sudah pernah dipakai PIN already in used400 0203Format mata uang tidak valid Invalid Currency format400 0204BIC Code tidak valid Invalid BIC Code400 0205ABA Code tidak valid Invalid ABA Code400 0206CHIPS Code tidak valid Invalid CHIPS Code400 0207NID Code tidak valid Invalid NID Code400 0208Country Code untuk NID Code tidak valid Invalid Country Code for NID Code401 0302Access Code salah Wrong Access Code400 0303Branch Code tidak ditemukan Unknown Branch Code400 0304User ID tidak ditemukan Unknown User ID401 0306Tidak berhak Not Authorized Privilege400 0401PRODUK/MATA UANG TIDAK DITEMUKAN PRODUCT/CURRENCY NOT FOUND400 0402LIMIT PRODUK TIDAK CUKUP PRODUCT LIMIT NOT ENOUGH400 0403LIMIT RELEASER TIDAK CUKUP LIMIT RELEASER NOT ENOUGH400 0404ERROR PDN LIMIT ERROR PDN LIMIT400 0405GAGAL VALIDASI REKENING PENERIMA FAILED TO VALIDATE BENEFICIARY ACCOUNT400 0406REKENING PENERIMA TIDAK VALID INVALID BENEFICIARY ACCOUNT400 0407SALDO TIDAK CUKUP INSUFFICIENT FUND404 0408TRANSAKSI TIDAK DITEMUKAN TRANSACTION NOT FOUND400 0409REKENING PENERIMA TUTUP CLOSED BENEFICIARY ACCOUNT500 0410TRANSAKSI TIDAK DAPAT DIPROSES TRANSACTION CAN NOT BE PROCESSED400 0412BUKAN MEMBER SWITCHING NON SWITCHING MEMBER400 0413REKENING TIDAK VALID INVALID ACCOUNT NUMBER400 0414DITOLAK KARENA SALDO TIDAK CUKUP REJECT DUE INSUFFICIENT FUND503 0415DITOLAK KARENA CUT OFF SISTEM BCA REJECT DUE TO CUT OFF SYSTEM BCA500 0901Error Database Error Database503 0902Sistem sedang dalam maintenance System Under Maintenance503 0903Awal hari error Begin of Day Error500 0999Error sistem System error

POST https://sandbox.bca.co.id/fire/transactions/to-account

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST /fire/transactions/to-account HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]

{
    "Authentication": {
        "CorporateID": "CORP1",
        "AccessCode": "Ya9wn025GgaKkcbDsTvR",
        "BranchCode": "BRANCH01",
        "UserID": "USERID0001",
        "LocalID": "LOCAL001"
    },
    "SenderDetails": {
        "FirstName": "First Name",
        "LastName": "",
        "DateOfBirth": "",
        "Address1": "Taman Cibaduyut RT1",
        "Address2": "",
        "City": "Kab  Bandung",
        "StateID": "",
        "PostalCode": "",
        "CountryID": "ID",
        "Mobile": "",
        "IdentificationType": "",
        "IdentificationNumber": "",
        "AccountNumber": ""
    },
    "BeneficiaryDetails": {
        "Name": "RAFI ZULFIKAR MUSTAQIM",
        "DateOfBirth": "",
        "Address1": "",
        "Address2": "",
        "City": "",
        "StateID": "",
        "PostalCode": "",
        "CountryID": "ID",
        "Mobile": "",
        "IdentificationType": "",
        "IdentificationNumber": "",
        "NationalityID": "",
        "Occupation": "",
        "BankCodeType": "BIC",
        "BankCodeValue": "CENAIDJAXXX",
        "BankCountryID": "ID",
        "BankAddress": "",
        "BankCity": "",
        "AccountNumber": "0106666011"
    },
    "TransactionDetails": {
        "CurrencyID": "IDR",
        "Amount": "100000",
        "PurposeCode": "550",
        "Description1": "",
        "Description2": "",
        "DetailOfCharges": "OUR",
        "SourceOfFund": "",
        "FormNumber": "56559859607940"
    }
}
```

POST https://sandbox.bca.co.id/fire/transactions/to-account

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "BeneficiaryDetails":
    {
        "Name":"monica gupt",
        "AccountNumber":"0106666011",
        "ServerBeneAccountName":"monica gupt"
    },
    "TransactionDetails":
    {
        "CurrencyID":"IDR",
        "Amount":"10000000",
        "Description1":"",
        "Description2":"",
        "FormNumber":"7632605701245868",
        "ReferenceNumber":"IDPYID01000INA17030000057",
        "ReleaseDateTime":"2017-10-30T10:37:47.380+07:00"
    },
    "StatusTransaction":"0000",
    "StatusMessage":"Success"
}
```

##### 6\. Pembatalan Kiriman Uang Tunai

Provides service for Cancellation “Cash Transfer” to Non account holder. Yout request must contain following information:

##### Payload

**Field****DataType****Mandatory****Description****Authentication**CorporateIDString(6)YAgent code used to connect to FIRe systemAccessCodeString(20)YPassword that will be read by system for eachBranchCodeString(8)YFI branch code of API clientUserIDString(12)YUser IDLocalIDString(15)YBranch code from FI sub branch of API client**TransactionDetails**FormNumberString(16)YFI RefAmountString(20)YTransaction nominal.

Format: Number, 17.2CurrencyIDString(3)NCurrency code

Result of the request will contains following information:

##### Response

**Field****DataType****Mandatory****Description**StatusTransactionString(4)YRespponseCode and Sub-ResponseCode of transactionStatusMessageString(100)YDescription of StatusTransaction**TransactionDetails**FormNumberString(16)NFI RefReleaseDateTimeString(24)Transaction release time.

Format : YYYY-MM-DDThh:mm:ssTZD

##### Error

HTTP CodeError CodeError Message (Indonesian)Error Message (English)400 ESB-14-001 HMAC tidak cocok HMAC mismatch400 ESB-14-002 Permintaan tidak valid Invalid request400 ESB-14-003 Timestamp tidak valid Invalid timestamp400 ESB-14-004 Parameter input tidak valid Invalid input parameter503 ESB-14-005 Sistem sedang dalam maintenance System under maintenance504 ESB-14-006 Timeout, silakan periksa mutasi rekening Timeout, please check your account statement504 ESB-14-007 Timeout Timeout401 ESB-14-008 client\_id/client\_secret/grant\_type tidak valid Invalid client\_id/client\_secret/grant\_type401 ESB-14-009 Tidak berhak Unauthorized429 ESB-14-010 Jumlah permintaan melebihi limit Limit request exceeded404 ESB-14-011 Service tidak ada Service doesn't exist401 ESB-14-012 Tidak berhak mengakses service ini Not allowed to access this service400 0101Field \[fieldName\] is required Field \[fieldName\] is required400 0102Field \[fieldName\] is larger than expected Field \[fieldName\] is larger than expected400 0103Data type is invalid at field \[fieldName\] Data type is invalid at field \[fieldName\]409 0201Ref pengirim sama Same Sender Ref409 0202PIN sudah pernah dipakai PIN already in used400 0203Format mata uang tidak valid Invalid Currency format400 0204BIC Code tidak valid Invalid BIC Code400 0205ABA Code tidak valid Invalid ABA Code400 0206CHIPS Code tidak valid Invalid CHIPS Code400 0207NID Code tidak valid Invalid NID Code400 0208Country Code untuk NID Code tidak valid Invalid Country Code for NID Code401 0302Access Code salah Wrong Access Code400 0303Branch Code tidak ditemukan Unknown Branch Code400 0304User ID tidak ditemukan Unknown User ID401 0306Tidak berhak Not Authorized Privilege400 0401PRODUK/MATA UANG TIDAK DITEMUKAN PRODUCT/CURRENCY NOT FOUND400 0402LIMIT PRODUK TIDAK CUKUP PRODUCT LIMIT NOT ENOUGH400 0403LIMIT RELEASER TIDAK CUKUP LIMIT RELEASER NOT ENOUGH400 0404ERROR PDN LIMIT ERROR PDN LIMIT400 0405GAGAL VALIDASI REKENING PENERIMA FAILED TO VALIDATE BENEFICIARY ACCOUNT400 0406REKENING PENERIMA TIDAK VALID INVALID BENEFICIARY ACCOUNT400 0407SALDO TIDAK CUKUP INSUFFICIENT FUND404 0408TRANSAKSI TIDAK DITEMUKAN TRANSACTION NOT FOUND400 0409REKENING PENERIMA TUTUP CLOSED BENEFICIARY ACCOUNT500 0410TRANSAKSI TIDAK DAPAT DIPROSES TRANSACTION CAN NOT BE PROCESSED400 0412BUKAN MEMBER SWITCHING NON SWITCHING MEMBER400 0413REKENING TIDAK VALID INVALID ACCOUNT NUMBER400 0414DITOLAK KARENA SALDO TIDAK CUKUP REJECT DUE INSUFFICIENT FUND503 0415DITOLAK KARENA CUT OFF SISTEM BCA REJECT DUE TO CUT OFF SYSTEM BCA500 0901Error Database Error Database503 0902Sistem sedang dalam maintenance System Under Maintenance503 0903Awal hari error Begin of Day Error500 0999Error sistem System error

POST https://sandbox.bca.co.id/fire/transactions/cash-transfer/cancel

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST /fire/transactions/cash-transfer/cancel HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]

{
    "Authentication":{
        "CorporateID":"BCAAPI",
        "AccessCode":"k6re8tG35hgNODSg3Y",
        "BranchCode":"BCA01",
        "UserID":"BCAUSERID001",
        "LocalID":"40115"
    },
    "TransactionDetails":{
        "FormNumber":"CT293 IDR2D",
        "Amount":"9350000.00",
        "CurrencyID":"IDR"
    }
}
```

POST https://sandbox.bca.co.id/fire/transactions/cash-transfer/cancel

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "TransactionDetails":
    {
        "FormNumber":"CT293_IDR2D",
        "ReleaseDateTime":"2016-04-20T04:00:23+07:00"
    },
    "StatusTransaction":"0000",
    "StatusMessage":"Success"
}
```

##### 7\. Transaksi Kiriman Uang Tunai

Provides service for transaction “Cash Transfer” to Non account holder.

Your request must contain following information:

##### Payload

**Field****DataType****Mandatory****Description****Authentication**CorporateIDString(6)YAgent code used to connect to FIRe systemAccessCodeString(20)YPassword that will be read by system for each transaction requestBranchCodeString(8)YFI branch code of API clientUserIDString(12)YUser IDLocalIDString(15)YBranch code from FI sub branch of API client**SenderDetails**FirstNameString(35)YSender’s first nameLastNameString(35)NSender’s last nameDateOfBirthString(8)NSender’s date of birth. Format : ddmmyyyyAddress1String(35)YSender’s first addressAddress2String(35)NSender’s second addressCityString(18)YSender’s cityStateIDString(2)NSender’s state IDPostalCodeString(10)NSender’s postal codeCountryIDString(2)YSender’s country initial code. Example : Indonesia = IDMobileString(20)NSender’s mobile phone numberIdentificationTypeString(2)NSender’s identity type. (PP = Passport, ID = Identity, DR = Dr. License)IdentificationNumberString(24)NSender’s identity number**BeneficiaryDetails**NameString(35)YBeneficiary’s nameDateOfBirthString(8)NBeneficiary’s date of birth. Format : ddmmyyyyAddress1String(35)NBeneficiary’s first addressAddress2String(35)NBeneficiary’s second addressCityString(35)NBeneficiary’s cityStateIDString(2)NBeneficiary’s state IDPostalCodeString(10)NBeneficiary’s postal codeCountryIDString(2)YBeneficiary’s country initial code. Example : Indonesia = IDMobileString(20)NBeneficiary’s mobile phone numberIdentificationTypeString(2)NBeneficiary’s identity type. (PP = Passport, ID = Identity, DR = Dr. License)IdentificationNumberString(24)NBeneficiary’s identity numberNationalityIDString(2)NCountry initial code of beneficiary’s nationality. Example : Indonesia = IDOccupationString(30)NBeneficiary occupation**TransactionDetails**PINString(35)NTransaction PIN number. If not available, BCA will provide itSecretQuestionString(100)NQuestionsSecretAnswerString(100)NAnswersCurrencyIDString(3)YTransaction currency between FI and FIRe based on contract (setting FI Country in TPS)AmountString(20)YTransaction nominal. Format: number(17,2)PurposeCodeString(3)YTransaction purpose.

011 EXPORT

012 IMPORT

030 BUSINESS/PRIVATE TRIP

040 EDUCATION

150 WORKER’S REMMITANCE 161 TAX AND FINE

162 GRANT

999 OTHERSDescription1String(35)NFirst descriptionDescription2String(35)NSecond descriptionDetailOfChargesString(3)YCharges (SHA / OUR)SourceOfFundString(100)YSource of transaction fundFormNumberString(16)YFI Ref

##### Response

**Field****DataType****Mandatory****Description**StatusTransactionString(4)YResponseCode and Sub- ResponseCode of transactionStatusMessageString(100)YDescription of StatusTransaction**BeneficiaryDetails**NameString(35)NBeneficiary’s name**TransactionDetails**PINString(35)NTransaction PIN number. If not available, BCA will provide itCurrencyIDString(3)NTransaction currency between FI and FIRe based on contract(setting FI Country in TPS)AmountString(20)NTransaction nominalDescription1String(35)NFirst descriptionDescription2String(35)NSecond descriptionFormNumberString(16)NFI RefReferenceNumberString(25)NFI RefReleaseDateTimeString(24)NTransaction release time.

Format : YYYY-MM-DDThh:mm:ssTZD

##### Error

HTTP CodeError CodeError Message (Indonesian)Error Message (English)400 ESB-14-001 HMAC tidak cocok HMAC mismatch400 ESB-14-002 Permintaan tidak valid Invalid request400 ESB-14-003 Timestamp tidak valid Invalid timestamp400 ESB-14-004 Parameter input tidak valid Invalid input parameter503 ESB-14-005 Sistem sedang dalam maintenance System under maintenance504 ESB-14-006 Timeout, silakan periksa mutasi rekening Timeout, please check your account statement504 ESB-14-007 Timeout Timeout401 ESB-14-008 client\_id/client\_secret/grant\_type tidak valid Invalid client\_id/client\_secret/grant\_type401 ESB-14-009 Tidak berhak Unauthorized429 ESB-14-010 Jumlah permintaan melebihi limit Limit request exceeded404 ESB-14-011 Service tidak ada Service doesn't exist401 ESB-14-012 Tidak berhak mengakses service ini Not allowed to access this service400 0101Field \[fieldName\] is required Field \[fieldName\] is required400 0102Field \[fieldName\] is larger than expected Field \[fieldName\] is larger than expected400 0103Data type is invalid at field \[fieldName\] Data type is invalid at field \[fieldName\]409 0201Ref pengirim sama Same Sender Ref409 0202PIN sudah pernah dipakai PIN already in used400 0203Format mata uang tidak valid Invalid Currency format400 0204BIC Code tidak valid Invalid BIC Code400 0205ABA Code tidak valid Invalid ABA Code400 0206CHIPS Code tidak valid Invalid CHIPS Code400 0207NID Code tidak valid Invalid NID Code400 0208Country Code untuk NID Code tidak valid Invalid Country Code for NID Code401 0302Access Code salah Wrong Access Code400 0303Branch Code tidak ditemukan Unknown Branch Code400 0304User ID tidak ditemukan Unknown User ID401 0306Tidak berhak Not Authorized Privilege400 0401PRODUK/MATA UANG TIDAK DITEMUKAN PRODUCT/CURRENCY NOT FOUND400 0402LIMIT PRODUK TIDAK CUKUP PRODUCT LIMIT NOT ENOUGH400 0403LIMIT RELEASER TIDAK CUKUP LIMIT RELEASER NOT ENOUGH400 0404ERROR PDN LIMIT ERROR PDN LIMIT400 0405GAGAL VALIDASI REKENING PENERIMA FAILED TO VALIDATE BENEFICIARY ACCOUNT400 0406REKENING PENERIMA TIDAK VALID INVALID BENEFICIARY ACCOUNT400 0407SALDO TIDAK CUKUP INSUFFICIENT FUND404 0408TRANSAKSI TIDAK DITEMUKAN TRANSACTION NOT FOUND400 0409REKENING PENERIMA TUTUP CLOSED BENEFICIARY ACCOUNT500 0410TRANSAKSI TIDAK DAPAT DIPROSES TRANSACTION CAN NOT BE PROCESSED400 0412BUKAN MEMBER SWITCHING NON SWITCHING MEMBER400 0413REKENING TIDAK VALID INVALID ACCOUNT NUMBER400 0414DITOLAK KARENA SALDO TIDAK CUKUP REJECT DUE INSUFFICIENT FUND503 0415DITOLAK KARENA CUT OFF SISTEM BCA REJECT DUE TO CUT OFF SYSTEM BCA500 0901Error Database Error Database503 0902Sistem sedang dalam maintenance System Under Maintenance503 0903Awal hari error Begin of Day Error500 0999Error sistem System error

POST https://sandbox.bca.co.id/fire/transactions/cash-transfer

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
POST /fire/transactions/cash-transfer HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]

{
    "Authentication": {
        "CorporateID": "PYTHSG",
        "AccessCode": "wArjfChZjj9dC3eJqz0X",
        "BranchCode": "PYTHSG01",
        "UserID": "PYTHSGAPI1",
        "LocalID": "PYTHSGAPI1"
    },
    "SenderDetails": {
        "FirstName": "Sutirah Binti",
        "LastName": "Hadi Suwito",
        "DateOfBirth": "",
        "Address1": "20 Pine Grove  06-02, Singapore",
        "Address2": "Singapore",
        "City": "SINGAPORE",
        "StateID": "",
        "PostalCode": "",
        "CountryID": "SG",
        "Mobile": "",
        "IdentificationType": "ID",
        "IdentificationNumber": "7729902",
        "AccountNumber": ""
    },
    "BeneficiaryDetails": {
        "Name": "HADI SUWITO",
        "DateOfBirth": "",
        "Address1": "",
        "Address2": "",
        "City": "",
        "StateID": "",
        "PostalCode": "",
        "CountryID": "ID",
        "Mobile": "",
        "IdentificationType": "",
        "IdentificationNumber": "",
        "NationalityID": "",
        "Occupation": "",
        "BankCodeType": "",
        "BankCodeValue": "",
        "BankCountryID": "",
        "BankAddress": "",
        "BankCity": "",
        "AccountNumber": ""
    },
    "TransactionDetails": {
        "PIN": "P2HBCA5581412",
        "CurrencyID": "IDR",
        "Amount": "5870890",
        "PurposeCode": "011",
        "Description1": "",
        "Description2": "",
        "DetailOfCharges": "OUR",
        "SourceOfFund": "SOF",
        "FormNumber": "5581412"
    }
}
```

POST https://sandbox.bca.co.id/fire/transactions/cash-transfer

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
{
    "BeneficiaryDetails":
    {
        "Name":"TEST"
    },
    "TransactionDetails":
    {
        "PIN":"477634423",
        "CurrencyID":"IDR",
        "Amount":"150000.00",
        "Description1":"",
        "Description2":"",
        "FormNumber":"477634423",
        "ReferenceNumber":"CITIID01000NON16040000099",
        "ReleaseDateTime":""
    },
    "StatusTransaction":"0003",
    "StatusMessage":"Ready to Encash"
}
```

#### Top Up Flazz

##### 1\. Top Up Flazz Mobile Operating System

Banking service to make Flazz top up transaction by using Merchant's device/reader/application.

To be able to use this feature, the following 5 API services must be accessed by the client:

##### a. Get Session Key

API service to request session key. Session key is formed by randomly generated key to ensure session security during communication between Merchant and BCA Host.

Your Request must contain following information:

##### Request Headers and URI Params

Field**Params Type**DataType**Mandatory**Descriptionchannel-id HeaderStringY95381 (WSID FLAZZ)credential-id HeaderString(36)YCopartner-id

##### Payload (Form Data)

FieldData Type**Mandatory**Descriptionreference\_idString(12)Y Reference ID is datetime transaction in milisecond (unix time), requires only the first 12 digits and the same reference\_id must be carried in 1 service API set for Top Up Flazz. Ex : 969677130000 Unique per transactionrequest\_dataString(198)YRequest data stream

##### Response

FieldDataTypeMandatoryDescriptionresponse\_dataString(198) NResponse data streamtransaction\_idString (16)Ntransaction\_id

##### b. Send topup request from merchant

API service to send Flazz balance top up requests.

Your Request must contain following information:

##### Request Headers and URI Params

Field**Params Type**DataTypeMandatoryDescriptionchannel-idHeaderStringY95381 (WSID FLAZZ)credential-idHeaderString(36)YCopartner-id

##### Payload (Form Data)

FieldData Type**Mandatory**Descriptionreference\_idString(12)YReference ID is datetime transaction in milisecond (unix time), requires only the first 12 digits and the same reference\_id must be carried in 1 service API set for Top Up Flazz. Ex : 969677130000 Unique per transactionrequest\_dataString(448)YRequest data stream

##### Response

FieldDataType**Mandatory**Descriptionresponse\_dataString(350)NResponse data stream

ref\_noString(6)Nref notransaction\_idString (16)Ntransaction\_id

##### c. Reversal

API service to send reversal request of Flazz top up

Your Request must contain following information:

##### Request Headers and URI Params

Field**Params Type**DataTypeMandatoryDescriptionchannel-idHeaderStringY95381 (WSID FLAZZ)credential-idHeaderString(36)YCopartner-id

##### Payload (Form Data)

FieldData Type**Mandatory**Descriptionreference\_idString(12)YReference ID is datetime transaction in milisecond (unix time), requires only the first 12 digits and the same reference\_id must be carried in 1 service API set for Top Up Flazz. Ex : 969677130000 Unique per transactionrequest\_dataString(544)YRequest data stream

##### Response

FieldDataType**Mandatory**Descriptionresponse\_dataStringNResponse data stream

transaction\_idString (16)Ntransaction\_id

##### d.ACK

API service to send ACK requests (Acknowledge) for successful / failed / lost-contact transactions of Flazz top up

Your Request must contain following information:

##### Request Headers and URI Params

Field**Params Type**DataTypeMandatoryDescriptionchannel-idHeaderStringY95381 (WSID FLAZZ)credential-idHeaderString(36)YCopartner-id

##### Payload (Form Data)

FieldData Type**Mandatory**Descriptionreference\_idString(12)YReference ID is datetime transaction in milisecond (unix time), requires only the first 12 digits and the same reference\_id must be carried in 1 service API set for Top Up Flazz. Ex : 969677130000 Unique per transactionrequest\_dataString(544)YRequest data stream

##### Response

FieldDataType**Mandatory**Descriptionresponse\_dataStringNResponse data stream

transaction\_idString (16)Ntransaction\_id

##### e.Card Verification

API service to send Card Verification requests

Your Request must contain following information:

##### Request Headers and URI Params

Field**Params Type**DataTypeMandatoryDescriptionchannel-idHeaderStringY95381 (WSID FLAZZ)credential-idHeaderString(36)YCopartner-id

##### Payload (Form Data)

FieldData Type**Mandatory**Descriptionreference\_idString(12)YReference ID is datetime transaction in milisecond (unix time), requires only the first 12 digits and the same reference\_id must be carried in 1 service API set for Top Up Flazz. Ex : 969677130000 Unique per transactionrequest\_dataString(16)YRequest data from online merchants to check Flazz card number
Response

FieldDataType**Mandatory**Descriptionresponse\_dataString(17)NFlazz card number


Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message (Indonesian)**Error Message (English)**Description400ESB-14-001HMAC tidak cocokHMAC mismatch400400ESB-14-002Permintaan tidak validInvalid request400400ESB-14-003Timestamp tidak validInvalid timestamp400400ESB-14-004parameter input tidak validInvalid input parameter400500ESB-14-005Sistem sedang dalam maintenanceSystem under maintenance500504ESB-14-007TimeoutTimeout504401ESB-14-008Client\_Id/Client\_Secret /Grant\_Type tidak validInvalid Client\_Id/Client\_Secret /Grant\_Type401401ESB-14-009Tidak berhakUnauthorized401404ESB-14-011Service tidak tersediaService doesn't exist404401ESB-14-012Tidak berhak mengakses service iniNot allowed to access this service401401ESB-14-014ChannelID/CredentialID tidak validInvalid ChannelID/CredentialID 401400ESB-14-015Content Type tidak valid Invalid Content Type400400ESB-14-016Format JSON tidak valid Invalid JSON format400400ESB-14-021API Key tidak validInvalid API Key400400ESB-14-022API Key/API Secret tidak cocokAPI Key/API Secret mismatch400504ESB-14-023Koneksi Timeout. Silakan cek transaksi kembali untuk memastikan keberhasilan transaksi Anda Connection timeout. Please reinquiry your transaction to ensure transaction status5044000002Transaksi GagalTransaction FailedCopartner is not registered4000012Transaksi GagalTransaction FailedCopartner is blocked4000022Transaksi GagalTransaction FailedCopartner max retry exceeded 4000032Transaksi GagalTransaction FailedCopartner session key not exist4001012Kartu ErrorCard ErrorCheck expired date failed4001022Transaksi gagal, nominal Top Up kurang dari minimumTransaction failed, Top Up amount less than minimumCheck min amount failed4001032Transaksi gagal, nominal Top Up melebihi saldo maksimumTransaction failed, Top Up amount exceeds maximum balanceCheck max balance failed4001052Transaksi gagalTransaction failedCheck mandatory field failed4001062Transaksi gagalTransaction failedParse request data failed4001102Transaksi gagalTransaction failedDuplicate transaction4001122Transaksi gagalTransaction failedCheck original data failed4001132Transaksi gagalTransaction failedCheck tid failed4001142Transaksi gagalTransaction failedCheck mid failed4001152Kartu errorCard errorCard blacklisted4001202Transaksi gagalTransaction failedTID and MID doesn’t exist4003013Transaksi gagalTransaction failedCheck topup limit timeout failed4003023Transaksi gagalTransaction failedRequest topup debit timeout failed4003045Transaksi gagalTransaction failedCheck atuf timeout failed4004662Transaksi tidak dapat diproses. Silahkan coba beberapa saat lagi.Transaction can not be processed. Please try again later.Error dari EAI apabila terjadi koneksi timeout saat memanggil ke service lain.4004045Maaf, transaksi anda tidak dapat diprosesSorry, your transaction can not be processedIdentifier Channel tidak diizinkan mengakses4004009Field (0) harus diisiMissing mandatory field (0)Field (0) harus diisi4004038Panjang field (0) tidak valid Field (0) length not validPanjang field (0) tidak valid4004039Isian input (0) tidak validInput for field (0) not validIsian input (0) tidak valid4004112Inputan String JSON tidak validInvalid JSON String inputFormat JSON input invalid 4004113Transaksi ditolakTransaction rejectedClientID tidak dapat mengakses ke service4004414Transaksi gagalTransaction failedGagal saat konversi(format data yang diterima tidak sesuai)4004415Transaksi gagalTransaction failedRecord yang ingin dihapus tidak ada.4004416Transaksi gagalTransaction failedFile Limit TopUp tidak ada4004417Transaksi gagalTransaction failedFile Limit TopUp tidak bisa dibuka 4004418Transaksi gagalTransaction failedFile Limit TopUp tidak bisa dibaca4004419Transaksi gagalTransaction failedRecord sudah ada ketika tambah record baru 4004420Transaksi gagalTransaction failedRecord tidak ada ketika Change/Update4004421Transaksi gagalTransaction failedGagal menulis record4004422Transaksi gagalTransaction failedGagal hapus record4004423Transaksi gagalTransaction failedFIID yang dikirimkan tidak sesuai/data FIID tidak ada4004424Transaksi gagalTransaction failedTransaksi code tidak sesuai(sesuai dengan data element 91) 4004425Transaksi gagalTransaction failedFIID tidak terdaftar4004426Transaksi gagalTransaction failedFile request tidak sesuai(sesuai dengan data element 101)4004428Transaksi gagalTransaction failedGagal ketika mencoba mencari nomor kartu4004429Transaksi gagal, jumlah top up melebihi limit bulananTransaction failed, Top Up amount exceeds monthly limitLimit kartu Flazz habis.4004430Transaksi gagalTransaction failedUnik id tidak ditemukan4004431Transaksi gagalTransaction failedKoneksi saat dari atau ke b24 terputus4004432Transaksi gagalTransaction failedKoneksi saat dari atau ke b24d itolak4004414Transaksi gagalTransaction failedGagal saat konversi(format data yang diterima tidak sesuai)4004415Transaksi gagalTransaction failedRecord yang ingin dihapus tidak ada.4004416Transaksi gagalTransaction failedFile Limit TopUp tidak ada4004417Transaksi gagalTransaction failedFile Limit TopUp tidak bisa dibuka 4004418Transaksi gagalTransaction failedFile Limit TopUp tidak bisa dibaca4004419Transaksi gagalTransaction failedRecord sudah ada ketika tambah record baru 4004420Transaksi gagalTransaction failedRecord tidak ada ketika Change/Update4004421Transaksi gagalTransaction failedGagal menulis record 4004422Transaksi gagalTransaction failedGagal hapus record 4004423Transaksi gagalTransaction failedFIID yang dikirimkan tidak sesuai/data FIID tidak ada4004424Transaksi gagalTransaction failedTransaksi code tidak sesuai(sesuai dengan data element 91) 4004425Transaksi gagalTransaction failedFIID tidak terdaftar4004426Transaksi gagalTransaction failedFile request tidak sesuai(sesuai dengan data element 101)4004427Transaksi gagalTransaction failedCategory selain C,F,T,R. 4004428Transaksi gagalTransaction failedGagal ketika mencoba mencari nomor kartu4004429Transaksi gagal, kartu sumber dana lebih dari limit Transaction failed, card for source of founds is over limitLimit habis kartu akses habis4004430Transaksi gagalTransaction failedUnik id tidak ditemukan4004431Transaksi gagalTransaction failedKoneksi saat dari atau ke b24 terputus4004432Transaksi gagalTransaction failedKoneksi saat dari atau ke b24 ditolak4000000Transaksi BerhasilTransaction SuccessTransaksi berhasil4005202Transaksi 0 rupiahTransaction amount 0Transaksi 0 Rupiah4005204Maaf saldo tidak mencukupi Sorry your balance is not enough Maaf saldo tidak mencukupi4005206Rekening ditolakAccount RejectedRekening ditolak4005700Maaf transaksi harus IDRSorry, currency must IDRMaaf transaksi harus IDR4005701Rekening ditolakAccount RejectedRekening ditolak 4005702Maaf rekening sudah tidak aktifSorry, your account is dormantMaaf rekening sudah tidak aktif4005711Maaf rekening anda masuk daftar hitamSorry, your account is blacklistRekening anda diblokir/masuk daftar hitam 4005999Sistem sedang tidak tersediaSystem unavailableSistem sedang tidak tersedia

POST /flazz/

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
REQUEST FOR GET SESSION KEY

POST https://sandbox.bca.co.id/flazz/v2/getKey HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
ChannelID: [channelID]
CredentialID: [credentialID]
{
    "reference_id" : "123456789012",
    "request_data" : "01010000202001171513224553594230303031454D424341303031885000015999824E180B6E8227061B4E8A275E3D373702021B4B5E821B063D4E0B370627024E843D3D274B181B8A1B375E4E8A4E020B203"
}

REQUEST FOR TOP UP REQUEST

POST https://sandbox.bca.co.id/flazz/v2/topup HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
ChannelID: [channelID]
CredentialID: [credentialID]
{
    "reference_id" : "123456789012",
    "request_data" : "01010024014500280000003800311900000000102020020317152011013049124553594230303031454D4243413030318850000159996E8282370684824E5E02848A8482188A064B06064E0627278A184B4E4E828A3D848A4B82028402064B8206023D1C0B1C1B0B826E1C063D6E185E6E0B4E821B844B37273718064E82024B4E1B0282185E4B37273718064E82024B4E1B0282185E4B37273718064E82024B4E1B0282185E4E0B3D181C4E1B27026E3D5E6E186E1B8A1C3D273782028A068A274B3D0B826E1AE5BCBB7C0226A4DB8C"
}

REQUEST FOR REVERSAL

POST https://sandbox.bca.co.id/flazz/v2/reversal HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
ChannelID: [channelID]
CredentialID: [credentialID]
{
    "reference_id" : "123456789012",
    "request_data" : ""01010024014500010036938400291613000000102020012315164520013049124553594230303031454D424341303031885000015999275E6E824B1B4E021B1B1C4E0606023D4E844E5E8A021B6E8A4B4B1C6E4E6E8A4E5E8A82826E374B1B1B0218188A068482061C278A5E0B180B8A6E8A3D0B8A823784186E0B4E8237024E020B6E02825E1B8A020B270B82270227378202183D846E84824E021B4B844B82273D02183D5E5E8A061C6E06068A8A061B024B1B3D4B068A1C02020B6E824E8A371B1C064E8A3D4B84270682278A5E84374E37061C843D4B84270682278A5E84374E37061C843D4B84270682278A5E84374E37061C84C0396097845AD44C26CC"
}

REQUEST FOR ACK

POST https://sandbox.bca.co.id/flazz/v2/ack HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
ChannelID: [channelID]
CredentialID: [credentialID]
{
    "reference_id" : "123456789012",
    "request_data" : ""01010024014500010036938400291613000000102020012315164520013049124553594230303031454D424341303031885000015999275E6E824B1B4E021B1B1C4E0606023D4E844E5E8A021B6E8A4B4B1C6E4E6E8A4E5E8A82826E374B1B1B0218188A068482061C278A5E0B180B8A6E8A3D0B8A823784186E0B4E8237024E020B6E02825E1B8A020B270B82270227378202183D846E84824E021B4B844B82273D02183D5E5E8A061C6E06068A8A061B024B1B3D4B068A1C02020B6E824E8A371B1C064E8A3D4B84270682278A5E84374E37061C843D4B84270682278A5E84374E37061C843D4B84270682278A5E84374E37061C84C0396097845AD44C26CC"
}

REQUEST FOR CARD VERIFICATION

POST https://api.klikbca.com/flazz/v2/cardver
-H channel-id 95381
-H credential-id 00ESBY0001
-H Authorization Bearer LYQG93kCXeQCwUWbhYcv5fnrRHum1esgN7nlIl3xmOWhfPIuicXzb9
-H X-BCA-Key fd2bd8c5-ff4d-41b6-8905-f0264f5c6b4f
-H X-BCA-Signature 652ce6723d7a0272c40d5812b3b16ecdaa689b74a04d6d859ed5ffcd7b233e2a
-H X-BCA-Timestamp 2020-02-18T09:27:31.198+07:00
-d {
 "reference_id": "123456789012",
 "request_data": " 0145007206508575"
 }
```

POST /flazz/

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
RESPONSE FOR GET SESSION KEY

{
    "response_data" : ""01010000202001171513224553594230303031454D4243413030318850000159996DD3D35D3422B7B022CF34B7B0CD6D34345DB07A7ACFA45DCFD3B7A4D3A46DA4CFB0D3D35334B0347D6E537A347A5DCDABB",
"transaction_id": "1234567800000000"
}

RESPONSE FOR REQUEST TOP UP

{
    "response_data" : "01240145002800000038000000100000000000160230202002191019275D5D6E5DB05D6D7A7DCF7A22227A6EB07D5D3422CD68CF6853A4A4CDB77AD35DD334347AB7A45DB7D33434B06D68D3D334B7CF7D92D35D34A46E7AD3B7CF34685D34CF34A453D36EB75DD368B76EA4925D34CF34A453D36EB75DD368B76EA4925D34CF34A453D36EB75DD368B76EA4925D34CF34A453D36EB75DD368B76EA4925D34CF34A453D36EB75DD368B76EA492F8B2",
"ref_no": "777777",
 "transaction_id": "0000112233445566"
}

RESPONSE FOR REVERSAL

{
"response_data": null,
 "transaction_id": "1234567800000000"
}

RESPONSE FOR CARD VERIFICATION

{
    "response_data" : "01450072065085750"
}

RESPONSE FOR ACK

{
    "response_data" : "null"
    "transaction_id": "1234567800000000"
}

Error Schema
{
"error_code" : "Error code",
"error_message" : {
"indonesian" : "Error Message (Indonesian)",
"english" : " Error Message (English)"
}
}
```

##### 2\. Top Up Flazz Non Mobile Operating System

Banking service to make Flazz top up transaction by using Merchant's device/reader.

To be able to use this feature, the following 4 API services must be accessed by the client:

##### a. Get Session Key

API service to request session key. Session key is formed by randomly generated key to ensure session security during communication between Merchant and BCA Host.

Your Request must contain following information:

##### Request Headers and URI Params

Field**Params Type**DataTypeMandatoryDescriptionchannel-id HeaderStringY95381 (WSID FLAZZ)credential-id HeaderString(36)YCopartner-id

##### Payload (Form Data)

FieldData Type**Mandatory**Descriptionreference\_idString(12)Y Reference ID is datetime transaction in milisecond (unix time), requires only the first 12 digits and the same reference\_id must be carried in 1 service API set for Top Up Flazz. Ex : 969677130000 Unique per transactionrequest\_dataString(166)YRequest data stream

##### Response

FieldDataType**Mandatory**Descriptionresponse\_dataString(166) NResponse data stream

##### b. Send topup request from merchant

API service to send Flazz balance top up requests.

Your Request must contain following information:

##### Request Headers and URI Params

Field**Params Type**DataTypeMandatoryDescriptionchannel-idHeaderStringY95381 (WSID FLAZZ)credential-idHeaderString(36)YCopartner-id

##### Payload (Form Data)

FieldData Type**Mandatory**Descriptionreference\_idString(12)YReference ID is datetime transaction in milisecond (unix time), requires only the first 12 digits and the same reference\_id must be carried in 1 service API set for Top Up Flazz. Ex : 969677130000 Unique per transactionrequest\_dataString(416)YRequest data stream

##### Response

FieldDataType**Mandatory**Descriptionresponse\_dataString(350)NResponse data stream

##### c. Reversal

API service to send reversal request of Flazz top up

Your Request must contain following information:

##### Request Headers and URI Params

Field**Params Type**DataTypeMandatoryDescriptionchannel-idHeaderStringY95381 (WSID FLAZZ)credential-idHeaderString(36)YCopartner-id

##### Payload (Form Data)

FieldData Type**Mandatory**Descriptionreference\_idString(12)YReference ID is datetime transaction in milisecond (unix time), requires only the first 12 digits and the same reference\_id must be carried in 1 service API set for Top Up Flazz. Ex : 969677130000 Unique per transactionrequest\_dataString(512)YRequest data stream

##### Response

FieldDataType**Mandatory**Descriptionresponse\_dataStringNResponse data stream

##### d.ACK

API service to send ACK requests (Acknowledge) for successful / failed / lost-contact transactions of Flazz top up

Your Request must contain following information:

##### Request Headers and URI Params

Field**Params Type**DataTypeMandatoryDescriptionchannel-idHeaderStringY95381 (WSID FLAZZ)credential-idHeaderString(36)YCopartner-id

##### Payload (Form Data)

FieldData Type**Mandatory**Descriptionreference\_idString(12)YReference ID is datetime transaction in milisecond (unix time), requires only the first 12 digits and the same reference\_id must be carried in 1 service API set for Top Up Flazz. Ex : 969677130000 Unique per transactionrequest\_dataString(512)YRequest data stream

##### Response

FieldDataType**Mandatory**Descriptionresponse\_dataStringNResponse data stream

Here is the list of error codes that can be returned.

HTTP CodeError CodeError Message (Indonesian)**Error Message (English)**Description400ESB-14-001HMAC tidak cocokHMAC mismatch400400ESB-14-002Permintaan tidak validInvalid request400400ESB-14-003Timestamp tidak validInvalid timestamp400400ESB-14-004parameter input tidak validInvalid input parameter400500ESB-14-005Sistem sedang dalam maintenanceSystem under maintenance500504ESB-14-007TimeoutTimeout504401ESB-14-008Client\_Id/Client\_Secret /Grant\_Type tidak validInvalid Client\_Id/Client\_Secret /Grant\_Type401401ESB-14-009Tidak berhakUnauthorized401404ESB-14-011Service tidak tersediaService doesn't exist404401ESB-14-012Tidak berhak mengakses service iniNot allowed to access this service401401ESB-14-014ChannelID/CredentialID tidak validInvalid ChannelID/CredentialID 401403ESB-14-019Koneksi tidak diperbolehkanConnection not allowed403400ESB-14-021API Key tidak validInvalid API Key400400ESB-14-022API Key/API Secret tidak cocokAPI Key/API Secret mismatch4004000002Transaksi GagalTransaction FailedCopartner is not registered4000012Transaksi GagalTransaction FailedCopartner is blocked4000022Transaksi GagalTransaction FailedCopartner max retry exceeded 4000032Transaksi GagalTransaction FailedCopartner session key not exist4001012Kartu ErrorCard ErrorCheck expired date failed4001022Transaksi gagal, nominal Top Up kurang dari minimumTransaction failed, Top Up amount less than minimumCheck min amount failed4001032Transaksi gagal, nominal Top Up melebihi saldo maksimumTransaction failed, Top Up amount exceeds maximum balanceCheck max balance failed4001052Transaksi gagalTransaction failedCheck mandatory field failed4001062Transaksi gagalTransaction failedParse request data failed4001102Transaksi gagalTransaction failedDuplicate transaction4001122Transaksi gagalTransaction failedCheck original data failed4001132Transaksi gagalTransaction failedCheck tid failed4001142Transaksi gagalTransaction failedCheck mid failed4001152Kartu errorCard errorCard blacklisted4003013Transaksi gagalTransaction failedCheck topup limit timeout failed4003023Transaksi gagalTransaction failedRequest topup debit timeout failed4003045Transaksi gagalTransaction failedCheck atuf timeout failed4004662Transaksi tidak dapat diproses. Silahkan coba beberapa saat lagi.Transaction can not be processed. Please try again later.Error dari EAI apabila terjadi koneksi timeout saat memanggil ke service lain.4004045Maaf, transaksi anda tidak dapat diprosesSorry, your transaction can not be processedIdentifier Channel tidak diizinkan mengakses4004009Field (0) harus diisiMissing mandatory field (0)Field (0) harus diisi4004038Panjang field (0) tidak valid Field (0) length not validPanjang field (0) tidak valid4004039Isian input (0) tidak validInput for field (0) not validIsian input (0) tidak valid4004112Inputan String JSON tidak validInvalid JSON String inputFormat JSON input invalid 4004113Transaksi ditolakTransaction rejectedClientID tidak dapat mengakses ke service4004414Transaksi gagalTransaction failedGagal saat konversi(format data yang diterima tidak sesuai)4004415Transaksi gagalTransaction failedRecord yang ingin dihapus tidak ada.4004416Transaksi gagalTransaction failedFile Limit TopUp tidak ada4004417Transaksi gagalTransaction failedFile Limit TopUp tidak bisa dibuka 4004418Transaksi gagalTransaction failedFile Limit TopUp tidak bisa dibaca4004419Transaksi gagalTransaction failedRecord sudah ada ketika tambah record baru 4004420Transaksi gagalTransaction failedRecord tidak ada ketika Change/Update4004421Transaksi gagalTransaction failedGagal menulis record4004422Transaksi gagalTransaction failedGagal hapus record4004423Transaksi gagalTransaction failedFIID yang dikirimkan tidak sesuai/data FIID tidak ada4004424Transaksi gagalTransaction failedTransaksi code tidak sesuai(sesuai dengan data element 91) 4004425Transaksi gagalTransaction failedFIID tidak terdaftar4004426Transaksi gagalTransaction failedFile request tidak sesuai(sesuai dengan data element 101)4004428Transaksi gagalTransaction failedGagal ketika mencoba mencari nomor kartu4004429Transaksi gagal, jumlah top up melebihi limit bulananTransaction failed, Top Up amount exceeds monthly limitLimit kartu Flazz habis.4004430Transaksi gagalTransaction failedUnik id tidak ditemukan4004431Transaksi gagalTransaction failedKoneksi saat dari atau ke b24 terputus4004432Transaksi gagalTransaction failedKoneksi saat dari atau ke b24d itolak4004414Transaksi gagalTransaction failedGagal saat konversi(format data yang diterima tidak sesuai)4004415Transaksi gagalTransaction failedRecord yang ingin dihapus tidak ada.4004416Transaksi gagalTransaction failedFile Limit TopUp tidak ada4004417Transaksi gagalTransaction failedFile Limit TopUp tidak bisa dibuka 4004418Transaksi gagalTransaction failedFile Limit TopUp tidak bisa dibaca4004419Transaksi gagalTransaction failedRecord sudah ada ketika tambah record baru 4004420Transaksi gagalTransaction failedRecord tidak ada ketika Change/Update4004421Transaksi gagalTransaction failedGagal menulis record 4004422Transaksi gagalTransaction failedGagal hapus record 4004423Transaksi gagalTransaction failedFIID yang dikirimkan tidak sesuai/data FIID tidak ada4004424Transaksi gagalTransaction failedTransaksi code tidak sesuai(sesuai dengan data element 91) 4004425Transaksi gagalTransaction failedFIID tidak terdaftar4004426Transaksi gagalTransaction failedFile request tidak sesuai(sesuai dengan data element 101)4004427Transaksi gagalTransaction failedCategory selain C,F,T,R. 4004428Transaksi gagalTransaction failedGagal ketika mencoba mencari nomor kartu4004429Transaksi gagal, kartu sumber dana lebih dari limit Transaction failed, card for source of founds is over limitLimit habis kartu akses habis4004430Transaksi gagalTransaction failedUnik id tidak ditemukan4004431Transaksi gagalTransaction failedKoneksi saat dari atau ke b24 terputus4004432Transaksi gagalTransaction failedKoneksi saat dari atau ke b24 ditolak4000000Transaksi BerhasilTransaction SuccessTransaksi berhasil4005202Transaksi 0 rupiahTransaction amount 0Transaksi 0 Rupiah4005204Maaf saldo tidak mencukupi Sorry your balance is not enough Maaf saldo tidak mencukupi4005206Rekening ditolakAccount RejectedRekening ditolak4005700Maaf transaksi harus IDRSorry, currency must IDRMaaf transaksi harus IDR4005701Rekening ditolakAccount RejectedRekening ditolak 4005702Maaf rekening sudah tidak aktifSorry, your account is dormantMaaf rekening sudah tidak aktif4005711Maaf rekening anda masuk daftar hitamSorry, your account is blacklistRekening anda diblokir/masuk daftar hitam 4005999Sistem sedang tidak tersediaSystem unavailableSistem sedang tidak tersedia

POST /flazz/

Request

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-white hljs
REQUEST FOR GET SESSION KEY

POST /flazz/getKey HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
ChannelID: [channelID]
CredentialID: [credentialID]
{
    "reference_id" : "123456789012",
    "request_data" : "01010000202001171513224553594230303031454D424341303031885000015999824E180B6E8227061B4E8A275E3D373702021B4B5E821B063D4E0B370627024E843D3D274B181B8A1B375E4E8A4E020B203"
}

REQUEST FOR TOP UP REQUEST

POST /flazz/topup HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
ChannelID: [channelID]
CredentialID: [credentialID]
{
    "reference_id" : "123456789012",
    "request_data" : "01010024014500280000003800311900000000102020020317152011013049124553594230303031454D4243413030318850000159996E8282370684824E5E02848A8482188A064B06064E0627278A184B4E4E828A3D848A4B82028402064B8206023D1C0B1C1B0B826E1C063D6E185E6E0B4E821B844B37273718064E82024B4E1B0282185E4B37273718064E82024B4E1B0282185E4B37273718064E82024B4E1B0282185E4E0B3D181C4E1B27026E3D5E6E186E1B8A1C3D273782028A068A274B3D0B826E1AE5BCBB7C0226A4DB8C"
}

REQUEST FOR REVERSAL

POST /flazz/reversal HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
ChannelID: [channelID]
CredentialID: [credentialID]
{
    "reference_id" : "123456789012",
    "request_data" : ""01010024014500010036938400291613000000102020012315164520013049124553594230303031454D424341303031885000015999275E6E824B1B4E021B1B1C4E0606023D4E844E5E8A021B6E8A4B4B1C6E4E6E8A4E5E8A82826E374B1B1B0218188A068482061C278A5E0B180B8A6E8A3D0B8A823784186E0B4E8237024E020B6E02825E1B8A020B270B82270227378202183D846E84824E021B4B844B82273D02183D5E5E8A061C6E06068A8A061B024B1B3D4B068A1C02020B6E824E8A371B1C064E8A3D4B84270682278A5E84374E37061C843D4B84270682278A5E84374E37061C843D4B84270682278A5E84374E37061C84C0396097845AD44C26CC"
}

REQUEST FOR ACK

POST /flazz/ack HTTP/1.1
Host: sandbox.bca.co.id
Authorization: Bearer [access_token]
Content-Type: application/json
Origin: [yourdomain.com]
X-BCA-Key: [api_key]
X-BCA-Timestamp: [timestamp]
X-BCA-Signature: [signature]
ChannelID: [channelID]
CredentialID: [credentialID]
{
    "reference_id" : "123456789012",
    "request_data" : ""01010024014500010036938400291613000000102020012315164520013049124553594230303031454D424341303031885000015999275E6E824B1B4E021B1B1C4E0606023D4E844E5E8A021B6E8A4B4B1C6E4E6E8A4E5E8A82826E374B1B1B0218188A068482061C278A5E0B180B8A6E8A3D0B8A823784186E0B4E8237024E020B6E02825E1B8A020B270B82270227378202183D846E84824E021B4B844B82273D02183D5E5E8A061C6E06068A8A061B024B1B3D4B068A1C02020B6E824E8A371B1C064E8A3D4B84270682278A5E84374E37061C843D4B84270682278A5E84374E37061C843D4B84270682278A5E84374E37061C84C0396097845AD44C26CC"
}
```

POST /flazz/

Response

![](/content/ReskinDeveloperApiBcaCoId/images/icons/copy.png)

```custom !bca-text-black hljs
RESPONSE FOR GET SESSION KEY

{
    "response_data" : ""01010000202001171513224553594230303031454D4243413030318850000159996DD3D35D3422B7B022CF34B7B0CD6D34345DB07A7ACFA45DCFD3B7A4D3A46DA4CFB0D3D35334B0347D6E537A347A5DCDABB"
}

RESPONSE FOR REQUEST TOP UP

{
    "response_data" : "01240145002800000038000000100000000000160230202002191019275D5D6E5DB05D6D7A7DCF7A22227A6EB07D5D3422CD68CF6853A4A4CDB77AD35DD334347AB7A45DB7D33434B06D68D3D334B7CF7D92D35D34A46E7AD3B7CF34685D34CF34A453D36EB75DD368B76EA4925D34CF34A453D36EB75DD368B76EA4925D34CF34A453D36EB75DD368B76EA4925D34CF34A453D36EB75DD368B76EA4925D34CF34A453D36EB75DD368B76EA492F8B2"
}

RESPONSE FOR REVERSAL

{
    "response_data" : "null"
}

RESPONSE FOR ACK

{
    "response_data" : "null"
}
```

![](/content/ReskinDeveloperApiBcaCoId/images/icons/close.png)

Anda akan meninggalkan developer.bca.co.id

Dengan mengklik ‘Lanjutkan’, Anda akan diarahkan menuju website di luar www.developer.bca.co.id yang tidak terafiliasi dengam BCA dan mungkin memiliki tingkat sekuriti yang berbeda. BCA tidak bertanggung jawab dan tidak mendukung, menjamin, mengendalikan konten, mengendalikan ketersediaan dan perspektif atas produk-produk atau layanan-layanan yang ditawarkan atau dinyatakan oleh website tersebut.

LanjutkanBatal

[![BCA](/-/media/Project/DeveloperApiBcaCoId/Images/Reskin/logo-bca-logo-only-light.svg)](/id)

Kantor Pusat

Menara BCA, Grand IndonesiaJl. MH Thamrin No. 1Jakarta 10310

[![](/-/media/Project/DeveloperApiBcaCoId/Images/Footer/location.png?h=17&w=16&hash=A6364B4F246908939DFADB9D9A7100F2) Lokasi BCA Lainnya](https://www.bca.co.id/lokasi-bca?utm_source=devapi)

Hubungi Kami

[![](/-/media/Project/DeveloperApiBcaCoId/Images/Footer/telp.png?h=17&w=17&hash=B6A39868F9A66FFF57D3576B3DF364DC)(021) 235 88000](tel:02123588000) [![](/-/media/Project/DeveloperApiBcaCoId/Images/Reskin/icon/whatsapp.png?h=64&w=64&hash=A901868E8C7DEE8AE1BEC44C02032373)+62 811 1500 998](https://wa.me/628111500998?text=%23API) [![](/-/media/Project/DeveloperApiBcaCoId/Images/Footer/customer-support.png?h=15&w=16&hash=90E2FD10BF1671C86069DB771E0F86D1)API Support](mailto:api_support@bca.co.id?subject=SUPPORT API BCA) [![](/-/media/Project/DeveloperApiBcaCoId/Images/Footer/contact-us.png?h=17&w=17&hash=A73CFE657904B80C91173AD35D5FC9BA)Hubungi Kami](/id/Hubungi-Kami)

Media Sosial

[![](/-/media/Project/DeveloperApiBcaCoId/Images/Footer/facebook.png?h=19&w=19&hash=163305E579FCFB8608066EFBAD840A0D)GoodLife BCA](http://www.facebook.com/GoodLifeBCA) [![](/-/media/Project/DeveloperApiBcaCoId/Images/Footer/instagram.png?h=19&w=19&hash=48996CBA46134FBADFD8BD3160FC6AC2)@goodlifebca](https://www.instagram.com/goodlifeBCA/)

[![](/-/media/Project/DeveloperApiBcaCoId/Images/Footer/x.png?h=17&w=17&hash=4F7C403052754340789D6BBA7595929C)@BankBCA](https://twitter.com/BankBCA) [![](/-/media/Project/DeveloperApiBcaCoId/Images/Footer/youtube.png?h=14&w=17&hash=3B0ECEE26E3435B04C2796E4661C7448)Solusi BCA](https://www.youtube.com/solusiBCA)

[Lihat Semua Media Sosial](https://www.bca.co.id/id/tentang-bca/media-riset/Social-Media?utm_source=devapi)

BCA berizin dan diawasi oleh Otoritas Jasa Keuangan & Bank Indonesia

BCA merupakan peserta penjaminan LPS. Maksimum nilai simpanan yang dijamin LPS per Nasabah per Bank adalah Rp2 miliar. Untuk cek Tingkat Bunga Penjaminan LPS, klik [di sini](https://apps.lps.go.id/BankPesertaLPSRate).

- [Kebijakan](https://www.bca.co.id/id/informasi/Kebijakan?utm_source=devapi)
- \|
- [Syarat penggunaan](/id/Syarat-Penggunaan)
- \|
- [RIPLAY](https://www.bca.co.id/-/media/Files/2023/20230927-final-rev-riplay-umum-api-bca-2109023-pdf-des)