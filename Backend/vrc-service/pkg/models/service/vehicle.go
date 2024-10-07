package service

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
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
