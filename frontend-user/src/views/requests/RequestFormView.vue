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
  PregnancyFlag,
  PermintaanStatus,
  Rhesus,
  UpdatePermintaanRequest,
} from '@/types/models'
import { toDateInputValue, toIsoDate, toIsoDateTime, toTimeInputValue } from '@/utils/format'
import { btn, ui } from '@/utils/ui'

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
  pernahHamil: PregnancyFlag | ''
  tanggalPermintaan: string
  jamPermintaan: string
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
  jamPermintaan: toTimeInputValue(new Date().toISOString()),
  status: 'dibuat',
})

const isEdit = computed(() => route.name === 'request-edit')
const requestId = computed(() => String(route.params.id || ''))
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

const optionalPregnancyFlag = (value: PregnancyFlag | '') => {
  return value || undefined
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
  pernahHamil: optionalPregnancyFlag(form.pernahHamil),
  status: form.status,
  tanggalPermintaan: toIsoDateTime(form.tanggalPermintaan, form.jamPermintaan),
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
  form.pernahHamil =
    request.pernahHamil === 'Y' || request.pernahHamil === 'N' ? request.pernahHamil : ''
  form.tanggalPermintaan = toDateInputValue(request.tanggalPermintaan)
  form.jamPermintaan = toTimeInputValue(request.tanggalPermintaan)
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
    <form class="grid gap-[18px]" @submit.prevent="openConfirm">
      <section :class="[ui.card, 'p-6']">
        <h2 :class="[ui.sectionTitle, 'mb-4']">Data Pasien</h2>
        <div :class="ui.formGrid">
          <div :class="ui.formField">
            <label :class="ui.formLabel" for="namaPasien">Nama pasien</label>
            <input id="namaPasien" v-model="form.namaPasien" :class="ui.formControl" required />
          </div>
          <div :class="ui.formField">
            <label :class="ui.formLabel" for="noRM">No. RM</label>
            <input id="noRM" v-model="form.noRM" :class="ui.formControl" />
          </div>
          <div :class="ui.formField">
            <label :class="ui.formLabel" for="tempatLahir">Tempat lahir</label>
            <input id="tempatLahir" v-model="form.tempatLahir" :class="ui.formControl" required />
          </div>
          <div :class="ui.formField">
            <label :class="ui.formLabel" for="tanggalLahir">Tanggal lahir</label>
            <input
              id="tanggalLahir"
              v-model="form.tanggalLahir"
              :class="ui.formControl"
              type="date"
              required
            />
          </div>
          <div :class="ui.formField">
            <label :class="ui.formLabel" for="jenisKelamin">Jenis kelamin</label>
            <select id="jenisKelamin" v-model="form.jenisKelamin" :class="ui.formControl" required>
              <option value="" disabled>Pilih jenis kelamin</option>
              <option value="L">Laki-laki</option>
              <option value="P">Perempuan</option>
            </select>
          </div>
          <div :class="ui.formField">
            <div class="grid grid-cols-[minmax(0,1fr)_120px] gap-2.5 max-sm:grid-cols-1">
              <div :class="ui.formField">
                <label :class="ui.formLabel" for="golonganDarah">Golongan darah</label>
                <select id="golonganDarah" v-model="form.golonganDarah" :class="ui.formControl">
                  <option value="">Belum diketahui</option>
                  <option value="A">A</option>
                  <option value="B">B</option>
                  <option value="AB">AB</option>
                  <option value="O">O</option>
                </select>
              </div>
              <div :class="ui.formField">
                <label :class="ui.formLabel" for="rhesusDarah">Rhesus</label>
                <select id="rhesusDarah" v-model="form.rhesusDarah" :class="ui.formControl">
                  <option value="">-</option>
                  <option value="+">+</option>
                  <option value="-">-</option>
                </select>
              </div>
            </div>
          </div>
        </div>
      </section>

      <section :class="[ui.card, 'p-6']">
        <h2 :class="[ui.sectionTitle, 'mb-4']">Data Medis</h2>
        <div :class="ui.formGrid">
          <div :class="ui.formField">
            <label :class="ui.formLabel" for="hemoglobin">Hemoglobin</label>
            <input
              id="hemoglobin"
              v-model="form.hemoglobin"
              :class="ui.formControl"
              type="number"
              step="0.1"
              min="0"
            />
          </div>
          <div :class="ui.formField">
            <label :class="ui.formLabel" for="ruangBagianKelas">Ruang/Bagian/Kelas</label>
            <input id="ruangBagianKelas" v-model="form.ruangBagianKelas" :class="ui.formControl" />
          </div>
          <div :class="ui.formField">
            <label :class="ui.formLabel" for="tanggalPermintaan">Tanggal permintaan</label>
            <input
              id="tanggalPermintaan"
              v-model="form.tanggalPermintaan"
              :class="ui.formControl"
              type="date"
              required
            />
          </div>
          <div :class="ui.formField">
            <label :class="ui.formLabel" for="jamPermintaan">Jam permintaan</label>
            <input
              id="jamPermintaan"
              v-model="form.jamPermintaan"
              :class="ui.formControl"
              type="time"
              required
            />
          </div>
          <div :class="[ui.formField, 'justify-end']">
            <label
              class="flex min-h-10 items-center gap-2.5 font-semibold text-gray-900"
              for="pernahTransfusi"
            >
              <input
                id="pernahTransfusi"
                v-model="form.pernahTransfusi"
                class="h-[18px] w-[18px] accent-blue-600"
                type="checkbox"
              />
              Pernah transfusi
            </label>
          </div>
          <div :class="[ui.formField, ui.formFieldFull]">
            <label :class="ui.formLabel" for="indikasiTransfusi">Indikasi transfusi</label>
            <textarea
              id="indikasiTransfusi"
              v-model="form.indikasiTransfusi"
              :class="ui.formTextarea"
            />
          </div>
          <div :class="[ui.formField, ui.formFieldFull]">
            <label :class="ui.formLabel" for="pernahHamil">Riwayat hamil</label>
            <input id="pernahHamil" v-model="form.pernahHamil" :class="ui.formControl" />
          </div>
        </div>
      </section>

      <section v-if="!isEdit" :class="[ui.card, 'p-6']">
        <div class="mb-3.5 flex items-start justify-between gap-4 max-sm:flex-col">
          <div>
            <h2 :class="ui.sectionTitle">Komponen Darah</h2>
            <p :class="ui.formHelp">Tambahkan kebutuhan komponen bila data sudah tersedia.</p>
          </div>
          <button
            type="button"
            :class="[btn('btnSecondary'), 'max-sm:w-full']"
            :disabled="komponenStore.isLoading || activeComponents.length === 0"
            @click="addDetailRow"
          >
            + Tambah Komponen
          </button>
        </div>

        <div
          v-if="komponenStore.error"
          class="mb-3 rounded-2xl border border-red-200 bg-red-50 p-3.5 text-sm text-red-800"
        >
          {{ komponenStore.error }}
        </div>

        <div
          v-if="detailRows.length === 0"
          class="rounded-2xl border border-dashed border-gray-200 p-3.5 text-sm text-gray-600"
        >
          Belum ada detail komponen pada permintaan ini.
        </div>

        <div v-else class="grid gap-3">
          <div
            v-for="row in detailRows"
            :key="row.key"
            class="grid grid-cols-[minmax(220px,1fr)_120px_100px_120px_44px] items-end gap-2.5 rounded-2xl border border-gray-100 p-3 max-xl:grid-cols-2 max-sm:grid-cols-1"
          >
            <div :class="ui.formField">
              <label :class="ui.formLabel" :for="`komponen-${row.key}`">Komponen</label>
              <select
                :id="`komponen-${row.key}`"
                v-model="row.komponenDarahId"
                :class="ui.formControl"
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
            <div :class="ui.formField">
              <label :class="ui.formLabel" :for="`detail-gol-${row.key}`">Golongan</label>
              <select
                :id="`detail-gol-${row.key}`"
                v-model="row.golonganDarah"
                :class="ui.formControl"
                required
              >
                <option value="A">A</option>
                <option value="B">B</option>
                <option value="AB">AB</option>
                <option value="O">O</option>
              </select>
            </div>
            <div :class="ui.formField">
              <label :class="ui.formLabel" :for="`detail-rhesus-${row.key}`">Rhesus</label>
              <select
                :id="`detail-rhesus-${row.key}`"
                v-model="row.rhesusDarah"
                :class="ui.formControl"
                required
              >
                <option value="+">+</option>
                <option value="-">-</option>
              </select>
            </div>
            <div :class="ui.formField">
              <label :class="ui.formLabel" :for="`jumlah-${row.key}`">Kantong</label>
              <input
                :id="`jumlah-${row.key}`"
                v-model="row.jumlahKantong"
                :class="ui.formControl"
                min="1"
                type="number"
                required
              />
            </div>
            <button
              type="button"
              :class="[btn('btnDanger'), ui.btnIcon, 'max-xl:w-full']"
              @click="removeDetailRow(row.key)"
            >
              x
            </button>
          </div>
        </div>
      </section>

      <div class="flex justify-end gap-2.5 max-sm:flex-col">
        <RouterLink
          :class="[btn('btnSecondary'), 'max-sm:w-full']"
          :to="isEdit ? `/requests/${requestId}` : '/requests'"
        >
          Batal
        </RouterLink>
        <button
          type="submit"
          :class="[btn('btnPrimary'), 'max-sm:w-full']"
          :disabled="requestsStore.isSubmitting"
        >
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
        <button
          type="button"
          :class="[btn('btnSecondary'), 'flex-1']"
          @click="isConfirmOpen = false"
        >
          Batal
        </button>
        <button
          type="button"
          :class="[btn('btnPrimary'), 'flex-1']"
          :disabled="requestsStore.isSubmitting"
          @click="submit"
        >
          {{ submitLabel }}
        </button>
      </template>
    </AppModal>
  </section>
</template>
