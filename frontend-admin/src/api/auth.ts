import { apiClient } from './client'
import type { ApiResponse, User } from '@/types/models'

export interface LoginRequest {
  username: string
  password: string
}

export interface LoginResponse {
  token: string
  user: User
}

export const authAPI = {
  loginAdmin(credentials: LoginRequest) {
    return apiClient.post<LoginResponse>('/auth/login/admin', credentials, { includeAuth: false })
  },

  logout() {
    localStorage.removeItem('authToken')
    localStorage.removeItem('authUser')
  },

  getMe() {
    return apiClient.get<User>('/admin/me')
  },

  updateProfile(data: Partial<User>) {
    return apiClient.put<User>('/admin/me', data)
  },
}
