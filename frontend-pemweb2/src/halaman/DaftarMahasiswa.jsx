import { useEffect, useState } from 'react'
import { Link, useLocation } from 'react-router-dom'
import { backend } from '../api/konfigurasi'
import { normalisasiGalat, peranPengguna } from '../api/kontrak'
import { ambilDaftarMahasiswa, hapusMahasiswa } from '../api/layananMahasiswa'

export default function DaftarMahasiswa() {
  const lokasi = useLocation()
  const [daftar, setDaftar] = useState([])
  const [meta, setMeta] = useState({ halaman: 1, total_halaman: 1, total: 0 })
  const [halaman, setHalaman] = useState(1)
  const [cari, setCari] = useState('')
  const [sedangMemuat, setSedangMemuat] = useState(true)
  const [pesanGalat, setPesanGalat] = useState('')
  const [pesanSukses, setPesanSukses] = useState(lokasi.state?.pesan || '')
  const [menghapus, setMenghapus] = useState(null)
  const [versiData, setVersiData] = useState(0)
  let pengguna = null
  try { pengguna = JSON.parse(sessionStorage.getItem(`pengguna_${backend}`) || 'null') } catch { /* cache is optional */ }
  const peran = peranPengguna(pengguna)
  // Laravel Modul 5 menerbitkan kemampuan mahasiswa:tulis hanya untuk admin.
  // Fiber Modul 8 mengizinkan admin/dosen menambah, tetapi hanya admin menghapus.
  const bolehHapus = peran === 'admin'
  const bolehTambah = peran === 'admin' || (backend === 'fiber' && peran === 'dosen')

  useEffect(() => {
    const controller = new AbortController()
    let aktif = true
    setSedangMemuat(true)
    setPesanGalat('')
    const penunda = setTimeout(async () => {
      try {
        const hasil = await ambilDaftarMahasiswa({ halaman, cari, perHalaman: 10 }, { signal: controller.signal })
        if (!aktif) return
        setDaftar(hasil.daftar)
        setMeta(hasil.meta)
      } catch (galat) {
        if (!aktif || galat.code === 'ERR_CANCELED') return
        setPesanGalat(normalisasiGalat(galat).pesan || 'Gagal memuat data mahasiswa')
      } finally {
        if (aktif) setSedangMemuat(false)
      }
    }, 400)

    return () => { aktif = false; clearTimeout(penunda); controller.abort() }
  }, [halaman, cari, versiData])

  async function tanganiHapus(mahasiswa) {
    if (!window.confirm(`Hapus mahasiswa ${mahasiswa.nama} (${mahasiswa.nim})?`)) return
    setMenghapus(mahasiswa.id)
    setPesanGalat('')
    setPesanSukses('')
    try {
      await hapusMahasiswa(mahasiswa.id)
      setPesanSukses(`Data ${mahasiswa.nama} berhasil dihapus.`)
      if (daftar.length === 1 && halaman > 1) setHalaman(h => h - 1)
      else setVersiData(v => v + 1)
    } catch (galat) {
      setPesanGalat(normalisasiGalat(galat).pesan)
    } finally {
      setMenghapus(null)
    }
  }

  return (
    <section>
      <div className="heading-row">
        <div>
          <p className="eyebrow">DASHBOARD MAHASISWA</p>
          <h1>Daftar mahasiswa</h1>
          <p className="muted">Data diambil dari REST API melalui Axios dengan autentikasi Bearer.</p>
        </div>
        {bolehTambah && <Link to="/mahasiswa/tambah" className="button button-primary">+ Tambah mahasiswa</Link>}
      </div>
      <div className="stat-row">
        <div className="stat"><strong>{meta.total}</strong><span>Total mahasiswa</span></div>
        <div className="stat"><strong>{meta.total_halaman}</strong><span>Total halaman</span></div>
        <div className="stat"><strong>{backend.toUpperCase()}</strong><span>Backend aktif</span></div>
      </div>
      <div className="paper">
        <div className="toolbar">
          <label htmlFor="cari" className="sr-only">Cari mahasiswa</label>
          <input id="cari" type="search" placeholder="Cari nama atau NIM..." value={cari} onChange={e => { setHalaman(1); setCari(e.target.value) }} />
          <button className="button button-outline" type="button" onClick={() => setVersiData(v => v + 1)}>Muat ulang</button>
        </div>
        {pesanSukses && <div className="alert alert-success" role="status">{pesanSukses}</div>}
        {pesanGalat && <div className="alert alert-error" role="alert">{pesanGalat}</div>}
        {sedangMemuat ? <div role="status" className="empty">Memuat data mahasiswa...</div> : (
          daftar.length === 0 ? <div className="empty">Data tidak ditemukan.</div> : (
            <div className="table-scroll">
              <table>
                <thead><tr><th>NIM</th><th>Nama</th><th>Angkatan</th><th>IPK</th><th>Status</th><th>Aksi</th></tr></thead>
                <tbody>{daftar.map(m => (
                  <tr key={m.id}>
                    <td className="monospace">{m.nim}</td>
                    <td><strong>{m.nama}</strong><small>{m.email}</small></td>
                    <td>{m.angkatan}</td><td>{m.ipk}</td>
                    <td><span className={`status ${m.aktif === false ? 'status-inactive' : ''}`}>{m.aktif === false ? 'Nonaktif' : 'Aktif'}</span></td>
                    <td><button type="button" className="button button-danger compact" disabled={menghapus === m.id || !bolehHapus} onClick={() => tanganiHapus(m)} title={!bolehHapus ? 'Hanya admin yang boleh menghapus' : 'Hapus mahasiswa'}>{menghapus === m.id ? 'Menghapus...' : 'Hapus'}</button></td>
                  </tr>
                ))}</tbody>
              </table>
            </div>
          )
        )}
        <div className="pagination">
          <span>Halaman {meta.halaman} dari {meta.total_halaman}</span>
          <div className="pagination-controls">
            <button type="button" className="button button-outline compact" disabled={halaman <= 1 || sedangMemuat} onClick={() => setHalaman(h => Math.max(1, h - 1))}>Sebelumnya</button>
            <button type="button" className="button button-outline compact" disabled={halaman >= meta.total_halaman || sedangMemuat} onClick={() => setHalaman(h => h + 1)}>Berikutnya</button>
          </div>
        </div>
      </div>
      {peran && peran !== 'admin' && <p className="muted role-note">Peran saat ini: {peran}. Endpoint hapus hanya untuk admin. Tambah tersedia untuk admin, atau dosen di backend Fiber.</p>}
    </section>
  )
}
