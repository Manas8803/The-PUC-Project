package main

import (
	"fmt"
	"strings"
	"net/http"
	"io"
)

func main() {

	url := "https://rto-vehicle-information-verification-india.p.rapidapi.com/api/v1/rc/vehicleinfo"

	payload := strings.NewReader("{\"reg_no\":\"GJ01JT0459\",\"consent\":\"Y\",\"consent_text\":\"I hear by declare my consent agreement for fetching my information via AITAN Labs API\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("x-rapidapi-key", "a412d23e14msh44f18b76cb0a9fcp125ee2jsn56618941a894")
	req.Header.Add("x-rapidapi-host", "rto-vehicle-information-verification-india.p.rapidapi.com")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}