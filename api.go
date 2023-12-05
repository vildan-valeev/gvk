package gvk

import "net/url"

const (
	APIVersion = "5.131"
)

type API struct {
	token string
	base  string
}

// NewAPI returns a new API object.
func NewAPI(token string) API {
	return API{
		token: token,
		base:  "https://api.vk.com/method/",
	}
}

// GetUpdates is used to receive incoming updates using long polling.
func (a API) GetUpdates(opts *UpdateOptions) (res APIResponseUpdate, err error) {
	return get[APIResponseUpdate](a.base, "getUpdates", urlValues(opts))
}

// GetUpdates is used to receive incoming updates using long polling.
func (a API) MessagesSend(text string, chatID int64, opts *MessagesSendOptions) (res APIResponseUpdate, err error) {
	var vals = make(url.Values)
	vals.Set("access_token", a.token)
	vals.Set("v", APIVersion)

	vals.Set("message", text)

	return get[APIResponseUpdate](a.base, "messages.send", addValues(vals, opts))
}

//// GetUpdates is used to receive incoming updates using long polling.
//func (a API) MessagesSend(opts *UpdateOptions) (res APIResponseUpdate, err error) {
//	return get[APIResponseUpdate](a.base, "getUpdates", urlValues(opts))
//}

// MessagesGetLongPollServer returns data required for connection to a Long Poll server.
//
// https://vk.com/dev/messages.getLongPollServer
func (a API) MessagesGetLongPollServer(opts *GetLongPollServerOptions) (res APIResponseGetLongPollServer, err error) {
	var vals = make(url.Values)

	vals.Set("access_token", a.token)
	vals.Set("v", APIVersion)
	return get[APIResponseGetLongPollServer](a.base, "groups.getLongPollServer", addValues(vals, opts))

}
