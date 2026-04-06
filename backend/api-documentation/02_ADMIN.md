# API Documentation - Admin

**Base URL:** `http://localhost:8080/api`

**All endpoints require Authorization header:**

```
Authorization: Bearer {token}
```

---

## Create Admin

**POST** `/admin`

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
    "createdAt": "2026-04-06T08:33:32Z",
    "updatedAt": "2026-04-06T08:33:32Z"
  }
}
```

**Response (400 Bad Request):**

```json
{
  "success": false,
  "message": "Invalid input",
  "errors": ["email must be valid", "username is required"]
}
```

---

## Get All Admins

**GET** `/admin?limit=20&offset=0`

**Query Parameters:**

- `limit` (optional): default 20, max 100
- `offset` (optional): default 0

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
      "createdAt": "2026-04-01T10:00:00Z",
      "updatedAt": "2026-04-06T08:33:32Z"
    },
    {
      "adminId": "ADM-002",
      "adminUsername": "admin2",
      "adminName": "Admin Two",
      "adminEmail": "admin2@example.com",
      "adminRole": "admin",
      "createdAt": "2026-04-06T08:33:32Z",
      "updatedAt": "2026-04-06T08:33:32Z"
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

## Get Admin by ID

**GET** `/admin/{id}`

**Path Parameters:**

- `id` (required): Admin ID

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
    "createdAt": "2026-04-01T10:00:00Z",
    "updatedAt": "2026-04-06T08:33:32Z"
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

## Update Admin

**PUT** `/admin/{id}`

**Path Parameters:**

- `id` (required): Admin ID

**Request:**

```json
{
  "adminUsername": "admin_updated",
  "adminPassword": "newpassword123",
  "adminName": "Admin One Updated",
  "adminEmail": "admin_updated@example.com",
  "adminRole": "admin"
}
```

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
    "createdAt": "2026-04-01T10:00:00Z",
    "updatedAt": "2026-04-06T09:00:00Z"
  }
}
```

---

## Delete Admin (Soft Delete)

**DELETE** `/admin/{id}`

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

---

## Restore Admin

**PUT** `/admin/restore/{id}`

**Path Parameters:**

- `id` (required): Admin ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Admin restored successfully",
  "data": null
}
```

---

## Admin Role Values

- `superadmin` - Super admin
- `admin` - Regular admin

---

## Error Codes

- `400` - Bad Request (Invalid input)
- `401` - Unauthorized (Missing/invalid token)
- `403` - Forbidden (No permission)
- `404` - Not Found
- `500` - Internal Server Error
