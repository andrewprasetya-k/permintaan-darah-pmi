package utils

import (
	"fmt"
	"time"
)

// ValidatePermintaanDarahInput validates permintaan darah input
func ValidatePermintaanDarahInput(hemoglobin *float64, tglLahir time.Time, tglPermintaan time.Time) error {
	now := time.Now()

	// Validate hemoglobin
	if hemoglobin != nil && *hemoglobin <= 0 {
		return fmt.Errorf("hemoglobin tidak valid")
	}
	if hemoglobin != nil && *hemoglobin > 20 {
		return fmt.Errorf("hemoglobin tidak valid")
	}

	// Validate tanggal lahir - cannot be in future
	if tglLahir.After(now) {
		return fmt.Errorf("tanggal lahir tidak bisa di masa depan")
	}

	//validate tanggal permintaan - cannot be in the past
	if tglPermintaan.Before(now) {
		return fmt.Errorf("tanggal permintaan tidak bisa di masa lalu")
	}

	// Validate age - cannot be more than 150 years old
	age := now.Year() - tglLahir.Year()
		if age < 0 || age > 150 {
			return fmt.Errorf("invalid tanggal lahir")
		}
	return nil
}

// ValidateDetailPermintaanDarahInput validates detail permintaan input
func ValidateDetailPermintaanDarahInput(jmlKantong int) error {
	if jmlKantong <= 0 {
		return fmt.Errorf("jumlah kantong harus lebih dari 0")
	}

	return nil
}

// ValidatePhoneNumber validates phone number format (Indonesia)
func ValidatePhoneNumber(phone string) error {
	if len(phone) < 10 {
		return fmt.Errorf("nomor telepon harus terdiri dari minimal 10 digit")
	}

	if len(phone) > 15 {
		return fmt.Errorf("nomor telepon harus terdiri dari maksimal 15 digit")
	}

	return nil
}
