import { apiClient } from '@/api/client'
import type { KomponenDarah } from '@/types/models'

export const komponenAPI = {
  getAll() {
    return apiClient.get<KomponenDarah[]>('/komponen-darah', {
      params: {
        limit: 100,
        offset: 0,
      },
    })
  },
}
