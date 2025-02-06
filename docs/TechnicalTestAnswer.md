# Jawaban Technical Test Golang Developer

## **Soal 1: REST API dengan Golang**

### **1. Jelaskan bagaimana Anda meng-handle error handling dan logging dalam kode Anda.**

- Menggunakan `log` package untuk mencatat error yang terjadi.
- Menangani error di setiap operasi penting seperti koneksi database dan request handler.

### **2. Bagaimana Anda memastikan kode Anda aman dari SQL injection?**

- Menggunakan ORM GORM dengan parameter binding untuk mencegah injection.
- Melakukan validasi input dengan validator sebelum memasukkan ke dalam database.

---

## **Soal 2: Unit Testing**

### **1. Jelaskan pentingnya unit testing dalam pengembangan software.**

- Memastikan fungsi berjalan dengan benar sebelum diterapkan dalam sistem.
- Membantu menemukan bug lebih awal dalam tahap pengembangan.
- Mempermudah refactoring kode tanpa mengganggu fungsionalitas sebelumnya.
- Meningkatkan kepercayaan terhadap kode karena telah diuji sebelumnya.

---

## **Soal 3: Integrasi dengan SOAP Service**

### **1. Bagaimana Anda akan mengimplementasikan integrasi ini di Golang?**

- Menggunakan package `github.com/hooklift/gowsdl` untuk mengenerate client dari WSDL.
- Mengirim request SOAP menggunakan struct yang sesuai dengan skema WSDL.
- Menggunakan `http.Client` untuk mengirim request dan membaca response XML.

### **2. Jelaskan library atau tools open source apa yang akan Anda gunakan.**

- `github.com/hooklift/gowsdl` untuk menghasilkan client SOAP.
- `encoding/xml` untuk parsing response XML dari layanan SOAP.

### **3. Contoh sederhana bagaimana meng-handle response XML dari SOAP service.**

```go
 type GetUserResponse struct {
     XMLName xml.Name `xml:"GetUserResponse"`
     Name    string   `xml:"Name"`
     Email   string   `xml:"Email"`
 }

 func parseSOAPResponse(responseBody []byte) (*GetUserResponse, error) {
     var userResponse GetUserResponse
     err := xml.Unmarshal(responseBody, &userResponse)
     if err != nil {
         return nil, err
     }
     return &userResponse, nil
 }
```

---

## **Soal 4: Message Broker (RabbitMQ)**

### **1. Jelaskan bagaimana Anda meng-handle error dan retry mechanism jika terjadi kegagalan.**

- Menggunakan **Dead Letter Exchange (DLX)** untuk menangani pesan gagal.
- Jika terjadi error dalam memproses pesan, pesan dikembalikan ke queue dengan delay tertentu.
- Menggunakan `defer` untuk memastikan koneksi RabbitMQ selalu tertutup dengan benar.

---

## **Soal 5: Containerization dengan Docker**

### **1. Jelaskan manfaat menggunakan containerization dalam pengembangan software.**

- **Konsistensi lingkungan**: Aplikasi berjalan dengan konfigurasi yang sama di berbagai environment.
- **Portabilitas**: Aplikasi dapat dijalankan di berbagai sistem tanpa setup manual.
- **Isolasi**: Setiap aplikasi berjalan dalam lingkungan terpisah.
- **Scalability**: Mempermudah deployment dan scaling menggunakan Kubernetes.

### **2. Bagaimana Anda mengintegrasikan Docker dengan CI/CD pipeline menggunakan tools open source seperti GitHub Actions atau GitLab CI/CD?**

- **GitHub Actions**:
  - Menambahkan workflow YAML di .github/workflows/ci.yml.
  - Build dan push Docker image ke Docker Hub/GitHub Container Registry.
  - Deploy container ke server setelah build selesai.
- **Contoh pipeline untuk GitHub Actions:**

```yaml
name: CI/CD Pipeline
on:
  push:
    branches:
      - main
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v2
      - name: Build Docker image
        run: docker build -t my-app .
      - name: Push to Docker Hub
        run: |
          echo "${{ secrets.DOCKER_PASSWORD }}" | docker login -u "${{ secrets.DOCKER_USERNAME }}" --password-stdin
          docker tag my-app my-dockerhub-username/my-app:latest
          docker push my-dockerhub-username/my-app:latest
```

- **GitLab CI/CD**:
  - Menambahkan pipeline di `.gitlab-ci.yml`.
  - Build dan push Docker image ke GitLab Container Registry.
  - Deploy aplikasi dengan Docker Compose atau Kubernetes.

---

Jika ada pertanyaan lebih lanjut atau revisi, silakan beri tahu! 🚀
