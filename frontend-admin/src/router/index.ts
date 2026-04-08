import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue'),
      meta: { requiresAuth: false, layout: 'blank' },
    },
    {
      path: '/',
      name: 'home',
      component: () => import('../views/HomeView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('../views/DashboardView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/permintaan',
      name: 'permintaan-list',
      component: () => import('../views/PermintaanListView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/permintaan/:id',
      name: 'permintaan-detail',
      component: () => import('../views/PermintaanDetailView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/admin',
      name: 'admin-list',
      component: () => import('../views/AdminListView.vue'),
      meta: { requiresAuth: true, roles: ['admin'] },
    },
    {
      path: '/rumah-sakit',
      name: 'rumah-sakit-list',
      component: () => import('../views/RumahSakitListView.vue'),
      meta: { requiresAuth: true, roles: ['admin'] },
    },
    {
      path: '/komponen',
      name: 'komponen-list',
      component: () => import('../views/KomponenListView.vue'),
      meta: { requiresAuth: true, roles: ['admin'] },
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: () => import('../views/NotFoundView.vue'),
    },
  ],
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  const requiresAuth = to.meta.requiresAuth !== false
  const requiredRoles = to.meta.roles as string[] | undefined
  const isAuthenticated = authStore.isAuthenticated
  const userRole = authStore.user?.role

  // Kalau page butuh auth tapi belum login
  if (requiresAuth && !isAuthenticated) {
    return next({ name: 'login', query: { redirect: to.fullPath } })
  }

  // Kalau sudah login dan ke login page, redirect ke home
  if (to.name === 'login' && isAuthenticated) {
    return next({ name: 'home' })
  }

  // Cek role-based access
  if (requiredRoles && isAuthenticated && userRole && !requiredRoles.includes(userRole)) {
    return next({ name: 'home' })
  }

  next()
})

export default router
