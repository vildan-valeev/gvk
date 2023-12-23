package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/vildan-valeev/gvk"
)

const (
	groupID = 194299208
	token   = "b30fae3f8d488e20cdbe041cbec9a0aa62e7c52e6107f97f97a9fd9007abe32223e1373cce590bfabf374"
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
	b.MessagesSend("goodbye...", &gvk.MessagesSendOptions{UserID: b.chatID})
	dsp.DelSession(b.chatID)
}

func (b *Bot) Update(update *gvk.Update) {
	b.state = b.state(update)
}

func (b *Bot) EntryHandler(update *gvk.Update) stateFn {
	if strings.HasPrefix(update.Object.MessageNew.Message.Text, "ping") {
		b.MessagesSend("pong", &gvk.MessagesSendOptions{UserID: b.chatID})
		return b.handleNext
	}

	b.MessagesSend("not understand...", &gvk.MessagesSendOptions{UserID: b.chatID})

	return b.EntryHandler
}

func (b *Bot) handleNext(update *gvk.Update) stateFn {

	b.MessagesSend("pong again )))", &gvk.MessagesSendOptions{
		UserID: b.chatID,
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
