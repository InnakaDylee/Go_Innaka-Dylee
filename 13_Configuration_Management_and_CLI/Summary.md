# Summary
## Command Line Interface (CLI)
### Apa itu Command Line?
* Command Line adalah salah satu user interface yang berinteraksi langsung dengan user

### User Interface
* User Interface adalah tampilan utama yang langsung berinteraksi dengan pengguna
* User Interface terbagi menjadi beberapa yaitu:
    * GUI (Graphical User Interface) merupakan user interface yang menampilkan gambar sehingga lebih mudah digunakan
    * CLI (Command Line Interface) merupakan user interface yang bebasis tulisan sehingga sedikit sulit untuk digunakan namun lebih cepat, powerful, efektif dan efisien

### Kenapa harus menggunakan Command Line?
* Commend line merupakan control utama untuk setiap OS ataupun aplikasi
* Cepat memanaje banyak tugas pada OS
* Memiliki kemampuan untuk mengirim scripts secara cepat 
* Dapat membantu menganalisa kesalahan bahkan permasalahan koneksi jaringan sekalipun
* Contoh command line interface yaitu: 
    * Shell = CLI for OS Services
        * UNIX Shell
        * Command Prompt (MSDOS)
    * Other app CLI
        * Python REPL
        * MySQL CLI client
        * Mongo shell
        * redis-cli
        * etc

### Intro to UNIX Shell
* Terminal & Shell
    ```
    me@linuxbox:~$
    ```
    * Shell Legend
        * me -> your username
        * linuxbox -> your hostname
        * ~ -> your current path (home)
        * $ -> Shell for normal user
        * `#` -> Shell for root user

* Normal User vs Root User
    * Normal User = dapat merubah /home/username dir only
    * Root User = dapat merubah / dir (all directory)
    * Normal User + Sudoers = dapat menjadi seperti root untuk sementara

* UNIX Shell
    * ZSH
    * BASH

### UNIX Shell Commands
* Directory
    * pwd
    * ls
    * mkdir
    * cd
    * rm
    * cp
    * mv
    * ln

* Files
    * create: touch
    * view: head, cat, tail, less
    * editor: vim, nano
    * permission: chown, chmod
    * different: diff

* Network
    * ping
    * ssh
    * netstat
    * nmap
    * ip addr (ifconfig)
    * wget
    * curl
    * telnet
    * netcat

* Utility
    * man
    * env
    * echo
    * date
    * which
    * watch
    * sudo
    * history
    * grep
    * locate

* Practice
    * ```
        drwxrwxrwx0 2 fransiskaariana staff 64 Feb 8 06:58 access_directory

    * drwxrwxrwx0 : permission
    * 2 : count of file
    * fransiskaariana : owned by (user)
    * staff : staff
    * 64 : size
    * Feb 8 06:58 : create date
    * access_directory : file/directory name
* ____________________________________
    * ```
        d | rwx | rwx | rwx
        ^    ^     ^     ^
        |    |     |     |<= r = read, w = write, x = execute for all other user
        |    |     |<= r = read, w = write, x = execute for the group owner
        |    |<= r = read, w = write, x = execute for the file owner
        |<= file type, - : untuk file reguler, d : untuk direktori

    * rwx = 111 (binary) => 7 (decimal) jadi jika pada ketiga permission adalah rwx maka ditulis 777

    * ```
        chmod 777 access_directory <= untuk memberikan permission penuh

### Shell Script
* Shell, program yang berfungsi sebagai penghubung antara user dan kernel
* Shell Script, merupakan bahasa pemrograman yang dicompile berbasis pada shell command
* Cara membuat file sh
    * ``` echo "#!/bin/sh" > my-script.sh
* Cara menjalankan file sh
    * ``` ./my-script.sh
* Cara membuat variable
    * ``` nama_variable="isi variable"
