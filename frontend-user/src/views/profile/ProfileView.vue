<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import AppModal from '@/components/feedback/AppModal.vue'
import { useFeedbackStore } from '@/stores/feedback'
import { useMyProfileStore } from '@/stores/my-profile'
import { btn, ui } from '@/utils/ui'

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
    <div :class="ui.pageHeader">
      <div>
        <p :class="ui.pageEyebrow">Profil</p>
        <h1 :class="ui.pageTitle">Profil Rumah Sakit</h1>
        <p :class="ui.pageSubtitle">Kelola data kontak yang tersimpan di backend rumah sakit.</p>
      </div>
    </div>

    <form :class="[ui.card, 'max-w-[860px] p-6']" @submit.prevent="isConfirmOpen = true">
      <div :class="ui.formGrid">
        <div :class="ui.formField">
          <label :class="ui.formLabel" for="nama">Nama rumah sakit</label>
          <input id="nama" v-model="form.nama" :class="ui.formControl" required />
        </div>
        <div :class="ui.formField">
          <label :class="ui.formLabel" for="nomorTelepon">Nomor telepon</label>
          <input id="nomorTelepon" v-model="form.nomorTelepon" :class="ui.formControl" required />
        </div>
        <div :class="ui.formField">
          <label :class="ui.formLabel" for="email">Email</label>
          <input id="email" v-model="form.email" :class="ui.formControl" type="email" />
        </div>
        <div :class="ui.formField">
          <label :class="ui.formLabel" for="password">Password baru</label>
          <input
            id="password"
            v-model="form.password"
            :class="ui.formControl"
            type="password"
            autocomplete="new-password"
          />
          <p :class="ui.formHelp">Kosongkan bila tidak ingin mengganti password.</p>
        </div>
        <div :class="[ui.formField, ui.formFieldFull]">
          <label :class="ui.formLabel" for="alamat">Alamat</label>
          <textarea id="alamat" v-model="form.alamat" :class="ui.formTextarea" required />
        </div>
      </div>

      <div class="mt-5 flex justify-end max-sm:block">
        <button
          type="submit"
          :class="[btn('btnPrimary'), 'max-sm:w-full']"
          :disabled="profileStore.isLoading"
        >
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
          :disabled="profileStore.isLoading"
          @click="submit"
        >
          Simpan
        </button>
      </template>
    </AppModal>
  </section>
</template>
