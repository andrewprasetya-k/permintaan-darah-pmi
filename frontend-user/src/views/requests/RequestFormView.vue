<script setup lang="ts">
import { computed, onMounted, reactive, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import AppModal from '@/components/feedback/AppModal.vue'
import { useAuthStore } from '@/stores/auth'
import { useFeedbackStore } from '@/stores/feedback'
import { useKomponenStore } from '@/stores/komponen'
import { useMyProfileStore } from '@/stores/my-profile'
import { useMyRequestsStore } from '@/stores/my-requests'
import type {
  BloodType,
  CreateDetailPermintaanDarahRequest,
  CreatePermintaanRequest,
  Gender,
  PermintaanStatus,
  Rhesus,
  UpdatePermintaanRequest,
} from '@/types/models'
import { toDateInputValue, toIsoDate } from '@/utils/format'

interface DetailFormRow {
  key: number
  komponenDarahId: string
  golonganDarah: BloodType
  rhesusDarah: Rhesus
  jumlahKantong: string
}

interface FormState {
  namaPasien: string
  noRM: string
  tempatLahir: string
  tanggalLahir: string
  jenisKelamin: Gender | ''
  golonganDarah: BloodType | ''
  rhesusDarah: Rhesus | ''
  hemoglobin: string
  ruangBagianKelas: string
  pernahTransfusi: boolean
  indikasiTransfusi: string
  pernahHamil: string
  tanggalPermintaan: string
  status: PermintaanStatus
}

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const profileStore = useMyProfileStore()
const requestsStore = useMyRequestsStore()
const komponenStore = useKomponenStore()
const feedbackStore = useFeedbackStore()

const isConfirmOpen = ref(false)
const detailKey = ref(1)
const detailRows = ref<DetailFormRow[]>([])

const form = reactive<FormState>({
  namaPasien: '',
  noRM: '',
  tempatLahir: '',
  tanggalLahir: '',
  jenisKelamin: '',
  golonganDarah: '',
  rhesusDarah: '',
  hemoglobin: '',
  ruangBagianKelas: '',
  pernahTransfusi: false,
  indikasiTransfusi: '',
  pernahHamil: '',
  tanggalPermintaan: toDateInputValue(new Date().toISOString()),
  status: 'dibuat',
})

const isEdit = computed(() => route.name === 'request-edit')
const requestId = computed(() => String(route.params.id || ''))
const title = computed(() => (isEdit.value ? 'Edit Permintaan' : 'Buat Permintaan'))
const submitLabel = computed(() => (isEdit.value ? 'Simpan Perubahan' : 'Kirim Permintaan'))
const activeComponents = computed(() =>
  komponenStore.components.filter((component) => component.isActive !== false),
)

const optionalString = (value: string) => {
  const trimmed = value.trim()
  return trimmed ? trimmed : undefined
}

const optionalNumber = (value: string) => {
  const parsed = Number(value)
  return Number.isFinite(parsed) ? parsed : undefined
}

const addDetailRow = () => {
  const firstComponent = activeComponents.value[0]
  detailRows.value.push({
    key: detailKey.value,
    komponenDarahId: firstComponent ? String(firstComponent.komponenId) : '',
    golonganDarah: form.golonganDarah || 'O',
    rhesusDarah: form.rhesusDarah || '+',
    jumlahKantong: '1',
  })
  detailKey.value += 1
}

const removeDetailRow = (key: number) => {
  detailRows.value = detailRows.value.filter((row) => row.key !== key)
}

const buildDetails = (): CreateDetailPermintaanDarahRequest[] | undefined => {
  const details = detailRows.value
    .filter((row) => row.komponenDarahId && row.jumlahKantong)
    .map((row) => ({
      komponenDarahId: Number(row.komponenDarahId),
      golonganDarah: row.golonganDarah,
      rhesusDarah: row.rhesusDarah,
      jumlahKantong: Number(row.jumlahKantong),
    }))

  return details.length > 0 ? details : undefined
}

const buildBasePayload = () => ({
  namaPasien: form.namaPasien.trim(),
  noRM: optionalString(form.noRM),
  tempatLahir: form.tempatLahir.trim(),
  tanggalLahir: toIsoDate(form.tanggalLahir),
  jenisKelamin: form.jenisKelamin as Gender,
  golonganDarah: form.golonganDarah || undefined,
  rhesusDarah: form.rhesusDarah || undefined,
  hemoglobin: optionalNumber(form.hemoglobin),
  ruangBagianKelas: optionalString(form.ruangBagianKelas),
  pernahTransfusi: form.pernahTransfusi,
  indikasiTransfusi: optionalString(form.indikasiTransfusi),
  pernahHamil: optionalString(form.pernahHamil),
  status: form.status,
  tanggalPermintaan: toIsoDate(form.tanggalPermintaan),
})

const populateForm = () => {
  const request = requestsStore.selectedRequest
  if (!request) {
    return
  }

  form.namaPasien = request.namaPasien
  form.noRM = request.noRM || ''
  form.tempatLahir = request.tempatLahir
  form.tanggalLahir = toDateInputValue(request.tanggalLahir)
  form.jenisKelamin = request.jenisKelamin
  form.golonganDarah = request.golonganDarah || ''
  form.rhesusDarah = request.rhesusDarah || ''
  form.hemoglobin = request.hemoglobin ? String(request.hemoglobin) : ''
  form.ruangBagianKelas = request.ruangBagianKelas || ''
  form.pernahTransfusi = request.pernahTransfusi
  form.indikasiTransfusi = request.indikasiTransfusi || ''
  form.pernahHamil = request.pernahHamil || ''
  form.tanggalPermintaan = toDateInputValue(request.tanggalPermintaan)
  form.status = request.status
}

const load = async () => {
  await Promise.allSettled([
    komponenStore.fetchAll(),
    profileStore.profile ? Promise.resolve(profileStore.profile) : profileStore.fetchMe(),
  ])

  if (!isEdit.value || !requestId.value) {
    return
  }

  try {
    const request = await requestsStore.fetchById(requestId.value)
    if (!requestsStore.canEdit(request)) {
      feedbackStore.showFlag({
        title: 'Permintaan tidak bisa diedit',
        message: 'Hanya permintaan berstatus dibuat yang bisa diedit rumah sakit.',
        variant: 'info',
      })
      await router.push(`/requests/${request.permintaanDarahId}`)
      return
    }
    populateForm()
  } catch (err) {
    feedbackStore.showFlag({
      title: 'Gagal memuat permintaan',
      message: err instanceof Error ? err.message : 'Coba muat ulang halaman.',
      variant: 'error',
    })
  }
}

const openConfirm = () => {
  isConfirmOpen.value = true
}

const submit = async () => {
  isConfirmOpen.value = false

  try {
    if (isEdit.value) {
      const rumahSakitId =
        requestsStore.selectedRequest?.rumahSakitId ||
        profileStore.profile?.rumahSakitId ||
        authStore.user?.id

      if (!rumahSakitId) {
        throw new Error('ID rumah sakit tidak tersedia')
      }

      const payload: UpdatePermintaanRequest = {
        ...buildBasePayload(),
        rumahSakitId,
      }
      const updated = await requestsStore.updateMyRequest(requestId.value, payload)
      feedbackStore.showFlag({
        title: 'Permintaan disimpan',
        message: 'Perubahan permintaan berhasil dikirim.',
        variant: 'success',
      })
      await router.push(`/requests/${updated.permintaanDarahId}`)
      return
    }

    const payload: CreatePermintaanRequest = {
      ...buildBasePayload(),
      details: buildDetails(),
    }
    const created = await requestsStore.create(payload)
    feedbackStore.showFlag({
      title: 'Permintaan dibuat',
      message: 'Permintaan darah berhasil dikirim ke PMI.',
      variant: 'success',
    })
    await router.push(`/requests/${created.permintaanDarahId}`)
  } catch (err) {
    feedbackStore.showFlag({
      title: isEdit.value ? 'Gagal menyimpan' : 'Gagal membuat permintaan',
      message: err instanceof Error ? err.message : 'Periksa kembali data permintaan.',
      variant: 'error',
    })
  }
}

onMounted(load)
watch(requestId, load)
</script>

<template>
  <section>
    <div class="page-header">
      <div>
        <p class="page-eyebrow">Permintaan</p>
        <h1 class="page-title">{{ title }}</h1>
        <p class="page-subtitle">
          Isi data pasien dan kebutuhan komponen darah sesuai kontrak backend permintaan darah.
        </p>
      </div>
      <RouterLink class="btn btn-secondary" :to="isEdit ? `/requests/${requestId}` : '/requests'">
        Batal
      </RouterLink>
    </div>

    <form class="request-form" @submit.prevent="openConfirm">
      <section class="card form-section">
        <h2 class="section-title">Data Pasien</h2>
        <div class="form-grid">
          <div class="form-field">
            <label class="form-label" for="namaPasien">Nama pasien</label>
            <input id="namaPasien" v-model="form.namaPasien" class="form-input" required />
          </div>
          <div class="form-field">
            <label class="form-label" for="noRM">No. RM</label>
            <input id="noRM" v-model="form.noRM" class="form-input" />
          </div>
          <div class="form-field">
            <label class="form-label" for="tempatLahir">Tempat lahir</label>
            <input id="tempatLahir" v-model="form.tempatLahir" class="form-input" required />
          </div>
          <div class="form-field">
            <label class="form-label" for="tanggalLahir">Tanggal lahir</label>
            <input
              id="tanggalLahir"
              v-model="form.tanggalLahir"
              class="form-input"
              type="date"
              required
            />
          </div>
          <div class="form-field">
            <label class="form-label" for="jenisKelamin">Jenis kelamin</label>
            <select id="jenisKelamin" v-model="form.jenisKelamin" class="form-select" required>
              <option value="" disabled>Pilih jenis kelamin</option>
              <option value="L">Laki-laki</option>
              <option value="P">Perempuan</option>
            </select>
          </div>
          <div class="form-field blood-fields">
            <div>
              <label class="form-label" for="golonganDarah">Golongan darah</label>
              <select id="golonganDarah" v-model="form.golonganDarah" class="form-select">
                <option value="">Belum diketahui</option>
                <option value="A">A</option>
                <option value="B">B</option>
                <option value="AB">AB</option>
                <option value="O">O</option>
              </select>
            </div>
            <div>
              <label class="form-label" for="rhesusDarah">Rhesus</label>
              <select id="rhesusDarah" v-model="form.rhesusDarah" class="form-select">
                <option value="">-</option>
                <option value="+">+</option>
                <option value="-">-</option>
              </select>
            </div>
          </div>
        </div>
      </section>

      <section class="card form-section">
        <h2 class="section-title">Data Medis</h2>
        <div class="form-grid">
          <div class="form-field">
            <label class="form-label" for="hemoglobin">Hemoglobin</label>
            <input
              id="hemoglobin"
              v-model="form.hemoglobin"
              class="form-input"
              type="number"
              step="0.1"
              min="0"
            />
          </div>
          <div class="form-field">
            <label class="form-label" for="ruangBagianKelas">Ruang/Bagian/Kelas</label>
            <input id="ruangBagianKelas" v-model="form.ruangBagianKelas" class="form-input" />
          </div>
          <div class="form-field">
            <label class="form-label" for="tanggalPermintaan">Tanggal permintaan</label>
            <input
              id="tanggalPermintaan"
              v-model="form.tanggalPermintaan"
              class="form-input"
              type="date"
              required
            />
          </div>
          <div class="form-field checkbox-field">
            <label class="checkbox-label" for="pernahTransfusi">
              <input id="pernahTransfusi" v-model="form.pernahTransfusi" type="checkbox" />
              Pernah transfusi
            </label>
          </div>
          <div class="form-field form-field-full">
            <label class="form-label" for="indikasiTransfusi">Indikasi transfusi</label>
            <textarea
              id="indikasiTransfusi"
              v-model="form.indikasiTransfusi"
              class="form-textarea"
            />
          </div>
          <div class="form-field form-field-full">
            <label class="form-label" for="pernahHamil">Riwayat hamil</label>
            <input id="pernahHamil" v-model="form.pernahHamil" class="form-input" />
          </div>
        </div>
      </section>

      <section v-if="!isEdit" class="card form-section">
        <div class="component-header">
          <div>
            <h2 class="section-title">Komponen Darah</h2>
            <p class="form-help">Tambahkan kebutuhan komponen bila data sudah tersedia.</p>
          </div>
          <button
            type="button"
            class="btn btn-secondary"
            :disabled="komponenStore.isLoading || activeComponents.length === 0"
            @click="addDetailRow"
          >
            + Tambah Komponen
          </button>
        </div>

        <div v-if="komponenStore.error" class="inline-error">
          {{ komponenStore.error }}
        </div>

        <div v-if="detailRows.length === 0" class="empty-row">
          Belum ada detail komponen pada permintaan ini.
        </div>

        <div v-else class="detail-row-list">
          <div v-for="row in detailRows" :key="row.key" class="detail-row">
            <div class="form-field">
              <label class="form-label" :for="`komponen-${row.key}`">Komponen</label>
              <select
                :id="`komponen-${row.key}`"
                v-model="row.komponenDarahId"
                class="form-select"
                required
              >
                <option value="" disabled>Pilih komponen</option>
                <option
                  v-for="component in activeComponents"
                  :key="component.komponenId"
                  :value="String(component.komponenId)"
                >
                  {{ component.komponenDarah }} ({{ component.komponenKode }})
                </option>
              </select>
            </div>
            <div class="form-field">
              <label class="form-label" :for="`detail-gol-${row.key}`">Golongan</label>
              <select
                :id="`detail-gol-${row.key}`"
                v-model="row.golonganDarah"
                class="form-select"
                required
              >
                <option value="A">A</option>
                <option value="B">B</option>
                <option value="AB">AB</option>
                <option value="O">O</option>
              </select>
            </div>
            <div class="form-field">
              <label class="form-label" :for="`detail-rhesus-${row.key}`">Rhesus</label>
              <select
                :id="`detail-rhesus-${row.key}`"
                v-model="row.rhesusDarah"
                class="form-select"
                required
              >
                <option value="+">+</option>
                <option value="-">-</option>
              </select>
            </div>
            <div class="form-field">
              <label class="form-label" :for="`jumlah-${row.key}`">Kantong</label>
              <input
                :id="`jumlah-${row.key}`"
                v-model="row.jumlahKantong"
                class="form-input"
                min="1"
                type="number"
                required
              />
            </div>
            <button
              type="button"
              class="btn btn-danger btn-icon remove-row"
              @click="removeDetailRow(row.key)"
            >
              x
            </button>
          </div>
        </div>
      </section>

      <div class="form-actions">
        <RouterLink class="btn btn-secondary" :to="isEdit ? `/requests/${requestId}` : '/requests'">
          Batal
        </RouterLink>
        <button type="submit" class="btn btn-primary" :disabled="requestsStore.isSubmitting">
          {{ requestsStore.isSubmitting ? 'Menyimpan...' : submitLabel }}
        </button>
      </div>
    </form>

    <AppModal
      :is-open="isConfirmOpen"
      :title="isEdit ? 'Simpan perubahan?' : 'Kirim permintaan?'"
      :description="
        isEdit
          ? 'Perubahan data permintaan akan disimpan.'
          : 'Permintaan darah akan dikirim ke backend PMI.'
      "
      width="sm"
      @close="isConfirmOpen = false"
    >
      <template #footer>
        <button type="button" class="btn btn-secondary" @click="isConfirmOpen = false">
          Batal
        </button>
        <button
          type="button"
          class="btn btn-primary"
          :disabled="requestsStore.isSubmitting"
          @click="submit"
        >
          {{ submitLabel }}
        </button>
      </template>
    </AppModal>
  </section>
</template>

<style scoped>
.request-form {
  display: grid;
  gap: 18px;
}

.form-section {
  padding: 24px;
}

.form-section .section-title {
  margin-bottom: 16px;
}

.blood-fields {
  display: grid;
  grid-template-columns: minmax(0, 1fr) 120px;
  gap: 10px;
}

.checkbox-field {
  justify-content: end;
}

.checkbox-label {
  display: flex;
  min-height: 42px;
  align-items: center;
  gap: 10px;
  color: var(--text);
  font-weight: 600;
}

.checkbox-label input {
  width: 18px;
  height: 18px;
  accent-color: var(--blue);
}

.component-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 16px;
  margin-bottom: 14px;
}

.inline-error,
.empty-row {
  border: 1px dashed var(--line-strong);
  border-radius: 16px;
  padding: 14px;
  color: var(--text-soft);
}

.inline-error {
  margin-bottom: 12px;
  border-color: #fecaca;
  background: #fef2f2;
  color: #991b1b;
}

.detail-row-list {
  display: grid;
  gap: 12px;
}

.detail-row {
  display: grid;
  grid-template-columns: minmax(220px, 1fr) 120px 100px 120px 44px;
  gap: 10px;
  align-items: end;
  border: 1px solid #f3f4f6;
  border-radius: 16px;
  padding: 12px;
}

.remove-row {
  margin-bottom: 1px;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

@media (max-width: 980px) {
  .detail-row {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }

  .remove-row {
    width: 100%;
  }
}

@media (max-width: 640px) {
  .component-header,
  .form-actions {
    flex-direction: column;
  }

  .component-header .btn,
  .form-actions .btn {
    width: 100%;
  }

  .blood-fields,
  .detail-row {
    grid-template-columns: 1fr;
  }
}
</style>
