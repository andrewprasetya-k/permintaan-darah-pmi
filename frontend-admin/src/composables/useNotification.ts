import { ref } from 'vue'

const hasPermission = ref(
  typeof window !== 'undefined' && 'Notification' in window && Notification.permission === 'granted',
)
const isSupported = ref(typeof window !== 'undefined' && 'Notification' in window)
let audioContext: AudioContext | null = null
let isAudioUnlockBound = false

const getAudioContext = () => {
  if (typeof window === 'undefined') {
    return null
  }

  if (!audioContext) {
    audioContext = new (window.AudioContext || (window as any).webkitAudioContext)()
  }

  return audioContext
}

const unlockAudio = () => {
  const ctx = getAudioContext()
  if (ctx?.state === 'suspended') {
    void ctx.resume()
  }
}

const bindAudioUnlock = () => {
  if (isAudioUnlockBound || typeof window === 'undefined') {
    return
  }

  isAudioUnlockBound = true
  window.addEventListener('pointerdown', unlockAudio, { once: true })
  window.addEventListener('keydown', unlockAudio, { once: true })
}

export const useNotification = () => {
  bindAudioUnlock()

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

    hasPermission.value = false
    return false
  }

  const playSound = () => {
    try {
      const ctx = getAudioContext()
      if (!ctx) {
        return
      }

      const beep = () => {
        const now = ctx.currentTime
        const osc = ctx.createOscillator()
        const gain = ctx.createGain()

        osc.connect(gain)
        gain.connect(ctx.destination)

        osc.frequency.value = 800
        gain.gain.setValueAtTime(0.3, now)
        gain.gain.exponentialRampToValueAtTime(0.01, now + 0.3)

        osc.start(now)
        osc.stop(now + 0.3)
      }

      if (ctx.state === 'suspended') {
        void ctx.resume().then(beep).catch((err) => {
          console.warn('Could not resume notification sound:', err)
        })
        return
      }

      beep()
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
    const body = `${data.namaPasien} (${data.golonganDarah}${data.rhesus})${data.rumahSakit ? ` - ${data.rumahSakit}` : ''}`

    const options: NotificationOptions = {
      body,
      tag: 'permintaan-baru',
      requireInteraction: false,
    }

    playSound()
    showNotification(title, options)
  }

  const notifyPermintaanUpdate = (data: {
    namaPasien: string
    statusLama?: string
    statusBaru?: string
  }) => {
    const title = 'Update Permintaan Darah'
    const statusText =
      data.statusLama && data.statusBaru
        ? `${data.statusLama} -> ${data.statusBaru}`
        : data.statusBaru
          ? `Status: ${data.statusBaru}`
          : 'Data permintaan diperbarui'
    const body = `${data.namaPasien}: ${statusText}`

    const options: NotificationOptions = {
      body,
      tag: 'permintaan-update',
      requireInteraction: false,
    }

    playSound()
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
