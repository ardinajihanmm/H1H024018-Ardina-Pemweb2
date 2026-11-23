import axios from 'axios'
import { baseURL, backend } from './konfigurasi'

const klien = axios.create({
  baseURL,
  timeout: 10000,
  headers: { Accept: 'application/json', 'Content-Type': 'application/json' },
})

klien.interceptors.request.use((konfigurasi) => {
  const token = sessionStorage.getItem(`token_${backend}`)
  if (token) konfigurasi.headers.Authorization = `Bearer ${token}`
  return konfigurasi
})

klien.interceptors.response.use(
  (respons) => respons,
  (galat) => {
    const status = galat.response?.status
    const url = String(galat.config?.url || '')
    if (status === 401 && !url.includes('/auth/login')) {
      sessionStorage.removeItem(`token_${backend}`)
      sessionStorage.removeItem(`pengguna_${backend}`)
      if (window.location.pathname !== '/login') window.location.replace('/login')
    }
    return Promise.reject(galat)
  },
)

export default klien
