package util

import (
	"fmt"
	"time"

	"github.com/Manas8803/The-PUC-Project__BackEnd/reg_expiration_job-service/pkg/models/service"
)

func IsPucExpired(v *service.Vehicle) (bool, error) {
	if v.PucUpto == nil {
		return false, fmt.Errorf("PucUpto date is nil")
	}

	pucUptoTime := time.Date(v.PucUpto.Year, time.Month(v.PucUpto.Month), v.PucUpto.Day, 0, 0, 0, 0, time.UTC)

	today := time.Now().UTC()
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)

	return pucUptoTime.Before(today), nil
}

func CheckWarningDays(v *service.Vehicle) (bool, error) {
	if v.PucUpto == nil {
		return false, fmt.Errorf("PucUpto date is nil")
	}

	pucExpiryDate := time.Date(v.PucUpto.Year, time.Month(v.PucUpto.Month), v.PucUpto.Day, 0, 0, 0, 0, time.UTC)

	currentDate := time.Now().UTC()

	diff := pucExpiryDate.Sub(currentDate)

	return diff.Hours() <= 5*24, nil
}
