# ORM & Struktur Kode

## ORM
ORM adalah teknik pemrograman yang digunakan untuk mengubah data antara sistem dengan tipe yang tidak kompatibel, dengan menggunakan bahasa pemrograman berorientasi objek.

**Keuntungan:**
1. Mengurangi penggunaan kueri berulang.
2. Mengambil data secara otomatis ke dalam objek yang siap digunakan.
3. Menyederhanakan proses penyaringan data sebelum disimpan dalam database.
4. Beberapa ORM memiliki fitur caching untuk kueri.

**Kelemahan:**
1. Menambahkan lapisan tambahan dalam kode yang dapat menimbulkan overhead proses.
2. Memuat data hubungan yang mungkin tidak diperlukan.
3. Kueri SQL mentah yang kompleks bisa sulit diimplementasikan dengan ORM, terutama jika melibatkan lebih dari 10 tabel yang di-join.
4. Beberapa fungsi SQL spesifik dari satu vendor mungkin tidak didukung atau tidak memiliki fungsi tertentu, seperti contohnya dalam MySQL dengan FORCE INDEX.

### Migrasi Database
Migrasi database adalah cara untuk memperbarui versi database agar sesuai dengan perubahan versi aplikasi. Perubahan ini bisa berupa pembaruan ke versi terbaru atau pengembalian ke versi sebelumnya.

**Hubungan Database dalam GORM (Go Object Relational Mapping):**
1. Belongs to: Cocok untuk hubungan yang saling berhubungan.
2. Has one: Digunakan untuk merepresentasikan entitas tunggal yang hanya memiliki satu entitas terkait.
3. Has many: Digunakan untuk merepresentasikan entitas tunggal yang memiliki banyak entitas terkait.
4. Many to many: Digunakan untuk merepresentasikan hubungan banyak-ke-banyak antara banyak entitas. Karakteristik utama dari hubungan ini adalah penggunaan tabel pivot.

**Transaksi Database dalam GORM:**
- Transaksi adalah urutan operasi database. Jika salah satu operasi gagal, maka seluruh transaksi dianggap gagal.
- Transaksi berguna untuk memastikan integritas dan konsistensi data.
- GORM mendukung transaksi database.

**Cara Menginstal GORM:**
- Anda dapat melihat panduan instalasi GORM di [sini](https://gorm.io/docs/).
- Panduan migrasi database GORM dapat ditemukan [di sini](https://gorm.io/docs/migration.html).

**Cara Menghubungkan Aplikasi Go dengan Database:**
1. Buat fungsi `InitDB()` untuk melakukan koneksi ke database.
2. Buat skema pengguna menggunakan struct (objek) dan buat fungsi `InitialMigration()` untuk membuat struktur database di MySQL.
3. Panggil fungsi `InitDB()` dan `InitialMigration()` untuk menggunakannya.
4. Buat `GetUserController()` untuk mengambil data pengguna dari database (menggunakan ORM) dan menampilkannya dalam respons.
5. Buat `CreateUserController()` untuk mengikat data pengguna dari klien dan menyimpannya ke dalam database.
6. Tentukan rute API RESTful dengan menentukan penggunaan controller.

## Struktur Kode
Struktur proyek dapat diatur menggunakan pola Model-View-Controller (MVC).

**Alasan menggunakan struktur ini:**
- Untuk mencapai modularitas dalam aplikasi.
- Untuk menerapkan pemisahan tanggung jawab.
- Mengurangi konflik versi yang mungkin terjadi.
