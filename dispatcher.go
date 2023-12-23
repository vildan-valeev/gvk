package gvk

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
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
	opts    UpdateOptions
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
	}
	err := d.updateServer(true)
	if err != nil {
		return nil, err
	}

	go d.listen()
	return d, nil
}

// DelSession deletes the Bot instance, seen as a session, from the
// map with all of them.
func (d *Dispatcher) DelSession(chatID int64) {
	d.mu.Lock()
	delete(d.sessionMap, chatID)
	d.mu.Unlock()
}

// AddSession allows to arbitrarily create a new Bot instance.
func (d *Dispatcher) AddSession(chatID int64) {
	d.mu.Lock()
	if _, isIn := d.sessionMap[chatID]; !isIn {
		d.sessionMap[chatID] = d.newBot(chatID)
	}
	d.mu.Unlock()
}

// Poll is a wrapper function for PollOptions.
func (d *Dispatcher) Poll() error {

	return d.PollOptions(true)
}

func (d *Dispatcher) PollOptions(dropPendingUpdates bool) error {
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
		time.Sleep(2 * time.Second)
		result, err := d.api.GetUpdates(&d.opts)
		if err != nil {
			return err
		}

		updates, err := result.Updates.UnmarshalCustom()
		if err != nil {
			return err
		}

		err = d.check(result)
		if err != nil {
			return err
		}
		//
		//if !dropPendingUpdates || !isFirstRun {
		//	for _, u := range response.Result {
		//		d.updates <- u
		//	}
		//}

		for _, u := range updates {
			fmt.Println("UPDATES!", u.Object.MessageNew)
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
		id := update.ChatID()
		// если ошибка и chat_ID не пришел
		if id == 0 {
			continue
		}

		bot := d.instance(id)
		go bot.Update(update)
	}
}

func (d *Dispatcher) updateServer(updateTs bool) error {
	serverSetting, err := d.api.GroupsGetLongPollServer(&GetLongPollServerOptions{GroupID: d.GroupID})
	if err != nil {
		return err
	}
	//TODO: повесить мютексы
	d.opts.Key = serverSetting.Response.Key
	d.opts.Server = serverSetting.Response.Server

	if updateTs {
		d.opts.Ts = serverSetting.Response.Ts
	}

	return nil
}

func (d *Dispatcher) autoSetting(ctx context.Context) error {
	// Updating LongPoll settings
	opts := SetLongPollSettingsOptions{
		GroupID:    d.GroupID,
		Enable:     1,
		APIVersion: APIVersion,
	}
	_, err := d.api.GroupsSetLongPollSettings(&opts)
	log.Println("SetLongPollSettings", err)
	return err
}

func (d *Dispatcher) check(r APIResponseUpdate) (err error) {
	switch r.Failed {
	case 0:
		d.opts.Ts = r.Ts
	case 1:
		d.opts.Ts = r.Ts
	case 2:
		err = d.updateServer(false)
	case 3:
		err = d.updateServer(true)
	default:
		log.Println("Dispatcher update check", err)
		err = &Failed{Code: r.Failed}
	}
	return nil
}

// Failed struct.
type Failed struct {
	Code int64
}

// Error returns the message of a Failed.
func (e Failed) Error() string {
	return fmt.Sprintf(
		"longpoll: failed code %d",
		e.Code,
	)
}
