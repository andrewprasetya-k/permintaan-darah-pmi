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

**GET** `/status-logs?limit=20&offset=0`

**Access:** Admin only

**Query Parameters:**

- `limit` (optional): default 20
- `offset` (optional): default 0

**Purpose:** View history of all status changes for blood requests

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [
    {
      "logId": 1,
      "permintaanDarahId": "PD04071430001",
      "adminId": "ADM-001",
      "adminNama": "Admin One",
      "statusFrom": "dibuat",
      "statusTo": "diproses",
      "notes": "Status berubah dari dibuat menjadi diproses. Alasan: Verifikasi completed",
      "createdAt": "2026-04-07T14:35:00Z"
    },
    {
      "logId": 2,
      "permintaanDarahId": "PD04071430001",
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

- `limit` (optional): default 20
- `offset` (optional): default 0

**Purpose:** View complete audit trail of all admin actions

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": [
    {
      "sysLogId": 1001,
      "sysUserId": "ADM-001",
      "sysUserNama": "Admin One",
      "sysUserRole": "superadmin",
      "sysAction": "LOGIN",
      "sysTargetTable": null,
      "sysTargetId": null,
      "sysNotes": "Super admin login from web",
      "sysUserAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
      "createdAt": "2026-04-07T14:00:00Z"
    },
    {
      "sysLogId": 1002,
      "sysUserId": "ADM-001",
      "sysUserNama": "Admin One",
      "sysUserRole": "superadmin",
      "sysAction": "CREATE",
      "sysTargetTable": "permintaan_darah",
      "sysTargetId": "PD04071430001",
      "sysNotes": "Created blood request for patient Budi Santoso",
      "sysUserAgent": "Mozilla/5.0...",
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
  "sysLogId": 1001,
  "sysUserId": "ADM-001",
  "sysUserNama": "Admin One",
  "sysUserRole": "superadmin",
  "sysAction": "LOGIN",
  "sysTargetTable": null,
  "sysTargetId": null,
  "sysNotes": "Admin Admin One login successfully",
  "sysUserAgent": null,
  "createdAt": "2026-04-07T14:00:00Z"
}
```

---

### Get Logs by User ID

**GET** `/system-logs/user/{userId}`

**Access:** Admin only

**Query Parameters:**

- `limit` (optional): default 20
- `offset` (optional): default 0

**Purpose:** View all actions performed by a specific user

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": [...],
  "pagination": {...}
}
```

---

### Get Logs by Action

**GET** `/system-logs/action/{action}`

**Access:** Admin only

**Query Parameters:**

- `limit` (optional): default 20
- `offset` (optional): default 0

**Path Parameters:**

- `action` (required): examples include CREATE | UPDATE | DELETE | SOFT_DELETE | RESTORE | LOGIN | LOGIN_FAILED | UPDATE_STATUS | ACTIVATE | DEACTIVATE

**Purpose:** View all logs for a specific action type

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": [...],
  "pagination": {...}
}
```

---

### Get Logs by Target Table

**GET** `/system-logs/table/{table}`

**Access:** Admin only

**Query Parameters:**

- `limit` (optional): default 20
- `offset` (optional): default 0

**Path Parameters:**

- `table` (required): admins | rumah_sakit | komponen_darah | permintaan_darah | etc.

**Purpose:** View all actions on a specific table

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": [...],
  "pagination": {...}
}
```

---

### Get Logs by Target ID

**GET** `/system-logs/target/{targetId}`

**Access:** Admin only

**Query Parameters:**

- `limit` (optional): default 20
- `offset` (optional): default 0

**Path Parameters:**

- `targetId` (required): ID of the entity (e.g., PD04071430001, ADM-001)

**Purpose:** View all actions on a specific entity

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
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
| ACTIVATE     | Component activated            |
| DEACTIVATE   | Component deactivated          |

---

## Notes

- Status logs track status changes with before/after states
- System logs track actions such as CREATE, UPDATE, DELETE, LOGIN, UPDATE_STATUS, ACTIVATE, and DEACTIVATE
- System logs capture user agent; `sys_ip_address` exists in the SQL schema but is not mapped in the current Go model/DTO
- Logs are immutable (cannot be edited or deleted)
- Used for compliance and security audits
- Admin names are captured at time of action (preserved even if admin is deleted)

---

**Last Updated:** 2026-04-27
