# Summary String - Advance Function - Pointer - Method - Struct and Interface

## String
* String merupakan sebuah tipe data yang dapat menampung character dari setiap inputan. tipe data string memiliki sebuah library strings yang memiliki beberapa fungsi yang dapat mengatur isi dari string itu sendiri.
    * Contoh fungsi yang dapat digunakan pada string yaitu :
        * len() : yang berguna untuk membaca panjang string
        * compare() : yang berguna untuk membandingkan isi dari dua string
        * contains() : yang berguna untuk melihat apakah pada suatu string terdapat huruf dari substring
        * substring : merupakan sebuah bagian dari sebuah string
        * replace() : yang berguna untuk mengganti isi dari string dari string lainnya
        * insert() : yang berguna untuk menambah isi sebuah string

## Advance Function
* Variadic Function
    * Untuk menghindari membuat slice sementara hanya untuk dilewatkan ke dalam fungsi.
    * Ketika jumlah parameter masukan tidak diketahui.
    * Untuk menyatakan niat Anda dalam meningkatkan keterbacaan.  

* Anonymous Function sebuah fungsi anonim adalah fungsi yang tidak memiliki nama. Hal ini berguna saat Anda ingin membuat fungsi dalam baris kode.
    ```
    func main() {
        // Anonymous function
        func() {
            fmt.Println("Welcome! to GeeksforGeeks")
        }()

        // Assigning an anonymous function to a variable
        value := func() {
            fmt.Println("Welcome! to GeeksforGeeks")
        }
        value()

        // Passing arguments in anonymous function
        func(sentence string) {
            fmt.Println(sentence)
        }("GeeksforGeeks")
    }
    ```

* Closure sebuah closure adalah jenis khusus dari fungsi anonim yang merujuk pada variabel-variabel yang dideklarasikan di luar fungsi itu sendiri. Dalam hal ini, kami akan menggunakan variabel yang tidak disertakan sebagai parameter ke dalam fungsi, tetapi sebaliknya variabel-variabel ini tersedia saat fungsi dideklarasikan.
    ```
    func newCounter() func() int {
        count := 0
        return func() int {
            count += 1
            return count
        }
    }
    ```

* Defer Function sebuah fungsi tertunda (deferred function) func hanya akan dijalankan setelah fungsi utama (parent func) kembali. Penggunaan multiple return juga bisa digunakan, dan mereka dijalankan seperti tumpukan (stack), satu per satu.
    ```
    func main() {
        defer func() {
            fmt.Println("Later")
        }()

        fmt.Println("First")
    }
    ```

## Pointer
* WHAT IS POINTER? Pointer adalah suatu variabel yang menyimpan alamat memori dari variabel lain. Pointer memiliki kemampuan untuk mengubah data yang ditunjuk olehnya.
    * Deklarasi Pointer
        ```
        func main() {
            var name string = "John"
            var nameAddress *string = &name
            fmt.Println("name (value)   :", name)                // John
            fmt.Println("name (address) :", &name)               // 0xc000010050
            fmt.Println("nameAddress (value)   :", *nameAddress) // John
            fmt.Println("nameAddress (address) :", nameAddress)  // 0xc000010050

            name = "Doe"
  
            fmt.Println("")  
            fmt.Println("name (value)   :", name) // Doe
            fmt.Println("name (address) :", &name) // 0xc20800a220
            fmt.Println("nameAddress (value)   :", *nameAddress) // Doe
            fmt.Println("nameAddress (address) :", nameAddress) // 0xc20800a220
        }
        ```
        * Perubahan pada variabel dengan referensi memori yang sama akan berdampak pada satu sama lain.

    * 2 Operator penting dalam pointer
        * Operator * Deferencing berguna untuk deklarasi pointer variable dan mengakses value dari sebuah alamat
        * Operator & Referencing berguna untuk mengembalikan alamat variable dan mengakses alamat variable ke sebuah pointer

## Struct
* Sebuah struktur adalah tipe yang didefinisikan oleh pengguna yang berisi kumpulan bidang/properti atau fungsi (metode) yang diberi nama.
    * Deklarasi Struct
        ```
        type Person struct {
            FirstName string
            LastName  string
            Age       int
        }
        ```
    
    *  Mengakses dan menggunakan sebuah struct
        ```
        var Person0 Person
        Person0.FirstName = "Muchson"
        Person0.LastName = "Ibi"
        Person0.Age = 27
        fmt.Println(Person0.FirstName, Person0.LastName, Person0.Age)
        ```

## Method
* Metode adalah fungsi yang melekat pada suatu tipe
    * Deklarasi Method
        ```
        func (receiver StructType) MethodName(parameterList) (returnTypes) {
        // block statement
        }
        ```

    * Keunggulan Method
        * Membantu Anda dalam menghasilkan kode dalam gaya berorientasi objek di Go.
        * Metode membantu menghindari konflik penamaan.
        * Pemanggilan metode jauh lebih mudah dibaca dan dipahami daripada pemanggilan fungsi.

## Interface
* Interface adalah kumpulan dari tanda tangan metode yang dapat diimplementasikan oleh suatu objek. Oleh karena itu, antarmuka mendefinisikan perilaku dari objek tersebut.
    * Deklarasi Interface
        ```
        type interface_name interface {
            method_name1 <return_type>
            method_name2 <return_type>
            method_name3 <return_type>
            ...
            method_namen <return_type>
        }
        ```

    * Implementasi Interface'
        ```
        type calculate interface {
            large() int
        }

        type square struct {
            side int
        }

        func (s square) large() int {
            return s.side * s.side
        }

        func main() {
            var dimResult calculate
            dimResult = square{10}
            fmt.Println("large square :", dimResult.large())
        }
        ```

## Package
* Package adalah kumpulan dari fungsi dan data.
    * Membuat variable akses ke package yang berbeda
        * var ageOfUniverse int <= package lain tidak dapat mengakses karena huruf pertama merupakan huruf kecil
        * var AgeOfUniverse int <= package lain dapat mengakses karena huruf pertama merupakan huruf besar

    * Contoh Package
        ```
        // aritmatika/package.go
        package aritmatika

        func Tambah(a, b int) int {
            return a + b
        }

        func Kurang(a, b int) int {
            return a - b
        }

        ```
       
        ```
        // main.go
        package main

        import (
            "aritmatika"
            "fmt"
        )

        func main() {
            fmt.Println(aritmatika.Tambah(2, 3))
        }

        ```

## Error Handling
* ERROR, PANIC & RECOVER
Jika pada bahasa pemrograman berbasis OOP terdapat try and catch, maka golang mempunya panic dan recover

* Error Handling Object
Jika Anda sedang membuat sebuah metode yang memerlukan pengembalian error jika terjadi masalah di tengah-tengah, gunakan paket 'errors' untuk tujuan tersebut. Mari lihat contoh sederhana berikut:
    ```
    package main

    import "fmt"

    import (
        "errors"
    )

    func myFunc(i int) (int, error) {
        if i <= 0 {
            return -1, errors.New("should be greater than zero")
        }
        return i, nil
    }

    func main() {
        result, err := myFunc(-1)
        fmt.Println(result, err)
    }
    ```

* Panic & Recover
    * Panic : Ketika runtime Go mendeteksi kesalahan-kesalahan semacam ini, terjadi panic.
    * Recover : Untuk menambahkan kemampuan untuk pulih dari kesalahan panic, Anda bisa menambahkan sebuah fungsi anonim atau mendefinisikan sebuah fungsi kustom dan memanggilnya dengan kata kunci 'defer' dari dalam metode, di mana panic mungkin terjadi dari panggilan internal lainnya.
        ```
        package main

        import "fmt"

        func myMethod() {
            defer func() {
                if err := recover(); err != nil {
                fmt.Println("Error Message:", err)
                }
            }()

            anOddCondition := true
            if anOddCondition {
                panic("I am panicking")
            }
        }

        func main() {
            myMethod()
        }
        ```



