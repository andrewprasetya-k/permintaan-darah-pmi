export interface ApiResponse<T> {
  success: boolean
  message: string
  data: T
  pagination?: PaginationMeta
  errors?: string[]
}

export interface PaginationMeta {
  total: number
  page: number
  limit: number
  offset: number
}

export interface User {
  id: string
  username: string
  role: 'rumah_sakit' | string
}

export interface LoginResponse extends User {
  token: string
}

export interface RumahSakit {
  rumahSakitId: string
  nama: string
  nomorTelepon: string
  alamat: string
  email?: string
  username: string
  isDeleted: boolean
  createdAt: string
  updatedAt: string
  deletedAt?: string
}

export interface UpdateRumahSakitProfileRequest {
  nama: string
  nomorTelepon: string
  alamat: string
  email?: string
  password?: string
}

export type PermintaanStatus = 'dibuat' | 'diproses' | 'bisa_diambil' | 'selesai' | 'dibatalkan'
export type Gender = 'L' | 'P'
export type BloodType = 'A' | 'B' | 'AB' | 'O'
export type Rhesus = '+' | '-'
export type PregnancyFlag = 'Y' | 'N'

export interface DetailPermintaanDarah {
  detailId: number
  permintaanDarahId: string
  komponenNama: string
  golonganDarah: BloodType
  rhesusDarah: Rhesus
  jumlahKantong: number
  tanggalDiperlukan: string
  createdAt: string
}

export interface PermintaanDarahListItem {
  permintaanDarahId: string
  namaPasien: string
  jenisKelamin: Gender
  golonganDarah?: BloodType
  rhesusDarah?: Rhesus
  tanggalPermintaan: string
  status: PermintaanStatus
  createdAt: string
  updatedAt: string
  deletedAt?: string
}

export interface PermintaanDarah extends PermintaanDarahListItem {
  rumahSakitId: string
  noRM?: string
  tempatLahir: string
  tanggalLahir: string
  hemoglobin?: number
  ruangBagianKelas?: string
  pernahTransfusi: boolean
  indikasiTransfusi?: string
  pernahHamil?: PregnancyFlag
  cancelReason?: string
  detailPermintaanDarah?: DetailPermintaanDarah[]
}

export interface CreateDetailPermintaanDarahRequest {
  komponenDarahId: number
  golonganDarah: BloodType
  rhesusDarah: Rhesus
  jumlahKantong: number
}

export interface CreatePermintaanRequest {
  rumahSakitId?: string
  namaPasien: string
  noRM?: string
  tempatLahir: string
  tanggalLahir: string
  jenisKelamin: Gender
  golonganDarah?: BloodType
  rhesusDarah?: Rhesus
  hemoglobin?: number
  ruangBagianKelas?: string
  pernahTransfusi: boolean
  indikasiTransfusi?: string
  pernahHamil?: PregnancyFlag
  status: PermintaanStatus
  cancelReason?: string
  tanggalPermintaan: string
  details?: CreateDetailPermintaanDarahRequest[]
}

export interface UpdatePermintaanRequest extends Omit<
  CreatePermintaanRequest,
  'rumahSakitId' | 'details'
> {
  rumahSakitId: string
}

export interface UpdatePermintaanStatusRequest {
  status: PermintaanStatus
  reason?: string
}

export interface KomponenDarah {
  komponenId: number
  komponenDarah: string
  komponenKode: string
  isActive?: boolean
}

export interface StatusSummary {
  dibuat: number
  diproses: number
  bisaDiambil: number
  selesai: number
  dibatalkan: number
  total: number
}

export interface WebSocketMessage<T = unknown> {
  type: string
  action: string
  entityId: string
  entityType: string
  data: T
  timestamp: string
}
