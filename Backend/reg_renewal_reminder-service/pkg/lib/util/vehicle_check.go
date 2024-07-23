package util

import (
	"fmt"
	"log"
	"time"

	"github.com/Manas8803/The-PUC-Project__BackEnd/reg_renewal_reminder-service/pkg/models/service"
)

func CheckRegNoIfExists(reg_no string) (bool, *service.Vehicle, error) {
	v, err := service.GetVehicleOnRegNo(reg_no)
	//* If an error occurs, return false with error
	if err != nil {
		return false, &service.Vehicle{}, err
	}

	//* If the vehicle does not exist in db return false with empty vehicle object
	if service.IsStructEmpty(v) {
		log.Println("NO VEHICLE EXISTS WITH THE GIVEN REG NO")
		return false, &service.Vehicle{}, nil
	}

	return true, v, nil
}

func IsNextCheckDateToday(v *service.Vehicle) (bool, error) {
	if v.LastCheckDate == nil {
		return false, fmt.Errorf("last check date is nil")
	}

	last_check_date := time.Date(v.LastCheckDate.Year, time.Month(v.LastCheckDate.Month), (v.LastCheckDate.Day), 0, 0, 0, 0, time.Local)

	next_check_date := last_check_date.AddDate(0, 0, 2)

	today := time.Now()
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.Local)
	log.Println(next_check_date, "------------", today)

	//* Next check date is before or equal to today
	return next_check_date.Before(today) || next_check_date.Equal(today), nil
}

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
