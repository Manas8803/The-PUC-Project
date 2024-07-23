package socket

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Manas8803/The-PUC-Project__BackEnd/vrc-service/pkg/models/service"
	"github.com/gorilla/websocket"
)

type Payload struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

func ReportAuthority(vehicle *service.Vehicle) error {
	url := os.Getenv("REPORT_WEBSOCKET_URL") 
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Failed to connect to WebSocket:", err)
	}
	defer conn.Close()

	data := Payload{
		Action: "report",
		Data:   vehicle,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Failed to marshal data to JSON:", err)
	}

	err = conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Fatal("Failed to send message:", err)
	}

	fmt.Println("Message sent successfully")
	return nil
}
