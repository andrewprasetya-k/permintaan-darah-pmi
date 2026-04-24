import { apiClient } from './client'
import type { Admin } from '@/types/models'

export interface CreateAdminRequest {
  adminUsername: string
  adminPassword: string
  adminName: string
  adminEmail: string
  adminRole: 'superadmin' | 'admin'
}

export interface UpdateAdminRequest {
  adminUsername: string
  adminPassword?: string
  adminName: string
  adminEmail: string
  adminRole: 'superadmin' | 'admin'
}

export const adminAPI = {
  create(data: CreateAdminRequest) {
    return apiClient.post<Admin>('/admin', data)
  },

  getAll(params?: Record<string, any>) {
    return apiClient.get<Admin[]>('/admin', { params })
  },

  getById(id: string) {
    return apiClient.get<Admin>(`/admin/${id}`)
  },

  update(id: string, data: UpdateAdminRequest) {
    return apiClient.put<Admin>(`/admin/${id}`, data)
  },

  delete(id: string) {
    return apiClient.delete<void>(`/admin/${id}`)
  },

  restore(id: string) {
    return apiClient.put<Admin>(`/admin/restore/${id}`, {})
  },
}
