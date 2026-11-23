export function payloadLogin(email, kataSandi, backend) {
  return backend === 'fiber'
    ? { email, kata_sandi: kataSandi }
    : { email, password: kataSandi }
}

export function parameterDaftar({ halaman = 1, cari = '', perHalaman = 10 }, backend) {
  return backend === 'fiber'
    ? { halaman, cari, per_halaman: perHalaman }
    : { page: halaman, cari, per_halaman: perHalaman }
}

export function normalisasiLogin(hasil) {
  const data = hasil?.data || hasil || {}
  const token = data.token || data.access_token || hasil?.token || hasil?.access_token
  if (!token || typeof token !== 'string') {
    throw new Error('Respons login tidak memuat token. Periksa kontrak API backend.')
  }
  const pengguna = data.pengguna || data.user || hasil?.user || null
  return { token, pengguna, pesan: hasil?.pesan || hasil?.message || 'Login berhasil' }
}

export function normalisasiDaftar(hasil, halamanFallback = 1) {
  // Mendukung {data:[],meta:{}}, {data:{data:[],...pagination}}, atau {data:[]}.
  const isi = hasil?.data
  const daftar = Array.isArray(isi)
    ? isi
    : Array.isArray(isi?.data)
      ? isi.data
      : []
  const metadata = hasil?.meta || (isi && !Array.isArray(isi) ? isi : {}) || {}
  const halaman = Number(metadata.halaman ?? metadata.current_page ?? halamanFallback)
  const totalHalaman = Number(metadata.total_halaman ?? metadata.last_page ?? 1)
  const total = Number(metadata.total ?? daftar.length)
  return {
    daftar,
    meta: {
      halaman: Number.isFinite(halaman) && halaman > 0 ? halaman : 1,
      total_halaman: Number.isFinite(totalHalaman) && totalHalaman > 0 ? totalHalaman : 1,
      total: Number.isFinite(total) ? total : daftar.length,
    },
  }
}

export function normalisasiGalat(galat) {
  // Fiber: galat:[{field,pesan}]. Laravel: errors:{nim:[...]}.
  const hasil = galat?.response?.data || galat?.data || {}
  const daftar = []
  if (Array.isArray(hasil.galat)) {
    for (const item of hasil.galat) {
      daftar.push({ field: item.field || item.Field || 'data', pesan: item.pesan || item.Pesan || 'Tidak valid' })
    }
  }
  if (hasil.errors && typeof hasil.errors === 'object' && !Array.isArray(hasil.errors)) {
    for (const [field, pesan] of Object.entries(hasil.errors)) {
      for (const teks of (Array.isArray(pesan) ? pesan : [pesan])) {
        daftar.push({ field, pesan: String(teks) })
      }
    }
  }
  const pesan = hasil.pesan || hasil.message || galat?.message || 'Permintaan gagal'
  return { pesan, daftar, status: galat?.response?.status || null }
}

export function peranPengguna(pengguna) {
  return (pengguna?.peran || pengguna?.role || '').toLowerCase()
}
