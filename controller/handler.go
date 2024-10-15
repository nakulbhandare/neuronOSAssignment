package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/neuronOS/cmd"
)

// CommandRequest for parsing incomming request
type CommandRequest struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

// CommandResponse for OutPut Response format
type CommandResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

// HandleCommand  for handling requests and executing the commands
func HandleCommand(cmdr cmd.Commander) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request ::", r)
		var req CommandRequest

		// decode the request body
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		var resp CommandResponse

		// request type check
		switch req.Type {
		// Execute If command is Ping
		case "ping":
			result, err := cmdr.Ping(req.Payload)
			if err != nil {
				resp = CommandResponse{Success: false, Error: err.Error()}
			} else {
				resp = CommandResponse{Success: true, Data: result}
			}

		// Execute If command is sysinfo
		case "sysinfo":
			result, err := cmdr.GetSystemInfo()
			if err != nil {
				resp = CommandResponse{Success: false, Error: err.Error()}
			} else {
				resp = CommandResponse{Success: true, Data: result}
			}
		default:
			// if there is command which is not supported by system should return error response with valid error message.
			resp = CommandResponse{Success: false, Error: "Unknown command"}
		}

		// returning the http Response in json format.
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
