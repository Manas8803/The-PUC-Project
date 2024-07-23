package responses

type UserResponse struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

type UserResponse_doc struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type ErrorResponse_doc struct {
	Message string `json:"message"`
}
