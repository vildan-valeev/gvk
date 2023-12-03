package main

import (
	"fmt"
	sdk "github.com/SevereCloud/vksdk/v2/api"
	"github.com/vildan-valeev/gvk"
	"log"
	"strings"
)

// Recursive type definition of the bot state function.
type stateFn func(event *gvk.GroupEvent) stateFn

type bot struct {
	chatID int
	state  stateFn
	name   string

	API *sdk.VK
}

const token = "b30fae3f8d488e20cdbe041cbec9a0aa62e7c52e6107f97f97a9fd9007abe32223e1373cce590bfabf374"

func newBot(chatID int) gvk.Bot {
	b := &bot{
		chatID: chatID,
		API:    sdk.NewVK(token),
	}

	b.state = b.handleMessage
	return b
}

func (b *bot) Update(update *gvk.GroupEvent) {
	b.state = b.state(update)
}

func (b *bot) handleMessage(update *gvk.GroupEvent) stateFn {
	if strings.HasPrefix(update.ParsedObj.Message.Text, "ping") {
		b.API.MessagesSend(sdk.Params{"user_ids": "", "peer_ids": ""})
		// Here we return b.handleName since next time we receive a message it
		// will be the new name.
		return b.handleName
	}

	return b.handleMessage
}

func (b *bot) handleName(update *gvk.GroupEvent) stateFn {
	b.name = update.ParsedObj.Message.Text
	b.API.MessagesSend(sdk.Params{"user_ids": "", "peer_ids": ""})
	// Here we return b.handleMessage since the next time we receive a message
	// it will be handled in the default way.
	return b.handleMessage
}

func main() {
	fmt.Print("Start!")

	vk := sdk.NewVK(token)

	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		log.Fatal(err)
	}

	dsp := gvk.NewLongPoll(token, group[0].ID, newBot)
	log.Println(dsp.Run())
}
