# Summary
## Database Schema, DDL, DML
### Database
* Database adalah sekumpulan data yang terorganisir
* Database Relationship
    * One to One
    * One to Many
    * Many to Many

* Tipe data MySQL : num, huruf, dan date
### DDL (Data Definition Language)
* DDL (Data Definition Language) digunakan untuk mendefinisikan, mengelola, dan mengubah struktur objek dalam database. Statement: Create, Drop, Rename, Alter.
* Berikut adalah beberapa pernyataan DDL dan fungsinya.
    * Kata "CREATE" digunakan untuk menciptakan objek baru dalam database, seperti tabel, indeks, atau tampilan.
        ```
        CREATE DATABASE databasename;
        USE DATABASE databasename;
        CREATE TABLE table_name (
            column1 datatype,
            column2 datatype,
            column3 datatype,
            ....
        );
        DROP TABLE table_name;
        ```

    * Kata "ALTER" berfungsi untuk mengubah struktur objek yang telah ada di dalam database. Ini digunakan untuk menambahkan, mengubah, atau menghapus kolom dalam tabel atau mengganti jenis data kolom.
        ```
        ALTER TABLE table_name 
        ADD column_name datatype;

        ALTER TABLE table_name 
        DROP COLUMN column_name;

        ALTER TABLE table_name 
        ALTER COLUMN column_name datatype;
        ```

    * Kata "DROP" digunakan untuk menghilangkan objek dari dalam database, seperti tabel, indeks, atau tampilan.
        ```
        DROP TABLE table_name;
        ```

    * TRUNCATE digunakan untuk menghapus semua data baris dalam tabel, sementara tetap menjaga struktur tabel itu sendiri. Truncate lebih efisien daripada DROP karena tidak memerlukan proses pembangunan ulang tabel.
       ```
       TRUNCATE TABLE table_name;
       ```

### DML (Data Manipulation Language)
* DML (Data Manipulation Language) digunakan untuk mengelola atau memanipulasi data dalam tabel dalam database
* Pernyataan DML fokus pada operasi seperti menambah, mengubah, menghapus, dan mengambil data dari tabel
* Berikut adalah beberapa pernyataan DML dan fungsinya.
    * Kata "INSERT" digunakan untuk memasukkan data baru ke dalam tabel.
        ```
        INSERT INTO table_name (column1, column2, ...)
        VALUES (value1, value2, ...);
        ```

    * Kata "UPDATE" digunakan untuk memperbarui atau mengubah data yang sudah ada dalam tabel.
        ```
        UPDATE table_name
        SET column1 = value1, column2 = value2, ...
        WHERE condition;
        ```

    * Kata "DELETE" digunakan untuk menghilangkan data dari tabel.
        ```
        DELETE FROM table_name
        WHERE condition;
        ```

    * Kata "SELECT" digunakan untuk mengambil informasi dari tabel.
        ```
        SELECT column1, column2, ...
        FROM table_name
        WHERE condition;

        SELECT * FROM table_name;
        ```

    * Kata "LIKE" digunakan dalam pernyataan SELECT untuk mencari kesamaan nilai kolom dengan suatu pola tertentu. Fungsinya adalah untuk mencari data yang mengandung substring atau karakteristik khusus dalam kolom.
        ```
        SELECT column1, column2, ...
        FROM table_name
        WHERE columnN LIKE pattern;
        ```
    
    * Kata "BETWEEN" digunakan dalam pernyataan SELECT untuk mencari data yang berada di antara nilai-nilai tertentu dalam kolom.
        ```
        SELECT column_name(s)
        FROM table_name
        WHERE column_name BETWEEN value1 AND value2;
        ```
    * Kata "AND" digunakan dalam pernyataan WHERE untuk menghubungkan dua atau lebih kondisi, di mana semua kondisi harus terpenuhi agar baris data dapat dipilih.
        ```
        SELECT column1, column2, ...
        FROM table_name
        WHERE condition1 AND condition2 AND condition3 ...;
        ```
    * Kata "OR" digunakan dalam pernyataan WHERE untuk menggabungkan dua atau lebih kondisi, di mana setidaknya satu dari kondisi tersebut harus benar agar baris data dapat dipilih.
        ```
        SELECT column1, column2, ...
        FROM table_name
        WHERE condition1 OR condition2 OR condition3 ...;
        ```
    * Kata "ORDER BY" digunakan dalam pernyataan SELECT untuk mengatur hasil berdasarkan satu atau lebih kolom, dengan pilihan pengurutan secara ascending (ASC) atau descending (DESC).
        ```
        SELECT column1, column2, ...
        FROM table_name
        ORDER BY column1, column2, ... ASC|DESC;
        ```
    * Kata "LIMIT" digunakan dalam pernyataan SELECT untuk menghadirkan pembatasan pada jumlah baris data yang dikeluarkan oleh kueri.
        ```
        SELECT * FROM table_name LIMIT 10;
        ```

### DCL (Data Control Language)
* DCL (Data Control Language) berfungsi untuk mengatur pengendalian dan otorisasi dalam database.
* Pernyataan DCL memungkinkan pengelola database untuk mengontrol siapa yang dapat mengakses dan melakukan operasi tertentu dalam database
* Berikut adalah beberapa pernyataan DML dan fungsinya.
    * Kata "GRANT" digunakan untuk memberikan hak istimewa atau izin tertentu kepada pengguna dalam database.
        ```
        GRANT permission_type ON object_name TO user_or_role;
        ```
    * Kata "REVOKE" digunakan untuk mencabut atau menghapus izin yang sebelumnya diberikan kepada pengguna dalam database.
       ```
       REVOKE permission_type ON object_name FROM user_or_role;
       ```