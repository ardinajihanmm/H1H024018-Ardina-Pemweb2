# Cara Mengerjakan & Mengecek Modul 10 — VS Code / PowerShell

**Ringkasnya:** Modul 10 meminta 10 langkah: OpenAPI, tes Laravel, tes Fiber, konfigurasi produksi Laravel, build Fiber untuk Linux, systemd, Nginx, build frontend, GitHub Actions, dan checklist keamanan. Proyek akhir (bagian E pada PDF) adalah tugas tambahan yang terpisah dan berkelompok.

## 0. Salin folder tanpa merusak proyekmu

ZIP ini memuat `latihan-laravel`, `latihan-fiber-api`, `frontend-pemweb2` dan file di akar repositori. Simpan backup proyek lama sebelum menggabungkan. **Jangan timpa `.env` yang sudah jalan dengan file contoh.** File rahasia, `.git`, `vendor`, `node_modules`, dan database lokal sengaja tidak disertakan.

## 1. Cek dokumentasi OpenAPI

File `openapi.yaml` di akar folder adalah dokumentasi gabungan dua backend berdasarkan rute yang ada. Buka di VS Code dengan ekstensi pratinjau OpenAPI atau tempel isinya di editor Swagger. Pastikan daftar endpoint /auth/login, /mahasiswa, /matakuliah muncul. Perhatikan bahwa Laravel (`password`, `page`, Sanctum) berbeda dari Fiber (`kata_sandi`, `halaman`, JWT).

## 2. Tes Laravel (penting: database TEST, bukan data asli!)

Pastikan `latihan-laravel/.env` lokal masih sesuai proyekmu. File `phpunit.xml` proyek sudah memakai `DB_CONNECTION=sqlite` dan `DB_DATABASE=:memory:`. PHP harus memiliki ekstensi `pdo_sqlite`. Dari PowerShell:

```powershell
cd "D:\laragon\Praktikum Pemweb II\latihan-laravel"
php -m
php artisan test --filter=MahasiswaApiTest
php artisan test
```

File yang ditambah: `tests/Feature/MahasiswaApiTest.php` dengan 7 tes. Jangan menyalin `DB_DATABASE` produksi ke `.env.testing`.

## 3. Tes Fiber

```powershell
cd "D:\laragon\Praktikum Pemweb II\latihan-fiber-api"
go version
go test ./...
go test ./... -cover
go vet ./...
```

Kebutuhan `go.mod` proyek asli: Go **1.27.0**. Pastikan toolchain itu tersedia. Tes Fiber: endpoint status asli, profil tanpa token, ID bukan angka, NIM tidak valid, serta 3 tes pola NIM. Semua dibuat tanpa membutuhkan MySQL untuk skenario yang diuji.

## 4. Tes frontend dan build

```powershell
cd "D:\laragon\Praktikum Pemweb II\frontend-pemweb2"
npm install
npm test
npm run build
```

Jika selesai, folder `dist` menjadi hasil frontend produksi. Jangan memasukkan secret ke variabel `VITE_` karena nilainya akan terbaca di browser.

## 5. Periksa GitHub Actions

File `.github/workflows/uji.yml` ada di akar folder proyek. Agar otomatis berjalan, unggah susunan folder yang sama ke **satu repo GitHub** dan lakukan push. Periksa tab Actions. Workflow memiliki 3 job: Laravel, Fiber, React. Versi Go otomatis dibaca dari `go.mod`, sehingga tidak memakai angka `1.23` yang ada pada contoh PDF tetapi tidak cocok dengan proyekmu.

## 6. Deployment lab / VM

Ikuti `deployment/README_DEPLOYMENT.md` untuk build Linux, service systemd, reverse proxy Nginx, SSL, dan perbedaan konfigurasi produksi. **Tidak bisa membuktikan deployment hanya dari file konfigurasi; harus diuji di VM/VPS.**

## 7. Screenshot yang dimasukkan ke laporan

1. `openapi.yaml` berhasil dirender di Swagger/preview.
2. Terminal Laravel `php artisan test` dengan jumlah tes dan statusnya.
3. Terminal Go `go test ./... -cover` dengan coverage.
4. Terminal React `npm test` dan `npm run build`.
5. Jika ada VM: `systemctl status`, `nginx -t`, dan curl melalui Nginx.
6. GitHub Actions tab Jobs ketika workflow selesai.

## 8. Bukti pengujian yang sudah bisa diverifikasi di lingkungan penyusunan

- `npm test` React: **7 lulus, 0 gagal**.
- File tes PHP sudah lolos pemeriksaan sintaks dengan `php -l`.
- Go source test sudah diformat dengan `gofmt` tetapi **tidak dapat dieksekusi di sini** karena Go yang tersedia 1.23.2, sementara `go.mod` memerlukan 1.27.0.
- Laravel PHPUnit belum bisa dieksekusi di sini karena ekstensi PDO SQLite tidak tersedia. Jalankan pada Laragon dengan ekstensi tersebut.
- Belum ada bukti server Linux, systemd, Nginx, atau HTTPS hidup; konfigurasinya baru **template deployment**.

## 9. Ketentuan proyek akhir terpisah (Bagian E)

Satu kelompok 2–3 mahasiswa, minimal 3 entitas yang saling berelasi, Laravel Sanctum sebagai REST API utama, Fiber untuk minimal satu fungsi domain tambahan, frontend yang mengonsumsi kedua layanan, OpenAPI, minimal 5 uji otomatis, riwayat kontribusi Git, README, Postman collection, laporan maksimal 15 halaman dan presentasi 15 menit.

**Catatan jujur:** ZIP ini menyelesaikan persiapan dan implementasi berkas Modul 10 pada proyek yang kamu unggah, **belum menyelesaikan seluruh proyek akhir**, sebab frontend Modul 9 masih menggunakan pilihan satu backend pada satu waktu. Agar memenuhi proyek akhir, frontend perlu benar-benar memakai dua layanan untuk fungsi berbeda dalam satu alur, plus bukti kontribusi kelompok dan demo.
