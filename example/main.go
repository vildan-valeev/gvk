package main

import (
	"fmt"
	"github.com/vildan-valeev/gvk"
	"log"
	"strings"
)

// Recursive type definition of the bot state function.
type stateFn func(event *gvk.Update) stateFn

type bot struct {
	chatID int64
	state  stateFn
	name   string

	gvk.API
}

const (
	groupID = 2354
	token   = "b30fae3f8d488e20cdbe041cbec9a0aa62e7c52e6107f97f97a9fd9007abe32223e1373cce590bfabf374"
)

func newBot(chatID int64) gvk.Bot {
	b := &bot{
		chatID: chatID,
		API:    gvk.NewAPI(token),
	}

	b.state = b.EntryHandler
	return b
}

func (b *bot) Update(update *gvk.Update) {
	b.state = b.state(update)
}

func (b *bot) EntryHandler(update *gvk.Update) stateFn {
	if strings.HasPrefix(update.MessageNew.Text, "ping") {
		b.MessagesSend("pong", update.MessageNew.FromID, &gvk.MessagesSendOptions{UserID: update.MessageNew.FromID, PeerID: 0})

		return b.handleNext
	}

	return b.EntryHandler
}

func (b *bot) handleNext(update *gvk.Update) stateFn {
	b.name = update.MessageNew.Text
	b.MessagesSend("pong again)))", update.MessageNew.FromID, &gvk.MessagesSendOptions{UserID: update.MessageNew.FromID, PeerID: 0})

	return b.EntryHandler
}

func main() {
	fmt.Print("Start!")

	dsp, err := gvk.NewDispatcher(token, groupID, newBot)

	log.Println(err)
	log.Println(dsp.Poll())
}
