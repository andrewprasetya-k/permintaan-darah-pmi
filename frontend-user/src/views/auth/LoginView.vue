<script setup lang="ts">
import { reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useFeedbackStore } from '@/stores/feedback'
import { useMyProfileStore } from '@/stores/my-profile'
import { btn, ui } from '@/utils/ui'

const authStore = useAuthStore()
const profileStore = useMyProfileStore()
const feedbackStore = useFeedbackStore()
const router = useRouter()
const route = useRoute()

const form = reactive({
  username: '',
  password: '',
})

const submit = async () => {
  try {
    await authStore.login(form.username.trim(), form.password)
    await profileStore.fetchMe().catch(() => undefined)
    feedbackStore.showFlag({
      title: 'Login berhasil',
      message: 'Portal rumah sakit siap digunakan.',
      variant: 'success',
    })
    const redirect = typeof route.query.redirect === 'string' ? route.query.redirect : '/'
    await router.push(redirect)
  } catch (err) {
    feedbackStore.showFlag({
      title: 'Login gagal',
      message: err instanceof Error ? err.message : 'Periksa username dan password.',
      variant: 'error',
    })
  }
}
</script>

<template>
  <main class="grid min-h-screen place-items-center bg-gray-50 p-6">
    <section class="w-[min(430px,100%)] rounded-2xl border border-gray-100 bg-white p-7 shadow-2xl">
      <div class="mb-7 flex items-center gap-3.5">
        <div
          class="grid h-14 w-14 place-items-center rounded-2xl bg-blue-50 font-bold text-blue-600"
        >
          PMI
        </div>
        <div>
          <p class="m-0 text-[13px] font-medium text-gray-500">Permintaan Darah</p>
          <h1 class="m-0 text-2xl font-semibold text-gray-900">Portal Rumah Sakit</h1>
        </div>
      </div>

      <form class="grid gap-4" @submit.prevent="submit">
        <div :class="ui.formField">
          <label :class="ui.formLabel" for="username">Username</label>
          <input
            id="username"
            v-model="form.username"
            :class="ui.formControl"
            type="text"
            autocomplete="username"
            required
            autofocus
          />
        </div>

        <div :class="ui.formField">
          <label :class="ui.formLabel" for="password">Password</label>
          <input
            id="password"
            v-model="form.password"
            :class="ui.formControl"
            type="password"
            autocomplete="current-password"
            required
          />
        </div>

        <button
          type="submit"
          :class="[btn('btnPrimary'), 'mt-1.5 w-full']"
          :disabled="authStore.isLoading"
        >
          {{ authStore.isLoading ? 'Masuk...' : 'Masuk' }}
        </button>
      </form>
    </section>
  </main>
</template>
