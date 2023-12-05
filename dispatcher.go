package gvk

import (
	"net/http"
	"sync"
)

type Bot interface {
	Update(*Update)
}

type NewBotFn func(chatId int64) Bot

type Dispatcher struct {
	sessionMap map[int64]Bot
	newBot     NewBotFn
	updates    chan *Update
	api        API
	httpServer *http.Server
	mu         sync.Mutex

	GroupID int64
	Server  string
	Key     string
	Ts      string
	Wait    int64
}

// NewDispatcher returns a new instance of the Dispatcher object.
// Calls the Update function of the bot associated with each chat ID.
// If a new chat ID is found, newBotFn will be called first.
func NewDispatcher(token string, groupID int64, newBotFn NewBotFn) (*Dispatcher, error) {
	d := &Dispatcher{
		api:        NewAPI(token),
		sessionMap: make(map[int64]Bot),
		newBot:     newBotFn,
		updates:    make(chan *Update),
		GroupID:    groupID,
		Wait:       25,
	}
	err := d.updateServer(true)
	if err != nil {
		return nil, err
	}
	go d.listen()
	return d, nil
}

// Poll is a wrapper function for PollOptions.
func (d *Dispatcher) Poll() error {
	return d.PollOptions(true, UpdateOptions{Timeout: 120})
}

func (d *Dispatcher) PollOptions(dropPendingUpdates bool, opts UpdateOptions) error {
	//var (
	//	timeout    = d.Wait
	//	isFirstRun = true //TODO: сброс апдейтов
	//)

	//// deletes webhook if present to run in long polling mode
	//if _, err := d.api.DeleteWebhook(dropPendingUpdates); err != nil {
	//	return err
	//}

	for {
		//TODO: сброс апдейтов
		//if isFirstRun {
		//	opts.Timeout = 0
		//}

		response, err := d.api.GetUpdates(&opts)
		if err != nil {
			return err
		}
		//
		//if !dropPendingUpdates || !isFirstRun {
		//	for _, u := range response.Result {
		//		d.updates <- u
		//	}
		//}
		for _, u := range response.Updates {
			d.updates <- u
		}
		//if l := len(response.Result); l > 0 {
		//	opts.Offset = response.Result[l-1].ID + 1
		//}
		//
		//if isFirstRun {
		//	isFirstRun = false
		//	opts.Timeout = timeout
		//}
	}
}

func (d *Dispatcher) instance(chatID int64) Bot {
	bot, ok := d.sessionMap[chatID]
	if !ok {
		bot = d.newBot(chatID)
		d.mu.Lock()
		d.sessionMap[chatID] = bot
		d.mu.Unlock()
	}
	return bot
}

func (d *Dispatcher) listen() {
	for update := range d.updates {
		bot := d.instance(update.ChatID())
		go bot.Update(update)
	}
}

func (d *Dispatcher) updateServer(updateTs bool) error {
	serverSetting, err := d.api.MessagesGetLongPollServer(&GetLongPollServerOptions{GroupID: d.GroupID})
	if err != nil {
		return err
	}

	d.Key = serverSetting.Key
	d.Server = serverSetting.Server

	if updateTs {
		d.Ts = serverSetting.Ts
	}

	return nil
}
