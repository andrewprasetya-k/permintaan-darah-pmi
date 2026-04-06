# API Documentation - Rumah Sakit

**Base URL:** `http://localhost:8080/api`

**All endpoints require Authorization header:**

```
Authorization: Bearer {token}
```

---

## Create Rumah Sakit

**POST** `/rumah-sakit`

**Request:**

```json
{
  "rsNama": "Rumah Sakit Harapan",
  "rsNoTelp": "021-1234567",
  "rsAlamat": "Jl. Merdeka No. 10, Jakarta",
  "rsEmail": "rs@example.com",
  "rsUsername": "rs_harapan",
  "rsPassword": "password123"
}
```

**Response (201 Created):**

```json
{
  "success": true,
  "message": "Rumah Sakit created successfully",
  "data": {
    "rsId": "RS001",
    "rsNama": "Rumah Sakit Harapan",
    "rsNoTelp": "021-1234567",
    "rsAlamat": "Jl. Merdeka No. 10, Jakarta",
    "rsEmail": "rs@example.com",
    "rsUsername": "rs_harapan",
    "createdAt": "2026-04-06T08:33:32Z",
    "updatedAt": "2026-04-06T08:33:32Z"
  }
}
```

---

## Get All Rumah Sakit

**GET** `/rumah-sakit?limit=20&offset=0`

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
      "rsId": "RS001",
      "rsNama": "Rumah Sakit Harapan",
      "rsNoTelp": "021-1234567",
      "rsAlamat": "Jl. Merdeka No. 10, Jakarta",
      "rsEmail": "rs@example.com",
      "rsUsername": "rs_harapan",
      "createdAt": "2026-04-06T08:33:32Z",
      "updatedAt": "2026-04-06T08:33:32Z"
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

## Get Rumah Sakit by ID

**GET** `/rumah-sakit/{id}`

**Path Parameters:**

- `id` (required): Rumah Sakit ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Rumah Sakit retrieved successfully",
  "data": {
    "rsId": "RS001",
    "rsNama": "Rumah Sakit Harapan",
    "rsNoTelp": "021-1234567",
    "rsAlamat": "Jl. Merdeka No. 10, Jakarta",
    "rsEmail": "rs@example.com",
    "rsUsername": "rs_harapan",
    "createdAt": "2026-04-06T08:33:32Z",
    "updatedAt": "2026-04-06T08:33:32Z"
  }
}
```

---

## Update Rumah Sakit

**PUT** `/rumah-sakit/{id}`

**Path Parameters:**

- `id` (required): Rumah Sakit ID

**Request:**

```json
{
  "rsNama": "Rumah Sakit Harapan Baru",
  "rsNoTelp": "021-7654321",
  "rsAlamat": "Jl. Merdeka No. 20, Jakarta",
  "rsEmail": "rs_baru@example.com",
  "rsUsername": "rs_harapan_baru",
  "rsPassword": "newpassword123"
}
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Rumah Sakit updated successfully",
  "data": {
    "rsId": "RS001",
    "rsNama": "Rumah Sakit Harapan Baru",
    "rsNoTelp": "021-7654321",
    "rsAlamat": "Jl. Merdeka No. 20, Jakarta",
    "rsEmail": "rs_baru@example.com",
    "rsUsername": "rs_harapan_baru",
    "createdAt": "2026-04-06T08:33:32Z",
    "updatedAt": "2026-04-06T09:00:00Z"
  }
}
```

---

## Delete Rumah Sakit (Soft Delete)

**DELETE** `/rumah-sakit/{id}`

**Path Parameters:**

- `id` (required): Rumah Sakit ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Rumah Sakit deleted successfully",
  "data": null
}
```

---

## Restore Rumah Sakit

**PUT** `/rumah-sakit/restore/{id}`

**Path Parameters:**

- `id` (required): Rumah Sakit ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Rumah Sakit restored successfully",
  "data": null
}
```

---

## Get Distinct Rumah Sakit Names

**GET** `/filter/rumah-sakit/`

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [
    {
      "rsId": "RS001",
      "rsNama": "Rumah Sakit Harapan"
    },
    {
      "rsId": "RS002",
      "rsNama": "Rumah Sakit Medika"
    }
  ]
}
```

---

## Error Codes

- `400` - Bad Request (Invalid input)
- `401` - Unauthorized (Missing/invalid token)
- `403` - Forbidden (No permission)
- `404` - Not Found
- `500` - Internal Server Error
