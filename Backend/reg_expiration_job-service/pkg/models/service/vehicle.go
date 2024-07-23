package service

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/Manas8803/The-PUC-Project__BackEnd/reg_expiration_job-service/pkg/models/db"
)

type Date struct {
	Year  int
	Month int
	Day   int
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
	LastCheckDate    *Date  `json:"last_check_date"`
}

func convertVehicleDynToVehicle(dynVehicle db.Vehicle) (Vehicle, error) {
	reg_upto_date, err := parseDate(dynVehicle.RegUpto)
	if err != nil {
		return Vehicle{}, err
	}

	puc_upto_date, err := parseDate(dynVehicle.PucUpto)
	if err != nil {
		return Vehicle{}, err
	}

	mobile, err := strconv.ParseInt(dynVehicle.Mobile, 10, 64)
	if err != nil {
		return Vehicle{}, err
	}

	last_reg_date, err := parseDate(dynVehicle.LastCheckDate)
	if err != nil {
		return Vehicle{}, err
	}

	return Vehicle{
		OwnerName:        dynVehicle.OwnerName,
		OfficeName:       dynVehicle.OfficeName,
		RegNo:            dynVehicle.RegNo,
		VehicleClassDesc: dynVehicle.VehicleClassDesc,
		Model:            dynVehicle.Model,
		RegUpto:          &reg_upto_date,
		VehicleType:      dynVehicle.VehicleType,
		Mobile:           mobile,
		PucUpto:          &puc_upto_date,
		LastCheckDate:    &last_reg_date,
	}, nil
}

func IsStructEmpty(obj interface{}) bool {
	value := reflect.ValueOf(obj)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return false
	}

	zero := reflect.Zero(value.Type())
	return value.Interface() == zero.Interface()
}

func parseDate(dateStr string) (Date, error) {
	parts := strings.Split(dateStr, "-")
	if len(parts) != 3 {
		return Date{}, fmt.Errorf("invalid date format: %s", dateStr)
	}

	day, err := strconv.Atoi(parts[0])
	if err != nil {
		return Date{}, err
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return Date{}, err
	}

	year, err := strconv.Atoi(parts[2])
	if err != nil {
		return Date{}, err
	}

	return Date{
		Year:  year,
		Month: month,
		Day:   day,
	}, nil
}

func convertVehicleToVehicleDyn(vehicle Vehicle) (db.Vehicle, error) {

	puc_upto := fmt.Sprint(vehicle.PucUpto.Day + vehicle.PucUpto.Month + vehicle.PucUpto.Year)
	reg_upto := fmt.Sprint(vehicle.RegUpto.Day + vehicle.RegUpto.Month + vehicle.RegUpto.Year)
	last_check_date := fmt.Sprint(vehicle.LastCheckDate.Day + vehicle.LastCheckDate.Month + vehicle.LastCheckDate.Year)
	mobile := fmt.Sprint(vehicle.Mobile)

	return db.Vehicle{
		OwnerName:        vehicle.OwnerName,
		OfficeName:       vehicle.OfficeName,
		RegNo:            vehicle.RegNo,
		VehicleClassDesc: vehicle.VehicleClassDesc,
		Model:            vehicle.Model,
		VehicleType:      vehicle.VehicleType,
		RegUpto:          reg_upto,
		PucUpto:          puc_upto,
		Mobile:           mobile,
		LastCheckDate:    last_check_date,
	}, nil
}

func GetAllVehicles() (*[]Vehicle, error) {
	vehicles_dyn, err := db.GetAllVehicles()
	if err != nil {
		log.Println("error in fetching vehicle list: ", err)
		return &[]Vehicle{}, err
	}

	log.Println(vehicles_dyn)
	vehicles := make([]Vehicle, 0, len(*vehicles_dyn))
	for _, dynVehicle := range *vehicles_dyn {
		vehicle, err := convertVehicleDynToVehicle(dynVehicle)
		if err != nil {
			log.Println("error in converting vehicles: ", err)
			return &[]Vehicle{}, err
		}
		vehicles = append(vehicles, vehicle)
	}

	return &vehicles, nil
}
func UpdateLastCheckDate(vehicle Vehicle) error {
	vehicle_dyn, err := convertVehicleToVehicleDyn(vehicle)
	if err != nil {
		log.Println("error converting service vehicle to db vehicle: ", err)
		return err
	}

	log.Println("update last check date: ", vehicle_dyn)

	err = db.UpdateLastCheckDate(vehicle_dyn)
	if err != nil {
		log.Println("error in saving data in table : ", err)
		return err
	}
	return nil
}
