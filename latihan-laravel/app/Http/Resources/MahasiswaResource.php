<?php

namespace App\Http\Resources;

use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class MahasiswaResource extends JsonResource
{
    public function toArray(Request $request): array
    {
        $fields = $request->query('fields');

        if (!$fields) {
            return [
                'id' => $this->id,
                'nim' => $this->nim,
                'nama' => $this->nama,
                'email' => $this->email,
                'angkatan' => $this->angkatan,
                'ipk' => (float) $this->ipk,
                'aktif' => $this->aktif,
                'program_studi' => $this->whenLoaded(
                    'programStudi',
                    function () {
                        return [
                            'id' => $this->programStudi->id,
                            'kode' => $this->programStudi->kode,
                            'nama' => $this->programStudi->nama,
                        ];
                    }
                ),
                'dibuat_pada' => $this->created_at?->toIso8601String(),
            ];
        }

        $fieldsDiminta = array_filter(
            array_map('trim', explode(',', $fields))
        );

        $data = [];

        if (in_array('id', $fieldsDiminta, true)) {
            $data['id'] = $this->id;
        }

        if (in_array('nim', $fieldsDiminta, true)) {
            $data['nim'] = $this->nim;
        }

        if (in_array('nama', $fieldsDiminta, true)) {
            $data['nama'] = $this->nama;
        }

        if (in_array('email', $fieldsDiminta, true)) {
            $data['email'] = $this->email;
        }

        if (in_array('angkatan', $fieldsDiminta, true)) {
            $data['angkatan'] = $this->angkatan;
        }

        if (in_array('ipk', $fieldsDiminta, true)) {
            $data['ipk'] = (float) $this->ipk;
        }

        if (in_array('aktif', $fieldsDiminta, true)) {
            $data['aktif'] = $this->aktif;
        }

        if (in_array('program_studi', $fieldsDiminta, true)) {
            $data['program_studi'] = $this->whenLoaded(
                'programStudi',
                function () {
                    return [
                        'id' => $this->programStudi->id,
                        'kode' => $this->programStudi->kode,
                        'nama' => $this->programStudi->nama,
                    ];
                }
            );
        }

        if (in_array('dibuat_pada', $fieldsDiminta, true)) {
            $data['dibuat_pada'] = $this->created_at?->toIso8601String();
        }

        return $data;
    }
}