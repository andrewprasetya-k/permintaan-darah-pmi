import { apiClient } from './client'
import type { ApiResponse } from '@/types/models'

export interface StatusLog {
  id: string
  permintaanID: string
  status: string
  keterangan: string
  createdAt: string
}

export interface SystemAccessLog {
  id: string
  userID: string
  action: string
  table: string
  targetID: string
  oldValue: string
  newValue: string
  timestamp: string
  ipAddress: string
}

export const logsAPI = {
  getStatusLogs(params?: Record<string, any>) {
    return apiClient.get<StatusLog[]>('/status-logs', { params })
  },

  getSystemLogs(params?: Record<string, any>) {
    return apiClient.get<SystemAccessLog[]>('/system-logs', { params })
  },

  getSystemLogById(id: string) {
    return apiClient.get<SystemAccessLog>(`/system-logs/${id}`)
  },

  getSystemLogsByUserId(userId: string, params?: Record<string, any>) {
    return apiClient.get<SystemAccessLog[]>(`/system-logs/user/${userId}`, { params })
  },

  getSystemLogsByAction(action: string, params?: Record<string, any>) {
    return apiClient.get<SystemAccessLog[]>(`/system-logs/action/${action}`, { params })
  },

  getSystemLogsByTable(table: string, params?: Record<string, any>) {
    return apiClient.get<SystemAccessLog[]>(`/system-logs/table/${table}`, { params })
  },

  getSystemLogsByTargetId(targetId: string, params?: Record<string, any>) {
    return apiClient.get<SystemAccessLog[]>(`/system-logs/target/${targetId}`, { params })
  },
}
