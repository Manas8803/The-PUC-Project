package service

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Manas8803/The-PUC-Project__BackEnd/vrc-service/pkg/models/db"
)

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

func parseDate(dateStr string) (*Date, error) {
	parts := strings.Split(dateStr, "-")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid date format: %s", dateStr)
	}

	year, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	day, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}

	return &Date{
		Year:  year,
		Month: month,
		Day:   day,
	}, nil
}

type Vehicle struct {
	OwnerName        string `json:"owner_name"`
	OfficeName       string `json:"office_name"`
	RegNo            string `json:"reg_no"`
	VehicleClassDesc string `json:"vehicle_class_desc"`
	Model            string `json:"model"`
	RegUpto          *Date  `json:"reg_upto"`
	VehicleType      string `json:"vehicle_type"`
	Mobile           int64  `json:"mobile"`
	PucUpto          *Date  `json:"puc_upto"`
	PucStatus        bool   `json:"puc_status"`
	LastCheckDate    *Date  `json:"last_check_date"`
}

func (v *Vehicle) FromJson(data []byte) error {
	var resp struct {
		Result struct {
			OwnerName        string `json:"owner_name"`
			OfficeName       string `json:"office_name"`
			RegNo            string `json:"reg_no"`
			VehicleClassDesc string `json:"vehicle_class_desc"`
			Model            string `json:"model"`
			RegUpto          string `json:"reg_upto"`
			VehicleType      string `json:"vehicle_type"`
			VehiclePuccDetails struct {
				PuccUpto string `json:"pucc_upto"`
			} `json:"vehicle_pucc_details"`
		} `json:"result"`
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return err
	}

	v.OwnerName = resp.Result.OwnerName
	v.OfficeName = resp.Result.OfficeName
	v.RegNo = resp.Result.RegNo
	v.VehicleClassDesc = resp.Result.VehicleClassDesc
	v.Model = resp.Result.Model
	v.VehicleType = resp.Result.VehicleType

	regUpto, err := parseDate(resp.Result.RegUpto)
	if err != nil {
		return fmt.Errorf("error parsing reg_upto: %v", err)
	}
	v.RegUpto = regUpto

	pucUpto, err := parseDate(resp.Result.VehiclePuccDetails.PuccUpto)
	if err != nil {
		return fmt.Errorf("error parsing puc_upto: %v", err)
	}
	v.PucUpto = pucUpto

	// Set PucStatus based on whether PucUpto is in the future
	now := time.Now()
	v.PucStatus = time.Date(v.PucUpto.Year, time.Month(v.PucUpto.Month), v.PucUpto.Day, 0, 0, 0, 0, time.UTC).After(now)

	// Set LastCheckDate to current date
	v.LastCheckDate = &Date{
		Year:  now.Year(),
		Month: int(now.Month()),
		Day:   now.Day(),
	}

	// Mobile number is not present in the provided JSON. If it's needed, you may have to get it from another source.
	v.Mobile = 0 // Set to default value or handle as needed

	return nil
}

func convertVehicleToVehicleDyn(vehicle Vehicle) (db.Vehicle, error) {

	puc_upto := fmt.Sprint(vehicle.PucUpto.Day + vehicle.PucUpto.Month + vehicle.PucUpto.Year)
	reg_upto := fmt.Sprint(vehicle.RegUpto.Day + vehicle.RegUpto.Month + vehicle.RegUpto.Year)
	last_check_date := fmt.Sprint(vehicle.LastCheckDate.Day + vehicle.LastCheckDate.Month + vehicle.LastCheckDate.Year)
	mobile := fmt.Sprint(vehicle.Mobile)

	return db.Vehicle{
		OwnerName:        vehicle.OwnerName,
		OfficeName:       vehicle.OfficeName,
		RegNo:            vehicle.OwnerName,
		VehicleClassDesc: vehicle.VehicleClassDesc,
		Model:            vehicle.Model,
		VehicleType:      vehicle.VehicleType,
		RegUpto:          reg_upto,
		PucUpto:          puc_upto,
		Mobile:           mobile,
		LastCheckDate:    last_check_date,
	}, nil
}

func SaveOrUpdateVehicle(vehicle Vehicle) error {
	vehicle_dyn, err := convertVehicleToVehicleDyn(vehicle)
	if err != nil {
		log.Println("error converting service vehicle to db vehicle: ", err)
		return err
	}

	err = db.SaveOrUpdateVehicle(vehicle_dyn)
	if err != nil {
		log.Println("error in saving data in table : ", err)
		return err
	}
	return nil
}