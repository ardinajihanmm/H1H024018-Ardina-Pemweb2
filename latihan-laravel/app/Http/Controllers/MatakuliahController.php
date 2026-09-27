<?php

namespace App\Http\Controllers;

use App\Models\MataKuliah;

class MatakuliahController extends Controller
{
    public function index()
    {
        $matakuliah = MataKuliah::all();

        return view('matakuliah.index', compact('matakuliah'));
    }

    public function show(MataKuliah $matakuliah)
    {
        return view('matakuliah.show', compact('matakuliah'));
    }
}