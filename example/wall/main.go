package main

import (
	"fmt"
	"github.com/vildan-valeev/gvk"
	"log"
	"strings"
)

const (
	groupID      = 194299208
	token        = "vk1.a.8-WDS15nXOxt3wU9pfCOHjAt0e7LiOZl8_u_su7PzFXVcgRgpJXZbDOB_cpBVKPOitdfTi9_Bp-oGQbrEeRr_ozjdHB3tbCtJOQFSe6VSEfX5C_IzvEUqZ3xnPlODyQTohFEh-EzhP9CcQEK5Ei6s0Xwy2L3JRSYIoup2wZFApIhpkGDTU5tdvQ09Skz7qMO3eg8fmqpM6jIzMfBIkZR6A"
	tokenPosting = "vk1.a.aeQLjRCZN_9B7QfHcz7XJYeRuVr6ty6qgYP6QJ7k-qURVARGmeAFK-VEoIyoK1Qit6sRUxEcC71w7fxZJeUTKbyxY2MffdW8wSN4zU6zT1I31GeUM081O59pa8n_aVwx86QWXVYbP81TiEGd9GIoJ5V3mxyPgJDuGHT0FBlkhvjr1ZJjbGZ7EjJQPdZZ7BSQ"
)

/*
https://api.vk.com/method/wall.post?access_token=vk1.a.F3PWupGyC4SKUFgQs51H1u7NOXcm8uWweVusWkmzWDwZha2uUzCeAItJGq4GBbwSQEnqIrHvVph6tt5xaHOc5w96q0GVKgJShvtUEIjtcEgj81CzXeaN8nJLEjrKN6ZVxpTU54Id45JOhY3sFBxg1giD45JxclmQXrk9FYxem6aRmXTpkSq-9hPnxuTyu5wb&v=5.131&message=pong&from_group=1&owner_id=-194299208

https://oauth.vk.com/authorize?client_id=7306727&display=page&redirect_uri=https://oauth.vk.com/blank.html&scope=friends&response_type=token&v=5.131

*/

type stateFn func(event *gvk.Update) stateFn

type Bot struct {
	chatID int64
	state  stateFn
	name   string

	postID int64
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
		poster := gvk.NewAPI(tokenPosting)
		opt := gvk.WallPostOptions{
			Message:   "pong",
			OwnerID:   -groupID,
			FromGroup: 1,
		}
		post, err := poster.WallPost(&opt)
		if err != nil {
			log.Fatalln(err)
		}
		b.postID = post.Response.PostID
		b.MessagesSend(&gvk.MessagesSendOptions{Message: "pong to channel...", UserID: b.chatID})
		return b.handleNext
	}

	b.MessagesSend(&gvk.MessagesSendOptions{Message: "not understand...", UserID: b.chatID})

	return b.EntryHandler
}

func (b *Bot) handleNext(update *gvk.Update) stateFn {
	poster := gvk.NewAPI(tokenPosting)
	opt := gvk.WallEditOptions{
		Message: "update pong",
		OwnerID: -groupID,
	}
	_, err := poster.WallEdit(&opt)
	log.Println(b.postID, err)
	b.MessagesSend(&gvk.MessagesSendOptions{
		Message: "pong again )))",
		UserID:  b.chatID,
	})

	return b.EntryHandler
}

func main() {
	fmt.Println("Start!")

	dsp, err := gvk.NewDispatcher(token, groupID, newBot)

	log.Println(err)
	log.Println(dsp.Poll())
}
