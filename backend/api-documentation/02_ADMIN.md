# API Documentation - Admin

**Base URL:** `http://localhost:8080/api`

**All endpoints require Authorization header:**

```
Authorization: Bearer {token}
```

**Access Control:**
- `/admin` (CRUD operations) - **Superadmin Only**
- `/admin/me` (Self Profile) - **Any Admin**

---

## Create Admin (Superadmin Only)

**POST** `/admin`

**Access:** Superadmin only

**Request:**

```json
{
  "adminUsername": "admin2",
  "adminPassword": "password123",
  "adminName": "Admin Two",
  "adminEmail": "admin2@example.com",
  "adminRole": "admin"
}
```

**Response (201 Created):**

```json
{
  "success": true,
  "message": "Admin created successfully",
  "data": {
    "adminId": "ADM-002",
    "adminUsername": "admin2",
    "adminName": "Admin Two",
    "adminEmail": "admin2@example.com",
    "adminRole": "admin",
    "isDeleted": false,
    "createdAt": "2026-04-07T10:30:45Z",
    "updatedAt": "2026-04-07T10:30:45Z"
  }
}
```

**Response (400 Bad Request):**

```json
{
  "success": false,
  "message": "Invalid input"
}
```

**Response (403 Forbidden):**

```json
{
  "error": "Forbidden: Super Admins only"
}
```

---

## Get All Admins (Superadmin Only)

**GET** `/admin?limit=20&offset=0&status=active`

**Access:** Superadmin only

**Query Parameters:**

- `limit` (optional): default 20
- `offset` (optional): default 0
- `status` (optional): `active` | `deleted` | `all` (default `active`)

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Admins retrieved successfully",
  "data": [
    {
      "adminId": "ADM-001",
      "adminUsername": "admin",
      "adminName": "Admin One",
      "adminEmail": "admin@example.com",
      "adminRole": "superadmin",
      "isDeleted": false,
      "createdAt": "2026-04-01T10:00:00Z",
      "updatedAt": "2026-04-07T10:30:45Z"
    },
    {
      "adminId": "ADM-002",
      "adminUsername": "admin2",
      "adminName": "Admin Two",
      "adminEmail": "admin2@example.com",
      "adminRole": "admin",
      "isDeleted": false,
      "createdAt": "2026-04-07T10:30:45Z",
      "updatedAt": "2026-04-07T10:30:45Z"
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

## Get Admin by ID (Superadmin Only)

**GET** `/admin/{id}`

**Access:** Superadmin only

**Path Parameters:**

- `id` (required): Admin ID (e.g., ADM-001)

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Admin retrieved successfully",
  "data": {
    "adminId": "ADM-001",
    "adminUsername": "admin",
    "adminName": "Admin One",
    "adminEmail": "admin@example.com",
    "adminRole": "superadmin",
    "isDeleted": false,
    "createdAt": "2026-04-01T10:00:00Z",
    "updatedAt": "2026-04-07T10:30:45Z"
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

## Update Admin (Superadmin Only)

**PUT** `/admin/{id}`

**Access:** Superadmin only

**Path Parameters:**

- `id` (required): Admin ID

**Request:**

```json
{
  "adminUsername": "admin_updated",
  "adminName": "Admin One Updated",
  "adminEmail": "admin_updated@example.com",
  "adminRole": "admin",
  "adminPassword": "newpassword123"
}
```

**Note:** `adminPassword` bersifat opsional saat update. Jika field ini tidak dikirim atau string kosong, password lama tetap dipakai.

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Admin updated successfully",
  "data": {
    "adminId": "ADM-001",
    "adminUsername": "admin_updated",
    "adminName": "Admin One Updated",
    "adminEmail": "admin_updated@example.com",
    "adminRole": "admin",
    "isDeleted": false,
    "createdAt": "2026-04-01T10:00:00Z",
    "updatedAt": "2026-04-07T11:00:00Z"
  }
}
```

---

## Delete Admin / Soft Delete (Superadmin Only)

**DELETE** `/admin/{id}`

**Access:** Superadmin only

**Path Parameters:**

- `id` (required): Admin ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Admin deleted successfully",
  "data": null
}
```

**Note:** This is a soft delete. Data is marked as deleted but preserved in database for audit trail.

---

## Restore Admin (Superadmin Only)

**PUT** `/admin/restore/{id}`

**Access:** Superadmin only

**Path Parameters:**

- `id` (required): Admin ID (must be previously deleted)

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Admin restored successfully",
  "data": {
    "adminId": "ADM-001",
    "adminUsername": "admin",
    "adminName": "Admin One",
    "adminEmail": "admin@example.com",
    "adminRole": "superadmin",
    "isDeleted": false,
    "createdAt": "2026-04-01T10:00:00Z",
    "updatedAt": "2026-04-07T11:40:00Z"
  }
}
```

---

## Get Own Profile (Any Admin)

**GET** `/admin/me`

**Access:** Any authenticated admin (superadmin or admin role)

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Admin retrieved successfully",
  "data": {
    "adminId": "ADM-001",
    "adminUsername": "admin",
    "adminName": "Admin One",
    "adminEmail": "admin@example.com",
    "adminRole": "superadmin",
    "isDeleted": false,
    "createdAt": "2026-04-01T10:00:00Z",
    "updatedAt": "2026-04-07T10:30:45Z"
  }
}
```

---

## Update Own Profile (Any Admin)

**PUT** `/admin/me`

**Access:** Any authenticated admin

**Request:**

```json
{
  "adminUsername": "admin",
  "adminName": "Admin Updated Name",
  "adminEmail": "newemail@example.com",
  "adminRole": "superadmin",
  "adminPassword": "newpassword123"
}
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Admin profile updated successfully",
  "data": {
    "adminId": "ADM-001",
    "adminUsername": "admin",
    "adminName": "Admin Updated Name",
    "adminEmail": "newemail@example.com",
    "adminRole": "superadmin",
    "isDeleted": false,
    "createdAt": "2026-04-01T10:00:00Z",
    "updatedAt": "2026-04-07T11:30:00Z"
  }
}
```

**Note:** Endpoint ini saat ini memakai DTO update admin yang sama dengan `PUT /admin/{id}`, jadi request tetap harus mengirim `adminUsername`, `adminName`, `adminEmail`, dan `adminRole`. `adminPassword` opsional.

---

## Admin Role Values

- `superadmin` - Super administrator (can manage other admins)
- `admin` - Regular administrator (cannot manage other admins)

---

## Notes

- All passwords are hashed with bcrypt (never stored in plain text)
- Admins can only update their own profile (except superadmin can update any admin)
- Soft delete preserves all data for audit trail
- All operations are logged in system_access_logs

---

**Last Updated:** 2026-04-27
