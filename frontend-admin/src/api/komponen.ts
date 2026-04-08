import { apiClient } from './client'
import type { ApiResponse } from '@/types/models'

export interface Komponen {
  id: string
  nama: string
  deskripsi: string
  status: 'active' | 'inactive'
  createdAt: string
  updatedAt: string
}

export interface CreateKomponenRequest {
  nama: string
  deskripsi: string
}

export interface UpdateKomponenRequest extends Partial<CreateKomponenRequest> {}

export const komponenAPI = {
  create(data: CreateKomponenRequest) {
    return apiClient.post<Komponen>('/komponen-darah', data)
  },

  getAll(params?: Record<string, any>) {
    return apiClient.get<Komponen[]>('/komponen-darah', { params })
  },

  getById(id: string) {
    return apiClient.get<Komponen>(`/komponen-darah/${id}`)
  },

  update(id: string, data: UpdateKomponenRequest) {
    return apiClient.put<Komponen>(`/komponen-darah/${id}`, data)
  },

  delete(id: string) {
    return apiClient.delete<void>(`/komponen-darah/${id}`)
  },

  activate(id: string) {
    return apiClient.put<Komponen>(`/komponen-darah/activate/${id}`, {})
  },

  deactivate(id: string) {
    return apiClient.put<Komponen>(`/komponen-darah/deactivate/${id}`, {})
  },
}
