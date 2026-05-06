import { defineStore } from 'pinia'
import { ref } from 'vue'
import {
  permintaanAPI,
  type CreatePermintaanRequest,
  type UpdatePermintaanRequest,
} from '@/api/permintaan'
import type { PaginationMeta, PermintaanDarah, WebSocketMessage } from '@/types/models'

const toWebSocketUrl = (apiBaseUrl: string) => {
  const url = new URL(apiBaseUrl)
  url.protocol = url.protocol === 'https:' ? 'wss:' : 'ws:'
  url.pathname = `${url.pathname.replace(/\/api\/?$/, '')}/api/ws/connect`
  return url
}

interface FetchPermintaanParams {
  search?: string
  status?: string
  limit?: number
  offset?: number
}

export const usePermintaanStore = defineStore('permintaan', () => {
  const requests = ref<PermintaanDarah[]>([])
  const selectedRequest = ref<PermintaanDarah | null>(null)
  const pagination = ref<PaginationMeta | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)
  const isRealtimeConnected = ref(false)

  let socket: WebSocket | null = null

  const fetchAll = async (params?: FetchPermintaanParams) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await permintaanAPI.getAll(params)
      requests.value = response.data
      pagination.value = response.pagination ?? null
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch requests'
    } finally {
      isLoading.value = false
    }
  }

  const fetchById = async (id: string) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await permintaanAPI.getById(id)
      selectedRequest.value = response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch request'
    } finally {
      isLoading.value = false
    }
  }

  const create = async (data: CreatePermintaanRequest) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await permintaanAPI.create(data)
      requests.value.push(response.data)
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to create request'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const update = async (id: string, data: UpdatePermintaanRequest) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await permintaanAPI.update(id, data)
      const index = requests.value.findIndex((r) => r.permintaanDarahId === id)
      if (index !== -1) {
        requests.value[index] = response.data
      }
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to update request'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const updateStatus = async (id: string, status: PermintaanDarah['status']) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await permintaanAPI.updateStatus(id, { status })
      const index = requests.value.findIndex((r) => r.permintaanDarahId === id)
      if (index !== -1) {
        requests.value[index] = response.data
      }
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to update status'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const deleteRequest = async (id: string) => {
    isLoading.value = true
    error.value = null
    try {
      await permintaanAPI.delete(id)
      requests.value = requests.value.filter((r) => r.permintaanDarahId !== id)
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to delete request'
      throw err
    } finally {
      isLoading.value = false
    }
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
        const message = JSON.parse(event.data) as WebSocketMessage<PermintaanDarah>
        if (message.entityType !== 'permintaan_darah') {
          return
        }

        // Refresh all permintaan data on any permintaan event
        void fetchAll()
      } catch {
        // Ignore malformed messages
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
    requests,
    selectedRequest,
    pagination,
    isLoading,
    error,
    isRealtimeConnected,
    fetchAll,
    fetchById,
    create,
    update,
    updateStatus,
    deleteRequest,
    connectRealtime,
    disconnectRealtime,
  }
})
