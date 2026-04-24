<script setup lang="ts">
import { computed, ref } from 'vue'
import { usePermintaanStore } from '@/stores/permintaan'
import { X, AlertCircle } from '@lucide/vue'
import { permintaanAPI } from '@/api/permintaan'
import AppModal from '@/components/feedback/AppModal.vue'
import AppFlag from '@/components/feedback/AppFlag.vue'

const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  close: []
  updated: []
}>()

const permintaanStore = usePermintaanStore()
const isUpdatingStatus = ref(false)
const statusError = ref<string | null>(null)
const showStatusDropdown = ref(false)
const pendingStatus = ref<string | null>(null)
const cancelReason = ref('')
const flag = ref<{ variant: 'success' | 'error'; title: string; message?: string } | null>(null)

const subtitle = computed(() => permintaanStore.selectedRequest?.namaPasien)

const statusStyle: Record<string, string> = {
  dibuat: 'bg-amber-50 text-amber-600',
  diproses: 'bg-blue-50 text-blue-600',
  bisa_diambil: 'bg-violet-50 text-violet-600',
  selesai: 'bg-green-50 text-green-600',
  dibatalkan: 'bg-red-50 text-red-600',
}

// Valid status transitions
const validNextStatuses = computed(() => {
  const current = permintaanStore.selectedRequest?.status
  const transitions: Record<string, string[]> = {
    dibuat: ['diproses', 'dibatalkan'],
    diproses: ['bisa_diambil', 'dibatalkan'],
    bisa_diambil: ['selesai', 'dibatalkan'],
    selesai: [],
    dibatalkan: [],
  }
  return transitions[current || 'dibuat'] || []
})

const updateStatus = async (newStatus: string) => {
  if (!permintaanStore.selectedRequest) return

  if (newStatus === 'dibatalkan') {
    pendingStatus.value = newStatus
    cancelReason.value = ''
    statusError.value = null
    showStatusDropdown.value = false
    return
  }

  await submitStatusUpdate(newStatus)
}

const closeCancellationDialog = () => {
  pendingStatus.value = null
  cancelReason.value = ''
  statusError.value = null
}

const submitStatusUpdate = async (newStatus: string, reason?: string) => {
  if (!permintaanStore.selectedRequest) return

  isUpdatingStatus.value = true
  statusError.value = null

  try {
    const response = await permintaanAPI.updateStatus(
      permintaanStore.selectedRequest.permintaanDarahId,
      { status: newStatus as any, reason },
    )
    permintaanStore.selectedRequest = response.data
    showStatusDropdown.value = false
    closeCancellationDialog()
    flag.value = {
      variant: 'success',
      title: 'Status diperbarui',
      message:
        newStatus === 'dibatalkan'
          ? 'Permintaan darah berhasil dibatalkan.'
          : `Status berhasil diubah menjadi ${newStatus}.`,
    }
    emit('updated')
  } catch (error: unknown) {
    statusError.value =
      error instanceof Error && error.message ? error.message : 'Gagal mengubah status'
    flag.value = {
      variant: 'error',
      title: 'Operasi gagal',
      message:
        error instanceof Error && error.message ? error.message : 'Gagal mengubah status permintaan darah.',
    }
  } finally {
    isUpdatingStatus.value = false
  }
}

const submitCancellation = async () => {
  if (!cancelReason.value.trim()) {
    statusError.value = 'Alasan pembatalan harus diisi'
    return
  }

  await submitStatusUpdate('dibatalkan', cancelReason.value.trim())
}
</script>

<template>
  <Teleport to="body">
    <AppFlag
      v-if="flag"
      :variant="flag.variant"
      :title="flag.title"
      :message="flag.message"
      @close="flag = null"
    />

    <AppModal
      :is-open="pendingStatus === 'dibatalkan'"
      title="Batalkan Permintaan Darah"
      description="Status dibatalkan memerlukan alasan yang akan dikirim ke payload backend."
      @close="closeCancellationDialog"
    >
      <textarea
        v-model="cancelReason"
        rows="4"
        placeholder="Masukkan alasan pembatalan"
        class="w-full rounded-xl border border-gray-200 bg-gray-50 px-3.5 py-2.5 text-sm text-gray-900 outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
      />

      <div
        v-if="statusError && pendingStatus === 'dibatalkan'"
        class="mt-3 flex items-center gap-2 rounded-xl border border-red-100 bg-red-50 p-3 text-xs text-red-600"
      >
        <AlertCircle :size="14" class="shrink-0" />
        {{ statusError }}
      </div>

      <template #footer>
        <button
          type="button"
          class="flex-1 rounded-xl bg-gray-100 px-4 py-2.5 text-sm font-medium text-gray-700 transition-colors hover:bg-gray-200"
          @click="closeCancellationDialog"
        >
          Tutup
        </button>
        <button
          type="button"
          :disabled="isUpdatingStatus"
          class="flex-1 rounded-xl bg-red-600 px-4 py-2.5 text-sm font-medium text-white transition-colors hover:bg-red-700 disabled:opacity-50"
          @click="submitCancellation"
        >
          {{ isUpdatingStatus ? 'Memproses...' : 'Batalkan Permintaan' }}
        </button>
      </template>
    </AppModal>

    <Transition name="backdrop">
      <div
        v-if="isOpen"
        class="fixed inset-0 bg-black/20 backdrop-blur-sm z-40"
        @click="$emit('close')"
      />
    </Transition>

    <Transition name="drawer">
      <div
        v-if="isOpen"
        class="fixed top-0 right-0 h-full bg-white shadow-2xl z-50 flex flex-col w-full lg:w-4/5 lg:rounded-tl-3xl lg:rounded-bl-3xl"
      >
        <!-- Header -->
        <div class="flex items-center justify-between px-10 py-8 border-b border-gray-200">
          <div>
            <h2 class="text-xl font-semibold text-gray-900">Detail Permintaan Darah</h2>
            <p class="text-sm text-gray-400 mt-0.5">{{ subtitle }}</p>
          </div>
          <button
            @click="$emit('close')"
            class="p-2 rounded-lg hover:bg-gray-100 text-gray-400 transition-colors"
          >
            <X :size="20" />
          </button>
        </div>

        <!-- Content -->
        <div v-if="permintaanStore.selectedRequest" class="flex-1 overflow-y-auto px-10 py-8">
          <div class="space-y-6 max-w-full">
            <div class="space-y-4">
              <!-- Nama Pasien -->
              <div>
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2">
                  Nama Pasien
                </label>
                <p class="text-base text-gray-900 font-medium">
                  {{ permintaanStore.selectedRequest.namaPasien }}
                </p>
              </div>

              <!-- No. RM & Tempat Lahir -->
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label
                    class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2"
                  >
                    No. Rekam Medis
                  </label>
                  <p class="text-base text-gray-900 font-medium">
                    {{ permintaanStore.selectedRequest.noRM || '-' }}
                  </p>
                </div>
                <div>
                  <label
                    class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2"
                  >
                    Tempat Lahir
                  </label>
                  <p class="text-base text-gray-900 font-medium">
                    {{ permintaanStore.selectedRequest.tempatLahir }}
                  </p>
                </div>
              </div>

              <!-- Jenis Kelamin & Tanggal Lahir -->
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label
                    class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2"
                  >
                    Jenis Kelamin
                  </label>
                  <p class="text-base text-gray-900 font-medium">
                    {{
                      permintaanStore.selectedRequest.jenisKelamin === 'L'
                        ? 'Laki-laki'
                        : 'Perempuan'
                    }}
                  </p>
                </div>
                <div>
                  <label
                    class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2"
                  >
                    Tanggal Lahir
                  </label>
                  <p class="text-base text-gray-900 font-medium">
                    {{
                      permintaanStore.selectedRequest.tanggalLahir
                        ? new Date(permintaanStore.selectedRequest.tanggalLahir).toLocaleDateString(
                            'id-ID',
                            {
                              year: 'numeric',
                              month: 'long',
                              day: 'numeric',
                            },
                          )
                        : '-'
                    }}
                  </p>
                </div>
              </div>

              <!-- Golongan Darah & Rhesus -->
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label
                    class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2"
                  >
                    Golongan Darah
                  </label>
                  <span
                    class="inline-flex items-center gap-1 px-2.5 py-1 bg-red-50 text-red-600 text-xs font-semibold rounded-lg"
                  >
                    {{ permintaanStore.selectedRequest.golonganDarah }}
                    {{ permintaanStore.selectedRequest.rhesusDarah }}
                  </span>
                </div>
                <div>
                  <label
                    class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2"
                  >
                    Hemoglobin
                  </label>
                  <p class="text-base text-gray-900 font-medium">
                    {{ permintaanStore.selectedRequest.hemoglobin }} g/dL
                  </p>
                </div>
              </div>

              <!-- Ruang/Bagian/Kelas -->
              <div>
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2">
                  Ruang/Bagian/Kelas
                </label>
                <p class="text-base text-gray-900 font-medium">
                  {{ permintaanStore.selectedRequest.ruangBagianKelas || '-' }}
                </p>
              </div>

              <!-- Pernah Transfusi -->
              <div>
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2">
                  Riwayat Transfusi
                </label>
                <p class="text-base text-gray-900 font-medium">
                  {{ permintaanStore.selectedRequest.pernahTransfusi ? 'Pernah' : 'Tidak Pernah' }}
                </p>
              </div>

              <!-- Indikasi Transfusi -->
              <div>
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2">
                  Indikasi Transfusi
                </label>
                <p class="text-base text-gray-900 font-medium">
                  {{ permintaanStore.selectedRequest.indikasiTransfusi || '-' }}
                </p>
              </div>

              <!-- Pernah Hamil -->
              <div>
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2">
                  Pernah Hamil
                </label>
                <p class="text-base text-gray-900 font-medium">
                  {{ permintaanStore.selectedRequest.pernahHamil || '-' }}
                </p>
              </div>

              <hr class="my-4 border-gray-100" />

              <!-- Tanggal Permintaan -->
              <div>
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2">
                  Tanggal Permintaan
                </label>
                <p class="text-sm text-gray-600">
                  {{
                    permintaanStore.selectedRequest.tanggalPermintaan
                      ? new Date(
                          permintaanStore.selectedRequest.tanggalPermintaan,
                        ).toLocaleDateString('id-ID', {
                          weekday: 'long',
                          year: 'numeric',
                          month: 'long',
                          day: 'numeric',
                          hour: '2-digit',
                          minute: '2-digit',
                        })
                      : '-'
                  }}
                </p>
              </div>

              <!-- Status -->
              <div>
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2">
                  Status
                </label>
                <div class="relative">
                  <div class="flex items-center gap-2">
                    <span
                      class="inline-flex items-center px-3 py-1.5 rounded-lg text-xs font-medium capitalize"
                      :class="
                        statusStyle[permintaanStore.selectedRequest.status] ??
                        'bg-gray-50 text-gray-600'
                      "
                    >
                      {{ permintaanStore.selectedRequest.status }}
                    </span>
                  </div>
                </div>

                <!-- Status Error -->
                <div
                  v-if="statusError"
                  class="mt-2 flex items-center gap-2 p-2 bg-red-50 border border-red-100 text-red-600 rounded-lg text-xs"
                >
                  <AlertCircle :size="14" class="shrink-0" />
                  {{ statusError }}
                </div>
              </div>

              <!-- Cancel Reason (jika ada) -->
              <div v-if="permintaanStore.selectedRequest.cancelReason">
                <label class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2">
                  Alasan Pembatalan
                </label>
                <p class="text-base text-red-600 font-medium">
                  {{ permintaanStore.selectedRequest.cancelReason }}
                </p>
              </div>

              <!-- Tanggal -->
              <div class="grid grid-cols-2 gap-6">
                <div>
                  <label
                    class="block text-xs font-medium text-gray-500 uppercase tracking-wide mb-2"
                  >
                    Tanggal Permintaan Dibuat
                  </label>
                  <p class="text-sm text-gray-600">
                    {{
                      new Date(permintaanStore.selectedRequest.createdAt).toLocaleDateString(
                        'id-ID',
                        {
                          weekday: 'long',
                          year: 'numeric',
                          month: 'long',
                          day: 'numeric',
                          hour: '2-digit',
                          minute: '2-digit',
                        },
                      )
                    }}
                  </p>
                </div>
              </div>

              <!-- Detail Permintaan Darah (Komponen yang Diminta) -->
              <div
                v-if="
                  permintaanStore.selectedRequest.detailPermintaanDarah &&
                  permintaanStore.selectedRequest.detailPermintaanDarah.length > 0
                "
                class="border-t pt-4 mt-4"
              >
                <h3 class="text-sm font-semibold text-gray-900 mb-3">
                  Komponen Darah yang Diminta
                </h3>
                <div class="space-y-2">
                  <div
                    v-for="detail in permintaanStore.selectedRequest.detailPermintaanDarah"
                    :key="detail.detailId"
                    class="flex items-center justify-between p-3 bg-gray-50 rounded-lg"
                  >
                    <div>
                      <p class="text-xs font-medium text-gray-600">
                        {{ detail.komponenNama || 'Komponen' }}
                      </p>
                      <p class="text-xs text-gray-500">
                        {{ detail.golonganDarah }}{{ detail.rhesusDarah }}
                      </p>
                    </div>
                    <div class="text-right">
                      <p class="text-sm font-semibold text-gray-900">{{ detail.jumlahKantong }}</p>
                      <p class="text-xs text-gray-500">kantong</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Close Button -->
            <div class="flex gap-3 pt-4">
              <button
                @click="$emit('close')"
                class="flex-1 py-2.5 bg-gray-100 hover:bg-gray-200 text-gray-600 text-sm font-medium rounded-xl transition-colors"
              >
                Tutup
              </button>
              <div class="relative flex-1">
                <button
                  v-if="validNextStatuses.length > 0"
                  @click="showStatusDropdown = !showStatusDropdown"
                  :disabled="isUpdatingStatus"
                  class="w-full py-2.5 bg-blue-600 hover:bg-blue-700 text-white text-sm font-medium rounded-xl transition-colors disabled:opacity-50"
                >
                  Ubah Status
                </button>
                <!-- Status Dropdown Menu - Muncul ke atas -->
                <div
                  v-if="showStatusDropdown && validNextStatuses.length > 0"
                  class="absolute bottom-full left-0 mb-2 bg-white border border-gray-200 rounded-lg shadow-lg z-20 w-full"
                >
                  <div
                    v-for="status in validNextStatuses"
                    :key="status"
                    @click="updateStatus(status)"
                    class="px-4 py-2 text-sm text-gray-700 hover:bg-blue-50 cursor-pointer transition-colors first:rounded-t-lg last:rounded-b-lg"
                  >
                    {{
                      status === 'dibatalkan'
                        ? 'Batalkan'
                        : status === 'bisa_diambil'
                          ? 'Bisa Diambil'
                          : status.charAt(0).toUpperCase() + status.slice(1)
                    }}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.backdrop-enter-active,
.backdrop-leave-active {
  transition: opacity 0.3s ease;
}
.backdrop-enter-from,
.backdrop-leave-to {
  opacity: 0;
}

.drawer-enter-active,
.drawer-leave-active {
  transition: transform 0.3s ease;
}
.drawer-enter-from,
.drawer-leave-to {
  transform: translateX(100%);
}
</style>
