import type { ApiResponse } from '@/types/models'

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'

interface RequestConfig {
  method: string
  headers: Record<string, string>
  body?: string
}

const getAuthToken = (): string | null => {
  return localStorage.getItem('authToken')
}

const buildHeaders = (includeAuth = true): Record<string, string> => {
  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
  }

  if (includeAuth) {
    const token = getAuthToken()
    if (token) {
      headers['Authorization'] = `Bearer ${token}`
    }
  }

  return headers
}

const handleResponse = async <T>(response: Response): Promise<ApiResponse<T>> => {
  if (!response.ok) {
    if (response.status === 401) {
      localStorage.removeItem('authToken')
      localStorage.removeItem('authUser')
      window.location.href = '/login'
    }

    const error = await response.json().catch(() => ({ message: 'Unknown error' }))
    throw new Error(error.message || `HTTP ${response.status}`)
  }

  return response.json()
}

class ApiClient {
  private baseUrl: string

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl
  }

  async request<T>(
    endpoint: string,
    options: {
      method?: string
      data?: any
      includeAuth?: boolean
      params?: Record<string, any>
    } = {},
  ): Promise<ApiResponse<T>> {
    const { method = 'GET', data, includeAuth = true, params = {} } = options

    let url = `${this.baseUrl}${endpoint}`

    // Add query parameters
    if (Object.keys(params).length > 0) {
      const queryString = new URLSearchParams(
        Object.entries(params).reduce(
          (acc, [key, value]) => {
            if (value === undefined || value === null || value === '') {
              return acc
            }
            acc[key] = String(value)
            return acc
          },
          {} as Record<string, string>,
        ),
      ).toString()
      url += `?${queryString}`
    }

    const config: RequestConfig = {
      method,
      headers: buildHeaders(includeAuth),
    }

    if (data && (method === 'POST' || method === 'PUT' || method === 'PATCH')) {
      config.body = JSON.stringify(data)
    }

    const response = await fetch(url, config)
    return handleResponse<T>(response)
  }

  get<T>(endpoint: string, options?: { params?: Record<string, any>; includeAuth?: boolean }) {
    return this.request<T>(endpoint, { ...options, method: 'GET' })
  }

  post<T>(
    endpoint: string,
    data?: any,
    options?: { params?: Record<string, any>; includeAuth?: boolean },
  ) {
    return this.request<T>(endpoint, { ...options, method: 'POST', data })
  }

  put<T>(
    endpoint: string,
    data?: any,
    options?: { params?: Record<string, any>; includeAuth?: boolean },
  ) {
    return this.request<T>(endpoint, { ...options, method: 'PUT', data })
  }

  patch<T>(
    endpoint: string,
    data?: any,
    options?: { params?: Record<string, any>; includeAuth?: boolean },
  ) {
    return this.request<T>(endpoint, { ...options, method: 'PATCH', data })
  }

  delete<T>(endpoint: string, options?: { params?: Record<string, any>; includeAuth?: boolean }) {
    return this.request<T>(endpoint, { ...options, method: 'DELETE' })
  }

  async download(endpoint: string, params: Record<string, any> = {}) {
    const url = new URL(`${this.baseUrl}${endpoint}`)
    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined && value !== null && value !== '') {
        url.searchParams.set(key, String(value))
      }
    })

    const token = getAuthToken()
    const headers: Record<string, string> = {}
    if (token) {
      headers.Authorization = `Bearer ${token}`
    }

    const response = await fetch(url.toString(), {
      method: 'GET',
      headers,
    })

    if (response.status === 401) {
      localStorage.removeItem('authToken')
      localStorage.removeItem('authUser')
      window.location.href = '/login'
    }

    if (!response.ok) {
      const error = await response.json().catch(() => ({ message: 'Unknown error' }))
      throw new Error(error.message || `HTTP ${response.status}`)
    }

    return {
      blob: await response.blob(),
      filename: parseFilename(response.headers.get('content-disposition')),
    }
  }
}

export const apiClient = new ApiClient(API_BASE_URL)

const parseFilename = (contentDisposition: string | null) => {
  if (!contentDisposition) {
    return undefined
  }

  const match = contentDisposition.match(/filename="?([^"]+)"?/i)
  return match?.[1]
}
