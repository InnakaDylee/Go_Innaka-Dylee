# Summary Data Structure

## WHAT IS ARRAY?
* Array adalah struktur data yang berisi sekelompok elemen, dapat mengandung satu jenis variabel dengan ukuran alokasi tetap. Berbagai jenis data dapat ditangani sebagai elemen dalam array, seperti data Numerik, String, Boolean.
    * Untuk mendeklarasikan sebuah array, Anda perlu menentukan jumlah elemen yang diinginkan di dalam tanda kurung siku [], diikuti oleh tipe elemen yang akan diisi dalam array tersebut.
        ```
        func main() {
            var primes [5]int
            var countries [5]string

            fmt.Println(reflect.ValueOf(primes).Kind())
            fmt.Println(reflect.ValueOf(countries).Kind())
        }
        ```

    * Memberi dan akses array element
        ```
        func main() {
            var countries [2]string
            
            countries[0] = "India"
            countries[1] = "Canada"
            
            fmt.Println(countries[0]) // Access the first element value
            fmt.Println(countries[1]) // Access the second element value
        }
        ```

    * Memberi nilai pada array saat dideklarasikan
        ```
        func main() {
            odd_numbers := [5]int{1, 3, 5, 7, 9}   // Intialized with values
            var even_numbers [5]int = [5]int{0, 2, 4} // Partial assignment

            fmt.Println(odd_numbers)
            fmt.Println(even_numbers)
        }
        ```
    
    * Mencari panjang array menggunakan len()
        ```
        func main() {
            primes := [...]int{2, 3, 3}

            fmt.Println(reflect.ValueOf(primes).Kind())
            fmt.Println(len(primes))
        }
        ```

    * Memberi nilai pada array tertentu
        ```
        func main() {
            even_numbers := [5]int{1: 2, 2: 4}
            fmt.Println(even_numbers)
        }
        ```

    * iterasi pada array
        ```
        for index := 0; index < len(primes); index++ {
            fmt.Println(primes[index])
        }
        ```
    
## WHAT IS SLICE?
* Slice adalah struktur data yang mirip dengan Array, tetapi ukuran alokasinya dapat berubah sesuai kebutuhan.
    * Membuat slice dari array, slice tidak sepenuhnya dinamis tetapi slice menggunakan referensi dari array
    * Panjang sebuah slice ditentukan saat seberapa banyak data yang masuk, dan setiap kapasitas slice penuh maka slice akan terus bertambah dengan melipat gandakan kapasitasnya
    * Cara membuat slice dari array
        ```
        func main() {
            var primes = [5]int{2, 3, 5, 7, 11}

            // Buat sebuah slice dari sebuah array
            var part_primes []int = primes[1:4]

            fmt.Println(reflect.ValueOf(part_primes).Kind())
            fmt.Println(part_primes)
        }

        ```
    
    * Membuat array dengan fungsi make()
        ```
        func make([]T, len, cap) []T
        ```

    * Fungsi yang dapat digunakan pada slice yaitu append() dan copy()
        * appand() berguna untuk menambah isi sebuah slice
        * copy() berguna untuk menyalin sebuah slice


## WHAT IS MAP?
* Sebuah map adalah struktur data yang mengatur data dalam bentuk pasangan kunci dan nilai di mana setiap kunci memiliki keunikannya sendiri.
    * Cara membuat map
        ```
        func main() {
            // long declaration
            var salary = map[string]int{}
            fmt.Println(salary)

            // long declaration with value
            var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
            fmt.Println(salary_a)

            // short declaration
            salary_b := map[string]int{}
            fmt.Println(salary_b)

            // using make
            var salary_c = make(map[string]int)
            salary_c["doe"] = 7000 // assign value
            fmt.Println(salary_c)
        }
        ```

    * Cara menggunakan map
        ```
        func main() {
            // long declaration with value
            var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
            fmt.Println(salary_a, len(salary_a))

            salary_a["nabilah"] = 7000 // assign value
            fmt.Println(salary_a)

            delete(salary_a, "iswanul") // remove value by key
            fmt.Println(salary_a)

            value, exist := salary_a["umam"] // check key is exist
            fmt.Println(value, exist)

            for key, value := range salary_a { // loop over maps
                fmt.Println("->", key, value)
            }
        }
        ```

## WHAT IS FUNCTION?
* Sebuah fungsi adalah sepotong kode yang dipanggil dengan menggunakan namanya. Fungsi adalah cara yang nyaman untuk membagi kode Anda menjadi blok-blok yang berguna. Ini memungkinkan kita untuk menulis kode yang bersih, teratur, dan modular.
    * Deklarasi fungsi
        ```
        func <name_function> () { <statements>}
        func main() {}

        func <name_function> () <type_return> { <statements> }
        func phi() float64 { return 3.14 }

        func <name_function> (<parameter>) { <statements> }
        func square(value int) int {
            return value * value
        }
        ```




