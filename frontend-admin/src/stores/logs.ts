import { computed, ref } from 'vue'
import { defineStore } from 'pinia'
import { logsAPI } from '@/api/logs'
import type { PaginationMeta, SystemAccessLog, WebSocketMessage } from '@/types/models'

const MAX_RECENT_ACTIVITIES = 10

const toWebSocketUrl = (apiBaseUrl: string) => {
  const url = new URL(apiBaseUrl)
  url.protocol = url.protocol === 'https:' ? 'wss:' : 'ws:'
  url.pathname = `${url.pathname.replace(/\/api\/?$/, '')}/api/ws/connect`
  return url
}

export const useLogsStore = defineStore('logs', () => {
  const systemLogs = ref<SystemAccessLog[]>([])
  const recentActivities = ref<SystemAccessLog[]>([])
  const isLoading = ref(false)
  const isRealtimeConnected = ref(false)
  const error = ref<string | null>(null)
  const pagination = ref<PaginationMeta | null>(null)

  let socket: WebSocket | null = null

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

  const pushRecentActivity = (log: SystemAccessLog) => {
    recentActivities.value = [log, ...recentActivities.value.filter((item) => item.sysLogId !== log.sysLogId)].slice(
      0,
      MAX_RECENT_ACTIVITIES,
    )
    systemLogs.value = [log, ...systemLogs.value.filter((item) => item.sysLogId !== log.sysLogId)]
  }

  const connectRealtime = () => {
    const token = localStorage.getItem('authToken')
    if (!token || socket) return

    const apiBaseUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'
    const wsUrl = toWebSocketUrl(apiBaseUrl)
    wsUrl.searchParams.set('token', token)

    socket = new WebSocket(wsUrl.toString())

    socket.onopen = () => {
      isRealtimeConnected.value = true
    }

    socket.onmessage = (event) => {
      try {
        const message = JSON.parse(event.data) as WebSocketMessage<SystemAccessLog>
        if (message.type === 'new_activity' && message.data) {
          pushRecentActivity(message.data)
        }
      } catch {
        // Ignore malformed messages from unrelated events
      }
    }

    socket.onerror = () => {
      error.value = 'Realtime connection error'
    }

    socket.onclose = () => {
      isRealtimeConnected.value = false
      socket = null
    }
  }

  const disconnectRealtime = () => {
    if (socket) {
      socket.close()
      socket = null
    }
    isRealtimeConnected.value = false
  }

  return {
    systemLogs,
    recentActivities,
    recentActivityItems,
    isLoading,
    isRealtimeConnected,
    error,
    pagination,
    fetchSystemLogs,
    connectRealtime,
    disconnectRealtime,
  }
})
