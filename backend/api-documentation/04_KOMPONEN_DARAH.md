# API Documentation - Komponen Darah

**Base URL:** `http://localhost:8080/api`

**All endpoints require Authorization header:**

```
Authorization: Bearer {token}
```

---

## Create Komponen Darah

**POST** `/komponen-darah`

**Request:**

```json
{
  "komNama": "Packed Red Cell (PRC)",
  "komKode": "PRC"
}
```

**Response (201 Created):**

```json
{
  "success": true,
  "message": "Created successfully",
  "data": {
    "komId": 1,
    "komNama": "Packed Red Cell (PRC)",
    "komKode": "PRC",
    "isActive": true
  }
}
```

---

## Get All Komponen Darah

**GET** `/komponen-darah?limit=20&offset=0`

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
      "komId": 1,
      "komNama": "Whole Blood",
      "komKode": "WB",
      "isActive": true
    },
    {
      "komId": 2,
      "komNama": "Packed Red Cell (PRC)",
      "komKode": "PRC",
      "isActive": true
    }
  ],
  "pagination": {
    "total": 2,
    "page": 1,
    "limit": 20,
    "offset": 0
  }
}
```

---

## Get Komponen Darah by ID

**GET** `/komponen-darah/{id}`

**Path Parameters:**

- `id` (required): Komponen Darah ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": {
    "komId": 1,
    "komNama": "Whole Blood",
    "komKode": "WB",
    "isActive": true
  }
}
```

---

## Update Komponen Darah

**PUT** `/komponen-darah/{id}`

**Path Parameters:**

- `id` (required): Komponen Darah ID

**Request:**

```json
{
  "komNama": "Whole Blood Updated",
  "komKode": "WB_UPD"
}
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": {
    "komId": 1,
    "komNama": "Whole Blood Updated",
    "komKode": "WB_UPD",
    "isActive": true
  }
}
```

---

## Activate Komponen Darah

**PUT** `/komponen-darah/activate/{id}`

**Path Parameters:**

- `id` (required): Komponen Darah ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": {
    "komId": 1,
    "komNama": "Whole Blood",
    "komKode": "WB",
    "isActive": true
  }
}
```

---

## Deactivate Komponen Darah

**PUT** `/komponen-darah/deactivate/{id}`

**Path Parameters:**

- `id` (required): Komponen Darah ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": {
    "komId": 1,
    "komNama": "Whole Blood",
    "komKode": "WB",
    "isActive": false
  }
}
```

---

## Delete Komponen Darah

**DELETE** `/komponen-darah/{id}`

**Path Parameters:**

- `id` (required): Komponen Darah ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": null
}
```

---

## Komponen Darah Kode Reference

- `WB` - Whole Blood
- `PRC` - Packed Red Cell
- `WE` - Washed Erythrocyte
- `FFP` - Fresh Frozen Plasma
- `PRP` - Platelet Rich Plasma
- `TC` - Thrombocyte Concentrate
- `CRY` - Cryoprecipitate
- `PRCL` - PRC Leukodepleted

---

## Error Codes

- `400` - Bad Request (Invalid input)
- `401` - Unauthorized (Missing/invalid token)
- `403` - Forbidden (No permission)
- `404` - Not Found
- `500` - Internal Server Error
