import { defineStore } from 'pinia'
import { ref } from 'vue'
import {
  permintaanAPI,
  type CreatePermintaanRequest,
  type UpdatePermintaanRequest,
} from '@/api/permintaan'
import type { PaginationMeta, PermintaanDarah } from '@/types/models'

interface FetchPermintaanParams {
  search?: string
  status?: string
  limit?: number
  offset?: number
  startDate?: string
  endDate?: string
}

type RealtimeHighlightType = 'new' | 'status' | 'updated'

interface RealtimeHighlight {
  type: RealtimeHighlightType
  highlightedUntil: number
}

const HIGHLIGHT_DURATION_MS = 10000

export const usePermintaanStore = defineStore('permintaan', () => {
  const requests = ref<PermintaanDarah[]>([])
  const selectedRequest = ref<PermintaanDarah | null>(null)
  const pagination = ref<PaginationMeta | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)
  const lastParams = ref<FetchPermintaanParams | undefined>()
  const realtimeHighlights = ref<Record<string, RealtimeHighlight>>({})
  const highlightTimers = new Map<string, ReturnType<typeof setTimeout>>()

  const fetchAll = async (params?: FetchPermintaanParams) => {
    if (params) {
      lastParams.value = params
    }
    isLoading.value = true
    error.value = null
    try {
      const response = await permintaanAPI.getAll(params || lastParams.value)
      requests.value = response.data
      pagination.value = response.pagination ?? null
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Gagal memuat permintaan'
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
      error.value = err instanceof Error ? err.message : 'Gagal memuat permintaan'
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
      error.value = err instanceof Error ? err.message : 'Gagal membuat permintaan'
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
      error.value = err instanceof Error ? err.message : 'Gagal memperbarui permintaan'
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
      error.value = err instanceof Error ? err.message : 'Gagal memperbarui status'
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
      error.value = err instanceof Error ? err.message : 'Gagal menghapus permintaan'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const exportExcel = async (params?: FetchPermintaanParams) => {
    try {
      return await permintaanAPI.exportExcel(params || lastParams.value)
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Gagal mengekspor permintaan'
      throw err
    }
  }

  const markRealtimeHighlight = (id: string, type: RealtimeHighlightType) => {
    if (highlightTimers.has(id)) {
      clearTimeout(highlightTimers.get(id))
    }

    realtimeHighlights.value = {
      ...realtimeHighlights.value,
      [id]: {
        type,
        highlightedUntil: Date.now() + HIGHLIGHT_DURATION_MS,
      },
    }

    highlightTimers.set(
      id,
      setTimeout(() => {
        const next = { ...realtimeHighlights.value }
        delete next[id]
        realtimeHighlights.value = next
        highlightTimers.delete(id)
      }, HIGHLIGHT_DURATION_MS),
    )
  }

  const clearRealtimeHighlight = (id: string) => {
    if (highlightTimers.has(id)) {
      clearTimeout(highlightTimers.get(id))
      highlightTimers.delete(id)
    }
    const next = { ...realtimeHighlights.value }
    delete next[id]
    realtimeHighlights.value = next
  }

  const getRealtimeHighlight = (id: string) => realtimeHighlights.value[id]

  return {
    requests,
    selectedRequest,
    pagination,
    isLoading,
    error,
    realtimeHighlights,
    fetchAll,
    fetchById,
    create,
    update,
    updateStatus,
    deleteRequest,
    exportExcel,
    markRealtimeHighlight,
    clearRealtimeHighlight,
    getRealtimeHighlight,
  }
})
