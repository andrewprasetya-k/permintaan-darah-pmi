import { apiClient } from './client'
import type { BloodRequest } from '@/types/models'

export interface CreatePermintaanRequest {
  rumahSakitId?: string
  namaPasien: string
  noRM?: string
  tempatLahir: string
  tanggalLahir: string
  jenisKelamin: 'L' | 'P'
  golonganDarah?: 'A' | 'B' | 'AB' | 'O'
  rhesusDarah?: '+' | '-'
  hemoglobin?: number
  ruangBagianKelas?: string
  pernahTransfusi: boolean
  indikasiTransfusi?: string
  pernahHamil?: string
  status: 'dibuat' | 'diproses' | 'bisa_diambil' | 'selesai' | 'dibatalkan'
  cancelReason?: string
  tanggalPermintaan: string
}

export interface UpdatePermintaanRequest extends CreatePermintaanRequest {
  rumahSakitId: string
}

export interface UpdateStatusRequest {
  status: 'dibuat' | 'diproses' | 'bisa_diambil' | 'selesai' | 'dibatalkan'
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

  updateStatus(id: string, data: UpdateStatusRequest) {
    return apiClient.put<BloodRequest>(`/permintaan-darah/update/${id}`, data)
  },

  delete(id: string) {
    return apiClient.delete<void>(`/permintaan-darah/${id}`)
  },
}
