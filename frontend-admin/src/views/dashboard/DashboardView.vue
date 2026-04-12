<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
</script>

<template>
  <div style="max-width: 1200px; margin: 0 auto">
    <!-- Stats Grid -->
    <div
      style="
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
        gap: 1rem;
        margin-bottom: 2rem;
      "
    >
      <div
        style="
          background-color: white;
          border-radius: 0.5rem;
          padding: 1.5rem;
          box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
        "
      >
        <p style="font-size: 0.875rem; color: #6b7280; margin-bottom: 0.5rem">Permintaan Baru</p>
        <p style="font-size: 2.5rem; font-weight: 700; color: #2563eb">0</p>
      </div>

      <div
        style="
          background-color: white;
          border-radius: 0.5rem;
          padding: 1.5rem;
          box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
        "
      >
        <p style="font-size: 0.875rem; color: #6b7280; margin-bottom: 0.5rem">Dalam Proses</p>
        <p style="font-size: 2.5rem; font-weight: 700; color: #f59e0b">0</p>
      </div>

      <div
        style="
          background-color: white;
          border-radius: 0.5rem;
          padding: 1.5rem;
          box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
        "
      >
        <p style="font-size: 0.875rem; color: #6b7280; margin-bottom: 0.5rem">Siap Diambil</p>
        <p style="font-size: 2.5rem; font-weight: 700; color: #10b981">0</p>
      </div>

      <div
        style="
          background-color: white;
          border-radius: 0.5rem;
          padding: 1.5rem;
          box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
        "
      >
        <p style="font-size: 0.875rem; color: #6b7280; margin-bottom: 0.5rem">Selesai</p>
        <p style="font-size: 2.5rem; font-weight: 700; color: #8b5cf6">0</p>
      </div>
    </div>

    <!-- Recent Activity -->
    <div
      style="
        background-color: white;
        border-radius: 0.5rem;
        padding: 1.5rem;
        box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
      "
    >
      <h2 style="font-size: 1.25rem; font-weight: 600; color: #111827; margin-bottom: 1rem">
        Aktivitas Terbaru
      </h2>
      <p style="color: #6b7280; text-align: center; padding: 2rem">Belum ada aktivitas</p>
    </div>
  </div>
</template>
