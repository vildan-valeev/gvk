package gvk

import (
	"net/url"
	"strconv"
)

const (
	APIVersion = "5.131"
)

// https://dev.vk.com/ru/api/community-messages/getting-started#Доступные инструменты и методы
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

// GetUpdates https://dev.vk.com/ru/api/bots-long-poll/getting-started
func (a API) GetUpdates(opts *UpdateOptions) (res ResponseUpdate, err error) {
	var vals = make(url.Values)
	vals.Set("access_token", a.token)
	vals.Set("v", APIVersion)

	return getUpdates(a.base, "", urlValues(opts))
}

func (a API) MessagesSend(text string, chatID int64, opts *MessagesSendOptions) (res ResponseMessagesSend, err error) {
	var vals = make(url.Values)
	vals.Set("message", text)
	vals.Set("chat_id", strconv.FormatInt(chatID, 10))

	return get[ResponseMessagesSend](a.base, "messages.send", addValues(vals, opts))
}

// MessagesGetLongPollServer https://dev.vk.com/ru/method/groups.getLongPollServer
func (a API) GroupsGetLongPollServer(opts *GetLongPollServerOptions) (res ResponseGetLongPollServer, err error) {
	var vals = make(url.Values)
	vals.Set("access_token", a.token)
	vals.Set("v", APIVersion)
	return get[ResponseGetLongPollServer](a.base, "groups.getLongPollServer", addValues(vals, opts))

}

// GroupsSetLongPollSettings https://dev.vk.com/ru/method/groups.setLongPollSettings
func (a API) GroupsSetLongPollSettings(opts *SetLongPollSettingsOptions) (res ResponseSetLongPollSettings, err error) {
	var vals = make(url.Values)

	vals.Set("access_token", a.token)
	vals.Set("v", APIVersion)
	return get[ResponseSetLongPollSettings](a.base, "groups.setLongPollSettings", addValues(vals, opts))

}
