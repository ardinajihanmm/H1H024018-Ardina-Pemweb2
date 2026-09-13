<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    private function getDaftarMatakuliah()
    {
        return [
            ['kode' => 'TKO101', 'nama' => 'Pemrograman Web II', 'sks' => 3],
            ['kode' => 'TKO102', 'nama' => 'Algoritma dan Struktur Data', 'sks' => 4],
            ['kode' => 'TKO103', 'nama' => 'Praktikum Sistem Operasi', 'sks' => 2],
            ['kode' => 'TKO104', 'nama' => 'Jaringan Komputer', 'sks' => 3],
            ['kode' => 'TKO105', 'nama' => 'Bahasa Inggris', 'sks' => 2],
        ];
    }

    public function index(Request $request)
    {
        $daftarMatakuliah = $this->getDaftarMatakuliah();
        $q = $request->query('q', '');

        // Fitur Pencarian Query String
        if (!empty($q)) {
            $daftarMatakuliah = array_filter($daftarMatakuliah, function ($item) use ($q) {
                return stripos($item['nama'], $q) !== false || stripos($item['kode'], $q) !== false;
            });
        }

        return view('matakuliah.index', [
            'daftarMatakuliah' => $daftarMatakuliah,
            'q' => $q,
        ]);
    }

    public function show(string $kode)
    {
        $daftarMatakuliah = $this->getDaftarMatakuliah();
        $matakuliah = null;

        foreach ($daftarMatakuliah as $item) {
            if ($item['kode'] === $kode) {
                $matakuliah = $item;
                break;
            }
        }

        return view('matakuliah.show', [
            'kode' => $kode,
            'matakuliah' => $matakuliah
        ]);
    }
}