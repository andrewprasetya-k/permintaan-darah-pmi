export interface ApiResponse<T> {
  success: boolean
  data: T
  message?: string
  error?: string
  pagination?: PaginationMeta
}

export interface PaginationMeta {
  total: number
  page: number
  limit: number
  offset: number
}

export interface PaginatedApiResponse<T> extends ApiResponse<T> {
  pagination?: PaginationMeta
}

export interface User {
  id: string
  username: string
  role: string
}

export interface Admin {
  adminId: string
  adminUsername: string
  adminName: string
  adminEmail: string
  adminRole: 'superadmin' | 'admin'
  isDeleted: boolean
  createdAt: string
  updatedAt: string
  deletedAt?: string
}

export interface RumahSakit {
  rumahSakitId: string
  nama: string
  nomorTelepon: string
  alamat: string
  email?: string
  username: string
  createdAt: string
  updatedAt: string
  deletedAt?: string
}

export interface KomponenDarah {
  komponenId: number
  komponenDarah: string
  komponenKode: string
  isActive: boolean
}

export interface PermintaanDarah {
  permintaanDarahId: string
  rumahSakitId: string
  namaPasien: string
  noRM?: string
  tempatLahir: string
  tanggalLahir: string
  jenisKelamin: 'L' | 'P'
  golonganDarah?: 'A' | 'B' | 'AB' | 'O'
  rhesusDarah?: '+' | '-'
  hemoglobin?: number
  ruangBagianKelas?: string
  pernahTransfusi: boolean
  indikasiTransfusi?: string
  pernahHamil?: string
  status: 'dibuat' | 'diproses' | 'bisa_diambil' | 'selesai' | 'dibatalkan'
  cancelReason?: string
  tanggalPermintaan: string
  createdAt: string
  updatedAt: string
  deletedAt?: string
  detailPermintaanDarah?: DetailPermintaanDarah[]
}

export interface DetailPermintaanDarah {
  detailId: number
  permintaanDarahId: string
  komponenNama: string
  golonganDarah: 'A' | 'B' | 'AB' | 'O'
  rhesusDarah: '+' | '-'
  jumlahKantong: number
  tanggalDiperlukan: string
  createdAt: string
}

export interface StatusLog {
  logId: string
  permintaanDarahId: string
  statusFrom?: string
  statusTo: string
  keterangan?: string
  createdAt: string
}

export interface SystemAccessLog {
  sysLogId: number
  sysUserId?: string
  sysUserNama: string
  sysUserRole: string
  sysAction: string
  sysTargetTable?: string
  sysTargetId?: string
  sysNotes: string
  sysUserAgent?: string
  createdAt: string
}

export interface WebSocketMessage<T = unknown> {
  type: string
  action: string
  entityId: string
  entityType: string
  data: T
  timestamp: string
}
