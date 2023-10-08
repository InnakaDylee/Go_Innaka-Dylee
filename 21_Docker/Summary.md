# Summary
## Docker

### Docker
Docker adalah sebuah platform perangkat lunak yang digunakan untuk mengembangkan, mengemas, dan menjalankan aplikasi dalam wadah (container) yang ringan. Ini memungkinkan pengembang untuk mengisolasi aplikasi beserta seluruh dependensinya dalam lingkungan yang konsisten, portabel, dan dapat diimplementasikan di berbagai platform, termasuk mesin fisik, mesin virtual, atau cloud.

### Docker Basics
1. **Image**
2. **Container**
3. **Engine**
4. **Registry**
5. **Control Plane**

### Container
Kontainer adalah unit standar yang dapat dijalankan yang berisi aplikasi dan semua dependensinya. Kontainer memungkinkan pengembang untuk mengemas aplikasi beserta seluruh komponennya (seperti pustaka, kode, variabel lingkungan, dll.) ke dalam lingkungan terisolasi yang konsisten dan portabel.

#### Perbedaan Container dan Mesin Virtual
Kontainer:
1. Abstraksi pada lapisan aplikasi.
2. Kontainer mengonsumsi ruang yang lebih sedikit dibandingkan dengan Mesin Virtual.
3. Menangani lebih banyak aplikasi dan memerlukan sumber daya yang lebih sedikit daripada Mesin Virtual.
4. Menghindari overhead Mesin Virtual dan sistem operasi.

Mesin Virtual:
1. Abstraksi pada lapisan perangkat keras fisik.
2. Setiap Mesin Virtual menyertakan salinan lengkap sistem operasi.
3. Waktu boot yang lebih lambat.

### Syntax Docker
- `FROM` -> Mendapatkan gambar dari registri Docker.
- `RUN` -> Menjalankan perintah bash saat membuat Kontainer.
- `ENV` -> Menetapkan variabel lingkungan di dalam Kontainer.
- `ADD` -> Menyalin file dengan beberapa proses lainnya.
- `COPY` -> Menyalin file.
- `WORKDIR` -> Mengatur direktori kerja.
- `ENTRYPOINT` -> Menjalankan perintah setelah selesai membangun Kontainer.
- `CMD` -> Menjalankan perintah tetapi dapat ditimpa.

### Pengenalan Penyebaran (Deployment)
Penyebaran adalah aktivitas yang bertujuan untuk menyebarkan aplikasi/produk yang telah dikembangkan oleh para pengembang, seringkali untuk mengubahnya dari status sementara menjadi permanen.

### Strategi Penyebaran
1. **Strategi Penyebaran Big-Bang**
2. **Strategi Penyebaran Rollout**
3. **Strategi Penyebaran Blue/Green**
4. **Strategi Penyebaran Canary**

### Docker Volume
Docker Volume dapat dianggap sebagai penyimpanan atau tempat penyimpanan data dalam sebuah kontainer. Secara alami, ketika membuat kontainer, kita tidak ingin data di dalam kontainer terhapus saat kontainer dimatikan. Oleh karena itu, Docker Volume dapat digunakan untuk mengelola persistensi data.

### Docker Network
Secara default, kontainer dalam Docker terisolasi satu sama lain. Anda tidak dapat melakukan permintaan API, misalnya, dari satu kontainer ke kontainer lainnya. Untuk mengaktifkan komunikasi antara kontainer, Anda perlu membuat dan mendaftarkannya dalam jaringan yang sama.
