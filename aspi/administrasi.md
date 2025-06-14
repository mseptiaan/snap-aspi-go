Administrasi - SNAP Developer Site

[![](/img/icon/snap-header-white-100.png)](/)

- [**Beranda**](/)
- [**Info**](/info)
- [**API**](/api-services)
- [**Direktori Publikasi**](/direktori-publikasi)
  - [**Sign In**](/sign-in)

[![](/img/icon/bi-logo-full-100.png)](/)

![](/img/card/administrasi-100.png)

### Administrasi

Overview


Guides


#### OVERVIEW

##### **PENGELOLAAN AKSES API**

Dalam penyelenggaraan _Open_ API Pembayaran berbasis SNAP diperlukan pedoman untuk pemberian akses dan _credential_ dari Penyedia Layanan kepada Pengguna Layanan. Hal ini mempertimbangkan bahwa _Open_ API Pembayaran berbasis SNAP yang disediakan oleh Penyedia Layanan, tidak dapat diakses secara langsung oleh publik. Penyedia Layanan mensyaratkan proses administrasi sebelum calon Pengguna Layanan dapat mengakses _Open_ API Pembayaran berbasis SNAP yang diselenggarakan oleh Penyedia Layanan.

Pada tahap proses administrasi, Penyedia Layanan baik Penyelenggara Jasa Pembayaran (PJP) Bank maupun PJP Lembaga Selain Bank, memerlukan sejumlah data guna mengetahui identitas dari calon Pengguna Layanan sebagai dasar agar dapat memberikan _credential_ dan akses.

#### GUIDES

##### **BENTUK PENGAJUAN API**

Dalam pengajuan kerjasama _Open_ API Pembayaran berbasis SNAP, calon Pengguna Layanan disyaratkan untuk mengisi formulir yang berisi informasi calon Pengguna Layanan dan layanan API yang akan digunakan. Berdasarkan informasi tersebut, Penyedia Layanan akan melakukan pemrosesan. Setelah pengajuan kerjasama disetujui, Penyedia Layanan akan mendaftarkan Pengguna Layanan sebagai pihak yang bekerjasama.

Formulir pengajuan kerjasama _Open_ API Pembayaran berbasis SNAP dapat berbentuk formulir tertulis ( _physical registration_) maupun formulir daring ( _online registration_).

##### **Formulir Tertulis ( _Physical Registration_)**

Dalam hal Penyedia Layanan hanya menyediakan formulir pengajuan kerjasama _Open_ API Pembayaran berbasis SNAP berbentuk formulir tertulis, maka calon Pengguna Layanan mengisi dan melengkapi formulir sebagaimana dimaksud. Formulir yang telah dilengkapi oleh calon Pengguna Layanan, akan diproses oleh Penyedia Layanan dan diinput secara manual pada sistem Penyedia Layanan.

##### **Formulir Daring ( _Online Registration_)**

Dalam hal Penyedia Layanan hanya menyediakan formulir pengajuan kerjasama _Open_ API Pembayaran berbasis SNAP berbentuk formulir daring, maka calon Pengguna Layanan mengisi dan melengkapi formulir pengajuan kerjasama secara daring baik melalui _website_ atau aplikasi yang disediakan oleh Penyedia Layanan.

##### **PERSETUJUAN TERHADAP KETENTUAN KERJASAMA LAYANAN PEMBAYARAN BERBASIS API**

Dalam proses pengajuan kerjasama _Open_ API Pembayaran berbasis SNAP, terdapat persyaratan kerjasama yang disepakati oleh calon Pengguna Layanan dan Penyedia Layanan. Hal ini untuk mencegah apabila dalam pelaksanaan atau penggunaan _Open_ API Pembayaran berbasis SNAP terjadi kesalahpahaman atau kelalaian di antara kedua belah pihak. Calon Pengguna Layanan memberikan persetujuan terhadap persyaratan kerjasama tersebut dengan menandatangani pernyataan di atas materai agar pernyataan tersebut berkekuatan hukum. Pernyataan yang telah ditandatangani oleh calon Pengguna Layanan disampaikan kepada Penyedia Layanan sebagai bukti bahwa calon Pengguna Layanan telah setuju dengan persyaratan yang ditetapkan oleh Penyedia Layanan.

##### **JENIS PENANGGUNG JAWAB API**

Dalam setiap kerjasama _Open_ API Pembayaran berbasis SNAP ditunjuk penanggung jawab sebagai pihak berwenang yang akan mewakili pihak Pengguna Layanan yang akan bekerjasama. Penunjukkan penanggung jawab bertujuan untuk mempermudah komunikasi antar pihak Penyedia Layanan dengan pihak Pengguna Layanan seperti untuk mempermudah proses diskusi atau penyampaian masalah.

Berikut merupakan jenis-jenis penanggung jawab yang diminta oleh Penyedia Layanan kepada calon Pengguna Layanan pada saat proses pengajuan kerja sama:

1. _Project Manager_
2. Penanggung jawab API Public Key
3. Penanggung jawab API Private Key
4. Penanggung jawab API _password key_ dan API Private Key
5. Penanggung jawab transaksi API atau pertukaran data
6. Penanggung jawab integrasi API
7. Penanggung jawab administratif dan kerahasiaan dokumen
8. Penanggung jawab indikasi _fraud_

Penyedia Layanan menetapkan jumlah minimum dan maksimum untuk masing-masing penanggung jawab API yaitu minimum 1 (satu) orang dan maksimum 2 (dua) orang.

##### **DATA ADMINISTRASI PENDAFTARAN BAGI CALON PENGGUNA LAYANAN**

Data-data administrasi yang disyaratkan untuk pendaftaran kerjasama _Open_ API Pembayaran berbasis SNAP terdiri dari informasi Pengguna Layanan, informasi legalitas, dan informasi pejabat berwenang. Data administrasi sebagaimana dimaksud adalah sebagai berikut, namun tidak terbatas pada:

1. Informasi calon Pengguna Layanan
01. Nama perusahaan
02. Nama dagang / merek (jika ada)
03. Klasifikasi bisnis (deskripsi dan industri)\*
       1. Jenis usaha / kategori usaha
       2. Kriteria usaha\*
       3. MCC ( _Merchant Category Code_)
04. Alamat perusahaan
05. Alamat korespondensi termasuk email
06. Alamat usaha / operasional
07. _Website_\*
08. Logo\*
09. _National Merchant_ ID
10. Nomor Pokok Wajib Pajak (NPWP) usaha
11. Nomor Surat Izin Usaha Perdagangan (SIUP)
12. Informasi rekening (Nama dan No. Rekening, Nama Bank)
13. Informasi tambahan\*
       1. Kesepakatan Komersial (sesuai _product inventory_)\*
       2. Informasi fitur (sesuai _product inventory_)\*
       3. Informasi konfigurasi merchant (nama dan _email merchant_)\*

\\* Opsional

1. Informasi legalitas
01. Salinan dokumen dasar pendirian badan termasuk dengan perubahannya (Anggaran Dasar, Akta Pendirian)
02. Salinan dokumen pengesahan dan perijinan yang masih berlaku dari instansi terkait
03. Dokumen perjanjian kerja sama
04. Asli surat penunjukan dari perusahaan sebagai Pengguna Layanan
05. Salinan NPWP / bukti potong pajak
06. Salinan SIUP
07. Salinan Nomor Induk Berusaha (NIB)\*\*
08. Salinan surat keterangan domisili usaha
09. Salinan Tanda Daftar Perusahaan (TDP)
10. Salinan surat pengukuhan Pengusaha Kena Pajak
11. Salinan rekening koran atau buku tabungan
12. Asli surat kuasa pendebetan rekening
13. Nama pemilik usaha (salinan identitas diri terakhir)
14. Nama penanggung jawab usaha dan pengurus (salinan identitas diri terakhir)
15. Foto lokasi perusahaan\*
16. Dokumentasi hasil pengujian API di _Developer Site_ SNAP

\\* Opsional

\\*\\* NIB menggantikan copy dokumen SIUP dan TDP

1. Informasi penanggung jawab usaha PIC atau pejabat berwenang dalam kontrak:

1. Nama pemilik usaha / pejabat berwenang
2. Nomor kontak PIC atau pejabat berwenang
3. Salinan identitas diri terakhir yang masih berlaku dari para pengurus atau pejabat berwenang yang bertindak untuk dan atas nama pemohon

##### **DATA-DATA TEKNIS PENGAJUAN API**

Selain data-data terkait informasi calon Pengguna Layanan, Penyedia Layanan juga meminta informasi data teknis untuk digunakan dalam melakukan konfigurasi koneksi antara kedua belah pihak. Data-data ini diperlukan agar komunikasi antar kedua pihak dapat berlangsung.

Data-data yang diperlukan oleh pihak Penyedia Layanan pada saat pengajuan kerjasama Layanan Pembayaran Berbasis API adalah sebagai berikut:

01. Fitur API yang ingin dipakai
02. Informasi apakah Pengguna Layanan akan menggunakan jalur koneksi yang sudah tersedia saat ini antara Pengguna Layanan dan Penyedia Layanan
03. Jenis koneksi yang akan digunakan (internet/ _Multi Protocol Label Switching_)
04. Informasi _whitelist_ IP
05. Informasi pemakaian _switcher_ sebagai perantara koneksi
06. _Callback_ URL (alamat domain aplikasi Pengguna Layanan)
07. _Public key_ _for digital signature_ _(for validate X-Signature in Access Token)_
08. _Subscription Package_ (opsional)
09. Limit transaksi dan dokumen _underlying_
10. Jam operasional penggunaan API

##### **PERUBAHAN DATA-DATA ADMINISTRASI DAN TEKNIS API**

Mempertimbangkan bahwa hingga saat ini belum ada Penyedia Layanan yang telah memiliki API untuk layanan perubahan data Pengguna Layanan baik data administrasi maupun teknis, maka perubahan data dilakukan dengan cara mengisi dan mengirimkan formulir tertulis sebagaimana pada saat pengajuan kerjasama atau dengan mengubah sebagian data pada _website_ atau aplikasi yang disediakan oleh Penyedia Layanan. Dalam hal, data yang dapat diubah melalui aplikasi yang disediakan Penyedia Layanan terbatas, maka apabila dibutuhkan perubahan data lainnya Pengguna Layanan harus menghubungi pihak Penyedia Layanan untuk perubahan data yang tidak tersedia pada _website_ atau aplikasi sebagaimana dimaksud. Pihak Penyedia Layanan membantu Pengguna Layanan untuk melakukan perubahan data berdasarkan informasi dari Pengguna Layanan.

##### **PENGIRIMAN _KEY_**

Pada kerjasama _Open_ API Pembayaran berbasis SNAP, diperlukan adanya penggunaan _key_ yaitu API Public Key dan API Private Key antara Pengguna Layanan dengan Penyedia Layanan. Penggunaan _key_ bertujuan untuk menjaga keamanan data transaksi yang dikirimkan dan untuk memastikan bahwa pengirim data adalah benar berasal dari pihak yang dikenal oleh Penyedia Layanan.

Proses pengiriman _key_ kepada Pengguna Layanan perlu menjaga faktor kerahasiaan agar tidak disalahgunakan oleh pihak yang tidak berwenang.

_Key_ untuk kerjasama Layanan Pembayaran Berbasis API dikirimkan melalui media _email_ atau aplikasi yang disediakan oleh Penyedia Layanan.

##### **Metode Penerimaan API Public Key**

Penerimaan API Public Key dilakukan dengan pilihan metode sebagai berikut:

1. API Public Key diterima oleh PIC yang ditunjuk; atau
2. API Public Key diterima dan diproses oleh sistem aplikasi.

Pilihan mekanisme penerimaan tersebut dimaksudkan untuk meningkatkan keamanan terhadap API Public Key agar tidak dengan mudah disalahgunakan atau dibocorkan oleh pihak tidak berwenang.

##### **Metode pengiriman API Public Key**

Selain menentukan metode penerimaan API Public Key, peningkatan pengamanan pada pengiriman API Public Key dilakukan melalui cara penggunaan enkripsi atau _password_ untuk melindungi keamanan data. API Public Key dienkripsi atau dilindungi menggunakan _password_ oleh Penyedia Layanan sebagai pihak pengirim. Selanjutnya, Pengguna Layanan sebagai pihak penerima perlu melakukan dekripsi atau menggunakan _password_ yang diketahuinya untuk melihat atau membuka data API Public Key.

Pengiriman API Public Key dilakukan dengan pilihan metode sebagai berikut:

1. API Public Key dikirimkan melalui _email_ dalam bentuk _zip file_ yang diberi _password_;
2. API Public Key dikirimkan melalui _email_ dalam bentuk _file_ terenkripsi menggunakan _PGP Encryption_;
3. API Public Key dikirimkan melalui _email_ dalam bentuk _file_ terenkripsi menggunakan _OpenSSL_ dan _public key partner_; atau
4. API Public Key dikirimkan secara _online_ (melalui aplikasi).

Mekanisme pengiriman API Public Keydan API sebagaimana dimaksud dilakukan dengan pilihan metode sebagai berikut:

1. Pengiriman dipisahkan dalam tiga _email_ berbeda, yaitu:

1. email 1 : Zip file berisi API Public Key
2. email 2 : password zip untuk membuka _file_ lampiran zip pada email 1
2. Pengiriman dipisahkan dalam dua _email_ berbeda, yaitu:

1. email 1 : Zip file berisi API Public Key
2. email 2 : password zip untuk membuka _file_ lampiran zip pada email 1
3. Pengiriman secara _online_ (melalui aplikasi)

##### **PENGELOLAAN _KEY_**

Pengelolaan _key_ baik API Public Key dan API Private Key oleh Penyedia Layanan maupun Pengguna Layanan dilaksanakan berdasarkan asas kerahasiaan. Pengelolaan _key_ sebagaimana dimaksud sekurang-kurangnya meliputi namun tidak terbatas pada:

1. Memiliki prosedur hak akses pengelolaan _key_.
2. Memiliki _database_ yang aman sebagai tempat penyimpanan _key_ dan hanya dapat diakses oleh pihak yang berwenang.
3. Memiliki prosedur pembuatan, pembaharuan, penghapusan, dan enkripsi/dekripsi _key_.
4. Menggunakan algoritma yang direkomendasikan lembaga standar internasional ataupun nasional dalam pembuatan _key._
5. Terdapat kejelasan versi _master key_ yang digunakan.
6. Memungkinkan dilakukan monitoring dan audit atas penggunaan _key_.

© 2025 - SNAP Developer Site