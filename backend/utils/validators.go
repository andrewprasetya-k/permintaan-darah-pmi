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
		return fmt.Errorf("hemoglobin must be greater than 0")
	}
	if hemoglobin != nil && *hemoglobin > 20 {
		return fmt.Errorf("hemoglobin value is unrealistic (max 20)")
	}

	// Validate tanggal lahir - cannot be in future
	if tglLahir.After(now) {
		return fmt.Errorf("tanggal lahir cannot be in the future")
	}

	//validate tanggal permintaan - cannot be in the past
	if tglPermintaan.Before(now) {
		return fmt.Errorf("tanggal permintaan cannot be in the past")
	}

	// Validate age - cannot be more than 150 years old
	age := now.Year() - tglLahir.Year()
		if age < 0 || age > 150 {
			return fmt.Errorf("invalid birth date - patient age must be between 0 and 150")
		}
	return nil
}

// ValidateDetailPermintaanDarahInput validates detail permintaan input
func ValidateDetailPermintaanDarahInput(jmlKantong int) error {
	if jmlKantong <= 0 {
		return fmt.Errorf("jumlah kantong must be greater than 0")
	}

	return nil
}

// ValidatePhoneNumber validates phone number format (Indonesia)
func ValidatePhoneNumber(phone string) error {
	if len(phone) < 10 {
		return fmt.Errorf("phone number must be at least 10 digits")
	}

	if len(phone) > 15 {
		return fmt.Errorf("phone number must not exceed 15 digits")
	}

	return nil
}
