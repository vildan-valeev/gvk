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

// https://dev.vk.com/ru/api/api-requests#Query-%D0%BF%D0%B0%D1%80%D0%B0%D0%BC%D0%B5%D1%82%D1%80%D1%8B
func (a API) defaultParams() url.Values {
	var vals = make(url.Values)

	vals.Set("access_token", a.token)
	vals.Set("v", APIVersion)
	return vals
}

// GetUpdates https://dev.vk.com/ru/api/bots-long-poll/getting-started
func (a API) GetUpdates(opts *UpdateOptions) (res APIResponseUpdate, err error) {
	return get[APIResponseUpdate](opts.Server, "", addValues(a.defaultParams(), opts))
}

// GroupsGetLongPollServer https://dev.vk.com/ru/method/groups.getLongPollServer
func (a API) GroupsGetLongPollServer(opts *GetLongPollServerOptions) (res APIResponseGetLongPollServer, err error) {
	return get[APIResponseGetLongPollServer](a.base, "groups.getLongPollServer", addValues(a.defaultParams(), opts))

}

// GroupsSetLongPollSettings https://dev.vk.com/ru/method/groups.setLongPollSettings
func (a API) GroupsSetLongPollSettings(opts *SetLongPollSettingsOptions) (res APIResponseSetLongPollSettings, err error) {
	return get[APIResponseSetLongPollSettings](a.base, "groups.setLongPollSettings", addValues(a.defaultParams(), opts))
}

// GroupsIsMember https://dev.vk.com/ru/method/groups.isMember
func (a API) GroupsIsMember(opts *GroupsIsMemberOptions) (res APIResponseGroupsIsMember, err error) {
	return get[APIResponseGroupsIsMember](a.base, "groups.isMember", addValues(a.defaultParams(), opts))
}

// MessagesSend https://dev.vk.com/ru/method/messages.send
func (a API) MessagesSend(opts *MessagesSendOptions) (res APIResponseMessagesSend, err error) {
	var vals = a.defaultParams()
	vals.Set("random_id", strconv.Itoa(int(rand.Uint32())))
	return get[APIResponseMessagesSend](a.base, "messages.send", addValues(vals, opts))
}

// MessagesEdit https://dev.vk.com/ru/method/messages.edit
func (a API) MessagesEdit(opts *MessagesEditOptions) (res APIResponseMessagesEdit, err error) {
	return get[APIResponseMessagesEdit](a.base, "messages.edit", addValues(a.defaultParams(), opts))
}

// MessagesSendMessageEventAnswer https://dev.vk.com/ru/method/messages.sendMessageEventAnswer
func (a API) MessagesSendMessageEventAnswer(opts *MessageEventAnswerOptions) (res APIResponseMessageEventAnswer, err error) {
	return get[APIResponseMessageEventAnswer](a.base, "messages.sendMessageEventAnswer", addValues(a.defaultParams(), opts))
}

// UsersGet https://dev.vk.com/ru/method/users.get
func (a API) UsersGet(opts *UsersGetOptions) (res APIResponseUsersGet, err error) {
	return get[APIResponseUsersGet](a.base, "users.get", addValues(a.defaultParams(), opts))
}
