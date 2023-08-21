# Summary Version Control

## Version Control
Version control merupakan cara untuk mengatur atau mengelola source code suatu program menjadi beberapa versi. Verion control biasanya digunakan untuk megatasi duplikasi dari sebuah file agar tidak terjadi perulangan file yang sama dengan cara membuat file versi terbaru dan menghapus file versi lama

## Version Control Tools
Terdapat beberapa tools diantaranya yaitu : Version Control System (VCS), Source Code Manager (SCM), dan Revision Control System (RCS). Dari ketiga tools tersebut Version Control System (VCS) yang sering digunakan

## Sejarah Version Control System 
1. Single User : SCCS -1972 Unix only, RCS - 1982 Cross platform, text only
1. Centralized : CVS - 1986 File focus, Perforce - 1995, Subversion - 2000 - track directory structure, Microsoft Team Foundation Server - 2005
1. Distributed : Git - 2005, Mercurial - 2005, Bazaar - 2005

Git merupakan salah satu version control system yang populer bagi kalangan developer dalam mengembangkan suatu aplikasi secara bersamaan

## Sejarah Git
Dibuat oleh Linux Torvalds pada 2005 yang sekaligus pencipta dari Linux Kernel, Git dibuat secara terdistribusi bukan tersentraliasi yang dimana sebuah tim developer dapat melakukan pengembangan aplikasi secara bersamaan dengan cara mensinkronisasi setiap perubahan kedalam remote computer/server

Setiap perubahan pada file dan folder dalam mengembangan suatu aplikasi akan disimpan kedalam Git Repository(Folder Project) yang dapat mempermudah dalam melihat riwayat perubahan yang ada, setiap perubahan dapat terlihat dengan fitur yang disebut commit yang dimiliki Git. git server sendiri lumayan rumit untuk dioperasikan sehingga munculah layanan yang mempermudah proses management yaitu Github

## Git Install
1. Install git pada mac : Download installer pada website https://git-scm.com/downloads > ikuti instruksi yang tersedia > cek status install dengan cara mengetikkan git --version pada terminal

1. Install git pada windows : Download installer pada website https://git-scm.com/downloads > Jalankan installer yang sudah didownload > ikuti panduan next dan finish

1. Install git pada linux : ketikkan sudo apt-get update pada terminal > setelah selesai ketikkan sudo apt-get install git >> cek status install dengam mengetikkan git --version pada terminal

## Git Setup
Git Init, Clone, config

* git config
* git config --global user.name "your-name"
* git config --global user.email "your-email"

* start with init
* git init
* git remote add remote_name remote_repo_url
* git push -u remote_name lokal_branch_name

* clone git repo
* git clone repo_url
* cd folder_project

## Saving Change
Working directory/local folder >>git add>> Staging area >>git commit>> Repoitory

Git Status, Add, Commit

* git status

* git add nama_file
* git add hello.go
* git add . (all)

* git commit -m "pesan commit"
pesan commit berguna untuk memberi informasi perubahan apa yang telah dilakukan saat melakukan commit

* git diff --staged (add staging area)

* git stash (stashing)

* git stash apply (re-applaying stash changes)

## Inspecting Repository
Git Log, Checkout, Reset

* git log --oneline (melihat perubahan yang lama)

* git checkout code_log (menampilkan perubahan)

* git reset code_log --soft/--hard (mereset perubahan)

Git Push, Fetch, Pull

* git remote -v
* git remote add origin remote_repo_url

* git fetch
* git pull origin branch_name

* git push origin main
* git push origin feature/login-user

## Git Branching
Git Branch

* git branch --list (melihat semua branch)
* git branch branch_name (membuat branch)
* git branch -D branch_name (menghapus branch)
* git branch -a (melihat branch remote)

Git Merge

* git checkout -b new-feature main (membuat branch berdasarkan branch utama)
* git add nama_file
* git commit -m "pesan commit"
* git checkout main
* git merge new-feature (lakukan penyatuan branch)
* git branch -d new-feature

Pull Request biasanya dilakukan untuk meminta penyatuan sebelum di satukan

## Workflow Collaboration
Workflow merupakan teknik untuk memisahkan tahapan pengembangan sebuah aplikasi dengan membuat beberapa branch untuk mengelola aplikasi dengan lebih efektif, beberapa hal penting yang harus diperhatikan yaitu : 
1. Usahakan branch master/main menjadi tempat hasil akhir dari sebuah aplikasi
1. Buat branch development/dev menjadi tempat utama saat aplikasi masih dalam tahapan pengembangan
1. Buat branch untuk setiap fitur yang dibuat dan usahakan untuk menggabungkan setiap fitur yang sudah siap kedalam branch development saja, agar tidak mengganggu branch utama master/main
1. Jika aplikasi sudah dianggap selesai barulah melakukan penggabungan branch development ke dalam branch utama master/main