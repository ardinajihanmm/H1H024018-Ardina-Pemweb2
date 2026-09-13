<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\MatakuliahController;

// Rute Utama
Route::get('/', function () {
    return view('welcome');
});

// Langkah 1: Rute Dasar
Route::get('/salam', function () {
    return 'Selamat datang di Pemrograman Web II';
});

// Langkah 2: Rute dengan Parameter Wajib
Route::get('/mahasiswa/{nim}', function (string $nim) { 
    return 'Data mahasiswa dengan NIM ' . $nim; 
}); 

// Langkah 2: Rute dengan Parameter Opsional
Route::get('/matakuliah/{kode?}', function (?string $kode = null) { 
    if ($kode === null) { 
        return 'Menampilkan seluruh matakuliah'; 
    } 

    return 'Menampilkan matakuliah kode ' . $kode; 
});

// Langkah 3: Membatasi Format Parameter (Hanya Angka)
Route::get('/semester/{angka}', function (string $angka) { 
    return 'Semester ke ' . $angka; 
})->whereNumber('angka');

Route::get('/data-mahasiswa', [MahasiswaController::class, 
'index'])->name('mahasiswa.index'); 
 
Route::get('/data-mahasiswa/{nim}', [MahasiswaController::class, 
'show'])->name('mahasiswa.show');

Route::get('/cari-mahasiswa', [MahasiswaController::class, 'cari']); 

// Rute Tugas Praktikum Matakuliah
Route::get('/data-matakuliah', [MatakuliahController::class, 'index'])->name('matakuliah.index');
Route::get('/data-matakuliah/{kode}', [MatakuliahController::class, 'show'])->name('matakuliah.show');