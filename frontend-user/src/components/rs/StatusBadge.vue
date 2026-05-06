<script setup lang="ts">
import type { PermintaanStatus } from '@/types/models'
import { statusLabels } from '@/utils/format'

const props = defineProps<{
  status: PermintaanStatus
}>()
</script>

<template>
  <span
    class="inline-flex min-h-[26px] items-center whitespace-nowrap rounded-lg px-2.5 py-1 text-xs font-medium leading-none"
    :class="{
      'bg-amber-50 text-amber-700': props.status === 'dibuat',
      'bg-blue-50 text-blue-700': props.status === 'diproses',
      'bg-violet-50 text-violet-700': props.status === 'bisa_diambil',
      'bg-emerald-50 text-emerald-700': props.status === 'selesai',
      'bg-red-50 text-red-700': props.status === 'dibatalkan',
    }"
  >
    {{ statusLabels[props.status] }}
  </span>
</template>
