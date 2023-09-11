# Summary
## Clean Code
* Clean code merupakan istilah untuk code yang mudah dipahami oleh pembaca, dan terstruktur rapih.
* Karakteristik clean code sendiri diantaranya :
    * Mudah dipahami
    * Mudah dibaca
    * Efektif dalam penulisan (singkat namun mendeskripikan)      
    * Konsisten dalam kata
    * Tidak ada pemborosan yang tidak perlu
    * Memberi penjelasan berupa komentar
    * Menggunakan fungsi
    * Formating

## Principle
* Prinsip KISS (Keep It Simple) adalah untuk menjaga agar fungsi yang dibuat fokus pada tugasnya yang seharusnya, tanpa campur tangan dalam tugas-tugas lain seperti memodifikasi B, memeriksa fungsi C, dan sebagainya.
* DRY (Don't Repeat Yourself) adalah prinsip yang mendorong kita untuk menghindari duplikasi kode yang sering terjadi akibat tindakan menyalin dan menempel. Untuk mengatasi masalah duplikasi kode ini, kita harus menciptakan fungsi-fungsi yang dapat digunakan berulang kali.

* Refactoring adalah proses memodifikasi struktur kode yang telah dibuat dengan cara mengubah bagian internalnya tanpa mengganggu perilaku eksternalnya. Prinsip KISS dan DRY dapat diterapkan melalui tindakan refactor ini.

Teknik refactoring:
- Membuat sebuah abstraksi
- Memcah kode dengan fungsi/class
- Perbaiki penamaan dan lokasi kode
- Deteksi kode yang memiliki duplikasi