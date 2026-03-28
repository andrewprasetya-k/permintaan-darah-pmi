package models

type RhesusEnum string

const (
	RhesusPlus  RhesusEnum = "+"
	RhesusMinus RhesusEnum = "-"
)

type BloodTypeEnum string

const (
	BloodTypeA  BloodTypeEnum = "A"
	BloodTypeB  BloodTypeEnum = "B"
	BloodTypeAB BloodTypeEnum = "AB"
	BloodTypeO  BloodTypeEnum = "O"
)

type GenderEnum string

const (
	GenderL GenderEnum = "L"
	GenderP GenderEnum = "P"
)

type PermintaanStatusEnum string

const (
	StatusDiterima     PermintaanStatusEnum = "diterima"
	StatusDiproses     PermintaanStatusEnum = "diproses"
	StatusBisaDiambil  PermintaanStatusEnum = "bisa_diambil"
	StatusSelesai      PermintaanStatusEnum = "selesai"
	StatusDibatalkan   PermintaanStatusEnum = "dibatalkan"
)

type AdminRoleEnum string

const (
	RoleAdmin      AdminRoleEnum = "admin"
	RoleSuperadmin AdminRoleEnum = "superadmin"
)