Keamanan - SNAP Developer Site

[![](/img/icon/snap-header-white-100.png)](/)

- [**Beranda**](/)
- [**Info**](/info)
- [**API**](/api-services)
- [**Direktori Publikasi**](/direktori-publikasi)
  - [**Sign In**](/sign-in)

[![](/img/icon/bi-logo-full-100.png)](/)

![](/img/card/keamanan-100.png)

### Keamanan

Overview


Guides


Code Snippets


Response Code


Aplikasi Pengujian


#### OVERVIEW

#### **STANDAR TEKNIS DAN STANDAR KEAMANAN**

Standar keamanan merupakan bagian dari Standar Nasional _Open_ API Pembayaran yang bertujuan untuk memastikan kerahasiaan data, integritas data dan sistem, serta ketersediaan layanan, mengatur mengenai standar untuk otentikasi, otorisasi, enkripsi untuk menjamin integritas dan kerahasiaan data, terdapatnya _business continuity plan_, maupun penerapan _fraud detection system_ untuk memitigasi potensi _fraud_. Selain mengacu pada standar keamanan tersebut, Penyedia Layanan dan Pengguna Layanan harus menerapkan kontrol dan perlindungan menyeluruh terhadap data dan informasi dari potensi risiko siber untuk melindungi sistem, data Konsumen maupun data terkait Penyedia Layanan dan/atau Pengguna Layanan.

#### **KOMPONEN STANDAR TEKNIS DAN STANDAR KEAMANAN**

Standar teknis dan keamanan dari Standar Nasional _Open_ API Pembayaran menstandarkan hal-hal sebagai berikut:

01. Tipe Arsitektur
02. Format Data
03. _Character Encoding_
04. Komponen _HTTP Method_
05. Komponen Struktur Format _Header - Access Token_ ( _Business to Business_ (B2B) dan _Business to Business to Consumer_ (B2B2C))
06. Komponen Struktur Format _Header – Transaction_ (B2B dan B2B2C)
07. Komponen _Server Authentication Method_
08. Komponen _Client Authentication Method_
09. Komponen Standar Enkripsi
10. Komponen _Secured Channel Communication_
11. Komponen Standardisasi _URI Path_
12. Komponen Standardisasi _Business Continuity Plan_ ( _Reliability_, _Availability_, dan _Scalability_)
13. Komponen Standardisasi Keamanan Lainnya

##### **Tipe Arsitektur API**

Tipe arsitektur yang digunakan **adalah** _**Representational State Transfer**_ **(REST) API**.

##### **Format Data**

Format data yang digunakan pada _request body_ dan _response body_ adalah _**JavaScript Object Notation**_ **(JSON)**.

##### _Character Encoding_

Standar _character encoding_ yang digunakan adalah **UTF-8**.

##### **Komponen _HTTP Method_**

_HTTP Method_ berfungsi sebagai identifikasi terhadap aksi yang ingin dilakukan pada suatu sumber daya ( _resource_) dengan komponen HTTP-Verbyang pada umumnya digunakan. _HTTP-Verb_ yang digunakanadalah:

1. _**POST**_ _Request_
2. _**GET**_ _Request_
3. _**DELETE**_ _Request_
4. _**PUT**_ _Request_

Sebagai pertimbangan keamanan, untuk _service_ _**get Access Token**_ menggunakan _**POST**_ _Request_. Untuk _services_ lainnya menggunakan HTTP-Verb yang disesuaikan untuk tipe operasi dan _resource_ yang diakses. Penggunaan HTTP _method_ untuk masing-masing _service_ disebutkan pada tabel informasi umum pada dokumen spesifikasi teknis SNAP.

##### **Komponen _Server Authorization_ dan _Authentication Method_**

Otorisasi adalah metode bagi Penyedia Layanan untuk memberikan akses _request_ API dari Pengguna Layanan. Standar yang digunakan adalah:

- _OAuth_ 2.0 sesuai RFC6749
- _Bearer token_ sesuai RFC6750

Dalam memberikan akses kepada Pengguna Layanan, Penyedia Layanan melakukan otentikasi untuk memvalidasi Pengguna Layanan oleh Penyedia Layanan. Sarana yang digunakan adalah\_credential\_ yang dipertukarkan pada saat proses pembentukan kerja sama, yaitu _client secret_ dan pasangan _public/private key_, yang digunakan bersama dengan algoritma kriptografi tertentu.

##### **Komponen _Client Authentication Method_**

_Client Authentication Method_ adalah metode otentikasi untuk memvalidasi konsumen. Standar _Two-Factor Authentication_ yang digunakan adalah:

1. _Short Message Service_ (SMS) TOTP ( _Time based One Time Password_)
2. SMS TOTP dengan 6 digit- numerik dengan durasi 5 menit;
3. _Personal Identification Number_ (PIN)
4. PIN dengan 6 digit-numerik
5. _Biometric_( _Fingerprint_ & _Face Recognition_)
6. Lainnya

##### **Komponen Standar Enkripsi**

Model enkripsi terhadap _message_ yang digunakan yaitu enkripsi asimetris dan simetris, menggunakan kombinasi _Private Key_ dan _Public Key_, dengan standar sebagai berikut:

1. _Standard Asymmetric Encryption Signature_:

1. **SHA256withRSA** dengan _Private Key_ ( **Kpriv** ) dan _Public Key_ ( **Kpub** ) (256 bits)
2. _Standard Symmetric Encryption Signature_
1. **HMAC\_SHA512** (512 bits)
3. _Standard Symmetric Encryption_
1. AES-256 dengan **client secret** sebagai _encryption key_.

##### **Komponen _Secured Channel Communication_**

\_Secured channel communication\_adalah kanal komunikasi yang aman untuk menjaga kerahasiaan _message_ yang dikirimkan. Standar yang akan digunakan adalah:

1. _Transport Layer Security_ (TLS) 1.3
2. Memiliki kemampuan untuk negosiasi ke TLS 1.2 namun dengan modul enkripsi minimum yang telah ditentukan sebagai berikut:
3. Memiliki kemampuan untuk negosiasi ke TLS 1.2 namun dengan modul enkripsi minimum yang telah ditentukan sebagai berikut:
   - TLS\_DHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
   - TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
   - TLS\_DHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384
   - TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384

Penggunaan TLS 1.2 dengan modul enkripsi minimum sebagaimana dimaksud pada angka 2 huruf a sampai dengan huruf d, hanya dapat diterapkan oleh Penyedia Layanan dan Pengguna Layanan sampai dengan tanggal 30 Juni 2026.

##### **Komponen Standardisasi _Uniform Resources Indentifier_ (URI\_) Path\_**

Standardisasi URI _resource path_ dengan format sebagai berikut:

_/\[domain\_......\_api\]//\[version\]/\[service-group\]/\[product-type\]_

1. \[domain\_....\_api\]
_The constant string of specific respective of PJP/Non-PJP api domain name_

2. \[version\]
_The version of the APIs expressed as_ / **v** \[major-version\] **.** \[minor-version\]/

3. \[service-group\]
_The service-group identifies the group of endpoints_

4. \[product-type\]
_Details of the resource if such service has another product definition underneath_


##### **Prinsip-Prinsip _Business Continuity Plan_ (BCP)**

Prinsip-prinsip standar BCP adalah sebagai berikut:

1. _**Reliability**_– untuk memastikan ketersediaan data dan layanan serta untuk menjamin kesinambungan proses bisnis.
2. _**Availability**_– memastikan sistem dan data tersedia untuk pengguna yang berwenang ketika mereka membutuhkannya. Melalui _Active-Active deployment_ atau _Active - Stand-By default_.
3. _**Scalability**_– memastikan layanan dari produk jasa keuangan memiliki _response time_ yang terukur.

Standar BCP adalah sebagai berikut:

No.Infrastruktur Pendukung Open API Pembayaran berbasis SNAPPersyaratan1Tipe Data Recovery Center untuk API Management   i.          HOT DRC (RTO: <1 Hour, RPO: <1 Hour)   ii.        Replikasi data harus mendukung SLA RTO & RPO < 1 jam.2Kategori data center yang digunakan untuk API ManagementRTO: <1 Hour RPO: <1 Hour3Terdapat regular backup database & transaction log·      Backup database (harian, mingguan, bulanan)·      Backup transaction log·      Retensi data & log : 10 tahun

##### **Standar Keamanan Lainnya**

1. **Ketersediaan Kebijakan Tertulis Terkait Sistem Informasi**

Penyedia Layanan dan Pengguna Layanan memiliki kebijakan atau prosedur tertulis terkait sistem informasi yang paling sedikit meliputi:

1)       Manajemen user4)    Pengembangan secure application2)       Manajemen siber5)     Change management3)       Pengamanan dan perlindungan data (termasuk penyimpanan data)6)     Tata kelola sistem informasi

1. **Pemenuhan Sertifikasi dan/atau Standar Keamanan dan Keandalan Sistem Informasi**

1. Penyedia Layanan dan Pengguna Layanan _Open_ API Pembayaran berbasis SNAP mengadopsi praktik-praktik umum terbaik dalam implementasi keamanan dan keandalan sistem informasi.
2. Penyedia dan Pengguna Layanan direkomendasikan memiliki sertifikasi dan/atau standar keamanan dan keandalan sistem informasi yang berlaku umum sesuai dengan jenis layanan yang diselenggarakan.
2. _**Fraud Detection System**_ **(FDS)**


FDS adalah _tools_ yang dipergunakan untuk mencegah, mendeteksi, memitigasi, menganalisis aktivitas _fraudulent_ pada saat aktivitas tersebut teridentifikasi masuk ke dalam sistem serta mampu memberikan informasi/ _alert_ kepada petugas yang berwenang.

_Open_ API Pembayaran berbasis SNAP dilengkapi dengan penerapan FDS.

FDS didukung oleh kebijakan/prosedur dan sumber daya manusia yang diperlukan dalam implementasi/operasional FDS.

Fitur yang direkomendasikan diimplementasikan dalam FDS namun tidak terbatas pada:

1. Memiliki fleksiblitas untuk mengkonfigurasi \_rules/\_parameter sebelum dan sesudah implementasi FDS
2. Memiliki kemampuan untuk menerima dan mengolah data _fraud_ yang bersumber dari luar
3. Memiliki kemampuan untuk menganalisis, memitigasi dan/atau memprioritaskan tindak lanjut berdasarkan potensi serangan/ _fraud_
4. Kemampuan mendeteksi/mencegah anomali transaksi
5. Memiliki kemampuan untuk mendeteksi/mencegah potensial _fraud_ sejak fase pendaftaran akun nasabah.

_Rules_/parameter yang direkomendasikan diimplementasikan dalam FDS namun tidak terbatas pada:

1)    Waktu transaksi5)   Nominal9)      Excessive login2)    Frekuensi transaksi6)  Negative balance10)   Device ID3)    Velocity\*)7)  Akun dormant11)   Fraudster ID/black list akun4)    Incorrect PIN/OTP/Password/other authentication method8)  Negara asal dan/atau negara tujuan transaksi12)   Lokasi transaksi\*)

\*)dalam hal transaksi mencakup data lokasi

1. **Pelaksanaan Audit Secara Berkala**

Penyedia Layanan dan Pengguna Layanan melakukan audit secara berkala terhadap implementasi SNAP. Audit dilakukan oleh auditor independen.

1. **Aspek Keamanan lainnya**
1. Adanya penerapan _whitelisted IP_ pada perangkat/aset yang digunakan untuk _Open_ API Pembayaran berbasis SNAP dan perangkat pendukung lainnya.
2. Memiliki _firewall_

_Open_ API Pembayaran berbasis SNAP dilengkapi dengan _Web Application Firewall_ baik menggunakan _Cloud Based_, _Network Based_ ataupun _Host-Based Firewall_ yang dapat melindungi dari _cyber attack_ seperti _cross-site-scripting_ (XSS), _cross-site forgery_, SQL _injection_, DDoS, _malware_ dan lain lain.

**Pengelolaan yang** **direkomendasikan diimplementasikan dalam** _**Firewall**_ **namun tidak terbatas pada:**

1)   Adanya dokumen firewall (tujuan, layanan pengguna firewall, rules)4)    Manajemen/monitoring network traffic2)   Access Control List (ACLs)5)    Pengujian firewall secara berkala3)   Rules antara lain packet filtering, antispoofing filter, user permit rules, permit management, alert untuk suspicious traffic dan traffic log6)    Pengkinian firewall secara reguler

#### GUIDES

##### **Komponen Struktur Format _Header_ – _Access Token_ (B2B dan B2B2C)**

Setiap Pengguna Layanan yang ingin melakukan akses terhadap layanan API yang terdaftar untuk model _use case_:

1. B2B (integrasi antaraPJP Penyedia Layanan dan Pengguna Layanan); atau
2. B2B2C (integrasi antara PJPPenyedia Layanan, Pengguna Layanan, dan Konsumen)

harus melakukan _access token_ _request_ terlebih dahulu dengan standar sebagai berikut:

**Komponen Struktur _Format Header_ – _Access Token Request_ (B2B)**

Service Code73NameAPI Access Token B2BVersion1.0HTTP MethodPOSTPath../{version}/access-token/b2b

Struktur _Format Header_ API untuk _Access Token Request_ (B2B):

AreaFieldAttributeTypeDescriptionHeaderContent-TypeMandatoryStringString represents indicate the media type of the resource (e.g. application/json, application/pdf)X-TIMESTAMPMandatoryStringClient's current local time in yyyy-MM- ddTHH:mm:ssTZD formatX-CLIENT- KEYMandatoryStringClient’s client\_id (PJP Name) (given at completion registration process )X-SIGNATUREMandatoryStringNon-Repudiation & Integrity checkingX-Signature : dengan algoritma asymmetric signature SHA256withRSA(Private\_Key, stringToSign). stringToSign = client\_ID + “\|” + X-TIMESTAMPBodygrantTypeMandatoryString“client\_credentials” : The client can request an access token using only its client credentials (or other supported means of authentication) when the client is requesting access to the protected resources under its control (OAuth 2.0: RFC 6749 & 6750)additionalInfoOptionalObjectAdditional Information

**Komponen Struktur _Format Header_ – _Access Token Response_ (B2B)**

Sebagai _response_ dari _access token request_, diatur standar dengan format sebagai berikut:

AreaFieldAttributeTypeDescriptionHeaderX-TIMESTAMPMandatoryStringClient's current local time in yyyy-MM- ddTHH:mm:ssTZD formatX-CLIENT- KEYMandatoryStringClient’s client\_id (PJP Name) (given at completion registration process )BodyresponseCodeConditionalStringRefer to standar data dan spesifikasi teknis part 6 (Response Code).If access token failed to generate, this value must be filled.responseMessageConditionalStringRefer to standar data dan spesifikasi teknis part 6 (Response Message).If access token failed to generate, this value must be filled.accessTokenMandatoryString (2048)A string representing an authorization issued to the client that used to access protected resourcestokenTypeMandatoryStringThe access token type provides the client with the information required to successfully utilize the access token to make a protected resource request (along with type-specific attributes)Token Type Value:•   “Bearer”: includes the access token string in the request•   “Mac”: issuing a Message Authentication Code (MAC) key together with the access token that is used to sign certain components of the HTTP requestsReference: OAuth2.0 RFC 6749 & 6750expiresInMandatoryStringSession expiry in seconds: 900 (15 menit)additionalInfoOptionalObjectAdditional information for custom use that are not provided by SNAP

**Komponen Struktur _Format Header_ – _Access Token Request_ (B2B2C)**

Service Code74NameAPI Access Token B2B2CVersion1.0HTTP MethodPOSTPath../{version}/access-token/b2b2c

Struktur _Format Header_ API untuk _Access Token Request_ (B2B2C):

AreaFieldAttributeTypeDescriptionContent-TypeMandatoryStringString represents indicate the media type of the resource (e.g. application/json, application/pdf)HeaderX-TIMESTAMPMandatoryStringClient's current local time in yyyy-MM-ddTHH:mm:ssTZD formatX-CLIENT-KEYMandatoryStringClient’s client\_id (PJP Name)(given at completion registration process)X-SIGNATUREMandatoryStringNon-Repudiation & Integrity checkingX-Signature : dengan algoritma asymmetric signature SHA256withRSA(Private\_Key, stringToSign). stringToSign = client\_ID + “\|” + X-TIMESTAMPgrantTypeMandatoryStringApply token request key type, can be AUTHORIZATION\_CODE or REFRESH\_TOKEN.BodyauthCodeConditionalString (256)The authorization code received after the User provides the consent. Mandatory if grantType = AUTHORIZATION\_CODErefreshTokenConditionalString (512)Refresh token to get a new accessToken where the User doesn't need to provide the consent again. Mandatory if grantType = REFRESH\_TOKEN. Refresh Token should be less than access token validity and will be manage by the PJP’s application to generate a new access\_tokenadditionalInfoOptionalObjectAdditional information for custom use that are not provided by SNAP

**Komponen Struktur _Format Header_ – _Access Token Response_ (B2B2C)**

Sebagai _response_ dari _access token request_ diatur standar dengan format sebagai berikut:

AreaFieldAttributeTypeDescriptionHeaderX-TIMESTAMPMandatoryStringClient's current local time in yyyy-MM-ddTHH:mm:ssTZD formatX-CLIENT-KEYMandatoryStringClient’s client\_id (PJP Name)(given at completion registration process)BodyresponseCodeConditionalStringRefer to standar data dan spesifikasi teknis part 6 (Response Code).If access token failed to generate, this value must be filled.responseMessageConditionalStringRefer to standar data dan spesifikasi teknis part 6 (Response Code)If access token failed to generate, this value must be filled.accessTokenMandatoryString (2048)A string representing an authorization issued to the client that used to access protected resources.tokenTypeMandatoryStringThe access token type provides the client with the information required to successfully utilize the access token to make a protected resource request (along with type-specific attributes)Token Type Value:•   “Bearer”: includes the access token string in the request•   “Mac”: issuing a Message Authentication Code (MAC) key together with the access token that is used to sign certain components of the HTTP requestsReference: OAuth2.0 RFC 6749 & 6750accessTokenExpiryTimeMandatoryStringTime when the accessToken will be expired.Access token valid time will be 15 daysformat ISO8601refreshTokenMandatoryStringA random string that can be used by specific client to get a refreshed accessToken to prolong the access to the User's resources.refreshTokenExpiryTimeMandatoryStringTime when the refreshToken will be expired. Refresh Token should be less than access token validity and will be manage by the PJP’s application to generate a new access\_tokenformat ISO8601additionalInfoOptionalObjectAdditional information for custom use that are not provided by SNAP

##### **Komponen Struktur Format _Header_ – _Transaction_ (B2B dan B2B2C)**

Standar struktur _format header_ untuk API level transaksi adalah sebagai berikut:

**Komponen Struktur _Format Header_ – _Transaction Request_ (B2B)**

Struktur _format header_ API untuk _transaction request_ (B2B):

AreaFieldAttributeTypeDescriptionHeaderContent-TypeMandatoryStringString represents indicate the media type of the resource (e.g. application/json, application/pdf)AuthorizationConditionalStringRepresents access\_token of a request; string starts with keyword “Bearer ” followed by access\_token (e.g. Bearer eyJraWQiOi...Jzc29zIiwiY)X-TIMESTAMPMandatoryStringClient's current local time in yyyy-MM-ddTHH:mm:ssTZD formatX-SIGNATUREMandatoryStringRepresents signature of a request.Identify Signature Type usedValue:1 - Symmetric Signature with Get Token2 - Asymmetric Signature without Get TokenDefault Value: 11.      Symetric-Signature :HMAC\_SHA512 (clientSecret, stringToSign) dengan formula stringToSign = HTTPMethod +”:“+ EndpointUrl +":"+ AccessToken +":“+ Lowercase(HexEncode(SHA-256(minify(RequestBody))))+ ":“ + TimeStamp2.  Asymetric-Signature :SHA256withRSA (clientSecret, stringToSign) dengan formulastringToSign = HTTPMethod +”:“+ EndpointUrl +":“+ Lowercase(HexEncode(SHA-256(minify(RequestBody)))) + ":“ + TimeStampCatatan:1.      Endpoint URL lengkap termasuk seluruh parameter pada URL terkait (Relative path, contoh: Path pada informasi umum setiap API service)2.      Untuk parameter minify(Request Body), dalam hal tidak terdapat Request Body maka digunakan string kosong.ORIGINOptionalStringOrigin Domain www.yourdomain.comX-PARTNER-IDMandatoryString (36)Unique ID for a partnerX-EXTERNAL- IDMandatoryString (36)Numeric String. Reference number that should be unique in the same dayCHANNEL-IDMandatoryString (5)PJP’s channel idDevice identification on which the API services is currently being accessed by the end user (customer)

**Contoh _Header_ – _Transaction Request_ (B2B):**

```http
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-17T10:55:00+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
CHANNEL-ID: 95221

```

**Komponen Struktur _Format Header_ – _Transaction Request_ (B2B2C)**

Struktur _format header_ API untuk _transaction request_ (B2B2C):

AreaFieldAttributeTypeDescriptionHeaderContent-TypeMandatoryStringString represents indicate the media type of the resource (e.g. application/json, application/pdf)AuthorizationMandatoryStringRepresents access\_token of a request; string starts with keyword “Bearer ” followed by access\_token (e.g. Bearer eyJraWQiOi...Jzc29zIiwiY)Authorization-CustomerMandatoryStringRepresents access\_token of a request belong customer; string starts with keyword “Bearer ” followed by access\_token (e.g. Bearer eyJrsWaiOi...Jzc523awiY)X-TIMESTAMPMandatoryStringClient's current local time in yyyy-MM-ddTHH:mm:ssTZD formatX-SIGNATUREMandatoryStringRepresents signature of a requestX-Signature : algoritma symmetric signature HMAC\_SHA512 (clientSecret, stringToSign) dengan formulastringToSign = HTTPMethod +”:“+ EndpointUrl +":"+ AccessToken +":“+ Lowercase(HexEncode(SHA-256(minify(RequestBody))))+ ":“ + TimeStampCatatan:1.    Endpoint URL lengkap termasuk seluruh parameter pada URL terkait (Relative path, contoh: Path pada informasi umum setiap API service)2.    Untuk parameter minify(Request Body), dalam hal tidak terdapat Request Body maka digunakan string kosong.ORIGINOptionalStringOrigin Domain www.yourdomain.comX-PARTNER-IDMandatoryString (36)Unique ID for a partnerX-EXTERNAL-IDMandatoryString (36)Numeric String. Reference number that should be unique in the same dayX-IP-ADDRESSOptionalString (15)IP address of the end user (customer) using IPv4 formatExample: 172.31.255.255X-DEVICE-IDMandatoryString (400)Device identification on which the API services is currently being accessed by the end user (customer)Sample:Web Application:Mozilla / 5.0(Windows NT 10.0; Win64; x64)AppleWebKit / 537.36(KHTML, like Gecko)Chrome / 75.0.3770.100 Safari / 537.36 OPR / 62.0.3331.99Mobile Application:Android: android-20013adf6cdd8123fiOS: 72635bdfd223yvjm7246nsdj34hd4559393kjh42CHANNEL-IDMandatoryString (5)PJP’s channel idDevice identification on which the API services is currently being accessed by the end user (customer)X-LATITUDEOptionalString (10)Location on which the API services is currently being accessed by the end user (customer)Refer to ISO 6709 Standard representation of geographic point location by coordinates±DD.DDDD format (without minutes and seconds)±DD = three-digit integer degrees part of latitude.DDDD = variable-length fraction part in degreesSample:New York City:Latitude: +40.75X-LONGITUDEOptionalString (10)Location on which the API services is currently being accessed by the end user (customer)Refer to ISO 6709 Standard representation of geographic point location by coordinates±DDD.DDDD format (without minutes and seconds)±DDD = four-digit integer degrees part of latitude.DDDD = variable-length fraction part in degreesSample:New York City:Longitude: -074.00

**Contoh _Header_ – _Transaction Request_ (B2B2C)**

```http
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
CHANNEL-ID: 95221
X-LATITUDE: -6.108841
X-LONGITUDE: 106.7782137

```

**Komponen Struktur _Format Header_ – _Transaction Response_ (B2B dan B2B2C)**

Struktur _format header_ API untuk _transaction response_ (B2B dan B2B2C):

AreaFieldAttributeTypeDescriptionHeaderContent-TypeMandatoryStringString represents indicate the media type of the resourceX-TIMESTAMPMandatoryStringClient's current local time in yyyy-MM-ddTHH:mm:ssTZD format

**Contoh _Header_ – _Transaction Response_ (B2B dan B2B2C)**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-21T10:30:34+07:00

```

#### CODE SNIPPETS

##### Signature Access Token

##### Sample Request

```http
POST /api/v1/utilities/signature-auth HTTP/1.1
Host: localhost:44339
X-TIMESTAMP: 2020-01-01T00:00:00+07:00
X-CLIENT-KEY: 962489e9-de5d-4eb7-92a4-b07d44d64bf4
Private_Key: xoiOeVwrNlstGotvioidyzCdnAj5x6KreeTeqbdWFIA=

```

##### Sample Response

```json
{
    "signature": "07abc7c30d245c0ecce3ef6c2a9ac76cd9ffaf6d0d090773b429c2b97437dc72047f46d9890abb2d6d8af7594ea19787e79ec80e388e2f6225b449c2e4d82e7df50f37c301424aede785935703c1c70235ba4e59f589f571218ce2dce4c061e598f0f38d1ac57f3feb52cf0c31078e3ceee8d796c53983fe1d38ebd71155aaa613700dc21f5a57941b787f921af7d287e72687d5242eb3063d543d5f5923f76db008cf4f56fb9c618f7f4bc8366ae70d88705617487754563e629119013fa0549e6645b397524b3dd2fa7e7f3fe9faf0fbf77da59f566861a3c510241fd4416ab7d0eba42d998e1178da51d607e0ef866607c458837c762323be53827d86e875"
}

```

##### Signature Service

##### Sample Request

```http
POST /api/v1/utilities/signature-service HTTP/1.1
Host: localhost:44339
X-TIMESTAMP: 2020-01-01T00:00:00+07:00
X-CLIENT-SECRET: xS3vNQQgJRemFF0SZfXkZOq3r7kQ9n5YJgK4Wg0tVCQ=
HttpMethod: POST
EndpoinUrl: /api/v1/balance-inquiry
AccessToken: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJhOGQ2YmVkNS05MzdkLTQzZTUtYTlkMi1hYWY0ODFlZjc2YTIiLCJjbGllbnRJZCI6IjZhZTk1N2M0LTI4NjMtNDcxMy1hY2NlLWJhMTJkZTYzNmNmYyIsIm5iZiI6MTYxMTM4NjM4NywiZXhwIjoxNjExMzg3Mjg3LCJpYXQiOjE2MTEzODYzODd9.nUillb6567_zkM6Ys35OOG-YWGoo7Ik1odPJn1tR-ao
Content-Type: application/json
Content-Length: 119

{
    "accountNo" : "2000200202",
    "clientId"  : "962489e9-de5d-4eb7-92a4-b07d44d64bf4",
    "reqMsgId"  : "a"
}

```

##### Sample Response

```json
{
    "signature": "06a7c024bd3927ecea861ddb8605f96b382cd09e8f0ed71a4c4e8c810627212dd6973ab495b405a14dbad54f9fe23f8873b33ebcc546e2766912b7de4c225ef5"
}

```

##### Access Token B2B

##### Sample Request

```http
POST /api/v1/access-token/b2b HTTP/1.1
Host: localhost:44339
X-TIMESTAMP: 2020-01-01T00:00:00+07:00
X-CLIENT-KEY: 962489e9-de5d-4eb7-92a4-b07d44d64bf4
X-SIGNATURE: 07abc7c30d245c0ecce3ef6c2a9ac76cd9ffaf6d0d090773b429c2b97437dc72047f46d9890abb2d6d8af7594ea19787e79ec80e388e2f6225b449c2e4d82e7df50f37c301424aede785935703c1c70235ba4e59f589f571218ce2dce4c061e598f0f38d1ac57f3feb52cf0c31078e3ceee8d796c53983fe1d38ebd71155aaa613700dc21f5a57941b787f921af7d287e72687d5242eb3063d543d5f5923f76db008cf4f56fb9c618f7f4bc8366ae70d88705617487754563e629119013fa0549e6645b397524b3dd2fa7e7f3fe9faf0fbf77da59f566861a3c510241fd4416ab7d0eba42d998e1178da51d607e0ef866607c458837c762323be53827d86e875
Content-Type: application/json
Content-Length: 41

{
   "grantType":"client_credentials",
   "additionalInfo":{

   }
}

```

##### Sample Response

```json
{
   "responseCode":"2007300",
   "responseMessage":"Successful",
   "accessToken":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJiZjFmM2Q3ZS1kOTA3LTRkOWItODJlNC02Y2IxZGYxOTBlOWUiLCJjbGllbnRJZCI6IjZhZTk1N2M0LTI4NjMtNDcxMy1hY2NlLWJhMTJkZTYzNmNmYyIsIm5iZiI6MTYxMTQ2ODg1NiwiZXhwIjoxNjExNDY5NzU2LCJpYXQiOjE2MTE0Njg4NTZ9.-7HRhcyEh4y0qsG2H3DRdu0AeYv3MEJHfWRKhRBYcNU",
   "tokenType":"Bearer",
   "expiresIn":"900",
   "additionalInfo":{

   }
}

```

##### Access Token B2B2C

##### Sample Request

```http
POST /api/v1/access-token/b2b2c HTTP/1.1
Host: localhost:44339
X-TIMESTAMP: 2020-01-01T00:00:00+07:00
X-CLIENT-KEY: 962489e9-de5d-4eb7-92a4-b07d44d64bf4
X-SIGNATURE: 07abc7c30d245c0ecce3ef6c2a9ac76cd9ffaf6d0d090773b429c2b97437dc72047f46d9890abb2d6d8af7594ea19787e79ec80e388e2f6225b449c2e4d82e7df50f37c301424aede785935703c1c70235ba4e59f589f571218ce2dce4c061e598f0f38d1ac57f3feb52cf0c31078e3ceee8d796c53983fe1d38ebd71155aaa613700dc21f5a57941b787f921af7d287e72687d5242eb3063d543d5f5923f76db008cf4f56fb9c618f7f4bc8366ae70d88705617487754563e629119013fa0549e6645b397524b3dd2fa7e7f3fe9faf0fbf77da59f566861a3c510241fd4416ab7d0eba42d998e1178da51d607e0ef866607c458837c762323be53827d86e875
Content-Type: application/json
Content-Length: 119

{
   "grantType":"authorization_code",
   "authCode":"a6975f82-d00a-4ddc-9633-087fefb6275e",
   "refreshToken":"83a58570-6795-11ec-90d6-0242ac120003",
   "additionalInfo":{

   }
}

```

##### Sample Response

```json
{
   "responseCode":"2007400",
   "responseMessage":"Successful",
   "accessToken":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIyMTFlZThiMi1hN2FlLTRhZGUtYmJlYS1mNzI3MDk3ZmQ0NmEiLCJjbGllbnRJZCI6IjZhZTk1N2M0LTI4NjMtNDcxMy1hY2NlLWJhMTJkZTYzNmNmYyIsIm5iZiI6MTYxMTQ2ODk3OCwiZXhwIjoxNjExNDY5ODc4LCJpYXQiOjE2MTE0Njg5Nzh9.KM7yz9GvuUaDR1bXwei4iO0h4e3g4o1Hct5Ie9VoBdo",
   "tokenType":"Bearer",
   "accessTokenExpiryTime":"2020-01-01T00:00:00+07:00",
   "refreshToken":"57d21fe3-ba9c-4f2d-9fde-eae669bbf80d",
   "refreshTokenExpiryTime":"2020-01-01T00:00:00+07:00",
   "additionalInfo":{

   }
}

```

#### RESPONSES CODE

##### Response Code

Response status merupakan informasi yang diberikan oleh service provider kepada service consumer pada response body, sebagai indikasi hasil dari pemrosesan request yang diterima.

Response status terdiri dari 2 komponen, yaitu kode (response code) dan deskripsinya (response message).

KomponenTipe DataLengthKeteranganresponseCodeString7response code = HTTP status code + service code + case coderesponseMessageString150

###### **Daftar** _**Response Code**_

CategoryHTTP CodeService CodeCase CodeResponse MessageDescriptionSuccess200any00SuccessfulSuccessfulSuccess202any00Request In ProgressTransaction still on processSystem400any00Bad RequestGeneral request failed error, including message parsing failed.Message400any01Invalid Field Format {field name}Invalid formatMessage400any02Invalid Mandatory Field {field name}Missing or invalid format on mandatory fieldSystem401any00Unauthorized. \[reason\]General unauthorized error (No Interface Def, API is Invalid, Oauth Failed, Verify Client Secret Fail, Client Forbidden Access API, Unknown Client, Key not Found)System401any01Invalid Token (B2B)Token found in request is invalid (Access Token Not Exist, Access Token Expiry)System401any02Invalid Customer TokenToken found in request is invalid (Access Token Not Exist, Access Token Expiry)System401any03Token Not Found (B2B)Token not found in the system. This occurs on any API that requires token as input parameterSystem401any04Customer Token Not FoundToken not found in the system. This occurs on any API that requires token as input parameterBusiness403any00Transaction ExpiredTransaction expiredSystem403any01Feature Not Allowed \[Reason\]This merchant is not allowed to call Direct Debit APIsBusiness403any02Exceeds Transaction Amount LimitExceeds Transaction Amount LimitBusiness403any03Suspected FraudSuspected FraudBusiness403any04Activity Count Limit ExceededToo many request, Exceeds Transaction Frequency LimitBusiness403any05Do Not HonorAccount or User status is abnormalSystem403any06Feature Not Allowed At This Time. \[reason\]Cut off In ProgressBusiness403any07Card BlockedThe payment card is blockedBusiness403any08Card ExpiredThe payment card is expiredBusiness403any09Dormant AccountThe account is dormantBusiness403any10Need To Set Token LimitNeed to set token limitSystem403any11OTP BlockedOTP has been blockedSystem403any12OTP Lifetime ExpiredOTP has been expiredSystem403any13OTP Sent To Cardholerinitiates request OTP to the issuerBusiness403any14Insufficient FundsInsufficient FundsBusiness403any15Transaction Not Permitted.\[reason\]Transaction Not PermittedBusiness403any16Suspend TransactionSuspend TransactionBusiness403any17Token Limit ExceededPurchase amount exceeds the token limit set priorBusiness403any18Inactive Card/Account/CustomerIndicates inactive accountBusiness403any19Merchant BlacklistedMerchant is suspended from calling any APIsBusiness403any20Merchant Limit ExceedMerchant aggregated purchase amount on that day exceeds the agreed limitBusiness403any21Set Limit Not AllowedSet limit not allowed on particular tokenBusiness403any22Token Limit InvalidThe token limit desired by the merchant is not within the agreed range between the merchant and the IssuerBusiness403any23Account Limit ExceedAccount aggregated purchase amount on that day exceeds the agreed limitBusiness404any00Invalid Transaction StatusInvalid transaction statusBusiness404any01Transaction Not FoundTransaction not foundSystem404any02Invalid RoutingInvalid RoutingSystem404any03Bank Not Supported By SwitchBank not supported by switchBusiness404any04Transaction CancelledTransaction is cancelled by customerBusiness404any05Merchant Is Not Registered For Card Registration ServicesMerchant is not registered for Card Registration servicesSystem404any06Need To Request OTPNeed to request OTPSystem404any07Journey Not FoundThe journeyId cannot be found in the systemBusiness404any08Invalid MerchantMerchant does not exist or status abnormalBusiness404any09No IssuerNo issuerSystem404any10Invalid API TransitionInvalid API transition within a journeyBusiness404any11Invalid Card/Account/Customer \[info\]/Virtual AccountCard information may be invalid, or the card account may be blacklisted, or Virtual Account number maybe invalid.Business404any12Invalid Bill/Virtual Account \[Reason\]The bill is blocked/ suspended/not found.Virtual account is suspend/not found.Business404any13Invalid AmountThe amount doesn't match with what supposed toBusiness404any14Paid BillThe bill has been paidSystem404any15Invalid OTPOTP is incorrectBusiness404any16Partner Not FoundPartner number can't be foundBusiness404any17Invalid TerminalTerminal does not exist in the systemBusiness404any18Inconsistent RequestInconsistent request parameter found for the same partner reference number/transaction idIt can be considered as failed in transfer debit, but it should be considered as success in transfer credit.Considered as success:\- Transfer credit = (i) Intrabank transfer; (ii) Interbank transfer; (iii) RTGS transfer; (iv) SKNBI transfer;\- Virtual account = (i) Payment VA; (ii) Payment to VA;\- Transfer debit = (i) Refund payment; (ii) Void;Considered as failed:\- Transfer credit = (i) Transfer to OTC;\- Transfer debit = (i) Direct debit payment; (ii) QR CPM payment; (iii) Auth payment; (iv) Capture;Business404any19Invalid Bill/Virtual AccountThe bill is expired.Virtual account is expired.System405any00Requested Function Is Not SupportedRequested function is not supportedBusiness405any01Requested Opearation Is Not AllowedRequested operation to cancel/refund transaction Is not allowed at this time.System409any00ConflictCannot use same X-EXTERNAL-ID in same daySystem409any01Duplicate partnerReferenceNoTransaction has previously been processed indicates the same partnerReferenceNo already successSystem429any00Too Many RequestsMaximum transaction limit exceededSystem500any00General ErrorGeneral ErrorSystem500Any01Internal Server ErrorUnknown Internal Server Failure, Please retry the process againSystem500Any02External Server ErrorBackend system failure, etcSystem504any00Timeouttimeout from the issuer

#### APLIKASI PENGUJIAN

×

Akses Terbatas, Mohon Sign Up untuk Dapat Mengakses Halaman Ini

© 2025 - SNAP Developer Site