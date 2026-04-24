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
  "permintaanDarahId": "PD04071430001",
  "komponenDarahId": 2,
  "golonganDarah": "O",
  "rhesusDarah": "+",
  "jumlahKantong": 3
}
```

**Response (201 Created):**

```json
{
  "success": true,
  "message": "Detail created successfully",
  "data": {
    "detailId": 2,
    "permintaanDarahId": "PD04071430001",
    "komponenNama": "Packed Red Cell (PRC)",
    "golonganDarah": "O",
    "rhesusDarah": "+",
    "jumlahKantong": 3,
    "tanggalDiperlukan": "0001-01-01T00:00:00Z",
    "createdAt": "2026-04-07T14:35:00Z"
  }
}
```

---

## Get All Details

**GET** `/detail-permintaan-darah?limit=20&offset=0`

**Access:** Rumah Sakit only

**Query Parameters:**

- `limit` (optional): default 20, max 100
- `offset` (optional): default 0

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
    "detailId": 1,
    "permintaanDarahId": "PD04071430001",
    "komponenNama": "Whole Blood",
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
  "permintaanDarahId": "PD04071430001",
  "komponenDarahId": 2,
  "golonganDarah": "O",
  "rhesusDarah": "+",
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
- Current list endpoint supports pagination only; it does not yet expose filter by `permintaanDarahId`

---

**Last Updated:** 2026-04-07
