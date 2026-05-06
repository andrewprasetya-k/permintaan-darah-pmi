import { createRouter, createWebHistory } from 'vue-router'
import AppLayout from '@/components/layout/AppLayout.vue'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/auth/LoginView.vue'),
      meta: { guestOnly: true },
    },
    {
      path: '/',
      component: AppLayout,
      meta: { requiresAuth: true },
      children: [
        {
          path: '',
          name: 'dashboard',
          component: () => import('@/views/dashboard/DashboardView.vue'),
          meta: {
            pageTitle: 'Dashboard',
            pageSubtitle: 'Ringkasan status permintaan darah rumah sakit',
          },
        },
        {
          path: 'requests',
          name: 'requests',
          component: () => import('@/views/requests/RequestsListView.vue'),
          meta: {
            pageTitle: 'Permintaan Saya',
            pageSubtitle: 'Daftar permintaan darah milik rumah sakit yang sedang login',
            pageActionKey: 'create-request',
          },
        },
        {
          path: 'requests/new',
          name: 'request-create',
          component: () => import('@/views/requests/RequestFormView.vue'),
          meta: {
            pageTitle: 'Buat Permintaan',
            pageSubtitle: 'Buat permintaan darah baru untuk rumah sakit',
          },
        },
        {
          path: 'requests/:id',
          name: 'request-detail',
          component: () => import('@/views/requests/RequestDetailView.vue'),
          meta: {
            pageTitle: 'Detail Permintaan',
            pageSubtitle: 'Lihat detail dan update status permintaan darah',
          },
        },
        {
          path: 'requests/:id/edit',
          name: 'request-edit',
          component: () => import('@/views/requests/RequestFormView.vue'),
          meta: {
            pageTitle: 'Edit Permintaan',
            pageSubtitle: 'Ubah data permintaan darah',
          },
        },
        {
          path: 'profile',
          name: 'profile',
          component: () => import('@/views/profile/ProfileView.vue'),
          meta: {
            pageTitle: 'Profil Rumah Sakit',
            pageSubtitle: 'Kelola informasi profil dan pengaturan rumah sakit',
          },
        },
      ],
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: () => import('@/views/NotFoundView.vue'),
    },
  ],
})

router.beforeEach((to) => {
  const authStore = useAuthStore()
  authStore.initializeFromStorage()

  if (to.meta.requiresAuth && (!authStore.isAuthenticated || !authStore.isRumahSakit)) {
    return { name: 'login', query: { redirect: to.fullPath } }
  }

  if (to.meta.guestOnly && authStore.isAuthenticated && authStore.isRumahSakit) {
    return { name: 'dashboard' }
  }

  return true
})

export default router
