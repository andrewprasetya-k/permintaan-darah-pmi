# API Documentation - Logs (Audit & Status Trails)

**Base URL:** `http://localhost:8080/api`

**All endpoints require Authorization header:**

```
Authorization: Bearer {token}
```

**Access Control:** All endpoints - **Admin Only**

---

## Status Logs (Status Change History)

### Get All Status Logs

**GET** `/status-logs?limit=20&offset=0&pdId={pdId}`

**Access:** Admin only

**Query Parameters:**

- `limit` (optional): default 20, max 100
- `offset` (optional): default 0
- `pdId` (optional): filter by blood request ID

**Purpose:** View history of all status changes for blood requests

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Status logs retrieved successfully",
  "data": [
    {
      "logId": 1,
      "pdId": "PD04071430001",
      "adminId": "ADM-001",
      "adminNama": "Admin One",
      "statusFrom": "dibuat",
      "statusTo": "diproses",
      "notes": "Status berubah dari dibuat menjadi diproses. Alasan: Verifikasi completed",
      "createdAt": "2026-04-07T14:35:00Z"
    },
    {
      "logId": 2,
      "pdId": "PD04071430001",
      "adminId": "ADM-001",
      "adminNama": "Admin One",
      "statusFrom": "diproses",
      "statusTo": "bisa_diambil",
      "notes": "Status berubah dari diproses menjadi bisa_diambil",
      "createdAt": "2026-04-07T15:30:00Z"
    }
  ],
  "pagination": {
    "total": 25,
    "page": 1,
    "limit": 20,
    "offset": 0
  }
}
```

---

## System Access Logs (Complete Action Audit Trail)

### Get All System Logs

**GET** `/system-logs?limit=20&offset=0`

**Access:** Admin only

**Query Parameters:**

- `limit` (optional): default 20, max 100
- `offset` (optional): default 0

**Purpose:** View complete audit trail of all admin actions

**Response (200 OK):**

```json
{
  "success": true,
  "message": "System logs retrieved successfully",
  "data": [
    {
      "sysLogId": 1001,
      "userId": "ADM-001",
      "userName": "Admin One",
      "userRole": "superadmin",
      "action": "LOGIN",
      "targetTable": null,
      "targetId": null,
      "notes": "Super admin login from web",
      "ipAddress": "192.168.1.100",
      "userAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
      "createdAt": "2026-04-07T14:00:00Z"
    },
    {
      "sysLogId": 1002,
      "userId": "ADM-001",
      "userName": "Admin One",
      "userRole": "superadmin",
      "action": "CREATE",
      "targetTable": "permintaan_darah",
      "targetId": "PD04071430001",
      "notes": "Created blood request for patient Budi Santoso",
      "ipAddress": "192.168.1.100",
      "userAgent": "Mozilla/5.0...",
      "createdAt": "2026-04-07T14:30:45Z"
    }
  ],
  "pagination": {
    "total": 500,
    "page": 1,
    "limit": 20,
    "offset": 0
  }
}
```

---

### Get System Log by ID

**GET** `/system-logs/{id}`

**Access:** Admin only

**Response (200 OK):**

```json
{
  "success": true,
  "message": "System log retrieved successfully",
  "data": {...}
}
```

---

### Get Logs by User ID

**GET** `/system-logs/user/{userId}`

**Access:** Admin only

**Query Parameters:**

- `limit` (optional): default 20, max 100
- `offset` (optional): default 0

**Purpose:** View all actions performed by a specific user

**Response (200 OK):**

```json
{
  "success": true,
  "message": "User logs retrieved successfully",
  "data": [...],
  "pagination": {...}
}
```

---

### Get Logs by Action

**GET** `/system-logs/action/{action}`

**Access:** Admin only

**Query Parameters:**

- `limit` (optional): default 20, max 100
- `offset` (optional): default 0

**Path Parameters:**

- `action` (required): CREATE | UPDATE | DELETE | RESTORE | LOGIN | LOGIN_FAILED

**Purpose:** View all logs for a specific action type

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Action logs retrieved successfully",
  "data": [...],
  "pagination": {...}
}
```

---

### Get Logs by Target Table

**GET** `/system-logs/table/{table}`

**Access:** Admin only

**Query Parameters:**

- `limit` (optional): default 20, max 100
- `offset` (optional): default 0

**Path Parameters:**

- `table` (required): admins | rumah_sakit | komponen_darah | permintaan_darah | etc.

**Purpose:** View all actions on a specific table

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Table logs retrieved successfully",
  "data": [...],
  "pagination": {...}
}
```

---

### Get Logs by Target ID

**GET** `/system-logs/target/{targetId}`

**Access:** Admin only

**Query Parameters:**

- `limit` (optional): default 20, max 100
- `offset` (optional): default 0

**Path Parameters:**

- `targetId` (required): ID of the entity (e.g., PD04071430001, ADM-001)

**Purpose:** View all actions on a specific entity

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Target logs retrieved successfully",
  "data": [...],
  "pagination": {...}
}
```

---

## Log Action Types

| Action       | Description                    |
|--------------|--------------------------------|
| CREATE       | New record created             |
| UPDATE       | Record updated                 |
| DELETE       | Record soft-deleted            |
| SOFT_DELETE  | Mark as deleted (same as DELETE)|
| RESTORE      | Restore deleted record         |
| LOGIN        | Successful user login          |
| LOGIN_FAILED | Failed login attempt           |
| UPDATE_STATUS| Status change (for requests)   |

---

## Notes

- Status logs track status changes with before/after states
- System logs track ALL actions (CREATE, UPDATE, DELETE, LOGIN)
- Both logs capture IP address and user agent for security
- Logs are immutable (cannot be edited or deleted)
- Used for compliance and security audits
- Admin names are captured at time of action (preserved even if admin is deleted)

---

**Last Updated:** 2026-04-07
