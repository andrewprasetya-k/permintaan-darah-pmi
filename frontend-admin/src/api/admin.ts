import { apiClient } from './client'
import type { ApiResponse, User } from '@/types/models'

export interface CreateAdminRequest {
  name: string
  email: string
  password: string
  role: 'admin' | 'user'
}

export interface UpdateAdminRequest extends Partial<CreateAdminRequest> {}

export const adminAPI = {
  create(data: CreateAdminRequest) {
    return apiClient.post<User>('/admin', data)
  },

  getAll(params?: Record<string, any>) {
    return apiClient.get<User[]>('/admin', { params })
  },

  getById(id: string) {
    return apiClient.get<User>(`/admin/${id}`)
  },

  update(id: string, data: UpdateAdminRequest) {
    return apiClient.put<User>(`/admin/${id}`, data)
  },

  delete(id: string) {
    return apiClient.delete<void>(`/admin/${id}`)
  },

  restore(id: string) {
    return apiClient.put<User>(`/admin/restore/${id}`, {})
  },
}
