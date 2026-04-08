import { apiClient } from './client'
import type { KomponenDarah } from '@/types/models'

export interface CreateKomponenRequest {
  komponenDarah: string
  komponenKode: string
  isActive?: boolean
}

export interface UpdateKomponenRequest {
  komponenDarah: string
  komponenKode: string
  isActive?: boolean
}

export const komponenAPI = {
  create(data: CreateKomponenRequest) {
    return apiClient.post<KomponenDarah>('/komponen-darah', data)
  },

  getAll(params?: Record<string, any>) {
    return apiClient.get<KomponenDarah[]>('/komponen-darah', { params })
  },

  getById(id: number) {
    return apiClient.get<KomponenDarah>(`/komponen-darah/${id}`)
  },

  update(id: number, data: UpdateKomponenRequest) {
    return apiClient.put<KomponenDarah>(`/komponen-darah/${id}`, data)
  },

  delete(id: number) {
    return apiClient.delete<void>(`/komponen-darah/${id}`)
  },

  activate(id: number) {
    return apiClient.put<KomponenDarah>(`/komponen-darah/activate/${id}`, {})
  },

  deactivate(id: number) {
    return apiClient.put<KomponenDarah>(`/komponen-darah/deactivate/${id}`, {})
  },
}
