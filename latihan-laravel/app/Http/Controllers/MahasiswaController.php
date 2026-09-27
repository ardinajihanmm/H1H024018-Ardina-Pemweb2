<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MahasiswaController extends Controller
{
    public function index()
    {
        $daftarMahasiswa = [
            [
                'nim' => 'H1A123001',
                'nama' => 'Andi Pratama',
                'angkatan' => 2023,
            ],
            [
                'nim' => 'H1A123002',
                'nama' => 'Budi Santoso',
                'angkatan' => 2023,
            ],
            [
                'nim' => 'H1A123003',
                'nama' => 'Citra Lestari',
                'angkatan' => 2023,
            ],
        ];

        return view('mahasiswa.index', compact('daftarMahasiswa'));
    }

    public function show($nim)
    {
        return view('mahasiswa.show', compact('nim'));
    }

    public function cari(Request $request)
    {
        $kataKunci = $request->query('cari');

        $daftarMahasiswa = [
            [
                'nim' => 'H1A123001',
                'nama' => 'Andi Pratama',
                'angkatan' => 2023,
            ],
            [
                'nim' => 'H1A123002',
                'nama' => 'Budi Santoso',
                'angkatan' => 2023,
            ],
            [
                'nim' => 'H1A123003',
                'nama' => 'Citra Lestari',
                'angkatan' => 2023,
            ],
        ];

        if ($kataKunci) {
            $daftarMahasiswa = array_filter(
                $daftarMahasiswa,
                function ($mahasiswa) use ($kataKunci) {
                    return str_contains(
                        strtolower($mahasiswa['nama']),
                        strtolower($kataKunci)
                    ) ||
                    str_contains(
                        strtolower($mahasiswa['nim']),
                        strtolower($kataKunci)
                    );
                }
            );
        }

        return view('mahasiswa.index', compact('daftarMahasiswa'));
    }
}