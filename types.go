package gvk

import "encoding/json"

type Response struct {
	Ts      string   `json:"ts"`
	Updates []Update `json:"updates"`
	Failed  int      `json:"failed"`
}

// GroupEvent struct.
type Update struct {
	Type    EventType       `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
	EventID string          `json:"event_id"`
	V       string          `json:"v"`
	Secret  string          `json:"secret"`
}

// ChatID returns the ID of the chat the update is coming from.
func (u Update) ChatID() int64 {
	switch {
	case u.Message != nil:
		return u.Message.Chat.ID
	case u.EditedMessage != nil:
		return u.EditedMessage.Chat.ID

	default:
		return 0
	}
}

// APIResponseBase is a base type that represents the incoming response from Telegram servers.
// Used by APIResponse* to slim down the implementation.
type APIResponseBase struct {
	Description string `json:"description,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Ok          bool   `json:"ok"`
}

// Base returns the APIResponseBase itself.
func (a APIResponseBase) Base() APIResponseBase {
	return a
}

type APIResponseUpdate struct {
	Result []*Update `json:"result,omitempty"`
	APIResponseBase
}
