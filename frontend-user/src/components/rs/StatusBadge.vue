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
  min-height: 28px;
  align-items: center;
  border: 1px solid;
  border-radius: 999px;
  padding: 5px 10px;
  font-size: 12px;
  font-weight: 800;
  line-height: 1;
  white-space: nowrap;
}

.status-dibuat {
  border-color: #d1d5db;
  background: #f9fafb;
  color: #374151;
}

.status-diproses {
  border-color: #bfdbfe;
  background: #eff6ff;
  color: #1d4ed8;
}

.status-bisa_diambil {
  border-color: #bbf7d0;
  background: #f0fdf4;
  color: #15803d;
}

.status-selesai {
  border-color: #86efac;
  background: #dcfce7;
  color: #166534;
}

.status-dibatalkan {
  border-color: #fecaca;
  background: #fef2f2;
  color: #b91c1c;
}
</style>
