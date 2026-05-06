import { ref } from 'vue'

export const useNotification = () => {
  const hasPermission = ref(false)
  const isSupported = ref(typeof window !== 'undefined' && 'Notification' in window)
  let audioContext: AudioContext | null = null

  const requestPermission = async () => {
    if (!isSupported.value) {
      console.warn('Browser notifications not supported')
      return false
    }

    if (Notification.permission === 'granted') {
      hasPermission.value = true
      return true
    }

    if (Notification.permission !== 'denied') {
      const permission = await Notification.requestPermission()
      hasPermission.value = permission === 'granted'
      return permission === 'granted'
    }

    return false
  }

  const playSound = () => {
    try {
      // Create simple beep using Web Audio API
      if (!audioContext) {
        audioContext = new (window.AudioContext || (window as any).webkitAudioContext)()
      }

      const ctx = audioContext
      const now = ctx.currentTime

      // Create oscillator
      const osc = ctx.createOscillator()
      const gain = ctx.createGain()

      osc.connect(gain)
      gain.connect(ctx.destination)

      // Short beep: 800Hz for 300ms
      osc.frequency.value = 800
      gain.gain.setValueAtTime(0.3, now)
      gain.gain.exponentialRampToValueAtTime(0.01, now + 0.3)

      osc.start(now)
      osc.stop(now + 0.3)
    } catch (err) {
      console.warn('Could not play notification sound:', err)
    }
  }

  const showNotification = (title: string, options?: NotificationOptions) => {
    if (!hasPermission.value) {
      console.warn('Notification permission not granted')
      return
    }

    try {
      const notification = new Notification(title, {
        icon: '/icon.png',
        badge: '/badge.png',
        ...options,
      })

      notification.onclick = () => {
        window.focus()
        notification.close()
      }

      playSound()
    } catch (err) {
      console.warn('Could not show notification:', err)
    }
  }

  const notifyNewPermintaan = (data: {
    namaPasien: string
    golonganDarah: string
    rhesus: string
    rumahSakit?: string
  }) => {
    const title = 'Permintaan Darah Baru'
    const options: NotificationOptions = {
      body: `${data.namaPasien} (${data.golonganDarah}${data.rhesus})${data.rumahSakit ? ` - ${data.rumahSakit}` : ''}`,
      tag: 'permintaan-baru',
      requireInteraction: false,
    }

    showNotification(title, options)
  }

  const notifyPermintaanUpdate = (data: {
    namaPasien: string
    statusLama: string
    statusBaru: string
  }) => {
    const title = '🔄 Update Permintaan Darah'
    const options: NotificationOptions = {
      body: `${data.namaPasien}: ${data.statusLama} → ${data.statusBaru}`,
      tag: 'permintaan-update',
      requireInteraction: false,
    }

    showNotification(title, options)
  }

  return {
    hasPermission,
    isSupported,
    requestPermission,
    showNotification,
    notifyNewPermintaan,
    notifyPermintaanUpdate,
    playSound,
  }
}
