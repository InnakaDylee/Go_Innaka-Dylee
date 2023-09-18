# Summary
## Intro RESTful API
### What's API ?
* API (Application Programming Interface) adalah sekumpulan aturan dan protokol yang digunakan oleh perangkat lunak untuk berkomunikasi antara satu dengan yang lain. Klien mengirimkan permintaan, kemudian server memprosesnya dan mengirimkan respons. Ini memungkinkan berbagai aplikasi untuk bekerja sama dan berbagi data.
* Berikut adalah beberapa Alat API yang berguna:
    * Katalon
    * Apache JMeter
    * SoapUI
    * Postman   


### REST
* REST adalah singkatan dari Representational State Transfer.
* REST (Representational State Transfer) adalah pendekatan arsitektur aplikasi untuk pengembangan layanan web.
* Pendekatan ini memusatkan perhatian pada sumber daya yang dapat diakses melalui URL dengan menggunakan operasi HTTP.
* REST bersifat stateless, mendukung berbagai representasi data, dan sesuai untuk sistem terdistribusi.
* Format permintaan (Request) dan tanggapan (Response) meliputi JSON, XML, dan SOAP.
* Metode permintaan HTTP termasuk GET, POST, PUT, DELETE, HEAD, OPTION, dan PATCH.


### JSON
* JSON (JavaScript Object Notation) adalah format data dalam teks yang digunakan untuk komunikasi data antara aplikasi.
* Data dalam JSON dinyatakan dalam bentuk pasangan "key-value".

### HTTP Respon Code
* 200 : OK
* 201 : Created
* 400 : Bad Request
* 404 : Not Found
* 401 : Unauthorized
* 405 : Method Not Allowed
* 500 : Internal Server Error

### Open API
* Open API (Application Programming Interface) adalah antarmuka perangkat lunak yang dibuka untuk umum yang memungkinkan developer perangkat lunak eksternal untuk mengakses dan berinteraksi dengan layanan atau sistem tertentu

### Postman
* Postman adalah sebuah aplikasi yang digunakan oleh pengembang perangkat lunak untuk menguji, mengelola, dan mendokumentasikan API (Application Programming Interface)
* Postman memungkinkan pengguna untuk mengirim permintaan HTTP ke API dan mengamati responsnya
* Hal ini memudahkan pengujian fungsionalitas API

### Package net/http
* Package net/http menyediakan implementasi client HTTP dan server

### Package Encoding/JSON
* berisi beberapa fungsi untuk operasi JSON
    * Decode json to object struct
    * Decode json to map [string] interface {} & interface{}
    * Decode array json to array object
    * Encode object to json string