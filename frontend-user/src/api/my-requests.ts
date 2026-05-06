import { apiClient } from '@/api/client'
import type {
  CreatePermintaanRequest,
  PermintaanDarah,
  PermintaanDarahListItem,
  UpdatePermintaanRequest,
  UpdatePermintaanStatusRequest,
} from '@/types/models'

export interface FetchMyRequestsParams {
  limit?: number
  offset?: number
}

export const myRequestsAPI = {
  getAll(params: FetchMyRequestsParams = {}) {
    return apiClient.get<PermintaanDarahListItem[]>('/permintaan-darah/my-requests', {
      params: {
        limit: params.limit ?? 100,
        offset: params.offset ?? 0,
      },
    })
  },

  getById(id: string) {
    return apiClient.get<PermintaanDarah>(`/permintaan-darah/${id}`)
  },

  create(payload: CreatePermintaanRequest) {
    return apiClient.post<PermintaanDarah>('/permintaan-darah', payload)
  },

  updateMyRequest(id: string, payload: UpdatePermintaanRequest) {
    return apiClient.put<PermintaanDarah>(`/permintaan-darah/my-requests/${id}`, payload)
  },

  updateStatus(id: string, payload: UpdatePermintaanStatusRequest) {
    return apiClient.put<PermintaanDarah>(`/permintaan-darah/update/${id}`, payload)
  },
}
