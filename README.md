# Final Task PBI - Rakamin Fullstack Developer

API ini dikembangkan sebagai bagian dari syarat kelulusan di PBI Rakamin Academy. API ini menawarkan fitur-fitur autentikasi dan otorisasi pengguna yang dirancang untuk memastikan keamanan dan integritas data. Fitur autentikasi memastikan bahwa hanya pengguna terdaftar yang dapat mengakses layanan, sementara otorisasi memastikan bahwa hanya pengguna yang telah login yang memiliki hak untuk mengubah data yang ada pada foto.

## Fitur-fitur utama API ini meliputi:
1. Autentikasi Pengguna: Proses ini memverifikasi identitas pengguna sebelum mereka dapat mengakses API. Pengguna harus menyediakan kredensial yang valid untuk login.
2. Otorisasi Pengguna: Setelah login, pengguna diberikan izin tertentu. Hanya pengguna yang sudah terautentikasi yang dapat melakukan perubahan data pada foto.
3. Manajemen Data Foto: Pengguna dapat mengunggah, mengubah, dan menghapus data foto, dengan batasan akses yang ditentukan oleh status login mereka.

API ini dibangun dengan mempertimbangkan kemudahan penggunaan dan keamanan, sehingga memastikan bahwa data yang diolah tetap aman dan hanya diakses oleh pihak yang berwenang.

# Cara Menjalankan API

Pastikan Anda telah menginstal **Go** dan **PostgreSQL** pada komputer Anda. Berikut adalah langkah-langkah untuk menjalankan API:

## 1. Clone Repositori

Pertama, clone repositori proyek ini dengan menggunakan perintah git berikut:

```sh
git clone https://github.com/your-username/final-task-pbi-rakamin-fullstack-Muhamad-Rafi-Rahmat-Alghifari.git
```

## 2. Masuk ke Direktori Proyek

Buka terminal dan navigasikan ke direktori proyek yang baru saja Anda clone:

```sh
cd final-task-pbi-rakamin-fullstack-Muhamad-Rafi-Rahmat-Alghifari
```

## 3. Buat File .env

Di dalam direktori proyek, buat file bernama `.env` dan isi dengan konfigurasi database PostgreSQL Anda. Berikut adalah contoh isi file `.env`:

```plaintext
PGHOST=
PGPORT=
PGUSER=
PGPASSWORD=
PGDATABASE=

PORT=
SECRET_KEY=
```


## 4. Instal Dependensi

Jalankan perintah berikut untuk menginstal semua dependensi yang dibutuhkan oleh proyek:

```sh
go mod tidy
```

Perintah ini akan membersihkan dan memperbarui modul yang digunakan dalam proyek, memastikan semua dependensi yang dibutuhkan telah terinstal.

## 5. Jalankan API

Setelah semua dependensi terinstal, Anda dapat menjalankan API dengan perintah berikut:

```sh
go run main.go
```

Setelah menjalankan perintah ini, API akan berjalan pada `localhost` dengan port default `8080`. Anda dapat mengakses API melalui URL berikut:

```
http://localhost:8080
```

[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://app.getpostman.com/run-collection/15816622-9595333d-034f-4410-b3ed-7a3f3f283eb5?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D15816622-9595333d-034f-4410-b3ed-7a3f3f283eb5%26entityType%3Dcollection%26workspaceId%3Df66cecf2-0e26-4001-9b0a-3c5e5fb47c8b#?env%5BENV_V1%5D=W3sia2V5IjoiQkFTRV9VUkwiLCJ2YWx1ZSI6ImxvY2FsaG9zdDozMDAwIiwiZW5hYmxlZCI6dHJ1ZSwidHlwZSI6ImRlZmF1bHQiLCJzZXNzaW9uVmFsdWUiOiJsb2NhbGhvc3Q6MzAwMCIsInNlc3Npb25JbmRleCI6MH0seyJrZXkiOiJUT0tFTiIsInZhbHVlIjoiZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SmxiV0ZwYkNJNklubGhlV0YwUUdWNFlXMXdiR1V1WTI5dElpd2lhV1FpT2pGOS50UHk4dDJzZ25QWHZUV3BaOEp3ZjRicXBDOXpSaWh3MmJISXdfSmxRLTcwIiwiZW5hYmxlZCI6dHJ1ZSwidHlwZSI6ImRlZmF1bHQiLCJzZXNzaW9uVmFsdWUiOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKbGJXRnBiQ0k2SW5saGVXRjBRR1Y0WVcxd2JHVXVZMjl0SWl3aWFXUWlPakY5LnRQeTh0MnNnblBYdlRXcFo4SndmNGJxcEM5elJpaHcyYkhJd19KbFEtNzAiLCJzZXNzaW9uSW5kZXgiOjF9XQ==)
