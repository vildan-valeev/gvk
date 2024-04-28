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
	if update.Object.MessageNew != nil {
		if strings.HasPrefix(update.Object.MessageNew.Message.Text, "ping") {
			b.MessagesSend(&gvk.MessagesSendOptions{Message: "pong to channel...", UserID: b.chatID})

			poster := gvk.NewAPI(token) // https://oauth.vk.com/authorize?client_id=APP_ID&display=page&redirect_uri=https://oauth.vk.com/blank.html&scope=wall,offline&response_type=token&v=5.199
			opt := gvk.WallPostOptions{
				Message:   "pong 1",
				OwnerID:   -groupID,
				FromGroup: 1,
			}

			post, err := poster.WallPost(&opt)
			log.Println(post.Response.PostID)
			if err != nil {
				log.Println(err)
			}

			return b.handleNext
		}

		b.MessagesSend(&gvk.MessagesSendOptions{Message: "not understand...", UserID: b.chatID})
	}

	return b.EntryHandler
}

func (b *Bot) handleNext(update *gvk.Update) stateFn {

	poster := gvk.NewAPI(token)
	opt := gvk.WallPostOptions{
		Message:   "pong 2",
		OwnerID:   -groupID,
		FromGroup: 1,
	}

	post, err := poster.WallPost(&opt)

	if err != nil {
		log.Println(err)
	}
	log.Println(post.Response.PostID)

	b.MessagesSend(&gvk.MessagesSendOptions{
		Message: "pong again )))",
		UserID:  b.chatID,
	})

	return b.EntryHandler
}

func main() {
	fmt.Println("Start wall!")

	dsp, err := gvk.NewDispatcher(token, groupID, newBot)

	log.Println(err)
	log.Println(dsp.Poll())
}

/*
https://api.vk.com/method/wall.edit?access_token=token&message=Откликнуться&owner_id=-194299208&post_id=206&v=5.199
*/
