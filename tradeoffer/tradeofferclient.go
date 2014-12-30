package tradeoffer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/smithfox/gosteam/bot"
	"github.com/smithfox/gosteam/community"
	"github.com/smithfox/gosteam/economy/inventory"
	"github.com/smithfox/gosteam/netutil"
	. "github.com/smithfox/gosteam/steamid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const apiUrl = "http://api.steampowered.com/IEconService/%s/v%d"

type TradeOfferClient struct {
	httpClient *http.Client
	b          *bot.BotRunTime
}

func NewTradeOfferClient(b *bot.BotRunTime) *TradeOfferClient {
	c := &TradeOfferClient{
		httpClient: new(http.Client),
		b:          b,
	}

	community.SetCookies(c.httpClient, b.WebSessionId(), b.WebSteamLogin())

	community.SetCookiesHttps(c.httpClient, b.WebSessionId(), b.WebSteamLogin(), b.WebSteamLoginSecure())
	return c
}

func (c *TradeOfferClient) SendOffer(partner SteamId, offer_url_token string, message string, obj *TradeOfferSendObj) (TradeOfferId, error) {
	json_tradeoffer, err := json.Marshal(obj)
	if err != nil {
		return 0, err
	}

	uuuu, _ := url.Parse("https://steamcommunity.com/")
	send_url := "https://steamcommunity.com/tradeoffer/new/send"
	refer_url := fmt.Sprintf("https://steamcommunity.com/tradeoffer/new/?partner=%d&token=%s", partner.GetAccountId(), offer_url_token)
	c.b.Debugf("refer_url=%s,cookie=%v\n", refer_url, c.httpClient.Jar.Cookies(uuuu))
	//fmt.Printf("refer_url=%s,cookie=%v\n", refer_url, c.httpClient.Jar.Cookies(uuuu))
	resp, err := c.httpClient.Do(netutil.NewPostForm1(send_url, refer_url, netutil.ToUrlValues(map[string]string{
		"sessionid":                 c.b.WebSessionId(),
		"serverid":                  "1",
		"partner":                   partner.ToString(),
		"tradeoffermessage":         message,
		"json_tradeoffer":           string(json_tradeoffer),
		"trade_offer_create_params": "{}",
	})))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	ss, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		c.b.Debugf("sendoffer response=%s\n, http status code=%d\n", ss, resp.StatusCode)
		return 0, errors.New("sendoffer error: status code not 200")
	}

	t := new(struct {
		Tradeofferid TradeOfferId
	})
	err = json.Unmarshal(ss, t)
	if err != nil {
		fmt.Printf("ss=%s\n", string(ss))
		return 0, err
	}
	return t.Tradeofferid, nil
}

func (c *TradeOfferClient) GetOffers() (*TradeOffers, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf(apiUrl, "GetTradeOffers", 1) + "?" + netutil.ToUrlValues(map[string]string{
		"key":                 string(c.b.ApiKey),
		"get_sent_offers":     "1",
		"get_received_offers": "1",
	}).Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	t := new(struct {
		Response *TradeOffers
	})

	ss, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(ss, t)
	//err = json.NewDecoder(ss).Decode(t)
	if err != nil {
		c.b.Debugf("ss=%s\n", string(ss))
		return nil, err
	}
	return t.Response, nil
}

type actionResult struct {
	Success bool
	Error   string
}

func (c *TradeOfferClient) action(method string, version uint, id TradeOfferId) error {
	resp, err := c.httpClient.Do(netutil.NewPostForm(fmt.Sprintf(apiUrl, method, version), netutil.ToUrlValues(map[string]string{
		"key":          string(c.b.ApiKey),
		"tradeofferid": strconv.FormatUint(uint64(id), 10),
	})))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New(method + " error: status code not 200")
	}
	return nil
}

func (c *TradeOfferClient) Decline(id TradeOfferId) error {
	return c.action("DeclineTradeOffer", 1, id)
}

func (c *TradeOfferClient) Cancel(id TradeOfferId) error {
	return c.action("CancelTradeOffer", 1, id)
}

func (c *TradeOfferClient) GetOwnInventory(contextId uint64, appId uint32) (*inventory.Inventory, error) {
	return inventory.GetOwnInventory(c.httpClient, contextId, appId)
}

func (c *TradeOfferClient) GetTheirInventory(other SteamId, contextId uint64, appId uint32) (*inventory.Inventory, error) {
	return inventory.GetFullInventory(func() (*inventory.PartialInventory, error) {
		return c.getPartialTheirInventory(other, contextId, appId, nil)
	}, func(start uint) (*inventory.PartialInventory, error) {
		return c.getPartialTheirInventory(other, contextId, appId, &start)
	})
}

func (c *TradeOfferClient) getPartialTheirInventory(other SteamId, contextId uint64, appId uint32, start *uint) (*inventory.PartialInventory, error) {
	data := map[string]string{
		"sessionid": c.b.WebSessionId(),
		"partner":   other.ToString(),
		"contextid": strconv.FormatUint(contextId, 10),
		"appid":     strconv.FormatUint(uint64(appId), 10),
	}
	if start != nil {
		data["start"] = strconv.FormatUint(uint64(*start), 10)
	}

	const baseUrl = "http://steamcommunity.com/tradeoffer/new/"
	req, err := http.NewRequest("GET", baseUrl+"partnerinventory/?"+netutil.ToUrlValues(data).Encode(), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Referer", baseUrl+"?partner="+fmt.Sprintf("%d", other))

	return inventory.DoInventoryRequest(c.httpClient, req)
}
