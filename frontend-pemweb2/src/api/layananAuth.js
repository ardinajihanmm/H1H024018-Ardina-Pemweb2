import klien from './klien'
import { backend } from './konfigurasi'
import { payloadLogin, normalisasiLogin } from './kontrak'

export async function login(email, kataSandi) {
  const respons = await klien.post('/auth/login', payloadLogin(email, kataSandi, backend))
  return normalisasiLogin(respons.data)
}

export async function ambilProfil() {
  const respons = await klien.get('/auth/profil')
  return respons.data
}

export async function logout() {
  // Kedua modul memakai POST /api/auth/logout; Laravel mencabut token Sanctum, Fiber memblokir JWT.
  const respons = await klien.post('/auth/logout')
  return respons.data
}
