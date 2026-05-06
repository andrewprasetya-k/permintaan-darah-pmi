<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import AppModal from '@/components/feedback/AppModal.vue'
import { useFeedbackStore } from '@/stores/feedback'
import { useMyProfileStore } from '@/stores/my-profile'

const profileStore = useMyProfileStore()
const feedbackStore = useFeedbackStore()
const isConfirmOpen = ref(false)

const form = reactive({
  nama: '',
  nomorTelepon: '',
  alamat: '',
  email: '',
  password: '',
})

const fillForm = () => {
  if (!profileStore.profile) {
    return
  }

  form.nama = profileStore.profile.nama
  form.nomorTelepon = profileStore.profile.nomorTelepon
  form.alamat = profileStore.profile.alamat
  form.email = profileStore.profile.email || ''
  form.password = ''
}

const submit = async () => {
  isConfirmOpen.value = false

  try {
    await profileStore.updateMe({
      nama: form.nama.trim(),
      nomorTelepon: form.nomorTelepon.trim(),
      alamat: form.alamat.trim(),
      email: form.email.trim() || undefined,
      password: form.password || undefined,
    })
    fillForm()
    feedbackStore.showFlag({
      title: 'Profil disimpan',
      message: 'Data rumah sakit berhasil diperbarui.',
      variant: 'success',
    })
  } catch (err) {
    feedbackStore.showFlag({
      title: 'Gagal menyimpan profil',
      message: err instanceof Error ? err.message : 'Periksa kembali data profil.',
      variant: 'error',
    })
  }
}

onMounted(async () => {
  try {
    await profileStore.fetchMe()
    fillForm()
  } catch (err) {
    feedbackStore.showFlag({
      title: 'Gagal memuat profil',
      message: err instanceof Error ? err.message : 'Coba muat ulang halaman.',
      variant: 'error',
    })
  }
})
</script>

<template>
  <section>
    <div class="page-header">
      <div>
        <p class="page-eyebrow">Profil</p>
        <h1 class="page-title">Profil Rumah Sakit</h1>
        <p class="page-subtitle">Kelola data kontak yang tersimpan di backend rumah sakit.</p>
      </div>
    </div>

    <form class="card profile-form" @submit.prevent="isConfirmOpen = true">
      <div class="form-grid">
        <div class="form-field">
          <label class="form-label" for="nama">Nama rumah sakit</label>
          <input id="nama" v-model="form.nama" class="form-input" required />
        </div>
        <div class="form-field">
          <label class="form-label" for="nomorTelepon">Nomor telepon</label>
          <input id="nomorTelepon" v-model="form.nomorTelepon" class="form-input" required />
        </div>
        <div class="form-field">
          <label class="form-label" for="email">Email</label>
          <input id="email" v-model="form.email" class="form-input" type="email" />
        </div>
        <div class="form-field">
          <label class="form-label" for="password">Password baru</label>
          <input id="password" v-model="form.password" class="form-input" type="password" autocomplete="new-password" />
          <p class="form-help">Kosongkan bila tidak ingin mengganti password.</p>
        </div>
        <div class="form-field form-field-full">
          <label class="form-label" for="alamat">Alamat</label>
          <textarea id="alamat" v-model="form.alamat" class="form-textarea" required />
        </div>
      </div>

      <div class="profile-actions">
        <button type="submit" class="btn btn-primary" :disabled="profileStore.isLoading">
          {{ profileStore.isLoading ? 'Menyimpan...' : 'Simpan Profil' }}
        </button>
      </div>
    </form>

    <AppModal
      :is-open="isConfirmOpen"
      title="Simpan profil?"
      description="Perubahan data rumah sakit akan dikirim ke backend."
      width="sm"
      @close="isConfirmOpen = false"
    >
      <template #footer>
        <button type="button" class="btn btn-secondary" @click="isConfirmOpen = false">Batal</button>
        <button type="button" class="btn btn-primary" :disabled="profileStore.isLoading" @click="submit">
          Simpan
        </button>
      </template>
    </AppModal>
  </section>
</template>

<style scoped>
.profile-form {
  max-width: 860px;
  padding: 18px;
}

.profile-actions {
  display: flex;
  justify-content: flex-end;
  margin-top: 18px;
}

@media (max-width: 640px) {
  .profile-actions .btn {
    width: 100%;
  }
}
</style>
