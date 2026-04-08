import { apiClient } from './client'

export interface LoginRequest {
  username: string
  password: string
}

export interface LoginResponse {
  id: string
  username: string
  role: string
  token: string
}

export const authAPI = {
  loginAdmin(credentials: LoginRequest) {
    return apiClient.post<LoginResponse>('/auth/login/admin', credentials, { includeAuth: false })
  },

  logout() {
    localStorage.removeItem('authToken')
    localStorage.removeItem('authUser')
  },
}
