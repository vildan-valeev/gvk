package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/vildan-valeev/gvk"
)

//const (
//	groupID = 194299208
//	token   = "b30fae3f8d488e20cdbe041cbec9a0aa62e7c52e6107f97f97a9fd9007abe32223e1373cce590bfabf374"
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
	fmt.Println("INCOME TEXT", update.Object.MessageNew.Message.Text)
	if strings.HasPrefix(update.Object.MessageNew.Message.Text, "ping") {
		buttons := make([][]gvk.Button, 0)
		buttons = append(buttons, []gvk.Button{
			{
				Color: gvk.ButtonColorPrimary,
				Action: gvk.ButtonAction{
					Type:  gvk.ButtonTypeText,
					Label: "Текст кнопки",
				},
			},
		})

		b.MessagesSend(&gvk.MessagesSendOptions{
			Message: "pong",
			UserID:  b.chatID,
			Keyboard: gvk.Keyboard{
				Inline:  false,
				Buttons: buttons,
			},
		})
		return b.handleNext
	}

	b.MessagesSend(&gvk.MessagesSendOptions{Message: "not understand...", UserID: b.chatID})

	return b.EntryHandler
}

func (b *Bot) handleNext(update *gvk.Update) stateFn {

	buttons := make([][]gvk.Button, 0)
	buttons = append(buttons, []gvk.Button{
		{
			Color: gvk.ButtonColorPrimary,
			Action: gvk.ButtonAction{
				Type:    gvk.ButtonTypeCallBack,
				Label:   "Жми колбек",
				Payload: `{"data":"loading"}`,
			},
		},
	})

	b.MessagesSend(&gvk.MessagesSendOptions{
		Message: "pong again )))",
		UserID:  b.chatID,
		Keyboard: gvk.Keyboard{
			Inline:  true,
			Buttons: buttons,
		},
	})

	return b.hanleCallback
}

func (b *Bot) hanleCallback(update *gvk.Update) stateFn {
	s := string(update.Object.MessageEvent.Payload)
	b.MessagesSend(&gvk.MessagesSendOptions{
		Message: fmt.Sprintf("Callback received %s", s),
		UserID:  b.chatID,
	})

	return b.EntryHandler
}

func main() {
	fmt.Print("Start with Keyboards! \n")

	dsp, err := gvk.NewDispatcher(token, groupID, newBot)

	log.Println(err)
	log.Println(dsp.Poll())
}
