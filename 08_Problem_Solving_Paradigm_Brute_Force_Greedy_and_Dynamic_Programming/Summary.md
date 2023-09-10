# Summary
## Problem Solving, Paradigm, Brute Force, Greedy, and Dynamic Programming

### What is Problem Solving Pradigm?
* Problem solving paradigm merupakan pendekatan untuk memecahkan permasalahan
* Setiap paradigma memiliki kelebihan dan kelemahan tertentu
* Pemilihan paradigma yang tepat tergantung pada sifat masalah yang akan dipecahkan
* Berikut beberapa paradigma yang sering digunakan dalam problem-solving
    * Complete Search (Brute Force)
    * Divide and Conquer
    * The Greedy approach
    * Dynamic Programming

### Brute Force
* Complete Search (Brute Force) merupakan metode memecahkan suatu masalah dengan mencari solusi menggunakan cara memeriksa semua kemungkinan solusi secara berurutan
* Brute Force digunakan ketika tidak ada algoritma lain yang tersedia
* Biasanya mudah digunakan
* Kompleksitas waktu tinggi jika kemungkinan yang terjadi besar
* Secara teoritis semua masalah dapat diselesaikan dengan menggunakan pendekatan Brute Force terutama ketika memiliki waktu tidak terbatas
* Contoh : mencari nilai minimal dan maksimal, etc.

### Divide and Conquer
* Divide and Conquer adalah paradigma pemecahan masalah dimana suatu masalah dibuat sesederhana mungkin dengan 'membagi' menjadi submasalah yang lebih kecil
* Setiap submasalah diselesaikan secara terpisah dan hasilnya digabungkan untuk mendapatkan solusi akhir
* Tahapan pada Devide dan Conquer antara lain.
    * Divide : membagi masalah yang besar menjadi masalah yang lebih kecil
    * Conquer : ketika masalah sudah cukup kecil untuk diselesaikan maka langsung diselesaikan
    * Combine : jika dibutuhkan maka perlu menggabungkan solusi dari submasalah yang lebih kecil menjadi solusi untuk masalah yang besar
* Contoh : algoritma pencarian Binary search, etc.

### The Greedy Approach
* Pendekatan ini memilih langkah terbaik yang tersedia pada setiap langkah agar mencapai solusi yang optimal
* efektif, tetapi tidak selalau menghasilkan solusi optimal
* Algoritma ini digunakan dalam masalah optimasi yang memungkinkan pemilihan langkah terbaik secara lokal
* Contoh : algorima Huffman Coding, Activity Selection, Djikstra algoritma, etc.

### Dynamic Programming
* Dynamic Programming adalah paradigma problem solving untuk memecahkan masalah kompleks dengan memecahkan menjadi serangkaian submasalah yang lebih kecil dan menghindari perhitungan berulang dengan menyimpan hasil dari setiap submasalah
* Dynamic Programming digunakan untuk mengatasi masalah dengan struktur rekursif atau tumpang tindih
* Berikut merupakan karakteristik Dynamic Programming
    * overlapping subproblem
        * Karakteristik ini mengacy pada sifat masalah yang memiliki submasalah yang rekursif atau tumpang tindih
        * Beberapa submasalah perlu dipecahkan lebih dari sekali
        * Solusi dari submasalah yang sama dapat digunakan kembali
        * Dynamic programming meyimpan hasil perhitungan kembali
        * Dynamic programming meyimpan hasil perhitungan submasalah untuk menghindari perhitungan berulang
    
    * optimal substructure property
        * Solusi optimal dari masalah besar dapat didapatkan dari solusi optimal submasalah yang lebih kecil
        * Struktur masalah memungkinkan penggunaan solusi optimal submasalah untuk mencapai solusi optimal secara global

* Berikut adalah metode pada Dynamic Programming.
    * Top-Dowm with Memoization
        * Penyelesaian masalah rekursif dengan menyimpan hasil perhitungan submasalah dalam memoization table
        * ketika submasalah ditemukan lagi, solusinya dapat ditemukan pada table daripada melakukan perhitungan ulang
        * mengurangi kompleksitas waktu
    
    * Bottom-Up with Tabulation
        * Menyelesaikan submasalah terkecil terlebih dahulu dan menyimpan hasilnya dalam tabel
        * Solusi untuk masalah lebih besar dibangun berdasarkan hasil yang telah disimpan dalam table
        * Efisien dalam penggunaan memori

* Contoh : algoritma Fibonacci dengan penyimpanan hasil perhitungan sebelumnya, etc.