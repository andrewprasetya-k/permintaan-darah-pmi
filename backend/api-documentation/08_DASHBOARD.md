# API Documentation - Dashboard

**Base URL:** `http://localhost:8080/api`

**All endpoints require Authorization header:**

```
Authorization: Bearer {token}
```

---

## Get Status Summary

**GET** `/dashboard/status-summary/{rumahSakitID}`

**Path Parameters:**

- `rumahSakitID` (required): Rumah Sakit ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": {
    "rumahSakitId": "RS001",
    "dibuat": 5,
    "diproses": 3,
    "bisaDiambil": 2,
    "selesai": 10,
    "dibatalkan": 1
  }
}
```

**Response (404 Not Found):**

```json
{
  "success": false,
  "message": "Data not found",
  "errors": []
}
```

---

## Status Summary Fields

- `dibuat` - Count of requests with status "dibuat"
- `diproses` - Count of requests with status "diproses"
- `bisaDiambil` - Count of requests with status "bisa_diambil"
- `selesai` - Count of requests with status "selesai"
- `dibatalkan` - Count of requests with status "dibatalkan"

---

## Error Codes

- `401` - Unauthorized (Missing/invalid token)
- `403` - Forbidden (No permission)
- `404` - Not Found (Rumah Sakit tidak ditemukan)
- `500` - Internal Server Error
