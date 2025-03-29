# API Bank Merchant

## Ikhtisar

API Bank Merchant adalah aplikasi berbasis Go yang menyediakan fungsionalitas untuk mengelola pedagang, pelanggan, dan transaksi. Ini memungkinkan pemrosesan pembayaran, pengelolaan akun pelanggan, dan penanganan riwayat transaksi.

## Fitur

-   Mengelola pedagang dan detailnya
-   Mengelola akun pelanggan
-   Memproses transaksi
-   Melihat riwayat transaksi

## Instalasi

1. Kloning repositori:

    ```bash
    git clone https://github.com/yourusername/merchant-bank-api.git
    ```

2. Arahkan ke direktori proyek:

    ```bash
    cd merchant-bank-api
    ```

3. Instal dependensi:
    ```bash
    go mod tidy
    ```

## Penggunaan

Untuk memulai server API, jalankan:

```bash
go run main.go
```

Server akan dimulai di `http://localhost:8080`.

## Endpoint API

-   **POST /register**: Mendaftar pelanggan baru
-   **POST /login**: Masuk sebagai pelanggan
-   **GET /merchants**: Mengambil daftar pedagang
-   **POST /transactions**: Memproses transaksi baru

## Kontribusi

Kontribusi sangat diterima! Silakan buka masalah atau kirim permintaan tarik.

## Lisensi

Proyek ini dilisensikan di bawah Lisensi MIT.
