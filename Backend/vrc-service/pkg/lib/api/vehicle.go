package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Manas8803/The-PUC-Project__BackEnd/vrc-service/pkg/models/service"
)

func GetVehicleInfoByRegNo(reg_no string) (*service.Vehicle, error) {
	url := os.Getenv("API_URL")

	payload := fmt.Sprintf(`{
        "reg_no": "%s",
        "consent": "Y",
        "consent_text": "I hear by declare my consent agreement for fetching my information via AITAN Labs API"
    }`, reg_no)

	payloadReader := strings.NewReader(payload)
	req, err := http.NewRequest("POST", url, payloadReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("x-rapidapi-key", os.Getenv("API_KEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("API_HOST"))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Check for non-200 status codes
	if res.StatusCode != http.StatusOK {
		var errResp struct {
			Message string `json:"message"`
		}
		if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, fmt.Errorf("error response with status code %d: %s", res.StatusCode, string(body))
		}
		switch res.StatusCode {
		case http.StatusTooManyRequests:
			return nil, fmt.Errorf("rate limit exceeded (429): %s", errResp.Message)
		case http.StatusNotFound:
			return nil, fmt.Errorf("resource not found (404): %s", errResp.Message)
		default:
			return nil, fmt.Errorf("error response with status code %d: %s", res.StatusCode, errResp.Message)
		}
	}

	v := &service.Vehicle{}
	err = v.FromJson(body)
	if err != nil {
		log.Println("Error in unmarshalling")
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}
	log.Println(v)

	return v, nil
}
