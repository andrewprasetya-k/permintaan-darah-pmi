export interface ApiResponse<T> {
  success: boolean
  data: T
  message?: string
  error?: string
}

export interface User {
  userId: string
  username: string
  role: string
}

export interface BloodRequest {
  id: string
  patientName: string
  bloodType: string
  quantity: number
  hospital: string
  urgency: 'urgent' | 'normal'
  status: 'pending' | 'approved' | 'rejected' | 'completed'
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
