# API Documentation - Authentication

**Base URL:** `http://localhost:8080/api`

---

## Admin Login

**POST** `/auth/login/admin`

**Rate Limited:** 5 attempts per minute

**Request:**

```json
{
  "username": "admin@example.com",
  "password": "password123"
}
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "id": "ADM-001",
    "username": "admin@example.com",
    "role": "superadmin",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiQURNLTAwMSIsInVzZXJuYW1lIjoiYWRtaW4iLCJyb2xlIjoic3VwZXJhZG1pbiIsImV4cCI6MTcxMjU2ODAwMH0..."
  }
}
```

**Response (401 Unauthorized):**

```json
{
  "success": false,
  "message": "Invalid credentials"
}
```

**Response (429 Too Many Requests):**

```json
{
  "success": false,
  "message": "Rate limit exceeded. Please try again in 45 seconds.",
  "retryAfterSeconds": 45
}
```

---

## Rumah Sakit Login

**POST** `/auth/login/rumah-sakit`

**Rate Limited:** 5 attempts per minute

**Request:**

```json
{
  "username": "rs_username",
  "password": "password123"
}
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "id": "RS001",
    "username": "rs_username",
    "role": "rumah_sakit",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiUlMwMDEiLCJ1c2VybmFtZSI6InJzX3VzZXJuYW1lIiwicm9sZSI6InJ1bWFoX3Nha2l0IiwiZXhwIjoxNzEyNTY4MDAwfQ..."
  }
}
```

**Response (401 Unauthorized):**

```json
{
  "success": false,
  "message": "Invalid credentials"
}
```

---

## Authentication Details

### Token Information

- **Type:** JWT (JSON Web Token)
- **Algorithm:** HS256
- **Expiry:** 24 hours
- **Claims:** `user_id`, `username`, `role`, `exp`

### Token Usage

Token returned from login must be included in Authorization header untuk semua protected endpoints:

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### WebSocket Authentication

Untuk WebSocket endpoint, token dikirim via query parameter (karena HTTP RFC 6455 tidak allow custom headers saat upgrade):

```
ws://localhost:8080/api/ws/connect?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

Lihat `GENERAL.md` untuk details WebSocket.

---

## Notes

- Token harus disimpan secara aman di client (localStorage atau secure cookie)
- Token otomatis expire setelah 24 jam (user perlu login lagi)
- Token tidak bisa di-refresh (harus login ulang untuk new token)
- Login endpoints menerapkan rate limiting ketat (5 attempts/menit) untuk prevent brute force
- Request/response payload sensitive (passwords) tidak di-log untuk security
- JWT middleware untuk protected endpoint mengembalikan payload sederhana seperti `{"error":"Invalid token"}` saat token missing, malformed, atau invalid

---

**Last Updated:** 2026-04-07
