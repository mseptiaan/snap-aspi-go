Transfer Kredit - SNAP Developer Site

[![](/img/icon/snap-header-white-100.png)](/)

- [**Beranda**](/)
- [**Info**](/info)
- [**API**](/api-services)
- [**Direktori Publikasi**](/direktori-publikasi)
  - [**Sign In**](/sign-in)

[![](/img/icon/bi-logo-full-100.png)](/)

![](/img/card/transfer-debit-100.png)

### Transfer Debit

##### Transfer Debit

- [Direct Debit](/api-services/transfer-debit/direct-debit)
- [CPM](/api-services/transfer-debit/cpm)
- [Auth Payment](/api-services/transfer-debit/auth-payment)
- [Direct Debit BI-FAST](/api-services/transfer-debit/direct-debit-bi-fast)

#### **TRANSFER DEBIT**

API Transfer Debit digunakan untuk melakukan pemindahanbukuan dana dari satu rekening ke rekening lain baik dalam PJP AIS yang sama maupun PJP AIS yang berbeda yang diinisiasi oleh pihak penerima (pihak terkredit).

##### **SKENARIO PENGGUNAAN API TRANSFER DEBIT**

![](/img/docs/5_1_use_case_diagram_api_transfer_debit.png)

_Use Case Diagram_ API Transfer Debit

Merujuk pada _use case diagram_, pemilik rekening yaitu Konsumen, Non-PJP Pengguna Layanan, serta PJP PIAS, dapat melakukan pemindahan dana atas rekeningnya pada PJP AIS Bank menggunakan API Direct Debit dan API CPM. Sedangkan pemindahbukuan dana pada rekening PJP AIS Lembaga Selain Bank menggunakan API Direct Debit, API CPM, dan API Auth Payment.

API Transfer Debit dapat digunakan dalam sejumlah skema sebagai berikut:

**Skema 1: Konsumen - Non-PJP Pengguna Layanan / PJP PIAS - PJP AIS Lembaga Selain Bank**

Dalam skema ini, konsumen dapat melakukan pembayaran kepada Non-PJP Pengguna Layanan menggunakan dana yang berasal dari rekening Konsumen pada PJP AIS Lembaga Selain Bank, menggunakan API Direct Debit Payment atau API Auth Payment. Sebelum pendebitan dapat dilakukan, Konsumen melakukan account binding antara rekeningnya pada PJP AIS Lembaga Selain Bank dengan akun pada Non-PJP Pengguna Layanan (merujuk pada Bagian I untuk account binding).

- Untuk transaksi direct debit, Non-PJP Pengguna Layanan atau PJP PIAS mendapatkan persetujuan dari Konsumen untuk dilakukan pendebitan (proses binding) atas kartu/rekeningnya menggunakan API enable direct debit. Non-PJP Pengguna Layanan dapat menggunakan API disable direct debit apabila konsumen mencabut persetujuan pendebitan kartu/rekeningnya (proses unbinding).
- Untuk transaksi auth payment, Non-PJP Pengguna Layanan maupun PJP PIAS dapat meminta PJP AIS untuk menahan sejumlah nominal tertentu untuk keperluan pembelian barang dan/atau jasa yang nilainya belum ditentukan menggunakan API Auth Payment sesuai dengan otorisasi yang diberikan Konsumen. Pada penyelesaian transaksi, pendebitan dilakukan sejumlah nilai transaksi menggunakan API Auth Capture, sedangkan kelebihan nominal yang ditahan dikembalikan menggunakan API Auth Void. Apabila terdapat kendala pada transaksi tersebut, dana yang ditahan dapat dikembalikan menggunakan API Auth Refund.

**Skema 2: Konsumen - Non-PJP Pengguna Layanan / PJP PIAS - PJP AIS Bank**

Dalam skema ini, konsumen dapat melakukan pembayaran kepada Non-PJP Pengguna Layanan menggunakan dana yang berasal dari rekening Konsumen pada PJP AIS Bank, menggunakan API Direct Debit Payment. Sebelum pendebitan dapat dilakukan, Konsumen melakukan card registration atau account binding antara kartu atau rekeningnya pada PJP AIS Bank dengan akun pada Non-PJP Pengguna Layanan (merujuk pada Bagian I untuk card registration atau account binding).

- Untuk transaksi direct debit, Non-PJP Pengguna Layanan, atau PJP PIAS mendapatkan persetujuan dari Konsumen untuk dilakukan pendebitan (proses binding) atas kartu/rekeningnya menggunakan API enable direct debit. Non-PJP Pengguna Layanan dapat menggunakan API disable direct debit apabila konsumen mencabut persetujuan pendebitan kartu/rekeningnya (proses unbinding).

**Skema 3: Konsumen - Non-PJP Pengguna Layanan - PJP AIns/PJP PIAS - PIP (Switching) - PJP AIS Bank/PJP AIS Lembaga Selain Bank**

- Skema ini merupakan transaksi menggunakan QR CPM. Konsumen dapat melakukan pembayaran kepada Non-PJP Pengguna Layanan menggunakan dana yang berasal dari rekening konsumen pada PJP AIS Bank atau PJP AIS Lembaga Selain Bank, menggunakan API CPM Payment.
- Sebelum pendebitan dapat dilakukan, konsumen memberikan otorisasi kepada PJP AIS Bank atau PJP AIS Lembaga Selain Bank sebagai penatausaha rekening menggunakan API Binding.
- Transaksi pembayaran ini dilakukan menggunakan QR yang dihasilkan dari API Generate QR.
- Apabila terjadi gangguan atau kesalahan dalam bertransaksi, Non-PJP Pengguna Layanan dapat melakukan pengecekan status transaksi mengggunakan API Query Payment.
- Non-PJP Pengguna Layanan dapat membatalkan transaksi menggunakan API Cancel Payment atas permintaan Konsumen.

**Skema 4:** **Non-PJP Pengguna Layanan / PJP PIAS / PJP AIS Lembaga Selain Bank - PJP AIS Bank - PIP (BI-FAST)**

- Skema ini digunakan pada layanan Direct Debit BI-FAST yaitu pendebitan secara berkala rekening nasabah tertagih oleh _biller_, umumnya dalam rangka pembayaran atas penggunaan layanan (listrik, telepon, dll). Nasabah memberikan _consent_ pendebitan rekeningnya dalam bentuk _e-mandate_.
- Registrasi e-Mandate dilakukan nasabah tertagih, dengan memberikan informasi tagihan yang akan dibayar secara berkala. Informasi ini diteruskan ke bank penagih untuk dilakukan validasi terkait informasi _biller_.
- Direct Debit Transfer diinisiasi oleh biller dengan mengirimkan daftar billing kepada bank. Selanjutnya bank menginisiasi debit transfer ke bank nasabah melalui BI-Fast.
- Notify Direct Debit Transfer digunakan untuk menginformasikan status seluruh transfer individual yang terdapat dalam direct debit message.

© 2025 - SNAP Developer Site