<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class UpdateMatakuliahRequest extends FormRequest
{
    public function authorize(): bool
    {
        return true;
    }

    public function rules(): array
    {
        $id = $this->route('matakuliah')->id;

        return [
            'kode' => [
                'sometimes',
                'string',
                'max:20',
                'unique:matakuliahs,kode,' . $id,
            ],

            'nama' => [
                'sometimes',
                'string',
                'max:100',
            ],

            'sks' => [
                'sometimes',
                'integer',
                'min:1',
                'max:6',
            ],

            'semester' => [
                'sometimes',
                'integer',
                'min:1',
                'max:8',
            ],
        ];
    }

    public function messages(): array
    {
        return [
            'kode.unique' => 'Kode mata kuliah sudah terdaftar',
            'sks.min' => 'SKS minimal 1',
            'sks.max' => 'SKS maksimal 6',
            'semester.min' => 'Semester minimal 1',
            'semester.max' => 'Semester maksimal 8',
        ];
    }
}