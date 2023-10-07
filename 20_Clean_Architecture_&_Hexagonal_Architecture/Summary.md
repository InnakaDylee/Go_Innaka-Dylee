# Ringkasan

## Clean Architecture and Hexagonal Architecture

### Clean Architecture
* Arsitektur Bersih adalah pendekatan desain perangkat lunak yang menekankan pemisahan berbagai lapisan dalam sebuah aplikasi. Tujuan dari implementasi Arsitektur Bersih adalah membuat perangkat lunak yang:
    * Tidak bergantung pada Kerangka Kerja (Frameworks)
    * Dapat diuji (Testable)
    * Tidak bergantung pada Antarmuka Pengguna (UI)
    * Tidak bergantung pada Basis Data (Database)
    * Tidak bergantung pada Ketergantungan Eksternal apa pun

* Arsitektur Bersih terdiri dari empat lapisan utama: Entitas (Entities), Kasus Pengguna (Use Cases), Antarmuka Pengguna (User Interface), dan Sumber Data (Data Sources).
    * Entities:
        * Ini adalah lapisan inti yang berisi definisi entitas bisnis dan logika bisnis murni.
        * Lapisan ini tidak boleh bergantung pada lapisan lain dan harus tetap stabil.
    * Use Cases:
        * Lapisan ini berisi implementasi kasus pengguna atau alur kerja yang mewakili operasi bisnis.
        * Ini memproses input dari antarmuka pengguna dan mengembalikan hasilnya.
    * User Interface:
        * Lapisan ini mengelola interaksi antara pengguna dan aplikasi, termasuk antarmuka pengguna (UI) atau antarmuka program (API).
        * Lapisan ini mengambil input dari pengguna, meneruskannya ke lapisan kasus pengguna, dan menampilkan hasilnya.
    * Data Sources:
        * Lapisan ini menangani akses ke sumber data eksternal seperti basis data, layanan web, atau penyimpanan berkas.
        * Lapisan ini menghubungkan aplikasi dengan sumber data eksternal.

* Arsitektur Bersih bertujuan menciptakan perangkat lunak yang lebih mudah dimengerti, dipelihara, dan dioptimalkan dengan menghindari ketergantungan yang salah dan menerapkan prinsip*prinsip desain yang baik.

### Hexagonal Architecture
* Arsitektur Hexagonal adalah pendekatan desain perangkat lunak yang bertujuan menciptakan aplikasi yang mudah dipelihara, tidak tergantung pada detail teknis, dan mudah diuji. Dalam Arsitektur Hexagonal, lapisan bisnis atau domain ditempatkan sebagai inti dari aplikasi, yang berisi logika bisnis utama. Konsep "port" digunakan untuk mendefinisikan antarmuka yang berfungsi sebagai kontrak untuk komunikasi dengan dunia luar. Konsep "Adapter" digunakan untuk implementasi kontrak ini dan bertanggung jawab untuk menghubungkan aplikasi dengan komponen eksternal seperti basis data atau antarmuka pengguna.

### Konteks dalam Golang
* Paket konteks (context package) adalah paket yang membawa informasi seperti batas waktu (deadline), sinyal pembatalan, atau nilai lain dalam permintaan API.

### Unit Tests
* Untuk menulis pengujian unit:
    * Buat berkas pengujian terpisah untuk setiap paket yang ingin diuji.
    * Berkas pengujian harus memiliki nama dengan format `nama_test.go`.
    * Impor paket `testing` dalam berkas pengujian.
    * Tulis fungsi yang dimulai dengan kata 'Test', diikuti dengan nama fungsi dan menerima parameter `*testing.T`.
    * Gunakan fungsi pengujian seperti `t.Errorf` untuk menunjukkan kesalahan.
    * Gunakan perintah `go test` untuk menjalankan pengujian.
