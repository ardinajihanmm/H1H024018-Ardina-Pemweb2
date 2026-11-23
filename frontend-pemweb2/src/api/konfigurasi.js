// Semua pilihan backend dikendalikan oleh SATU konfigurasi: VITE_BACKEND.
// Jangan simpan secret JWT, API key, atau password pada variabel VITE_.
export const backend = (import.meta.env.VITE_BACKEND || 'fiber').toLowerCase()

if (!['laravel', 'fiber'].includes(backend)) {
  throw new Error('VITE_BACKEND harus bernilai fiber atau laravel')
}

export const baseURL = backend === 'laravel'
  ? import.meta.env.VITE_API_LARAVEL
  : import.meta.env.VITE_API_FIBER

if (!baseURL) {
  throw new Error(`URL API ${backend} belum diatur pada .env`)
}
