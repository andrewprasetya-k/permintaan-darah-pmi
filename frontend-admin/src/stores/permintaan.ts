import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { permintaanAPI, type CreatePermintaanRequest, type UpdatePermintaanRequest } from '@/api/permintaan'
import type { BloodRequest } from '@/types/models'

export const usePermintaanStore = defineStore('permintaan', () => {
  const requests = ref<BloodRequest[]>([])
  const myRequests = ref<BloodRequest[]>([])
  const selectedRequest = ref<BloodRequest | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const fetchAll = async (params?: Record<string, any>) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await permintaanAPI.getAll(params)
      requests.value = response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch requests'
    } finally {
      isLoading.value = false
    }
  }

  const fetchMyRequests = async (params?: Record<string, any>) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await permintaanAPI.getMyRequests(params)
      myRequests.value = response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch my requests'
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
      const index = requests.value.findIndex((r) => r.id === id)
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

  const updateMyRequest = async (id: string, data: UpdatePermintaanRequest) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await permintaanAPI.updateMyRequest(id, data)
      const index = myRequests.value.findIndex((r) => r.id === id)
      if (index !== -1) {
        myRequests.value[index] = response.data
      }
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to update my request'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const updateStatus = async (id: string, status: BloodRequest['status']) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await permintaanAPI.updateStatus(id, status)
      const index = requests.value.findIndex((r) => r.id === id)
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
      requests.value = requests.value.filter((r) => r.id !== id)
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to delete request'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  return {
    requests,
    myRequests,
    selectedRequest,
    isLoading,
    error,
    fetchAll,
    fetchMyRequests,
    fetchById,
    create,
    update,
    updateMyRequest,
    updateStatus,
    deleteRequest,
  }
})
