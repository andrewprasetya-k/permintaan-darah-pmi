<script setup lang="ts">
import { reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useFeedbackStore } from '@/stores/feedback'
import { useMyProfileStore } from '@/stores/my-profile'

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
  <main class="login-page">
    <section class="login-panel">
      <div class="login-brand">
        <div class="login-mark">PMI</div>
        <div>
          <p>Permintaan Darah</p>
          <h1>Portal Rumah Sakit</h1>
        </div>
      </div>

      <form class="login-form" @submit.prevent="submit">
        <div class="form-field">
          <label class="form-label" for="username">Username</label>
          <input
            id="username"
            v-model="form.username"
            class="form-input"
            type="text"
            autocomplete="username"
            required
            autofocus
          />
        </div>

        <div class="form-field">
          <label class="form-label" for="password">Password</label>
          <input
            id="password"
            v-model="form.password"
            class="form-input"
            type="password"
            autocomplete="current-password"
            required
          />
        </div>

        <button type="submit" class="btn btn-primary login-submit" :disabled="authStore.isLoading">
          {{ authStore.isLoading ? 'Masuk...' : 'Masuk' }}
        </button>
      </form>
    </section>
  </main>
</template>

<style scoped>
.login-page {
  display: grid;
  min-height: 100vh;
  place-items: center;
  padding: 24px;
  background:
    linear-gradient(135deg, rgba(220, 38, 38, 0.08), transparent 42%),
    var(--bg);
}

.login-panel {
  width: min(430px, 100%);
  border: 1px solid var(--line);
  border-radius: 8px;
  background: var(--surface);
  padding: 28px;
  box-shadow: var(--shadow);
}

.login-brand {
  display: flex;
  align-items: center;
  gap: 14px;
  margin-bottom: 28px;
}

.login-mark {
  display: grid;
  width: 54px;
  height: 54px;
  place-items: center;
  border-radius: 8px;
  background: var(--red);
  color: #ffffff;
  font-weight: 900;
}

.login-brand p,
.login-brand h1 {
  margin: 0;
}

.login-brand p {
  color: var(--text-muted);
  font-size: 13px;
  font-weight: 800;
}

.login-brand h1 {
  font-size: 24px;
  font-weight: 900;
}

.login-form {
  display: grid;
  gap: 16px;
}

.login-submit {
  margin-top: 6px;
  width: 100%;
}
</style>
