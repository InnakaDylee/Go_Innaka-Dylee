# Summary
## Recursive, Number Theory, Sorting, and Searching

### Recursive
* Rekursi adalah konsep dalam pemrograman di mana sebuah fungsi atau algoritma memanggil dirinya sendiri selama eksekusinya. Hal ini memungkinkan pemecahan masalah yang kompleks menjadi lebih sederhana dengan membaginya menjadi masalah yang lebih kecil yang memiliki struktur serupa.
* Jika masalah yang terjadi cukup kecil, maka fungsi rekursif dapat memberikan jawaban langsung
* Jika masalah yang terjadi cukup besar, maka fungsi rekursif akan memanggil dirinya sendiri untuk menyelesaikan masalah tersebut
* Dengan menggunakan rekursif, akan sangat mudah untuk untuk melihat dan merancang alur penyelesaian
* Strategi rekursif : Base Case dan Recurrence relations
* Contoh masalah : faktorial
    ```
    func factorial(n int) int {
        if n == 1 {
            return 1
        } else {
            return n * factorial(n-1)
        }
    }
    ```

### Number Theory
* Number theory merupakan sebuah cabang matematika yang mempelajari tentang bilangan bulat (integers).
* Beberapa contoh topik dari number theory sebagai berikut :
    * Prime Number
        ```
        func checkPrime(number int) bool{
            if number < 2 {
                return false
            }
            sqrtNumber := int(math.Sqrt(float64(number)))
            for i := 2; i <= sqrtNumber; i++ {
                if number % i == 0 {
                    return false
                }
            }
            return true
        }
        ```

    * Greatest Common Divisor (GCD)
        ```
        func gcd(a int, b int) int {
            if b == 0 {
                return a
            }
            return gcd(b, a % b)
        }
        ```
    
    * Least Common Multiple (LCM)
        ```
        func lcm (a int, b int) int {
            return a * (b / gcd(a, b))
        }
        ```
    
### Searching
* Searching adalah proses pencarian data berdasarkan posisi nilai yang ada pada daftar nilai
* Berikut beberapa contoh searching: 
    * Linear Search - O(n)
        ```
        func linearSearch(elements []int, x int) int {
            for i := 0; i < len(elements); i++{
                if elements[i] == x {
                    return i
                }
            }
            return -1
        }
        ```
    
    * Builtins Search
        ```
        elements := []int{12, 15, 15, 19, 24, 31, 53, 59, 60}
        value := 31
        findIndex := sort.SearchInts(elements, value)
        if value == elements[findIndex] {
            fmt.Println("value found in elements")
        } else {
            fmt.Println("value not found in elements")
        }
        ```

### Sorting
* Sorting adalah proses pengurutan data dalam urutan tertentu berdasarkan nilai elemenya
* Urutan yang sering digunakan yaitu numerik dan abjad, diurutkan dari yang awal maupun akhir
* Berikut contoh sorting yaitu :
    * Selection sort - O(n^2)
        ```
        func selectionSort(elements []int) []int {
            n := len(elements)
            for k := 0; k < n; k++ {
                minimal := k
                for j := k + 1; j < n; j++{
                    if elements[j] < elements[minimal]{
                        minimal = j
                    }
                }
                elements[k], elements[minimal] = elements[minimal], elements[k]
            }
            return elements
        }
        ```

    * Counting sort - O(n+k)
        ```
        func countingSort(elements []int, k int) []int {
            count := make([]int, k + 1)
            for i := 0; i < len(elements); i++{
                count[elements[i]]++
            }
            counter := 0
            for i := 0; i < k + 1; i++{
                for j := 0; j < count[i]; j++{
                    elements[counter] = i
                    counter += 1
                }
            }
            return elements
        }
        ```

    * Merge sort - O(log n)