# Summary Time Complexity & Space Complexity

## Big-O Notation
* Notasi Big-O digunakan untuk mengevaluasi tingkat kompleksitas dari sebuah algoritma
* Kompleksitas algoritma dibagi menjadi dua bagian, yakni Time Complexity dan Space Complexity

## Time Complexity
* Time Complexity adalah teknik mengukur seberapa lama waktu yang diperlukan oleh sebuah algoritma untuk menyelesaikan sebuah masalah tergantung pada input yang diberikan.Berikut jenis time complexity.

    * O(1) — Constant 
    Constant berarti bahwa jumlah input yang diberikan kepada algoritma tidak akan memengaruhi waktu eksekusi algoritma tersebut.
        ```
        go
        func dominant (n int) int {
            var result int = 0
            result = result + 10
            return result
        }
        ```


    * O(n) — Linear
    Linear berarti jumlah waktu proses (runtime) akan berbanding lurus dengan seberapa banyak inputan.
        ```
        go
        func dominant (n int) int {
        result := 0
            for i := 0; i < n; i++ {
                result += 1
            }
            result = result + 10
            return result
        }
        ```


    * O(n+m) — Linear
        ```
        go
        func linear (n int, m int) int {
            result := 0
            for i := 0; i < n; i++ {
                result += 1
            }
            for j := 0; j < m; j++ {
                result += 1
            }
            return result
        }
        ```


    * O(log n) — Logaritmic
    Logarithmic Time mengacu pada situasi di mana saat input sebesar n diberikan kepada suatu fungsi, jumlah langkah yang dilakukan oleh fungsi tersebut mengurang secara proporsional berdasarkan suatu faktor.
        ```
        go
        func logarithmic(n int)int {
        result := 0
            for n > 1 {
                n /= 2
                result += 1
            }
            return result
        }
        ```


    * O(n^2) — Quadratic
    Quadratic Time terjadi ketika waktu eksekusi dari suatu fungsi adalah sebanding dengan n^2, di mana n merupakan jumlah input dari fungsi tersebut.
        ```
        go
        func quadratic(n int) int {
            result := 0
            for i := 0; i < n; i++ {
                for j := 0; j < n; j++ {
                result += 1
                } 
            }
            return result
        }
        ```
        
    
    * Beberapa Time Complexity lainnya diataranya, yaitu : O(2^n) — Exponential, O(n!) — Factorial

## Space Complexity
* Space Complexity mengukur seberapa efisien algoritma dalam menggunakan memori saat menyelesaikan sebuah masalah tergantung pada ukuran input yang diberikan. Jenis dari space complexity sama seperti time complexity


