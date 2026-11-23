import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { backend } from '../api/konfigurasi'
import { logout } from '../api/layananAuth'
import { normalisasiGalat, peranPengguna } from '../api/kontrak'

export default function TataLetak({ children }) {
  const navigasi = useNavigate()
  const [memuat, setMemuat] = useState(false)
  const [galat, setGalat] = useState('')
  let pengguna = null
  try { pengguna = JSON.parse(sessionStorage.getItem(`pengguna_${backend}`) || 'null') } catch { /* invalid cached data */ }
  const peran = peranPengguna(pengguna)

  async function tanganiLogout() {
    setMemuat(true)
    setGalat('')
    try {
      await logout()
      sessionStorage.removeItem(`token_${backend}`)
      sessionStorage.removeItem(`pengguna_${backend}`)
      navigasi('/login', { replace: true })
    } catch (error) {
      setGalat(`${normalisasiGalat(error).pesan}. Logout belum dikonfirmasi server.`)
    } finally {
      setMemuat(false)
    }
  }

  return (
    <div className="app-shell">
      <header className="topbar">
        <div className="topbar-inner">
          <Link to="/mahasiswa" className="brand">
            <span className="brand-icon">P</span>
            <span>PEMWEB <strong>II</strong></span>
          </Link>
          <div className="topbar-actions">
            <span className="backend-tag">API {backend.toUpperCase()}</span>
            {pengguna && <span className="user-tag">{pengguna.nama || pengguna.name || pengguna.email} {peran ? `· ${peran}` : ''}</span>}
            <button type="button" className="button button-ghost compact" onClick={tanganiLogout} disabled={memuat}>
              {memuat ? 'Memproses...' : 'Keluar'}
            </button>
          </div>
        </div>
      </header>
      <main className="page-container">
        {galat && <div className="alert alert-error" role="alert">{galat}</div>}
        {children}
      </main>
      <footer className="site-footer">Praktikum Pemrograman Web II</footer>
    </div>
  )
}
