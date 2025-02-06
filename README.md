# Dokumentasi - Technical Test Golang Developer

## 1. Setup Project

### a. Clone Repository
```sh
git clone https://github.com/Ayyasy123/pt-aka-tech-vision.git
cd pt-aka-tech-vision
```

### b. Install Dependencies
```sh
go mod tidy
```

### c. Setup Environment
Buat file `.env` berdasarkan template `.env.example` dan isi dengan konfigurasi database PostgreSQL

## 2. Menjalankan Aplikasi

Jalankan seluruh layanan yang dibutuhkan menggunakan Docker Compose:
```sh
docker-compose up -d
```

Aplikasi akan berjalan di `http://localhost:8080`.

## 3. Endpoint API

### a. Membuat User
**Endpoint:** `POST /users`
**Request Body:**
```json
{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "password": "secret"
}
```
**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "johndoe@example.com"
}
```

## 4. Unit Testing
### a. Menjalankan Unit Test
```sh
go test ./...
```

## 5. Integrasi SOAP
Menggunakan `github.com/hooklift/gowsdl` untuk berkomunikasi dengan layanan SOAP.

## 6. RabbitMQ
```sh
cd rabbitMQ
```
### a. Mengirim Pesan ke Queue
```sh
go run send/send.go
```

### b. Menerima Pesan dari Queue
```sh
go run receive/receive.go
```

## 7. Penjelasan Tambahan
- **Error Handling:** Menggunakan `log` dan middleware untuk menangani error.
- **Keamanan SQL Injection:** Menggunakan ORM GORM dengan parameterized query.
- **Containerization:** Menggunakan multi-stage build untuk mengoptimalkan image.

