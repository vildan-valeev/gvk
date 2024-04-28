package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/vildan-valeev/gvk"
)

type stateFn func(event *gvk.Update) stateFn

type Bot struct {
	chatID int64
	state  stateFn
	name   string

	gvk.API
}

var dsp *gvk.Dispatcher

func newBot(chatID int64) gvk.Bot {
	b := &Bot{
		chatID: chatID,
		API:    gvk.NewAPI(token),
	}
	go b.selfDestruct(time.After(time.Minute))
	b.state = b.EntryHandler
	return b
}

func (b *Bot) selfDestruct(timech <-chan time.Time) {
	<-timech
	b.MessagesSend(&gvk.MessagesSendOptions{Message: "goodbye...", UserID: b.chatID})
	dsp.DelSession(b.chatID)
}

func (b *Bot) Update(update *gvk.Update) {
	b.state = b.state(update)
}

func (b *Bot) EntryHandler(update *gvk.Update) stateFn {
	if strings.HasPrefix(update.Object.MessageNew.Message.Text, "ping") {
		b.MessagesSend(&gvk.MessagesSendOptions{Message: "pong", UserID: b.chatID})
		return b.handleNext
	}

	b.MessagesSend(&gvk.MessagesSendOptions{Message: "not understand...", UserID: b.chatID})

	return b.EntryHandler
}

func (b *Bot) handleNext(update *gvk.Update) stateFn {

	b.MessagesSend(&gvk.MessagesSendOptions{
		Message: "pong again )))",
		UserID:  b.chatID,
	})

	return b.EntryHandler
}

func main() {
	fmt.Print("Start!")
	var err error

	dsp, err = gvk.NewDispatcher(token, groupID, newBot)

	log.Println(err)
	log.Println(dsp.Poll())
}
