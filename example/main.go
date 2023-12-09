package main

import (
	"fmt"
	"github.com/vildan-valeev/gvk"

	"log"
	"strings"
)

type stateFn func(event *gvk.Update) stateFn

type Bot struct {
	chatID int64
	state  stateFn
	name   string

	gvk.API
}

const (
	groupID = 194299208
	token   = "b30fae3f8d488e20cdbe041cbec9a0aa62e7c52e6107f97f97a9fd9007abe32223e1373cce590bfabf374"
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
	log.Printf("ПОЛУЧЕНО СООБЩЕНИЕ! %s \n", update.MessageNew.Message.Text)
	if strings.HasPrefix(update.MessageNew.Message.Text, "ping") {
		b.MessagesSend(
			"pong",
			&gvk.MessagesSendOptions{
				UserID: b.chatID,
				//PeerID:   update.MessageNew.Message.PeerID,
				//RandomID: 1,
			})

		return b.handleNext
	}

	b.MessagesSend(
		"not understand...",
		&gvk.MessagesSendOptions{
			UserID: b.chatID,
			//PeerID:   update.MessageNew.Message.PeerID,
			//RandomID: 2,
		})

	return b.EntryHandler
}

func (b *Bot) handleNext(update *gvk.Update) stateFn {
	b.name = update.MessageNew.Message.Text
	b.MessagesSend(
		"pong again )))",
		&gvk.MessagesSendOptions{
			UserID: b.chatID,
			//PeerID:   update.MessageNew.Message.PeerID,
			//RandomID: 0,
		})

	return b.EntryHandler
}

func main() {
	fmt.Print("Start!")

	dsp, err := gvk.NewDispatcher(token, groupID, newBot)

	log.Println(err)
	log.Println(dsp.Poll())
}
