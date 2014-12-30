package bot

import (
	"encoding/json"
	"fmt"
	. "github.com/smithfox/gosteam/apath"
	. "github.com/smithfox/gosteam/internal"
	"github.com/smithfox/gosteam/log"
	. "github.com/smithfox/gosteam/steamid"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync/atomic"
)

/*
{
    "Username":"dddddd",
    "Password":"xxxxx",
    "DisplayName":"ssss",
    "LogLevel":"Debug",
    "AutoStart": "true"
},
*/

type APIKey string

//compatible c# steambot config file
type BotConf struct {
	Username    string
	Password    string
	DisplayName string
	LogLevel    string //DEBUG,INFO,WARN,ERROR
	ApiKey      APIKey //steam can recognize bot with apikey
}

type BotsConf struct {
	Bots []*BotConf
}

var gBotsByDisplayName map[string]*BotConf = make(map[string]*BotConf)
var gBotsByUsername map[string]*BotConf = make(map[string]*BotConf)

func LoadBotConfs() error {
	json_conf_file_path := filepath.Join(AppPath(), "bots", "bots_conf.json")
	var botsconf BotsConf
	bb, err := ioutil.ReadFile(json_conf_file_path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bb, &botsconf)
	if err != nil {
		return err
	}

	if len(botsconf.Bots) < 1 {
		return fmt.Errorf("Not find any botconf from configfile=%s\n", json_conf_file_path)
	}

	for _, botconf := range botsconf.Bots {
		botconf.DisplayName = strings.ToLower(botconf.DisplayName)

		if _, ok := gBotsByDisplayName[botconf.DisplayName]; ok {
			return fmt.Errorf("Duplicate bot conf with displayname=%s\n", botconf.DisplayName)
		}
		if _, ok := gBotsByUsername[botconf.Username]; ok {
			return fmt.Errorf("Duplicate bot conf with username=%s\n", botconf.Username)
		}
		gBotsByDisplayName[botconf.DisplayName] = botconf
		gBotsByUsername[botconf.Username] = botconf
	}

	return nil
}

func GetBotConf(bot_display_name string) *BotConf {
	bot_display_name = strings.ToLower(bot_display_name)
	botconf, _ := gBotsByDisplayName[bot_display_name]
	return botconf
}

type EventBus struct {
	events chan interface{}
}

func newEventBus() *EventBus {
	return &EventBus{
		events: make(chan interface{}, 3),
	}
}

func (c *EventBus) Events() <-chan interface{} {
	return c.events
}

func (c *EventBus) Emit(event interface{}) {
	c.events <- event
}

// Emits an error formatted with fmt.Errorf.
func (c *EventBus) EmitErrorf(format string, a ...interface{}) {
	c.Emit(fmt.Errorf(format, a...))
}

type SentryHash []byte

type BotRunTime struct {
	// these need to be 64 bit aligned for sync/atomic on 32bit
	sessionId int32
	_         int32
	*log.Logger
	*BotConf
	*EventBus
	steamId        SteamId
	currentJobId   uint64
	AuthCode       string
	SentryFileHash SentryHash
	webSessionId   string
	// The `steamLogin` cookie required to use the steam website. Already URL-escaped.
	// It is only available after calling LogOn().
	webSteamLogin       string
	webSteamLoginSecure string
	webLoginKey         string
	// ConnectionTimeout time.Duration
}

func NewBotRunTime(botconf *BotConf, authcode string) *BotRunTime {
	bot_lower_name := strings.ToLower(botconf.DisplayName)
	log_dir := filepath.Join(AppPath(), "bots", "logs")
	log1 := log.NewLog(log_dir, bot_lower_name)
	log1.ConsoleOutput(true)

	eventbus := newEventBus()

	sentry_filepath := filepath.Join(AppPath(), "bots", "sentry", bot_lower_name+".sentry")
	sentry, err := ioutil.ReadFile(sentry_filepath)
	if err != nil {
		log1.Warnf("Fail to read sentry file: %s\n", sentry_filepath)
	}
	botruntime := &BotRunTime{
		Logger:         log1,
		BotConf:        botconf,
		EventBus:       eventbus,
		SentryFileHash: sentry,
		AuthCode:       authcode,
	}
	return botruntime
}

func (b *BotRunTime) WriteSentry(sentry SentryHash) {
	bot_lower_name := strings.ToLower(b.BotConf.DisplayName)
	sentry_filepath := filepath.Join(AppPath(), "bots", "sentry", bot_lower_name+".sentry")
	err := ioutil.WriteFile(sentry_filepath, sentry, 0666)

	if err != nil {
		b.Warnf("Fail to write sentry file: %s\n", sentry_filepath)
	}
}

func (b *BotRunTime) GetNextJobId() JobId {
	return JobId(atomic.AddUint64(&b.currentJobId, 1))
}

func (b *BotRunTime) SetSessionId(ss int32) {
	atomic.StoreInt32(&b.sessionId, ss)
}

func (b *BotRunTime) SessionId() int32 {
	return atomic.LoadInt32(&b.sessionId)
}

func (b *BotRunTime) SetSteamId(tmp_steamid SteamId) {
	b.steamId = tmp_steamid
}

func (b *BotRunTime) SteamId() SteamId {
	return b.steamId
}

/**
The `sessionid` cookie required to use the steam website.
This cookie may contain a characters that will need to be URL-escaped, otherwise
Steam (probably) interprets is as a string.
When used as an URL paramter this is automatically escaped by the Go HTTP package.
*/
func (b *BotRunTime) WebSessionId() string {
	return b.webSessionId
}

func (b *BotRunTime) SetWebSessionId(websessionid string) {
	b.webSessionId = websessionid
}

func (b *BotRunTime) WebSteamLoginSecure() string {
	return b.webSteamLoginSecure
}

func (b *BotRunTime) SetWebSteamLoginSecure(web_steam_login_secure string) {
	b.webSteamLoginSecure = web_steam_login_secure
}

func (b *BotRunTime) WebSteamLogin() string {
	return b.webSteamLogin
}

func (b *BotRunTime) SetWebSteamLogin(web_steam_login string) {
	b.webSteamLogin = web_steam_login
}

func (b *BotRunTime) WebLoginKey() string {
	return b.webLoginKey
}

func (b *BotRunTime) SetWebLoginKey(web_login_key string) {
	b.webLoginKey = web_login_key
}
