package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/vildan-valeev/gvk"
)

type stateFn func(event *gvk.Update) stateFn

type Bot struct {
	chatID int64
	state  stateFn
	name   string

	gvk.API
}

const (
	groupID = 1234567
	token   = "token here"
)

func newBot(chatID int64) gvk.Bot {
	b := &Bot{
		chatID: chatID,
		API:    gvk.NewAPI(token),
	}

	b.state = b.EntryHandler
	return b
}

func (b *Bot) Update(update *gvk.Update) {
	b.state = b.state(update)
}

func (b *Bot) EntryHandler(update *gvk.Update) stateFn {
	if strings.HasPrefix(update.MessageNew.Message.Text, "ping") {
		b.MessagesSend("pong", &gvk.MessagesSendOptions{UserID: b.chatID})
		return b.handleNext
	}

	b.MessagesSend("not understand...", &gvk.MessagesSendOptions{UserID: b.chatID})

	return b.EntryHandler
}

func (b *Bot) handleNext(update *gvk.Update) stateFn {
	b.name = update.MessageNew.Message.Text
	b.MessagesSend("pong again )))", &gvk.MessagesSendOptions{
		UserID: b.chatID})

	return b.EntryHandler
}

func main() {
	fmt.Print("Start!")

	dsp, err := gvk.NewDispatcher(token, groupID, newBot)

	log.Println(err)
	log.Println(dsp.Poll())
}
