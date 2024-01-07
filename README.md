# Sistem Informasi Peminjaman Buku Perpustakaan

Produk ini dirancang untuk membantu perpustakaan mengelola koleksi buku dan denda dengan lebih efisien. Fitur utama sistem ini adalah memudahkan dalam manajemen peminjaman buku terhadap member perpustakaan dan kemampuannya untuk menyesuaikan tarif denda untuk setiap buku atau kategori buku tertentu.

## Spesifikasi / Fitur

### [ Guest User ]

- Melihat / Mencari List dan Detail Buku
- Melihat / Mencari List dan Detail Feedback
- Memberikan Feedback sebagai Anonymous

### [ Member ]

- Melihat / Mencari List dan Detail Buku
- Meminjam Buku melalui **Librarian / Staff**
- Membayar Denda Keterlambatan / Kerusakan / Kehilangan
- Melihat / Mencari List dan Detail Feedback
- Memberikan Feedback

### [ Librarian / Staff ]

- Melihat, Mencari, Menambah, Menghapus, Mengubah Data Buku
- Melihat, Mencari, Menambah, Menghapus, Mengubah Data Penerbit
- Melihat, Mencari, Menambah, Menghapus, Mengubah Data Penulis dan hubungannya dengan Buku (Authorship)
- Melihat, Mencari, Menambah, Menghapus, Mengubah Data Peminjaman
- Melihat / Mencari List dan Detail Feedback
- Memberikan Reply pada Feedback
- Menghapus Feedback

## Teknologi

- Framework: **Echo**
- ORM: **GORM**
- Database: **MySQL/MariaDB**
- Deployment: **Google Cloud Platform**
- Code Structure: **Clean Architecture**
- Authentication: **JSONWebtoken**
- Image Server: **Cloudinary**
- Payment Gateway: **Midtrans**

## Link Postman

[Spesifikasi API](https://documenter.getpostman.com/view/28340333/2s9YsJBXiG)

## ERD

![image](https://raw.githubusercontent.com/dev4ult/go-perpustakaan/main/erd-perpus.png)
