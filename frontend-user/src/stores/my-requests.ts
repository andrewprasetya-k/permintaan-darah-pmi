import { computed, ref } from 'vue'
import { defineStore } from 'pinia'
import { myRequestsAPI, type FetchMyRequestsParams } from '@/api/my-requests'
import type {
  CreatePermintaanRequest,
  PermintaanDarah,
  PermintaanDarahListItem,
  PermintaanStatus,
  UpdatePermintaanRequest,
} from '@/types/models'

const cancelableStatuses: PermintaanStatus[] = ['dibuat', 'diproses', 'bisa_diambil']

const toListItem = (request: PermintaanDarah | PermintaanDarahListItem): PermintaanDarahListItem => ({
  permintaanDarahId: request.permintaanDarahId,
  namaPasien: request.namaPasien,
  jenisKelamin: request.jenisKelamin,
  golonganDarah: request.golonganDarah,
  rhesusDarah: request.rhesusDarah,
  tanggalPermintaan: request.tanggalPermintaan,
  status: request.status,
  createdAt: request.createdAt,
  updatedAt: request.updatedAt,
  deletedAt: request.deletedAt,
})

export const useMyRequestsStore = defineStore('myRequests', () => {
  const requests = ref<PermintaanDarahListItem[]>([])
  const selectedRequest = ref<PermintaanDarah | null>(null)
  const isLoading = ref(false)
  const isSubmitting = ref(false)
  const isDownloading = ref(false)
  const error = ref<string | null>(null)
  const pagination = ref<FetchMyRequestsParams | null>(null)

  const activeRequests = computed(() =>
    requests.value.filter((request) => request.status !== 'selesai' && request.status !== 'dibatalkan'),
  )

  const recentRequests = computed(() => requests.value.slice(0, 5))

  const replaceInList = (request: PermintaanDarah | PermintaanDarahListItem) => {
    const item = toListItem(request)
    const index = requests.value.findIndex(
      (existing) => existing.permintaanDarahId === item.permintaanDarahId,
    )

    if (index >= 0) {
      requests.value[index] = item
    } else {
      requests.value.unshift(item)
    }
  }

  const fetchAll = async (params: FetchMyRequestsParams = {}) => {
    isLoading.value = true
    error.value = null

    try {
      const response = await myRequestsAPI.getAll(params)
      requests.value = response.data
      pagination.value = {
        limit: response.pagination?.limit ?? params.limit ?? 100,
        offset: response.pagination?.offset ?? params.offset ?? 0,
      }
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Gagal memuat permintaan'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const fetchById = async (id: string) => {
    isLoading.value = true
    error.value = null

    try {
      const response = await myRequestsAPI.getById(id)
      selectedRequest.value = response.data
      replaceInList(response.data)
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Gagal memuat detail permintaan'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const create = async (payload: CreatePermintaanRequest) => {
    isSubmitting.value = true
    error.value = null

    try {
      const response = await myRequestsAPI.create(payload)
      selectedRequest.value = response.data
      replaceInList(response.data)
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Gagal membuat permintaan'
      throw err
    } finally {
      isSubmitting.value = false
    }
  }

  const updateMyRequest = async (id: string, payload: UpdatePermintaanRequest) => {
    isSubmitting.value = true
    error.value = null

    try {
      const response = await myRequestsAPI.updateMyRequest(id, payload)
      selectedRequest.value = response.data
      replaceInList(response.data)
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Gagal menyimpan permintaan'
      throw err
    } finally {
      isSubmitting.value = false
    }
  }

  const cancelRequest = async (id: string, reason: string) => {
    isSubmitting.value = true
    error.value = null

    try {
      const response = await myRequestsAPI.updateStatus(id, {
        status: 'dibatalkan',
        reason,
      })
      selectedRequest.value = response.data
      replaceInList(response.data)
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Gagal membatalkan permintaan'
      throw err
    } finally {
      isSubmitting.value = false
    }
  }

  const confirmPickup = async (id: string) => {
    isSubmitting.value = true
    error.value = null

    try {
      const response = await myRequestsAPI.updateStatus(id, {
        status: 'selesai',
      })
      selectedRequest.value = response.data
      replaceInList(response.data)
      return response.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Gagal mengonfirmasi pengambilan'
      throw err
    } finally {
      isSubmitting.value = false
    }
  }

  const canEdit = (request: PermintaanDarah | PermintaanDarahListItem | null) =>
    request?.status === 'dibuat'

  const canCancel = (request: PermintaanDarah | PermintaanDarahListItem | null) =>
    Boolean(request && cancelableStatuses.includes(request.status))

  const canConfirmPickup = (request: PermintaanDarah | PermintaanDarahListItem | null) =>
    request?.status === 'bisa_diambil'

  const downloadBlanko = async (id: string) => {
    isDownloading.value = true
    error.value = null

    try {
      return await myRequestsAPI.downloadBlanko(id)
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Gagal mengunduh blanko'
      throw err
    } finally {
      isDownloading.value = false
    }
  }

  return {
    requests,
    selectedRequest,
    isLoading,
    isSubmitting,
    isDownloading,
    error,
    pagination,
    activeRequests,
    recentRequests,
    fetchAll,
    fetchById,
    create,
    updateMyRequest,
    cancelRequest,
    confirmPickup,
    downloadBlanko,
    canEdit,
    canCancel,
    canConfirmPickup,
  }
})
