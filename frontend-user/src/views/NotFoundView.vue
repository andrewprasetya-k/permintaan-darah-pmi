<template>
  <main class="not-found">
    <section class="empty-state">
      <div>
        <h1>Halaman tidak ditemukan</h1>
        <p>Alamat yang dibuka tidak tersedia di portal rumah sakit.</p>
        <RouterLink class="btn btn-primary" to="/">Kembali ke Dashboard</RouterLink>
      </div>
    </section>
  </main>
</template>

<style scoped>
.not-found {
  display: grid;
  min-height: 100vh;
  place-items: center;
  padding: 24px;
}

.not-found h1 {
  margin: 0 0 8px;
  font-size: 24px;
  font-weight: 900;
}

.not-found .btn {
  margin-top: 18px;
}
</style>
