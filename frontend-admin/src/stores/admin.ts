import { defineStore } from 'pinia'
import { ref } from 'vue'
import { adminAPI, type CreateAdminRequest, type UpdateAdminRequest } from '@/api/admin'
import type { Admin, PaginationMeta } from '@/types/models'

export type AdminFilterStatus = 'active' | 'deleted' | 'all'

export const useAdminStore = defineStore('admin', () => {
  const admins = ref<Admin[]>([])
  const selectedAdmin = ref<Admin | null>(null)
  const myProfile = ref<Admin | null>(null)
  const currentFilter = ref<AdminFilterStatus>('active')
  const pagination = ref<PaginationMeta | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const fetchAll = async (params?: Record<string, any>) => {
    isLoading.value = true
    error.value = null
    try {
      const nextFilter = (params?.status as AdminFilterStatus | undefined) ?? currentFilter.value
      currentFilter.value = nextFilter
      const response = await adminAPI.getAll({ ...params, status: nextFilter })
      admins.value = response.data
      pagination.value = response.pagination ?? null
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
      const index = admins.value.findIndex((a) => a.adminId === id)
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
      await fetchAll({ status: currentFilter.value })
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
      await fetchAll({ status: currentFilter.value })
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to restore admin'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const fetchMe = async () => {
    isLoading.value = true
    error.value = null
    try {
      const response = await adminAPI.getMe()
      myProfile.value = response.data
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch profile'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const updateMe = async (data: UpdateAdminRequest) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await adminAPI.updateMe(data)
      myProfile.value = response.data
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to update profile'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  return {
    admins,
    selectedAdmin,
    myProfile,
    currentFilter,
    pagination,
    isLoading,
    error,
    fetchAll,
    fetchById,
    create,
    update,
    deleteAdmin,
    restore,
    fetchMe,
    updateMe,
  }
})
