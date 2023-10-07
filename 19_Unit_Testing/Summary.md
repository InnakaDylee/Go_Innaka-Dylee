# Summary
## Unit Testing
### Survey
* Apa yang dimaksud dengan software testing? Pengujian perangkat lunak adalah proses untuk menganalisis elemen-elemen perangkat lunak guna mengidentifikasi perbedaan antara kondisi saat ini dan kebutuhan yang diperlukan (yaitu, cacat), serta untuk mengevaluasi fitur-fitur perangkat lunak.

* Tujuan dari pengujian adalah:
    * Mencegah regresi
    * Meningkatkan keyakinan dalam melakukan refactoring
    * Memperbaiki desain kode
    * Mendokumentasikan kode
    * Memfasilitasi skalabilitas tim pengembangan

* Tingkat pengujian yang umum digunakan adalah sebagai berikut:
    * Pengujian Antarmuka Pengguna (UI), juga dikenal sebagai pengujian End-to-End, menguji interaksi keseluruhan melalui antarmuka pengguna.
    * Pengujian Integrasi, menguji modul atau subsistem tertentu melalui API.
    * Pengujian Unit, menguji bagian terkecil yang dapat diuji dari aplikasi, yaitu operasi logis tunggal, melalui metode.

### Framework

* Framework menyediakan alat dan kerangka kerja yang diperlukan untuk menjalankan pengujian secara efisien.

### Struktur

* Ada dua pola umum dalam pengorganisasian file pengujian:
    * Menempatkan file pengujian di dalam folder khusus untuk pengujian.
    * Menyimpan file pengujian bersama dengan file produksi.

### Runner

* Runner adalah alat yang digunakan untuk menjalankan file pengujian. Idealnya, pelari ini bekerja dalam mode jam tangan, di mana pengujian secara otomatis dijalankan ketika ada perubahan dalam kode sumber, sehingga meningkatkan efisiensi dalam pengembangan berbasis TDD (Test-Driven Development). Pemilihan pelari yang berjalan dengan kecepatan tinggi sangat penting.

### Mocking

* Penting untuk memastikan bahwa kasus uji Anda berdiri sendiri, sehingga Anda tidak perlu menguji:
    * Modul pihak ketiga
    * Basis data
    * API pihak ketiga atau sistem eksternal lainnya
    * Kelas objek yang telah diuji di tempat lain

### Coverage
* Pengembang perlu memastikan bahwa telah mencakup dengan baik semua bagian kode yang mereka buat. Alat cakupan (coverage tools) digunakan untuk menganalisis kode aplikasi saat pengujian berjalan, sehingga dapat memastikan cakupan yang memadai.
