Transfer Kredit - SNAP Developer Site

[![](/img/icon/snap-header-white-100.png)](/)

- [**Beranda**](/)
- [**Info**](/info)
- [**API**](/api-services)
- [**Direktori Publikasi**](/direktori-publikasi)
  - [**Sign In**](/sign-in)

[![](/img/icon/bi-logo-full-100.png)](/)

![](/img/card/transfer-credit-100.png)

### Transfer Kredit

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

##### **TRANSFER KREDIT**

API Transfer Kredit digunakan untuk melakukan pemindahbukuan dana dari satu rekening ke rekening lain baik dalam PJP AIS yang sama maupun PJP AIS yang berbeda yang diinisiasi oleh pihak pengirim (pihak terdebit).

###### **SKENARIO PENGGUNAAN API TRANSFER KREDIT**

![](/img/docs/4_1_use_case_diagram_api_transfer_credit.png)

_Use Case Diagram_ API Transfer Credit

Merujuk pada _use case diagram_, pemilik rekening yaitu Non-PJP Pengguna Layanan, atau PJP PIAS dapat melakukan pemindahanbukuan dana atas rekeningnya menggunakan:

- API internal account inquiry, external account inquiry, intrabank transfer, interbank transfer, transfer RTGS, transfer SKNBI, dan transaction status inquiry yang disediakan oleh PJP AIS Bank; dan
- API inquiry transaction, customer top up, bulk cashin, transfer to bank, transfer to OTC, CPM, MPM, transaction status inquiry, auth payment yang disediakan oleh PJP AIS Lembaga Selain Bank.

Pemilik rekening yaitu Konsumen, dapat melakukan pemindahanbukuan dana atas rekeningnya menggunakan:

- API Transfer to Bank, API Transfer to OTC, CPM, dan MPM yang disediakan oleh PJP AIS Lembaga Selain Bank.

API Transfer Kredit dapat digunakan dalam sejumlah skema sebagai berikut:

**Skema 1: Non-PJP Pengguna Layanan / PJP PIAS / PJP AIS Lembaga Selain Bank - PJP AIS Bank**

Dalam skema ini, Non-PJP Pengguna Layanan, PJP PIAS, atau PJP AIS Lembaga Selain Bank dapat melakukan pemindabukuan dana yang bersumber dari rekeningnya (skema B2B) melalui layanan transfer kredit pada PJP AIS Bank menggunakan API intrabank transfer, interbank transfer, request for payment, interbank transfer (bulk), transfer RTGS, dan transfer SKNBI.

Dalam skema ini, Non-PJP Pengguna Layanan dapat terhubung langsung ke PJP AIS maupun melalui PJP PIAS/PJP AIS Lembaga Selain Bank.

- Intrabank transfer:

API Intrabank Transfer digunakan untuk melakukan pemindahbukuan dana dari rekening Non-PJP Pengguna Layanan, PJP PIAS, atau PJP AIS Lembaga Selain Bank ke rekening tujuan pada satu PJP AIS Bank.

- Interbank Transfer:

API Interbank Transfer digunakan untuk melakukan pemindahbukuan dana dari satu rekening Non-PJP Pengguna Layanan, PJP PIAS, atau PJP AIS Lembaga Selain Bank, pada PJP AIS Bank ke rekening tujuan pada PJP AIS Bank lain melalui perantaraan Penyelenggara Infrastruktur Pembayaran (PIP) seperti GPN atau BI-FAST.

- Request for Payment:

API Request for Payment digunakan untuk melakukan penagihan oleh Non-PJP Pengguna Layanan, PJP PIAS, atau PJP AIS Lembaga Selain Bank. Apabila pihak tertagih menyetujui tagihan, proses pembayaran tagihan dilakukan menggunakan API Intrabank atau API Interbank.

- Interbank Transfer - Bulk:

API Interbank Transfer - Bulk digunakan untuk melakukan pemindahbukuan dana dari 1 (satu) rekening dengan tujuan lebih dari 1 (satu) rekening pada 1 (satu) atau lebih bank. Proses transaksi ini dapat diinisiasi oleh Non-PJP Pengguna Layanan atau PJP PIAS, dalam bentuk instruksi _bulk_. Transaksi kemudian diteruskan oleh PJP AIS Bank pengirim ke BI-FAST untuk dilakukan _de-bulking_ dan diteruskan ke PJP AIS Bank penerima dalam bentuk instruksi individual.

- Transfer RTGS:

API Transfer RTGS digunakan untuk melakukan pemindahbukuan dana dari satu rekening Non-PJP Pengguna Layanan, PJP PIAS, atau PJP AIS Lembaga Selain Bank, pada PJP AIS Bank ke rekening tujuan pada PJP AIS Bank lain melalui perantaraan Penyelenggara Infrastruktur Pembayaran (PIP) BI-RTGS.

- Transfer SKNBI:

API Transfer SKNBI digunakan untuk melakukan pemindahbukuan dana dari satu rekening Non-PJP Pengguna Layanan, PJP PIAS, atau PJP AIS Lembaga Selain Bank, pada PJP AIS Bank ke rekening tujuan pada PJP AIS Bank lain melalui perantaraan Penyelenggara Infrastruktur Pembayaran (PIP) SKNBI.

- Transfer _Virtual Account_:

Pembayaran menggunakan _virtual account_ pada dasarnya adalah melakukan transfer ke suatu nomor rekening yang bersifat _virtual_ dan ketika dilakukan pemindahbukuan ke dalam nomor rekening tersebut, akan langsung masuk ke nomor rekening asli yang terhubung ke _virtual account_ tersebut. Dengan memanfaatkan _virtual account_, Pengguna Layanan dapat mengidentifikasi tujuan pembayaran dari setiap transfer yang diterima. Informasi pembayaran beserta nomor _virtual account_ akan tercantum pada mutasi di rekening koran Pengguna Layanan.

Skenario penggunaan _virtual account_ terbagi menjadi:

- _Create_ VA

Pada skenario ini, nomor _virtual account_ dikelola oleh Penyedia Layanan. Permintaan pembuatan nomor _virtual account_ dilakukan oleh Non-PJP Pengguna Layanan kepada Penyedia Layanan melalui API.

- _Inquiry payment_

Pada skenario ini, nomor _virtual account_ dikelola oleh Non-PJP Pengguna Layanan. Setiap kali terdapat pembayaran maka Penyedia Layanan akan melakukan inquiry ke Non-PJP Pengguna Layanan untuk mendapatkan detail dari pembayaran/tagihan. Penyedia Layanan memberikan notifikasi secara _realtime_ kepada Non-PJP Pengguna Layanan melalui API setiap kali terjadi pembayaran.

- _Fund Transfer to_ VA

Pada skenario ini, Pengguna Layanan bekerja sama dengan Penyedia Layanan untuk menjadi antarmuka proses pembayaran _virtual account_ dari Pengguna Layanan lainnya.

Nomor _virtual account_ dapat dihasilkan untuk menerima satu pembayaran maupun digunakan secara berulang. Pembayaran menggunakan _virtual account_ dimungkinkan juga untuk dilakukan antar PJP melalui infrastruktur pembayaran (GPN atau BI-FAST).

Sebelum melakukan pemindahbukuan Intrabank, dilakukan validasi nomor dan nama rekening tujuan menggunakan API Internal Account Inquiry.

Sebelum melakukan pemindahbukuan Interbank, RTGS, atau SKNBI, dilakukan validasi nomor dan nama rekening tujuan menggunakan API External Account Inquiry.

Apabila terjadi gangguan pada sistem, time out, atau hal-hal lain yang dapat menyebabkan perbedaan interpretasi atas penyelesaian suatu transaksi, maka pihak yang menginisiasi transaksi (Non-PJP Pengguna Layanan, PJP PIAS, atau PJP AIS Lembaga Selain Bank) dapat mengakses API Transaction Status Inquiry yang disediakan oleh PJP AIS Bank.

**Skema 2: Non-PJP Pengguna Layanan - PJP AIS Lembaga Selain Bank**

Dalam skema ini, Non-PJP Pengguna Layanan dapat melakukan pemindahbukuan atas dana pada rekeningnya kepada rekening Konsumen melalui layanan transfer kredit pada PJP AIS Lembaga Selain Bank menggunakan API bulk cashin. PJP AIS Lembaga Selain Bank menginformasikan status transaksi bulk cashin ke Non-PJP Pengguna Layanan menggunakan API _notify bulk_.

**Skema 3: Konsumen - Non-PJP Pengguna Layanan/PJP PIAS - PJP AIS Lembaga Selain Bank**

Dalam skema ini, konsumen dapat melakukan top up saldo uang elektronik, transfer to OTC, dan transfer to bank, pada PJP AIS Lembaga Selain Bank melalui Non-PJP Pengguna Layanan menggunakan API Customer Top Up, Transfer to OTC, dan Transfer to Bank.

- Top up saldo uang elektronik

API customer top up digunakan untuk pemindabukuan dana dari rekening Non-PJP Pengguna Layanan ke rekening konsumen. Sebelum pemindahbukuan, Non-PJP Pengguna Layanan melakukan validasi rekening konsumen menggunakan API Account Inquiry.

- Transfer to OTC

API Transfer to OTC digunakan untuk pemindabukuan dana dari rekening Konsumen ke rekening Non-PJP Pengguna Layanan. Sebelum pemindahbukuan, Konsumen melakukan inisiasi transaksi OTC untuk mendapatkan token dari PJP AIS Lembaga Selain Bank pengelola rekeningnya.

- Transfer to bank

API Transfer to bank digunakan untuk pemindahbukuan dari rekening konsumen pada PJP AIS Lembaga Selain Bank ke rekening konsumen pada PJP AIS Bank, yang diinisiasi oleh konsumen melalui Non-PJP Pengguna Layanan. Sebelum pemindahbukuan, Non-PJP Pengguna Layanan melakukan validasi akun konsumen menggunakan API inquiry transaction.

Apabila terjadi gangguan pada sistem, time out, atau hal-hal lain yang dapat menyebabkan perbedaan interpretasi atas penyelesaian suatu transaksi, maka pihak yang menginisiasi transaksi (Non-PJP Pengguna Layanan, atau PJP PIAS) dapat mengakses API transaction status inquiry yang disediakan oleh PJP AIS Lembaga Selain Bank.

**Skema 4: QR MPM (Konsumen - Non-PJP Pengguna Layanan - PJP PIAS - PIP (Switching) - PJP AIS)**

Skema ini merupakan transaksi sesuai standar MPM. Konsumen dapat melakukan pembayaran kepada Non-PJP Pengguna Layanan menggunakan dana yang berasal dari rekening konsumen pada PJP AIS Bank atau PJP AIS Lembaga Selain Bank, menggunakan API MPM Payment.

Transaksi pembayaran ini dilakukan menggunakan QR yang dihasilkan dari API Generate QR dengan alternatif sebagai berikut:

- Pada mode redirection, konsumen akan diarahkan ke halaman web checkout Non-PJP Pengguna Layanan. PJP PIAS meminta PJP AIS Bank atau PJP AIS Lembaga Selain Bank untuk melakukan decode QR menggunakan API Decode QR. PJP PIAS melakukan request OTT untuk mendapatkan otorisasi menggunakan API Apply OTT, selanjutnya OTT disertakan pada redirect URL.
- Pada mode statis, konsumen diarahkan pada halaman web PJP AIS untuk memasukkan nominal.

Apabila terjadi kesalahan dalam transaksi, konsumen dapat melakukan pembatalan menggunakan API Cancel Payment.

Apabila terjadi gangguan pada sistem, time out, atau hal-hal lain yang dapat menyebabkan perbedaan interpretasi atas penyelesaian suatu transaksi, maka:

- Pihak yang menginisiasi transaksi mengkonfirmasi status transaksi dengan mengakses API Query Payment yang disediakan oleh PJP AIS.
- Pihak yang menyelesaikan transaksi (PJP AIS) menyampaikan status transaksi menggunakan API Payment Notify.

© 2025 - SNAP Developer Site