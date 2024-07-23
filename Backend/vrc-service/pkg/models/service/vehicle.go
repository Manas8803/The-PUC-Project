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
		return &Date{}, fmt.Errorf("invalid date format: %s", dateStr)
	}

	day, err := strconv.Atoi(parts[0])
	if err != nil {
		return &Date{}, err
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return &Date{}, err
	}

	year, err := strconv.Atoi(parts[2])
	if err != nil {
		return &Date{}, err
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
	var resp map[string]interface{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return err
	}

	result, ok := resp["result"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid response format")
	}

	v.OwnerName = result["owner_name"].(string)
	v.OfficeName = result["office_name"].(string)
	v.RegNo = result["reg_no"].(string)
	v.VehicleClassDesc = result["vehicle_class_desc"].(string)
	v.Model = result["model"].(string)

	regUptoStr, ok := result["reg_upto"].(string)
	if ok {

		reg_upto, err := parseDate(regUptoStr)
		if err != nil {
			log.Println("error parsing date: ", err)
			return err
		}
		v.RegUpto = reg_upto
	}

	v.VehicleType = result["vehicle_type"].(string)
	mobile_no, err := strconv.ParseFloat(fmt.Sprint(result["mobile_no"]), 64)
	if err != nil {
		log.Println("error parsing mobile no. : ", err)
		return err
	}
	v.Mobile = int64(mobile_no)

	pucUptoStr, ok := result["vehicle_pucc_details"].(map[string]interface{})["pucc_upto"].(string)
	if ok {
		puc_upto, err := parseDate(pucUptoStr)
		if err != nil {
			log.Println("error parsing date : ", err)
			return err
		}
		v.PucUpto = puc_upto
	}
	today := time.Now()

	v.LastCheckDate = &Date{Day: today.Day(), Month: int(today.Month()), Year: today.Year()}
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
	log.Println("error in saving data in table : ", err)
	if err != nil {
		return err
	}
	return nil
}
