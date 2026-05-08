import type { BloodType, PermintaanStatus, Rhesus } from '@/types/models'

export const statusLabels: Record<PermintaanStatus, string> = {
  dibuat: 'Dibuat',
  diproses: 'Diproses',
  bisa_diambil: 'Bisa Diambil',
  selesai: 'Selesai',
  dibatalkan: 'Dibatalkan',
}

export const statusDescriptions: Record<PermintaanStatus, string> = {
  dibuat: 'Menunggu proses PMI',
  diproses: 'Sedang diproses PMI',
  bisa_diambil: 'Siap diambil rumah sakit',
  selesai: 'Selesai',
  dibatalkan: 'Dibatalkan',
}

export const formatDate = (value?: string) => {
  if (!value) {
    return '-'
  }

  const date = new Date(value)
  if (Number.isNaN(date.getTime())) {
    return '-'
  }

  return new Intl.DateTimeFormat('id-ID', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
  }).format(date)
}

export const formatDateTime = (value?: string) => {
  if (!value) {
    return '-'
  }

  const date = new Date(value)
  if (Number.isNaN(date.getTime())) {
    return '-'
  }

  return new Intl.DateTimeFormat('id-ID', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  }).format(date)
}

export const toDateInputValue = (value?: string) => {
  if (!value) {
    return ''
  }

  const date = new Date(value)
  if (Number.isNaN(date.getTime())) {
    return ''
  }

  return date.toISOString().slice(0, 10)
}

export const toTimeInputValue = (value?: string) => {
  if (!value) {
    return ''
  }

  const date = new Date(value)
  if (Number.isNaN(date.getTime())) {
    return ''
  }

  return date.toTimeString().slice(0, 5)
}

export const toIsoDate = (value: string) => {
  if (!value) {
    return ''
  }

  return new Date(`${value}T00:00:00`).toISOString()
}

export const toIsoDateTime = (dateValue: string, timeValue: string) => {
  if (!dateValue) {
    return ''
  }

  return new Date(`${dateValue}T${timeValue || '00:00'}:00`).toISOString()
}

export const bloodLabel = (blood?: BloodType, rhesus?: Rhesus) => {
  if (!blood && !rhesus) {
    return '-'
  }
  return `${blood ?? ''}${rhesus ?? ''}`
}
