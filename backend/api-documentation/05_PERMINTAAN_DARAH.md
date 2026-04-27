# API Documentation - Permintaan Darah (Blood Requests)

**Base URL:** `http://localhost:8080/api`

**All endpoints require Authorization header:**

```
Authorization: Bearer {token}
```

**Access Control:**
- Create/List/Detail/Update/Status update - **Admin or Rumah Sakit**
- `DELETE /permintaan-darah/{id}` - **Admin or Superadmin only** (checked inside controller)
- `GET /permintaan-darah/my-requests` and `PUT /permintaan-darah/my-requests/{id}` - **Rumah Sakit only** with ownership check

---

## Create Blood Request

**POST** `/permintaan-darah`

**Access:** Admin or Rumah Sakit

**Request:**

```json
{
  "rumahSakitId": "RS001",
  "namaPasien": "Budi Santoso",
  "noRM": "RM123",
  "tempatLahir": "Jakarta",
  "tanggalLahir": "1990-05-15T00:00:00Z",
  "jenisKelamin": "L",
  "golonganDarah": "O",
  "rhesusDarah": "+",
  "hemoglobin": 12.5,
  "ruangBagianKelas": "ICU",
  "pernahTransfusi": false,
  "indikasiTransfusi": "Transfusi rutin",
  "pernahHamil": null,
  "status": "dibuat",
  "tanggalPermintaan": "2026-04-10T00:00:00Z",
  "details": [
    {
      "komponenDarahId": 1,
      "golonganDarah": "O",
      "rhesusDarah": "+",
      "jumlahKantong": 2
    }
  ]
}
```

**Response (201 Created):**

```json
{
  "success": true,
  "message": "Created successfully",
  "data": {
    "permintaanDarahId": "PD04071430001",
    "rumahSakitId": "RS001",
    "namaPasien": "Budi Santoso",
    "noRM": "RM123",
    "tempatLahir": "Jakarta",
    "tanggalLahir": "1990-05-15T00:00:00Z",
    "jenisKelamin": "L",
    "golonganDarah": "O",
    "rhesusDarah": "+",
    "hemoglobin": 12.5,
    "status": "dibuat",
    "tanggalPermintaan": "2026-04-10T00:00:00Z",
    "createdAt": "2026-04-07T14:30:45Z",
    "updatedAt": "2026-04-07T14:30:45Z",
    "detailPermintaanDarah": [
      {
        "detailId": 1,
        "permintaanDarahId": "PD04071430001",
        "komponenNama": "Whole Blood",
        "golonganDarah": "O",
        "rhesusDarah": "+",
        "jumlahKantong": 2,
        "tanggalDiperlukan": "0001-01-01T00:00:00Z",
        "createdAt": "2026-04-07T14:30:45Z"
      }
    ]
  }
}
```

---

## Get All Blood Requests

**GET** `/permintaan-darah?limit=20&offset=0&status=dibuat&search=Budi`

**Access:** Admin or Rumah Sakit

**Query Parameters:**

- `limit` (optional): default 20
- `offset` (optional): default 0
- `status` (optional): dibuat | diproses | bisa_diambil | selesai | dibatalkan
- `search` (optional): matches request ID, patient name, medical record number, blood type, rhesus, or combined blood type+rhesus
- `rsID` (optional): filter by rumah sakit ID
- `golDarah` (optional): A | B | AB | O
- `startDate` (optional): `YYYY-MM-DD`, compared against `tanggalPermintaan`
- `endDate` (optional): `YYYY-MM-DD`, compared against `tanggalPermintaan`

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [...],
  "pagination": {
    "total": 50,
    "page": 1,
    "limit": 20,
    "offset": 0
  }
}
```

---

## Get Own Blood Requests (Rumah Sakit)

**GET** `/permintaan-darah/my-requests?limit=20&offset=0`

**Access:** Rumah Sakit only

**Query Parameters:**

- `limit` (optional): default 20
- `offset` (optional): default 0

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [...],
  "pagination": {
    "total": 15,
    "page": 1,
    "limit": 20,
    "offset": 0
  }
}
```

---

## Get Blood Request by ID

**GET** `/permintaan-darah/{id}`

**Access:** Admin or Rumah Sakit. Current controller does not enforce ownership for this endpoint.

**Path Parameters:**

- `id` (required): Request ID (e.g., PD04071430001)

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {
    "permintaanDarahId": "PD04071430001",
    "rumahSakitId": "RS001",
    "namaPasien": "Budi Santoso",
    "status": "diproses",
    "createdAt": "2026-04-07T14:30:45Z",
    "updatedAt": "2026-04-07T15:00:00Z",
    "detailPermintaanDarah": [...]
  }
}
```

---

## Update Blood Request

**PUT** `/permintaan-darah/{id}`

**Access:** Admin or Rumah Sakit. Current controller does not enforce ownership for this endpoint; Rumah Sakit-specific ownership check is available on `/permintaan-darah/my-requests/{id}`.

**Request:**

```json
{
  "rumahSakitId": "RS001",
  "namaPasien": "Budi Santoso",
  "noRM": "RM123",
  "tempatLahir": "Jakarta",
  "tanggalLahir": "1990-05-15T00:00:00Z",
  "jenisKelamin": "L",
  "golonganDarah": "O",
  "rhesusDarah": "+",
  "hemoglobin": 13.0,
  "ruangBagianKelas": "VVIP",
  "pernahTransfusi": false,
  "indikasiTransfusi": "Transfusi rutin",
  "pernahHamil": null,
  "status": "diproses",
  "tanggalPermintaan": "2026-04-10T00:00:00Z"
}
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {...}
}
```

---

## Update Own Blood Request (Rumah Sakit)

**PUT** `/permintaan-darah/my-requests/{id}`

**Access:** Rumah Sakit only (own requests)

**Request:**

```json
{
  "rumahSakitId": "RS001",
  "namaPasien": "Budi Santoso",
  "noRM": "RM123",
  "tempatLahir": "Jakarta",
  "tanggalLahir": "1990-05-15T00:00:00Z",
  "jenisKelamin": "L",
  "golonganDarah": "O",
  "rhesusDarah": "+",
  "hemoglobin": 13.0,
  "ruangBagianKelas": "ICU",
  "pernahTransfusi": false,
  "indikasiTransfusi": "Transfusi emergensi",
  "pernahHamil": null,
  "status": "dibuat",
  "tanggalPermintaan": "2026-04-10T00:00:00Z"
}
```

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Request updated successfully",
  "data": {...}
}
```

---

## Update Blood Request Status

**PUT** `/permintaan-darah/update/{id}`

**Access:** Admin or Rumah Sakit. For Rumah Sakit users, the service checks that the request belongs to their hospital.

**Request:**

```json
{
  "status": "diproses",
  "reason": null
}
```

**Status Transition Handling:**

- Backend currently blocks changes only when current status is `selesai` or `dibatalkan`.
- Backend requires `reason` when new status is `dibatalkan`.
- Frontend may enforce a stricter workflow (`dibuat` → `diproses` → `bisa_diambil` → `selesai`), but that transition matrix is not enforced in the service.

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Status updated successfully",
  "data": {
    "permintaanDarahId": "PD04071430001",
    "status": "diproses",
    "updatedAt": "2026-04-07T15:15:00Z"
  }
}
```

**Response (400 Bad Request) - Invalid Status:**

```json
{
  "success": false,
  "message": "cannot update status of a completed or cancelled request"
}
```

---

## Delete Blood Request

**DELETE** `/permintaan-darah/{id}`

**Access:** Admin

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": null
}
```

---

## Blood Request Status Values

- `dibuat` - Newly created
- `diproses` - Being processed
- `bisa_diambil` - Ready to pick up
- `selesai` - Completed
- `dibatalkan` - Cancelled

---

## Blood Type & Rhesus Values

**Blood Types:** A, B, AB, O

**Rhesus:** +, -

**Gender:** L (Male), P (Female)

---

## Notes

- Request ID format: PD + DDMMYYHHMM + 3-digit sequence
- Patient data is flattened (duplicated per request, not normalized)
- Status changes are logged in status_logs for audit trail
- Real-time WebSocket notifications on status changes
- Requests marked selesai or dibatalkan cannot be modified
- Request dates should be sent as RFC3339 timestamps because the DTO uses `time.Time`.
- `tanggalDiperlukan` exists in the detail response DTO, but the current model/create flow does not persist it, so it may appear as the zero timestamp.

---

**Last Updated:** 2026-04-27
