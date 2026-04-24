# API Documentation - Complete Reference

**Project:** Permintaan Darah PMI - Blood Donation Request Management System

**Base URL:** `http://localhost:8080/api`

**Version:** 1.0.0

**Last Updated:** 2026-04-07

---

## 📚 Documentation Index

### Getting Started
1. **[GENERAL.md](GENERAL.md)** - General guidelines, authentication, rate limiting, enums
2. **[01_AUTH.md](01_AUTH.md)** - Authentication (login endpoints)

### API Endpoints by Module

#### User Management
3. **[02_ADMIN.md](02_ADMIN.md)** - Admin CRUD + self profile (Superadmin/Admin)
4. **[03_RUMAH_SAKIT.md](03_RUMAH_SAKIT.md)** - Hospital CRUD + self profile (Admin/Rumah Sakit)

#### Master Data
5. **[04_KOMPONEN_DARAH.md](04_KOMPONEN_DARAH.md)** - Blood component management (Admin only)

#### Core Business
6. **[05_PERMINTAAN_DARAH.md](05_PERMINTAAN_DARAH.md)** - Blood request CRUD + status management (Admin/Rumah Sakit)

#### Reporting & Monitoring
7. **[07_LOGS.md](07_LOGS.md)** - Audit logs (Admin only)
8. **[08_DASHBOARD.md](08_DASHBOARD.md)** - Dashboard analytics (Admin/Rumah Sakit)

---

## 🔐 Authentication & Authorization

### Login

```bash
# Admin Login
POST /auth/login/admin
{
  "username": "admin@example.com",
  "password": "password123"
}

# Rumah Sakit Login
POST /auth/login/rumah-sakit
{
  "username": "rs_username",
  "password": "password123"
}
```

### Token Usage

All protected endpoints require Authorization header:

```
Authorization: Bearer {token}
```

Token validity: **24 hours**

### WebSocket Connection

For real-time updates, connect via WebSocket with token in query parameter:

```
ws://localhost:8080/api/ws/connect?token=JWT_TOKEN
```

See **GENERAL.md** for WebSocket details.

---

## 👥 Access Control

### Superadmin
- ✅ Create/manage other admins
- ✅ Manage all hospitals
- ✅ Manage blood components
- ✅ View all blood requests
- ✅ View all logs
- ✅ Access all endpoints

### Admin
- ✅ Cannot create other admins
- ✅ Can manage hospitals & components
- ✅ View all blood requests
- ✅ Update request status
- ✅ View audit logs
- ❌ Cannot access admin management

### Rumah Sakit (Hospital)
- ✅ Create/manage own blood requests
- ✅ View own requests + shared requests
- ✅ Update own request status
- ❌ Cannot manage other hospitals
- ❌ Cannot access admin/component management
- ❌ Cannot view logs

---

## 📊 Main Workflows

### 1. Blood Request Workflow

```
Rumah Sakit creates request (status: dibuat)
                    ↓
Admin reviews & updates (status: diproses)
                    ↓
Admin marks ready (status: bisa_diambil)
                    ↓
Hospital picks up (status: selesai)
```

OR at any point: Rumah Sakit or Admin can mark `dibatalkan` with reason.

### 2. Real-Time Updates

- Status changes broadcast to all connected WebSocket clients
- All actions logged in system_access_logs
- Status changes also logged in status_logs for audit trail

---

## 📋 Common Query Parameters

Most list endpoints support:

- `limit` - Results per page (default: 20, max: 100)
- `offset` - Starting position (default: 0)

Example:
```
GET /api/permintaan-darah?limit=50&offset=100
```

---

## ⚡ Rate Limiting

Different endpoints have different limits:

| Endpoint Type       | Limit           | Window    |
|-------------------|-----------------|-----------|
| Login              | 5 requests      | 1 minute  |
| Protected API      | 100 requests    | 1 minute  |
| Strict endpoints   | 30 requests     | 1 minute  |

Response when limit exceeded:

```
HTTP/1.1 429 Too Many Requests
Retry-After: 45
```

---

## 📅 Standard Response Format

### Success Response

```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... },
  "pagination": {
    "total": 100,
    "page": 1,
    "limit": 20,
    "offset": 0
  }
}
```

### Error Response

```json
{
  "success": false,
  "message": "Error description"
}
```

### HTTP Status Codes

- `200` - Success
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized
- `403` - Forbidden
- `404` - Not Found
- `429` - Too Many Requests (rate limit)
- `500` - Server Error

---

## 🔄 Enum Values

### Status Values

- `dibuat` - Newly created
- `diproses` - Being processed
- `bisa_diambil` - Ready to pick up
- `selesai` - Completed
- `dibatalkan` - Cancelled

### Blood Type

- `A`, `B`, `AB`, `O`

### Rhesus Factor

- `+` (positive)
- `-` (negative)

### Gender

- `L` - Laki-laki (Male)
- `P` - Perempuan (Female)

### Admin Role

- `superadmin` - Super administrator
- `admin` - Regular administrator

### Log Actions

- `CREATE` - Create
- `UPDATE` - Update
- `DELETE` - Delete
- `RESTORE` - Restore
- `LOGIN` - Login
- `LOGIN_FAILED` - Failed login

---

## 🛠️ Common Use Cases

### For Admin Dashboard
- GET `/dashboard/status-summary/all` - Overview seluruh rumah sakit
- GET `/system-logs` - Audit trail
- GET `/permintaan-darah` - All requests
- PUT `/permintaan-darah/update/{id}` - Update status

### For Hospital Portal
- POST `/permintaan-darah` - Create request
- GET `/permintaan-darah/my-requests` - Own requests
- PUT `/permintaan-darah/my-requests/{id}` - Update own request
- GET `/dashboard/status-summary/{hospitalId}` - Own dashboard (rumah sakit hanya akan menerima data miliknya sendiri)

### For Monitoring (WebSocket)
```javascript
const ws = new WebSocket('ws://localhost:8080/api/ws/connect?token=JWT_TOKEN');
ws.onmessage = (event) => {
  const message = JSON.parse(event.data);
  console.log('Update:', message.type, message.data);
};
```

---

## 📞 Support & Questions

For issues or questions:
1. Check relevant documentation file
2. Review GENERAL.md for common issues
3. Check HTTP status codes
4. Verify authentication token is valid
5. Check rate limiting headers if getting 429 errors

---

## 🔄 API Version History

| Version | Date       | Changes           |
|---------|------------|-------------------|
| 1.0.0   | 2026-04-07 | Initial release   |

---

**Maintained by:** Development Team  
**Last Updated:** 2026-04-07  
**Status:** ✅ Production Ready
