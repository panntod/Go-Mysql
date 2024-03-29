# GO MySql 🚀

## Deskripsi

Dalam repo ini, belajar pemrograman Golang menggunakan basis data MySQL. Mempelajari cara terhubung ke MySQL, mengirim perintah SQL ke MySQL, dan menangani masalah SQL Injection. Menggunakan database driver yang disediakan oleh Third-party dari [Go Developer](https://github.com/go-sql-driver/mysql).

## Fitur

- Sistem CRUD ✨
- Akses MySql 📊
- Driver GO 🚗
- Test 🧪
- Rest API 🌐

## Cara Membuat Aplikasi

### Langkah 1: 

Pastikan sudah menginstal go.msi atau download menggunakan link ini:

[Download Go for Windows](https://go.dev/dl/go1.21.5.windows-amd64.msi).

### Langkah 2: 

Jalan kan perintah ini untuk membuat package main:
```
go mod init main
```

### Langkah 3: 

Jalan kan perintah ini untuk mendapatkan package driver:
```
go get -u github.com/go-sql-driver/mysql
```

### Langkah 4: 

Jalan kan perintah ini untuk mendapatkan framework gin:
```
go get -u go get -u github.com/gin-gonic/gin
```
atau jalankan perintah ini:
```
go mod tidy
```

### Langkah 5: 

Buat Database MySql Dengan menjalankan Perintah:
```sql
CREATE DATABASE go_crud;
```

### Langkah 6: 

Buat Table go-user MySql Dengan menjalankan Perintah:
```sql
CREATE TABLE `go-user` (
  `id` int(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `nama` varchar(255) NOT NULL,
  `umur` int(10) NOT NULL,
  `alamat` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

Atau Import Table dari [Tabel Go-User](https://github.com/panntod/Go-Mysql/tree/main/MySql)

### Langkah 7: 

Jalankan code dengan perintah
```go
go run main.go
```

## Urutan Pembuatan Source Code

1. **database_test.go**
2. **database.go**
3. **sql_test.go**
4. **main.go**
    - Connection
    - Main
      - Router
        - Post
        - Get
        - Get By Id
        - Update
        - Delete
      - Run App
5. **request.rest**   
    - Request
      - Post
      - Get
      - Get By Id
      - Update
      - Delete

