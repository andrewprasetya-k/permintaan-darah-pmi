export interface ApiResponse<T> {
  success: boolean
  data: T
  message?: string
  error?: string
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

export interface BloodRequest {
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
  detailPermintaanDarahId: number
  permintaanDarahId: string
  komponenId: number
  jumlahKantong: number
  createdAt: string
  updatedAt: string
}

export interface BloodInventory {
  id: string
  bloodType: string
  quantity: number
  expiredAt: string
}

export interface Donation {
  id: string
  donorName: string
  bloodType: string
  quantity: number
  location: string
  donationDate: string
  status: 'pending' | 'completed' | 'rejected'
  createdAt: string
  updatedAt: string
}

export interface StatusLog {
  id: string
  permintaanDarahId: string
  status: string
  keterangan: string
  createdAt: string
}

export interface SystemAccessLog {
  id: string
  userId: string
  action: string
  table: string
  targetId: string
  oldValue: string
  newValue: string
  timestamp: string
  ipAddress: string
}
