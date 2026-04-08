import { createApp } from 'vue'
import { createPinia } from 'pinia'
import './styles/global.css'

import App from './App.vue'
import router from './router'
import { useAuthStore } from './stores/auth'

const app = createApp(App)

const pinia = createPinia()
app.use(pinia)
app.use(router)

const authStore = useAuthStore()
authStore.initializeFromStorage()

app.mount('#app')
