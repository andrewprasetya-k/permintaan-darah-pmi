# API Documentation - Authentication

**Base URL:** `http://localhost:8080/api`

---

## Admin Login

**POST** `/auth/login/admin`

**Request:**

```json
{
  "username": "admin",
  "password": "password123"
}
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "userId": "ADM-001",
    "username": "admin",
    "role": "superadmin",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**Response (401 Unauthorized):**

```json
{
  "success": false,
  "message": "Invalid credentials",
  "errors": []
}
```

---

## Rumah Sakit Login

**POST** `/auth/login/rumah-sakit`

**Request:**

```json
{
  "username": "rs_username",
  "password": "password123"
}
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "userId": "RS001",
    "username": "rs_username",
    "role": "rumah_sakit",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**Response (401 Unauthorized):**

```json
{
  "success": false,
  "message": "Invalid credentials",
  "errors": []
}
```

---

## Notes

- **Token Expiry:** 24 hours
- **Token Location:** Authorization header with "Bearer" prefix
- **Format:** `Authorization: Bearer {token}`
- Login endpoints return token yang harus disimpan di client
- Semua protected endpoints memerlukan valid token di header
