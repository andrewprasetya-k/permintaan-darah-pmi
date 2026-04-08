import { defineStore } from 'pinia'
import { ref } from 'vue'
import { komponenAPI, type CreateKomponenRequest, type UpdateKomponenRequest } from '@/api/komponen'
import type { KomponenDarah } from '@/types/models'

export const useKomponenStore = defineStore('komponen', () => {
  const komponens = ref<KomponenDarah[]>([])
  const selectedKomponen = ref<KomponenDarah | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const fetchAll = async (params?: Record<string, any>) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await komponenAPI.getAll(params)
      komponens.value = response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch components'
    } finally {
      isLoading.value = false
    }
  }

  const fetchById = async (id: number) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await komponenAPI.getById(id)
      selectedKomponen.value = response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch component'
    } finally {
      isLoading.value = false
    }
  }

  const create = async (data: CreateKomponenRequest) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await komponenAPI.create(data)
      komponens.value.push(response.data)
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to create component'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const update = async (id: number, data: UpdateKomponenRequest) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await komponenAPI.update(id, data)
      const index = komponens.value.findIndex((k) => k.komponenId === id)
      if (index !== -1) {
        komponens.value[index] = response.data
      }
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to update component'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const deleteKomponen = async (id: number) => {
    isLoading.value = true
    error.value = null
    try {
      await komponenAPI.delete(id)
      komponens.value = komponens.value.filter((k) => k.komponenId !== id)
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to delete component'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const activate = async (id: number) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await komponenAPI.activate(id)
      const index = komponens.value.findIndex((k) => k.komponenId === id)
      if (index !== -1) {
        komponens.value[index] = response.data
      }
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to activate component'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const deactivate = async (id: number) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await komponenAPI.deactivate(id)
      const index = komponens.value.findIndex((k) => k.komponenId === id)
      if (index !== -1) {
        komponens.value[index] = response.data
      }
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to deactivate component'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  return {
    komponens,
    selectedKomponen,
    isLoading,
    error,
    fetchAll,
    fetchById,
    create,
    update,
    deleteKomponen,
    activate,
    deactivate,
  }
})
