package model

import "encoding/json"

type Response struct {
	Message string          `json:"message"`
	Data    json.RawMessage `jsons:"data"`
}
