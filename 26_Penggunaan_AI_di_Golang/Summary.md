# Summary
## Penggunaan AI di Golang
### Mengapa Memilih Go?
Beberapa alasan pemilihan bahasa pemrograman Go adalah:
- Dikembangkan oleh Google untuk pengembangan aplikasi server-side dan berbasis cloud.
- Go merupakan bahasa yang dikompilasi (compiled), sehingga lebih cepat dalam eksekusi dibandingkan dengan bahasa yang diinterpretasi (contoh: Python).
- Mendukung konkurensi (concurrency) melalui fitur goroutines, yang memungkinkan untuk menjalankan banyak tugas secara simultan tanpa memberatkan penggunaan memori.
- Memiliki pustaka standar (standard libraries) yang sangat kaya, didukung oleh komunitas pengembang yang kuat.

### Pustaka Go untuk Pembelajaran Mesin (Machine Learning)
* GoLearn : Merupakan proyek open source yang menonjolkan sederhana dan kemampuan penyesuaian yang tinggi.
* Gorgonia : Mempermudah penulisan dan evaluasi persamaan matematika yang melibatkan matriks multidimensi. Ini adalah pustaka tingkat rendah (low-level library) yang mirip dengan Theano, tetapi memiliki tingkat abstraksi lebih tinggi daripada Tensorflow.
* Gonum : Sebuah rangkaian pustaka yang dirancang untuk membuat penulisan algoritma numerik dan sains lebih produktif, dengan kinerja tinggi, dan skalabilitas. Mirip dengan numpy dan scipy.
* Gomind

### AI as a Service - AIaaS
AIaaS adalah alat kecerdasan buatan yang siap digunakan (pre-built or off-the-shelf AI tools). Beberapa perusahaan penyedia AIaaS termasuk AWS, GCP, Microsoft Azure, IBM, dan OpenAI.

* Jenis-jenis AIaaS
    * Bots/Chatbots
        * Amazon Lex (AWS)
        * Azure Bot Service (Microsoft Azure)
        * DialogFlow (GCP)

    * APIs
        * Amazon Rekognition
        * Amazon Comprehend
        * Azure Cognitive Service
        * Azure Speech Services
        * Google Cloud Vision
        * Cloud Natural Language

    * Machine Learning
        * Amazon SageMaker
        * Azure Machine Learning
        * Google Cloud AI

* Kelebihan
    * Mengurangi biaya
    * Mode yang rendah (Low-mode)
    * Proses implementasi cepat
    * Fleksibilitas
    * Kegunaan
    * Skalabilitas
    * Kustomisasi

* Kekurangan
    * Biaya jangka panjang yang dapat meningkat
    * Masalah privasi mungkin muncul
    * Keamanan
    * Kurangnya transparansi
    * Risiko ketergantungan pada penyedia layanan (vendor lock-in)

### API OpenAI
* Tujuan:
    * Merancang prompt yang efektif (prompt engineering)
    * Membangun aplikasi berbasis AI dengan menggunakan API OpenAI

* Harga API OpenAI
    * Mulai dengan gratis, yang memberikan kredit gratis senilai $5 yang dapat digunakan selama tiga bulan pertama setelah pendaftaran.
    * Pembayaran sesuai penggunaan, dengan tarif berbeda untuk setiap model.

* Langkah-langkah untuk mendapatkan kunci API OpenAI:
    * Kunjungi situs web OpenAI.
    * Daftar jika Anda belum memiliki akun, atau masuk jika Anda sudah memiliki akun.
    * Akses halaman kunci API (API keys) atau melalui tautan berikut: API Keys.
    * Setelah Anda berada di halaman kunci API, klik tombol "create new secret key".
