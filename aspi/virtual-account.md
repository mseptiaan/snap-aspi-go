Virtual Account - SNAP Developer Site

[![](/img/icon/snap-header-white-100.png)](/)

- [**Beranda**](/)
- [**Info**](/info)
- [**API**](/api-services)
- [**Direktori Publikasi**](/direktori-publikasi)
  - [**Sign In**](/sign-in)

[![](/img/icon/bi-logo-full-100.png)](/)

![](/img/card/transfer-credit-100.png)

### [Transfer Kredit](/api-services/transfer-kredit)

##### Virtual Account

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

#### API Virtual Account

![](/img/docs/4_11_sequence_diagram_inquiry_payment_virtual_account.png)

Sequence Diagram Inquiry Payment - Virtual Account

![](/img/docs/4_12_sequence_diagram_create_virtual_account.png)

Sequence Diagram Create Virtual Account

![](/img/docs/4_13_fund_transfer_to_virtual_account.png)

Fund Transfer to Virtual Account

##### API Virtual Account - Inquiry

**Informasi Umum**

Service Code24NameAPI Virtual Account - InquiryVersion1.0HTTP MethodPOSTPath.../{version}/transfer-va/inquiry

##### API Virtual Account - Payment

**Informasi Umum**

Service Code25NameAPI Virtual Account - PaymentVersion1.0HTTP MethodPOSTPath.../{version}/transfer-va/payment

##### API Virtual Account - Inquiry Status

**Informasi Umum**

Service Code26NameAPI Virtual Account - Inquiry StatusVersion1.0HTTP MethodPOSTPath.../{version}/transfer-va/status

##### API Virtual Account - Create VA

**Informasi Umum**

Service Code27NameAPI Virtual Account - Create VAVersion1.0HTTP MethodPOSTPath.../{version}/transfer-va/create-va

##### API Virtual Account - Update VA

**Informasi Umum**

Service Code28NameAPI Virtual Account - Update VAVersion1.0HTTP MethodPUTPath.../{version}/transfer-va/update-va

##### API Virtual Account - Update Status VA

**Informasi Umum**

Service Code29NameAPI Virtual Account - Update Status VAVersion1.0HTTP MethodPUTPath.../{version}/transfer-va/update-status

##### API Virtual Account - Inquiry VA

**Informasi Umum**

Service Code30NameAPI Virtual Account - Inquiry VAVersion1.0HTTP MethodPOSTPath.../{version}/transfer-va/inquiry-va

##### API Virtual Account - Delete VA

**Informasi Umum**

Service Code31NameAPI Virtual Account - Delete VAVersion1.0HTTP MethodDELETEPath.../{version}/transfer-va/delete-va

##### API Virtual Account - Inquiry Payment to VA from Intra Bank

**Informasi Umum**

Service Code32NameAPI Virtual Account - Inquiry Payment to VA from Intra BankVersion1.0HTTP MethodPOSTPath.../{version}/transfer-va/inquiry-intrabank

##### API Virtual Account - Payment to VA from Intra Bank

**Informasi Umum**

Service Code33NameAPI Virtual Account - Payment to VA from Intra BankVersion1.0HTTP MethodPOSTPath.../{version}/transfer-va/payment-intrabank

##### API Virtual Account - Notification for Payment to VA from Intra Bank

**Informasi Umum**

Service Code34NameAPI Virtual Account - Notification for Payment to VA from Intra BankVersion1.0HTTP MethodPOSTPath.../{version}/transfer-va/notify-payment-intrabank

##### API Virtual Account - Get Report

**Informasi Umum**

Service Code35NameAPI Virtual Account - Get ReportVersion1.0HTTP MethodGETPath.../{version}/transfer-va/report

#### GUIDES

#### Spesifikasi Parameter Header dan Body API Virtual Account

##### API Virtual Account - Inquiry

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerServiceIdStringM8Derivative of X-PARTNER-ID, similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number(up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNotrxDateInitDateO25PJP internal system datetime with timezone, which follows the ISO-8601 standardchannelCodeNumberO4Channel code based on ISO 18245languageStringO2Language code based on ISO 639-1amountObjectOvalueStringM16,2Transaction Amount.Nominal inputted by Customer with 2 decimalcurrencyStringM3Currency (ISO4217)hashedSourceAccountNoStringC32Source account number in hashTo be filled without hash if there is a request from the sender/payer or ifthe consent from sender/payer has been granted. This is subject toArticle 8 paragraph 5 of Law No. 3 of 2011 concerning Fund Transfers. Alsocheck for other provisions, such as the PPATK regulation.sourceBankCodeStringC11Source account bank code To be filled if there is a request from the sender or if the consent from sender has been granted.This is subject to Article 8 paragraph 5 of Law No. 3 of 2011 concerning Fund Transfers. Also checkfor other provisions, such as the PPATK regulation.passAppStringO64Key for 3rd party to access API like client secretinquiryRequestIdStringM128Unique identifier for this Inquiry.Generated by PJP.additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response CoderesponseMessageStringM150Response DescriptionvirtualAccountDataObjectMinquiryStatusStringO2Status of inquiryinquiryReasonObjectOReason for Inquiry Status multi languageenglishStringO64Reason for Inquiry Status in EnglishindonesiaStringO64Reason for Inquiry Status in BahasapartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number(up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNoStringvirtualAccountNameStringM255Customer namevirtualAccountEmailStringO255Customer emailvirtualAccountPhoneStringO30Customer's phone numberFormat: 62xxxxxxxxxxxxxinquiryRequestIdStringM128From Inquiry RequesttotalAmountObjectOvalueStringM16,2Transaction Amount.Total Amount with 2 decimalcurrencyStringM3Currency (ISO4217)subCompanyStringO5Sub Company code generated by PartnerbillDetailsArray of ObjectsOArray with maximum 24 ObjectsbillCodeStringO2Bill code for Customer choosebillNoStringO18Bill number from PartnerbillNameStringO20Bill NamebillShortNameStringO10Bill Name to shown tobillDescriptionObjectOBill DescriptionenglishStringO18Bill Description in EnglishindonesiaStringO18Bill Description in BahasabillSubCompanyStringC5Partner's product code.Mandatory if subCompany sentbillAmountObjectOvalueStringM16,2Transaction AmountCurrencyStringM3Currency (ISO4217)billAmountLabelStringO25Label for billAmountbillAmountValueStringO25Value that will be shown for billAmountadditionalInfoObjectOunlimitedAdditional Information for custom use for each billfreeTextsArray of ObjectsOArray with maximum 25 ObjectsenglishStringO32Will be shown in ChannelindonesiaStringO32Will be shown in ChannelvirtualAccountTrxTypeStringO1Type of Virtual AccountC = Closed PaymentO = Open PaymentI = PartialM = Minimum - only can be paid once with minimum amountL = MaximumN = Open Minimum - can be paid multiple with minimum amountX = Open Maximum - can be paid multiple withfeeAmountObjectOvalueStringM16,2Transaction Amount.Nominal inputted by Customer with 2 decimalcurrencyStringM3Currency (ISO4217)additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Virtual Account - Payment

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number(up to 20 digits) . partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNameStringO255Customer namevirtualAccountEmailStringO255Customer emailvirtualAccountPhoneStringO30Customer's phone numberFormat: 62xxxxxxxxxxxxxtrxIdStringC64Unique identifier generated by PartnerMandatory if Payment comes from the Create VA RequestpaymentRequestIdStringM128Unique identifier generated by PJP.If Payment comes from the Inquiry process, this value must be the same with inquiryRequestId.channelCodeNumberO4Channel code based on ISO 18245hashedSourceAccountNoStringC32Source account number in hashTo be filled without hash if there is a request from the sender/payer or ifthe consent from sender/payer has been granted. This is subject toArticle 8 paragraph 5 of Law No. 3 of 2011 concerning Fund Transfers. Alsocheck for other provisions, such as the PPATK regulation.sourceBankCodeStringC11Source account bank code To be filled if there is a request from the sender or if the consent from sender has been granted.This is subject to Article 8 paragraph 5 of Law No. 3 of 2011 concerning Fund Transfers. Also checkfor other provisions, such as the PPATK regulation.paidAmountObjectMvalueStringM16,2Paid Amount with 2 decimalcurrencyStringM3Currency (ISO4217)cumulativePaymentAmountObjectOvalueStringM16,2Transaction AmountcurrencyStringM3Currency (ISO4217)paidBillsStringO6Hexadecimal format of binary of flag of paid billstotalAmountObjectOvalueStringM16,2Transaction Amount.Total amount from Inquiry with 2 decimalcurrencyStringM3Currency (ISO4217)trxDateTimeDateO25PJP internal system datetime with timezone, which follows the ISO-8601 standardreferenceNoStringO64Payment auth code generated by PJPjournalNumStringO6Sequence journal number in PJP Core SystempaymentTypeStringO1Type of paymentflagAdviseStringO1Status is this a retry notificationsubCompanyStringO5Sub Company code generated by PartnerbillDetailsArray of ObjectsOArray with maximum 24 ObjectsbillCodeStringO2From Inquiry ResponsebillNoStringO18From Inquiry ResponsebillNameStringO20From Inquiry ResponsebillShortNameStringO10From Inquiry ResponsebillDescriptionObjectOFrom Inquiry ResponseenglishStringO18From Inquiry ResponseindonesiaStringO18From Inquiry ResponsebillSubCompanyStringO5From Inquiry ResponsebillAmountObjectOvalueStringM16,2Transaction Amount.From Inquiry ResponsecurrencyStringM3Currency (ISO4217)additionalInfoObjectOunlimitedFrom Inquiry ResponsebillReferenceNoNumberO15Bill auth code generated by PJPfreeTextsArray of ObjectsOArray with maximum 25 ObjectsenglishStringO32Will be shown in ChannelindonesiaStringO32Will be shown in ChanneladditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response CoderesponseMessageStringM150Response DescriptionvirtualAccountDataObjectMpaymentFlagReasonObjectOReason for Payment Status multi languageindonesiaStringO200Reason for Payment Status in EnglishenglishStringO200Reason for inquiryStatus in BahasapartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number(up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNameStringM255Customer namevirtualAccountEmailStringO255Customer emailvirtualAccountPhoneStringO30Customer's phone numberFormat: 62xxxxxxxxxxxxxtrxIdStringO32From Payment RequestpaymentRequestIdStringM128From Payment RequestpaidAmountObjectOvalueStringM16,2Transaction Amount.From Payment RequestcurrencyStringM3Currency (ISO4217)paidBillsStringO6From Payment RequesttotalAmountObjectOvalueStringM16,2Transaction Amount.From Payment RequestcurrencyStringM3Currency (ISO4217)trxDateTimeDateO25From Payment RequestreferenceNoStringO15From Payment RequestjournalNumStringO6From Payment RequestpaymentTypeStringO1From Payment RequestflagAdviseStringO1From Payment RequestpaymentFlagStatusStringO2Status for Payment Flag from PartnerbillDetailsArray of ObjectsOArray with maximum 24 ObjectsbillerReferenceIdO64From Inquiry ResponsebillCodeStringO2From Inquiry ResponsebillNoStringO18From Inquiry ResponsebillNameStringO20From Inquiry ResponsebillShortNameStringO10From Inquiry ResponsebillDescriptionObjectOFrom Inquiry ResponseenglishStringO18From Inquiry ResponseindonesiaStringO18From Inquiry ResponsebillSubCompanyStringO5From Inquiry ResponsebillAmountObjectOvalueStringM16,2Transaction Amount.From Inquiry ResponsecurrencyStringM3Currency (ISO4217)additionalInfoObjectOunlimitedFrom Inquiry ResponsestatusStringO2Payment status for specific BillreasonObjectOReason for Payment Status for specific Bill multi languageenglishStringO64Reason for Payment Status for specific Bill in EnglishindonesiaStringO64Reason for Payment Status for specific Bill in BahasafreeTextsArray of ObjectsOArray with maximum 25 ObjectsenglishStringO32Will be shown in ChannelindonesiaStringO32Will be shown in ChanneladditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Virtual Account - Inquiry Status

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNoinquiryRequestIdStringC128Unique identifier from Inquiry.For use case Bill stored in Partner.If not send, will return array of transaction based on virtualAccountNopaymentRequestIdStringO128Unique identifier from Payment was generated by PJP.additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response CoderesponseMessageStringM150Response DescriptionvirtualAccountDataObjectMpaymentFlagReasonObjectOReason for Payment Status multi languageenglishStringO200Reason for Payment Status in EnglishindonesiaStringO200Reason for inquiryStatus in BahasapartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number(up to 20 digits)virtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits)inquiryRequestIdStringM128Unique identifier from InquirypaymentRequestIdStringC128Unique identifier for this Payment from PJP. Mandatory if Payment happened.paidAmountObjectOvalueStringM16,2Paid Amount with 2 decimalcurrencyStringM3CurrencyCurrency of amount based on ISO 4217paidBillsStringO6Hexadecimal format of binary of flag of paid billstotalAmountObjectOvalueStringM16,2Transaction Amount.Total amount from Inquiry with 2 decimalcurrencyStringM3Currency (ISO4217)trxDateTimeDateO25PJP internal system datetime with timezone, which follows the ISO-8601 standardtransactionDateDateO25Payment datetime when the payment happenedreferenceNoStringO15Payment auth code generated by PJPpaymentTypeStringO1Type of paymentflagAdviseStringO1Status is this a retry notificationpaymentFlagStatusStringO2Status for Payment FlagbillDetailsArray of ObjectsOArray with maximum 24 ObjectsbillCodeStringO2Bill code for Customer choosebillNoStringO18Bill number from PartnerbillNameStringO20Bill NamebillShortNameStringO10Bill Name to shown tobillDescriptionObjectOBill DescriptionenglishStringO18Bill Description in EnglishindonesiaStringO18Bill Description in BahasabillSubCompanyStringO5Partner's product codebillAmountObjectOvalueStringM16,2Transaction Amount.Nominal inputted by Customer with 2 decimalcurrencyStringM3Currency (ISO4217)additionalInfoObjectOunlimitedAdditional Information for custom use for each billbillReferenceNoNumberO15Bill auth code generated by PJPstatusStringO2Payment status for specific BillreasonObjectO2Reason for Payment Status for specific Bill multi languageenglishStringO64Reason for Payment Status for specific Bill in EnglishindonesiaStringO64Reason for Payment Status for specific Bill in BahasafreeTextsArray of ObjectsOArray with maximum 25 ObjectsenglishStringO32Will be shown in ChannelindonesiaStringO32Will be shown in ChanneladditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Virtual Account - Create VA

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerServiceIdStringO8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringO20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringO28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNameStringM255Customer namevirtualAccountEmailStringO255Customer emailvirtualAccountPhoneStringO30Customer's phone numberFormat: 62xxxxxxxxxxxxxtrxIdStringM64Transaction ID in Partner systemtotalAmountObjectOvalueStringM16,2Transaction Amount.Total Amount with 2 decimalcurrencyStringM3Currency (ISO4217)billDetailsArray of ObjectsOArray with maximum 24 ObjectsbillCodeStringO2Bill code for Customer choosebillNoStringO18Bill number from PartnerbillNameStringO20Bill NamebillShortNameStringO10Bill Name to shown tobillDescriptionObjectOBill DescriptionenglishStringO18Bill Description in EnglishindonesiaStringO18Bill Description in BahasabillSubCompanyStringO5Partner's product codebillAmountObjectOvalueStringM16,2Transaction Amount.Nominal inputted by Customer with 2 decimalcurrencyStringM3Currency (ISO4217)additionalInfoObjectOunlimitedAdditional Information for custom use for each billfreeTextsArray of ObjectsOArray with maximum 25 ObjectsenglishStringO32Will be shown in ChannelindonesiaStringO32Will be shown in ChannelvirtualAccountTrxTypeStringO1Type of Virtual AccountC = Closed PaymentO = Open PaymentI = PartialM = Minimum - only can be paid once with minimum amountL = MaximumN = Open Minimum - can be paid multiple with minimum amountX = Open Maximum - can be paid multiple withfeeAmountArray of ObjectsOvalueStringM16,2Transaction AmountNominal inputted by Customer with 2 decimalcurrencyStringM3Currency (ISO4217)expiredDateStringO25Expiration date for Virtual Account.Format expired date : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response CoderesponseMessageStringM150Response DescriptionvirtualAccountDataObjectMpartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNameStringM255Customer namevirtualAccountEmailStringO255Customer emailvirtualAccountPhoneStringO30Customer's phone numberFormat: 62xxxxxxxxxxxxxtrxIdStringM32from Create VA RequesttotalAmountObjectOvalueStringM16,2Transaction Amount.Total Amount with 2 decimal.currencyStringM3Currency.Currency of amount based on ISO 4217.billDetailsArray of ObjectsOArray with maximum 24 ObjectsbillCodeStringO2Bill code for Customer choosebillNoStringO18Bill number from PartnerbillNameStringO20Bill NamebillShortNameStringO10Bill Name to shown tobillDescriptionObjectOBill DescriptionenglishStringO18Bill Description in EnglishindonesiaStringO18Bill Description in BahasabillSubCompanyStringO5Partner's product codebillAmountObjectOvalueStringM16,2Transaction Amount.Nominal inputted by Customer with 2 decimal.currencyStringM3Currency (ISO4217)additionalInfoObjectOunlimitedAdditional Information for custom use for each billfreeTextsArray of ObjectsOArray with maximum 25 ObjectsenglishStringO32Will be shown in ChannelindonesiaStringO32Will be shown in ChannelvirtualAccountTrxTypeStringO1Type of Virtual AccountC = Closed PaymentO = Open PaymentI = PartialM = Minimum - only can be paid once with minimum amountL = MaximumN = Open Minimum - can be paid multiple with minimum amountX = Open Maximum - can be paid multiple withfeeAmountObjectOvalueStringM16,2Transaction Amount.Nominal inputted by Customer with 2 decimal.currencyStringM3Currency (ISO4217)expiredDateStringO25Expiration date for Virtual Account.Format expired date : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Virtual Account - Update VA

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNameStringM255Customer namevirtualAccountEmailStringO255Customer emailvirtualAccountPhoneStringO30Customer's phone numberFormat: 62xxxxxxxxxxxxxtrxIdStringM64Transaction ID in Partner systemtotalAmountObjectOvalueStringM16,2Transaction Amount.Total Amount with 2 decimal. (ISO4217)currencyStringM3Currency.Currency of amount based on ISO 4217.billDetailsArray of ObjectsOArray with maximum 24 ObjectsbillCodeStringO2Bill code for Customer choosebillNoStringO18Bill number from PartnerbillNameStringO20Bill NamebillShortNameStringO10Bill Name to shown tobillDescriptionObjectOBill DescriptionenglishStringO18Bill Description in EnglishindonesiaStringO18Bill Description in BahasabillSubCompanyStringO5Partner's product codebillAmountObjectOvalueStringM16,2Transaction Amount.Nominal inputted by Customer with 2 decimal.currencyStringM3Currency (ISO4217)additionalInfoObjectOunlimitedAdditional Information for custom use for each billfreeTextsArray of ObjectsOArray with maximum 25 ObjectsenglishStringO32Will be shown in ChannelindonesiaStringO32Will be shown in ChannelvirtualAccountTrxTypeStringO1Type of Virtual AccountC = Closed PaymentO = Open PaymentI = PartialM = Minimum - only can be paid once with minimum amountL = MaximumN = Open Minimum - can be paid multiple with minimum amountX = Open Maximum - can be paid multiple withfeeAmountObjectOvalueStringM16,2Transaction Amount.Nominal inputted by Customer with 2 decimal.currencyStringM3Currency (ISO4217)expiredDateStringO25Expiration date for Virtual Account.Format expired date : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response CoderesponseMessageStringM150Response DescriptionvirtualAccountDataObjectMpartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNameStringM255Customer namevirtualAccountEmailStringO255Customer emailvirtualAccountPhoneStringO30Customer's phone numberFormat: 62xxxxxxxxxxxxxtrxIdStringM32Transaction ID in Partner systemtotalAmountObjectOvalueStringM16,2Transaction Amount.Bill Amount with 2 decimal.currencyStringM3Currency.Currency of amount based on ISO 4217.billDetailsArray of ObjectsOArray with maximum 24 ObjectsbillCodeStringO2Bill code for Customer choosebillNoStringO18Bill number from PartnerbillNameStringO20Bill NamebillShortNameStringO10Bill Name to shown tobillDescriptionObjectOBill DescriptionenglishStringO18Bill Description in EnglishindonesiaStringO18Bill Description in BahasabillSubCompanyStringO5Partner's product codebillAmountObjectOvalueStringM16,2Transaction Amount.Nominal inputted by Customer with 2 decimal.currencyStringM3Currency (ISO 4217)additionalInfoObjectOunlimitedAdditional Information for custom use for each billvirtualAccountTrxTypeStringO1Type of Virtual AccountC = Closed PaymentO = Open PaymentI = PartialM = Minimum - only can be paid once with minimum amountL = MaximumN = Open Minimum - can be paid multiple with minimum amountX = Open Maximum - can be paid multiple withfeeAmountObjectOvalueStringM16,2Transaction Amount.Nominal inputted by Customer with 2 decimal.currencyStringM3Currency (ISO4217)expiredDateStringO25Expiration date for Virtual Account.Format expired date : (ISO 8601) YYYY-MM-DDThh:mm:sslastUpdateDateStringO25Last update date for Virtual Account.Format last update : (ISO 8601) YYYY-MM-DDThh:mm:sspaymentDateStringO25Payment date for Virtual Account.Format payment date : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Virtual Account - Update Status VA

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNotrxIdStringM64Transaction ID in Partner systempaidStatusStringM1Y = PaidN = Not Paid

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response CoderesponseMessageStringM150Response DescriptionvirtualAccountDataObjectMpartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNameStringM255Customer namevirtualAccountEmailStringO255Customer emailvirtualAccountPhoneStringO30Customer's phone numberFormat: 62xxxxxxxxxxxxxtrxIdStringM64Transaction ID in Partner systemtotalAmountObjectOvalueStringM16,2Transaction Amount.Bill Amount with 2 decimal.currencyStringM3Currency.Currency of amount based on ISO 4217.virtualAccountTrxTypeStringO1Type of Virtual AccountC = Closed PaymentO = Open PaymentI = PartialM = Minimum - only can be paid once with minimum amountL = MaximumN = Open Minimum - can be paid multiple with minimum amountX = Open Maximum - can be paid multiple withfeeAmountObjectOvalueStringM16,2Transaction Amount.Nominal inputted by Customer with 2 decimal.currencyStringM3Currency (ISO4217)expiredDateStringO25Expiration date for Virtual Account.Format expired date : (ISO 8601) YYYY-MM-DDThh:mm:sslastUpdateDateStringO25Last update date for Virtual Account.Format last update : (ISO 8601) YYYY-MM-DDThh:mm:sspaymentDateStringO25Payment date for Virtual Account.Format payment date : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Virtual Account - Inquiry VA

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNotrxIdStringM64Transaction ID in Partner system

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response CoderesponseMessageStringM150Response DescriptionvirtualAccountDataObjectsMpartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNameStringM255Customer namevirtualAccountEmailStringO255Customer emailvirtualAccountPhoneStringO30Customer's phone numberFormat:62xxxxxxxxxxxxxtrxIdStringM32Transaction ID in Partner systemtotalAmountObjectOvalueStringM16,2Bill Amount with 2 decimal.currencyStringM3Currency.Currency of amount based on ISO 4217.billDetailsArray of ObjectsOArray with maximum 24 ObjectsbillCodeStringO2Bill code for Customer choosebillNoStringO18Bill number from PartnerbillNameStringO20Bill NamebillShortNameStringO10Bill Name to shown tobillDescriptionObjectOBill DescriptionenglishStringO18Bill Description in EnglishindonesiaStringO18Bill Description in BahasabillSubCompanyStringO5Partner's product codebillAmountObjectOvalueStringM16,2Customer with 2 decimal.currencyStringM3Currency (ISO4217)additionalInfoObjectOunlimitedAdditional Information for custom use for each billfreeTextsArray of ObjectsOArray with maximum 25 ObjectsenglishStringO32Will be shown in ChannelindonesiaStringO32Will be shown in ChannelvirtualAccountTrxTypeStringO1Type of Virtual AccountC = Closed PaymentO = Open PaymentI = PartialM = Minimum - only can be paid once with minimum amountL = MaximumN = Open Minimum - can be paid multiple with minimum amountX = Open Maximum - can be paid multiple withfeeAmountObjectOvalueStringM16,2Transaction Amount.Nominal inputted byCustomer with 2 decimal.currencyStringM3Currency (ISO4217)expiredDateStringO25Expiration date for Virtual Account.Format expired date : (ISO 8601) YYYY-MM-DDThh:mm:sslastUpdateDateStringO25Last update date for Virtual Account.Format last update : (ISO 8601) YYYY-MM-DDThh:mm:sspaymentDateStringO25Payment date for Virtual Account.Format payment date : (ISO 8601) YYYY-MM-DDThh:mm:ssadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Virtual Account - Delete VA

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNotrxIdStringO64Transaction ID in Partner systemadditionalInfoObjectOunlimitedAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response CoderesponseMessageStringM150Response DescriptionvirtualAccountDataObjectMpartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNotrxIdStringO12Transaction ID in Partner systemadditionalInfoObjectOunlimitedAdditional information for custom use that are not provided by SNAP

##### API Virtual Account - Inquiry Payment to VA from Intra Bank

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNopartnerReferenceNoStringO128Unique identifier for this Payment. Generated by Partner.customerNoStringM20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits)trxDateTimeDateO25PJP internal system datetime with timezone, which follows the ISO-8601 standardchannelCodeNumberO4Channel code based on ISO 18245languageStringO2Language code based on ISO 639-1amountObjectOvalueStringM16,2Nominal inputted by Customer with 2 decimal.currencyStringM3Currency (ISO4217)sourceAccountNoStringO32Source account numbersourceAccountTypeStringO1D = Current AccountS = Saving AccountadditionalInfoObjectOunlimitedAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response CoderesponseMessageStringM150Response DescriptionvirtualAccountDataObjectMinquiryStatusStringO2Status of inquiryinquiryReasonObjectOReason for Inquiry Status multi languageenglishStringO64Reason for Inquiry Status in EnglishindonesiaStringO64Reason for Inquiry Status in BahasapartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNopartnerReferenceNoStringO128Unique identifier for this Payment. Generated by Partner.customerNoStringM20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNameStringO255Customer namevirtualAccountEmailStringO255Customer emailvirtualAccountPhoneStringO30Customer's phone numberFormat: 62xxxxxxxxxxxxxsourceAccountNoStringO32Source account numbersourceAccountTypeStringO1D = Current AccountS = Saving AccountinquiryRequestIdStringO128Inquiry identifier for Inquiry. Generated by PJP.totalAmountObjectOvalueStringM16,2Total Amount with 2 decimal.currencyStringM3Currency (ISO4217)billDetailsArray of ObjectsOArray with maximum 24 ObjectsbillCodeStringO2Bill code for Customer choosebillNoStringO18Bill number from PartnerbillNameStringO20Bill NamebillShortNameStringO10Bill Name to shown tobillDescriptionObjectOBill DescriptionenglishStringO18Bill Description in EnglishindonesiaStringO18Bill Description in BahasabillSubCompanyStringO5Partner's product codebillAmountObjectOvalueStringM16,2Transaction Amount.Nominal inputted by Customer with 2 decimal.currencyStringM3Currency (ISO4217)billAmountLabelStringO25Label for billAmountbillAmountValueObjectO25Value that will be shown for billAmountadditionalInfoObjectOunlimitedAdditional Information for custom use for each billfreeTextsArray of ObjectsOArray with maximum 25 ObjectsenglishStringO32Will be shown in ChannelindonesiaStringO32Will be shown in ChannelvirtualAccountTrxTypeStringC1Type of Virtual AccountC = Closed PaymentO = Open PaymentI = PartialM = Minimum - only can be paid once with minimum amountL = MaximumN = Open Minimum - can be paid multiple with minimum amountX = Open Maximum - can be paid multiple withfeeAmountObjectOvalueStringM16,2Nominal inputted by Customer with 2 decimal.currencyStringM3Currency (ISO4217)productNameStringO30Product CategoryadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Virtual Account - Payment to VA from Intra Bank

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNoreferenceNoStringO15Trancation Identifier on Service Provider SystemvirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNameStringO255Customer namevirtualAccountEmailStringO255Customer emailvirtualAccountPhoneStringO30Customer's phone numberFormat: 62xxxxxxxxxxxxxsourceAccountNoStringO32Source account numbersourceAccountTypeStringO1D = Current AccountS = Saving AccountinquiryRequestIdStringO128From Inquiry RequestpartnerReferenceNoStringM128From Payment Request.paidAmountObjectMvalueStringM16,2Paid Amount with 2 decimal.currencyStringM3Currency.From Inquiry Response.cumulativePaymentAmountObjectOvalueStringM16,2Cumulative Amount from virtualAccountNo paid multiple times.currencyStringM3Currency (ISO4217)paidBillsStringO6Hexadecimal format of binary of flag of paid billstotalAmountObjectOvalueStringM16,2Total amount from Inquiry with 2 decimal.currencyStringM3Currency (ISO4217)trxDateTimeDateO25PJP internal system datetime with timezone, which follows the ISO-8601 standardjournalNumStringO6Sequence journal number in PJP Core SystempaymentTypeStringO1Type of paymentflagAdviseStringO1Status is this a retry notificationpaymentStatusStringO20Status of payment requestbillDetailsArray of ObjectsOArray with maximum 24 ObjectsbillCodeStringO2From Inquiry ResponsebillNoStringO18From Inquiry ResponsebillNameStringO20From Inquiry ResponsebillShortNameStringO10From Inquiry ResponsebillDescriptionObjectOFrom Inquiry ResponseenglishStringO18From Inquiry ResponseindonesiaStringO18From Inquiry ResponsebillSubCompanyStringO5From Inquiry ResponsebillAmountObjectOvalueStringM16,2From Inquiry Response.currencyStringM3Currency (ISO4217)additionalInfoObjectOunlimitedFrom Inquiry ResponsebillReferenceNoNumberO15Bill auth code generated by PJPfreeTextsArray of ObjectsOArray with maximum 25 ObjectsenglishStringO32Will be shown in ChannelindonesiaStringO32Will be shown in ChannelfeeAmountObjectOvalueStringM16,2Nominal inputted by Customer with 2 decimalcurrencyStringM3Currency (ISO4217)additionalInfoObjectOAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response CoderesponseMessageStringM150Response DescriptionvirtualAccountdataObjectMpaymentFlagReasonObjectOReason for Payment Status multi languageenglishStringO200Reason for Payment Status in EnglishindonesiaStringO200Reason for inquiryStatus in BahasapartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNameStringM255Customer namevirtualAccountEmailStringO255Customer emailvirtualAccountPhoneStringO30Customer's phone numberFormat: 62xxxxxxxxxxxxxsourceAccountNoStringO32Source account numbersourceAccountTypeStringO1D = Current AccountS = Saving AccountinquiryRequestIdStringO128From Payment RequestpaymentRequestIdStringO128Unique Identifier for Payment. Generated by PJP and has the value with generated paymentRequestId to Partner Bille.partnerReferenceNoStringM128From Payment Request.referenceNoStringO128Generated by PJP.paidAmountObjectOvalueStringM16,2Transaction Amount.From Payment Request.currencyStringM3Currency.From Payment Request.paidBillsStringO6From Payment RequesttotalAmountObjectOvalueStringM16,2From Payment Request.currencyStringM3Currency (ISO4217)trxDateTimeDateO25From Payment RequestreferenceNoStringO15From Payment RequestjournalNumStringO6Sequence journal number in PJP Core SystempaymentTypeStringO1From Payment RequestflagAdviseStringO1From Payment RequestbillDetailsArray of ObjectsOArray with maximum 24 ObjectsbillCodeStringO2From Inquiry ResponsebillNoStringO18From Inquiry ResponsebillNameStringO20From Inquiry ResponsebillShortNameStringO10From Inquiry ResponsebillDescriptionObjectOFrom Inquiry ResponseenglishStringO18From Inquiry ResponseindonesiaStringO18From Inquiry ResponsebillSubCompanyStringO5From Inquiry ResponsebillAmountObjectOvalueStringM16,2From Inquiry Response.currencyStringM3Currency (ISO4217)additionalInfoObjectOunlimitedFrom Inquiry ResponsestatusStringO2Payment status for specific BillreasonStringO2Reason for Payment Status for specific Bill multi languageindonesiaStringO64Reason for Payment Status for specific Bill in EnglishenglishStringO64Reason for Payment Status for specific Bill in BahasafreeTextsArray of ObjectsOArray with maximum 25 ObjectsenglishStringO32Will be shown in ChannelindonesiaStringO32Will be shown in ChannelfeeAmountObjectOvalueStringM16,2Nominal inputted by Customer with 2 decimal.currencyStringM3Currency (ISO4217)productNameStringO30Product CategoryadditionalInfoObjectOAdditional information for custom use that are not provided by SNAP

##### API Virtual Account - Notification for Payment to VA from Intra Bank

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerServiceIdStringNumberM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringNumberM20Unique number(up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNoinquiryRequestIdStringO128From Inquiry RequestpaymentRequestIdStringO128Unique identifier generated by PJP.If Payment comes from the Inquiry process, this value must be the same with inquiryRequestId.partnerReferenceNoStringO128Unique identifier for this Payment. Generated by Partner.trxDateTimeDateO25PJP internal system datetime with timezone, which follows the ISO-8601 standardpaymentStatusStringO20Status of payment requestpaymentFlagReasonObjectOReason for Payment Status multi languageenglishStringO200Reason for Payment Status in EnglishindonesiaStringO200Reason for inquiryStatus in BahasaadditionalInfoObjectOunlimitedAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response CoderesponseMessageStringM150Response DescriptionvirtualAccountdataObjectOpaymentFlagReasonObjectOReason for Payment Status multi languageenglishStringO200Reason for Payment Status in EnglishindonesiaStringO200Reason for inquiryStatus in BahasapartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNoinquiryRequestIdStringO128From Inquiry RequestpaymentRequestIdStringO128Unique identifier generated by PJP.If Payment comes from the Inquiry process, this value must be the same with inquiryRequestId.partnerReferenceNoStringO128Unique identifier for this Payment. Generated by Partner.trxDateTimeDateO25PJP internal system datetime with timezone, which follows the ISO-8601 standardpaymentStatusStringO20Status of payment requestadditionalInfoObjectOunlimitedAdditional information for custom use that are not provided by SNAP

##### API Virtual Account - Get Report

**Request Body**

ParameterData TypeMandatoryLengthDescriptionpartnerServiceIdNumberM8From Payment Request.Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNostartDateStringO10Start Date for Reportyyyy-MM-ddISO-8601startTimeStringO14Start Time for Report.HH:mmIf used, must send startDateIf startDate exists without startTime, default startTime = 00:00ISO-8601endDateStringO10End Date for Reportyyyy-MM-ddIf used, must send startDateISO-8601endTimeStringO14End Time for Report.HH:mmIf used, must send endDate.If endDate exists without endTime, default endTime = 23:59ISO-8601additionalInfoObjectOunlimitedAdditional information for custom use that are not provided by SNAP

**Response Body**

ParameterData TypeMandatoryLengthDescriptionresponseCodeStringM7Response CoderesponseMessageStringM150Response DescriptionvirtualAccountdataArray of ObjectsMpaymentFlagReasonObjectOReason for Payment Status multi languageenglishStringO200Reason for Payment Status in EnglishindonesiaStringO200Reason for inquiryStatus in BahasapartnerServiceIdStringM8Derivative of X-PARTNER-ID , similar to company code,8 digit left padding space. partnerServiceId + customerNo or virtualAccountNocustomerNoStringM20Unique number(up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNoStringM28partnerServiceId (8 digit left padding space) + customerNo (up to 20 digits). partnerServiceId + customerNo or virtualAccountNovirtualAccountNameStringM255Customer namevirtualAccountEmailStringO255Customer emailvirtualAccountPhoneStringO30Customer's phone numberFormat: 62xxxxxxxxxxxxxsourceAccountNoStringO32Source account numbersourceAccountTypeStringO1D = Current AccountS = Saving AccounttrxIdStringO64From Payment RequestinquiryRequestIdStringO128From Payment RequestpaymentRequestIdStringO128From Payment RequestpaidAmountObjectOvalueStringM16,2From Payment Request.currencyStringM3Currency (ISO4217)paidBillsStringO6From Payment RequesttotalAmountObjectOvalueStringM16,2From Payment Request.currencyStringM3Currency (ISO4217)trxDateTimeDateO25From Payment RequestreferenceNoStringO15From Payment RequestjournalNumStringO6Sequence journal number in PJP Core SystempaymentTypeStringO1From Payment RequestflagAdviseStringO1From Payment RequestbillDetailsArray of ObjectsOArray with maximum 24 ObjectsbillCodeStringO2From Inquiry ResponsebillNoStringO18From Inquiry ResponsebillNameStringO20From Inquiry ResponsebillShortNameStringO10From Inquiry ResponsebillDescriptionObjectOFrom Inquiry ResponseenglishStringO18From Inquiry ResponseindonesiaStringO18From Inquiry ResponsebillSubCompanyStringO5From Inquiry ResponsebillAmountObjectOvalueStringM16,2From Inquiry Response.currencyStringM3Currency (ISO4217)additionalInfoObjectOunlimitedFrom Inquiry ResponsestatusStringO2Payment status for specific BillreasonStringO2Reason for Payment Status for specific Bill multi language>>\> englishStringO64Reason for Payment Status for specific Bill in English>>\> indonesiaStringO64Reason for Payment Status for specific Bill in BahasafreeTextsArray of ObjectsOFrom Inquiry Response.Array with maximum 25 ObjectsenglishStringO32From Inquiry ResponseindonesiaStringO32From Inquiry ResponseadditionalInfoObjectOunlimitedAdditional information for custom use that are not provided by SNAP

#### CODE SNIPPETS

#### Code Snippets API Virtual Account

##### API Virtual Account - Inquiry

**Sample Request**

```http
POST .../v1.0/transfer-va/inquiry HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T14:56:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
X-ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.1617169
X-LONGITUDE: 106.6643946
CHANNEL-ID: 95221

{
   "partnerServiceId":"  88899",
   "customerNo":"12345678901234567890",
   "virtualAccountNo":"  08889912345678901234567890",
   "txnDateInit":"20201231T235959Z",
   "channelCode":6011,
   "language":"ID",
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "hashedSourceAccountNo":"abcdefghijklmnopqrstuvwxyz123456",
   "sourceBankCode":"008",
   "passApp":"abcdefghijklmnopqrstuvwxyz",
   "inquiryRequestId":"abcdef-123456-abcdef",
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
     "responseCode":"2002400",
     "responseMessage":"Success",
     "virtualAccountData":{
                          "inquiryStatus":"00",
                          "inquiryReason":{
                                             "english":"Success",
                                             "indonesia":"Sukses"
                                           },
                          "partnerServiceId":" 088899",
                          "customerNo":"12345678901234567890",
                          "virtualAccountNo":" 08889912345678901234567890",
                          "virtualAccountName":"Jokul Doe",
                          "virtualAccountEmail":"john@email.com",
                          "virtualAccountPhone":"6281828384858",
                          "inquiryRequestId":"abcdef-123456-abcdef",
                          "totalAmount":{
                                             "value":"12345678.00",
                                             "currency":"IDR"
                                        },
                          "subCompany":"12345",
                          "billDetails":[
                                         {
                                            "billCode":"01",
                                            "billNo":"123456789012345678",
                                            "billName":"Bill A for Jan",
                                            "billShortName":"Bill A",
                                            "billDescription":{
                                                                "english":"Maintenance",
                                                                "indonesia":"Pemeliharaan"
                                                              },
                                            "billSubCompany":"00001",
                                            "billAmount":{
                                                           "value":"12345678.00",
                                                           "currency":"IDR"
                                                         },
                                            "billAmountLabel":"Total Tagihan",
                                            "billAmountValue":"Rp. 50.000,-",
                                            "additionalInfo":{

                                            }
                                         }
                                        ],
                          "freeTexts":[
                                         {
                                            "english":"Free text",
                                            "indonesia":"Tulisan bebas"
                                         }
                                      ],
                          "virtualAccountTrxType":"C",
                          "feeAmount":{
                                         "value":"12345678.00",
                                         "currency":"IDR"
                                      }
                          },
     "additionalInfo":{
                         "deviceId":"12345679237",
                         "channel":"mobilephone"
                      }
}

```

##### API Virtual Account - Payment

**Sample Request**

```http
POST .../v1.0/transfer-va/payment HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T14:56:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
X-ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.1617169
X-LONGITUDE: 106.6643946
CHANNEL-ID: 95221

{
   "partnerServiceId":"  088899",
   "customerNo":"12345678901234567890",
   "virtualAccountNo":"  08889912345678901234567890",
   "virtualAccountName":"Jokul Doe",
   "virtualAccountEmail":"jokul@email.com",
   "virtualAccountPhone":"6281828384858",
   "trxId":"abcdefgh1234",
   "paymentRequestId":"abcdef-123456-abcdef",
   "channelCode":6011,
   "hashedSourceAccountNo":"abcdefghijklmnopqrstuvwxyz123456",
   "sourceBankCode":"008",
   "paidAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "cumulativePaymentAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "paidBills":"95000",
   "totalAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "trxDateTime":"20201231T235959Z",
   "referenceNo":"123456789012345",
   "journalNum":"123456",
   "paymentType":1,
   "flagAdvise":"Y",
   "subCompany":"12345",
   "billDetails":[
      {
         "billCode":"01",
         "billNo":"123456789012345678",
         "billName":"Bill A for Jan",
         "billShortName":"Bill A",
         "billDescription":{
            "english":"Maintenance",
            "indonesia":"Pemeliharaan"
         },
         "billSubCompany":"00001",
         "billAmount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "additionalInfo":{

         },
         "billReferenceNo":"123456789012345"
      }
   ],
   "freeTexts":[
      {
         "english":"Free text",
         "indonesia":"Tulisan bebas"
      }
   ],
   "additionalInfo":{

   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-21T14:36:19+07:00

{
   "responseCode":"2002500",
   "responseMessage":"Success",
   "virtualAccountData":{
      "paymentFlagReason":{
         "english":"Success",
         "indonesia":"Sukses"
      },
      "partnerServiceId":" 088899",
      "customerNo":"12345678901234567890",
      "virtualAccountNo":"  08889912345678901234567890",
      "virtualAccountName":"Jokul Doe",
      "virtualAccountEmail":"jokul@email.com",
      "virtualAccountPhone":"6281828384858",
      "trxId":"abcdefgh1234",
      "paymentRequestId":"abcdef-123456-abcdef",
      "paidAmount":{
         "value":"12345678.00",
         "currency":"IDR"
      },
      "paidBills":"95000",
      "totalAmount":{
         "value":"12345678.00",
         "currency":"IDR"
      },
      "trxDateTime":"20201231T235959Z",
      "referenceNo":"123456789012345",
      "journalNum":"123456",
      "paymentType":1,
      "flagAdvise":"Y",
      "paymentFlagStatus":"00",
      "billDetails":[
         {
            "billerReferenceId":"123456789012345678",
            "billCode":"01",
            "billNo":"123456789012345678",
            "billName":"Bill A for Jan",
            "billShortName":"Bill A",
            "billDescription":{
               "english":"Maintenance",
               "indonesia":"Pemeliharaan"
            },
            "billSubCompany":"00001",
            "billAmount":{
               "value":"12345678.00",
               "currency":"IDR"
            },
            "additionalInfo":{

            },
            "status":"00",
            "reason":{
               "english":"Success",
               "indonesia":"Sukses"
            }
         }
      ],
      "freeTexts":[
         {
            "english":"Free text",
            "indonesia":"Tulisan bebas"
         }
      ]
   },
   "additionalInfo":{

   }
}

```

##### API Virtual Account - Inquiry Status

**Sample Request**

```http
POST .../v1.0/transfer-va/status HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T14:56:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
X-ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.1617169
X-LONGITUDE: 106.6643946
CHANNEL-ID: 95221

{
   "partnerServiceId":"  088899",
   "customerNo":12345678901234567890,
   "virtualAccountNo":"  08889912345678901234567890",
   "inquiryRequestId":"abcdef-123456-abcdef",
   "paymentRequestId":"abcdef-123456-abcdef",
   "additionalInfo":{

   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-21T14:36:19+07:00

{
   "responseCode":"2002600",
   "responseMessage":"Success",
   "virtualAccountData":{
      "paymentFlagReason":{
         "english":"Success",
         "indonesia":"Sukses"
      },
      "partnerServiceId":"  088899",
      "customerNo":"12345678901234567890",
      "virtualAccountNo":"  08889912345678901234567890",
      "inquiryRequestId":"abcdef-123456-abcdef",
      "paymentRequestId":"abcdef-123456-abcdef",
      "paidAmount":{
          "value":"12345678.00",
          "currency":"IDR"
       },
      "paidBills":"95000",
      "totalAmount":{
          "value":"12345678.00",
          "currency":"IDR"
       },
      "trxDateTime":"20201231T235959Z",
      "transactionDate":"20201230T235959Z",
      "referenceNo":"123456789012345",
      "paymentType":1,
      "flagAdvise":"Y",
      "paymentFlagStatus":"00",
      "billDetails":[
         {
            "billCode":"01",
            "billNo":"123456789012345678",
            "billName":"Bill A for Jan",
            "billShortName":"Bill A",
            "billDescription":{
               "english":"Maintenance",
               "indonesia":"Pemeliharaan"
            },
            "billSubCompany":"00001",
            "billAmount":{
               "value":"12345678.00",
               "currency":"IDR"
            },
            "additionalInfo":{

            },
            "billReferenceNo":"123456789012345",
            "status":"00",
            "reason":{
               "english":"Success",
               "indonesia":"Sukses"
            }
         }
      ],
      "freeTexts":[
         {
            "english":"Free text",
            "indonesia":"Tulisan bebas"
         }
      ]
   },
   "additionalInfo":{

   }
}

```

##### API Virtual Account - Create VA

**Sample Request**

```http
POST .../v1.0/transfer-va/create-va HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T14:56:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
X-ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.1617169
X-LONGITUDE: 106.6643946
CHANNEL-ID: 95221

{
   "partnerServiceId":"  088899",
   "customerNo":"12345678901234567890",
   "virtualAccountNo":"  08889912345678901234567890",
   "virtualAccountName":"Jokul Doe",
   "virtualAccountEmail":"jokul@email.com",
   "virtualAccountPhone":"6281828384858",
   "trxId":"abcdefgh1234",
   "totalAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "billDetails":[
      {
         "billCode":"01",
         "billNo":"123456789012345678",
         "billName":"Bill A for Jan",
         "billShortName":"Bill A",
         "billDescription":{
            "english":"Maintenance",
            "indonesia":"Pemeliharaan"
         },
         "billSubCompany":"00001",
         "billAmount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "additionalInfo":{

         }
      }
   ],
   "freeTexts":[
      {
         "english":"Free text",
         "indonesia":"Tulisan bebas"
      }
   ],
   "virtualAccountTrxType":"C",
   "feeAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "expiredDate":"2020-12-31T23:59:59+07:00",
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
   "responseCode":"2002700",
   "responseMessage":"Success",
   "virtualAccountData":{
      "partnerServiceId":"  088899",
      "customerNo":"12345678901234567890",
      "virtualAccountNo":"  08889912345678901234567890",
      "virtualAccountName":"Jokul Doe",
      "virtualAccountEmail":"jokul@email.com",
      "virtualAccountPhone":"6281828384858",
      "trxId":"abcdefgh1234",
      "totalAmount":{
         "value":"12345678.00",
         "currency":"IDR"
      },
      "billDetails":[
         {
            "billCode":"01",
            "billNo":"123456789012345678",
            "billName":"Bill A for Jan",
            "billShortName":"Bill A",
            "billDescription":{
               "english":"Maintenance",
               "indonesia":"Pemeliharaan"
            },
            "billSubCompany":"00001",
            "billAmount":{
               "value":"12345678.00",
               "currency":"IDR"
            },
            "additionalInfo":{

            }
         }
      ],
      "freeTexts":[
         {
            "english":"Free text",
            "indonesia":"Tulisan bebas"
         }
      ],
      "virtualAccountTrxType":"C",
      "feeAmount":{
         "value":"12345678.00",
         "currency":"IDR"
      },
      "expiredDate":"2020-12-31T23:59:59+07:00",
      "additionalInfo":{
         "deviceId":"12345679237",
         "channel":"mobilephone"
      }
   }
}

```

##### API Virtual Account - Update VA

**Sample Request**

```http
PUT .../v1.0/transfer-va/update-va HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T14:56:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
X-ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.1617169
X-LONGITUDE: 106.6643946
CHANNEL-ID: 95221

{
   "partnerServiceId":"  088899",
   "customerNo":"12345678901234567890",
   "virtualAccountNo":"  08889912345678901234567890",
   "virtualAccountName":"Jokul Doe",
   "virtualAccountEmail":"jokul@email.com",
   "virtualAccountPhone":"6281828384858",
   "trxId":"abcdefgh1234",
   "totalAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "billDetails":[
      {
         "billCode":"01",
         "billNo":"123456789012345678",
         "billName":"Bill A for Jan",
         "billShortName":"Bill A",
         "billDescription":{
            "english":"Maintenance",
            "indonesia":"Pemeliharaan"
         },
         "billSubCompany":"00001",
         "billAmount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "additionalInfo":{

         }
      }
   ],
   "freeTexts":[
      {
         "english":"Free text",
         "indonesia":"Tulisan bebas"
      }
   ],
   "virtualAccountTrxType":"C",
   "feeAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "expiredDate":"2020-12-31T23:59:59+07:00",
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
   "responseCode":"2002800",
   "responseMessage":"Success",
   "virtualAccountData":{
      "partnerServiceId":"  088899",
      "customerNo":"12345678901234567890",
      "virtualAccountNo":"  08889912345678901234567890",
      "virtualAccountName":"Jokul Doe",
      "virtualAccountEmail":"jokul@email.com",
      "virtualAccountPhone":"6281828384858",
      "trxId":"abcdefgh1234",
      "totalAmount":{
         "value":"12345678.00",
         "currency":"IDR"
      },
      "billDetails":[
         {
            "billCode":"01",
            "billNo":"123456789012345678",
            "billName":"Bill A for Jan",
            "billShortName":"Bill A",
            "billDescription":{
               "english":"Maintenance",
               "indonesia":"Pemeliharaan"
            },
            "billSubCompany":"00001",
            "billAmount":{
               "value":"12345678.00",
               "currency":"IDR"
            },
            "additionalInfo":{

            }
         }
      ],
      "virtualAccountTrxType":"C",
      "feeAmount":{
         "value":"12345678.00",
         "currency":"IDR"
      },
      "expiredDate":"2020-12-31T23:59:59+07:00",
      "lastUpdateDate":"2020-12-31T23:59:59+07:00",
      "paymentDate":"2020-12-31T23:59:59+07:00",
      "additionalInfo":{
         "deviceId":"12345679237",
         "channel":"mobilephone"
      }
   }
}

```

##### API Virtual Account - Update Status VA

**Sample Request**

```http
PUT .../v1.0/transfer-va/update-status HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T14:56:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
X-ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.1617169
X-LONGITUDE: 106.6643946
CHANNEL-ID: 95221

{
   "partnerServiceId":"  088899",
   "customerNo":"12345678901234567890",
   "virtualAccountNo":"  08889912345678901234567890",
   "trxId":"abcdefgh1234",
   "paidStatus":"Y"
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-21T14:36:19+07:00

{
   "responseCode":"2002900",
   "responseMessage":"Success",
   "virtualAccountData":{
      "partnerServiceId":"  088899",
      "customerNo":"12345678901234567890",
      "virtualAccountNo":"  08889912345678901234567890",
      "virtualAccountName":"Jokul Doe",
      "virtualAccountEmail":"jokul@email.com",
      "virtualAccountPhone":"6281828384858",
      "trxId":"abcdefgh1234",
      "totalAmount":{
         "value":"12345678.00",
         "currency":"IDR"
      },
      "virtualAccountTrxType":"C",
      "feeAmount":{
         "value":"12345678.00",
         "currency":"IDR"
      },
      "expiredDate":"2020-12-31T23:59:59+07:00",
      "lastUpdateDate":"2020-12-31T23:59:59+07:00",
      "paymentDate":"2020-12-31T23:59:59+07:00",
      "additionalInfo":{
         "deviceId":"12345679237",
         "channel":"mobilephone"
      }
   }
}

```

##### API Virtual Account - Inquiry VA

**Sample Request**

```http
POST .../v1.0/transfer-va/inquiry-va HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T14:56:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
X-ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.1617169
X-LONGITUDE: 106.6643946
CHANNEL-ID: 95221

{
   "partnerServiceId":"  088899",
   "customerNo":"12345678901234567890",
   "virtualAccountNo":"  08889912345678901234567890",
   "trxId":"abcdefgh1234"
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-21T14:36:19+07:00

{
   "responseCode":2003000,
   "responseMessage":"Success",
   "virtualAccountData":{
      "partnerServiceId":"  088899",
      "customerNo":"12345678901234567890",
      "virtualAccountNo":"  08889912345678901234567890",
      "virtualAccountName":"Jokul Doe",
      "virtualAccountEmail":"jokul@email.com",
      "virtualAccountPhone":"6281828384858",
      "trxId":"abcdefgh1234",
      "totalAmount":{
         "value":"12345678.00",
         "currency":"IDR"
      },
      "billDetails":[
         {
            "billCode":"01",
            "billNo":"123456789012345678",
            "billName":"Bill A for Jan",
            "billShortName":"Bill A",
            "billDescription":{
               "english":"Maintenance",
               "indonesia":"Pemeliharaan"
            },
            "billSubCompany":"00001",
            "billAmount":{
               "value":"12345678.00",
               "currency":"IDR"
            },
            "additionalInfo":{

            }
         }
      ],
      "freeTexts":[
         {
            "english":"Free text",
            "indonesia":"Tulisan bebas"
         }
      ],
      "virtualAccountTrxType":"C",
      "feeAmount":{
         "value":"12345678.00",
         "currency":"IDR"
      },
      "expiredDate":"2020-12-31T23:59:59+07:00",
      "lastUpdateDate":"2020-12-31T23:59:59+07:00",
      "paymentDate":"2020-12-31T23:59:59+07:00",
      "additionalInfo":{
         "deviceId":"12345679237",
         "channel":"mobilephone"
      }
   }
}

```

##### API Virtual Account - Delete VA

**Sample Request**

```http
DELETE .../v1.0/transfer-va/delete-va HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T14:56:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
X-ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.1617169
X-LONGITUDE: 106.6643946
CHANNEL-ID: 95221

{
   "partnerServiceId":"  088899",
   "customerNo":"12345678901234567890",
   "virtualAccountNo":"  08889912345678901234567890",
   "trxId":"abcdefgh1234",
   "additionalInfo":{

   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-21T14:36:19+07:00

{
   "responseCode":"2003100",
   "responseMessage":"Success",
   "virtualAccountData":{
      "partnerServiceId":"  088899",
      "customerNo":"12345678901234567890",
      "virtualAccountNo":"  08889912345678901234567890",
      "trxId":"abcdefgh1234",
      "additionalInfo":{

      }
   }
}

```

##### API Virtual Account - Inquiry Payment to VA from Intra Bank

**Sample Request**

```http
POST .../v1.0/transfer-va/inquiry-intrabank HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T14:56:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
X-ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.1617169
X-LONGITUDE: 106.6643946
CHANNEL-ID: 95221

{
   "partnerServiceId":"  088899",
   "partnerReferenceNo":"abcdef-123456-abcdef",
   "customerNo":12345678901234567890,
   "virtualAccountNo":"  08889912345678901234567890",
   "trxDateTime":"2020-10-20T17:56:57",
   "channelCode":6011,
   "language":"ID",
   "amount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "sourceAccountNo":"1234567890",
   "sourceAccountType":"S",
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
   "responseCode":2003200,
   "responseMessage":"Success",
   "virtualAccountdata":{
      "inquiryStatus":"00",
      "inquiryReason":{
         "english":"Success",
         "indonesia":"Sukses"
      },
      "partnerServiceId":"  088899",
      "partnerReferenceNo":"abcdef-123456-abcdef",
      "customerNo":"12345678901234567890",
      "virtualAccountNo":"  08889912345678901234567890",
      "virtualAccountName":"Jokul Doe",
      "virtualAccountEmail":"john@email.com",
      "virtualAccountPhone":"6281828384858",
      "sourceAccountNo":"1234567890",
      "sourceAccountType":"S",
      "inquiryRequestId":"abcdef-123456-abcdef",
      "totalAmount":{
         "value":"12345678.00",
         "currency":"IDR"
      },
      "billDetails":[
         {
            "billCode":"01",
            "billNo":"123456789012345678",
            "billName":"Bill A for Jan",
            "billShortName":"Bill A",
            "billDescription":{
               "english":"Maintenance",
               "indonesia":"Pemeliharaan"
            },
            "billSubCompany":"00001",
            "billAmount":{
               "value":"12345678.00",
               "currency":"IDR"
            },
            "billAmountLabel":"Total Tagihan",
            "billAmountValue":"Rp. 50.000,-",
            "additionalInfo":{

            }
         }
      ],
      "freeTexts":[
         {
            "english":"Free text",
            "indonesia":"Tulisan bebas"
         }
      ],
      "virtualAccountTrxType":"C",
      "feeAmount":{
         "value":"12345678.00",
         "currency":"IDR"
      },
      "productName":"Pendidikan",
      "additionalInfo":{
         "deviceId":"12345679237",
         "channel":"mobilephone"
      }
   }
}

```

##### API Virtual Account - Payment to VA from Intra Bank

**Sample Request**

```http
POST .../v1.0/transfer-va/payment-intrabank HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T14:56:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
X-ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.1617169
X-LONGITUDE: 106.6643946
CHANNEL-ID: 95221

{
   "partnerServiceId":"  088899",
   "customerNo":12345678901234567890,
   "referenceNo":123456789012345,
   "virtualAccountNo":"  08889912345678901234567890",
   "virtualAccountName":"Jokul Doe",
   "virtualAccountEmail":"jokul@email.com",
   "virtualAccountPhone":"6281828384858",
   "sourceAccountNo":"1234567890",
   "sourceAccountType":"S",
   "inquiryRequestId":"abcdef-123456-abcdef",
   "partnerReferenceNo":"abcdef-123456-abcdef",
   "paidAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "cumulativePaymentAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "paidBills":"950000",
   "totalAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "trxDateTime":"2020-10-20T17:56:57",
   "journalNum":"123456",
   "paymentType":1,
   "flagAdvise":"Y",
   "paymentStatus":"In Progress",
   "billDetails":[
      {
         "billCode":"01",
         "billNo":"123456789012345678",
         "billName":"Bill A for Jan",
         "billShortName":"Bill A",
         "billDescription":{
            "english":"Maintenance",
            "indonesia":"Pemeliharaan"
         },
         "billSubCompany":"00001",
         "billAmount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "additionalInfo":{

         },
         "billReferenceNo":"123456789012345"
      }
   ],
   "freeTexts":[
      {
         "english":"Free text",
         "indonesia":"Tulisan bebas"
      }
   ],
   "feeAmount":{
      "value":"12345678.00",
      "currency":"IDR"
   },
   "additionalInfo":{

   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-21T14:36:19+07:00

{
   "responseCode":"2003300",
   "responseMessage":"Success",
   "virtualAccountdata":{
      "paymentFlagReason":{
         "english":"Success",
         "indonesia":"Sukses"
      },
      "partnerServiceId":"  088899",
      "customerNo":"12345678901234567890",
      "virtualAccountNo":"  08889912345678901234567890",
      "virtualAccountName":"Jokul Doe",
      "virtualAccountEmail":"jokul@email.com",
      "virtualAccountPhone":"6281828384858",
      "sourceAccountNo":"1234567890",
      "sourceAccountType":"S",
      "inquiryRequestId":"abcdef-123456-abcdef",
      "paymentRequestId":"abcdef-123456-abcdef",
      "partnerReferenceNo":"abcdef-123456-abcdef",
      "referenceNo":"abcdef-123456-abcdef",
      "paidAmount":{
         "value":"12345678.00",
         "currency":"IDR"
      },
      "paidBills":"95000",
      "totalAmount":{
         "value":"12345678.00",
         "currency":"IDR"
      },
      "trxDateTime":"2020-10-20T17:56:57",
      "referenceNo":"123456789012345",
      "journalNum":"123456",
      "paymentType":1,
      "flagAdvise":"Y",
      "billDetails":[
         {
            "billCode":"01",
            "billNo":"123456789012345678",
            "billName":"Bill A for Jan",
            "billShortName":"Bill A",
            "billDescription":{
               "english":"Maintenance",
               "indonesia":"Pemeliharaan"
            },
            "billSubCompany":"00001",
            "billAmount":{
               "value":"12345678.00",
               "currency":"IDR"
            },
            "additionalInfo":{

            },
            "status":"00",
            "reason":{
               "english":"Success",
               "indonesia":"Sukses"
            }
         }
      ],
      "freeTexts":[
         {
            "english":"Free text",
            "indonesia":"Tulisan bebas"
         }
      ],
      "feeAmount":{
         "value":"12345678.00",
         "currency":"IDR"
      },
      "productName":"Pendidikan"
   },
   "additionalInfo":{

   }
}

```

##### API Virtual Account - Notification for Payment to VA from Intra Bank

**Sample Request**

```http
POST .../v1.0/transfer-va/notify-payment-intrabank HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T14:56:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
X-ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.1617169
X-LONGITUDE: 106.6643946
CHANNEL-ID: 95221

{
   "partnerServiceId":"  088899",
   "customerNo":12345678901234567890,
   "virtualAccountNo":"  08889912345678901234567890",
   "inquiryRequestId":"abcdef-123456-abcdef",
   "paymentRequestId":"abcdef-123456-abcdef",
   "partnerReferenceNo":"abcdef-123456-abcdef",
   "trxDateTime":"2020-12-21T14:56:11+07:00",
   "paymentStatus":"Success",
   "paymentFlagReason":{
      "english":"Success",
      "indonesia":"Sukses"
   },
   "additionalInfo":{

   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-21T14:36:19+07:00
X-SIGNATURE: 85be8171923ac135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5

{
   "responseCode":2003400,
   "responseMessage":"Success",
   "virtualAccountdata":{
      "paymentFlagReason":{
         "english":"Success",
         "indonesia":"Sukses"
      },
      "partnerServiceId":"  088899",
      "customerNo":12345678901234567890,
      "virtualAccountNo":"  08889912345678901234567890",
      "inquiryRequestId":"abcdef-123456-abcdef",
      "paymentRequestId":"abcdef-123456-abcdef",
      "partnerReferenceNo":"abcdef-123456-abcdef",
      "trxDateTime":"2020-12-21T14:56:11+07:00",
      "paymentStatus":"In Progress",
      "additionalInfo":{

      }
   }
}

```

##### API Virtual Account - Get Report

**Sample Request**

```http
POST .../v1.0/transfer-va/report HTTP/1.2
Content-type: application/json
Authorization: Bearer gp9HjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
Authorization-Customer: Bearer fa8sjjEj813Y9JGoqwOeOPWbnt4CUpvIJbU1mMU4a11MNDZ7Sg5u9a"
X-TIMESTAMP: 2020-12-21T14:56:11+07:00
X-SIGNATURE: 85be817c55b2c135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5
X-ORIGIN: www.hostname.com
X-PARTNER-ID: 82150823919040624621823174737537
X-EXTERNAL-ID: 41807553358950093184162180797837
X-IP-ADDRESS: 172.24.281.24
X-DEVICE-ID: 09864ADCASA
X-LATITUDE: -6.1617169
X-LONGITUDE: 106.6643946
CHANNEL-ID: 95221

{
   "partnerServiceId":"  088899",
   "startDate":"2020-12-31",
   "startTime":"14:56:11+07:00",
   "endDate":"2021-12-31",
   "endTime":"14:56:11+07:00",
   "additionalInfo":{
   }
}

```

**Sample Response**

```http
Content-type: application/json
X-TIMESTAMP: 2020-12-21T14:36:19+07:00
X-SIGNATURE: 85be8171923ac135157c7e89f52499bf0c25ad6eeebe04a986e8c862561b19a5

{
   "responseCode":2003500,
   "responseMessage":"Success",
   "virtualAccountdata":[
      {
         "paymentFlagReason":{
            "english":"Success",
            "indonesia":"Sukses"
         },
         "partnerServiceId":"  088899",
         "customerNo":"12345678901234567890",
         "virtualAccountNo":"  08889912345678901234567890",
         "virtualAccountName":"Jokul Doe",
         "virtualAccountEmail":"jokul@email.com",
         "virtualAccountPhone":"6281828384858",
         "sourceAccountNo":"1234567890",
         "sourceAccountType":"S",
         "trxId":"abcdefgh1234",
         "inquiryRequestId":"abcdef-123456-abcdef",
         "paymentRequestId":"abcdef-123456-abcdef",
         "paidAmount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "paidBills":"95000",
         "totalAmount":{
            "value":"12345678.00",
            "currency":"IDR"
         },
         "trxDateTime":"20201231T235959Z",
         "referenceNo":"123456789012345",
         "journalNum":"123456",
         "paymentType":1,
         "flagAdvise":"Y",
         "billDetails":[
            {
               "billCode":"01",
               "billNo":"123456789012345678",
               "billName":"Bill A for Jan",
               "billShortName":"Bill A",
               "billDescription":{
                  "english":"Maintenance",
                  "indonesia":"Pemeliharaan"
               },
               "billSubCompany":"00001",
               "billAmount":{
                  "value":"12345678.00",
                  "currency":"IDR"
               },
               "additionalInfo":{

               },
               "status":"00",
               "reason":{
                  "english":"Success",
                  "indonesia":"Sukses"
               }
            }
         ],
         "freeTexts":[
            {
               "english":"Free text",
               "indonesia":"Tulisan bebas"
            }
         ],
         "additionalInfo":{

         }
      }
   ]
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