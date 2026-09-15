<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use App\Models\Mahasiswa;

class MataKuliah extends Model
{
    use HasFactory;

    protected $table = 'matakuliahs';

    protected $fillable = [
        'kode',
        'nama',
        'sks',
        'semester',
    ];

    public function mahasiswas()
    {
        return $this->belongsToMany(
            Mahasiswa::class,
            'mahasiswa_matakuliah'
        )->withPivot('nilai');
    }
}