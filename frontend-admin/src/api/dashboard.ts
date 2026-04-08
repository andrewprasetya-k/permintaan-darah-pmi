import { apiClient } from './client'
import type { ApiResponse } from '@/types/models'

export interface StatusSummary {
  pending: number
  approved: number
  rejected: number
  completed: number
}

export const dashboardAPI = {
  getStatusSummary(rumahSakitID: string) {
    return apiClient.get<StatusSummary>(
      `/dashboard/status-summary/${rumahSakitID}`,
    )
  },
}
