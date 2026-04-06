# API Documentation - Status Logs & System Access Logs

**Base URL:** `http://localhost:8080/api`

**All endpoints require Authorization header:**

```
Authorization: Bearer {token}
```

---

## Get All Status Logs

**GET** `/status-logs?limit=20&offset=0`

**Query Parameters:**

- `limit` (optional): default 20, max 100
- `offset` (optional): default 0

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [
    {
      "logId": 1,
      "logPdId": "PD0406082826001",
      "logAdminId": "ADM-001",
      "logAdminNama": "Admin One",
      "logStatusFrom": "dibuat",
      "logStatusTo": "diproses",
      "logNotes": "Status updated to diproses",
      "createdAt": "2026-04-06T08:45:00Z"
    }
  ],
  "pagination": {
    "total": 1,
    "page": 1,
    "limit": 20,
    "offset": 0
  }
}
```

---

## Get All System Access Logs

**GET** `/system-logs?limit=20&offset=0`

**Query Parameters:**

- `limit` (optional): default 20, max 100
- `offset` (optional): default 0

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [
    {
      "sysLogId": 1,
      "sysUserId": "ADM-001",
      "sysUserNama": "Admin One",
      "sysUserRole": "admin",
      "sysAction": "CREATE",
      "sysTargetTable": "permintaan_darah",
      "sysTargetId": "PD0406082826001",
      "sysNotes": "Created new permintaan darah",
      "sysIpAddress": "127.0.0.1",
      "sysUserAgent": "Mozilla/5.0...",
      "createdAt": "2026-04-06T08:33:32Z"
    }
  ],
  "pagination": {
    "total": 1,
    "page": 1,
    "limit": 20,
    "offset": 0
  }
}
```

---

## Get System Access Log by ID

**GET** `/system-logs/{id}`

**Path Parameters:**

- `id` (required): System Log ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": {
    "sysLogId": 1,
    "sysUserId": "ADM-001",
    "sysUserNama": "Admin One",
    "sysUserRole": "admin",
    "sysAction": "CREATE",
    "sysTargetTable": "permintaan_darah",
    "sysTargetId": "PD0406082826001",
    "sysNotes": "Created new permintaan darah",
    "sysIpAddress": "127.0.0.1",
    "sysUserAgent": "Mozilla/5.0...",
    "createdAt": "2026-04-06T08:33:32Z"
  }
}
```

---

## Get System Access Logs by User ID

**GET** `/system-logs/user/{userId}?limit=20&offset=0`

**Path Parameters:**

- `userId` (required): User ID

**Query Parameters:**

- `limit` (optional): default 20
- `offset` (optional): default 0

**Response (200 OK):** Same as GetAll

---

## Get System Access Logs by Action

**GET** `/system-logs/action/{action}?limit=20&offset=0`

**Path Parameters:**

- `action` (required): Action type (CREATE, UPDATE, DELETE, SOFT_DELETE, RESTORE, LOGIN, etc)

**Query Parameters:**

- `limit` (optional): default 20
- `offset` (optional): default 0

**Response (200 OK):** Same as GetAll

---

## Get System Access Logs by Target Table

**GET** `/system-logs/table/{table}?limit=20&offset=0`

**Path Parameters:**

- `table` (required): Table name (admins, rumah_sakit, permintaan_darah, etc)

**Query Parameters:**

- `limit` (optional): default 20
- `offset` (optional): default 0

**Response (200 OK):** Same as GetAll

---

## Get System Access Logs by Target ID

**GET** `/system-logs/target/{targetId}?limit=20&offset=0`

**Path Parameters:**

- `targetId` (required): Target ID (can be any entity ID)

**Query Parameters:**

- `limit` (optional): default 20
- `offset` (optional): default 0

**Response (200 OK):** Same as GetAll

---

## Action Types

- `CREATE` - Create new record
- `UPDATE` - Update record
- `DELETE` - Soft delete record
- `SOFT_DELETE` - Mark as deleted
- `RESTORE` - Restore deleted record
- `LOGIN` - User login
- `LOGIN_FAILED` - Failed login attempt

---

## Error Codes

- `400` - Bad Request (Invalid input)
- `401` - Unauthorized (Missing/invalid token)
- `403` - Forbidden (No permission)
- `404` - Not Found
- `500` - Internal Server Error
