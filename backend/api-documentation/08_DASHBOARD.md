# API Documentation - Dashboard (Analytics & Reporting)

**Base URL:** `http://localhost:8080/api`

**All endpoints require Authorization header:**

```
Authorization: Bearer {token}
```

**Access Control:** Admin or Rumah Sakit

---

## Status Summary (Dashboard Stats)

**GET** `/dashboard/status-summary/{rumahSakitID}`

**Access:** Admin (all hospitals) or Rumah Sakit (own hospital only)

**Path Parameters:**

- `rumahSakitID` (required): Hospital ID (e.g., RS001, or "all" for admin to see all)

**Purpose:** Get summary of blood request statuses for dashboard display

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Status summary retrieved successfully",
  "data": {
    "rumahSakitId": "RS001",
    "rumahSakitNama": "Rumah Sakit Pusat",
    "totalRequests": 50,
    "statusBreakdown": {
      "dibuat": 5,
      "diproses": 10,
      "bisa_diambil": 8,
      "selesai": 25,
      "dibatalkan": 2
    },
    "recentRequests": [
      {
        "permintaanDarahId": "PD04071430001",
        "namaPasien": "Budi Santoso",
        "status": "diproses",
        "tanggalPermintaan": "2026-04-07T14:30:45Z",
        "tanggalUpdated": "2026-04-07T15:00:00Z"
      },
      {
        "permintaanDarahId": "PD04071430002",
        "namaPasien": "Ani Wijaya",
        "status": "bisa_diambil",
        "tanggalPermintaan": "2026-04-07T13:20:00Z",
        "tanggalUpdated": "2026-04-07T14:45:00Z"
      }
    ]
  }
}
```

---

## Usage Examples

### Admin - View All Hospitals Summary

```bash
GET /api/dashboard/status-summary/all
Authorization: Bearer {token}
```

**Response:** Combined stats for all hospitals

---

### Admin - View Specific Hospital Summary

```bash
GET /api/dashboard/status-summary/RS001
Authorization: Bearer {token}
```

**Response:** Stats for RS001 only

---

### Rumah Sakit - View Own Hospital Summary

```bash
GET /api/dashboard/status-summary/RS001
Authorization: Bearer {token}
```

**Note:** Rumah Sakit can only access their own hospital data. If they try to access another hospital, they'll get 403 Forbidden.

---

## Dashboard Metrics

### Status Breakdown

| Status         | Meaning                                      |
|----------------|----------------------------------------------|
| dibuat         | Newly created, awaiting processing           |
| diproses       | Being processed by blood bank                |
| bisa_diambil   | Ready for hospital to pick up                |
| selesai        | Completed, blood has been picked up          |
| dibatalkan     | Cancelled (no longer needed)                 |

### Key Metrics

- **totalRequests** - Total number of requests for selected hospital
- **statusBreakdown** - Count of requests in each status
- **recentRequests** - Latest 10 requests with current status

---

## Notes

- Dashboard stats are real-time
- Admin can view all hospitals or specific hospital
- Rumah Sakit can only view their own data
- Used for monitoring and reporting purposes
- Recent requests help track urgent or pending requests

---

**Last Updated:** 2026-04-07
