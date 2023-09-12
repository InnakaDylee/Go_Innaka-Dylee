# Summary
## Join, Union, Agregasi, Subquery, Function - DBMS
### Join
* Join adalah klausa untuk menggabungkan record dari dua atau lebih table
* Berikut adalah tipe SQL Join.
    * Inner Join adalah operasi penggabungan data yang menggabungkan baris yang memiliki nilai yang sesuai di kedua tabel. Hasil dari operasi ini hanya akan mencakup baris yang memiliki pertandingan dalam kedua tabel.
        ```
        SELECT column_name(s)
        FROM table1
        INNER JOIN table2
        ON table1.column_name = table2.column_name;
        ```
    * Left Join adalah operasi penggabungan data yang menggabungkan semua baris dari tabel sumber (tabel kiri) dengan baris yang memiliki kesesuaian di tabel target (tabel kanan). Jika tidak ada kesesuaian yang ditemukan di tabel target, maka kolom-kolom yang berasal dari tabel target akan memiliki nilai NULL dalam hasilnya.
        ```
        SELECT column_name(s)
        FROM table1
        LEFT JOIN table2
        ON table1.column_name = table2.column_name;
        ```
    * Right Join adalah kebalikan dari Left Join. Ini menggabungkan semua baris dari tabel kanan dengan baris yang cocok di tabel kiri. Jika tidak ada pertandingan di tabel kiri, nilai-nilai di tabel kiri akan menjadi NULL
        ```
        SELECT column_name(s)
        FROM table1
        RIGHT JOIN table2
        ON table1.column_name = table2.column_name;
        ```
### Union
* Union adalah operasi dalam SQL yang digunakan untuk menggabungkan hasil dari dua atau lebih pernyataan SELECT menjadi satu hasil tunggal
    ```
    SELECT column1, column2 FROM table1
    UNION
    SELECT column1, column2 FROM table2;
     ```
### Agregasi
* Fungsi Agregasi adalah jenis fungsi di mana nilai-nilai dari beberapa baris digabungkan atau dihitung bersama-sama untuk menghasilkan sebuah nilai ringkasan tunggal.
* Berikut beberapa fungsi Agregasi SQL.
    * MIN and MAX

        ```
        SELECT MIN(column_name)
        FROM table_name
        WHERE condition;
        ```

        ```SELECT MAX(column_name)
        FROM table_name
        WHERE condition;
        ```
    * SUM
        ```
        SELECT SUM(column_name)
        FROM table_name
        WHERE condition;
        ```
    * AVG
        ```
        SELECT AVG(column_name)
        FROM table_name
        WHERE condition;
        ```
    * COUNT
        ```
        SELECT COUNT(column_name)
        FROM table_name
        WHERE condition;
        ```
    * HAVING
        ```
        SELECT column_name(s)
        FROM table_name
        WHERE condition
        GROUP BY column_name(s)
        HAVING condition
        ORDER BY column_name(s);
        ```

### Subquery
* Subquery, juga dikenal sebagai Inner query atau Nested query, merupakan sebuah pertanyaan atau query SQL yang terdapat di dalam query SQL lainnya.
* Subquey digunakan untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil
    ```
    (SELECT [DISTINCT] subquery_select_argument
    FROM {table_name | view_name}
    {table_name | view_name} ...
    [WHERE search_conditions]
    [GROUP BY aggregate_expression [, aggregate_expression] ...]
    [HAVING search_conditions])
    ```
  
### Function
* Function adalah sebuah kumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilnya
    ```
    DELIMITER //

    CREATE FUNCTION function_name [ (parameter datatype [, parameter datatype]) ]
    RETURNS return_datatype

    BEGIN

    declaration_section

    executable_section

    END;

    DELIMITER ;
    ```
    ```
    CREATE TRIGGER trigger_name
    {BEFORE | AFTER} {INSERT | UPDATE| DELETE }
    ON table_name FOR EACH ROW
    trigger_body;
    ```
    