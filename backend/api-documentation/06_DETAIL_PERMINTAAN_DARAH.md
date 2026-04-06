# API Documentation - Detail Permintaan Darah

**Base URL:** `http://localhost:8080/api`

**All endpoints require Authorization header:**

```
Authorization: Bearer {token}
```

---

## Create Detail Permintaan Darah

**POST** `/detail-permintaan-darah`

**Request:**

```json
{
  "dpdPdId": "PD0406082826001",
  "dpdKomId": 1,
  "dpdGolonganDarah": "O",
  "dpdRhesus": "+",
  "dpdJmlKantong": 2,
  "dpdTglDiperlukan": "2026-04-08T10:00:00Z"
}
```

**Response (201 Created):**

```json
{
  "success": true,
  "message": "Created successfully",
  "data": {
    "dpdId": 1,
    "dpdPdId": "PD0406082826001",
    "komponenNama": "Packed Red Cell (PRC)",
    "golonganDarah": "O",
    "rhesusDarah": "+",
    "jmlKantong": 2,
    "tglDiperlukan": "2026-04-08T10:00:00Z",
    "createdAt": "2026-04-06T08:33:32Z"
  }
}
```

---

## Get All Detail Permintaan Darah

**GET** `/detail-permintaan-darah?limit=20&offset=0`

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
      "dpdId": 1,
      "dpdPdId": "PD0406082826001",
      "komponenNama": "Packed Red Cell (PRC)",
      "golonganDarah": "O",
      "rhesusDarah": "+",
      "jmlKantong": 2,
      "tglDiperlukan": "2026-04-08T10:00:00Z",
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

## Get Detail Permintaan Darah by ID

**GET** `/detail-permintaan-darah/{id}`

**Path Parameters:**

- `id` (required): Detail Permintaan Darah ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": {
    "dpdId": 1,
    "dpdPdId": "PD0406082826001",
    "komponenNama": "Packed Red Cell (PRC)",
    "golonganDarah": "O",
    "rhesusDarah": "+",
    "jmlKantong": 2,
    "tglDiperlukan": "2026-04-08T10:00:00Z",
    "createdAt": "2026-04-06T08:33:32Z"
  }
}
```

---

## Update Detail Permintaan Darah

**PUT** `/detail-permintaan-darah/{id}`

**Path Parameters:**

- `id` (required): Detail Permintaan Darah ID

**Request:**

```json
{
  "dpdPdId": "PD0406082826001",
  "dpdKomId": 2,
  "dpdGolonganDarah": "A",
  "dpdRhesus": "-",
  "dpdJmlKantong": 3,
  "dpdTglDiperlukan": "2026-04-09T10:00:00Z"
}
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": {
    "dpdId": 1,
    "dpdPdId": "PD0406082826001",
    "komponenNama": "Fresh Frozen Plasma (FFP)",
    "golonganDarah": "A",
    "rhesusDarah": "-",
    "jmlKantong": 3,
    "tglDiperlukan": "2026-04-09T10:00:00Z",
    "createdAt": "2026-04-06T08:33:32Z"
  }
}
```

---

## Delete Detail Permintaan Darah

**DELETE** `/detail-permintaan-darah/{id}`

**Path Parameters:**

- `id` (required): Detail Permintaan Darah ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": null
}
```

---

## Golongan Darah Values

- `A`
- `B`
- `AB`
- `O`

## Rhesus Values

- `+` - Positive
- `-` - Negative

---

## Error Codes

- `400` - Bad Request (Invalid input)
- `401` - Unauthorized (Missing/invalid token)
- `403` - Forbidden (No permission)
- `404` - Not Found
- `500` - Internal Server Error
