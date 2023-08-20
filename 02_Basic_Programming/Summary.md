# Summary Basic Programming
## Apa itu Go?
Go merupakan bahasa yang relative baru di dunia pemrogramman, Go memiliki keunggulan yaitu kemudahan untuk digunakan, simple, dapat diandalkan dan efisien. Go juga merupakan bahasa yang bagus untuk pengembangan yang membutuhkan kecepatan.

## Kenapa harus Go?
1. Mudah digunakan
1. Dapat digunakan sebagai static type languages maupun dynamic languages
1. Penulisan Syntax yang ringan/mudah
1. Dibangun untuk skala yang besar, jaringan komputer dan menggunakan multi-core hardware
1. Open Source
1. Digunakan perusahaan besar

## Bagaimana cara menggunakan Go?
1. Harus memilih IDE untuk menulis code. contoh IDE : VS Code, Goland, dll.
1. Install Golang dari website https://go.dev/dl/
1. Buat File dengan ekstensi .go
1. Go siap digunakan

## Apa saja yang dipelajar dalam basic programming go?
### 1. Mengetahui Proses Compiling
Pada Go terdapat 3 tahapan compiling yaitu: 
Source Code > Compile > Run

### 2. Command Go pada terminal
* go run : eksekusi program go
* go build : compile program
* go install : seperti go build dan melanjutkan proses install
* go test : untuk melakukan test menggunakan sufix _test.go
* go get : download package go

### 3. Dasar pemrograman pada go
1.Package "fmt" berguna untuk melakukan input output. Syntax : fmt.Printf(), fmt.Scanln(), dll.

2.Variable berguna untuk memberi nama pada data dan variable memiliki type data ( boolean, numeric, string), adapun variable constants yang tidak dapat diubah setelah dideklarasikan. Syntax: var nama_variable type_data, nama_variable := value, const nama_variable = value

3.Operator beguna untuk memproses variable yang sudah diinputkan. Berikut beberapa contoh operator :
* Operator Aritmatika : +, -, *, /, %, ++, --
* Operator Comparison : >, <, >=, <=, ==, !=
* Operator logika     : &&, ||, !
* Operator Bitwise    : &, |, ^, <<, >>
* Miscellunous        : &, * (Pointer)

4.Control structures berguna jika ada pengkondisia dan perulangan. Berikut Contohnya

1.Pengkondisian merupakan suatu kasus dimana memiliki output berbeda dari suatu input. contoh: jika inputan lebih dari 0 maka bilangan positif dan jika inputan kurang dari 0 maka bilangan negatif. Berikut Syntax :
* if (kondisi)
* if (kondisi) else
* if (kondisi) else if (kondisi) else
* switch variable case 1 case 2 default

2.Perulangan merupakan suatu kasus dimana memperbolehkan code untuk melakukan proses berulang kali. Berikut Syntax :
* for init; condition; post{}
* for condition{}
* for init; condition; post { if condition{ continue } if condition{ break }}


