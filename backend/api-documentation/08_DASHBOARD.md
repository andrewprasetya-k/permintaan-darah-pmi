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
  "message": "Operation successful",
  "data": {
    "dibuat": 5,
    "diproses": 10,
    "bisaDiambil": 8,
    "selesai": 25,
    "dibatalkan": 2,
    "total": 50
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

**Note:** Rumah Sakit token will always be resolved to their own hospital ID by the controller, so the path parameter is ignored for that role.

---

## Dashboard Metrics

### Returned Metrics

- `dibuat` - Count of newly created requests
- `diproses` - Count of requests being processed
- `bisaDiambil` - Count of requests ready to be picked up
- `selesai` - Count of completed requests
- `dibatalkan` - Count of cancelled requests
- `total` - Total count across all statuses

---

## Notes

- Dashboard stats are derived from current database state
- Admin can view all hospitals or specific hospital
- Rumah Sakit can only view their own data
- Current endpoint returns status counts only
- Recent requests and richer analytics are not implemented yet

---

**Last Updated:** 2026-04-27
