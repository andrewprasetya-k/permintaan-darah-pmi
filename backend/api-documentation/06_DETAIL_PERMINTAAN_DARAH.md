# API Documentation - Detail Permintaan Darah

**Base URL:** `http://localhost:8080/api`

**All endpoints require Authorization header:**

```
Authorization: Bearer {token}
```

**Access Control:** All endpoints - **Rumah Sakit Only**

---

## Create Detail

**POST** `/detail-permintaan-darah`

**Access:** Rumah Sakit only

**Request:**

```json
{
  "pdId": "PD04071430001",
  "komId": 2,
  "golonganDarah": "O",
  "rhesusDarah": "+",
  "jumlahKantong": 3,
  "tanggalDiperlukan": "2026-04-10T15:00:00Z"
}
```

**Response (201 Created):**

```json
{
  "success": true,
  "message": "Detail created successfully",
  "data": {
    "dpdId": 2,
    "pdId": "PD04071430001",
    "komId": 2,
    "komNama": "Packed Red Cell (PRC)",
    "golonganDarah": "O",
    "rhesusDarah": "+",
    "jumlahKantong": 3,
    "tanggalDiperlukan": "2026-04-10T15:00:00Z",
    "createdAt": "2026-04-07T14:35:00Z"
  }
}
```

---

## Get All Details

**GET** `/detail-permintaan-darah?limit=20&offset=0&pdId={pdId}`

**Access:** Rumah Sakit only

**Query Parameters:**

- `limit` (optional): default 20, max 100
- `offset` (optional): default 0
- `pdId` (optional): filter by blood request ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Details retrieved successfully",
  "data": [...],
  "pagination": {
    "total": 5,
    "page": 1,
    "limit": 20,
    "offset": 0
  }
}
```

---

## Get Detail by ID

**GET** `/detail-permintaan-darah/{id}`

**Access:** Rumah Sakit only

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Detail retrieved successfully",
  "data": {
    "dpdId": 1,
    "pdId": "PD04071430001",
    "komId": 1,
    "komNama": "Whole Blood",
    "golonganDarah": "O",
    "rhesusDarah": "+",
    "jumlahKantong": 2,
    "tanggalDiperlukan": "2026-04-10T10:00:00Z",
    "createdAt": "2026-04-07T14:30:45Z"
  }
}
```

---

## Update Detail

**PUT** `/detail-permintaan-darah/{id}`

**Access:** Rumah Sakit only

**Request:**

```json
{
  "jumlahKantong": 4,
  "tanggalDiperlukan": "2026-04-11T10:00:00Z"
}
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Detail updated successfully",
  "data": {...}
}
```

---

## Delete Detail

**DELETE** `/detail-permintaan-darah/{id}`

**Access:** Rumah Sakit only

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Detail deleted successfully",
  "data": null
}
```

---

## Notes

- Each detail specifies one blood component with quantity and needed date
- One request (permintaan_darah) can have multiple details
- Used when rumah sakit needs multiple components for one patient
- Component must be active (isActive: true)

---

**Last Updated:** 2026-04-07
