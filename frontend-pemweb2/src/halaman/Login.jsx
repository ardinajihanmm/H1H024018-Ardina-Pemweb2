import { useState } from 'react'
import { Navigate, useNavigate } from 'react-router-dom'
import { backend } from '../api/konfigurasi'
import { normalisasiGalat } from '../api/kontrak'
import { login } from '../api/layananAuth'

export default function Login() {
  const [email, setEmail] = useState('')
  const [kataSandi, setKataSandi] = useState('')
  const [pesanGalat, setPesanGalat] = useState('')
  const [sedangMemuat, setSedangMemuat] = useState(false)
  const navigasi = useNavigate()

  if (sessionStorage.getItem(`token_${backend}`)) {
    return <Navigate to="/mahasiswa" replace />
  }

  async function tanganiKirim(peristiwa) {
    peristiwa.preventDefault()
    setPesanGalat('')
    setSedangMemuat(true)
    try {
      const hasil = await login(email, kataSandi)
      sessionStorage.setItem(`token_${backend}`, hasil.token)
      sessionStorage.setItem(`pengguna_${backend}`, JSON.stringify(hasil.pengguna))
      navigasi('/mahasiswa', { replace: true })
    } catch (galat) {
      setPesanGalat(normalisasiGalat(galat).pesan)
    } finally {
      setSedangMemuat(false)
    }
  }

  return (
    <div className="login-screen">
      <div className="login-panel">
        <div className="login-mark">P</div>
        <p className="eyebrow">PRAKTIKUM PEMROGRAMAN WEB II</p>
        <h1>Masuk ke sistem</h1>
        <p className="muted">Masuk menggunakan akun yang telah dibuat</p>
        {pesanGalat && <div role="alert" className="alert alert-error">{pesanGalat}</div>}
        <form onSubmit={tanganiKirim} className="form-stack">
          <label htmlFor="email">Email</label>
          <input id="email" type="email" autoComplete="username" value={email} onChange={e => setEmail(e.target.value)} required placeholder="nama@example.com" />
          <label htmlFor="kata-sandi">Kata sandi</label>
          <input id="kata-sandi" type="password" autoComplete="current-password" value={kataSandi} onChange={e => setKataSandi(e.target.value)} required placeholder="Masukkan kata sandi" />
          <button type="submit" className="button button-primary full" disabled={sedangMemuat}>
            {sedangMemuat ? 'Memproses...' : 'Masuk'}
          </button>
        </form>
        <p className="login-note">Backend: <strong>{backend.toUpperCase()}</strong> · {backend === 'laravel' ? 'Sanctum API Token' : 'JWT'} (Bearer)</p>
      </div>
    </div>
  )
}
