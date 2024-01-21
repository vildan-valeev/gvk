package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/vildan-valeev/gvk"
)

//
//const (
//	groupID      = 194299208
//	token        = "vk1.a.8-WDS15nXOxt3wU9pfCOHjAt0e7LiOZl8_u_su7PzFXVcgRgpJXZbDOB_cpBVKPOitdfTi9_Bp-oGQbrEeRr_ozjdHB3tbCtJOQFSe6VSEfX5C_IzvEUqZ3xnPlODyQTohFEh-EzhP9CcQEK5Ei6s0Xwy2L3JRSYIoup2wZFApIhpkGDTU5tdvQ09Skz7qMO3eg8fmqpM6jIzMfBIkZR6A"
//	tokenPosting = "vk1.a.F3PWupGyC4SKUFgQs51H1u7NOXcm8uWweVusWkmzWDwZha2uUzCeAItJGq4GBbwSQEnqIrHvVph6tt5xaHOc5w96q0GVKgJShvtUEIjtcEgj81CzXeaN8nJLEjrKN6ZVxpTU54Id45JOhY3sFBxg1giD45JxclmQXrk9FYxem6aRmXTpkSq-9hPnxuTyu5wb"
//)

type stateFn func(event *gvk.Update) stateFn

type Bot struct {
	chatID int64
	state  stateFn
	name   string

	gvk.API
}

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

	dsp, err := gvk.NewDispatcher(token, groupID, newBot)

	log.Println(err)
	log.Println(dsp.Poll())
}
