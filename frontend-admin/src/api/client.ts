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
}

export const apiClient = new ApiClient(API_BASE_URL)
