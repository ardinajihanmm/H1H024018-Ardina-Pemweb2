import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom'
import Login from './halaman/Login'
import DaftarMahasiswa from './halaman/DaftarMahasiswa'
import TambahMahasiswa from './halaman/TambahMahasiswa'
import RuteTerlindungi from './komponen/RuteTerlindungi'
import TataLetak from './komponen/TataLetak'

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Navigate to="/mahasiswa" replace />} />
        <Route path="/login" element={<Login />} />
        <Route
          path="/mahasiswa"
          element={<RuteTerlindungi><TataLetak><DaftarMahasiswa /></TataLetak></RuteTerlindungi>}
        />
        <Route
          path="/mahasiswa/tambah"
          element={<RuteTerlindungi><TataLetak><TambahMahasiswa /></TataLetak></RuteTerlindungi>}
        />
        <Route path="*" element={<Navigate to="/mahasiswa" replace />} />
      </Routes>
    </BrowserRouter>
  )
}
