<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { AlertCircle, Shield, UserRound, Mail, AtSign, Eye, EyeOff } from '@lucide/vue'
import { useAdminStore } from '@/stores/admin'
import { useAuthStore } from '@/stores/auth'
import AppModal from '@/components/feedback/AppModal.vue'
import AppFlag from '@/components/feedback/AppFlag.vue'

const adminStore = useAdminStore()
const authStore = useAuthStore()
const isSubmitting = ref(false)
const showPassword = ref(false)
const showSubmitDialog = ref(false)
const flag = ref<{ variant: 'success' | 'error'; title: string; message?: string } | null>(null)
const formData = ref({
  adminUsername: '',
  adminName: '',
  adminEmail: '',
  adminRole: 'admin' as 'admin' | 'superadmin',
  adminPassword: '',
})

const profile = computed(() => adminStore.myProfile)

const hydrateForm = () => {
  if (!adminStore.myProfile) return
  formData.value = {
    adminUsername: adminStore.myProfile.adminUsername,
    adminName: adminStore.myProfile.adminName,
    adminEmail: adminStore.myProfile.adminEmail,
    adminRole: adminStore.myProfile.adminRole,
    adminPassword: '',
  }
}

onMounted(async () => {
  await adminStore.fetchMe()
  hydrateForm()
})

const handleSubmit = async () => {
  isSubmitting.value = true
  try {
    const updatedProfile = await adminStore.updateMe(formData.value)
    if (authStore.user) {
      authStore.user = {
        ...authStore.user,
        username: updatedProfile.adminUsername,
      }
      localStorage.setItem('authUser', JSON.stringify(authStore.user))
    }
    hydrateForm()
    flag.value = {
      variant: 'success',
      title: 'Profil diperbarui',
      message: 'Perubahan profil berhasil disimpan.',
    }
  } catch (error) {
    flag.value = {
      variant: 'error',
      title: 'Operasi gagal',
      message: error instanceof Error ? error.message : 'Gagal memperbarui profil.',
    }
  } finally {
    isSubmitting.value = false
    showSubmitDialog.value = false
  }
}

const openSubmitDialog = () => {
  showSubmitDialog.value = true
}

const closeSubmitDialog = () => {
  if (isSubmitting.value) return
  showSubmitDialog.value = false
}
</script>

<template>
  <div class="mx-auto flex h-full min-h-0 max-w-4xl flex-col">
    <AppFlag
      v-if="flag"
      :variant="flag.variant"
      :title="flag.title"
      :message="flag.message"
      @close="flag = null"
    />

    <AppModal
      :is-open="showSubmitDialog"
      title="Simpan Perubahan Profil"
      description="Perubahan data profil akun Anda akan disimpan."
      @close="closeSubmitDialog"
    >
      <template #footer>
        <button
          type="button"
          class="flex-1 rounded-xl bg-gray-100 px-4 py-2.5 text-sm font-medium text-gray-700 transition-colors hover:bg-gray-200"
          @click="closeSubmitDialog"
        >
          Batal
        </button>
        <button
          type="button"
          :disabled="isSubmitting || adminStore.isLoading"
          class="flex-1 rounded-xl bg-blue-600 px-4 py-2.5 text-sm font-medium text-white transition-colors hover:bg-blue-700 disabled:opacity-50"
          @click="handleSubmit"
        >
          {{ isSubmitting ? 'Menyimpan...' : 'Simpan' }}
        </button>
      </template>
    </AppModal>

    <div class="grid min-h-0 flex-1 gap-6 lg:grid-cols-[320px_minmax(0,1fr)]">
      <section class="rounded-2xl border border-gray-100 bg-white p-6">
        <div class="flex items-center gap-4">
          <div
            class="flex h-14 w-14 items-center justify-center rounded-2xl bg-blue-50 text-blue-600"
          >
            <UserRound :size="24" />
          </div>
          <div>
            <h2 class="text-lg font-semibold text-gray-900">
              {{ profile?.adminName || 'Admin' }}
            </h2>
            <p class="mt-1 text-sm text-gray-500">{{ profile?.adminUsername || '-' }}</p>
          </div>
        </div>

        <div class="mt-6 space-y-4 border-t border-gray-100 pt-6">
          <div class="rounded-xl bg-gray-50 px-4 py-3">
            <p class="text-xs font-medium uppercase tracking-wide text-gray-400">Role</p>
            <div
              class="mt-2 inline-flex items-center gap-2 rounded-lg bg-white px-3 py-1.5 text-sm font-medium text-gray-700"
            >
              <Shield :size="14" class="text-blue-600" />
              {{ profile?.adminRole === 'superadmin' ? 'Superadmin' : 'Admin' }}
            </div>
          </div>
          <div class="rounded-xl bg-gray-50 px-4 py-3">
            <p class="text-xs font-medium uppercase tracking-wide text-gray-400">Dibuat</p>
            <p class="mt-2 text-sm text-gray-700">
              {{ profile?.createdAt ? new Date(profile.createdAt).toLocaleString('id-ID') : '-' }}
            </p>
          </div>
          <div class="rounded-xl bg-gray-50 px-4 py-3">
            <p class="text-xs font-medium uppercase tracking-wide text-gray-400">Diperbarui</p>
            <p class="mt-2 text-sm text-gray-700">
              {{ profile?.updatedAt ? new Date(profile.updatedAt).toLocaleString('id-ID') : '-' }}
            </p>
          </div>
        </div>
      </section>

      <section class="rounded-2xl border border-gray-100 bg-white">
        <div class="border-b border-gray-100 px-6 py-5">
          <h2 class="text-lg font-semibold text-gray-900">Perbarui Profil</h2>
          <p class="mt-1 text-sm text-gray-500">Akun ini memakai endpoint profil admin sendiri.</p>
        </div>

        <form class="space-y-5 px-6 py-6" @submit.prevent="openSubmitDialog">
          <div>
            <label class="mb-1.5 block text-xs font-medium uppercase tracking-wide text-gray-500">
              Username
            </label>
            <div class="relative">
              <AtSign
                :size="16"
                class="pointer-events-none absolute left-3 top-1/2 -translate-y-1/2 text-gray-400"
              />
              <input
                v-model="formData.adminUsername"
                type="text"
                required
                class="w-full rounded-xl border border-gray-200 bg-gray-50 py-2.5 pl-9 pr-3 text-sm text-gray-900 outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>
          </div>

          <div>
            <label class="mb-1.5 block text-xs font-medium uppercase tracking-wide text-gray-500">
              Nama
            </label>
            <div class="relative">
              <UserRound
                :size="16"
                class="pointer-events-none absolute left-3 top-1/2 -translate-y-1/2 text-gray-400"
              />
              <input
                v-model="formData.adminName"
                type="text"
                required
                class="w-full rounded-xl border border-gray-200 bg-gray-50 py-2.5 pl-9 pr-3 text-sm text-gray-900 outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>
          </div>

          <div>
            <label class="mb-1.5 block text-xs font-medium uppercase tracking-wide text-gray-500">
              Email
            </label>
            <div class="relative">
              <Mail
                :size="16"
                class="pointer-events-none absolute left-3 top-1/2 -translate-y-1/2 text-gray-400"
              />
              <input
                v-model="formData.adminEmail"
                type="email"
                required
                class="w-full rounded-xl border border-gray-200 bg-gray-50 py-2.5 pl-9 pr-3 text-sm text-gray-900 outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
              />
            </div>
          </div>

          <div>
            <label class="mb-1.5 block text-xs font-medium uppercase tracking-wide text-gray-500">
              Password
            </label>
            <div class="relative">
              <input
                v-model="formData.adminPassword"
                :type="showPassword ? 'text' : 'password'"
                class="w-full rounded-xl border border-gray-200 bg-gray-50 px-3 py-2.5 pr-10 text-sm text-gray-900 outline-none transition-all focus:bg-white focus:border-blue-400 focus:ring-2 focus:ring-blue-100"
                placeholder="Kosongkan jika tidak diubah"
              />
              <button
                type="button"
                class="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 transition-colors hover:text-gray-600"
                @click="showPassword = !showPassword"
              >
                <Eye v-if="!showPassword" :size="16" />
                <EyeOff v-else :size="16" />
              </button>
            </div>
          </div>

          <div
            v-if="adminStore.error"
            class="flex items-center gap-2 rounded-xl border border-red-100 bg-red-50 p-3 text-xs text-red-600"
          >
            <AlertCircle :size="14" class="shrink-0" />
            {{ adminStore.error }}
          </div>

          <div class="flex justify-end">
            <button
              type="submit"
              :disabled="isSubmitting || adminStore.isLoading"
              class="inline-flex min-w-40 items-center justify-center rounded-xl bg-blue-600 px-4 py-2.5 text-sm font-medium text-white transition-colors hover:bg-blue-700 disabled:opacity-50"
            >
              {{ isSubmitting ? 'Menyimpan...' : 'Simpan Perubahan' }}
            </button>
          </div>
        </form>
      </section>
    </div>
  </div>
</template>
