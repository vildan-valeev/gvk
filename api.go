package gvk

import (
	"fmt"
	"net/url"
)

// API is the object that contains all the functions that wrap those of the Telegram Bot API.
type API struct {
	token string
	base  string
}

// NewAPI returns a new API object.
func NewAPI(token string) API {
	return API{
		token: token,
		base:  fmt.Sprintf("https://api.vk.com/method/%s/", token),
	}
}

// GetUpdates is used to receive incoming updates using long polling.
func (a API) GetUpdates(opts *UpdateOptions) (res APIResponseUpdate, err error) {
	return get[APIResponseUpdate](a.base, "getUpdates", urlValues(opts))
}

// MessagesSend is used to send text messages.
func (a API) MessagesSend(text string, chatID int64, opts *MessageOptions) (res APIResponseMessage, err error) {
	var vals = make(url.Values)

	vals.Set("text", text)
	vals.Set("chat_id", itoa(chatID))
	return get[APIResponseMessage](a.base, "sendMessage", addValues(vals, opts))
}
