package gvk

import (
	"math/rand"
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
func (a API) GetUpdates(opts *UpdateOptions) (res APIResponseUpdate, err error) {
	var vals = make(url.Values)
	vals.Set("access_token", a.token)
	vals.Set("v", APIVersion)

	return get[APIResponseUpdate](opts.Server, "", addValues(vals, opts))
}

// MessagesGetLongPollServer https://dev.vk.com/ru/method/groups.getLongPollServer
func (a API) GroupsGetLongPollServer(opts *GetLongPollServerOptions) (res APIResponseGetLongPollServer, err error) {
	var vals = make(url.Values)
	vals.Set("access_token", a.token)
	vals.Set("v", APIVersion)
	return get[APIResponseGetLongPollServer](a.base, "groups.getLongPollServer", addValues(vals, opts))

}

// GroupsSetLongPollSettings https://dev.vk.com/ru/method/groups.setLongPollSettings
func (a API) GroupsSetLongPollSettings(opts *SetLongPollSettingsOptions) (res APIResponseSetLongPollSettings, err error) {
	var vals = make(url.Values)

	vals.Set("access_token", a.token)
	vals.Set("v", APIVersion)
	return get[APIResponseSetLongPollSettings](a.base, "groups.setLongPollSettings", addValues(vals, opts))

}

func (a API) MessagesSend(text string, opts *MessagesSendOptions) (res APIResponseMessagesSend, err error) {
	var vals = make(url.Values)
	vals.Set("message", text)
	vals.Set("access_token", a.token)
	vals.Set("v", APIVersion)
	vals.Set("random_id", strconv.Itoa(int(rand.Uint32())))
	return get[APIResponseMessagesSend](a.base, "messages.send", addValues(vals, opts))
}

func (a API) MessagesSendMessageEventAnswer(text string, opts *MessageEventAnswerOptions) (res APIResponseMessageEventAnswer, err error) {
	var vals = make(url.Values)
	vals.Set("message", text)
	vals.Set("access_token", a.token)
	vals.Set("v", APIVersion)
	vals.Set("random_id", strconv.Itoa(int(rand.Uint32())))
	return get[APIResponseMessageEventAnswer](a.base, "messages.sendMessageEventAnswer", addValues(vals, opts))
}

// UsersGet https://dev.vk.com/ru/method/users.get
func (a API) UsersGet(opts *UsersGetOptions) (res APIResponseUsersGet, err error) {
	var vals = make(url.Values)
	vals.Set("access_token", a.token)
	vals.Set("v", APIVersion)
	vals.Set("random_id", strconv.Itoa(int(rand.Uint32())))
	return get[APIResponseUsersGet](a.base, "users.get", addValues(vals, opts))
}
