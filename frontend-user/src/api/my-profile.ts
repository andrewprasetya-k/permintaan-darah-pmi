import { apiClient } from '@/api/client'
import type { RumahSakit, UpdateRumahSakitProfileRequest } from '@/types/models'

export const myProfileAPI = {
  getMe() {
    return apiClient.get<RumahSakit>('/rumah-sakit/me')
  },

  updateMe(payload: UpdateRumahSakitProfileRequest) {
    return apiClient.put<RumahSakit>('/rumah-sakit/me', payload)
  },
}
