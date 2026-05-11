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

export const usePermintaanStore = defineStore('permintaan', () => {
  const requests = ref<PermintaanDarah[]>([])
  const selectedRequest = ref<PermintaanDarah | null>(null)
  const pagination = ref<PaginationMeta | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)
  const lastParams = ref<FetchPermintaanParams | undefined>()

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

  const exportExcel = async (params?: FetchPermintaanParams) => {
    try {
      return await permintaanAPI.exportExcel(params || lastParams.value)
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to export requests'
      throw err
    }
  }

  return {
    requests,
    selectedRequest,
    pagination,
    isLoading,
    error,
    fetchAll,
    fetchById,
    create,
    update,
    updateStatus,
    deleteRequest,
    exportExcel,
  }
})
