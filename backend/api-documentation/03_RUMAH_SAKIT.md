# API Documentation - Rumah Sakit (Hospital Management)

**Base URL:** `http://localhost:8080/api`

**All endpoints require Authorization header:**

```
Authorization: Bearer {token}
```

**Access Control:**
- `/rumah-sakit` (CRUD operations) - **Admin Only**
- `/rumah-sakit/me` (Self Profile) - **Rumah Sakit Only**

---

## Create Rumah Sakit (Admin Only)

**POST** `/rumah-sakit`

**Access:** Admin or Superadmin only

**Request:**

```json
{
  "nama": "Rumah Sakit Pusat",
  "nomorTelepon": "021-1234567",
  "alamat": "Jl. Merdeka No. 10, Jakarta",
  "email": "info@rspusat.com",
  "username": "rspusat",
  "password": "password123"
}
```

**Response (201 Created):**

```json
{
  "success": true,
  "message": "Hospital created successfully",
  "data": {
    "rumahSakitId": "RS001",
    "nama": "Rumah Sakit Pusat",
    "nomorTelepon": "021-1234567",
    "alamat": "Jl. Merdeka No. 10, Jakarta",
    "email": "info@rspusat.com",
    "username": "rspusat",
    "createdAt": "2026-04-07T10:30:45Z",
    "updatedAt": "2026-04-07T10:30:45Z"
  }
}
```

---

## Get All Rumah Sakit (Admin Only)

**GET** `/rumah-sakit?limit=20&offset=0`

**Access:** Admin or Superadmin only

**Query Parameters:**

- `limit` (optional): default 20, max 100
- `offset` (optional): default 0

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Hospitals retrieved successfully",
  "data": [
    {
      "rumahSakitId": "RS001",
      "nama": "Rumah Sakit Pusat",
      "nomorTelepon": "021-1234567",
      "alamat": "Jl. Merdeka No. 10, Jakarta",
      "email": "info@rspusat.com",
      "username": "rspusat",
      "createdAt": "2026-04-01T10:00:00Z",
      "updatedAt": "2026-04-07T10:30:45Z"
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

## Get Rumah Sakit by ID (Admin Only)

**GET** `/rumah-sakit/{id}`

**Access:** Admin or Superadmin only

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Hospital retrieved successfully",
  "data": {
    "rumahSakitId": "RS001",
    "nama": "Rumah Sakit Pusat",
    "nomorTelepon": "021-1234567",
    "alamat": "Jl. Merdeka No. 10, Jakarta",
    "email": "info@rspusat.com",
    "username": "rspusat",
    "createdAt": "2026-04-01T10:00:00Z",
    "updatedAt": "2026-04-07T10:30:45Z"
  }
}
```

---

## Update Rumah Sakit (Admin Only)

**PUT** `/rumah-sakit/{id}`

**Access:** Admin or Superadmin only

**Request:**

```json
{
  "nama": "Rumah Sakit Pusat - Updated",
  "nomorTelepon": "021-9876543",
  "alamat": "Jl. Merdeka No. 10, Jakarta Pusat",
  "email": "updated@rspusat.com",
  "password": "newpassword123"
}
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Hospital updated successfully",
  "data": {
    "rumahSakitId": "RS001",
    "nama": "Rumah Sakit Pusat - Updated",
    "nomorTelepon": "021-9876543",
    "alamat": "Jl. Merdeka No. 10, Jakarta Pusat",
    "email": "updated@rspusat.com",
    "username": "rspusat",
    "createdAt": "2026-04-01T10:00:00Z",
    "updatedAt": "2026-04-07T11:00:00Z"
  }
}
```

---

## Delete Rumah Sakit (Admin Only)

**DELETE** `/rumah-sakit/{id}`

**Access:** Admin or Superadmin only

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Hospital deleted successfully",
  "data": null
}
```

---

## Restore Rumah Sakit (Admin Only)

**PUT** `/rumah-sakit/restore/{id}`

**Access:** Admin or Superadmin only

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Hospital restored successfully",
  "data": null
}
```

---

## Get Own Profile (Rumah Sakit Only)

**GET** `/rumah-sakit/me`

**Access:** Only rumah_sakit role

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Profile retrieved successfully",
  "data": {
    "rumahSakitId": "RS001",
    "nama": "Rumah Sakit Pusat",
    "nomorTelepon": "021-1234567",
    "alamat": "Jl. Merdeka No. 10, Jakarta",
    "email": "info@rspusat.com",
    "username": "rspusat",
    "createdAt": "2026-04-01T10:00:00Z",
    "updatedAt": "2026-04-07T10:30:45Z"
  }
}
```

---

## Update Own Profile (Rumah Sakit Only)

**PUT** `/rumah-sakit/me`

**Access:** Only rumah_sakit role

**Request:**

```json
{
  "nama": "Rumah Sakit Pusat",
  "nomorTelepon": "021-9876543",
  "alamat": "Jl. Merdeka No. 10, Jakarta Selatan",
  "email": "newemail@rspusat.com",
  "password": "newpassword123"
}
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Profile updated successfully",
  "data": {
    "rumahSakitId": "RS001",
    "nama": "Rumah Sakit Pusat",
    "nomorTelepon": "021-9876543",
    "alamat": "Jl. Merdeka No. 10, Jakarta Selatan",
    "email": "newemail@rspusat.com",
    "username": "rspusat",
    "createdAt": "2026-04-01T10:00:00Z",
    "updatedAt": "2026-04-07T11:30:00Z"
  }
}
```

---

## Get Distinct Rumah Sakit Names (Admin Only)

**GET** `/filter/rumah-sakit/`

**Access:** Admin or Superadmin only

**Purpose:** Get list of all unique hospital names for filtering

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Hospital names retrieved successfully",
  "data": [
    "Rumah Sakit Pusat",
    "Rumah Sakit Cabang",
    "Rumah Sakit Satelit"
  ]
}
```

---

**Last Updated:** 2026-04-07
