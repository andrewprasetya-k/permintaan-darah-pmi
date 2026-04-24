# Quick Reference - API Endpoints Cheat Sheet

**Base URL:** `http://localhost:8080/api`

---

## 🔐 Authentication

```bash
# Admin Login
POST /auth/login/admin
{"username": "admin@example.com", "password": "pass123"}
→ Returns: {token, userId, role}

# Hospital Login
POST /auth/login/rumah-sakit
{"username": "rs_username", "password": "pass123"}
→ Returns: {token, userId, role}

# All requests after login:
Authorization: Bearer {token}

# WebSocket:
ws://localhost:8080/api/ws/connect?token={token}
```

---

## 👥 Admin Management (Superadmin Only)

```bash
POST   /admin                    # Create admin
GET    /admin                    # List admins
GET    /admin/:id                # Get admin
PUT    /admin/:id                # Update admin
DELETE /admin/:id                # Delete admin (soft)
PUT    /admin/restore/:id        # Restore admin
GET    /admin/me                 # Get own profile (any admin)
PUT    /admin/me                 # Update own profile
```

---

## 🏥 Hospital Management (Admin Only)

```bash
POST   /rumah-sakit              # Create hospital
GET    /rumah-sakit              # List hospitals
GET    /rumah-sakit/:id          # Get hospital
PUT    /rumah-sakit/:id          # Update hospital
DELETE /rumah-sakit/:id          # Delete (soft)
PUT    /rumah-sakit/restore/:id  # Restore hospital
GET    /rumah-sakit/me           # Get own (rumah sakit)
PUT    /rumah-sakit/me           # Update own (rumah sakit)
GET    /filter/rumah-sakit/      # Get distinct names (admin)
```

---

## 🩸 Blood Components (Admin Only)

```bash
POST   /komponen-darah           # Create component
GET    /komponen-darah           # List components
GET    /komponen-darah/:id       # Get component
PUT    /komponen-darah/:id       # Update component
DELETE /komponen-darah/:id       # Delete component
PUT    /komponen-darah/activate/:id      # Activate
PUT    /komponen-darah/deactivate/:id    # Deactivate
```

---

## 🩹 Blood Requests (Admin & Rumah Sakit)

```bash
# Create & List
POST   /permintaan-darah         # Create (admin/rumah sakit)
GET    /permintaan-darah         # List all (admin) / own+shared (rumah sakit)
GET    /permintaan-darah/:id     # Get detail

# Rumah Sakit (Own Requests)
GET    /permintaan-darah/my-requests        # List own
PUT    /permintaan-darah/my-requests/:id    # Update own

# Admin (All Requests)
PUT    /permintaan-darah/:id                # Update any
DELETE /permintaan-darah/:id                # Delete any
PUT    /permintaan-darah/update/:id         # Change status

# Status Workflow
dibuat → diproses → bisa_diambil → selesai
  ↓        ↓           ↓
dibatalkan (any point, needs reason)
```

---

## 📋 Request Details (Rumah Sakit Only)

```bash
POST   /detail-permintaan-darah  # Create detail
GET    /detail-permintaan-darah  # List details
GET    /detail-permintaan-darah/:id    # Get detail
PUT    /detail-permintaan-darah/:id    # Update detail
DELETE /detail-permintaan-darah/:id    # Delete detail
```

---

## 📊 Logs & Audit (Admin Only)

```bash
# Status Change Logs
GET    /status-logs              # List status changes

# System Access Logs (filtering)
GET    /system-logs              # List all
GET    /system-logs/:id          # Get specific
GET    /system-logs/user/:userId # By user
GET    /system-logs/action/:action        # By action
GET    /system-logs/table/:table          # By table
GET    /system-logs/target/:targetId      # By target

# Dashboard
GET    /dashboard/status-summary/:rsId    # Get stats
# :rsId = hospital ID or "all" (admin only)
```

---

## 📡 WebSocket

```bash
# Connect
ws://localhost:8080/api/ws/connect?token={JWT_TOKEN}

# Messages Received
{
  "type": "status_change" | "update_permintaan",
  "action": "UPDATE",
  "entityId": "PD...",
  "entityType": "permintaan_darah",
  "data": {...},
  "timestamp": "2026-04-07T..."
}

# Status Check
GET    /ws/status                # Get active connections count
```

---

## 📈 Query Parameters

```bash
# Pagination (most GET endpoints)
?limit=20&offset=0

# Filtering (permintaan_darah)
?status=dibuat&search=Budi&golDarah=O&startDate=2026-04-01&endDate=2026-04-30

# Filtering (logs)
?limit=50&offset=100
```

---

## 🛑 HTTP Status Codes

| Code | Meaning                              |
|------|--------------------------------------|
| 200  | ✅ OK - Success                      |
| 201  | ✅ Created - Resource created        |
| 400  | ❌ Bad Request - Invalid input       |
| 401  | ❌ Unauthorized - Invalid token      |
| 403  | ❌ Forbidden - No permission         |
| 404  | ❌ Not Found - Resource not exist    |
| 429  | ❌ Too Many Requests - Rate limited  |
| 500  | ❌ Server Error                      |

---

## ⚡ Rate Limits

| Endpoint     | Limit      | Window   |
|--------------|-----------|----------|
| Login        | 5 req     | 1 minute |
| API (protect)| 100 req   | 1 minute |

---

## 🔑 Status Values

- `dibuat` - Created
- `diproses` - Processing
- `bisa_diambil` - Ready
- `selesai` - Completed
- `dibatalkan` - Cancelled

---

## 👤 Roles

- `superadmin` - Super admin
- `admin` - Regular admin
- `rumah_sakit` - Hospital

---

## 🩸 Blood Type & Gender

**Types:** A, B, AB, O

**Rhesus:** +, -

**Gender:** L (Male), P (Female)

---

## 📝 Example Requests

### Login & Get Token
```bash
curl -X POST http://localhost:8080/api/auth/login/admin \
  -H "Content-Type: application/json" \
  -d '{"username":"admin@example.com","password":"pass123"}'

# Save token from response
TOKEN="eyJhbGciOiJIUzI1NiIs..."
```

### Create Blood Request
```bash
curl -X POST http://localhost:8080/api/permintaan-darah \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "rsId":"RS001",
    "namaPasien":"Budi",
    "tanggalLahir":"1990-05-15",
    "jenisKelamin":"L",
    "golonganDarah":"O",
    "details":[{"komId":1,"jumlahKantong":2}]
  }'
```

### List Blood Requests with Filter
```bash
curl -X GET "http://localhost:8080/api/permintaan-darah?status=dibuat&limit=10" \
  -H "Authorization: Bearer $TOKEN"
```

### Update Request Status
```bash
curl -X PUT http://localhost:8080/api/permintaan-darah/update/PD04071430001 \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"status":"diproses"}'
```

---

## 📚 Full Documentation

- **README.md** - Start here
- **GENERAL.md** - Authentication, enums, rate limiting
- **01_AUTH.md** - Login endpoints
- **02_ADMIN.md** - Admin management
- **03_RUMAH_SAKIT.md** - Hospital management
- **04_KOMPONEN_DARAH.md** - Blood components
- **05_PERMINTAAN_DARAH.md** - Blood requests
- **06_DETAIL_PERMINTAAN_DARAH.md** - Request details
- **07_LOGS.md** - Audit trails
- **08_DASHBOARD.md** - Dashboard/stats

---

**Last Updated:** 2026-04-07  
**Token Expiry:** 24 hours  
**Base URL:** http://localhost:8080/api
