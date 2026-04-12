import { apiClient } from './client'
import type { DetailPermintaanDarah } from '@/types/models'

export interface CreateDetailRequest {
  pdId: string
  komId: number
  golonganDarah: 'A' | 'B' | 'AB' | 'O'
  rhesusDarah: '+' | '-'
  jumlahKantong: number
}

export interface UpdateDetailRequest {
  jumlahKantong?: number
  tanggalDiperlukan?: string
}

export const detailPermintaanAPI = {
  create(data: CreateDetailRequest) {
    return apiClient.post<DetailPermintaanDarah>('/detail-permintaan-darah', data)
  },

  getAll(params?: Record<string, any>) {
    return apiClient.get<DetailPermintaanDarah[]>('/detail-permintaan-darah', { params })
  },

  getById(id: number) {
    return apiClient.get<DetailPermintaanDarah>(`/detail-permintaan-darah/${id}`)
  },

  update(id: number, data: UpdateDetailRequest) {
    return apiClient.put<DetailPermintaanDarah>(`/detail-permintaan-darah/${id}`, data)
  },

  delete(id: number) {
    return apiClient.delete<void>(`/detail-permintaan-darah/${id}`)
  },
}
