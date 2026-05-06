<script setup lang="ts">
import { btn, ui } from '@/utils/ui'
</script>

<template>
  <main class="grid min-h-screen place-items-center p-6">
    <section :class="ui.emptyState">
      <div>
        <h1 class="mb-2 text-2xl font-semibold text-gray-900">Halaman tidak ditemukan</h1>
        <p :class="ui.emptyCopy">Alamat yang dibuka tidak tersedia di portal rumah sakit.</p>
        <RouterLink :class="[btn('btnPrimary'), 'mt-5']" to="/">Kembali ke Dashboard</RouterLink>
      </div>
    </section>
  </main>
</template>
