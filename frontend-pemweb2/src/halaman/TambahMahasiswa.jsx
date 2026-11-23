import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { normalisasiGalat } from '../api/kontrak'
import { simpanMahasiswa } from '../api/layananMahasiswa'

const nilaiAwal = {
  program_studi_id: '1', nim: '', nama: '', email: '',
  angkatan: '2024', ipk: '3.50', aktif: true,
}

export default function TambahMahasiswa() {
  const [form, setForm] = useState(nilaiAwal)
  const [galat, setGalat] = useState('')
  const [daftarGalat, setDaftarGalat] = useState([])
  const [sedangMemuat, setSedangMemuat] = useState(false)
  const navigasi = useNavigate()

  function ubah(e) {
    const { name, type, checked, value } = e.target
    setForm(sebelumnya => ({ ...sebelumnya, [name]: type === 'checkbox' ? checked : value }))
  }

  async function tanganiSimpan(e) {
    e.preventDefault()
    setSedangMemuat(true)
    setGalat('')
    setDaftarGalat([])
    try {
      await simpanMahasiswa({
        program_studi_id: Number(form.program_studi_id),
        nim: form.nim.trim(),
        nama: form.nama.trim(),
        email: form.email.trim(),
        angkatan: Number(form.angkatan),
        ipk: Number(form.ipk),
        aktif: form.aktif,
      })
      navigasi('/mahasiswa', { replace: true, state: { pesan: 'Mahasiswa baru berhasil disimpan.' } })
    } catch (error) {
      const hasil = normalisasiGalat(error)
      setGalat(hasil.pesan)
      setDaftarGalat(hasil.daftar)
    } finally {
      setSedangMemuat(false)
    }
  }

  return (
    <section>
      <div className="heading-row">
        <div><p className="eyebrow">DATA MAHASISWA</p><h1>Tambah mahasiswa</h1><p className="muted">Pesan validasi dari backend (422) ditampilkan di bawah formulir.</p></div>
        <Link to="/mahasiswa" className="button button-outline">Kembali</Link>
      </div>
      <div className="paper form-paper">
        <div className="callout">format NIM yang diterima adalah huruf <strong>H</strong> diikuti angka, contoh <code>H12345678</code>. Pastikan ID program studi tersedia di database.</div>
        {galat && <div className="alert alert-error" role="alert"><strong>{galat}</strong>{daftarGalat.length > 0 && <ul className="error-list">{daftarGalat.map((g,i) => <li key={`${g.field}-${i}`}><strong>{g.field}:</strong> {g.pesan}</li>)}</ul>}</div>}
        <form noValidate onSubmit={tanganiSimpan} className="form-grid">
          <div className="field"><label htmlFor="program_studi_id">ID Program Studi</label><input id="program_studi_id" name="program_studi_id" type="number" min="1" value={form.program_studi_id} onChange={ubah}/></div>
          <div className="field"><label htmlFor="nim">NIM</label><input id="nim" name="nim" value={form.nim} onChange={ubah} placeholder="H12345678" /></div>
          <div className="field"><label htmlFor="nama">Nama lengkap</label><input id="nama" name="nama" value={form.nama} onChange={ubah} placeholder="Nama mahasiswa" /></div>
          <div className="field"><label htmlFor="email">Email</label><input id="email" name="email" type="email" value={form.email} onChange={ubah} placeholder="nama@example.com" /></div>
          <div className="field"><label htmlFor="angkatan">Angkatan</label><input id="angkatan" name="angkatan" type="number" value={form.angkatan} onChange={ubah} /></div>
          <div className="field"><label htmlFor="ipk">IPK</label><input id="ipk" name="ipk" type="number" step="0.01" min="0" max="4" value={form.ipk} onChange={ubah} /></div>
          <label className="checkbox-field"><input id="aktif" name="aktif" type="checkbox" checked={form.aktif} onChange={ubah} /><span>Mahasiswa aktif</span></label>
          <div className="form-actions"><Link to="/mahasiswa" className="button button-outline">Batal</Link><button type="submit" className="button button-primary" disabled={sedangMemuat}>{sedangMemuat ? 'Menyimpan...' : 'Simpan mahasiswa'}</button></div>
        </form>
      </div>
    </section>
  )
}
