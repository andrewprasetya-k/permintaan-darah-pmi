import { apiClient } from './client'
import type { PermintaanDarah } from '@/types/models'

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
    return apiClient.post<PermintaanDarah>('/permintaan-darah', data)
  },

  getAll(params?: Record<string, any>) {
    return apiClient.get<PermintaanDarah[]>('/permintaan-darah', { params })
  },

  getMyRequests(params?: Record<string, any>) {
    return apiClient.get<PermintaanDarah[]>('/permintaan-darah/my-requests', { params })
  },

  getById(id: string) {
    return apiClient.get<PermintaanDarah>(`/permintaan-darah/${id}`)
  },

  update(id: string, data: UpdatePermintaanRequest) {
    return apiClient.put<PermintaanDarah>(`/permintaan-darah/${id}`, data)
  },

  updateMyRequest(id: string, data: UpdatePermintaanRequest) {
    return apiClient.put<PermintaanDarah>(`/permintaan-darah/my-requests/${id}`, data)
  },

  updateStatus(id: string, data: UpdateStatusRequest) {
    return apiClient.put<PermintaanDarah>(`/permintaan-darah/update/${id}`, data)
  },

  delete(id: string) {
    return apiClient.delete<void>(`/permintaan-darah/${id}`)
  },
}
