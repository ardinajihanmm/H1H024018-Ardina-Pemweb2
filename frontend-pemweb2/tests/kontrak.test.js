import test from 'node:test'
import assert from 'node:assert/strict'
import {
  normalisasiDaftar, normalisasiGalat, normalisasiLogin,
  parameterDaftar, payloadLogin, peranPengguna,
} from '../src/api/kontrak.js'

test('payload login Fiber menggunakan kata_sandi; Laravel menggunakan password', () => {
  assert.deepEqual(payloadLogin('a@contoh.com', '12345678', 'fiber'), { email: 'a@contoh.com', kata_sandi: '12345678' })
  assert.deepEqual(payloadLogin('a@contoh.com', '12345678', 'laravel'), { email: 'a@contoh.com', password: '12345678' })
})

test('parameter pagination sesuai backend', () => {
  assert.deepEqual(parameterDaftar({ halaman: 2, cari: 'Budi', perHalaman: 10 }, 'fiber'), { halaman: 2, cari: 'Budi', per_halaman: 10 })
  assert.deepEqual(parameterDaftar({ halaman: 2, cari: 'Budi', perHalaman: 10 }, 'laravel'), { page: 2, cari: 'Budi', per_halaman: 10 })
})

test('list Fiber dinormalisasi', () => {
  const r = normalisasiDaftar({ data: [{ id: 1 }], meta: { halaman: 3, total_halaman: 7, total: 64 } })
  assert.deepEqual(r.daftar, [{ id: 1 }])
  assert.deepEqual(r.meta, { halaman: 3, total_halaman: 7, total: 64 })
})

test('list Laravel dinormalisasi termasuk paginator bersarang', () => {
  const r = normalisasiDaftar({ data: { data: [{ id: 2 }], current_page: 2, last_page: 4, total: 38 } })
  assert.deepEqual(r.daftar, [{ id: 2 }])
  assert.deepEqual(r.meta, { halaman: 2, total_halaman: 4, total: 38 })
})

test('Laravel dengan meta terpisah', () => {
  const r = normalisasiDaftar({ data: [{ id: 3 }], meta: { current_page: 1, last_page: 2, total: 11 } })
  assert.deepEqual(r.meta, { halaman: 1, total_halaman: 2, total: 11 })
})

test('error Fiber dan Laravel dinormalisasi', () => {
  const fiber = normalisasiGalat({ response: { status: 422, data: { pesan: 'Data tidak valid', galat: [{ field: 'Nim', pesan: 'NIM tidak sesuai' }] } } })
  assert.equal(fiber.status, 422)
  assert.deepEqual(fiber.daftar, [{ field: 'Nim', pesan: 'NIM tidak sesuai' }])
  const laravel = normalisasiGalat({ response: { status: 422, data: { message: 'Validation failed', errors: { nim: ['NIM wajib diisi'] } } } })
  assert.deepEqual(laravel.daftar, [{ field: 'nim', pesan: 'NIM wajib diisi' }])
})

test('normalisasi hasil login dan role', () => {
  assert.deepEqual(normalisasiLogin({ data: { token: 'abc', pengguna: { peran: 'admin' } } }), { token: 'abc', pengguna: { peran: 'admin' }, pesan: 'Login berhasil' })
  assert.equal(peranPengguna({ role: 'Admin' }), 'admin')
  assert.throws(() => normalisasiLogin({ data: {} }), /tidak memuat token/)
})
