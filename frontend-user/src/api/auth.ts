import { apiClient } from '@/api/client'
import type { LoginResponse } from '@/types/models'

export interface LoginRequest {
  username: string
  password: string
}

export const authAPI = {
  loginRumahSakit(payload: LoginRequest) {
    return apiClient.post<LoginResponse>('/auth/login/rumah-sakit', payload, {
      includeAuth: false,
    })
  },
}
