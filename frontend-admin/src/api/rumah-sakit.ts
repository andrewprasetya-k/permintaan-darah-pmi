import { apiClient } from './client'
import type { ApiResponse } from '@/types/models'

export interface RumahSakit {
  id: string
  nama: string
  alamat: string
  kota: string
  provinsi: string
  telepon: string
  email: string
  createdAt: string
  updatedAt: string
}

export interface CreateRumahSakitRequest {
  nama: string
  alamat: string
  kota: string
  provinsi: string
  telepon: string
  email: string
}

export interface UpdateRumahSakitRequest extends Partial<CreateRumahSakitRequest> {}

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

  getMe() {
    return apiClient.get<RumahSakit>('/rumah-sakit/me')
  },

  updateProfile(data: UpdateRumahSakitRequest) {
    return apiClient.put<RumahSakit>('/rumah-sakit/me', data)
  },

  getDistinctNama() {
    return apiClient.get<string[]>('/filter/rumah-sakit/')
  },
}
