import { apiClient } from './client'
import type { ApiResponse, BloodRequest } from '@/types/models'

export interface CreatePermintaanRequest {
  patientName: string
  bloodType: string
  quantity: number
  hospital: string
  urgency: 'urgent' | 'normal'
}

export interface UpdatePermintaanRequest extends Partial<CreatePermintaanRequest> {
  status?: BloodRequest['status']
}

export const permintaanAPI = {
  create(data: CreatePermintaanRequest) {
    return apiClient.post<BloodRequest>('/permintaan-darah', data)
  },

  getAll(params?: Record<string, any>) {
    return apiClient.get<BloodRequest[]>('/permintaan-darah', { params })
  },

  getMyRequests(params?: Record<string, any>) {
    return apiClient.get<BloodRequest[]>('/permintaan-darah/my-requests', { params })
  },

  getById(id: string) {
    return apiClient.get<BloodRequest>(`/permintaan-darah/${id}`)
  },

  update(id: string, data: UpdatePermintaanRequest) {
    return apiClient.put<BloodRequest>(`/permintaan-darah/${id}`, data)
  },

  updateMyRequest(id: string, data: UpdatePermintaanRequest) {
    return apiClient.put<BloodRequest>(`/permintaan-darah/my-requests/${id}`, data)
  },

  updateStatus(id: string, status: BloodRequest['status']) {
    return apiClient.put<BloodRequest>(`/permintaan-darah/update/${id}`, { status })
  },

  delete(id: string) {
    return apiClient.delete<void>(`/permintaan-darah/${id}`)
  },
}
