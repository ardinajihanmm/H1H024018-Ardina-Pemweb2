import { Navigate } from 'react-router-dom'
import { backend } from '../api/konfigurasi'

export default function RuteTerlindungi({ children }) {
  const token = sessionStorage.getItem(`token_${backend}`)
  return token ? children : <Navigate to="/login" replace />
}
