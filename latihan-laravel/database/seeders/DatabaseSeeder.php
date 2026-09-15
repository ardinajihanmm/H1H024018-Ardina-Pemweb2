<?php

namespace Database\Seeders;

use App\Models\User;
use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;
use App\Models\Mahasiswa;

class DatabaseSeeder extends Seeder
{
    use WithoutModelEvents;

    /**
     * Seed the application's database.
     */
    public function run(): void
    {
        // Membuat data program studi
        $this->call(ProgramStudiSeeder::class);

        // Membuat 30 data mahasiswa
        Mahasiswa::factory()->count(30)->create();

        // Membuat data mata kuliah
        $this->call(MataKuliahSeeder::class);

        // Membuat relasi mahasiswa dengan mata kuliah
        $this->call(MahasiswaMataKuliahSeeder::class);
    }
}