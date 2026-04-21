import { apiClient } from './client'

export interface StatusSummary {
  dibuat: number
  diproses: number
  bisaDiambil: number
  selesai: number
  dibatalkan: number
  total: number
}

export const dashboardAPI = {
  getStatusSummary(rumahSakitID: string) {
    return apiClient.get<StatusSummary>(
      `/dashboard/status-summary/${rumahSakitID}`,
    )
  },
}
