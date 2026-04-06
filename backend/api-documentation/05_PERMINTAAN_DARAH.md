# API Documentation - Permintaan Darah

**Base URL:** `http://localhost:8080/api`

**All endpoints require Authorization header:**

```
Authorization: Bearer {token}
```

---

## Create Permintaan Darah

**POST** `/permintaan-darah`

**Request:**

```json
{
  "rumahSakitId": "RS001",
  "namaPasien": "John Doe",
  "noRM": "RM-001",
  "tempatLahir": "Jakarta",
  "tanggalLahir": "1990-01-15",
  "jenisKelamin": "L",
  "golonganDarah": "O",
  "rhesusDarah": "+",
  "hemoglobin": 12.5,
  "ruangBagianKelas": "ICU",
  "pernahTransfusi": false,
  "indikasiTransfusi": "Anemia",
  "pernahHamil": null,
  "status": "dibuat",
  "tanggalPermintaan": "2026-04-06T08:33:32Z"
}
```

**Response (201 Created):**

```json
{
  "success": true,
  "message": "Created successfully",
  "data": {
    "permintaanDarahId": "PD0406082826001",
    "rumahSakitId": "RS001",
    "namaPasien": "John Doe",
    "noRM": "RM-001",
    "tempatLahir": "Jakarta",
    "tanggalLahir": "1990-01-15T00:00:00Z",
    "jenisKelamin": "L",
    "golonganDarah": "O",
    "rhesusDarah": "+",
    "hemoglobin": 12.5,
    "ruangBagianKelas": "ICU",
    "pernahTransfusi": false,
    "indikasiTransfusi": "Anemia",
    "pernahHamil": null,
    "status": "dibuat",
    "cancelReason": null,
    "tanggalPermintaan": "2026-04-06T08:33:32Z",
    "createdAt": "2026-04-06T08:33:32Z",
    "updatedAt": "2026-04-06T08:33:32Z"
  }
}
```

---

## Get All Permintaan Darah

**GET** `/permintaan-darah?limit=20&offset=0&status=dibuat&rsID=RS001&gender=L&golDarah=O`

**Query Parameters:**

- `limit` (optional): default 20, max 100
- `offset` (optional): default 0
- `status` (optional): dibuat, diproses, bisa_diambil, selesai, dibatalkan
- `rsID` (optional): filter by rumah sakit
- `gender` (optional): L, P
- `golDarah` (optional): A, B, AB, O
- `startDate` (optional): ISO 8601 date format
- `endDate` (optional): ISO 8601 date format

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [
    {
      "permintaanDarahId": "PD0406082826001",
      "rumahSakitId": "RS001",
      "namaPasien": "John Doe",
      "noRM": "RM-001",
      "tempatLahir": "Jakarta",
      "tanggalLahir": "1990-01-15T00:00:00Z",
      "jenisKelamin": "L",
      "golonganDarah": "O",
      "rhesusDarah": "+",
      "hemoglobin": 12.5,
      "ruangBagianKelas": "ICU",
      "pernahTransfusi": false,
      "indikasiTransfusi": "Anemia",
      "status": "dibuat",
      "tanggalPermintaan": "2026-04-06T08:33:32Z",
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

## Get Permintaan Darah by ID

**GET** `/permintaan-darah/{id}`

**Path Parameters:**

- `id` (required): Permintaan Darah ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": {
    "permintaanDarahId": "PD0406082826001",
    "rumahSakitId": "RS001",
    "namaPasien": "John Doe",
    "noRM": "RM-001",
    "tempatLahir": "Jakarta",
    "tanggalLahir": "1990-01-15T00:00:00Z",
    "jenisKelamin": "L",
    "golonganDarah": "O",
    "rhesusDarah": "+",
    "hemoglobin": 12.5,
    "ruangBagianKelas": "ICU",
    "pernahTransfusi": false,
    "indikasiTransfusi": "Anemia",
    "status": "dibuat",
    "tanggalPermintaan": "2026-04-06T08:33:32Z",
    "createdAt": "2026-04-06T08:33:32Z",
    "updatedAt": "2026-04-06T08:33:32Z",
    "detailPermintaanDarah": [
      {
        "dpdId": 1,
        "dpdPdId": "PD0406082826001",
        "komponenNama": "Packed Red Cell (PRC)",
        "golonganDarah": "O",
        "rhesusDarah": "+",
        "jmlKantong": 2,
        "tglDiperlukan": "2026-04-08T10:00:00Z",
        "createdAt": "2026-04-06T08:33:32Z"
      }
    ]
  }
}
```

---

## Get Permintaan Darah by Rumah Sakit ID

**GET** `/permintaan-darah/rumah-sakit/{rsId}?limit=20&offset=0`

**Path Parameters:**

- `rsId` (required): Rumah Sakit ID

**Query Parameters:**

- `limit` (optional): default 20
- `offset` (optional): default 0

**Response (200 OK):** Same as GetAll

---

## Update Permintaan Darah

**PUT** `/permintaan-darah/{id}`

**Path Parameters:**

- `id` (required): Permintaan Darah ID

**Request:**

```json
{
  "rumahSakitId": "RS001",
  "namaPasien": "John Doe Updated",
  "noRM": "RM-001",
  "tempatLahir": "Jakarta",
  "tanggalLahir": "1990-01-15",
  "jenisKelamin": "L",
  "golonganDarah": "O",
  "rhesusDarah": "+",
  "hemoglobin": 13.0,
  "ruangBagianKelas": "ICCU",
  "pernahTransfusi": false,
  "indikasiTransfusi": "Anemia",
  "status": "diproses",
  "tanggalPermintaan": "2026-04-06T08:33:32Z"
}
```

**Response (200 OK):** Same as Create

---

## Update Permintaan Darah Status

**PUT** `/permintaan-darah/update/{id}`

**Path Parameters:**

- `id` (required): Permintaan Darah ID

**Request:**

```json
{
  "status": "bisa_diambil",
  "cancelReason": null
}
```

**Response (200 OK):** Same as Create

---

## Delete Permintaan Darah (Soft Delete)

**DELETE** `/permintaan-darah/{id}`

**Path Parameters:**

- `id` (required): Permintaan Darah ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": null
}
```

---

## Restore Permintaan Darah

**PUT** `/permintaan-darah/restore/{id}`

**Path Parameters:**

- `id` (required): Permintaan Darah ID

**Response (200 OK):**

```json
{
  "success": true,
  "message": "Operation successful",
  "data": null
}
```

---

## Status Values

- `dibuat` - Newly created
- `diproses` - Being processed
- `bisa_diambil` - Ready to pick up
- `selesai` - Completed
- `dibatalkan` - Cancelled

## Jenis Kelamin Values

- `L` - Laki-laki (Male)
- `P` - Perempuan (Female)

## Golongan Darah Values

- `A`
- `B`
- `AB`
- `O`

## Rhesus Values

- `+` - Positive
- `-` - Negative

---

## Error Codes

- `400` - Bad Request (Invalid input)
- `401` - Unauthorized (Missing/invalid token)
- `403` - Forbidden (No permission)
- `404` - Not Found
- `500` - Internal Server Error
