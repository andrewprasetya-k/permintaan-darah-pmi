import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useNotification } from '@/composables/useNotification'
import { usePermintaanStore } from './permintaan'
import { useLogsStore } from './logs'
import { useDashboardStore } from './dashboard'
import type { WebSocketMessage, PermintaanDarah, SystemAccessLog } from '@/types/models'

const toWebSocketUrl = (apiBaseUrl: string) => {
  const url = new URL(apiBaseUrl)
  url.protocol = url.protocol === 'https:' ? 'wss:' : 'ws:'
  url.pathname = `${url.pathname.replace(/\/api\/?$/, '')}/api/ws/connect`
  return url
}

export const useWebsocketStore = defineStore('websocket', () => {
  const isConnected = ref(false)
  const error = ref<string | null>(null)
  let socket: WebSocket | null = null
  let reconnectTimer: ReturnType<typeof setTimeout> | null = null
  let shouldReconnect = false

  const { notifyNewPermintaan, notifyPermintaanUpdate } = useNotification()
  const permintaanStore = usePermintaanStore()
  const logsStore = useLogsStore()
  const dashboardStore = useDashboardStore()

  const clearReconnectTimer = () => {
    if (reconnectTimer) {
      clearTimeout(reconnectTimer)
      reconnectTimer = null
    }
  }

  const scheduleReconnect = () => {
    if (!shouldReconnect || reconnectTimer) return

    reconnectTimer = setTimeout(() => {
      reconnectTimer = null
      connect()
    }, 4000)
  }

  const connect = () => {
    const token = localStorage.getItem('authToken')
    if (!token || socket) return

    shouldReconnect = true
    clearReconnectTimer()

    const apiBaseUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'
    const wsUrl = toWebSocketUrl(apiBaseUrl)
    wsUrl.searchParams.set('token', token)

    socket = new WebSocket(wsUrl.toString())

    socket.onopen = () => {
      isConnected.value = true
      logsStore.setRealtimeConnected(true)
      error.value = null
    }

    socket.onmessage = (event) => {
      try {
        const message = JSON.parse(event.data) as WebSocketMessage

        // Handle Permintaan Darah events
        if (message.entityType === 'permintaan_darah') {
          const pd = message.data as PermintaanDarah

          if (message.type === 'new_permintaan' && message.action === 'CREATE') {
            permintaanStore.markRealtimeHighlight(pd.permintaanDarahId, 'new')
            notifyNewPermintaan({
              namaPasien: pd.namaPasien,
              golonganDarah: pd.golonganDarah || 'N/A',
              rhesus: pd.rhesusDarah || '',
              rumahSakit: pd.rumahSakitId,
            })
            // Refresh counts on dashboard if we are there
            void dashboardStore.fetchStatusSummary('all')
          } else if (message.type === 'status_change' && message.action === 'UPDATE') {
            permintaanStore.markRealtimeHighlight(pd.permintaanDarahId, 'status')
            notifyPermintaanUpdate({
              namaPasien: pd.namaPasien,
              statusBaru: pd.status,
            })
            void dashboardStore.fetchStatusSummary('all')
          } else if (message.type === 'update_permintaan' && message.action === 'UPDATE') {
            permintaanStore.markRealtimeHighlight(pd.permintaanDarahId, 'updated')
            notifyPermintaanUpdate({
              namaPasien: pd.namaPasien,
            })
            void dashboardStore.fetchStatusSummary('all')
          }

          // Refresh requests list if active
          void permintaanStore.fetchAll()
        }

        // Handle System Logs / Activity
        if (message.type === 'new_activity' && message.data) {
          logsStore.pushRecentActivity(message.data as SystemAccessLog)
        }
      } catch (e) {
        console.error('Error processing WebSocket message:', e)
      }
    }

    socket.onerror = () => {
      error.value = 'Realtime connection error'
      isConnected.value = false
      logsStore.setRealtimeConnected(false)
    }

    socket.onclose = () => {
      isConnected.value = false
      logsStore.setRealtimeConnected(false)
      socket = null
      scheduleReconnect()
    }
  }

  const disconnect = () => {
    shouldReconnect = false
    clearReconnectTimer()
    if (socket) {
      socket.close()
      socket = null
    }
    isConnected.value = false
    logsStore.setRealtimeConnected(false)
  }

  return {
    isConnected,
    error,
    connect,
    disconnect,
  }
})
