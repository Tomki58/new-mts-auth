package serializer

import "encoding/json"

// Response is a structure that represents server's responses.
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   error       `json:"error"`
}

// SerializeResponseJson serializes the server response in JSON format.
func SerializeResponseJson(data interface{}) ([]byte, error) {
	response := new(Response)

	switch data := data.(type) {
	case error:
		response.Error = data

	default:
		response.Success = true
		response.Data = data
	}

	return json.Marshal(response)
}
