package api

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Manas8803/The-PUC-Project__BackEnd/vrc-service/pkg/models/service"
)

func GetVehicleInfoByRegNo(reg_no string) (*service.Vehicle, error) {

	url := os.Getenv("API_URL")

	payload := strings.NewReader("{\n    \"reg_no\": " + reg_no + ",\n    \"consent\": \"Y\",\n    \"consent_text\": \"I hear by declare my consent agreement for fetching my information via AITAN Labs API\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", os.Getenv("API_KEY"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("API_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println()
		return &service.Vehicle{}, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var v *service.Vehicle
	v.FromJson(body)

	return v, err
}
