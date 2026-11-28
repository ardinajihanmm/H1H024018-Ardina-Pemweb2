# Panduan Deployment Modul 10 (Bahasa Mahasiswa)

> **Status:** Konfigurasi *siap dicoba* di VM Linux, **belum ada deployment publik yang berhasil diverifikasi**. File ini bukan bukti server sudah online.

## A. Persiapan dan keamanan

1. Backup database lebih dulu. Uji semua perubahan pada database uji, jangan database kuliah yang dipakai sehari-hari.
2. Gunakan VM Ubuntu/Linux yang terpisah. Jangan simpan `.env` asli di GitHub, ZIP tugas, atau screenshot. Kunci JWT, kunci API, dan sandi DB harus unik di mesin server.
3. Jangan membagikan port MySQL (3306) ke internet. Jalankan Fiber pada `127.0.0.1:3000` bila memungkinkan; gunakan firewall untuk membatasi akses port 3000.
4. Gunakan HTTPS untuk rilis sungguhan; contoh HTTP `pemweb2-lab.conf` hanya untuk pengujian LAN/VM tertutup.

## B. Laravel 13

Di dalam `latihan-laravel` pada VM, buat file `.env` **sendiri** berdasar `.env.production.example`. Isi APP_KEY dengan `php artisan key:generate` (jangan memakai teks contoh). Jalankan:

```bash
composer install --no-dev --optimize-autoloader
php artisan config:cache
php artisan route:cache
php artisan view:cache
php artisan migrate --force
```

**Catatan:** `php artisan serve` hanya untuk demonstrasi lokal/VM; aplikasi produksi sebaiknya memakai PHP-FPM dan root Nginx di folder `public/`.

## C. Fiber v3 — membangun binary Linux

Jika menjalankan build di Windows PowerShell:

```powershell
cd "D:\\laragon\\Praktikum Pemweb II\\latihan-fiber-api"
New-Item -ItemType Directory -Force bin | Out-Null
$env:CGO_ENABLED = "0"
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -ldflags="-s -w" -o bin/api ./cmd/api
Remove-Item Env:GOOS,Env:GOARCH,Env:CGO_ENABLED -ErrorAction SilentlyContinue
```

Binary Linux hanya dapat *dijalankan* pada server Linux yang cocok dengan arsitektur target. Alternatif build langsung di VM:

```bash
mkdir -p bin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/api ./cmd/api
```

Pindahkan binary dan file kerja ke `/var/www/pemweb2-api`. Buat `/etc/pemweb2/fiber.env` yang berisi variabel seperti `.env.production.example` dengan nilai rahasia yang sebenarnya (izin berkas root: `chmod 600`). Jangan mengunggah file contoh yang belum diganti ke server publik.

Salin `deployment/systemd/pemweb2-fiber.service` ke `/etc/systemd/system/`, lalu jalankan:

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now pemweb2-fiber
sudo systemctl status pemweb2-fiber
sudo journalctl -u pemweb2-fiber -n 50 --no-pager
```

**Penting:** kode Fiber saat ini melakukan koneksi database dan migrasi otomatis ketika startup. Uji pada database cadangan sebelum menjalankannya di produksi. Pastikan pengguna `www-data` dapat membaca working directory yang dibutuhkan.

## D. Nginx reverse proxy

Untuk demo VM lokal, gunakan `deployment/nginx/pemweb2-lab.conf`. Salin ke `/etc/nginx/sites-available/pemweb2` kemudian:

```bash
sudo ln -s /etc/nginx/sites-available/pemweb2 /etc/nginx/sites-enabled/pemweb2
sudo nginx -t
sudo systemctl reload nginx
```

Uji dengan `curl -i http://IP_VM/fiber/api/status`. Konfigurasi produksi di `pemweb2-production.conf.example` baru dapat dipakai **setelah** DNS, TLS, PHP-FPM, serta izin direktori benar-benar disiapkan. Subpath `/laravel/` membutuhkan pengaturan URL Laravel dan routing yang cocok, sehingga uji semua endpoint sebelum dipakai publik.

## E. Frontend React

Sebelum `npm run build`, salin `.env.production.example` menjadi `.env.production` dan sesuaikan URL API **HTTPS** dengan server aktif. Kemudian:

```bash
npm ci
npm test
npm run build
```

Folder `dist/` adalah hasil build statis. Hidangkan melalui Nginx, bukan dengan Vite dev server. Pastikan konfigurasi server mendukung SPA fallback ke `index.html`.

## F. Bukti yang harus benar-benar dibuat

- `php artisan test` Laravel (hasil lulus dan gagal)
- `go test ./... -cover` Fiber
- `npm test` dan `npm run build` React
- `sudo nginx -t` serta `systemctl status pemweb2-fiber` pada VM
- Hasil `curl -i` melalui reverse proxy, dan HTTPS (bila publik)

Jangan menulis “deployment berhasil” dalam laporan sebelum layanan benar-benar dapat diakses sesuai target.
