# API Documentation - Komponen Darah (Blood Components)

**Base URL:** `http://localhost:8080/api`

**All endpoints require Authorization header:**

```
Authorization: Bearer {token}
```

**Access Control:** All endpoints - **Admin Only**

---

## Create Komponen Darah (Admin Only)

**POST** `/komponen-darah`

**Access:** Admin or Superadmin only

**Request:**

```json
{
  "komponenDarah": "Packed Red Cell (PRC)",
  "komponenKode": "PRC",
  "isActive": true
}
```

**Response (201 Created):**

```json
{
  "success": true,
  "message": "Blood component created successfully",
  "data": {
    "komponenId": 2,
    "komponenDarah": "Packed Red Cell (PRC)",
    "komponenKode": "PRC",
    "isActive": true
  }
}
```

---

## Get All Komponen Darah (Admin Only)

**GET** `/komponen-darah?limit=20&offset=0`

**Access:** Admin or Superadmin only

**Query Parameters:**

- `limit` (optional): default 20, max 100
- `offset` (optional): default 0

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Blood components retrieved successfully",
  "data": [
    {
      "komponenId": 1,
      "komponenDarah": "Whole Blood",
      "komponenKode": "WB",
      "isActive": true
    },
    {
      "komponenId": 2,
      "komponenDarah": "Packed Red Cell (PRC)",
      "komponenKode": "PRC",
      "isActive": true
    },
    {
      "komponenId": 3,
      "komponenDarah": "Washed Erythrocyte (WE)",
      "komponenKode": "WE",
      "isActive": true
    }
  ],
  "pagination": {
    "total": 8,
    "page": 1,
    "limit": 20,
    "offset": 0
  }
}
```

---

## Get Komponen Darah by ID (Admin Only)

**GET** `/komponen-darah/{id}`

**Access:** Admin or Superadmin only

**Path Parameters:**

- `id` (required): Component ID (e.g., 1, 2, 3)

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Blood component retrieved successfully",
  "data": {
    "komponenId": 1,
    "komponenDarah": "Whole Blood",
    "komponenKode": "WB",
    "isActive": true
  }
}
```

**Response (404 Not Found):**

```json
{
  "success": false,
  "message": "Data not found"
}
```

---

## Update Komponen Darah (Admin Only)

**PUT** `/komponen-darah/{id}`

**Access:** Admin or Superadmin only

**Path Parameters:**

- `id` (required): Component ID

**Request:**

```json
{
  "komponenDarah": "Whole Blood - Updated",
  "komponenKode": "WB_UPD",
  "isActive": true
}
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Blood component updated successfully",
  "data": {
    "komponenId": 1,
    "komponenDarah": "Whole Blood - Updated",
    "komponenKode": "WB_UPD",
    "isActive": true
  }
}
```

---

## Activate Komponen Darah (Admin Only)

**PUT** `/komponen-darah/activate/{id}`

**Access:** Admin or Superadmin only

**Purpose:** Set component as active (available for use)

**Path Parameters:**

- `id` (required): Component ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Blood component activated successfully",
  "data": {
    "komponenId": 1,
    "komponenDarah": "Whole Blood",
    "komponenKode": "WB",
    "isActive": true
  }
}
```

---

## Deactivate Komponen Darah (Admin Only)

**PUT** `/komponen-darah/deactivate/{id}`

**Access:** Admin or Superadmin only

**Purpose:** Set component as inactive (no longer available)

**Path Parameters:**

- `id` (required): Component ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Blood component deactivated successfully",
  "data": {
    "komponenId": 1,
    "komponenDarah": "Whole Blood",
    "komponenKode": "WB",
    "isActive": false
  }
}
```

---

## Delete Komponen Darah (Admin Only)

**DELETE** `/komponen-darah/{id}`

**Access:** Admin or Superadmin only

**Path Parameters:**

- `id` (required): Component ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Blood component deleted successfully",
  "data": null
}
```

**Note:** Hard delete (not soft delete like admins/hospitals)

---

## Blood Component Reference

### Pre-seeded Components

| ID | Name                           | Code | Description                          |
|----|--------------------------------|------|--------------------------------------|
| 1  | Whole Blood                    | WB   | Complete blood with all components   |
| 2  | Packed Red Cell (PRC)          | PRC  | Red cells only                       |
| 3  | Washed Erythrocyte (WE)        | WE   | Thoroughly washed red cells          |
| 4  | Fresh Frozen Plasma (FFP)      | FFP  | Plasma with coagulation factors      |
| 5  | Platelet Rich Plasma (PRP)     | PRP  | Plasma with platelets                |
| 6  | Thrombocyte Concentrate (TC)   | TC   | Platelet concentrate                 |
| 7  | Cryoprecipitate                | CRY  | Fibrinogen concentrate               |
| 8  | PRC Leukodepleted              | PRCL | Red cells with WBC removed           |

---

## Notes

- Components are used in blood requests (detail_permintaan_darah)
- Only active components can be used in new requests
- Component activation/deactivation affects availability without hard delete
- All operations are logged in system_access_logs

---

**Last Updated:** 2026-04-07
