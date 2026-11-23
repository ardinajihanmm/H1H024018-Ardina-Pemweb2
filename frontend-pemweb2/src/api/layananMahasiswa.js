import klien from './klien'
import { backend } from './konfigurasi'
import { parameterDaftar, normalisasiDaftar } from './kontrak'

export async function ambilDaftarMahasiswa(filter, opsi = {}) {
  const respons = await klien.get('/mahasiswa', {
    params: parameterDaftar(filter, backend),
    signal: opsi.signal,
  })
  return normalisasiDaftar(respons.data, filter.halaman)
}

export async function ambilDetailMahasiswa(id) {
  const respons = await klien.get(`/mahasiswa/${id}`)
  return respons.data
}

export async function simpanMahasiswa(data) {
  const respons = await klien.post('/mahasiswa', data)
  return respons.data
}

export async function hapusMahasiswa(id) {
  const respons = await klien.delete(`/mahasiswa/${id}`)
  return respons.data
}
