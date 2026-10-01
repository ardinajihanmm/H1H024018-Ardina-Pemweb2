<?php

use App\Http\Controllers\Api\AuthController;
use App\Http\Controllers\Api\MahasiswaController;
use App\Http\Controllers\Api\MatakuliahController;
use App\Http\Controllers\Api\ProgramStudiController;
use Illuminate\Support\Facades\Route;

Route::get('/status', function () {
    return response()->json([
        'sukses' => true,
        'pesan' => 'API Pemweb II aktif',
        'waktu' => now()->toIso8601String(),
    ]);
});

// =========================
// AUTENTIKASI
// =========================

Route::post('/auth/register', [AuthController::class, 'register']);

Route::post('/auth/login', [AuthController::class, 'login'])
    ->middleware('throttle:5,1');

// =========================
// ROUTE TERPROTEKSI
// =========================

Route::middleware('auth:sanctum')->group(function () {

    // Profil
    Route::get('/auth/profil', [AuthController::class, 'profil']);

    // Ubah password
    Route::put('/auth/password', [AuthController::class, 'ubahPassword']);

    // Logout
    Route::post('/auth/logout', [AuthController::class, 'logout']);

    // Logout semua perangkat
    Route::post('/auth/logout-semua', [AuthController::class, 'logoutSemua']);

    // =========================
    // MAHASISWA - BACA
    // =========================

    Route::get('/mahasiswa', [MahasiswaController::class, 'index']);

    Route::get('/mahasiswa/{mahasiswa}', [MahasiswaController::class, 'show']);

    // =========================
    // MAHASISWA - TULIS
    // =========================

    Route::middleware('ability:mahasiswa:tulis')->group(function () {

        Route::post('/mahasiswa', [MahasiswaController::class, 'store']);

        Route::put('/mahasiswa/{mahasiswa}', [MahasiswaController::class, 'update']);

        Route::patch('/mahasiswa/{mahasiswa}', [MahasiswaController::class, 'update']);
    });

    // =========================
    // HAPUS - KHUSUS ADMIN
    // =========================

    Route::middleware('peran.admin')->group(function () {

        Route::delete(
            '/mahasiswa/{mahasiswa}',
            [MahasiswaController::class, 'destroy']
        );
    });

    // =========================
    // MATA KULIAH
    // =========================

    Route::apiResource('matakuliah', MatakuliahController::class);

    // =========================
    // PROGRAM STUDI
    // =========================

    Route::get(
        '/program-studi/{programStudi}/mahasiswa',
        [ProgramStudiController::class, 'mahasiswa']
    );
});