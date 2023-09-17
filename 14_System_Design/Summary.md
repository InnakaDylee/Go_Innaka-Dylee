# Summary

## System Design

### Diagram Design
* Suatu gambaran visual merupakan penyajian grafis dari data atau informasi yang digunakan untuk menggambarkan hubungan, pola, atau struktur yang ada dalam data tersebut.
* **Alat-alat Desain Diagram**
    * Smartdraw
    * Lucidchart
    * draw.io
* Di bawah ini adalah jenis*jenis diagram yang sering digunakan dalam desain perangkat lunak.
    * **Flowchart**
        * Konsep ini bertujuan untuk menggambarkan satu proses.
        * Terdiri dari Proses, Keputusan, dan Terminator.
    * **Use Case**
        * Diagram use case merangkum bagaimana pengguna sistem (yang disebut aktor) berinteraksi dengan sistem dengan menjelaskan detail tindakan mereka dan tanggapan yang mereka berikan dalam interaksi tersebut.
    * **Diagram Entity Relationship (ERD)**
        * ERD adalah alat yang digunakan untuk menggambarkan hubungan antara entitas seperti individu, objek, atau konsep dalam suatu sistem serta bagaimana interaksi antara mereka terjadi.
    * **High-Level Architecture (HLA)**
        * Ini adalah desain keseluruhan sistem secara umum.


### Characteristics of Distributed Systems
* Scalability
    * Scalability adalah kemampuan sistem, proses, atau jaringan untuk berkembang dan mengelola peningkatan permintaan.
    * Sistem terdistribusi yang dapat terus berkembang untuk mendukung peningkatan jumlah pekerjaan dianggap dapat diskalakan.
    * Sistem mungkin harus diskalakan karena berbagai alasan seperti peningkatan volume data atau peningkatan jumlah pekerjaan.
    * Sistem yang dapat diskalakan ingin mencapai penskalaan tanpa kehilangan performa.

* Reliability
    * Reliability adalah probabilitas sistem akan gagal dalam periode tertentu.
    * Sistem terdistribusi dianggap dapat di reliability jika sistem tersebut tetap memberikan layanannya bahkan ketika satu atau beberapa komponen perangkat lunak atau perangkat kerasnya gagal.
    * Reliability mewakili salah satu karakteristik utama dari setiap sistem terdistribusi.

* Availability
    * Availability adalah waktu sistem tetap beroperasi untuk menjalankan fungsi yang dibutuhkan dalam periode tertentu.
    * Hal ini merujuk pada metrik sederhana yang mengukur seberapa lama sistem, layanan, atau mesin tetap beroperasi dalam keadaan normal.

* Efficiency
    * Efficiency adalah kemampuan sistem untuk menghasilkan hasil yang diinginkan dengan menggunakan sumber daya yang tersedia secara optimal
    * Efektivitas sistem terdistribusi sangat terkait dengan cara sistem mengelola sumber daya, membagi pekerjaan, dan meningkatkan efisiensi komunikasi antara komponennya.

* Serviceaility or Manageability
    * Serviceability atau Manageability merujuk pada tingkat kesederhanaan dan kecepatan ketika melakukan perbaikan atau pemeliharaan suatu sistem.
    * Jika waktu yang dibutuhkan untuk memperbaiki sistem yang mengalami kegagalan meningkat, maka ketersediaannya akan berkurang.
    * Aspek-aspek yang perlu dipertimbangkan dalam pengelolaan meliputi kemudahan dalam mendiagnosis dan memahami masalah, kemampuan untuk melakukan modifikasi, serta tingkat kesederhanaan dalam operasionalisasi sistem.


### Horizontal Scaling vs Vertical Scaling
* Horizontal Scaling (Scaling out)
    * Horizontal Scaling adalah tindakan meningkatkan kapasitas sistem dengan menambahkan lebih banyak mesin atau server ke dalam infrastrukturnya.
    * Dengan mendistribusikan beban kerja ke beberapa mesin, sistem dapat menangani peningkatan beban dengan lebih baik, yang membuatnya lebih efisien dalam mengatasi permintaan yang lebih besar.


* Vertical Scaling (Scaling up)
    * Vertical Scaling adalah upaya meningkatkan kapasitas sistem dengan meningkatkan daya komputasi pada satu mesin atau server yang sudah ada.
    * Ini melibatkan peningkatan perangkat keras seperti CPU, RAM, atau kapasitas penyimpanan pada server tunggal untuk meningkatkan kemampuan sistem dalam menangani beban yang lebih besar.

### Job/Work Queue
- Job Queue/Batch Queue mirip dengan daftar pekerjaan yang diorganisir oleh perangkat lunak penjadwalan untuk pelaksanaan.
- Work Queue adalah suatu kerangka kerja yang digunakan untuk mengembangkan aplikasi besar yang mengatur pekerjaan dan melibatkan eksekusi di ribuan komputer yang tersebar di berbagai kluster, cloud, dan jaringan.

### Load Balancing
* Load Balancing berguna dalam membagi lalu lintas ke sekelompok server dengan tujuan meningkatkan responsifitas dan ketersediaan aplikasi, situs web, atau database.meningkatkan daya tanggap dan ketersediaan aplikasi, situs web, atau database
* Selain itu, Load Balancing memonitor status seluruh sumber daya saat mengatur distribusi permintaan, dan jika ada server yang tidak tersedia atau tidak responsif, atau mengalami tingkat kesalahan yang tinggi, Load Balancing akan menghentikan pengiriman lalu lintas ke server tersebut.

### Monolithic vs Microservices
* Monolitik
    * Basis kode tunggal dengan modul-modul.
    * Modul dibagi berdasarkan fitur bisnis atau teknis.
    * Menggunakan sistem build tunggal dan satu file biner tunggal yang dapat dieksekusi.

* Microservices
    * Layanan yang dapat diterapkan secara independen dan terkait dengan domain bisnis.
    * Berkomunikasi melalui jaringan.
    * Menawarkan berbagai pilihan pemecahan masalah.
    * Arsitektur berbasis beberapa microservices yang berkolaborasi.

### SQL vs NoSQL
* ACID
    * Atomicity: Transaksi adalah entitas tunggal yang harus dilaksanakan sepenuhnya atau gagal sama sekali.
    * Consistency: Transaksi harus menjaga konsistensi basis data sebelum dan setelah pelaksanaan, mematuhi semua aturan integritas.
    * Isolation: Transaksi harus berjalan seolah-olah adalah satu-satunya transaksi yang berjalan, terisolasi dari transaksi lainnya.
    * Durability: Hasil transaksi harus bertahan bahkan setelah terjadi kegagalan sistem, memastikan data tetap aman dan tidak hilang.

* SQL (Structured Query Language)
    * Mengacu pada basis data relasional.
    * Terstruktur dengan skema yang telah ditentukan sebelumnya.
    * Menggunakan tabel untuk mengorganisasi data dengan baris dan kolom.
    * Cocok untuk data dengan struktur yang ketat dan telah didefinisikan sebelumnya.
    * Contoh DBMS: MySQL, PostgreSQL, Oracle.

* NoSQL (Not Only SQL)
    * Mengacu pada basis data non-relasional.
    * Tidak terstruktur dan memiliki skema yang fleksibel.
    * Menggunakan skema yang lebih dinamis, memungkinkan penyimpanan data yang berbeda dalam berbagai format seperti dokumen dan grafik.
    * Data tidak diorganisasi dalam tabel, dan model data lebih variatif.
    * Cocok untuk data tanpa struktur yang ketat atau yang memerlukan akses horizontal.
    * Contoh DBMS: MongoDB, Cassandra, Redis.

### Caching
* Caching adalah metode penyimpanan sementara data dalam memori atau penyimpanan yang lebih cepat.
* Bertujuan untuk meningkatkan kinerja dan responsivitas sistem dengan mengurangi waktu akses ke data yang sering digunakan.
* Data yang sering diakses disimpan dalam cache agar akses lebih cepat.

### Database Indexing
* Database indexing melibatkan pembuatan struktur data tambahan.
* Digunakan untuk mempercepat pencarian dan penyortiran data dalam basis data.
* Menyimpan data yang telah diurutkan dengan cara tertentu sehingga query dapat menemukan data lebih efisien.

### Database Replication
* Database replication adalah proses menggandakan data secara real-time dari satu basis data ke basis data lainnya.
* Bertujuan untuk redundancy, ketersediaan, dan skalabilitas.
* Memastikan data yang sama tersedia di berbagai server, meningkatkan ketersediaan dan keandalan data.
