import { apiClient } from './client'
import type { RumahSakit } from '@/types/models'

export interface CreateRumahSakitRequest {
  nama: string
  nomorTelepon: string
  alamat: string
  email?: string
  username: string
  password: string
}

export interface UpdateRumahSakitRequest {
  nama: string
  nomorTelepon: string
  alamat: string
  email?: string
  password?: string
}

export const rumahSakitAPI = {
  create(data: CreateRumahSakitRequest) {
    return apiClient.post<RumahSakit>('/rumah-sakit', data)
  },

  getAll(params?: Record<string, any>) {
    return apiClient.get<RumahSakit[]>('/rumah-sakit', { params })
  },

  getById(id: string) {
    return apiClient.get<RumahSakit>(`/rumah-sakit/${id}`)
  },

  update(id: string, data: UpdateRumahSakitRequest) {
    return apiClient.put<RumahSakit>(`/rumah-sakit/${id}`, data)
  },

  delete(id: string) {
    return apiClient.delete<void>(`/rumah-sakit/${id}`)
  },

  restore(id: string) {
    return apiClient.put<RumahSakit>(`/rumah-sakit/restore/${id}`, {})
  },
}
