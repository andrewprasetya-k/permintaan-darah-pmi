import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/utils/LoginView.vue'),
      meta: {
        requiresAuth: false,
        roles: ['admin', 'superadmin'],
        title: 'Login',
        subtitle: 'Masuk untuk mengelola permintaan darah',
      },
    },
    {
      path: '/',
      name: 'dashboard',
      component: () => import('../views/dashboard/DashboardView.vue'),
      meta: {
        requiresAuth: true,
        roles: ['admin', 'superadmin'],
        title: 'Dashboard',
        subtitle: 'Ringkasan aktivitas dan statistik terbaru',
      },
    },
    {
      path: '/permintaan',
      name: 'permintaan-list',
      component: () => import('../views/permintaan-darah/PermintaanDarah.vue'),
      meta: {
        requiresAuth: true,
        roles: ['admin', 'superadmin'],
        title: 'Permintaan Darah',
        subtitle: 'Daftar seluruh permintaan darah masuk',
      },
    },
    {
      path: '/permintaan/:id',
      name: 'permintaan-detail',
      component: () => import('../views/permintaan-darah/PermintaanDetailView.vue'),
      meta: { requiresAuth: true, roles: ['admin', 'superadmin'], title: 'Detail Permintaan' },
    },
    {
      path: '/admin',
      name: 'admin-list',
      component: () => import('../views/admin/AdminListView.vue'),
      meta: {
        requiresAuth: true,
        roles: ['admin', 'superadmin'],
        title: 'Kelola Admin',
        subtitle: 'Kelola akun admin yang terdaftar',
      },
    },
    {
      path: '/rumah-sakit',
      name: 'rumah-sakit-list',
      component: () => import('../views/rumah-sakit/RumahSakit.vue'),
      meta: {
        requiresAuth: true,
        roles: ['admin', 'superadmin'],
        title: 'Kelola Rumah Sakit',
        subtitle: 'Kelola rumah sakit yang terdaftar',
      },
    },
    {
      path: '/komponen',
      name: 'komponen-list',
      component: () => import('../views/komponen-darah/KomponenDarah.vue'),
      meta: {
        requiresAuth: true,
        roles: ['admin', 'superadmin'],
        title: 'Kelola Komponen Darah',
        subtitle: 'Kelola komponen darah yang tersedia',
      },
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: () => import('../views/utils/NotFoundView.vue'),
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
