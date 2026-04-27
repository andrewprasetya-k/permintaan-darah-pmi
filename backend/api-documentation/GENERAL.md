# API Documentation - General

**Base URL:** `http://localhost:8080/api`

---

## Standard Response Format

### Success Response

```json
{
  "success": true,
  "message": "Operation description",
  "data": {
    // Response data here
  },
  "pagination": {
    "total": 100,
    "page": 1,
    "limit": 20,
    "offset": 0
  }
}
```

**Note:**

- `pagination` field hanya ada untuk endpoint yang mengembalikan list data
- `pagination` field diabaikan jika tidak ada (omitempty)

---

### Error Response

```json
{
  "success": false,
  "message": "Error message",
  "errors": ["Additional error detail 1", "Additional error detail 2"]
}
```

JWT/RBAC middleware adalah pengecualian. Saat token missing/invalid atau role tidak sesuai, middleware mengembalikan payload sederhana seperti:

```json
{
  "error": "Invalid token"
}
```

---

## HTTP Status Codes

| Code | Meaning                                   |
| ---- | ----------------------------------------- |
| 200  | OK - Request successful                   |
| 201  | Created - Resource created successfully   |
| 400  | Bad Request - Invalid input data          |
| 401  | Unauthorized - Missing or invalid token   |
| 403  | Forbidden - Access denied / No permission |
| 404  | Not Found - Resource not found            |
| 409  | Conflict - Resource already exists        |
| 422  | Unprocessable Entity - Validation error   |
| 500  | Internal Server Error - Server error      |

---

## Pagination

### Query Parameters

- `limit` - Number of items per page (default: 20)
- `offset` - Starting position (default: 0)

### Pagination Response

```json
{
  "pagination": {
    "total": 150, // Total number of items in database
    "page": 2, // Current page number (calculated from offset/limit)
    "limit": 20, // Items per page
    "offset": 20 // Current offset
  }
}
```

### Example Request

```
GET /api/admin?limit=50&offset=100
```

**Result:** 50 items starting from position 100 (page 3)

---

## Authentication

### Login Flow

1. Call `/auth/login/admin` or `/auth/login/rumah-sakit`
2. Get token from response
3. Use token in Authorization header for all protected endpoints

### Authorization Header Format

```
Authorization: Bearer {token}
```

### Token Details

- **Type:** JWT (JSON Web Token)
- **Expiry:** 24 hours
- **Algorithm:** HS256
- **Claims:** `user_id`, `username`, `role`, `exp`
- **Location:** Authorization header
- **Format:** `Bearer {token}`

### Example Protected Request

```
GET /api/admin
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

---

## Enum Values Reference

### Admin Role

- `superadmin` - Super administrator
- `admin` - Regular administrator

### Permintaan Status

- `dibuat` - Newly created
- `diproses` - Being processed
- `bisa_diambil` - Ready to pick up
- `selesai` - Completed
- `dibatalkan` - Cancelled

### Gender

- `L` - Laki-laki (Male)
- `P` - Perempuan (Female)

### Blood Type

- `A`
- `B`
- `AB`
- `O`

### Rhesus Factor

- `+` - Positive
- `-` - Negative

### System Log Actions

- `CREATE` - Create new record
- `UPDATE` - Update record
- `DELETE` - Soft delete record
- `SOFT_DELETE` - Mark as deleted
- `RESTORE` - Restore deleted record
- `LOGIN` - User login
- `LOGIN_FAILED` - Failed login attempt
- `UPDATE_STATUS` - Permintaan status changed
- `ACTIVATE` - Component activated
- `DEACTIVATE` - Component deactivated

---

## Common Error Messages

| Message             | Status | Meaning                        |
| ------------------- | ------ | ------------------------------ |
| Invalid credentials | 401    | Wrong username or password     |
| Data not found      | 404    | Resource tidak ada di database |
| Invalid input       | 400    | Ada field yang tidak valid     |
| Unauthorized        | 401    | Token missing atau expired     |
| Forbidden           | 403    | User tidak punya permission    |
| Data already exists | 409    | Resource sudah ada (duplicate) |

---

## Date Format

All dates menggunakan ISO 8601 format dengan timezone UTC:

```
2026-04-06T08:33:32Z
```

Format:

- Year: YYYY (4 digits)
- Month: MM (2 digits)
- Day: DD (2 digits)
- Hour: HH (2 digits, 24-hour format)
- Minute: mm (2 digits)
- Second: ss (2 digits)
- Timezone: Z (UTC indicator)

---

## WebSocket Authentication

### Token Location for WebSocket

Karena WebSocket upgrade HTTP tidak allow custom headers (RFC 6455), token harus dikirim via query parameter:

```
ws://localhost:8080/api/ws/connect?token=JWT_TOKEN
```

### WebSocket Flow

1. Connect ke `/ws/connect` dengan token di query param
2. JWTMiddleware validate token dari query param
3. Extract userID & userRole dari token claims
4. Receive real-time messages: `{type, action, entityId, entityType, data, timestamp}`

### WebSocket Message Format

```json
{
  "type": "status_change",
  "action": "UPDATE",
  "entityId": "PD20240407123456",
  "entityType": "permintaan_darah",
  "data": {
    "permintaanDarahId": "PD20240407123456",
    "namaPasien": "Budi Santoso",
    "status": "diproses",
    "createdAt": "2026-04-07T10:30:45Z",
    "updatedAt": "2026-04-07T11:00:12Z"
  },
  "timestamp": "2026-04-07T11:00:12Z"
}
```

### Supported Message Types

- `new_permintaan` - Permintaan dibuat
- `update_permintaan` - Permintaan diupdate
- `status_change` - Permintaan status berubah
- `new_activity` - System activity log baru

### Connection Features

- **Timeout:** 24 jam (tetap connected selama admin session)
- **Heartbeat:** Ping every 30 seconds untuk prevent network timeout
- **Auto Response:** Browser otomatis respond dengan pong

---

## Rate Limiting

Berbagai endpoint memiliki rate limiting untuk prevent abuse:

| Endpoint        | Limit           | Window  |
| --------------- | --------------- | ------- |
| Login endpoints | 5 requests      | 1 menit |
| Protected API   | 100 requests    | 1 menit |
| Strict endpoints| 30 requests     | 1 menit |

Catatan: helper strict rate limiter tersedia di kode, tetapi belum dipasang pada route di `routes/api.go` saat ini.

Response header saat rate limit exceeded:

```
HTTP/1.1 429 Too Many Requests
Retry-After: 45
X-RateLimit-Limit: 5
X-RateLimit-Remaining: 0
X-RateLimit-Reset: 1712481234
```

---

## Dokumentasi Files

Setiap endpoint diorganisir per modul:

- `01_AUTH.md` - Authentication endpoints (Login)
- `02_ADMIN.md` - Admin management (CRUD)
- `03_RUMAH_SAKIT.md` - Rumah Sakit management (CRUD)
- `04_KOMPONEN_DARAH.md` - Blood component management (CRUD)
- `05_PERMINTAAN_DARAH.md` - Blood request management (CRUD + Status)
- `07_LOGS.md` - Status & system access logs
- `08_DASHBOARD.md` - Dashboard & analytics
- `GENERAL.md` - This file (General guidelines)

---

## Access Control Summary

| Role        | Can Access                                                        |
| ----------- | ----------------------------------------------------------------- |
| Superadmin  | Admin CRUD, hospital CRUD, components, all permintaan, all logs  |
| Admin       | Hospital view/update, components, all permintaan, all logs       |
| Rumah Sakit | Own profile, own permintaan, own request status update, dashboard |

---

**Last Updated:** 2026-04-27  
**API Version:** 1.0.0
