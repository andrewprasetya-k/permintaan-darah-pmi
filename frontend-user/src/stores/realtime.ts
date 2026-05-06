import { ref, shallowRef } from 'vue'
import { defineStore } from 'pinia'
import { WS_BASE_URL } from '@/api/client'
import { useMyRequestsStore } from '@/stores/my-requests'
import type { WebSocketMessage } from '@/types/models'

type ConnectionState = 'idle' | 'connecting' | 'connected' | 'disconnected' | 'error'

export const useRealtimeStore = defineStore('realtime', () => {
  const socket = shallowRef<WebSocket | null>(null)
  const status = ref<ConnectionState>('idle')
  const lastMessage = ref<WebSocketMessage | null>(null)
  const error = ref<string | null>(null)
  let reconnectTimer: number | null = null
  let activeToken: string | null = null

  const clearReconnect = () => {
    if (reconnectTimer !== null) {
      window.clearTimeout(reconnectTimer)
      reconnectTimer = null
    }
  }

  const handleMessage = (message: WebSocketMessage) => {
    lastMessage.value = message

    if (message.entityType !== 'permintaan_darah') {
      return
    }

    const requestsStore = useMyRequestsStore()
    void requestsStore.fetchAll()

    if (requestsStore.selectedRequest?.permintaanDarahId === message.entityId) {
      void requestsStore.fetchById(message.entityId)
    }
  }

  const connect = (token: string | null) => {
    if (!token) {
      return
    }

    activeToken = token
    clearReconnect()

    if (socket.value) {
      socket.value.close()
      socket.value = null
    }

    status.value = 'connecting'
    error.value = null

    const ws = new WebSocket(`${WS_BASE_URL}/ws/connect?token=${encodeURIComponent(token)}`)
    socket.value = ws

    ws.onopen = () => {
      status.value = 'connected'
      error.value = null
    }

    ws.onmessage = (event: MessageEvent<string>) => {
      try {
        handleMessage(JSON.parse(event.data) as WebSocketMessage)
      } catch {
        error.value = 'Pesan realtime tidak valid'
      }
    }

    ws.onerror = () => {
      status.value = 'error'
      error.value = 'Koneksi realtime bermasalah'
    }

    ws.onclose = () => {
      socket.value = null
      status.value = activeToken ? 'disconnected' : 'idle'

      if (activeToken) {
        reconnectTimer = window.setTimeout(() => {
          connect(activeToken)
        }, 4000)
      }
    }
  }

  const disconnect = () => {
    activeToken = null
    clearReconnect()

    if (socket.value) {
      socket.value.close()
      socket.value = null
    }

    status.value = 'idle'
  }

  return {
    status,
    lastMessage,
    error,
    connect,
    disconnect,
  }
})
