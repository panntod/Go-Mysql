# GO MySql

## Deskripsi

Dalam repo ini, belajar pemrograman Golang menggunakan basis data MySQL. mempelajari cara terhubung ke MySQL, mengirim perintah SQL ke MySQL dan menangani masalah SQL Injection. Menggunakan database driver yang disediakan oleh Third-party dari [go-drive](https://github.com/go-sql-driver/mysql).

## Fitur

- Sistem CRUD GO-User
- Akses MySql
- Driver GO
- Test

## Cara Membuat Aplikasi

### Langkah 1: 

Pastikan sudah menginstal go.msi atau download menggunakan link ini:
```
https://go.dev/dl/go1.21.5.windows-amd64.msi
```

### Langkah 2: 

Jalan kan perintah ini untuk membuat package gosql:
```
go mod init go-sql
```

### Langkah 2: 

Jalan kan perintah ini untuk mendapatkan package driver:
```
go get -u github.com/go-sql-driver/mysql
```

## Urutan Pembuatan Source Code

1. **database_test.go**
2. **database.go**
3. **sql_test.go**
