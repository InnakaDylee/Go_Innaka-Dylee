package main


type kendaraan struct { // nama struct usahakan menggunakan bahasa inggris karena bahasa internasional ( case jika team atau project manager merupakan orang asing sehingga mempermudah )

totalroda int // nama attribut struct menggunakan bahasa inggris
//tab size agar mudah dijelaskan dan line space dihapus agar tidak pemborosan panjang code
kecepatanperjam int // nama attribut struct menggunakan bahasa inggris

}


type mobil struct { // nama struct usahakan menggunakan bahasa inggris karena bahasa internasional
//tab size agar mudah dijelaskan dan line space dihapus agar tidak pemborosan panjang code
kendaraan // ubah penamaan struct menggunakan bahasa inggris

}


func (m *mobil) berjalan() {// ubah penamaan parameter menggunakan bahasa inggris dan parameter m dijelaskan lebih detail
//tab size agar mudah dijelaskan dan line space dihapus agar tidak pemborosan panjang code
m.tambahkecepatan(10) // penamaan method/fungsi menggunakan bahasa inggris		
//tab size agar mudah dijelaskan dan line space dihapus agar tidak pemborosan panjang code
}


func (m *mobil) tambahkecepatan(kecepatanbaru int) {// ubah penamaan parameter menggunakan bahasa inggris dan parameter m pada parameter struct dijelaskan lebih detail
													// penamaan method/fungsi menggunakan bahasa inggris dan penamaan huruf pertama pada setiap kata menggunakan huruf kapital
m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru // penamaan parameter menggunakan bahasa inggris dan setiap huruf pertama setelah kalimat pertama menggunakan huruf kapitas contoh: newSpeed

}


func main() {
//tab size agar mudah dijelaskan dan line space dihapus agar tidak pemborosan panjang code
mobilcepat := mobil{} // ubah penamaan variable menggunakan bahasa inggris

mobilcepat.berjalan()

mobilcepat.berjalan()

mobilcepat.berjalan()


mobillamban := mobil{}

mobillamban.berjalan()

}

