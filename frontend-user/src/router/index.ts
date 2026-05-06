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
        },
        {
          path: 'requests',
          name: 'requests',
          component: () => import('@/views/requests/RequestsListView.vue'),
        },
        {
          path: 'requests/new',
          name: 'request-create',
          component: () => import('@/views/requests/RequestFormView.vue'),
        },
        {
          path: 'requests/:id',
          name: 'request-detail',
          component: () => import('@/views/requests/RequestDetailView.vue'),
        },
        {
          path: 'requests/:id/edit',
          name: 'request-edit',
          component: () => import('@/views/requests/RequestFormView.vue'),
        },
        {
          path: 'profile',
          name: 'profile',
          component: () => import('@/views/profile/ProfileView.vue'),
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
