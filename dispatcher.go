package gvk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SevereCloud/vksdk/v2"
	"github.com/SevereCloud/vksdk/v2/api"
	poll "github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Bot interface {
	Update(*GroupEvent)
}

type NewBotFn func(chatId int) Bot

type LongPoll struct {
	sessionMap map[int]Bot
	newBot     NewBotFn
	updates    chan *GroupEvent
	VK         *api.VK
	Client     *http.Client
	mu         sync.Mutex

	GroupID int
	Server  string
	Key     string
	Ts      string
	Wait    int
}

// NewDispatcher returns a new instance of the Dispatcher object.
// Calls the Update function of the bot associated with each chat ID.
// If a new chat ID is found, newBotFn will be called first.
func NewLongPoll(token string, groupID int, newBotFn NewBotFn) *LongPoll {
	d := &LongPoll{
		VK:         api.NewVK(token),
		GroupID:    groupID,
		sessionMap: make(map[int]Bot),
		newBot:     newBotFn,
		updates:    make(chan *GroupEvent),
		Client:     http.DefaultClient,
		Wait:       25,
	}

	err := d.updateServer(true)
	if err != nil {
		log.Fatal(err)
	}

	go d.listen()

	return d
}

func (lp *LongPoll) updateServer(updateTs bool) error {
	params := api.Params{
		"group_id": lp.GroupID,
	}

	serverSetting, err := lp.VK.GroupsGetLongPollServer(params)
	if err != nil {
		return err
	}

	lp.Key = serverSetting.Key
	lp.Server = serverSetting.Server

	if updateTs {
		lp.Ts = serverSetting.Ts
	}

	return nil
}

type Response struct {
	Ts      string       `json:"ts"`
	Updates []GroupEvent `json:"updates"`
	Failed  int          `json:"failed"`
}

func (lp *LongPoll) check(ctx context.Context) (response Response, err error) {
	u := fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=%d", lp.Server, lp.Key, lp.Ts, lp.Wait)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return response, err
	}

	resp, err := lp.Client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	response, err = parseResponse(resp.Body)
	if err != nil {
		return response, err
	}

	err = lp.checkResponse(response)

	return response, err
}

func parseResponse(reader io.Reader) (response Response, err error) {
	decoder := json.NewDecoder(reader)
	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return response, err
		}

		t, ok := token.(string)
		if !ok {
			continue
		}

		switch t {
		case "failed":
			raw, err := decoder.Token()
			if err != nil {
				return response, err
			}

			response.Failed = int(raw.(float64))
		case "updates":
			var updates []GroupEvent

			err = decoder.Decode(&updates)
			if err != nil {
				return response, err
			}

			response.Updates = updates
		case "ts":
			// can be a number in the response with "failed" field: {"ts":8,"failed":1}
			// or string, e.g. {"ts":"8","updates":[]}
			rawTs, err := decoder.Token()
			if err != nil {
				return response, err
			}

			if ts, isNumber := rawTs.(float64); isNumber {
				response.Ts = strconv.Itoa(int(ts))
			} else {
				response.Ts = rawTs.(string)
			}
		}
	}

	return response, err
}

func (lp *LongPoll) checkResponse(response Response) (err error) {
	switch response.Failed {
	case 0:
		lp.Ts = response.Ts
	case 1:
		lp.Ts = response.Ts
	case 2:
		err = lp.updateServer(false)
	case 3:
		err = lp.updateServer(true)
	default:
		err = &poll.Failed{response.Failed}
	}

	return
}

func (lp *LongPoll) autoSetting(ctx context.Context) error {
	params := api.Params{
		"group_id":    lp.GroupID,
		"enabled":     true,
		"api_version": vksdk.API,
	}.WithContext(ctx)
	//for _, event := range lp.ListEvents() {
	//	params[string(event)] = true
	//}

	// Updating LongPoll settings
	_, err := lp.VK.GroupsSetLongPollSettings(params)

	return err
}

// Run handler.
func (lp *LongPoll) Run() error {
	return lp.RunWithContext(context.Background())
}

// RunWithContext handler.
func (lp *LongPoll) RunWithContext(ctx context.Context) error {
	return lp.run(ctx)
}

func (lp *LongPoll) run(ctx context.Context) error {
	//ctx, lp.cancel = context.WithCancel(ctx)

	if err := lp.autoSetting(ctx); err != nil {
		return err
	}

	for {
		select {
		case _, ok := <-ctx.Done():
			if !ok {
				return nil
			}
		default:
			resp, err := lp.check(ctx)
			if err != nil {
				return err
			}

			//ctx = context.WithValue(ctx, internal.LongPollTsKey, resp.Ts)

			for _, u := range resp.Updates {

				lp.updates <- &u
			}

		}
	}
}

func (lp *LongPoll) instance(chatID int) Bot {
	bot, ok := lp.sessionMap[chatID]
	if !ok {
		bot = lp.newBot(chatID)
		lp.mu.Lock()
		lp.sessionMap[chatID] = bot
		lp.mu.Unlock()
	}
	return bot
}

func (lp *LongPoll) listen() {
	for update := range lp.updates {
		bot := lp.instance(update.ParsedObj.Message.PeerID)
		go bot.Update(update)
	}
}

//// Shutdown gracefully shuts down the longpoll without interrupting any active connections.
//func (lp *LongPoll) Shutdown() {
//	if lp.cancel != nil {
//		lp.cancel()
//	}
//}

//// FullResponse handler.
//func (lp *LongPoll) FullResponse(f func(Response)) {
//	lp.funcFullResponseList = append(lp.funcFullResponseList, f)
//}
//

//
//// DelSession deletes the Bot instance, seen as a session, from the
//// map with all of them.
//func (lp *LongPoll) DelSession(chatID int64) {
//	d.mu.Lock()
//	delete(d.sessionMap, chatID)
//	d.mu.Unlock()
//}
//
//// AddSession allows to arbitrarily create a new Bot instance.
//func (lp *LongPoll) AddSession(chatID int64) {
//	d.mu.Lock()
//	if _, isIn := d.sessionMap[chatID]; !isIn {
//		d.sessionMap[chatID] = d.newBot(chatID)
//	}
//	d.mu.Unlock()
//}
