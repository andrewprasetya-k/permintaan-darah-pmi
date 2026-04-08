import { defineStore } from 'pinia'
import { ref } from 'vue'
import { adminAPI, type CreateAdminRequest, type UpdateAdminRequest } from '@/api/admin'
import type { User } from '@/types/models'

export const useAdminStore = defineStore('admin', () => {
  const admins = ref<User[]>([])
  const selectedAdmin = ref<User | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const fetchAll = async (params?: Record<string, any>) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await adminAPI.getAll(params)
      admins.value = response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch admins'
    } finally {
      isLoading.value = false
    }
  }

  const fetchById = async (id: string) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await adminAPI.getById(id)
      selectedAdmin.value = response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch admin'
    } finally {
      isLoading.value = false
    }
  }

  const create = async (data: CreateAdminRequest) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await adminAPI.create(data)
      admins.value.push(response.data)
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to create admin'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const update = async (id: string, data: UpdateAdminRequest) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await adminAPI.update(id, data)
      const index = admins.value.findIndex((a) => a.id === id)
      if (index !== -1) {
        admins.value[index] = response.data
      }
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to update admin'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const deleteAdmin = async (id: string) => {
    isLoading.value = true
    error.value = null
    try {
      await adminAPI.delete(id)
      admins.value = admins.value.filter((a) => a.id !== id)
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to delete admin'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const restore = async (id: string) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await adminAPI.restore(id)
      admins.value.push(response.data)
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to restore admin'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  return {
    admins,
    selectedAdmin,
    isLoading,
    error,
    fetchAll,
    fetchById,
    create,
    update,
    deleteAdmin,
    restore,
  }
})
