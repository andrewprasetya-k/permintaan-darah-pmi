<script setup lang="ts">
import type { PermintaanStatus } from '@/types/models'
import { statusLabels } from '@/utils/format'

const props = defineProps<{
  status: PermintaanStatus
}>()
</script>

<template>
  <span class="status-badge" :class="`status-${props.status}`">
    {{ statusLabels[props.status] }}
  </span>
</template>

<style scoped>
.status-badge {
  display: inline-flex;
  min-height: 26px;
  align-items: center;
  border-radius: 8px;
  padding: 5px 10px;
  font-size: 12px;
  font-weight: 500;
  line-height: 1;
  white-space: nowrap;
}

.status-dibuat {
  background: #fffbeb;
  color: #b45309;
}

.status-diproses {
  background: #eff6ff;
  color: #1d4ed8;
}

.status-bisa_diambil {
  background: #f5f3ff;
  color: #6d28d9;
}

.status-selesai {
  background: #ecfdf5;
  color: #166534;
}

.status-dibatalkan {
  background: #fef2f2;
  color: #b91c1c;
}
</style>
