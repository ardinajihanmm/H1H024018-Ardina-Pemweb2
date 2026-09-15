<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use App\Models\MataKuliah;

class MataKuliahSeeder extends Seeder
{
    public function run(): void
    {
        $daftar = [
            ['kode' => 'TK101', 'nama' => 'Dasar Sistem Komputer', 'sks' => 3, 'semester' => 1],
            ['kode' => 'TK102', 'nama' => 'Jaringan Komputer', 'sks' => 3, 'semester' => 2],
            ['kode' => 'TK103', 'nama' => 'Pemrograman Web', 'sks' => 3, 'semester' => 3],
            ['kode' => 'TK104', 'nama' => 'Basis Data', 'sks' => 3, 'semester' => 2],
            ['kode' => 'TK105', 'nama' => 'Algoritma dan Pemrograman', 'sks' => 3, 'semester' => 1],
        ];

        foreach ($daftar as $item) {
            MataKuliah::create($item);
        }
    }
}