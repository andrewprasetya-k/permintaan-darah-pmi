<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const logout = () => {
  authStore.logout()
  router.push('/login')
}
</script>

<template>
  <div style="min-height: 100vh; background-color: #f3f4f6;">
    <nav style="background-color: white; box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);">
      <div style="max-width: 80rem; margin: 0 auto; padding: 1rem 1.5rem; display: flex; justify-content: space-between; align-items: center;">
        <h1 style="font-size: 1.5rem; font-weight: bold;">Admin Dashboard</h1>
        <div style="display: flex; align-items: center; gap: 1rem;">
          <span style="color: #374151;">{{ authStore.user?.username || 'Loading...' }}</span>
          <button
            @click="logout"
            style="padding: 0.5rem 1rem; background-color: #dc2626; color: white; border-radius: 0.375rem; border: none; cursor: pointer; font-weight: 500;"
            @mouseenter="($event.target as HTMLElement).style.backgroundColor = '#b91c1c'"
            @mouseleave="($event.target as HTMLElement).style.backgroundColor = '#dc2626'"
          >
            Logout
          </button>
        </div>
      </div>
    </nav>

    <div style="max-width: 80rem; margin: 0 auto; padding: 1.5rem;">
      <div style="background-color: white; padding: 1.5rem; border-radius: 0.5rem; box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);">
        <h2 style="font-size: 1.5rem; font-weight: bold; margin-bottom: 1rem;">Selamat Datang!</h2>
        <p style="color: #4b5563;">
          Halo {{ authStore.user?.username || '(username)' }}, selamat datang di Admin Dashboard Permintaan Darah PMI
        </p>
        <p style="font-size: 0.875rem; color: #6b7280; margin-top: 0.5rem;">Role: {{ authStore.user?.role || '(role)' }}</p>

        <div style="margin-top: 2rem; display: grid; grid-template-columns: repeat(2, 1fr); gap: 1rem;">
          <router-link
            to="/permintaan"
            style="padding: 1rem; background-color: #dbeafe; border-radius: 0.375rem; text-decoration: none; color: inherit; transition: background-color 0.2s;"
            @mouseenter="($event.target as HTMLElement).style.backgroundColor = '#bfdbfe'"
            @mouseleave="($event.target as HTMLElement).style.backgroundColor = '#dbeafe'"
          >
            <h3 style="font-weight: 600; color: #1e3a8a;">Permintaan Darah</h3>
          </router-link>

          <router-link
            to="/dashboard"
            style="padding: 1rem; background-color: #dcfce7; border-radius: 0.375rem; text-decoration: none; color: inherit; transition: background-color 0.2s;"
            @mouseenter="($event.target as HTMLElement).style.backgroundColor = '#bbf7d0'"
            @mouseleave="($event.target as HTMLElement).style.backgroundColor = '#dcfce7'"
          >
            <h3 style="font-weight: 600; color: #166534;">Dashboard</h3>
          </router-link>

          <router-link
            v-if="authStore.user?.role === 'superadmin' || authStore.user?.role === 'admin'"
            to="/admin"
            style="padding: 1rem; background-color: #e9d5ff; border-radius: 0.375rem; text-decoration: none; color: inherit; transition: background-color 0.2s;"
            @mouseenter="($event.target as HTMLElement).style.backgroundColor = '#d8b4fe'"
            @mouseleave="($event.target as HTMLElement).style.backgroundColor = '#e9d5ff'"
          >
            <h3 style="font-weight: 600; color: #581c87;">Manajemen Admin</h3>
          </router-link>

          <router-link
            v-if="authStore.user?.role === 'superadmin' || authStore.user?.role === 'admin'"
            to="/rumah-sakit"
            style="padding: 1rem; background-color: #fed7aa; border-radius: 0.375rem; text-decoration: none; color: inherit; transition: background-color 0.2s;"
            @mouseenter="($event.target as HTMLElement).style.backgroundColor = '#fdba74'"
            @mouseleave="($event.target as HTMLElement).style.backgroundColor = '#fed7aa'"
          >
            <h3 style="font-weight: 600; color: #92400e;">Manajemen Rumah Sakit</h3>
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>
