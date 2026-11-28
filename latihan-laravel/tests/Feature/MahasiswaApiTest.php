<?php

namespace Tests\Feature;

use App\Models\Mahasiswa;
use App\Models\ProgramStudi;
use App\Models\User;
use Illuminate\Foundation\Testing\RefreshDatabase;
use Laravel\Sanctum\Sanctum;
use Tests\TestCase;

class MahasiswaApiTest extends TestCase
{
    use RefreshDatabase;

    private function loginSebagaiAdmin(): void
    {
        Sanctum::actingAs(
            User::factory()->create(['peran' => 'admin']),
            ['mahasiswa:baca', 'mahasiswa:tulis']
        );
    }

    public function test_tanpa_token_daftar_mahasiswa_ditolak(): void
    {
        $this->getJson('/api/mahasiswa')->assertStatus(401)
            ->assertJsonPath('sukses', false);
    }

    public function test_admin_dapat_melihat_daftar_dengan_pagination(): void
    {
        $this->loginSebagaiAdmin();

        $this->getJson('/api/mahasiswa?per_halaman=10')
            ->assertOk()
            ->assertJsonStructure([
                'data',
                'links',
                'meta' => ['current_page', 'last_page', 'total'],
            ]);
    }

    public function test_admin_dapat_membuat_mahasiswa_valid(): void
    {
        $this->loginSebagaiAdmin();
        $prodi = ProgramStudi::create([
            'kode' => 'TK', 'nama' => 'Teknik Komputer', 'jenjang' => 'S1',
        ]);

        $this->postJson('/api/mahasiswa', [
            'program_studi_id' => $prodi->id,
            'nim' => 'H125010',
            'nama' => 'Gita Permata',
            'email' => 'gita.m10@example.com',
            'angkatan' => 2025,
            'ipk' => 3.70,
            'aktif' => true,
        ])->assertCreated()->assertJsonPath('sukses', true);

        $this->assertDatabaseHas('mahasiswas', ['nim' => 'H125010']);
    }

    public function test_input_mahasiswa_tidak_valid_ditolak(): void
    {
        $this->loginSebagaiAdmin();

        $this->postJson('/api/mahasiswa', ['nama' => 'Ab'])
            ->assertStatus(422)
            ->assertJsonPath('sukses', false)
            ->assertJsonStructure(['galat']);
    }

    public function test_token_mahasiswa_tanpa_hak_tulis_ditolak(): void
    {
        Sanctum::actingAs(
            User::factory()->create(['peran' => 'mahasiswa']),
            ['mahasiswa:baca']
        );

        $this->postJson('/api/mahasiswa', ['nama' => 'Tes'])
            ->assertForbidden();
    }

    public function test_mahasiswa_tidak_boleh_menghapus_data(): void
    {
        $prodi = ProgramStudi::create([
            'kode' => 'IF', 'nama' => 'Informatika', 'jenjang' => 'S1',
        ]);
        $mahasiswa = Mahasiswa::create([
            'program_studi_id' => $prodi->id,
            'nim' => 'H125011',
            'nama' => 'Mahasiswa Contoh',
            'email' => 'contoh.m10@example.com',
            'angkatan' => 2025,
            'ipk' => 3.30,
            'aktif' => true,
        ]);
        Sanctum::actingAs(
            User::factory()->create(['peran' => 'mahasiswa']),
            ['mahasiswa:baca']
        );

        $this->deleteJson('/api/mahasiswa/'.$mahasiswa->id)->assertForbidden();
        $this->assertDatabaseHas('mahasiswas', ['id' => $mahasiswa->id]);
    }

    public function test_admin_dapat_menghapus_data(): void
    {
        $prodi = ProgramStudi::create([
            'kode' => 'TE', 'nama' => 'Teknik Elektro', 'jenjang' => 'S1',
        ]);
        $mahasiswa = Mahasiswa::create([
            'program_studi_id' => $prodi->id,
            'nim' => 'H125012',
            'nama' => 'Mahasiswa Hapus',
            'email' => 'hapus.m10@example.com',
            'angkatan' => 2025,
            'ipk' => 3.00,
            'aktif' => true,
        ]);
        $this->loginSebagaiAdmin();

        $this->deleteJson('/api/mahasiswa/'.$mahasiswa->id)
            ->assertOk()->assertJsonPath('sukses', true);
        $this->assertDatabaseMissing('mahasiswas', ['id' => $mahasiswa->id]);
    }
}
