<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Requests\StoreMatakuliahRequest;
use App\Http\Requests\UpdateMatakuliahRequest;
use App\Http\Resources\MatakuliahResource;
use App\Models\MataKuliah;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    public function index(Request $request)
    {
        $kueri = MataKuliah::query();

        if ($request->filled('cari')) {
            $kataKunci = $request->query('cari');

            $kueri->where(function ($sub) use ($kataKunci) {
                $sub->where('kode', 'like', '%' . $kataKunci . '%')
                    ->orWhere('nama', 'like', '%' . $kataKunci . '%');
            });
        }

        $urutan = $request->query('urut', 'kode');
        $arah = $request->query('arah', 'asc');

        $kolomDiizinkan = [
            'kode',
            'nama',
            'sks',
            'semester',
        ];

        if (in_array($urutan, $kolomDiizinkan, true)) {
            $kueri->orderBy(
                $urutan,
                $arah === 'desc' ? 'desc' : 'asc'
            );
        }

        $perHalaman = min(
            $request->integer('per_halaman', 10),
            100
        );

        return MatakuliahResource::collection(
            $kueri->paginate($perHalaman)
        );
    }

    public function store(StoreMatakuliahRequest $request): JsonResponse
    {
        $matakuliah = MataKuliah::create(
            $request->validated()
        );

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mata kuliah berhasil dibuat',
            'data' => new MatakuliahResource($matakuliah),
        ], 201);
    }

    public function show(MataKuliah $matakuliah): JsonResponse
    {
        return response()->json([
            'sukses' => true,
            'data' => new MatakuliahResource($matakuliah),
        ]);
    }

    public function update(
        UpdateMatakuliahRequest $request,
        MataKuliah $matakuliah
    ): JsonResponse {
        $matakuliah->update(
            $request->validated()
        );

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mata kuliah berhasil diperbarui',
            'data' => new MatakuliahResource($matakuliah),
        ]);
    }

    public function destroy(MataKuliah $matakuliah): JsonResponse
    {
        $matakuliah->delete();

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mata kuliah berhasil dihapus',
        ]);
    }
}