import { defineStore } from 'pinia'
import { ref } from 'vue'
import { rumahSakitAPI, type CreateRumahSakitRequest, type UpdateRumahSakitRequest } from '@/api/rumah-sakit'
import type { RumahSakit } from '@/types/models'

export const useRumahSakitStore = defineStore('rumahSakit', () => {
  const hospitals = ref<RumahSakit[]>([])
  const selectedHospital = ref<RumahSakit | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const fetchAll = async (params?: Record<string, any>) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await rumahSakitAPI.getAll(params)
      hospitals.value = response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch hospitals'
    } finally {
      isLoading.value = false
    }
  }

  const fetchById = async (id: string) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await rumahSakitAPI.getById(id)
      selectedHospital.value = response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch hospital'
    } finally {
      isLoading.value = false
    }
  }

  const create = async (data: CreateRumahSakitRequest) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await rumahSakitAPI.create(data)
      hospitals.value.push(response.data)
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to create hospital'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const update = async (id: string, data: UpdateRumahSakitRequest) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await rumahSakitAPI.update(id, data)
      const index = hospitals.value.findIndex((h) => h.rumahSakitId === id)
      if (index !== -1) {
        hospitals.value[index] = response.data
      }
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to update hospital'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const deleteHospital = async (id: string) => {
    isLoading.value = true
    error.value = null
    try {
      await rumahSakitAPI.delete(id)
      hospitals.value = hospitals.value.filter((h) => h.rumahSakitId !== id)
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to delete hospital'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const restore = async (id: string) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await rumahSakitAPI.restore(id)
      hospitals.value.push(response.data)
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to restore hospital'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  return {
    hospitals,
    selectedHospital,
    isLoading,
    error,
    fetchAll,
    fetchById,
    create,
    update,
    deleteHospital,
    restore,
  }
})
