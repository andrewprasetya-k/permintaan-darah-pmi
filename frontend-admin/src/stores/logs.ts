import { computed, ref } from 'vue'
import { defineStore } from 'pinia'
import { logsAPI } from '@/api/logs'
import type { PaginationMeta, SystemAccessLog } from '@/types/models'
import type { StatusLog } from '@/types/models'

const MAX_RECENT_ACTIVITIES = 10

export const useLogsStore = defineStore('logs', () => {
  const systemLogs = ref<SystemAccessLog[]>([])
  const statusLogs = ref<StatusLog[]>([])
  const recentActivities = ref<SystemAccessLog[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)
  const pagination = ref<PaginationMeta | null>(null)

  const recentActivityItems = computed(() => recentActivities.value.slice(0, MAX_RECENT_ACTIVITIES))

  const fetchSystemLogs = async (params?: Record<string, any>) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await logsAPI.getSystemLogs(params)
      systemLogs.value = response.data
      pagination.value = response.pagination ?? null
      if (!params || Number(params.offset ?? 0) === 0) {
        recentActivities.value = response.data.slice(0, MAX_RECENT_ACTIVITIES)
      }
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch system logs'
    } finally {
      isLoading.value = false
    }
  }

  const fetchStatusLogs = async (params?: Record<string, any>) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await logsAPI.getStatusLogs(params)
      statusLogs.value = response.data
      pagination.value = response.pagination ?? null
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch status logs'
    } finally {
      isLoading.value = false
    }
  }

  const fetchSystemLogsByUserId = async (userId: string, params?: Record<string, any>) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await logsAPI.getSystemLogsByUserId(userId, params)
      systemLogs.value = response.data
      pagination.value = response.pagination ?? null
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch system logs'
    } finally {
      isLoading.value = false
    }
  }

  const fetchSystemLogsByAction = async (action: string, params?: Record<string, any>) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await logsAPI.getSystemLogsByAction(action, params)
      systemLogs.value = response.data
      pagination.value = response.pagination ?? null
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch system logs'
    } finally {
      isLoading.value = false
    }
  }

  const fetchSystemLogsByTable = async (table: string, params?: Record<string, any>) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await logsAPI.getSystemLogsByTable(table, params)
      systemLogs.value = response.data
      pagination.value = response.pagination ?? null
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch system logs'
    } finally {
      isLoading.value = false
    }
  }

  const fetchSystemLogsByTargetId = async (targetId: string, params?: Record<string, any>) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await logsAPI.getSystemLogsByTargetId(targetId, params)
      systemLogs.value = response.data
      pagination.value = response.pagination ?? null
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch system logs'
    } finally {
      isLoading.value = false
    }
  }

  const pushRecentActivity = (log: SystemAccessLog) => {
    recentActivities.value = [log, ...recentActivities.value.filter((item) => item.sysLogId !== log.sysLogId)].slice(
      0,
      MAX_RECENT_ACTIVITIES,
    )
    systemLogs.value = [log, ...systemLogs.value.filter((item) => item.sysLogId !== log.sysLogId)]
  }

  return {
    systemLogs,
    statusLogs,
    recentActivities,
    recentActivityItems,
    isLoading,
    error,
    pagination,
    fetchSystemLogs,
    fetchStatusLogs,
    fetchSystemLogsByUserId,
    fetchSystemLogsByAction,
    fetchSystemLogsByTable,
    fetchSystemLogsByTargetId,
    pushRecentActivity,
  }
})
