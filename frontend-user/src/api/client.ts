import type { ApiResponse } from '@/types/models'

export const AUTH_TOKEN_KEY = 'rsAuthToken'
export const AUTH_USER_KEY = 'rsAuthUser'

const trimTrailingSlash = (value: string) => value.replace(/\/$/, '')

export const API_BASE_URL = trimTrailingSlash(
  import.meta.env.VITE_API_URL || 'http://localhost:8080/api',
)

type Params = Record<string, string | number | boolean | null | undefined>

interface RequestOptions {
  method?: string
  data?: unknown
  includeAuth?: boolean
  params?: Params
}

const getAuthToken = (): string | null => localStorage.getItem(AUTH_TOKEN_KEY)

const buildHeaders = (includeAuth: boolean, hasBody: boolean): Record<string, string> => {
  const headers: Record<string, string> = {}

  if (hasBody) {
    headers['Content-Type'] = 'application/json'
  }

  if (includeAuth) {
    const token = getAuthToken()
    if (token) {
      headers.Authorization = `Bearer ${token}`
    }
  }

  return headers
}

const clearAuthStorage = () => {
  localStorage.removeItem(AUTH_TOKEN_KEY)
  localStorage.removeItem(AUTH_USER_KEY)
}

const responseMessage = (body: Partial<ApiResponse<unknown>>, fallback: string) => {
  if (Array.isArray(body.errors) && body.errors.length > 0) {
    return body.errors.join(', ')
  }
  return body.message || fallback
}

class ApiClient {
  constructor(private readonly baseUrl: string) {}

  async request<T>(endpoint: string, options: RequestOptions = {}): Promise<ApiResponse<T>> {
    const { method = 'GET', data, includeAuth = true, params = {} } = options
    const url = new URL(`${this.baseUrl}${endpoint}`)

    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined && value !== null && value !== '') {
        url.searchParams.set(key, String(value))
      }
    })

    const hasBody = data !== undefined && data !== null
    const response = await fetch(url.toString(), {
      method,
      headers: buildHeaders(includeAuth, hasBody),
      body: hasBody ? JSON.stringify(data) : undefined,
    })

    const body = await response.json().catch(() => null)

    if (response.status === 401) {
      clearAuthStorage()
    }

    if (!response.ok) {
      throw new Error(
        body && typeof body === 'object'
          ? responseMessage(body as Partial<ApiResponse<unknown>>, `HTTP ${response.status}`)
          : `HTTP ${response.status}`,
      )
    }

    if (!body || typeof body !== 'object') {
      throw new Error('Respons backend tidak valid')
    }

    const apiBody = body as ApiResponse<T>
    if (apiBody.success === false) {
      throw new Error(responseMessage(apiBody, 'Operasi gagal'))
    }

    return apiBody
  }

  get<T>(endpoint: string, options?: Omit<RequestOptions, 'method' | 'data'>) {
    return this.request<T>(endpoint, { ...options, method: 'GET' })
  }

  post<T>(endpoint: string, data?: unknown, options?: Omit<RequestOptions, 'method' | 'data'>) {
    return this.request<T>(endpoint, { ...options, method: 'POST', data })
  }

  put<T>(endpoint: string, data?: unknown, options?: Omit<RequestOptions, 'method' | 'data'>) {
    return this.request<T>(endpoint, { ...options, method: 'PUT', data })
  }

  delete<T>(endpoint: string, options?: Omit<RequestOptions, 'method' | 'data'>) {
    return this.request<T>(endpoint, { ...options, method: 'DELETE' })
  }

  async download(endpoint: string, params: Params = {}) {
    const url = new URL(`${this.baseUrl}${endpoint}`)

    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined && value !== null && value !== '') {
        url.searchParams.set(key, String(value))
      }
    })

    const response = await fetch(url.toString(), {
      method: 'GET',
      headers: buildHeaders(true, false),
    })

    if (response.status === 401) {
      clearAuthStorage()
    }

    if (!response.ok) {
      const body = await response.json().catch(() => null)
      throw new Error(
        body && typeof body === 'object'
          ? responseMessage(body as Partial<ApiResponse<unknown>>, `HTTP ${response.status}`)
          : `HTTP ${response.status}`,
      )
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
