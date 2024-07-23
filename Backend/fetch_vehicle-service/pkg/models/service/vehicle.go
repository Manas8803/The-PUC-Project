package service

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Manas8803/The-PUC-Project__BackEnd/fetch_vehicle-service/pkg/models/db"
)

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
func convertDBVehicleToServiceVehicle(dbVehicle *db.Vehicle) (Vehicle, error) {
	var serviceVehicle Vehicle

	serviceVehicle.OwnerName = dbVehicle.OwnerName
	serviceVehicle.OfficeName = dbVehicle.OfficeName
	serviceVehicle.RegNo = dbVehicle.RegNo
	serviceVehicle.VehicleClassDesc = dbVehicle.VehicleClassDesc
	serviceVehicle.Model = dbVehicle.Model
	serviceVehicle.VehicleType = dbVehicle.VehicleType

	// Parse RegUpto
	regUpto, err := parseDate(dbVehicle.RegUpto)
	if err != nil {
		return Vehicle{}, fmt.Errorf("failed to parse reg_upto date: %w", err)
	}
	serviceVehicle.RegUpto = regUpto

	// Parse Mobile
	mobile, err := strconv.ParseInt(dbVehicle.Mobile, 10, 64)
	if err != nil {
		return Vehicle{}, fmt.Errorf("failed to parse mobile: %w", err)
	}
	serviceVehicle.Mobile = mobile

	// Parse PucUpto
	pucUpto, err := parseDate(dbVehicle.PucUpto)
	if err != nil {
		return Vehicle{}, fmt.Errorf("failed to parse puc_upto date: %w", err)
	}
	serviceVehicle.PucUpto = pucUpto

	// Parse LastCheckDate
	lastCheckDate, err := parseDate(dbVehicle.LastCheckDate)
	if err != nil {
		return Vehicle{}, fmt.Errorf("failed to parse last_check_date: %w", err)
	}
	serviceVehicle.LastCheckDate = lastCheckDate

	// Set PucStatus
	serviceVehicle.PucStatus = dbVehicle.PucUpto != ""

	return serviceVehicle, nil
}

func FetchVehicles(office_name string) ([]Vehicle, error) {
	db_vehicles, err := db.FetchVehicles(office_name)
	if err != nil {
		log.Println("Unable to fetch vehicles:", err)
		return nil, err
	}

	var service_vehicles []Vehicle
	for _, dbVehicle := range db_vehicles {
		serviceVehicle, err := convertDBVehicleToServiceVehicle(dbVehicle)
		if err != nil {
			log.Printf("Failed to convert DB vehicle to service vehicle: %v", err)
			continue
		}
		service_vehicles = append(service_vehicles, serviceVehicle)
	}

	return service_vehicles, nil
}
