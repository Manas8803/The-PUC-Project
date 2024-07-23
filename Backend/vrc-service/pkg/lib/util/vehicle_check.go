package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Manas8803/The-PUC-Project__BackEnd/vrc-service/pkg/models/service"
)

func IsPucExpired(v *service.Vehicle) (bool, error) {
	if v.PucUpto == nil {
		return false, fmt.Errorf("PucUpto date is nil")
	}

	pucUptoTime := time.Date(v.PucUpto.Year, time.Month(v.PucUpto.Month), v.PucUpto.Day, 0, 0, 0, 0, time.UTC)

	today := time.Now().UTC()
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)
	v.PucStatus = pucUptoTime.Before(today)

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

func UpdateLastCheckDate() *service.Date {
	today := time.Now().Format("02-01-2006")
	today_date := parseDate(today)
	return &today_date
}

func parseDate(dateStr string) service.Date {
	parts := strings.Split(dateStr, "-")

	day, err := strconv.Atoi(parts[0])
	if err != nil {
		return service.Date{}
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return service.Date{}
	}

	year, err := strconv.Atoi(parts[2])
	if err != nil {
		return service.Date{}
	}

	return service.Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
}
