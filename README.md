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
Buat file `.env` berdasarkan template `.env.example` dan isi dengan konfigurasi database PostgreSQL, RabbitMQ, dan SOAP service.

## 2. Menjalankan Aplikasi

### a. Menjalankan Database PostgreSQL
Pastikan PostgreSQL sudah terinstal dan jalankan dengan:
```sh
sudo systemctl start postgresql
```
Buat database sesuai konfigurasi dalam `.env`.

### b. Menjalankan RabbitMQ
```sh
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management
```

### c. Menjalankan Aplikasi
```sh
go run main.go
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

