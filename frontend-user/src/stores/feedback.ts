import { defineStore } from 'pinia'
import { ref } from 'vue'

type FlagVariant = 'success' | 'error' | 'info'

interface FlagState {
  title: string
  message?: string
  variant: FlagVariant
}

export const useFeedbackStore = defineStore('feedback', () => {
  const flag = ref<FlagState | null>(null)
  let timer: number | null = null

  const clearTimer = () => {
    if (timer !== null) {
      window.clearTimeout(timer)
      timer = null
    }
  }

  const showFlag = (nextFlag: FlagState, duration = 3500) => {
    clearTimer()
    flag.value = nextFlag
    timer = window.setTimeout(() => {
      flag.value = null
      timer = null
    }, duration)
  }

  const clearFlag = () => {
    clearTimer()
    flag.value = null
  }

  return {
    flag,
    showFlag,
    clearFlag,
  }
})
