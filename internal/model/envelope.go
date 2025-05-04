package model

import "encoding/json"

type Envelope struct {
	EventID string
	Items   []EnvelopeItem
}

type EnvelopeItem struct {
	Type    string          `json:"type"`
	Data    json.RawMessage `json:"data"`
	Headers map[string]string
}
